package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Volume struct {
	Name         string
	Size         string
	AccessMode   string
	ReplicaCount string
	DataLocality string
	DataEngine   string
	Encrypted    bool
	DataSource   string
	Frontend     string
}

func createVolume(name string) string {
	volume := Volume{
		Name:         name,
		Size:         "1024",
		AccessMode:   "rwo",
		ReplicaCount: "3",
		DataLocality: "best-effort",
		DataEngine:   "v1",
		Encrypted:    false,
		DataSource:   "local",
		Frontend:     "blockdev",
	}

	body, err := json.Marshal(volume)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "https://longhorn.dwcl-master.k8s.system.deweloper.cloud/v1/volumes", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic YWRtaW46ZG9tZG9ta3Vyc3VudTM1XzM1")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	// return body as string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(bodyBytes)
}

func main() {
	fmt.Println(createVolume("test"))
}
