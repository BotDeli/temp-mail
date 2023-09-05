package decoder

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeJSONToArray[T any](response *http.Response) ([]T, error) {
	var content []T

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &content)
	return content, err
}

func DecodeJSON[T any](response *http.Response) (T, error) {
	var content T

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return content, err
	}

	err = json.Unmarshal(bytes, &content)
	return content, err
}
