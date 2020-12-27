package impemas_a_pi

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Hinit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("<h1>Welconme to AP-I-MPEMASA</h1>"))

}

func HgetBancos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bancs, err := json.Marshal(getBanck())
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bancs)

}

func HgetCount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	w.Header().Set("Content-Type", "application/json")
	if user != "" {
		cuentas, err := json.Marshal(geCuentas(user))
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(cuentas)
		return
	}

	w.WriteHeader(405)
	w.Write([]byte("Usuario no espesificado"))

}

func HgetClientes(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	w.Header().Set("Content-Type", "application/json")

	if user != "" {
		clientes, err := json.Marshal(getClientes(user))

		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(clientes)
		return
	}
	w.WriteHeader(http.StatusNoContent)

	w.Write([]byte(`USER REQUIRED! ":(  `))

}

func HgetVentas(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	w.Header().Set("Content-Type", "application/json")
	if user != "" {
		ventas, err := json.Marshal(getVentas(user))
		w.WriteHeader(http.StatusOK)
		w.Write(ventas)
		createCsvCopy(getVentas(user))
		if err != nil {
			panic(err)

		}

	}

}
