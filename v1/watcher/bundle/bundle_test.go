package bundle_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/crec-sdk-ext-dta/v1/watcher/bundle"
)

func TestBundle_Get(t *testing.T) {
	b := bundle.Get()
	require.NotNil(t, b)
	assert.Equal(t, "dta.v1", b.Service)
	assert.NotEmpty(t, b.Contracts, "bundle should have contracts")
	assert.NotEmpty(t, b.Events, "bundle should have events")
}

func TestBundle_EventsSetup(t *testing.T) {
	b := bundle.Get()

	contractNames := make(map[string]bool)
	for _, c := range b.Contracts {
		contractNames[c.Name] = true
	}

	for _, evt := range b.Events {
		t.Run(evt.Name, func(t *testing.T) {
			assert.NotEmpty(t, evt.Name, "event Name must not be empty")
			assert.NotEmpty(t, evt.TriggerContract, "event TriggerContract must not be empty")
			assert.NotEmpty(t, evt.Description, "event Description must not be empty")

			// TriggerContract must reference a real contract in the bundle
			assert.True(t, contractNames[evt.TriggerContract],
				"TriggerContract %q is not in the bundle's Contracts list", evt.TriggerContract)

			// ParamsSchema must be present and valid JSON
			require.NotNil(t, evt.ParamsSchema,
				"ParamsSchema must not be nil — check that ParamsSchemas[%q] exists in params_schema_gen.go", evt.Name)
			require.True(t, len(evt.ParamsSchema) > 0,
				"ParamsSchema must not be empty — check that ParamsSchemas[%q] exists in params_schema_gen.go", evt.Name)

			var params map[string]interface{}
			err := json.Unmarshal(evt.ParamsSchema, &params)
			require.NoError(t, err, "ParamsSchema must be valid JSON")
			assert.Equal(t, "object", params["type"], "ParamsSchema should be type:object")
			assert.NotNil(t, params["properties"], "ParamsSchema should have properties")

			props, ok := params["properties"].(map[string]interface{})
			require.True(t, ok, "properties should be an object")
			assert.NotEmpty(t, props, "ParamsSchema properties should not be empty")

			// DataSchema, if present, must be valid JSON
			if evt.DataSchema != nil {
				var data map[string]interface{}
				err := json.Unmarshal(evt.DataSchema, &data)
				require.NoError(t, err, "DataSchema must be valid JSON")
				assert.Equal(t, "object", data["type"], "DataSchema should be type:object")
			}
		})
	}
}

func TestBundle_ParamsSchemaCoverage(t *testing.T) {
	b := bundle.Get()

	// Every event must have a corresponding entry in the generated ParamsSchemas map
	for _, evt := range b.Events {
		t.Run(evt.Name, func(t *testing.T) {
			schema, exists := bundle.ParamsSchemas[evt.Name]
			require.True(t, exists,
				"ParamsSchemas map is missing entry for %q — add it to bundleEvents in gen/config.go and re-run the generator", evt.Name)
			assert.NotEmpty(t, schema,
				"ParamsSchemas[%q] is empty — check the ABI has this event", evt.Name)
		})
	}

	// Every generated ParamsSchema should be referenced by at least one event
	eventNames := make(map[string]bool)
	for _, evt := range b.Events {
		eventNames[evt.Name] = true
	}
	for name := range bundle.ParamsSchemas {
		t.Run("generated/"+name, func(t *testing.T) {
			assert.True(t, eventNames[name],
				"ParamsSchemas has entry %q but no matching event in the events list — either add the event or remove it from bundleEvents in gen/config.go", name)
		})
	}
}

func TestBundle_NoDuplicateEvents(t *testing.T) {
	b := bundle.Get()
	seen := make(map[string]bool)
	for _, evt := range b.Events {
		assert.False(t, seen[evt.Name], "duplicate event %q in events list", evt.Name)
		seen[evt.Name] = true
	}
}

func TestBundle_DataSchemaConsistency(t *testing.T) {
	b := bundle.Get()

	// Events with no enrichment should have nil DataSchema
	noEnrichment := map[string]bool{
		"DistributorRegistered": true,
		"FundAdminRegistered":   true,
	}

	// Settlement events should use the settlement schema
	settlementEvents := map[string]bool{
		"DTASettlementOpened": true,
		"DTASettlementClosed": true,
	}

	for _, evt := range b.Events {
		t.Run(evt.Name, func(t *testing.T) {
			if noEnrichment[evt.Name] {
				assert.Nil(t, evt.DataSchema, "event %q should have no DataSchema (no enrichment)", evt.Name)
				return
			}

			require.NotNil(t, evt.DataSchema, "event %q should have a DataSchema (it has enrichment)", evt.Name)

			if settlementEvents[evt.Name] {
				var schema map[string]interface{}
				require.NoError(t, json.Unmarshal(evt.DataSchema, &schema))
				props, ok := schema["properties"].(map[string]interface{})
				require.True(t, ok)
				assert.Contains(t, props, "on_chain", "settlement DataSchema should include on_chain")
				assert.Contains(t, props, "off_chain", "settlement DataSchema should include off_chain")
				assert.Contains(t, props, "requests", "settlement DataSchema should include requests")
			}
		})
	}
}
