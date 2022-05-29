package protoprocessor

import (
	"testing"
	"time"

	commonpb "github.com/census-instrumentation/opencensus-proto/gen-go/agent/common/v1"
	metricspb "github.com/census-instrumentation/opencensus-proto/gen-go/metrics/v1"
	resourcepb "github.com/census-instrumentation/opencensus-proto/gen-go/resource/v1"
	"github.com/google/go-cmp/cmp"
        "go.opentelemetry.io/collector/translator/internaldata"
	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/model/otlp"
	"go.opentelemetry.io/collector/model/pdata"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type metricData struct {
	Node     *commonpb.Node
	Resource *resourcepb.Resource
	Metrics  []*metricspb.Metric
}

func createDoubleMetricsDataWithTimestamp(name string, timestamp *timestamppb.Timestamp) metricData {
	doubleGauge := metricspb.MetricDescriptor{
		Name:      name,
		Type:      metricspb.MetricDescriptor_GAUGE_DOUBLE,
		LabelKeys: []*metricspb.LabelKey{{Key: "label1"}, {Key: "label2"}},
	}
	return metricData{
		Metrics: []*metricspb.Metric{
			{
				MetricDescriptor: &doubleGauge,
				Timeseries: []*metricspb.TimeSeries{
					{
						LabelValues: []*metricspb.LabelValue{
							{Value: "label1-value1", HasValue: true},
							{Value: "label2-value1", HasValue: true},
						},
						Points:         []*metricspb.Point{{Value: &metricspb.Point_DoubleValue{DoubleValue: 0}, Timestamp: timestamp}},
						StartTimestamp: timestamp,
					},
				},
			},
		},
	}
}

func metricsFromOtlpProtoBytes(data []byte) (pdata.Metrics, error) {
	unmarshaler := otlp.NewProtobufMetricsUnmarshaler()
	return unmarshaler.UnmarshalMetrics(data)
}

func collectMds(nextMs []pdata.Metrics) []metricData {
	var nextMds []metricData
	for _, m := range nextMs {
		rms := m.ResourceMetrics()
		for i := 0; i < rms.Len(); i++ {
			node, resource, metrics := internaldata.ResourceMetricsToOC(rms.At(i))
			nextMds = append(nextMds, metricData{node, resource, metrics})
		}
	}
	return nextMds
}

func TestMetricsFromOtlpProtoBytesOldVersion2(t *testing.T) {
	expectedMd := createDoubleMetricsDataWithTimestamp("test", timestamppb.New(time.Unix(1642012747, 0)))
	metricBytesVersion1_1 := []byte{10, 86, 10, 0, 18, 82, 10, 0, 18, 78, 10, 4, 116, 101, 115, 116, 42, 70, 10, 68, 10, 23, 10, 6, 108, 97, 98, 101, 108, 49, 18, 13, 108, 97, 98, 101, 108, 49, 45, 118, 97, 108, 117, 101, 49, 10, 23, 10, 6, 108, 97, 98, 101, 108, 50, 18, 13, 108, 97, 98, 101, 108, 50, 45, 118, 97, 108, 117, 101, 49, 17, 0, 46, 153, 197, 228, 153, 201, 22, 25, 0, 46, 153, 197, 228, 153, 201, 22}

	ms, err := metricsFromOtlpProtoBytes(metricBytesVersion1_1)
	assert.NoError(t, err)
	md := collectMds([]pdata.Metrics{ms.Clone()})
	assert.Equal(t, 1, len(md))
	if diff := cmp.Diff(expectedMd, md[0], protocmp.Transform()); diff != "" {
		t.Errorf("metricsFromOtlpProtoBytes(metricsData v1.1.0-anthos.7 version) returned unexpected diff (-want +got):\n%s, with metrics data = %s", diff, expectedMd)
	}
}

