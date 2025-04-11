package ports

type AlertFetcher interface {
	FetchAlerts() ([]map[string]interface{}, error)
}
