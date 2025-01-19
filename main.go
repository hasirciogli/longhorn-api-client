package main

import (
	"fmt"

	"github.com/hasirciogli/longhorn-api-client/apis"
	"github.com/hasirciogli/longhorn-api-client/models"
)

func main() {
	basicAuth := "example http authed user:password (base64 encoded)"
	config := models.NewConfig("your web endpoint", &basicAuth)

	res := apis.GetVolumes(config)

	if res.Error != nil {
		fmt.Println(res.Error.Message)
		return
	}

	fmt.Println(res.Volumes.Data[0].Name)
}
