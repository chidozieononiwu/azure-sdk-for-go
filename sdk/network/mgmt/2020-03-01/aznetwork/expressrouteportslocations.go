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

// ExpressRoutePortsLocationsOperations contains the methods for the ExpressRoutePortsLocations group.
type ExpressRoutePortsLocationsOperations interface {
	// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said peering location.
	Get(ctx context.Context, locationName string) (*ExpressRoutePortsLocationResponse, error)
	// List - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location. Available bandwidths can only be obtained when retrieving a specific peering location.
	List() (ExpressRoutePortsLocationListResultPager, error)
}

// expressRoutePortsLocationsOperations implements the ExpressRoutePortsLocationsOperations interface.
type expressRoutePortsLocationsOperations struct {
	*Client
	subscriptionID string
}

// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said peering location.
func (client *expressRoutePortsLocationsOperations) Get(ctx context.Context, locationName string) (*ExpressRoutePortsLocationResponse, error) {
	req, err := client.getCreateRequest(locationName)
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
func (client *expressRoutePortsLocationsOperations) getCreateRequest(locationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
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
func (client *expressRoutePortsLocationsOperations) getHandleResponse(resp *azcore.Response) (*ExpressRoutePortsLocationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ExpressRoutePortsLocationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePortsLocation)
}

// getHandleError handles the Get error response.
func (client *expressRoutePortsLocationsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location. Available bandwidths can only be obtained when retrieving a specific peering location.
func (client *expressRoutePortsLocationsOperations) List() (ExpressRoutePortsLocationListResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &expressRoutePortsLocationListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ExpressRoutePortsLocationListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ExpressRoutePortsLocationListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ExpressRoutePortsLocationListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *expressRoutePortsLocationsOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *expressRoutePortsLocationsOperations) listHandleResponse(resp *azcore.Response) (*ExpressRoutePortsLocationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ExpressRoutePortsLocationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ExpressRoutePortsLocationListResult)
}

// listHandleError handles the List error response.
func (client *expressRoutePortsLocationsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
