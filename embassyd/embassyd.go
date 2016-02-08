package embassyd

import (
	"fmt"
	"net/http"

	"github.com/flebel/embassy/ambassadors"
	generic "github.com/flebel/embassy/ambassadors/generic"
	jsonip "github.com/flebel/embassy/ambassadors/jsonip"
	ping "github.com/flebel/embassy/ambassadors/ping"
	"github.com/gorilla/mux"
)

type simplePerformer func() (int, string, []byte, error)

var performers = map[string]simplePerformer{
	jsonip.Name: jsonip.Perform,
	ping.Name:   ping.Perform,
}

var errorHandler = func(w http.ResponseWriter) {
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

func GenericHandler(ambassador config.Ambassador) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errorHandler(w)

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

func SimpleHandler(perform simplePerformer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer errorHandler(w)

		statusCode, contentType, body, err := perform()
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
		if ambassador.Ambassador == generic.Name {
			handler = GenericHandler(ambassador)
		} else {
			handler = SimpleHandler(performers[ambassador.Ambassador])
		}
		if handler != nil {
			r.HandleFunc(ambassador.Path, handler.(func(w http.ResponseWriter, r *http.Request))).Methods(ambassador.HTTPVerb)
		}
	}
	http.ListenAndServe(listen, r)
}
