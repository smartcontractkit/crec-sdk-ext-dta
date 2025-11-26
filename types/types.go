// Package types provides shared types and functions for DTA event decoding.
// It contains the canonical struct definitions used by both v0 and v1 DTA event handlers.
package types

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	apiClient "github.com/smartcontractkit/crec-api-go/client"
)

// ConcreteEvent represents any decoded concrete event payload.
type ConcreteEvent interface{}

// VerifiableEvent represents an event structure that encapsulates data about the event,
// its metadata, and associated blockchain transaction details.
type VerifiableEvent struct {
	CreatedAt   time.Time         `json:"created_at"`
	Event       Event             `json:"event"`
	Metadata    Metadata          `json:"metadata"`
	Parameters  map[string]string `json:"parameters"`
	Transaction Transaction       `json:"transaction"`

	// ConcreteEvent holds the decoded concrete event based on the event name
	// and using the fields in VerifiableEvent.Metadata.WorkflowEvent.Attributes
	ConcreteEvent ConcreteEvent `json:"-"`
}

// Event contains the core event data from the blockchain.
type Event struct {
	Name        string            `json:"name"`
	Address     string            `json:"address"`
	Service     string            `json:"service"`
	LogIndex    int               `json:"log_index"`
	Parameters  map[string]string `json:"parameters"`
	TopicHash   string            `json:"topic_hash"`
	BlockNumber int               `json:"block_number"`
}

// Metadata contains chain and workflow information for an event.
type Metadata struct {
	ChainSelector string        `json:"chain_selector"`
	Network       string        `json:"network"`
	WorkflowEvent WorkflowEvent `json:"workflowEvent"`
}

// WorkflowEvent contains workflow-specific event attributes.
type WorkflowEvent struct {
	Component      string   `json:"component"`
	Attributes     Attrs    `json:"attributes"`
	ProcessLabels  []string `json:"process_labels"`
	EventTypeLabel string   `json:"event_type_label"`
}

// Transaction contains blockchain transaction details.
type Transaction struct {
	Hash          string `json:"hash"`
	ChainSelector string `json:"chain_selector"`
	Timestamp     int    `json:"timestamp"`
	BlockNumber   int    `json:"block_number"`
}

// Attribute represents a single event attribute with metadata.
type Attribute struct {
	Key        string `json:"key"`
	OnChain    bool   `json:"on_chain"`
	Value      string `json:"value"`
	Visibility string `json:"visibility"`
}

// Attrs is a map of attribute names to Attribute values.
type Attrs map[string]Attribute

// Has checks if the specified key exists in the Attrs map.
// Returns true if the key is present; otherwise, false.
func (a Attrs) Has(key string) bool {
	_, ok := a[key]
	return ok
}

// Get retrieves the value and existence status of the specified key from the Attrs map.
// Returns the value and true if key exists, otherwise an empty string and false.
func (a Attrs) Get(key string) (string, bool) {
	v, ok := a[key]
	return v.Value, ok
}

// Require retrieves the value of the specified key from the Attrs map.
// Returns an error if the key is missing or its value is empty.
func (a Attrs) Require(key string) (string, error) {
	if v, ok := a.Get(key); ok && v != "" {
		return v, nil
	}
	return "", fmt.Errorf("missing required attribute %q", key)
}

// Default returns the value associated with the specified key if it exists and is non-empty;
// otherwise, it returns the provided default value.
func (a Attrs) Default(key, def string) string {
	if v, ok := a.Get(key); ok && v != "" {
		return v
	}
	return def
}

// UnmarshalFunc is a function that unmarshals JSON bytes into a VerifiableEvent
// with version-specific ConcreteEvent population.
type UnmarshalFunc func(data []byte, v *VerifiableEvent) error

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and converts it
// to a VerifiableEvent. The provided unmarshal function is used to populate the ConcreteEvent
// field with version-specific event types.
func DecodeFromEvent(ctx context.Context, event apiClient.Event, unmarshal UnmarshalFunc) (VerifiableEvent, error) {
	// Extract the WatcherEventPayload from the Event.Payload union
	watcherPayload, err := event.Payload.AsWatcherEventPayload()
	if err != nil {
		return VerifiableEvent{}, fmt.Errorf("failed to extract watcher event payload: %w", err)
	}

	// Convert the data map to string parameters
	params := make(map[string]string)
	for k, v := range watcherPayload.Event.Data {
		params[k] = fmt.Sprintf("%v", v)
	}

	// Build attributes from the data
	attrs := make(Attrs)
	for k, v := range watcherPayload.Event.Data {
		attrs[k] = Attribute{
			Key:   k,
			Value: fmt.Sprintf("%v", v),
		}
	}
	// Add metadata attributes
	if watcherPayload.Event.Metadata != nil {
		for k, v := range *watcherPayload.Event.Metadata {
			attrs[k] = Attribute{
				Key:   k,
				Value: fmt.Sprintf("%v", v),
			}
		}
	}

	// Add event_type attribute for event name resolution
	attrs["event_type"] = Attribute{
		Key:   "event_type",
		Value: watcherPayload.Event.EventName,
	}

	// Build the VerifiableEvent structure
	ve := VerifiableEvent{
		CreatedAt: watcherPayload.Event.Timestamp,
		Event: Event{
			Name:       watcherPayload.Event.EventName,
			Address:    watcherPayload.Address,
			LogIndex:   watcherPayload.Event.LogIndex,
			Parameters: params,
			TopicHash:  watcherPayload.Event.TopicHash,
		},
		Metadata: Metadata{
			// TODO: Previously we were given ChainId, DTAO will need to handle this change in the future.
			ChainSelector: watcherPayload.ChainSelector,
			WorkflowEvent: WorkflowEvent{
				Attributes: attrs,
			},
		},
		Parameters: params,
		Transaction: Transaction{
			Hash:          watcherPayload.Transaction.Hash,
			// TODO: Previously we were given ChainId, DTAO will need to handle this change in the future.
			ChainSelector: watcherPayload.ChainSelector,
			Timestamp:     int(watcherPayload.Transaction.Timestamp),
		},
	}

	// Use JSON marshal/unmarshal to trigger version-specific ConcreteEvent population
	jsonBytes, err := json.Marshal(ve)
	if err != nil {
		return VerifiableEvent{}, fmt.Errorf("failed to marshal verifiable event: %w", err)
	}

	var result VerifiableEvent
	if err := unmarshal(jsonBytes, &result); err != nil {
		return VerifiableEvent{}, fmt.Errorf("failed to unmarshal verifiable event: %w", err)
	}

	return result, nil
}

// GetEventName returns the event name from the workflow attributes or outer event name.
// The parseFunc is used to validate and convert the event name string.
func GetEventName[T ~string](v VerifiableEvent, unknown T, parseFunc func(string) (T, bool)) T {
	var name T
	if attr, ok := v.Metadata.WorkflowEvent.Attributes["event_type"]; ok {
		if ev, ok := parseFunc(attr.Value); ok {
			name = ev
		}
	}
	if name == "" && v.Event.Name != "" {
		if ev, ok := parseFunc(v.Event.Name); ok {
			name = ev
		}
	}
	if name == "" {
		return unknown
	}
	return name
}

