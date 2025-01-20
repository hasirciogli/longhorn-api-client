package tests

import (
	"fmt"

	"github.com/hasirciogli/longhorn-api-client/apis"
	"github.com/hasirciogli/longhorn-api-client/models"
)

type Tests struct {
	BasicAuth *string
	Endpoint  string
}

func (t *Tests) TestCreateVolume() {
	config := models.NewConfig(t.Endpoint, t.BasicAuth)

	response := apis.CreateVolume(apis.VolumeCreateRequest{
		Name:             "pvc-xxxxx",
		Size:             "1Gi",
		NumberOfReplicas: 3,
		Frontend:         "blockdev",
	}, config)

	if response.Error != nil {
		fmt.Println("Volume creation failed: %v", response.Error)
	}

	fmt.Println(response.Message)
}

func (t *Tests) TestInspectVolume() {
	config := models.NewConfig(t.Endpoint, t.BasicAuth)

	response := apis.InspectVolume(apis.VolumeInspectRequest{
		VolumeName: "pvc-xxxx",
	}, config)

	if response.Error != nil {
		fmt.Println("Volume inspection failed: %v", response.Error.Message)
	}

	fmt.Println(response.Volume.Controllers[0].ActualSize)
}

func (t *Tests) TestDeleteVolume() {
	config := models.NewConfig(t.Endpoint, t.BasicAuth)

	response := apis.DeleteVolume(apis.VolumeDeleteRequest{
		VolumeName: "pvc-xxxx",
	}, config)

	if response.Error != nil {
		fmt.Println("Volume deletion failed: %v", response.Error.Message)
	}

	fmt.Println(response.Message)
}

func (t *Tests) TestListVolumes() {
	config := models.NewConfig(t.Endpoint, t.BasicAuth)

	response := apis.GetVolumes(config)

	if response.Error != nil {
		fmt.Println("Volume listing failed: %v", response.Error.Message)
	}

	fmt.Println(response.Volumes)
}
