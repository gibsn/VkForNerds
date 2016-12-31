package api

import (
	"encoding/json"
	"fmt"
	"log"
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

type UserJson struct {
	id        uint64
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type ErrorJson struct {
	Code int    `json:"error_code"`
	Msg  string `json:"error_msg"`
}

type Dialog struct {
	Uid          uint64
	FullName     string
	ReadState    int
	Title        string //useful for group chats, "..." for tet-a-tet
	FirstMessage string
	InputBuf     string
	// Date      uint64
	// Messages  []Messages
}

var apiUrl = "https://api.vk.com/method/"

func NewApi(token string) *Api {
	api := &Api{
		AccessToken: token,
	}

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

	defer response.Body.Close()
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

	return dialogs
}

func checkError(body *json.Decoder) *ErrorJson {
	_, _ = body.Token()
	respType, _ := body.Token()

	if respType == "error" {
		var errorJson ErrorJson
		body.Decode(&errorJson)
		return &errorJson
	}

	return nil
}

// TODO batch requests
func (this *Api) ResolveNameByUid(uid uint64) string {
	params := &map[string]string{
		"user_ids": fmt.Sprintf("%d", uid),
		"fields":   "first_name,last_name",
	}

	response := this.request("users.get", params)
	if response == nil {
		return ""
	}

	defer response.Body.Close()
	body := json.NewDecoder(response.Body)

	if apiErr := checkError(body); apiErr != nil {
		log.Println(apiErr)
		return ""
	}

	_, _ = body.Token()

	var userJson UserJson
	if err := body.Decode(&userJson); err != nil {
		log.Println(err)
		return ""
	}

	return userJson.FirstName + " " + userJson.LastName
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
