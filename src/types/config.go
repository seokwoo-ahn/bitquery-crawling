package types

type Config struct {
	Datasource DataSource
}

type DataSource struct {
	DataSource struct {
		Name   string `json:"name"`
		Url    string `json:"url"`
		ApiKey string `json:"apiKey"`
	}
}
