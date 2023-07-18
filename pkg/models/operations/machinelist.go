// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MachineListRequest struct {
	Limit  *int64 `queryParam:"style=form,explode=true,name=limit"`
	Offset *int64 `queryParam:"style=form,explode=true,name=offset"`
}

func (o *MachineListRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *MachineListRequest) GetOffset() *int64 {
	if o == nil {
		return nil
	}
	return o.Offset
}

type MachineListResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Default response
	MachineListDefaultTextPlainString *string
}

func (o *MachineListResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MachineListResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MachineListResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MachineListResponse) GetMachineListDefaultTextPlainString() *string {
	if o == nil {
		return nil
	}
	return o.MachineListDefaultTextPlainString
}
