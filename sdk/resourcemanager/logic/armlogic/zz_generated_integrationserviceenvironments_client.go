//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// IntegrationServiceEnvironmentsClient contains the methods for the IntegrationServiceEnvironments group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentsClient() instead.
type IntegrationServiceEnvironmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewIntegrationServiceEnvironmentsClient creates a new instance of IntegrationServiceEnvironmentsClient with the specified values.
func NewIntegrationServiceEnvironmentsClient(con *arm.Connection, subscriptionID string) *IntegrationServiceEnvironmentsClient {
	return &IntegrationServiceEnvironmentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginCreateOrUpdateOptions) (IntegrationServiceEnvironmentsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
	if err != nil {
		return IntegrationServiceEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result := IntegrationServiceEnvironmentsCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationServiceEnvironmentsClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return IntegrationServiceEnvironmentsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &IntegrationServiceEnvironmentsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) createOrUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *IntegrationServiceEnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, integrationServiceEnvironment)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *IntegrationServiceEnvironmentsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) Delete(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsDeleteOptions) (IntegrationServiceEnvironmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return IntegrationServiceEnvironmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return IntegrationServiceEnvironmentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *IntegrationServiceEnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *IntegrationServiceEnvironmentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) Get(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsGetOptions) (IntegrationServiceEnvironmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationServiceEnvironmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationServiceEnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationServiceEnvironmentsClient) getHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsGetResponse, error) {
	result := IntegrationServiceEnvironmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironment); err != nil {
		return IntegrationServiceEnvironmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *IntegrationServiceEnvironmentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets a list of integration service environments by resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) ListByResourceGroup(resourceGroup string, options *IntegrationServiceEnvironmentsListByResourceGroupOptions) *IntegrationServiceEnvironmentsListByResourceGroupPager {
	return &IntegrationServiceEnvironmentsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroup, options)
		},
		advancer: func(ctx context.Context, resp IntegrationServiceEnvironmentsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationServiceEnvironmentListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *IntegrationServiceEnvironmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroup string, options *IntegrationServiceEnvironmentsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *IntegrationServiceEnvironmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsListByResourceGroupResponse, error) {
	result := IntegrationServiceEnvironmentsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentListResult); err != nil {
		return IntegrationServiceEnvironmentsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *IntegrationServiceEnvironmentsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Gets a list of integration service environments by subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) ListBySubscription(options *IntegrationServiceEnvironmentsListBySubscriptionOptions) *IntegrationServiceEnvironmentsListBySubscriptionPager {
	return &IntegrationServiceEnvironmentsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp IntegrationServiceEnvironmentsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationServiceEnvironmentListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *IntegrationServiceEnvironmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *IntegrationServiceEnvironmentsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Logic/integrationServiceEnvironments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *IntegrationServiceEnvironmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentsListBySubscriptionResponse, error) {
	result := IntegrationServiceEnvironmentsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentListResult); err != nil {
		return IntegrationServiceEnvironmentsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *IntegrationServiceEnvironmentsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Restart - Restarts an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) Restart(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsRestartOptions) (IntegrationServiceEnvironmentsRestartResponse, error) {
	req, err := client.restartCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
	if err != nil {
		return IntegrationServiceEnvironmentsRestartResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return IntegrationServiceEnvironmentsRestartResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationServiceEnvironmentsRestartResponse{}, client.restartHandleError(resp)
	}
	return IntegrationServiceEnvironmentsRestartResponse{RawResponse: resp}, nil
}

// restartCreateRequest creates the Restart request.
func (client *IntegrationServiceEnvironmentsClient) restartCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentsRestartOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/restart"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// restartHandleError handles the Restart error response.
func (client *IntegrationServiceEnvironmentsClient) restartHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) BeginUpdate(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginUpdateOptions) (IntegrationServiceEnvironmentsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
	if err != nil {
		return IntegrationServiceEnvironmentsUpdatePollerResponse{}, err
	}
	result := IntegrationServiceEnvironmentsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("IntegrationServiceEnvironmentsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return IntegrationServiceEnvironmentsUpdatePollerResponse{}, err
	}
	result.Poller = &IntegrationServiceEnvironmentsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates an integration service environment.
// If the operation fails it returns the *ErrorResponse error type.
func (client *IntegrationServiceEnvironmentsClient) update(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, integrationServiceEnvironment, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *IntegrationServiceEnvironmentsClient) updateCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, integrationServiceEnvironment IntegrationServiceEnvironment, options *IntegrationServiceEnvironmentsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, integrationServiceEnvironment)
}

// updateHandleError handles the Update error response.
func (client *IntegrationServiceEnvironmentsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
