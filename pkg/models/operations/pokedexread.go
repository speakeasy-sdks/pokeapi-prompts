// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PokedexReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PokedexReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type PokedexReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	PokedexReadDefaultTextPlainString *string
}

func (o *PokedexReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokedexReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokedexReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PokedexReadResponse) GetPokedexReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.PokedexReadDefaultTextPlainString
}
