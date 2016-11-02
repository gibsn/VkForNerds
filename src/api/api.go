package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// https://oauth.vk.com/authorize?client_id=5680126&display=mobile&redirect_uri=https://oauth.vk.com/blank.html%20&scope=messages,offline&response_type=code&v=5.59
// https://oauth.vk.com/access_token?client_id=5680126&client_secret=ehxgcUW4eGAArVVwx6Cd&redirect_uri=https://oauth.vk.com/blank.html&code=c1213133c1054533d4

type Api struct {
	AccessToken string
}

type DialogJson struct {
	Date      uint64
	Out       int
	Uid       uint64
	ReadState int `json:"read_state"`
	Title     string
	Body      string
}

type Dialog struct {
	Uid          uint64
	ReadState    int
	Title        string //useful for group chats, "..." for tet-a-tet
	FirstMessage string
	InputBuf     string
	// Date      uint64
	// Messages  []Messages
}

var apiUrl = "https://api.vk.com/method/"

func NewApi() *Api {
	api := &Api{}

	return api
}

//TODO
func Auth() {

}

func (this *Api) RequestDialogsHeaders() []Dialog {
	response := this.request("messages.getDialogs", &map[string]string{})
	if response == nil {
		return nil
	}

	body := json.NewDecoder(response.Body)

	//hacking over the poorly designed array in response
	for i := 0; i < 4; i++ {
		_, _ = body.Token()
	}

	var dialogJson DialogJson
	var dialogs []Dialog
	for body.More() {
		body.Decode(&dialogJson)

		newDialog := &Dialog{
			Uid:          dialogJson.Uid,
			ReadState:    dialogJson.ReadState,
			Title:        dialogJson.Title,
			FirstMessage: dialogJson.Body,
			InputBuf:     "",
		}
		dialogs = append(dialogs, *newDialog)
	}

	response.Body.Close()

	return dialogs
}

func (this *Api) request(method string, params *map[string]string) *http.Response {
	url, err := url.Parse(apiUrl + method)

	query := url.Query()
	for key, value := range *params {
		query.Set(key, value)
	}

	query.Set("access_token", this.AccessToken)
	url.RawQuery = query.Encode()

	response, err := http.Get(url.String())
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return response
}
