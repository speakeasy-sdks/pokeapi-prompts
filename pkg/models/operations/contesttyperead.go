// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ContestTypeReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *ContestTypeReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type ContestTypeReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	ContestTypeReadDefaultTextPlainString *string
}

func (o *ContestTypeReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ContestTypeReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ContestTypeReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ContestTypeReadResponse) GetContestTypeReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.ContestTypeReadDefaultTextPlainString
}
