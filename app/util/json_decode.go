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

func JSONDecoderProduct(body io.ReadCloser, structBody model.Product) (model.Product, error) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&structBody)
	if err != nil {
		return model.Product{}, err
	}

	return structBody, err
}

func JSONDecoderOrder(body io.ReadCloser, structBody model.Order) (model.Order, error) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&structBody)
	if err != nil {
		return model.Order{}, err
	}

	return structBody, err
}

func JSONDecoderOderDetail(body io.ReadCloser, structBody model.OderDetail) (model.OderDetail, error) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&structBody)
	if err != nil {
		return model.OderDetail{}, err
	}

	return structBody, err
}
