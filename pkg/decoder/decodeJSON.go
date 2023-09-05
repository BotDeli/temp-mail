package decoder

import (
	"encoding/json"
	"io"
	"net/http"
)

func DecodeJSON[T any](response *http.Response) ([]T, error) {
	var content []T

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &content)
	return content, err
}
