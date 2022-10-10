package types

type Config struct {
	Datasource DataSource
	Chain      string `json:"chain"`
	ScanType   string `json:"scanType"`
}

type DataSource struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	ApiKey string `json:"apiKey"`
}
