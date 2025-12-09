package types_test

import (
	"testing"

	"github.com/smartcontractkit/crec-sdk-ext-dta/types"
	"github.com/stretchr/testify/require"
)

func TestAttrs_Has(t *testing.T) {
	tests := []struct {
		name   string
		attrs  types.Attrs
		key    string
		expect bool
	}{
		{
			name:   "key exists",
			attrs:  types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:    "foo",
			expect: true,
		},
		{
			name:   "key does not exist",
			attrs:  types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:    "baz",
			expect: false,
		},
		{
			name:   "empty attrs",
			attrs:  types.Attrs{},
			key:    "foo",
			expect: false,
		},
		{
			name:   "nil attrs",
			attrs:  nil,
			key:    "foo",
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.attrs.Has(tt.key)
			require.Equal(t, tt.expect, got)
		})
	}
}

func TestAttrs_Get(t *testing.T) {
	tests := []struct {
		name        string
		attrs       types.Attrs
		key         string
		expectValue string
		expectOK    bool
	}{
		{
			name:        "key exists with value",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "foo",
			expectValue: "bar",
			expectOK:    true,
		},
		{
			name:        "key exists with empty value",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: ""}},
			key:         "foo",
			expectValue: "",
			expectOK:    true,
		},
		{
			name:        "key does not exist",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "baz",
			expectValue: "",
			expectOK:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, ok := tt.attrs.Get(tt.key)
			require.Equal(t, tt.expectOK, ok)
			require.Equal(t, tt.expectValue, value)
		})
	}
}

func TestAttrs_Require(t *testing.T) {
	tests := []struct {
		name        string
		attrs       types.Attrs
		key         string
		expectValue string
		expectErr   bool
	}{
		{
			name:        "key exists with value",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "foo",
			expectValue: "bar",
			expectErr:   false,
		},
		{
			name:        "key exists with empty value",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: ""}},
			key:         "foo",
			expectValue: "",
			expectErr:   true,
		},
		{
			name:        "key does not exist",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "baz",
			expectValue: "",
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := tt.attrs.Require(tt.key)
			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), "missing required attribute")
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectValue, value)
			}
		})
	}
}

func TestAttrs_Default(t *testing.T) {
	tests := []struct {
		name        string
		attrs       types.Attrs
		key         string
		defaultVal  string
		expectValue string
	}{
		{
			name:        "key exists with value",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "foo",
			defaultVal:  "default",
			expectValue: "bar",
		},
		{
			name:        "key exists with empty value returns default",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: ""}},
			key:         "foo",
			defaultVal:  "default",
			expectValue: "default",
		},
		{
			name:        "key does not exist returns default",
			attrs:       types.Attrs{"foo": {Key: "foo", Value: "bar"}},
			key:         "baz",
			defaultVal:  "default",
			expectValue: "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value := tt.attrs.Default(tt.key, tt.defaultVal)
			require.Equal(t, tt.expectValue, value)
		})
	}
}

// EventType is a test type for GetEventName
type EventType string

const (
	EventTypeKnown   EventType = "KnownEvent"
	EventTypeUnknown EventType = "Unknown"
)

func parseEventType(s string) (EventType, bool) {
	if s == string(EventTypeKnown) {
		return EventTypeKnown, true
	}
	return "", false
}

func TestGetEventName(t *testing.T) {
	tests := []struct {
		name   string
		event  types.VerifiableEvent
		expect EventType
	}{
		{
			name: "event type from attributes",
			event: types.VerifiableEvent{
				Metadata: types.Metadata{
					WorkflowEvent: types.WorkflowEvent{
						Attributes: types.Attrs{
							"event_type": {Value: "KnownEvent"},
						},
					},
				},
			},
			expect: EventTypeKnown,
		},
		{
			name: "event type from event name when attrs empty",
			event: types.VerifiableEvent{
				Event: types.Event{
					Name: "KnownEvent",
				},
				Metadata: types.Metadata{
					WorkflowEvent: types.WorkflowEvent{
						Attributes: types.Attrs{},
					},
				},
			},
			expect: EventTypeKnown,
		},
		{
			name: "unknown event type",
			event: types.VerifiableEvent{
				Event: types.Event{
					Name: "SomeUnknownEvent",
				},
				Metadata: types.Metadata{
					WorkflowEvent: types.WorkflowEvent{
						Attributes: types.Attrs{
							"event_type": {Value: "SomeUnknownEvent"},
						},
					},
				},
			},
			expect: EventTypeUnknown,
		},
		{
			name:   "empty event returns unknown",
			event:  types.VerifiableEvent{},
			expect: EventTypeUnknown,
		},
		{
			name: "attributes take precedence over event name",
			event: types.VerifiableEvent{
				Event: types.Event{
					Name: "SomeOtherEvent",
				},
				Metadata: types.Metadata{
					WorkflowEvent: types.WorkflowEvent{
						Attributes: types.Attrs{
							"event_type": {Value: "KnownEvent"},
						},
					},
				},
			},
			expect: EventTypeKnown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := types.GetEventName(tt.event, EventTypeUnknown, parseEventType)
			require.Equal(t, tt.expect, got)
		})
	}
}

