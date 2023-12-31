// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LocationAreaReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *LocationAreaReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type LocationAreaReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	Res *string
}

func (o *LocationAreaReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LocationAreaReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LocationAreaReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LocationAreaReadResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
