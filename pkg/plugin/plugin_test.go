package plugin_test

import (
	"context"
	"testing"

	"github.com/computestdev/grafana-cloudwatch-alarm-datasource/pkg/plugin"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// This is where the tests for the datasource backend live.
func TestQueryData(t *testing.T) {
	ds := plugin.CloudWatchAlarmDatasource{}

	resp, err := ds.QueryData(
		context.Background(),
		&backend.QueryDataRequest{
			Queries: []backend.DataQuery{
				{RefID: "A"},
			},
		},
	)
	if err != nil {
		t.Error(err)
	}

	if len(resp.Responses) != 1 {
		t.Fatal("QueryData must return a response")
	}
}