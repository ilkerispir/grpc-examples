package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"weather/api"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	api.RegisterWeatherServiceServer(srv, &myWeatherService{})
	fmt.Println("starting server...")
	panic(srv.Serve(lis))
}

type myWeatherService struct {
	api.UnimplementedWeatherServiceServer
}

func (m *myWeatherService) ListCities(ctx context.Context,
	req *api.ListCitiesRequest) (*api.ListCitiesResponse, error) {

	return &api.ListCitiesResponse{
		Items: []*api.CityEntry{
			&api.CityEntry{
				CityCode: "tr_ank",
				CityName: "Ankara",
			},
			&api.CityEntry{
				CityCode: "tr_ist",
				CityName: "Istanbul",
			},
		},
	}, nil
}

func (m *myWeatherService) QueryWeather(req *api.WeatherRequest,
	resp api.WeatherService_QueryWeatherServer) error {

	for {
		err := resp.Send(&api.WeatherResponse{
			Temperature: rand.Float32()*10 + 10})
		if err != nil {
			break
		}
	}

	return nil
}
