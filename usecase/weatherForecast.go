package usecase

import (
	"context"
	"crservice/dal"
	"crservice/model"
	"fmt"
)

type WeatherForecastUseCase struct {
	dal dal.IWeatherForecastDAL
}

func NewWeatherForecastUseCase(dal dal.IWeatherForecastDAL) *WeatherForecastUseCase {
	return &WeatherForecastUseCase{
		dal: dal,
	}
}

func (u *WeatherForecastUseCase) GetMarineWaveHeight(ctx *context.Context, request *model.MarineWaveHeightForecastReq) (*model.MarineForecast, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}

	marineForecastEntity, err := u.dal.GetMarineWaveHeight(ctx, *request)
	if err != nil {
		return nil, err
	}

	return model.ParseMarineForecastEntityToModel(marineForecastEntity), nil
}
