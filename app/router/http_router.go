package router

import (
	"crservice/app/config"
	"crservice/app/middleware/context"
	"crservice/dal"
	"crservice/handler"
	"crservice/usecase"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(config *config.Config) *gin.Engine {
	router := gin.Default()
	router.Use(context.CreateContextWIthTimeout())

	dalInstance := dal.NewWeatherForecastDAL(config)
	useCaseInstance := usecase.NewWeatherForecastUseCase(dalInstance)
	handlerInstance := handler.NewWeatherForecastHandler(useCaseInstance)

	router.GET("weatherforecast/marinewaveheight/", handlerInstance.GetMarineWaveHeight)

	return router
}
