package ambassador

import (
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
	return config.HandleResponse(resp, err)
}
