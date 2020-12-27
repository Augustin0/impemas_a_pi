package impemas_a_pi

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
)

func getApp() *firebase.App {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://simplerep-91d30.firebaseio.com",
	}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		fmt.Printf("Error al iniciar la app %s", err)
	}
	return app
}
