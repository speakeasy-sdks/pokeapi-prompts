// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EggGroupReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *EggGroupReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type EggGroupReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	EggGroupReadDefaultTextPlainString *string
}

func (o *EggGroupReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EggGroupReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EggGroupReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EggGroupReadResponse) GetEggGroupReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.EggGroupReadDefaultTextPlainString
}
