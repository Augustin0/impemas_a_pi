package impemas_a_pi

import (
	"context"
)

type Cliente struct {
	Id        string `json:"id"`
	Nombre    string `json:"nombre"`
	Telefono  string `json:"telefono"`
	Direccion string `json:"direccion"`
	Rnc       string `json:"rnc"`
	Color     string `json:"color"`
	Estado    string `json:"estado"`
}

func getClientes(user string) map[string]Cliente {
	ctx := context.Background()

	cliente, err := getApp().Database(ctx)
	if err != nil {
		panic(err)
	}
	if user != "" {
		ref := cliente.NewRef("ORGS/ORG_" + user + "/DataClient")
		var clientes map[string]Cliente
		if err := ref.Get(ctx, &clientes); err != nil {
			panic(`PROBABLEMENTE NO ESTAS CONECTADO A INTERNET "(:(  `)
		}
		return clientes
	} else {
		panic(`USUARIO NO IDENTIFICADO "(:(      `)
	}

}
