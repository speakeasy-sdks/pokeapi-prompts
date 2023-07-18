// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PokeathlonStatListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *PokeathlonStatListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *PokeathlonStatListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type PokeathlonStatListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	PokeathlonStatListDefaultTextPlainString *string
}

func (o *PokeathlonStatListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokeathlonStatListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokeathlonStatListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PokeathlonStatListResponse) GetPokeathlonStatListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.PokeathlonStatListDefaultTextPlainString
}
