// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type StatReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type StatReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	StatReadDefaultTextPlainString *string
}
