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
	"time"
)

// SubnetsOperations contains the methods for the Subnets group.
type SubnetsOperations interface {
	// BeginCreateOrUpdate - Creates or updates a subnet in the specified virtual network.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*SubnetResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (SubnetPoller, error)
	// BeginDelete - Deletes the specified subnet.
	BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets the specified subnet by virtual network and resource group.
	Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*SubnetResponse, error)
	// List - Gets all subnets in a virtual network.
	List(resourceGroupName string, virtualNetworkName string) (SubnetListResultPager, error)
	// BeginPrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
	BeginPrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*HTTPResponse, error)
	// ResumePrepareNetworkPolicies - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumePrepareNetworkPolicies(token string) (HTTPPoller, error)
	// BeginUnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
	BeginUnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*HTTPResponse, error)
	// ResumeUnprepareNetworkPolicies - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUnprepareNetworkPolicies(token string) (HTTPPoller, error)
}

// subnetsOperations implements the SubnetsOperations interface.
type subnetsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates a subnet in the specified virtual network.
func (client *subnetsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*SubnetResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, virtualNetworkName, subnetName, subnetParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.createOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("subnetsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &subnetPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*SubnetResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *subnetsOperations) ResumeCreateOrUpdate(token string) (SubnetPoller, error) {
	pt, err := resumePollingTracker("subnetsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &subnetPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *subnetsOperations) createOrUpdateCreateRequest(resourceGroupName string, virtualNetworkName string, subnetName string, subnetParameters Subnet) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(subnetParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *subnetsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*SubnetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := SubnetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Subnet)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *subnetsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified subnet.
func (client *subnetsOperations) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, virtualNetworkName, subnetName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.deleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("subnetsOperations.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *subnetsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("subnetsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *subnetsOperations) deleteCreateRequest(resourceGroupName string, virtualNetworkName string, subnetName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *subnetsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *subnetsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets the specified subnet by virtual network and resource group.
func (client *subnetsOperations) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*SubnetResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, virtualNetworkName, subnetName, subnetsGetOptions)
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
func (client *subnetsOperations) getCreateRequest(resourceGroupName string, virtualNetworkName string, subnetName string, subnetsGetOptions *SubnetsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	if subnetsGetOptions != nil && subnetsGetOptions.Expand != nil {
		query.Set("$expand", *subnetsGetOptions.Expand)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *subnetsOperations) getHandleResponse(resp *azcore.Response) (*SubnetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := SubnetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Subnet)
}

// getHandleError handles the Get error response.
func (client *subnetsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all subnets in a virtual network.
func (client *subnetsOperations) List(resourceGroupName string, virtualNetworkName string) (SubnetListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName, virtualNetworkName)
	if err != nil {
		return nil, err
	}
	return &subnetListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *SubnetListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.SubnetListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.SubnetListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *subnetsOperations) listCreateRequest(resourceGroupName string, virtualNetworkName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
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
func (client *subnetsOperations) listHandleResponse(resp *azcore.Response) (*SubnetListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := SubnetListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.SubnetListResult)
}

// listHandleError handles the List error response.
func (client *subnetsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PrepareNetworkPolicies - Prepares a subnet by applying network intent policies.
func (client *subnetsOperations) BeginPrepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*HTTPResponse, error) {
	req, err := client.prepareNetworkPoliciesCreateRequest(resourceGroupName, virtualNetworkName, subnetName, prepareNetworkPoliciesRequestParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.prepareNetworkPoliciesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("subnetsOperations.PrepareNetworkPolicies", "location", resp, client.prepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *subnetsOperations) ResumePrepareNetworkPolicies(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("subnetsOperations.PrepareNetworkPolicies", token, client.prepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// prepareNetworkPoliciesCreateRequest creates the PrepareNetworkPolicies request.
func (client *subnetsOperations) prepareNetworkPoliciesCreateRequest(resourceGroupName string, virtualNetworkName string, subnetName string, prepareNetworkPoliciesRequestParameters PrepareNetworkPoliciesRequest) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/PrepareNetworkPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(prepareNetworkPoliciesRequestParameters)
}

// prepareNetworkPoliciesHandleResponse handles the PrepareNetworkPolicies response.
func (client *subnetsOperations) prepareNetworkPoliciesHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.prepareNetworkPoliciesHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// prepareNetworkPoliciesHandleError handles the PrepareNetworkPolicies error response.
func (client *subnetsOperations) prepareNetworkPoliciesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UnprepareNetworkPolicies - Unprepares a subnet by removing network intent policies.
func (client *subnetsOperations) BeginUnprepareNetworkPolicies(ctx context.Context, resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*HTTPResponse, error) {
	req, err := client.unprepareNetworkPoliciesCreateRequest(resourceGroupName, virtualNetworkName, subnetName, unprepareNetworkPoliciesRequestParameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.unprepareNetworkPoliciesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("subnetsOperations.UnprepareNetworkPolicies", "location", resp, client.unprepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *subnetsOperations) ResumeUnprepareNetworkPolicies(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("subnetsOperations.UnprepareNetworkPolicies", token, client.unprepareNetworkPoliciesHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// unprepareNetworkPoliciesCreateRequest creates the UnprepareNetworkPolicies request.
func (client *subnetsOperations) unprepareNetworkPoliciesCreateRequest(resourceGroupName string, virtualNetworkName string, subnetName string, unprepareNetworkPoliciesRequestParameters UnprepareNetworkPoliciesRequest) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualNetworks/{virtualNetworkName}/subnets/{subnetName}/UnprepareNetworkPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	urlPath = strings.ReplaceAll(urlPath, "{subnetName}", url.PathEscape(subnetName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPost, *u)
	return req, req.MarshalAsJSON(unprepareNetworkPoliciesRequestParameters)
}

// unprepareNetworkPoliciesHandleResponse handles the UnprepareNetworkPolicies response.
func (client *subnetsOperations) unprepareNetworkPoliciesHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.unprepareNetworkPoliciesHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// unprepareNetworkPoliciesHandleError handles the UnprepareNetworkPolicies error response.
func (client *subnetsOperations) unprepareNetworkPoliciesHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
