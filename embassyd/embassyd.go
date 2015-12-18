package embassyd

import (
	"net/http"

	jsonip "github.com/flebel/embassy/ambassadors/jsonip"
	"github.com/gorilla/mux"
)

func JsonIpHandler(w http.ResponseWriter, r *http.Request) {
	statusCode, contentType, body, err := jsonip.Perform()
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", contentType)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
	} else {
		w.Write(body)
	}
}

func StartNewEmbassyD() {
	r := mux.NewRouter()
	r.HandleFunc("/external_ip", JsonIpHandler)
	http.ListenAndServe("localhost:8000", r)
}
