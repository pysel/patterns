// Observer Pattern
package main

type TVStation interface {
	update(ForecastData) TVStation
}

type WeatherForecast struct {
	ForecastData
	Subscribers []*TVStation
}

func NewWeatherForecast(data ForecastData, subs []*TVStation) WeatherForecast {
	return WeatherForecast{data, subs}
}

func (wf *WeatherForecast) addSubscriber(sub TVStation) {
	wf.Subscribers = append(wf.Subscribers, &sub)
}

func (wf *WeatherForecast) removeSubscriber(sub TVStation) {
	ind := mustFindIndex(wf.Subscribers, sub)
	wf.Subscribers = remove(wf.Subscribers, ind)
}

func (wf *WeatherForecast) updateData(newData ForecastData) {
	wf.ForecastData = newData
	wf.notifySubs()
}

func (wf WeatherForecast) notifySubs() {
	for _, sub := range wf.Subscribers {
		(*sub).update(wf.ForecastData)
	}
}

// As an example I use an imitation of a weather forcast server as a subject and TV stations as observers (clients)
func main() {
	// no factory method to save time AND because objects are empty when initializing them.
	observer1 := TVstation1{}
	observer2 := TVstation2{}

	subject := NewWeatherForecast(NewForecastData(15, 5), nil)

	subject.addSubscriber(observer1)
	subject.addSubscriber(observer2)
	subject.notifySubs() // manually run to view changes

	subject.updateData(NewForecastData(10, 10)) // automatically updates all observers when data from subject changes

	subject.removeSubscriber(observer1)

	subject.updateData(NewForecastData(1, 1)) // only shows data from TVstation2 because observer1 was removed
}
