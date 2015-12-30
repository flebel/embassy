package ambassador

import (
	"io/ioutil"
	"net/http"

	"github.com/flebel/embassy/ambassadors"
)

var Name = "generic"

func Perform(ambassador config.Ambassador) (int, string, []byte, error) {
	fetch := config.HTTPVerbFunctionMap[ambassador.HTTPVerb].(func(string) (*http.Response, error))
	resp, err := fetch(ambassador.Url)
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
