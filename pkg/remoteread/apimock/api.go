// Code generated by mockery 2.7.4. DO NOT EDIT.

package apimock

import (
	context "context"

	model "github.com/prometheus/common/model"
	mock "github.com/stretchr/testify/mock"

	time "time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

// LabelNames provides a mock function with given fields: ctx, matches, startTime, endTime
func (_m *API) LabelNames(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]string, v1.Warnings, error) {
	ret := _m.Called(ctx, matches, startTime, endTime)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time, time.Time) []string); ok {
		r0 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, time.Time, time.Time) error); ok {
		r2 = rf(ctx, matches, startTime, endTime)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LabelValues provides a mock function with given fields: ctx, label, matches, startTime, endTime
func (_m *API) LabelValues(ctx context.Context, label string, matches []string, startTime time.Time, endTime time.Time) (model.LabelValues, v1.Warnings, error) {
	ret := _m.Called(ctx, label, matches, startTime, endTime)

	var r0 model.LabelValues
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, time.Time, time.Time) model.LabelValues); ok {
		r0 = rf(ctx, label, matches, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.LabelValues)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, time.Time, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, label, matches, startTime, endTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, []string, time.Time, time.Time) error); ok {
		r2 = rf(ctx, label, matches, startTime, endTime)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Query provides a mock function with given fields: ctx, query, ts
func (_m *API) Query(ctx context.Context, query string, ts time.Time) (model.Value, v1.Warnings, error) {
	ret := _m.Called(ctx, query, ts)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) model.Value); ok {
		r0 = rf(ctx, query, ts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Value)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, query, ts)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, time.Time) error); ok {
		r2 = rf(ctx, query, ts)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// QueryRange provides a mock function with given fields: ctx, query, r
func (_m *API) QueryRange(ctx context.Context, query string, r v1.Range) (model.Value, v1.Warnings, error) {
	ret := _m.Called(ctx, query, r)

	var r0 model.Value
	if rf, ok := ret.Get(0).(func(context.Context, string, v1.Range) model.Value); ok {
		r0 = rf(ctx, query, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.Value)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, string, v1.Range) v1.Warnings); ok {
		r1 = rf(ctx, query, r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, v1.Range) error); ok {
		r2 = rf(ctx, query, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Series provides a mock function with given fields: ctx, matches, startTime, endTime
func (_m *API) Series(ctx context.Context, matches []string, startTime time.Time, endTime time.Time) ([]model.LabelSet, v1.Warnings, error) {
	ret := _m.Called(ctx, matches, startTime, endTime)

	var r0 []model.LabelSet
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time, time.Time) []model.LabelSet); ok {
		r0 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.LabelSet)
		}
	}

	var r1 v1.Warnings
	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time, time.Time) v1.Warnings); ok {
		r1 = rf(ctx, matches, startTime, endTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(v1.Warnings)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, []string, time.Time, time.Time) error); ok {
		r2 = rf(ctx, matches, startTime, endTime)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}