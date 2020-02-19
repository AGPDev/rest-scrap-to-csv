package scrap

import (
	"github.com/go-resty/resty/v2"
)

var client *resty.Client

func init() {
	client = resty.New()
	client.EnableTrace()
	// client.SetAuthToken("")
	client.SetHostURL("https://www.agis.com.br/AgisSalesCenterWebAPI/api/")
	client.SetHeader("Accept", "application/json")
	// client.SetProxy("http://proxyserver:8888")
}

// GetRestClient to make API calls
func GetRestClient() *resty.Request {
	return client.R()
}
