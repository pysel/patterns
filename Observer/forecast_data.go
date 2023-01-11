package main

type ForecastData struct {
	temperature uint32
	wind        int64
}

func NewForecastData(tmp uint32, wind int64) ForecastData {
	return ForecastData{temperature: tmp, wind: wind}
}
