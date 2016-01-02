package ambassador

import (
	"io/ioutil"
	"net/http"
)

const Name = "jsonip"

func Perform() (int, string, []byte, error) {
	resp, err := http.Get("http://jsonip.com")
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
