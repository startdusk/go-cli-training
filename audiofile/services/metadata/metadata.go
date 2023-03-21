package metadata

import (
	"audiofile/internal/interfaces"
	"audiofile/storage"
	"fmt"
	"net/http"
)

type MetadataService struct {
	Server  *http.Server
	Storage interfaces.Storage
}

func CreateMetadataService(port int, storage interfaces.Storage) *MetadataService {
	mux := http.NewServeMux()
	srv := &MetadataService{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
		Storage: storage,
	}
	mux.HandleFunc("/upload", srv.upload)
	mux.HandleFunc("/request", srv.getByID)
	mux.HandleFunc("/list", srv.list)
	return srv
}

func Run(port int) error {
	m := CreateMetadataService(port, storage.FlatFile{})
	return m.Server.ListenAndServe()
}
