// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type BerryListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

type BerryListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	BerryListDefaultTextPlainString *string
}
