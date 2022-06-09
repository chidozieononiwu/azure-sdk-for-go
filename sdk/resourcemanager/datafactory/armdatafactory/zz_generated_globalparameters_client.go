//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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

// GlobalParametersClient contains the methods for the GlobalParameters group.
// Don't use this type directly, use NewGlobalParametersClient() instead.
type GlobalParametersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewGlobalParametersClient creates a new instance of GlobalParametersClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGlobalParametersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GlobalParametersClient, error) {
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
	client := &GlobalParametersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Global parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// globalParameterName - The global parameter name.
// defaultParam - Global parameter resource definition.
// options - GlobalParametersClientCreateOrUpdateOptions contains the optional parameters for the GlobalParametersClient.CreateOrUpdate
// method.
func (client *GlobalParametersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, defaultParam GlobalParameterResource, options *GlobalParametersClientCreateOrUpdateOptions) (GlobalParametersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, globalParameterName, defaultParam, options)
	if err != nil {
		return GlobalParametersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalParametersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalParametersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GlobalParametersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, defaultParam GlobalParameterResource, options *GlobalParametersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/globalParameters/{globalParameterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if globalParameterName == "" {
		return nil, errors.New("parameter globalParameterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalParameterName}", url.PathEscape(globalParameterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, defaultParam)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GlobalParametersClient) createOrUpdateHandleResponse(resp *http.Response) (GlobalParametersClientCreateOrUpdateResponse, error) {
	result := GlobalParametersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalParameterResource); err != nil {
		return GlobalParametersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a Global parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// globalParameterName - The global parameter name.
// options - GlobalParametersClientDeleteOptions contains the optional parameters for the GlobalParametersClient.Delete method.
func (client *GlobalParametersClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, options *GlobalParametersClientDeleteOptions) (GlobalParametersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, globalParameterName, options)
	if err != nil {
		return GlobalParametersClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalParametersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GlobalParametersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return GlobalParametersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GlobalParametersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, options *GlobalParametersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/globalParameters/{globalParameterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if globalParameterName == "" {
		return nil, errors.New("parameter globalParameterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalParameterName}", url.PathEscape(globalParameterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Global parameter
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// globalParameterName - The global parameter name.
// options - GlobalParametersClientGetOptions contains the optional parameters for the GlobalParametersClient.Get method.
func (client *GlobalParametersClient) Get(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, options *GlobalParametersClientGetOptions) (GlobalParametersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, globalParameterName, options)
	if err != nil {
		return GlobalParametersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GlobalParametersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GlobalParametersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GlobalParametersClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, globalParameterName string, options *GlobalParametersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/globalParameters/{globalParameterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if globalParameterName == "" {
		return nil, errors.New("parameter globalParameterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{globalParameterName}", url.PathEscape(globalParameterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GlobalParametersClient) getHandleResponse(resp *http.Response) (GlobalParametersClientGetResponse, error) {
	result := GlobalParametersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalParameterResource); err != nil {
		return GlobalParametersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists Global parameters
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-06-01
// resourceGroupName - The resource group name.
// factoryName - The factory name.
// options - GlobalParametersClientListByFactoryOptions contains the optional parameters for the GlobalParametersClient.ListByFactory
// method.
func (client *GlobalParametersClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *GlobalParametersClientListByFactoryOptions) *runtime.Pager[GlobalParametersClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[GlobalParametersClientListByFactoryResponse]{
		More: func(page GlobalParametersClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GlobalParametersClientListByFactoryResponse) (GlobalParametersClientListByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return GlobalParametersClientListByFactoryResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return GlobalParametersClientListByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return GlobalParametersClientListByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFactoryHandleResponse(resp)
		},
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *GlobalParametersClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *GlobalParametersClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/globalParameters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *GlobalParametersClient) listByFactoryHandleResponse(resp *http.Response) (GlobalParametersClientListByFactoryResponse, error) {
	result := GlobalParametersClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GlobalParameterListResponse); err != nil {
		return GlobalParametersClientListByFactoryResponse{}, err
	}
	return result, nil
}