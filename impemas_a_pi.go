package impemas_a_pi

import (

	//"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

/********************************************************************************************

 ________________________________
|                                |
|     SERVIDOR  IMPEMAS-A-PI     |
|           ○ ○ ○ ○              |
|________________________________|

**********************************************************************************************/

func services() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", Hinit).Methods("GET")
	r.HandleFunc("/apimpemasa/bancos", HgetBancos).Methods("GET")
	r.HandleFunc("/apimpemasa/cuentas-bancarias/{user}", HgetCount).Methods("GET")
	r.HandleFunc("/apimpemasa/clientes/{user}", HgetClientes).Methods("GET")
	r.HandleFunc("/apimpemasa/Ventas/{user}", HgetVentas).Methods("GET")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Println("SIRVIENDO EN EL PUERTO LOCAL :8080")
	}

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 20 << 20,
	}

	log.Println("SIRVIENDO EN EL PUERTO LOCAL :8080")
	log.Fatal(server.ListenAndServe())
}
