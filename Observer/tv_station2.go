package main

import "fmt"

type TVstation2 struct {
	ForecastData
}

func (s TVstation2) update(newData ForecastData) TVStation {
	s.ForecastData = newData
	fmt.Println("New data for TVStation2 is: ", newData)
	return s
}
