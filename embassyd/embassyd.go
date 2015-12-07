package embassyd

import (
	"net/http"

	"github.com/flebel/embassy/ambassadors/url"
	"github.com/gorilla/mux"
)

func JsonIpHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(ambassador.Name + "\n"))
}

func StartNewEmbassyD() {
	r := mux.NewRouter()
	r.HandleFunc("/jsonip", JsonIpHandler)
	http.ListenAndServe(":8000", r)
}
