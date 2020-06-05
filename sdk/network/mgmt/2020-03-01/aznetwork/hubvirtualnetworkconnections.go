// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package aznetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// HubVirtualNetworkConnectionsOperations contains the methods for the HubVirtualNetworkConnections group.
type HubVirtualNetworkConnectionsOperations interface {
	// Get - Retrieves the details of a HubVirtualNetworkConnection.
	Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (*HubVirtualNetworkConnectionResponse, error)
	// List - Retrieves the details of all HubVirtualNetworkConnections.
	List(resourceGroupName string, virtualHubName string) (ListHubVirtualNetworkConnectionsResultPager, error)
}

// hubVirtualNetworkConnectionsOperations implements the HubVirtualNetworkConnectionsOperations interface.
type hubVirtualNetworkConnectionsOperations struct {
	*Client
	subscriptionID string
}

// Get - Retrieves the details of a HubVirtualNetworkConnection.
func (client *hubVirtualNetworkConnectionsOperations) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string) (*HubVirtualNetworkConnectionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualHubName, connectionName)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client *hubVirtualNetworkConnectionsOperations) getCreateRequest(resourceGroupName string, virtualHubName string, connectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *hubVirtualNetworkConnectionsOperations) getHandleResponse(resp *azcore.Response) (*HubVirtualNetworkConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := HubVirtualNetworkConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.HubVirtualNetworkConnection)
}

// getHandleError handles the Get error response.
func (client *hubVirtualNetworkConnectionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Retrieves the details of all HubVirtualNetworkConnections.
func (client *hubVirtualNetworkConnectionsOperations) List(resourceGroupName string, virtualHubName string) (ListHubVirtualNetworkConnectionsResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, virtualHubName)
	if err != nil {
		return nil, err
	}
	return &listHubVirtualNetworkConnectionsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ListHubVirtualNetworkConnectionsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListHubVirtualNetworkConnectionsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListHubVirtualNetworkConnectionsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *hubVirtualNetworkConnectionsOperations) listCreateRequest(resourceGroupName string, virtualHubName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listHandleResponse handles the List response.
func (client *hubVirtualNetworkConnectionsOperations) listHandleResponse(resp *azcore.Response) (*ListHubVirtualNetworkConnectionsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ListHubVirtualNetworkConnectionsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListHubVirtualNetworkConnectionsResult)
}

// listHandleError handles the List error response.
func (client *hubVirtualNetworkConnectionsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
