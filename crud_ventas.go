package impemas_a_pi

import (
	"context"
	"log"
)

type Venta struct {
	ITBIS            int           `json:"ITBIS"`
	Cliente          string        `json:"cliente"`
	Deposito         interface{}   `json:"deposito"`
	Depositos        []interface{} `json:"depositos"`
	Fecha            string        `json:"fecha"`
	FechaVencimiento string        `json:"fechaVencimiento"`
	Id               string        `json:"id"`
	Numero           string        `json:"numero"`
	Rnc              string        `json:"rnc"`
	SubTotal         string        `json:"subTotal"`
	Tipo             string        `json:"tipo"`
	Total            string        `json:"total"`
}

func getVentas(user string) map[string]map[string]Venta {
	ctx := context.Background()
	cliente, err := getApp().Database(ctx)
	if err != nil {
		log.Printf("Error al tratar de acceder al cliente de base de datos %s", err)
	}
	if user != "" {
		ref := cliente.NewRef("ORGS/ORG_" + user + "/Ventas")
		var ventas map[string]map[string]Venta
		if err := ref.Get(ctx, &ventas); err != nil {
			log.Printf("No existe el Cliente %s no existe", user)
		}
		return ventas
	} else {
		panic(`User Required !! ":(`)
	}

}
