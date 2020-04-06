package all

import (
	_ "github.com/influxdata/telegraf/plugins/outputs/amon"
	_ "github.com/influxdata/telegraf/plugins/outputs/amqp"
	_ "github.com/influxdata/telegraf/plugins/outputs/application_insights"
	_ "github.com/influxdata/telegraf/plugins/outputs/azure_monitor"
	_ "github.com/influxdata/telegraf/plugins/outputs/cloud_pubsub"
	_ "github.com/influxdata/telegraf/plugins/outputs/cloudwatch"
	_ "github.com/influxdata/telegraf/plugins/outputs/cratedb"
	_ "github.com/influxdata/telegraf/plugins/outputs/datadog"
	_ "github.com/influxdata/telegraf/plugins/outputs/discard"
	_ "github.com/influxdata/telegraf/plugins/outputs/elasticsearch"
	_ "github.com/influxdata/telegraf/plugins/outputs/exec"
	_ "github.com/influxdata/telegraf/plugins/outputs/file"
	_ "github.com/influxdata/telegraf/plugins/outputs/graphite"
	_ "github.com/influxdata/telegraf/plugins/outputs/graylog"
	_ "github.com/influxdata/telegraf/plugins/outputs/health"
	_ "github.com/influxdata/telegraf/plugins/outputs/http"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/influxdb_v2"
	_ "github.com/influxdata/telegraf/plugins/outputs/instrumental"
	_ "github.com/influxdata/telegraf/plugins/outputs/kafka"
	_ "github.com/influxdata/telegraf/plugins/outputs/kinesis"
	_ "github.com/influxdata/telegraf/plugins/outputs/librato"
	_ "github.com/influxdata/telegraf/plugins/outputs/mqtt"
	_ "github.com/influxdata/telegraf/plugins/outputs/nats"
	_ "github.com/influxdata/telegraf/plugins/outputs/nsq"
	_ "github.com/influxdata/telegraf/plugins/outputs/opentsdb"
	_ "github.com/influxdata/telegraf/plugins/outputs/prometheus_client"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann"
	_ "github.com/influxdata/telegraf/plugins/outputs/riemann_legacy"
	_ "github.com/influxdata/telegraf/plugins/outputs/socket_writer"
	_ "github.com/influxdata/telegraf/plugins/outputs/stackdriver"
	_ "github.com/influxdata/telegraf/plugins/outputs/syslog"
	_ "github.com/influxdata/telegraf/plugins/outputs/warp10"
	_ "github.com/influxdata/telegraf/plugins/outputs/wavefront"
)