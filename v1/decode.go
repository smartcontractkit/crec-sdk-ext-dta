package v1

import (
	"context"
	"fmt"

	apiClient "github.com/smartcontractkit/crec-api-go/client"
)

// ConcreteEvent represents any decoded concrete event payload.
type ConcreteEvent interface{}

// DecodedEvent wraps WatcherEventPayload with a decoded ConcreteEvent.
type DecodedEvent struct {
	apiClient.WatcherEventPayload
	ConcreteEvent ConcreteEvent
}

// EventName returns the parsed event name from the payload.
func (e DecodedEvent) EventName() EventName {
	name, ok := parseEventName(e.Event.EventName)
	if !ok {
		return EventUnknown
	}
	return name
}

// DecodeFromEvent extracts the WatcherEventPayload from an apiClient.Event and decodes
// the ConcreteEvent based on the event type.
func DecodeFromEvent(_ context.Context, event apiClient.Event) (DecodedEvent, error) {
	payload, err := event.Payload.AsWatcherEventPayload()
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("extract watcher payload: %w", err)
	}

	// Convert Data map to string params for decoders
	params := make(map[string]string, len(payload.Event.Data))
	for k, v := range payload.Event.Data {
		params[k] = fmt.Sprintf("%v", v)
	}

	name, ok := parseEventName(payload.Event.EventName)
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event name: %s", payload.Event.EventName)
	}
	decoder, ok := eventDecoders[name]
	if !ok {
		return DecodedEvent{}, fmt.Errorf("unsupported event decoder for event name: %s", name)
	}

	concrete, err := decoder(params, payload.Transaction.Hash)
	if err != nil {
		return DecodedEvent{}, fmt.Errorf("decode %s: %w", name, err)
	}

	return DecodedEvent{
		WatcherEventPayload: payload,
		ConcreteEvent:       concrete,
	}, nil
}
