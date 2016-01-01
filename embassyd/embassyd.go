package embassyd

import (
	"fmt"
	"net/http"

	"github.com/flebel/embassy/ambassadors"
	generic "github.com/flebel/embassy/ambassadors/generic"
	jsonip "github.com/flebel/embassy/ambassadors/jsonip"
	"github.com/gorilla/mux"
)

var ErrorHandler = func(w http.ResponseWriter) {
	if e := recover(); e != nil {
		var err string
		switch x := e.(type) {
		case string:
			err = x
		default:
			err = fmt.Sprintf("%v", x)
		}
		w.WriteHeader(500)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(err))
	}
}

func JsonIpHandler(w http.ResponseWriter, r *http.Request) {
	defer ErrorHandler(w)

	statusCode, contentType, body, err := jsonip.Perform()
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", contentType)
	if err != nil {
		w.Write([]byte("Error: " + err.Error()))
	} else {
		w.Write(body)
	}
}

func GenericHandler(ambassador config.Ambassador) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer ErrorHandler(w)

		statusCode, contentType, body, err := generic.Perform(ambassador)
		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", contentType)
		if err != nil {
			w.Write([]byte("Error: " + err.Error()))
		} else {
			w.Write(body)
		}
	}
}

func StartNewEmbassyD(ambassadors []config.Ambassador, listen string) {
	r := mux.NewRouter()
	for _, ambassador := range ambassadors {
		var handler interface{} = nil
		if ambassador.Ambassador == "generic" {
			handler = GenericHandler(ambassador)
		} else if ambassador.Ambassador == "jsonip" {
			handler = JsonIpHandler
		}
		if handler != nil {
			r.HandleFunc(ambassador.Path, handler.(func(w http.ResponseWriter, r *http.Request)))
		}
	}
	http.ListenAndServe(listen, r)
}
