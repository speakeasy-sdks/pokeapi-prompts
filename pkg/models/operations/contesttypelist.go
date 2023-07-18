// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type ContestTypeListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *ContestTypeListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ContestTypeListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type ContestTypeListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	ContestTypeListDefaultTextPlainString *string
}

func (o *ContestTypeListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ContestTypeListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ContestTypeListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ContestTypeListResponse) GetContestTypeListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.ContestTypeListDefaultTextPlainString
}
