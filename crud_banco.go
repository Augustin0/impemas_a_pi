package impemas_a_pi

import (
	"context"
	"fmt"
)

// **MODELO** BancName
type BancName struct {
	BA string `json:"ba"`
	BP string `json:"bp"`
	SB string `json:"sb"`
}

////Get BancCount
func getBanck() BancName {
	ctx := context.Background()
	client, er := getApp().Database(ctx)
	if er != nil {
		fmt.Printf("Error al tratar de acceser al cliente de base de datos :%s", er)
	}
	var bancSelect BancName
	ref := client.NewRef("BancName")
	if err := ref.Get(ctx, &bancSelect); err != nil {
		fmt.Printf("Error al tratar de leer los en el nodo BancName :%s ", err)
	}

	return bancSelect
}

// **MODELO** Cuenta Bancaria
