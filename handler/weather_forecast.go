package handler

import (
	"crservice/app/middleware/context"
	"crservice/model"
	"crservice/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WeatherForecastHandler struct {
	weatherForecastUseCase usecase.IWeatherForecastUseCase
}

func NewWeatherForecastHandler(useCase usecase.IWeatherForecastUseCase) *WeatherForecastHandler {
	return &WeatherForecastHandler{
		weatherForecastUseCase: useCase,
	}
}

func (h *WeatherForecastHandler) GetMarineWaveHeight(c *gin.Context) {
	ctx := context.NewContext(c)
	if ctx == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
		return
	}
	var params model.MarineWaveHeightForecastReq

	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil})
		return
	}
	if params.Longitude == "" || params.Latitude == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil})
		return
	}

	result, err := h.weatherForecastUseCase.GetMarineWaveHeight(&ctx, &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}
