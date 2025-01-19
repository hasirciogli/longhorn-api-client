package apis

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

type Volume struct {
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

func CreateVolume(volume Volume) mai {
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

	CreateVolume()
}
