// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MachineReadRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *MachineReadRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type MachineReadResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Default response
	MachineReadDefaultTextPlainString *string
}

func (o *MachineReadResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MachineReadResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MachineReadResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MachineReadResponse) GetMachineReadDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.MachineReadDefaultTextPlainString
}
