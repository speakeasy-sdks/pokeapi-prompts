// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MoveListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type MoveListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	MoveListDefaultTextPlainString *string
}
