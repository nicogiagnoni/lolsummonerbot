package riot

var apiKey string

func SetAPIKey(key string) {
	apiKey = key
}

func GetAPIKey() string {
	return apiKey
}
