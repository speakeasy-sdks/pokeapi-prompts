// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ContestEffectReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ContestEffectReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ContestEffectReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	ContestEffectReadDefaultTextPlainString *string
}

func (o *ContestEffectReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ContestEffectReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ContestEffectReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ContestEffectReadResponse) GetContestEffectReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.ContestEffectReadDefaultTextPlainString
}
