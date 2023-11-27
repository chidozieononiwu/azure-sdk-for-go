//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CheckVirtualNetworkSubnetUsageClient contains the methods for the CheckVirtualNetworkSubnetUsage group.
// Don't use this type directly, use NewCheckVirtualNetworkSubnetUsageClient() instead.
type CheckVirtualNetworkSubnetUsageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCheckVirtualNetworkSubnetUsageClient creates a new instance of CheckVirtualNetworkSubnetUsageClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCheckVirtualNetworkSubnetUsageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CheckVirtualNetworkSubnetUsageClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CheckVirtualNetworkSubnetUsageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Execute - Get virtual network subnet usage for a given vNet resource id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-05-01
//   - locationName - The name of the location.
//   - parameters - The required parameters for creating or updating a server.
//   - options - CheckVirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the CheckVirtualNetworkSubnetUsageClient.Execute
//     method.
func (client *CheckVirtualNetworkSubnetUsageClient) Execute(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *CheckVirtualNetworkSubnetUsageClientExecuteOptions) (CheckVirtualNetworkSubnetUsageClientExecuteResponse, error) {
	var err error
	const operationName = "CheckVirtualNetworkSubnetUsageClient.Execute"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.executeCreateRequest(ctx, locationName, parameters, options)
	if err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	resp, err := client.executeHandleResponse(httpResp)
	return resp, err
}

// executeCreateRequest creates the Execute request.
func (client *CheckVirtualNetworkSubnetUsageClient) executeCreateRequest(ctx context.Context, locationName string, parameters VirtualNetworkSubnetUsageParameter, options *CheckVirtualNetworkSubnetUsageClientExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/checkVirtualNetworkSubnetUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// executeHandleResponse handles the Execute response.
func (client *CheckVirtualNetworkSubnetUsageClient) executeHandleResponse(resp *http.Response) (CheckVirtualNetworkSubnetUsageClientExecuteResponse, error) {
	result := CheckVirtualNetworkSubnetUsageClientExecuteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkSubnetUsageResult); err != nil {
		return CheckVirtualNetworkSubnetUsageClientExecuteResponse{}, err
	}
	return result, nil
}
