package weather

type Details struct {
	AirTemperature float32 `json:"air_temperature"`
}

type Instant struct {
	Details Details `json:"details"`
}

type TimeSeriesData struct {
	Instant Instant `json:"instant"`
}

type TimeSeries struct {
	Time string         `json:"time"`
	Data TimeSeriesData `json:"data"`
}

type Properties struct {
	Timeseries []TimeSeries `json:"timeseries"`
}

type YrLocationForecast struct {
	Properties Properties `json:"properties"`
	Type       string     `json:"type"`
}
