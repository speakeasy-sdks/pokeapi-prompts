// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type EncounterConditionValueReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *EncounterConditionValueReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type EncounterConditionValueReadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	EncounterConditionValueReadDefaultTextPlainString *string
}

func (o *EncounterConditionValueReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *EncounterConditionValueReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *EncounterConditionValueReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *EncounterConditionValueReadResponse) GetEncounterConditionValueReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.EncounterConditionValueReadDefaultTextPlainString
}
