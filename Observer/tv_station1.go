package main

import "fmt"

type TVstation1 struct {
	ForecastData
}

func (s TVstation1) update(newData ForecastData) TVStation {
	s.ForecastData = newData
	fmt.Println("New data for TVStation1 is: ", newData)
	return s
}
