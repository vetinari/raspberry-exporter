package collector

import (
	"regexp"
	"strconv"

	"github.com/derknerd/raspberry-exporter/utils"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	GpuClock  string = "core"
	EmmcClock string = "emmc"
	ArmClock  string = "arm"
)

var (
	clockRegex = regexp.MustCompile(`frequency\(\d*\)=|(\n)|(\r)`)
)

func (c *VcGenCmdCollector) getClock(desc *prometheus.Desc, device string) prometheus.Metric {
	clock, err := utils.ExecuteVcGen(c.VcGenCmd,"measure_clock", device)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	clock = clockRegex.ReplaceAllString(clock, "")

	clockFloat, err := strconv.ParseFloat(clock, 64)

	if err != nil {
		return prometheus.NewInvalidMetric(desc, err)
	}

	return prometheus.MustNewConstMetric(
		desc,
		prometheus.GaugeValue,
		clockFloat,
	)
}
