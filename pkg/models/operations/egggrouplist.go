// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EggGroupListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *EggGroupListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *EggGroupListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type EggGroupListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	EggGroupListDefaultTextPlainString *string
}

func (o *EggGroupListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EggGroupListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EggGroupListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EggGroupListResponse) GetEggGroupListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.EggGroupListDefaultTextPlainString
}
