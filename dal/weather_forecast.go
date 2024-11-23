package dal

import (
	"context"
	"crservice/app/config"
	"crservice/entity"
	"crservice/library"
	"crservice/library/http_request"
	"crservice/model"
	"encoding/json"
	"fmt"
)

const waveHeightMeasureType = "hourly"
const waveHeightMeasureName = "wave_height"

type WeatherForecastDAL struct {
	config *config.Config
}

func NewWeatherForecastDAL(cfg *config.Config) *WeatherForecastDAL {
	return &WeatherForecastDAL{config: cfg}
}

func (d *WeatherForecastDAL) GetMarineWaveHeight(ctx *context.Context, params model.MarineWaveHeightForecastReq) (*entity.MarineForecast, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is nil")
	}

	getUrl := d.config.MarineForecast.Url
	paramsMap, paramsErr := generateMarineWaveHeightForecastParams(params)
	if paramsErr != nil {
		return nil, paramsErr
	}

	response, err := http_request.Get(getUrl, paramsMap)
	if err != nil {
		return nil, err
	}

	marineForecast, forecast, err2 := parseMarineWaveHeightResponse(err, response)
	if err2 != nil {
		return forecast, err2
	}
	return marineForecast, nil
}

func parseMarineWaveHeightResponse(err error, response []byte) (*entity.MarineForecast, *entity.MarineForecast, error) {
	var marineForecast *entity.MarineForecast
	err = json.Unmarshal(response, &marineForecast)
	if err != nil {
		return nil, nil, err
	}
	return marineForecast, nil, nil
}

func generateMarineWaveHeightForecastParams(params model.MarineWaveHeightForecastReq) (map[string]string, error) {
	paramsMap, err := library.StructToStringMap(params)
	if err != nil {
		return nil, err
	}
	paramsMap[waveHeightMeasureType] = waveHeightMeasureName
	return paramsMap, err
}
