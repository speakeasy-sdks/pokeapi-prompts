// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MoveTargetListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MoveTargetListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MoveTargetListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type MoveTargetListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	MoveTargetListDefaultTextPlainString *string
}

func (o *MoveTargetListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MoveTargetListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MoveTargetListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MoveTargetListResponse) GetMoveTargetListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.MoveTargetListDefaultTextPlainString
}
