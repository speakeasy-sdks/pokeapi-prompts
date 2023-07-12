// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type VersionGroupListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type VersionGroupListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	VersionGroupListDefaultTextPlainString *string
}
