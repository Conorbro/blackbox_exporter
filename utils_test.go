package main

import (
	"fmt"
	"testing"

	dto "github.com/prometheus/client_model/go"
)

// Check if expected results are in the registry
func checkRegistryResults(expRes map[string]float64, mfs []*dto.MetricFamily, t *testing.T) {
	for i := range mfs {
		if _, ok := expRes[mfs[i].GetName()]; ok {
			if expRes[mfs[i].GetName()] == mfs[i].Metric[0].GetGauge().GetValue() {
			} else {
				t.Fatal(fmt.Errorf("Expected %v %v, got %v %v", mfs[i].GetName(), expRes[mfs[i].GetName()], mfs[i].GetName(), mfs[i].Metric[0].GetGauge().GetValue()))
			}
		}
	}
}
