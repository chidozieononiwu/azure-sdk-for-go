//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// EndpointCertificatesClient contains the methods for the EndpointCertificates group.
// Don't use this type directly, use NewEndpointCertificatesClient() instead.
type EndpointCertificatesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEndpointCertificatesClient creates a new instance of EndpointCertificatesClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEndpointCertificatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EndpointCertificatesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &EndpointCertificatesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a certificate used on the endpoint with the given id.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// endpointType - Type of the endpoint whose certificate the customer is looking for.
// options - EndpointCertificatesClientGetOptions contains the optional parameters for the EndpointCertificatesClient.Get
// method.
func (client *EndpointCertificatesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, endpointType string, options *EndpointCertificatesClientGetOptions) (EndpointCertificatesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, endpointType, options)
	if err != nil {
		return EndpointCertificatesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return EndpointCertificatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EndpointCertificatesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *EndpointCertificatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, endpointType string, options *EndpointCertificatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/endpointCertificates/{endpointType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if endpointType == "" {
		return nil, errors.New("parameter endpointType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointType}", url.PathEscape(endpointType))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EndpointCertificatesClient) getHandleResponse(resp *http.Response) (EndpointCertificatesClientGetResponse, error) {
	result := EndpointCertificatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointCertificate); err != nil {
		return EndpointCertificatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByInstancePager - List certificates used on endpoints on the target instance.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// managedInstanceName - The name of the managed instance.
// options - EndpointCertificatesClientListByInstanceOptions contains the optional parameters for the EndpointCertificatesClient.ListByInstance
// method.
func (client *EndpointCertificatesClient) NewListByInstancePager(resourceGroupName string, managedInstanceName string, options *EndpointCertificatesClientListByInstanceOptions) *runtime.Pager[EndpointCertificatesClientListByInstanceResponse] {
	return runtime.NewPager(runtime.PageProcessor[EndpointCertificatesClientListByInstanceResponse]{
		More: func(page EndpointCertificatesClientListByInstanceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EndpointCertificatesClientListByInstanceResponse) (EndpointCertificatesClientListByInstanceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EndpointCertificatesClientListByInstanceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EndpointCertificatesClientListByInstanceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EndpointCertificatesClientListByInstanceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByInstanceHandleResponse(resp)
		},
	})
}

// listByInstanceCreateRequest creates the ListByInstance request.
func (client *EndpointCertificatesClient) listByInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *EndpointCertificatesClientListByInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/endpointCertificates"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByInstanceHandleResponse handles the ListByInstance response.
func (client *EndpointCertificatesClient) listByInstanceHandleResponse(resp *http.Response) (EndpointCertificatesClientListByInstanceResponse, error) {
	result := EndpointCertificatesClientListByInstanceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointCertificateListResult); err != nil {
		return EndpointCertificatesClientListByInstanceResponse{}, err
	}
	return result, nil
}