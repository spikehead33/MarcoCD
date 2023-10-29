package pkg

import (
	"marcocd/pkg/appcontroller"
	"marcocd/pkg/repomonitor"
	"net/http"
)

func NewServer(
	appcontroller appcontroller.AppController,
	repomonitor repomonitor.RepoMonitor) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/refresh", nil)
	mux.HandleFunc("/sync", nil)
	return mux
}

func StartServer(addr string, m *http.ServeMux) error {
	if err := http.ListenAndServe(addr, m); err != nil {
		return err
	}

	return nil
}
