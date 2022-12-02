package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"yaweather/config"
	"yaweather/model"
)

func main() {

	cfg := config.GetConfig(path.Join("etc", ".env"))

	condition := map[string]string{

		"clear":                  "ясно",
		"partly-cloudy":          "малооблачно,",
		"cloudy":                 "облачно с прояснениями",
		"overcast":               "пасмурно",
		"drizzle":                "морось",
		"light-rain":             "небольшой дождь",
		"rain":                   "дождь",
		"moderate-rain":          "умеренно сильный дождь",
		"heavy-rain":             "сильный дождь",
		"continuous-heavy-rain":  "длительный сильный дождь",
		"showers":                "ливень",
		"wet-snow":               "дождь со снегом",
		"light-snow":             "небольшой снег",
		"snow":                   "снег",
		"snow-showers":           "снегопад",
		"hail":                   "град",
		"thunderstorm":           "гроза",
		"thunderstorm-with-rain": "дождь с грозой",
		"thunderstorm-with-hail": "гроза с градом",
	}

	windDirectory := map[string]string{
		"nw": "северо-западное",
		"n":  "северное",
		"ne": "северо-восточное",
		"e":  "восточное",
		"se": "юго-восточное",
		"s":  "южное",
		"sw": "юго-западное",
		"w":  "западное",
		"c":  "штиль",
	}

	w, err := getWeather(cfg) // Норильск
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Широта: %f Долгота: %f\n", w.Info.Lon, w.Info.Lat)
	fmt.Printf("Фактическая температура: %d По ощющениям: %d\n", w.Fact.Temp, w.Fact.FeelsLike)
	fmt.Printf("Влажность: %d%% \n", w.Fact.Humidity)
	fmt.Printf("Скорость ветра: %.1fм.с Напрвление ветра: %s\n", w.Fact.WindSpeed, windDirectory[w.Fact.WindDir])
	fmt.Printf("Давление в паскалях %dПа\nММ ртутного столба: %dмм/рт.ст\n", w.Fact.PressurePA, w.Fact.PressureMM)
	fmt.Printf("%s", condition[w.Fact.Condition])
}

func getWeather(cfg *config.Config) (model.WeatherStruct, error) {

	q := fmt.Sprintf("https://api.weather.yandex.ru/v2/informers/?lat=%f&lon=%f&lang=ru_RU", cfg.Latitude, cfg.Longitude)

	var w model.WeatherStruct

	client := http.Client{}
	request, err := http.NewRequest("GET", q, nil)
	if err != nil {
		return model.WeatherStruct{}, err
	}

	request.Header.Add("X-Yandex-API-Key", cfg.Token)

	resp, err := client.Do(request)
	if err != nil {
		return model.WeatherStruct{}, err
	}
	defer resp.Body.Close()

	fmt.Println("request status: ", resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&w)
	if err != nil {
		return model.WeatherStruct{}, err
	}

	return w, nil

}
