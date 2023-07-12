// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package pokeapi

import (
	"PokeAPI/pkg/models/operations"
	"PokeAPI/pkg/utils"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type encounterConditionValue struct {
	sdkConfiguration sdkConfiguration
}

func newEncounterConditionValue(sdkConfig sdkConfiguration) *encounterConditionValue {
	return &encounterConditionValue{
		sdkConfiguration: sdkConfig,
	}
}

func (s *encounterConditionValue) EncounterConditionValueList(ctx context.Context, request operations.EncounterConditionValueListRequest) (*operations.EncounterConditionValueListResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/api/v2/encounter-condition-value/"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.EncounterConditionValueListResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	default:
		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			out := string(rawBody)
			res.EncounterConditionValueListDefaultTextPlainString = &out
		}
	}

	return res, nil
}

func (s *encounterConditionValue) EncounterConditionValueRead(ctx context.Context, request operations.EncounterConditionValueReadRequest) (*operations.EncounterConditionValueReadResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url, err := utils.GenerateURL(ctx, baseURL, "/api/v2/encounter-condition-value/{id}/", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.EncounterConditionValueReadResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	default:
		switch {
		case utils.MatchContentType(contentType, `text/plain`):
			out := string(rawBody)
			res.EncounterConditionValueReadDefaultTextPlainString = &out
		}
	}

	return res, nil
}