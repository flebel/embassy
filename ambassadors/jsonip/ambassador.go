package ambassador

import (
	"net/http"

	"github.com/flebel/embassy/ambassadors"
)

const Name = "jsonip"

func Perform(amb config.Ambassador) (int, string, []byte, error) {
	resp, err := http.Get("http://jsonip.com")
	defer resp.Body.Close()
	return config.HandleResponse(resp, err)
}
