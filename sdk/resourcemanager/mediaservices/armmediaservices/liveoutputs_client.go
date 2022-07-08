//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LiveOutputsClient contains the methods for the LiveOutputs group.
// Don't use this type directly, use NewLiveOutputsClient() instead.
type LiveOutputsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewLiveOutputsClient creates a new instance of LiveOutputsClient with the specified values.
// subscriptionID - The unique identifier for a Microsoft Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewLiveOutputsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LiveOutputsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &LiveOutputsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreate - Creates a new live output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// liveOutputName - The name of the live output.
// parameters - Live Output properties needed for creation.
// options - LiveOutputsClientBeginCreateOptions contains the optional parameters for the LiveOutputsClient.BeginCreate method.
func (client *LiveOutputsClient) BeginCreate(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*runtime.Poller[LiveOutputsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LiveOutputsClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LiveOutputsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - Creates a new live output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
func (client *LiveOutputsClient) create(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *LiveOutputsClient) createCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, parameters LiveOutput, options *LiveOutputsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// liveOutputName - The name of the live output.
// options - LiveOutputsClientBeginDeleteOptions contains the optional parameters for the LiveOutputsClient.BeginDelete method.
func (client *LiveOutputsClient) BeginDelete(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*runtime.Poller[LiveOutputsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[LiveOutputsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[LiveOutputsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes a live output. Deleting a live output does not delete the asset the live output is writing to.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
func (client *LiveOutputsClient) deleteOperation(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LiveOutputsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a live output.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// liveOutputName - The name of the live output.
// options - LiveOutputsClientGetOptions contains the optional parameters for the LiveOutputsClient.Get method.
func (client *LiveOutputsClient) Get(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientGetOptions) (LiveOutputsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, liveEventName, liveOutputName, options)
	if err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LiveOutputsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LiveOutputsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, liveOutputName string, options *LiveOutputsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs/{liveOutputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	if liveOutputName == "" {
		return nil, errors.New("parameter liveOutputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveOutputName}", url.PathEscape(liveOutputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LiveOutputsClient) getHandleResponse(resp *http.Response) (LiveOutputsClientGetResponse, error) {
	result := LiveOutputsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveOutput); err != nil {
		return LiveOutputsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the live outputs of a live event.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2021-11-01
// resourceGroupName - The name of the resource group within the Azure subscription.
// accountName - The Media Services account name.
// liveEventName - The name of the live event, maximum length is 32.
// options - LiveOutputsClientListOptions contains the optional parameters for the LiveOutputsClient.List method.
func (client *LiveOutputsClient) NewListPager(resourceGroupName string, accountName string, liveEventName string, options *LiveOutputsClientListOptions) *runtime.Pager[LiveOutputsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LiveOutputsClientListResponse]{
		More: func(page LiveOutputsClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LiveOutputsClientListResponse) (LiveOutputsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, liveEventName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.ODataNextLink)
			}
			if err != nil {
				return LiveOutputsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return LiveOutputsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LiveOutputsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LiveOutputsClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, liveEventName string, options *LiveOutputsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaservices/{accountName}/liveEvents/{liveEventName}/liveOutputs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if liveEventName == "" {
		return nil, errors.New("parameter liveEventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{liveEventName}", url.PathEscape(liveEventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LiveOutputsClient) listHandleResponse(resp *http.Response) (LiveOutputsClientListResponse, error) {
	result := LiveOutputsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LiveOutputListResult); err != nil {
		return LiveOutputsClientListResponse{}, err
	}
	return result, nil
}