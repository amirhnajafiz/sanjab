package http

import (
	"encoding/json"
	"net/http"

	"github.com/amirhnajafiz/sanjab/internal/worker"
)

type Handler struct {
	AppMetrics Metrics
	Workers    []worker.Worker
}

type detail struct {
	Resource string `json:"resource"`
	Status   string `json:"status"`
}

// Health status return ok if the service is up
func (h Handler) Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Metrics returns the metrics data
func (h Handler) Metrics(w http.ResponseWriter, _ *http.Request) {
	bytes, err := json.Marshal(h.AppMetrics.Pull())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	_, _ = w.Write(bytes)
}

// Worker returns a status of internal workers
func (h Handler) Worker(w http.ResponseWriter, _ *http.Request) {
	var details []detail

	for _, wo := range h.Workers {
		details = append(details, detail{
			Resource: wo.GetResource(),
			Status:   wo.GetStatus(),
		})
	}

	bytes, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
	}

	_, _ = w.Write(bytes)
}
