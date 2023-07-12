// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MoveAilmentReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

type MoveAilmentReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	MoveAilmentReadDefaultTextPlainString *string
}
