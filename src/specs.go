package main

import (
	"os"
	"strconv"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

type argumentList struct {
	sdkArgs.DefaultArgumentList
}

const (
	integrationName    = "com.newrelic.specs"
	integrationVersion = "0.1.0"
)

var (
	args argumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Read cpuinfo data
	stat, err := linuxproc.ReadCPUInfo("/proc/cpuinfo")
	if err != nil {
		log.Error(err.Error())
	}

	// Retrieve local entity
	e1 := i.LocalEntity()

	// Loop processors and create dataset
	for _, s := range stat.Processors {
		metricSet := NewMetricSet("processorSpecSample", e1)
		SetMetric(metricSet, "cacheSize", s.CacheSize)
		SetMetric(metricSet, "coreId", s.CoreId)
		SetMetric(metricSet, "cores", s.Cores)
		SetMetric(metricSet, "flags", s.Flags)
		SetMetric(metricSet, "mhz", s.MHz)
		SetMetric(metricSet, "model", s.Model)
		SetMetric(metricSet, "modelName", s.ModelName)
		SetMetric(metricSet, "physicalId", s.PhysicalId)
		SetMetric(metricSet, "vendorId", s.VendorId)
	}

	if err = i.Publish(); err != nil {
		log.Error(err.Error())
	}
}

// SetMetric is a helper function for pushing metrics together with NewMetricSet
func SetMetric(metricSet *metric.Set, key string, val interface{}) {
	switch val.(type) {
	case float64:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case uint16:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case uint32:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case uint64:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case int:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case int32:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case int64:
		metricSet.SetMetric(key, val, metric.GAUGE)
	case bool:
		metricSet.SetMetric(key, strconv.FormatBool(val.(bool)), metric.ATTRIBUTE)
	case string:
		if val.(string) != "" {
			metricSet.SetMetric(key, val, metric.ATTRIBUTE)
		}
	}
}

// NewMetricSet creates a new Metric set for use with SetMetric
func NewMetricSet(event string, entity *integration.Entity) *metric.Set {
	metricSet := entity.NewMetricSet(event)

	return metricSet
}
