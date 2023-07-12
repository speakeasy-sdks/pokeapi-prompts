// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PokedexReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type PokedexReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	PokedexReadDefaultTextPlainString *string
}