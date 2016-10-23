package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type Api struct {
	accessToken string
}

type Dialog struct {
	Date      uint64
	Out       int
	Uid       uint64
	ReadState int `json:"read_state"`
	Title     string
	Body      string
}

//https://oauth.vk.com/authorize?client_id=5680126&display=mobile&redirect_uri=https://oauth.vk.com/blank.html%20&scope=messages,offline&response_type=code&v=5.59
//https://oauth.vk.com/access_token?client_id=5680126&client_secret=ehxgcUW4eGAArVVwx6Cd&redirect_uri=https://oauth.vk.com/blank.html&code=c1213133c1054533d4
var apiUrl = "https://api.vk.com/method/"

func NewApi() *Api {
	api := &Api{}
	api.accessToken = ""

	return api
}

func Auth() {

}

func (this *Api) RequestDialogsHeaders() {
	response := this.request("messages.getDialogs", &map[string]string{})

	body := json.NewDecoder(response.Body)

	//hacking over the poorly designed array in response
	for i := 0; i < 4; i++ {
		_, _ = body.Token()
	}

	var dialog Dialog
	for body.More() {
		body.Decode(&dialog)
		fmt.Println(dialog)
	}

	response.Body.Close()
}

func (this *Api) request(method string, params *map[string]string) *http.Response {
	url, err := url.Parse(apiUrl + method)

	query := url.Query()
	for key, value := range *params {
		query.Set(key, value)
	}

	query.Set("access_token", this.accessToken)
	url.RawQuery = query.Encode()

	response, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return response
}
