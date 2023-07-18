// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type BerryFirmnessListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *BerryFirmnessListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *BerryFirmnessListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type BerryFirmnessListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	BerryFirmnessListDefaultTextPlainString *string
}

func (o *BerryFirmnessListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *BerryFirmnessListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *BerryFirmnessListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *BerryFirmnessListResponse) GetBerryFirmnessListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.BerryFirmnessListDefaultTextPlainString
}
