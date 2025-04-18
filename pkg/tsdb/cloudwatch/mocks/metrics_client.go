package mocks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"

	"github.com/grafana/grafana/pkg/tsdb/cloudwatch/models/resources"
	"github.com/stretchr/testify/mock"
)

type FakeMetricsClient struct {
	mock.Mock
}

func (m *FakeMetricsClient) ListMetricsWithPageLimit(_ context.Context, params *cloudwatch.ListMetricsInput) ([]resources.MetricResponse, error) {
	args := m.Called(params)
	return args.Get(0).([]resources.MetricResponse), args.Error(1)
}
