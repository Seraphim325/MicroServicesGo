package handler

import (
	"broker/cmd/api/model"
	"encoding/json"
	"net/http"
)

type Broker struct {
}

func (b *Broker) Handler(w http.ResponseWriter, r *http.Request) {
	payload := model.JsonResponse{
		Message: "Hit the broker",
		Error:   false,
	}

	out, _ := json.MarshalIndent(payload, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)
}
