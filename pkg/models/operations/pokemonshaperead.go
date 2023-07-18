// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PokemonShapeReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *PokemonShapeReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type PokemonShapeReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	PokemonShapeReadDefaultTextPlainString *string
}

func (o *PokemonShapeReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PokemonShapeReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PokemonShapeReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PokemonShapeReadResponse) GetPokemonShapeReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.PokemonShapeReadDefaultTextPlainString
}
