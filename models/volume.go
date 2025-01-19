package models

type VolumesResponse struct {
	CreateTypes map[string]string `json:"createTypes"`

	Data []struct {
		AccessMode string `json:"accessMode"`
		// Actions    []struct {
		// 	Activate                          string `json:"activate"`
		// 	Attach                            string `json:"attach"`
		// 	CancelExpansion                   string `json:"cancelExpansion"`
		// 	Detach                            string `json:"detach"`
		// 	EngineUpgrade                     string `json:"engineUpgrade"`
		// 	Expand                            string `json:"expand"`
		// 	PvCreate                          string `json:"pvCreate"`
		// 	RecurringJobAdd                   string `json:"recurringJobAdd"`
		// 	RecurringJobDelete                string `json:"recurringJobDelete"`
		// 	RecurringJobList                  string `json:"recurringJobList"`
		// 	ReplicaRemove                     string `json:"replicaRemove"`
		// 	SnapshotBackup                    string `json:"snapshotBackup"`
		// 	SnapshotCRCreate                  string `json:"snapshotCRCreate"`
		// 	SnapshotCRDelete                  string `json:"snapshotCRDelete"`
		// 	SnapshotCRGet                     string `json:"snapshotCRGet"`
		// 	SnapshotCRList                    string `json:"snapshotCRList"`
		// 	SnapshotCreate                    string `json:"snapshotCreate"`
		// 	SnapshotDelete                    string `json:"snapshotDelete"`
		// 	SnapshotGet                       string `json:"snapshotGet"`
		// 	SnapshotList                      string `json:"snapshotList"`
		// 	SnapshotPurge                     string `json:"snapshotPurge"`
		// 	SnapshotRevert                    string `json:"snapshotRevert"`
		// 	TrimFilesystem                    string `json:"trimFilesystem"`
		// 	UpdateBackupCompressionMethod     string `json:"updateBackupCompressionMethod"`
		// 	UpdateDataLocality                string `json:"updateDataLocality"`
		// 	UpdateFreezeFilesystemForSnapshot string `json:"updateFreezeFilesystemForSnapshot"`
		// 	UpdateReplicaAutoBalance          string `json:"updateReplicaAutoBalance"`
		// 	UpdateReplicaCount                string `json:"updateReplicaCount"`
		// 	UpdateReplicaDiskSoftAntiAffinity string `json:"updateReplicaDiskSoftAntiAffinity"`
		// 	UpdateReplicaSoftAntiAffinity     string `json:"updateReplicaSoftAntiAffinity"`
		// 	UpdateReplicaZoneSoftAntiAffinity string `json:"updateReplicaZoneSoftAntiAffinity"`
		// 	UpdateSnapshotDataIntegrity       string `json:"updateSnapshotDataIntegrity"`
		// 	UpdateSnapshotMaxCount            string `json:"updateSnapshotMaxCount"`
		// 	UpdateSnapshotMaxSize             string `json:"updateSnapshotMaxSize"`
		// 	UpdateUnmapMarkSnapChainRemoved   string `json:"updateUnmapMarkSnapChainRemoved"`
		// } `json:"actions"`
		BackingImage            string `json:"backingImage"`
		BackupCompressionMethod string `json:"backupCompressionMethod"`
		BackupStatus            []struct {
			Actions   interface{} `json:"actions"`
			BackupURL string      `json:"backupURL"`
			Error     string      `json:"error"`
			Id        string      `json:"id"`
			Links     interface{} `json:"links"`
			Progress  int         `json:"progress"`
			Replica   string      `json:"replica"`
			Size      string      `json:"size"`
			Snapshot  string      `json:"snapshot"`
			State     string      `json:"state"`
		} `json: "backupStatus"`
		CloneStatus struct {
			AttemptCount         int    `json:"attemptCount"`
			NextAllowedAttemptAt string `json:"nextAllowedAttemptAt"`
			Snapshot             string `json:"snapshot"`
			SourceVolume         string `json:"sourceVolume"`
			State                string `json:"state"`
		} `json: "cloneStatus"`

		Controllers []struct {
			ActualSize                       string `json:"actualSize"`
			Address                          string `json:"address"`
			CurrentImage                     string `json:"currentImage"`
			Endpoint                         string `json:"endpoint"`
			HostId                           string `json:"hostId"`
			Image                            string `json:"image"`
			InstanceManagerName              string `json:"instanceManagerName"`
			IsExpanding                      bool   `json:"isExpanding"`
			LastExpansionError               string `json:"lastExpansionError"`
			LastExpansionFailedAt            string `json:"lastExpansionFailedAt"`
			LastRestoredBackup               string `json:"lastRestoredBackup"`
			Name                             string `json:"name"`
			RequestedBackupRestore           string `json:"requestedBackupRestore"`
			Running                          bool   `json:"running"`
			Size                             string `json:"size"`
			UnmapMarkSnapChainRemovedEnabled bool   `json:"unmapMarkSnapChainRemovedEnabled"`
		} `json:"controllers"`

		Created                     string   `json:"created"`
		CurrentImage                string   `json:"currentImage"`
		DataEngine                  string   `json:"dataEngine"`
		DataLocality                string   `json:"dataLocality"`
		DataSource                  string   `json:"dataSource"`
		DisableFrontend             bool     `json:"disableFrontend"`
		DiskSelector                []string `json:"diskSelector"`
		Encrypted                   bool     `json:"encrypted"`
		FreezeFilesystemForSnapshot string   `json:"freezeFilesystemForSnapshot"`
		FromBackup                  string   `json:"fromBackup"`
		Frontend                    string   `json:"frontend"`
		Id                          string   `json:"id"`
		Image                       string   `json:"image"`

		KubernetesStatus struct {
			LastPVCRefAt string `json:"lastPVCRefAt"`
			LastPodRefAt string `json:"lastPodRefAt"`
			Namespace    string `json:"namespace"`
			PvName       string `json:"pvName"`
			PvStatus     string `json:"pvStatus"`
			PvcName      string `json:"pvcName"`

			WorkloadsStatus []struct {
				PodName      string `json:"podName"`
				PodStatus    string `json:"podStatus"`
				WorkloadName string `json:"workloadName"`
				WorkloadType string `json:"workloadType"`
			} `json:"workloadsStatus"`
		} `json:"kubernetesStatus"`

		LastAttachedBy   string `json:"lastAttachedBy"`
		LastBackup       string `json:"lastBackup"`
		LastBackupAt     string `json:"lastBackupAt"`
		Migratable       bool   `json:"migratable"`
		Name             string `json:"name"`
		NodeSelector     []any  `json:"nodeSelector"`
		NumberOfReplicas int    `json:"numberOfReplicas"`

		PurgeStatus []struct {
			Actions   interface{} `json:"actions"`
			Error     string      `json:"error"`
			IsPurging bool        `json:"isPurging"`
			Links     interface{} `json:"links"`
			Progress  int         `json:"progress"`
			Replica   string      `json:"replica"`
			State     string      `json:"state"`
		} `json:"purgeStatus"`

		Ready         bool `json:"ready"`
		RebuildStatus []struct {
			Actions      interface{} `json:"actions"`
			Error        string      `json:"error"`
			IsRebuilding bool        `json:"isRebuilding"`
			Links        interface{} `json:"links"`
			Progress     int         `json:"progress"`
			Replica      string      `json:"replica"`
			State        string      `json:"state"`
		} `json:"rebuildStatus"`
		RecurringJobSelector        interface{} `json:"recurringJobSelector"`
		ReplicaAutoBalance          string      `json:"replicaAutoBalance"`
		ReplicaDiskSoftAntiAffinity string      `json:"replicaDiskSoftAntiAffinity"`
		ReplicaSoftAntiAffinity     string      `json:"replicaSoftAntiAffinity"`
		ReplicaZoneSoftAntiAffinity string      `json:"replicaZoneSoftAntiAffinity"`

		Replicas []struct {
			Address             string `json:"address"`
			CurrentImage        string `json:"currentImage"`
			DataEngine          string `json:"dataEngine"`
			DataPath            string `json:"dataPath"`
			DiskID              string `json:"diskID"`
			DiskPath            string `json:"diskPath"`
			FailedAt            string `json:"failedAt"`
			HostId              string `json:"hostId"`
			Image               string `json:"image"`
			InstanceManagerName string `json:"instanceManagerName"`
			Mode                string `json:"mode"`
			Name                string `json:"name"`
			Running             bool   `json:"running"`
		} `json:"replicas"`

		RestoreInitiated bool `json:"restoreInitiated"`
		RestoreRequired  bool `json:"restoreRequired"`

		RestoreStatus []struct {
			Actions      interface{} `json:"actions"`
			BackupURL    string      `json:"backupURL"`
			Error        string      `json:"error"`
			Filename     string      `json:"filename"`
			IsRestoring  bool        `json:"isRestoring"`
			LastRestored string      `json:"lastRestored"`
			Links        interface{} `json:"links"`
			Progress     int         `json:"progress"`
			Replica      string      `json:"replica"`
			State        string      `json:"state"`
		} `json:"restoreStatus"`

		RestoreVolumeRecurringJob string `json:"restoreVolumeRecurringJob"`
		RevisionCounterDisabled   bool   `json:"revisionCounterDisabled"`
		Robustness                string `json:"robustness"`
		ShareEndpoint             string `json:"shareEndpoint"`
		ShareState                string `json:"shareState"`
		Size                      string `json:"size"`
		SnapshotDataIntegrity     string `json:"snapshotDataIntegrity"`
		SnapshotMaxCount          int    `json:"snapshotMaxCount"`
		SnapshotMaxSize           string `json:"snapshotMaxSize"`
		StaleReplicaTimeout       int    `json:"staleReplicaTimeout"`
		Standby                   bool   `json:"standby"`
		State                     string `json:"state"`
		Type                      string `json:"type"`
		UnmapMarkSnapChainRemoved string `json:"unmapMarkSnapChainRemoved"`

		VolumeAttachment struct {
			Attachments map[string]struct {
				AttachmentID   string `json:"attachmentID"`
				AttachmentType string `json:"attachmentType"`
				Conditions     []struct {
					LastProbeTime      string `json:"lastProbeTime"`
					LastTransitionTime string `json:"lastTransitionTime"`
					Message            string `json:"message"`
					Reason             string `json:"reason"`
					Status             string `json:"status"`
					Type               string `json:"type"`
				} `json:"conditions"`
				NodeID     string            `json:"nodeID"`
				Parameters map[string]string `json:"parameters"`
				Satisfied  bool              `json:"satisfied"`
			} `json:"attachments"`
			Volume string `json:"volume"`
		} `json:"volumeAttachment"`
	} `json:"data"`
	ResourceType string `json:"resourceType"`
	Type         string `json:"type"`
}
