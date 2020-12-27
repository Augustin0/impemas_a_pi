package impemas_a_pi

import (
	"context"
	"log"
)

type Cuenta struct {
	Banco  string `json:"banco"`
	Color  string `json:"color"`
	Estado string `json:"estado"`
	Id     string `json:"id"`
	Numero string `json:"numero"`
}

func geCuentas(user string) map[string]Cuenta {
	ctx := context.Background()
	client, err := getApp().Database(ctx)
	if err != nil {
		panic(err)
	}
	if user != "" {
		ref := client.NewRef("ORGS/ORG_" + user + "/DataCount")
		var cuentas = make(map[string]Cuenta)
		if err := ref.Get(ctx, &cuentas); err != nil {
			panic("USUARIO NO CONECTADO :( !")
		}
		return cuentas
	} else {
		panic("USUARIO NO IDENTIFICADO  :( !")

	}

}

func putCuenta(user string, cuenta Cuenta) map[string]Cuenta {
	ctx := context.Background()
	cliente, err := getApp().Database(ctx)
	if err != nil {
		log.Printf("Error al tratar de acceder al cliente de base de datos %s", err)

	}
	ref, error := cliente.NewRef("ORGS/ORG_"+user+"/DataClient").Push(ctx, nil)
	if error != nil {
		panic(error)
	}
	cuenta.Id = ref.Key
	if err := ref.Set(ctx, &cuenta); err != nil {
		panic(err)
	}
	return geCuentas(user)

}
