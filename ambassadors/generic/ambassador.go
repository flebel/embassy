package ambassador

import (
	"io/ioutil"
	"net/http"

	"github.com/flebel/embassy/ambassadors"
)

const Name = "generic"

type Configuration struct {
	URL      string
	HTTPVerb string
}

func Perform(amb config.Ambassador) (int, string, []byte, error) {
	conf := Configuration{}
	config.ParseConfiguration(amb.Configuration, &conf)

	fetch := config.HTTPVerbFunctionMap[conf.HTTPVerb].(func(string) (*http.Response, error))
	resp, err := fetch(conf.URL)
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
