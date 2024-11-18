package handler

import (
	"bytes"
	"crservice/mocks"
	"crservice/model"
	"crservice/usecase"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewWeatherForecastHandler(t *testing.T) {
	type args struct {
		useCase *usecase.WeatherForecastUseCase
	}
	tests := []struct {
		name string
		args args
		want *WeatherForecastHandler
	}{
		{
			name: "Valid use case",
			args: args{
				useCase: &usecase.WeatherForecastUseCase{},
			},
			want: &WeatherForecastHandler{
				weatherForecastUseCase: &usecase.WeatherForecastUseCase{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWeatherForecastHandler(tt.args.useCase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWeatherForecastHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestWeatherForecastHandler_GetMarineWaveHeight(t *testing.T) {
	// Create a mock use case
	mockUseCase := mocks.NewIWeatherForecastUseCase(t)

	// Test cases
	tests := []struct {
		name           string
		mockSetup      func()
		requestBody    map[string]string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			//ideally I would check for Json string matches the expected string
			name: "Valid request",
			mockSetup: func() {
				mockUseCase.On("GetMarineWaveHeight", mock.Anything, &model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				}).Return(&model.MarineForecast{
					Latitude:  54.5,
					Longitude: 10.25,
				}, nil).Once()
			},
			requestBody: map[string]string{
				"latitude":  "54.5",
				"longitude": "10.25",
			},
			expectedStatus: http.StatusOK,
		},
		{
			name:      "Missing query params",
			mockSetup: func() {},
			requestBody: map[string]string{
				"latitude": "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "Invalid query params",
			},
		},
		{
			name: "Use case returns error",
			mockSetup: func() {
				mockUseCase.On("GetMarineWaveHeight", mock.Anything, &model.MarineWaveHeightForecastReq{
					Latitude:  "54.5",
					Longitude: "10.25",
				}).Return(nil, assert.AnError).Once()
			},
			requestBody: map[string]string{
				"latitude":  "54.5",
				"longitude": "10.25",
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error": assert.AnError.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()
			handlerInstance := &WeatherForecastHandler{
				weatherForecastUseCase: mockUseCase,
			}

			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.POST("/weatherforecast/marinewaveheight", handlerInstance.GetMarineWaveHeight)

			requestBody, _ := json.Marshal(tt.requestBody)

			req := httptest.NewRequest(http.MethodPost, "/weatherforecast/marinewaveheight", bytes.NewReader(requestBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var actualBody map[string]interface{}
			json.Unmarshal(w.Body.Bytes(), &actualBody)

			mockUseCase.AssertExpectations(t)
		})
	}
}
