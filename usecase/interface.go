package usecase

import (
	"context"
	"crservice/model"
)

//go:generate mockery --name=IWeatherForecastUseCase --dir=./ --output=../mocks

type IWeatherForecastUseCase interface {
	GetMarineWaveHeight(ctx *context.Context, request *model.MarineWaveHeightForecastReq) (*model.MarineForecast, error)
}
