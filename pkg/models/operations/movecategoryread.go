// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MoveCategoryReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *MoveCategoryReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type MoveCategoryReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	Res *string
}

func (o *MoveCategoryReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveCategoryReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveCategoryReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveCategoryReadResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
