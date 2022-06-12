package main

import (
	"context"
	"fmt"
	"io"
	"weather/api"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := api.NewWeatherServiceClient(conn)

	ctx := context.Background()

	resp, err := client.ListCities(ctx, &api.ListCitiesRequest{})
	if err != nil {
		panic(err)
	}

	fmt.Println("cities: ")
	for _, city := range resp.Items {
		fmt.Printf("\t%s: %s\n", city.GetCityCode(), city.CityName)
	}

	stream, err := client.QueryWeather(ctx, &api.WeatherRequest{
		CidyCode: "tr_ank",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("weather in Ankara: ")
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("\t temprature: %.2f\n", msg.GetTemperature())
	}
	fmt.Println("server stopped sending")
}
