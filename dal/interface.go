package dal

import (
	"context"
	"crservice/entity"
	"crservice/model"
)

//go:generate mockery --name=IWeatherForecastDAL --dir=./ --output=../mocks

type IWeatherForecastDAL interface {
	GetMarineWaveHeight(ctx *context.Context, params model.MarineWaveHeightForecastReq) (*entity.MarineForecast, error)
}
