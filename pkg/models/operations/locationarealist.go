// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type LocationAreaListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *LocationAreaListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *LocationAreaListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type LocationAreaListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	LocationAreaListDefaultTextPlainString *string
}

func (o *LocationAreaListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *LocationAreaListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *LocationAreaListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *LocationAreaListResponse) GetLocationAreaListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.LocationAreaListDefaultTextPlainString
}
