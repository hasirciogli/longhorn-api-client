package apis

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hasirciogli/longhorn-api-client/models"
	"github.com/hasirciogli/longhorn-api-client/utils"
)

/**
{
    "name": "22",
    "size": "1Gi",
    "numberOfReplicas": 3,
    "frontend": "blockdev",
    "dataLocality": "best-effort",
    "accessMode": "rwo",
    "backingImage": "",
    "dataSourceType": "",
    "dataEngine": "v1",
    "encrypted": false,
    "nodeSelector": [],
    "diskSelector": [],
    "snapshotDataIntegrity": "ignored",
    "snapshotMaxCount": 0,
    "snapshotMaxSize": "0Gi",
    "snapshotSizeUnit": "Gi",
    "replicaAutoBalance": "ignored",
    "unmapMarkSnapChainRemoved": "ignored",
    "replicaSoftAntiAffinity": "ignored",
    "replicaZoneSoftAntiAffinity": "ignored",
    "replicaDiskSoftAntiAffinity": "ignored",
    "freezeFilesystemForSnapshot": "ignored",
    "revisionCounterDisabled": true,
    "dataSource": "",
    "staleReplicaTimeout": 20,
    "fromBackup": ""
}
*/

type VolumeCreateRequest struct {
	Name                        string   `json:"name"`
	Size                        string   `json:"size"`
	NumberOfReplicas            int      `json:"numberOfReplicas"`
	Frontend                    string   `json:"frontend"`
	DataLocality                string   `json:"dataLocality"`
	AccessMode                  string   `json:"accessMode"`
	BackingImage                string   `json:"backingImage"`
	DataSourceType              string   `json:"dataSourceType"`
	DataEngine                  string   `json:"dataEngine"`
	Encrypted                   bool     `json:"encrypted"`
	NodeSelector                []string `json:"nodeSelector"`
	DiskSelector                []string `json:"diskSelector"`
	SnapshotDataIntegrity       string   `json:"snapshotDataIntegrity"`
	SnapshotMaxCount            int      `json:"snapshotMaxCount"`
	SnapshotMaxSize             string   `json:"snapshotMaxSize"`
	SnapshotSizeUnit            string   `json:"snapshotSizeUnit"`
	ReplicaAutoBalance          string   `json:"replicaAutoBalance"`
	UnmapMarkSnapChainRemoved   string   `json:"unmapMarkSnapChainRemoved"`
	ReplicaSoftAntiAffinity     string   `json:"replicaSoftAntiAffinity"`
	ReplicaZoneSoftAntiAffinity string   `json:"replicaZoneSoftAntiAffinity"`
	ReplicaDiskSoftAntiAffinity string   `json:"replicaDiskSoftAntiAffinity"`
	FreezeFilesystemForSnapshot string   `json:"freezeFilesystemForSnapshot"`
	RevisionCounterDisabled     bool     `json:"revisionCounterDisabled"`
	DataSource                  string   `json:"dataSource"`
	StaleReplicaTimeout         int      `json:"staleReplicaTimeout"`
	FromBackup                  string   `json:"fromBackup"`
}

type VolumeCreateResponse struct {
	Error *struct {
		Message    string         `json:"message"`
		StatusCode int            `json:"statusCode"`
		Status     bool           `json:"status"`
		Response   *http.Response `json:"response"`
	} `json:"error"`
	Response   *http.Response `json:"response"`
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Status     bool           `json:"status"`
}

func CreateVolume(volume VolumeCreateRequest, sConfig *models.SConfig) *VolumeCreateResponse {
	// Set default values for empty fields
	if volume.Name == "" {
		volume.Name = "default-name" // Set a default name
	}
	if volume.Size == "" {
		volume.Size = "1Gi" // Set a default size
	}
	if volume.NumberOfReplicas == 0 {
		volume.NumberOfReplicas = 3 // Set a default number of replicas
	}
	if volume.Frontend == "" {
		volume.Frontend = "blockdev" // Set a default frontend
	}
	if volume.DataLocality == "" {
		volume.DataLocality = "best-effort" // Set a default data locality
	}
	if volume.AccessMode == "" {
		volume.AccessMode = "rwo" // Set a default access mode
	}
	if volume.DataEngine == "" {
		volume.DataEngine = "v1" // Set a default data engine
	}

	if volume.SnapshotDataIntegrity == "" {
		volume.SnapshotDataIntegrity = "ignored" // Set a default snapshot data integrity
	}

	if volume.SnapshotMaxCount == 0 {
		volume.SnapshotMaxCount = 0 // Set a default snapshot max count
	}

	if volume.SnapshotMaxSize == "" {
		volume.SnapshotMaxSize = "0Gi" // Set a default snapshot max size
	}

	if volume.SnapshotSizeUnit == "" {
		volume.SnapshotSizeUnit = "Gi" // Set a default snapshot size unit
	}

	body, err1 := json.Marshal(volume)
	if err1 != nil {
		return &VolumeCreateResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err1.Error(),
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Message:    err1.Error(),
			Status:     false,
			StatusCode: 500,
			Response:   nil,
		}
	}

	resp, err := utils.Requester("POST", "/v1/volumes", body, sConfig)
	if err != nil {
		return &VolumeCreateResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err.Message,
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Message:    err.Message,
			Status:     false,
			StatusCode: 500,
			Response:   nil,
		}
	}

	if resp.StatusCode > 299 {
		return &VolumeCreateResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    "Volume creation failed",
				StatusCode: resp.StatusCode,
				Status:     false,
				Response:   resp,
			},
			Message:    "Volume creation failed",
			Status:     false,
			StatusCode: resp.StatusCode,
			Response:   resp,
		}
	}

	return &VolumeCreateResponse{
		Message:    "Volume created successfully",
		Status:     true,
		StatusCode: resp.StatusCode,
		Response:   resp,
	}
}

type VolumesResponse struct {
	Error *struct {
		Message    string         `json:"message"`
		StatusCode int            `json:"statusCode"`
		Status     bool           `json:"status"`
		Response   *http.Response `json:"response"`
	} `json:"error"`
	Volumes *models.VolumesResponse `json:"volumes"`
}

func GetVolumes(sConfig *models.SConfig) *VolumesResponse {
	resp, err := utils.Requester("GET", "/v1/volumes", nil, sConfig)
	if err != nil {
		return &VolumesResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err.Message,
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Volumes: nil,
		}
	}

	if resp.StatusCode > 299 {
		return &VolumesResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    "Volume list failed",
				StatusCode: resp.StatusCode,
				Status:     false,
				Response:   resp,
			},
			Volumes: nil,
		}
	}

	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return &VolumesResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err2.Error(),
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Volumes: nil,
		}
	}

	var volumes *models.VolumesResponse
	err3 := json.Unmarshal(body, &volumes)
	if err3 != nil {
		return &VolumesResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err3.Error(),
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Volumes: nil,
		}
	}

	return &VolumesResponse{
		Error:   nil,
		Volumes: volumes,
	}
}

type VolumeDeleteRequest struct {
	VolumeName string `json:"volumeName"`
}

type VolumeDeleteResponse struct {
	Error *struct {
		Message    string         `json:"message"`
		StatusCode int            `json:"statusCode"`
		Status     bool           `json:"status"`
		Response   *http.Response `json:"response"`
	} `json:"error"`
	Message    string         `json:"message"`
	Status     bool           `json:"status"`
	StatusCode int            `json:"statusCode"`
	Response   *http.Response `json:"response"`
}

func DeleteVolume(volume VolumeDeleteRequest, sConfig *models.SConfig) *VolumeDeleteResponse {
	resp, err := utils.Requester("DELETE", "/v1/volumes/"+volume.VolumeName, nil, sConfig)
	if err != nil {
		return &VolumeDeleteResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err.Message,
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
			Message:    err.Message,
			Status:     false,
			StatusCode: 500,
			Response:   nil,
		}
	}

	if resp.StatusCode > 299 {
		return &VolumeDeleteResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    "Volume deletion failed",
				StatusCode: resp.StatusCode,
				Status:     false,
				Response:   resp,
			},
			Message:    "Volume deletion failed",
			Status:     false,
			StatusCode: resp.StatusCode,
			Response:   resp,
		}
	}

	return &VolumeDeleteResponse{
		Error:      nil,
		Message:    "Volume deleted successfully",
		Status:     true,
		StatusCode: resp.StatusCode,
		Response:   resp,
	}
}

type VolumeInspectRequest struct {
	VolumeName string `json:"volumeName"`
}

type VolumeInspectResponse struct {
	Error *struct {
		Message    string         `json:"message"`
		StatusCode int            `json:"statusCode"`
		Status     bool           `json:"status"`
		Response   *http.Response `json:"response"`
	} `json:"error"`
	Volume *models.VolumeInspectResponse `json:"volume"`
}

func InspectVolume(volume VolumeInspectRequest, sConfig *models.SConfig) *VolumeInspectResponse {
	resp, err := utils.Requester("GET", "/v1/volumes/"+volume.VolumeName, nil, sConfig)
	if err != nil {
		return &VolumeInspectResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err.Message,
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
		}
	}

	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return &VolumeInspectResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err2.Error(),
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
		}
	}

	var inspectedVolume *models.VolumeInspectResponse
	err3 := json.Unmarshal(body, &inspectedVolume)
	if err3 != nil {
		return &VolumeInspectResponse{
			Error: &struct {
				Message    string         `json:"message"`
				StatusCode int            `json:"statusCode"`
				Status     bool           `json:"status"`
				Response   *http.Response `json:"response"`
			}{
				Message:    err3.Error(),
				StatusCode: 500,
				Status:     false,
				Response:   nil,
			},
		}
	}

	return &VolumeInspectResponse{
		Error:  nil,
		Volume: inspectedVolume,
	}
}
