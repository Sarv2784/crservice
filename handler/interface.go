package handler

import (
	"context"
)

type IWeatherForecastHandler interface {
	GetMarineWaveHeight(ctx *context.Context)
}
