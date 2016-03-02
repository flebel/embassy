package ambassador

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/flebel/embassy/ambassadors"
)

const Name = "pushover"

type Configuration struct {
	Token   string
	User    string
	Message string
}

func Perform(amb config.Ambassador) (int, string, []byte, error) {
	conf := Configuration{}
	config.ParseConfiguration(amb.Configuration, &conf)

	resp, err := http.PostForm("https://api.pushover.net/1/messages.json",
		url.Values{"token": {conf.Token}, "user": {conf.User}, "message": {conf.Message}})
	defer resp.Body.Close()
	if err != nil {
		return http.StatusInternalServerError, "", nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, "", nil, err
	}
	return resp.StatusCode, resp.Header["Content-Type"][0], body, nil
}
