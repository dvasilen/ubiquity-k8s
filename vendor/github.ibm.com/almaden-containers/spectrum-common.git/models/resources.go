package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateRequest struct {
	Name string
	Opts map[string]interface{}
}

type ActivateResponse struct {
	Implements []string
}

func (r *ActivateResponse) WriteResponse(w http.ResponseWriter) {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Error marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(data))
}

type GenericResponse struct {
	Err string
}

func (r *GenericResponse) WriteResponse(w http.ResponseWriter) {
	if r.Err != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Error marshalling response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(data))
}

type GenericRequest struct {
	Name string
}

type MountResponse struct {
	Mountpoint string
	Err        string
}

func (r *MountResponse) WriteResponse(w http.ResponseWriter) {
	if r.Err != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Error marshalling Get response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(data))
}

type VolumeMetadata struct {
	Name       string
	Mountpoint string
}

type SpectrumConfig struct {
	FilesetId  string `json:"fileset"`
	Filesystem string `json:"filesystem"`
}
type GetResponse struct {
	Volume VolumeMetadata
	Err    string
}

func (r *GetResponse) WriteResponse(w http.ResponseWriter) {
	if r.Err != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Error marshalling Get response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(data))
}

type ListResponse struct {
	Volumes []VolumeMetadata
	Err     string
}

func (r *ListResponse) WriteResponse(w http.ResponseWriter) {
	if r.Err != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Errorf("Error marshalling Get response: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(data))
}

type FlexVolumeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Device  string `json:"device"`
}

type FlexVolumeMountOptions struct {
	MountPath   string `json:"mount_path"`
	MountDevice string `json:"mount_device"`
	FSType      string `json:"fstype"`
}

type FlexVolumeAttachOptions struct {
	VolumeId    string `json:"volume_id"`
	Size        int    `json:"size"`
	VolumeGroup string `json:"volume_group"`
	FileSet     string `json:"fileset"`
	Path        string `json:"path"`
}