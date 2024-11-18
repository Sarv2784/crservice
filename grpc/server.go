package grpc

import (
	"context"
	"crservice/app/config"
	"crservice/dal"
	"crservice/model"
	"crservice/usecase"
	"fmt"
	"net"

	pb "crservice/grpc/proto"

	"google.golang.org/grpc"
)

type grpcServer struct {
	pb.UnimplementedWeatherForecastServiceServer
	useCase usecase.WeatherForecastUseCase
}

func NewGRPCServer(cfg *config.Config) error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPCServer.Host, cfg.GRPCServer.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	dalInstance := dal.NewWeatherForecastDAL(cfg)
	useCaseInstance := usecase.NewWeatherForecastUseCase(dalInstance)

	s := grpc.NewServer()
	pb.RegisterWeatherForecastServiceServer(s, &grpcServer{
		useCase: *useCaseInstance,
	})

	fmt.Printf("gRPC server listening at %s:%d\n", cfg.GRPCServer.Host, cfg.GRPCServer.Port)
	return s.Serve(lis)
}

func (s *grpcServer) GetMarineWaveHeight(ctx context.Context, req *pb.MarineWaveHeightForecastReq) (*pb.MarineForecast, error) {
	reqModel := &model.MarineWaveHeightForecastReq{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
	data, err := s.useCase.GetMarineWaveHeight(&ctx, reqModel)
	if err != nil {
		return nil, err
	}

	return ConvertToProtobufMarineForecast(*data), nil
}

func ConvertToProtobufMarineForecast(marineForecast model.MarineForecast) *pb.MarineForecast {
	return &pb.MarineForecast{
		Latitude:             marineForecast.Latitude,
		Longitude:            marineForecast.Longitude,
		GenerationTimeMs:     marineForecast.GenerationTimeMs,
		UtcOffsetSeconds:     int32(marineForecast.UTCOffsetSeconds),
		Timezone:             marineForecast.Timezone,
		TimezoneAbbreviation: marineForecast.TimezoneAbbreviation,
		Elevation:            marineForecast.Elevation,
		HourlyUnits: &pb.HourlyUnits{
			Time:       marineForecast.HourlyUnits.Time,
			WaveHeight: marineForecast.HourlyUnits.WaveHeight,
		},
		Hourly: &pb.Hourly{
			Time:       marineForecast.Hourly.Time,
			WaveHeight: marineForecast.Hourly.WaveHeight,
		},
	}
}
