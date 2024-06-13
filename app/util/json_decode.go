package util

import (
	"encoding/json"
	"io"

	"github.com/emanuellers/cake_store_api/model"
)

func JSONDecoderClient(body io.ReadCloser, structBody model.Client) (model.Client, error) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&structBody)
	if err != nil {
		return model.Client{}, err
	}

	return structBody, err
}
