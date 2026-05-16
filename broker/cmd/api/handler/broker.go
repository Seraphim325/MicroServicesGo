package handler

import (
	"broker/cmd/api/helper"
	"broker/cmd/api/model"
	"net/http"
)

type Broker struct {
}

func (b *Broker) Handler(w http.ResponseWriter, r *http.Request) {
	payload := model.JsonResponse{
		Message: "Hit the broker",
		Error:   false,
	}

	_ = helper.WriteJson(w, http.StatusAccepted, payload)
}
