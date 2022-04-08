package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/exp/slices"
)

const (
	contentType     string = "Content-Type"
	applicationJson string = "application/json"
)

func ParseRequestAs[TPayload any](r *http.Request) (payload TPayload, err error) {

	hct := r.Header[contentType]
	if _, ok := slices.BinarySearch(hct, applicationJson); !ok {
		err = fmt.Errorf("invalid request %s (expected %q)", contentType, applicationJson)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err = decoder.Decode(&payload)
	if err == io.EOF {
		err = errors.New("Missing request body")
		return
	}

	if e, ok := err.(*json.SyntaxError); ok {
		err = fmt.Errorf("%s (position: %d)", e, e.Offset)
	}

	return
}
