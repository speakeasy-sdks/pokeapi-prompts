// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EncounterConditionValueListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *EncounterConditionValueListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *EncounterConditionValueListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type EncounterConditionValueListResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	EncounterConditionValueListDefaultTextPlainString *string
}

func (o *EncounterConditionValueListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EncounterConditionValueListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EncounterConditionValueListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EncounterConditionValueListResponse) GetEncounterConditionValueListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.EncounterConditionValueListDefaultTextPlainString
}
