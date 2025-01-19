# Longhorn API Client Library

## Introduction
This library serves as a sample implementation for interacting with the Longhorn API. It is recommended to write your own code while using this library, as it is intended for demonstration purposes only. However, if configured correctly, it should work without errors. I believe that improvements in error handling are necessary, as this is absolutely essential.

## Example: Creating a Volume

```go
package main

import (
	"fmt"

	"github.com/hasirciogli/longhorn-api-client/apis"
	"github.com/hasirciogli/longhorn-api-client/models"
	"k8s.io/apimachinery/pkg/util/uuid"
)

func main() {
	myuuid := uuid.NewUUID()
	basicAuth := "example http authed user:password (base64 encoded)"
	config := models.NewConfig("your web endpoint", &basicAuth)

	res := apis.CreateVolume(apis.VolumeCreateRequest{
		Name:             string("pvc-" + myuuid),
		Size:             "1Gi",
		NumberOfReplicas: 3,
		Frontend:         "blockdev",
	}, config)
	if res.Error != nil {
		fmt.Println(res.Error.Message)
		return
	}

	fmt.Println(res.Message)
}
```

By correctly setting the endpoint and basic auth information as shown in the code above, it will work without errors. If there is no authentication and you are running internally, simply pass `nil` as a parameter. That will handle it :D

## Example: Listing Volumes

```go
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
```

Currently, these are the functions available. I created this library to meet the needs of my company. If necessary, I can add more functions. Feel free to contribute with commits as well.