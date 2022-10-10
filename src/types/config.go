package types

type Config struct {
	Datasource DataSource
}

type DataSource struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	ApiKey string `json:"apiKey"`
}
