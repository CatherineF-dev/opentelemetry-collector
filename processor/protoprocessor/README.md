1. 
```
git checkout tags/v0.31.0 -b v0.31.0
```

2. run this test failed. 
```
go test -timeout 30s -run ^TestMetricsFromOtlpProtoBytesOldVersion2$
```
```
--- FAIL: TestMetricsFromOtlpProtoBytesOldVersion2 (0.01s)
    proto_test.go:79: metricsFromOtlpProtoBytes(metricsData v1.1.0-anthos.7 version) returned unexpected diff (-want +got):
          protoprocessor.metricData{
          	Node:     Inverse(protocmp.Transform, protocmp.Message{"@invalid": bool(true), "@type": s"opencensus.proto.agent.common.v1.Node"}),
          	Resource: Inverse(protocmp.Transform, protocmp.Message{"@invalid": bool(true), "@type": s"opencensus.proto.resource.v1.Resource"}),
          	Metrics: []*v1.Metric{
          		Inverse(protocmp.Transform, protocmp.Message{
          			"@type":             s"opencensus.proto.metrics.v1.Metric",
          			"metric_descriptor": protocmp.Message{"@type": s"opencensus.proto.metrics.v1.MetricDescriptor", "label_keys": []protocmp.Message{{"@type": s"opencensus.proto.metrics.v1.LabelKey", "key": string("label1")}, {"@type": s"opencensus.proto.metrics.v1.LabelKey", "key": string("label2")}}, "name": string("test"), "type": s"GAUGE_DOUBLE"},
          			"timeseries": []protocmp.Message{
          				{
          					"@type":        s"opencensus.proto.metrics.v1.TimeSeries",
          					"label_values": []protocmp.Message{{"@type": s"opencensus.proto.metrics.v1.LabelValue", "has_value": bool(true), "value": string("label1-value1")}, {"@type": s"opencensus.proto.metrics.v1.LabelValue", "has_value": bool(true), "value": string("label2-value1")}},
          					"points": []protocmp.Message{
          						{
          							"@type":        s"opencensus.proto.metrics.v1.Point",
        - 							"double_value": float64(0),
          							"timestamp":    protocmp.Message{"@type": s"google.protobuf.Timestamp", "seconds": int64(1642012747)},
          						},
          					},
          					"start_timestamp": protocmp.Message{"@type": s"google.protobuf.Timestamp", "seconds": int64(1642012747)},
          				},
          			},
          		}),
          	},
          }
        , with metrics data = {<nil> <nil> [metric_descriptor:{name:"test"  type:GAUGE_DOUBLE  label_keys:{key:"label1"}  label_keys:{key:"label2"}}  timeseries:{start_timestamp:{seconds:1642012747}  label_values:{value:"label1-value1"  has_value:true}  label_values:{value:"label2-value1"  has_value:true}  points:{timestamp:{seconds:1642012747}  double_value:0}}]}
FAIL
exit status 1
FAIL	go.opentelemetry.io/collector/processor/protoprocessor	0.123s
```
