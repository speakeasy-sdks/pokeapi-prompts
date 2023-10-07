// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AbilityReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AbilityReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type AbilityReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	AbilityReadDefaultTextPlainString *string
}

func (o *AbilityReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AbilityReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AbilityReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AbilityReadResponse) GetAbilityReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.AbilityReadDefaultTextPlainString
}
