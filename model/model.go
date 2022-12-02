package model

type WeatherStruct struct {
	Now   int64  `json:"now"`
	NowDt string `json:"now_dt"`
	Info  struct {
		Lat float32 `json:"lat"`
		Lon float32 `json:"lon"`
		Url string  `json:"url"`
	} `json:"info"`
	Fact struct {
		Temp       int     `json:"temp"`
		FeelsLike  int     `json:"feels_like"`
		TempWater  int     `json:"temp_water"`
		Icon       string  `json:"icon"`
		Condition  string  `json:"condition"`
		WindSpeed  float32 `json:"wind_speed"`
		WindGust   float32 `json:"wind_gust"`
		WindDir    string  `json:"wind_dir"`
		PressureMM int     `json:"pressure_mm"`
		PressurePA int     `json:"pressure_pa"`
		Humidity   int     `json:"humidity"`
		DayTime    string  `json:"dayTime"`
		Polar      bool    `json:"polar"`
		Season     string  `json:"season"`
		ObsTime    int64   `json:"obs_time"`
		AccumPrec  int     `json:"accum_prec"`
	} `json:"fact"`
	Forecast struct {
		Date     string `json:"date"`
		DateTS   int64  `json:"date_tS"`
		Week     int    `json:"week"`
		Sunrise  string `json:"sunrise"`
		Sunset   string `json:"sunset"`
		MoonCode int    `json:"moon_code"`
		MoonText string `json:"moon_text"`
		Parts    []struct {
			PartName  string `json:"part_name,omitempty"`
			TempMin   int    `json:"temp_min,omitempty"`
			TempMax   int    `json:"temp_max,omitempty"`
			TempAvg   int    `json:"temp_avg,omitempty"`
			FeelsLike int    `json:"feels_like,omitempty"`
			Icon      string `json:"icon,omitempty"`
			Condition string `json:"condition,omitempty"`
			Daytime   string `json:"daytime,omitempty"`
			//Polar      bool    `json:"polar,omitempty"`
			WindSpeed  float32 `json:"wind_speed,omitempty"`
			WindGust   float32 `json:"wind_gust,omitempty"`
			WindDir    string  `json:"wind_dir,omitempty"`
			PressureMM int     `json:"pressure_mm,omitempty"`
			PressurePA int     `json:"pressure_pa,omitempty"`
			Humidity   int     `json:"humidity,omitempty"`
			PrecMM     float32 `json:"prec_mm,omitempty"`
			PrecPeriod int     `json:"prec_period,omitempty"`
			PrecProb   int     `json:"prec_prob,omitempty"`
		} `json:"parts"`
		PartName   string `json:"part_name"`
		TempMin    int    `json:"temp_min"`
		TempMax    int    `json:"temp_max"`
		TempAvg    int    `json:"temp_avg"`
		FeelsLike  int    `json:"feels_like"`
		Icon       string `json:"icon"`
		Condition  string `json:"condition"`
		Daytime    string `json:"daytime"`
		Polar      string `json:"polar,omitempty"`
		WindSpeed  int    `json:"wind_speed"`
		WindGust   int    `json:"wind_gust"`
		WindDir    string `json:"wind_dir"`
		PressureMM int    `json:"pressure_mm"`
		PressurePA int    `json:"pressure_pa"`
		Humidity   int    `json:"humidity"`
		PrecMM     int    `json:"prec_mm"`
		PrecPeriod int    `json:"prec_period"`
		PrecProb   int    `json:"prec_prob"`
	} `json:"forecast,omitempty"`
}
