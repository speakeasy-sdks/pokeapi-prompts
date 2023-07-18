// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type AbilityListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *AbilityListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *AbilityListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type AbilityListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	AbilityListDefaultTextPlainString *string
}

func (o *AbilityListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AbilityListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AbilityListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *AbilityListResponse) GetAbilityListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.AbilityListDefaultTextPlainString
}
