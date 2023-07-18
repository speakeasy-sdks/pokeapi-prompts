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
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	LocationAreaReadDefaultTextPlainString *string
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

func (o *LocationAreaReadResponse) GetLocationAreaReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.LocationAreaReadDefaultTextPlainString
}
