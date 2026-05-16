package helper

import (
	"broker/cmd/api/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const MAX_BYTES int = 1024 * 1024 // one megabyte

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {

	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_BYTES))

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(data); err != nil {
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("Body must have only single JSON value")
	}

	return nil
}

func WriteJson(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if _, err := w.Write(out); err != nil {
		return err
	}

	return nil
}

func ErrorJson(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	payload := model.JsonResponse{
		Error:   true,
		Message: err.Error(),
	}

	return WriteJson(w, statusCode, payload)
}
