// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PokemonFormListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *PokemonFormListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *PokemonFormListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type PokemonFormListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	PokemonFormListDefaultTextPlainString *string
}

func (o *PokemonFormListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokemonFormListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokemonFormListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PokemonFormListResponse) GetPokemonFormListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.PokemonFormListDefaultTextPlainString
}
