package add_tag

import (
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/processors"
)

const sampleConfig = `
`

type Addtag struct {
        Tags        map[string]string `toml:"if_tags"`
	Measurement string `toml:"measurement"`
	Add         map[string]string `toml:"add_tags"`
}

func (a *Addtag) SampleConfig() string {
	return sampleConfig
}

func (a *Addtag) Description() string {
	return "Add tags, given others tag present"
}

func (a *Addtag) Init() error {
	return nil
}

func (a *Addtag) Apply(in ...telegraf.Metric) []telegraf.Metric {
	for _, point := range in {
		if a.Measurement != point.Name() { continue }
                lenTags := len(a.Tags)
		for tag, value := range a.Tags {
			v, ok := point.GetTag(tag)
			if !(ok && (value == v)) { continue } 
                        lenTags--
			if lenTags <= 0 {
                		for k, t := range a.Add { point.AddTag(k,t) }
				break
			}
		}
	}
	return in
}

func init() {
	processors.Add("add_tag", func() telegraf.Processor {
		return &Addtag{}
	})
}
