// Code generated by mockery v2.47.0. DO NOT EDIT.

package mocks

import (
	context "context"
	model "crservice/model"

	mock "github.com/stretchr/testify/mock"
)

// IWeatherForecastUseCase is an autogenerated mock type for the IWeatherForecastUseCase type
type IWeatherForecastUseCase struct {
	mock.Mock
}

// GetMarineWaveHeight provides a mock function with given fields: ctx, request
func (_m *IWeatherForecastUseCase) GetMarineWaveHeight(ctx *context.Context, request *model.MarineWaveHeightForecastReq) (*model.MarineForecast, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetMarineWaveHeight")
	}

	var r0 *model.MarineForecast
	var r1 error
	if rf, ok := ret.Get(0).(func(*context.Context, *model.MarineWaveHeightForecastReq) (*model.MarineForecast, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*context.Context, *model.MarineWaveHeightForecastReq) *model.MarineForecast); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.MarineForecast)
		}
	}

	if rf, ok := ret.Get(1).(func(*context.Context, *model.MarineWaveHeightForecastReq) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIWeatherForecastUseCase creates a new instance of IWeatherForecastUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIWeatherForecastUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IWeatherForecastUseCase {
	mock := &IWeatherForecastUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
