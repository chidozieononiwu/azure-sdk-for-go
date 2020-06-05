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

// VpnServerConfigurationsOperations contains the methods for the VpnServerConfigurations group.
type VpnServerConfigurationsOperations interface {
	// BeginCreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VpnServerConfiguration) (*VpnServerConfigurationResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (VpnServerConfigurationPoller, error)
	// BeginDelete - Deletes a VpnServerConfiguration.
	BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves the details of a VpnServerConfiguration.
	Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*VpnServerConfigurationResponse, error)
	// List - Lists all the VpnServerConfigurations in a subscription.
	List() (ListVpnServerConfigurationsResultPager, error)
	// ListByResourceGroup - Lists all the vpnServerConfigurations in a resource group.
	ListByResourceGroup(resourceGroupName string) (ListVpnServerConfigurationsResultPager, error)
	// UpdateTags - Updates VpnServerConfiguration tags.
	UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject) (*VpnServerConfigurationResponse, error)
}

// vpnServerConfigurationsOperations implements the VpnServerConfigurationsOperations interface.
type vpnServerConfigurationsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
func (client *vpnServerConfigurationsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VpnServerConfiguration) (*VpnServerConfigurationResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters)
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
	pt, err := createPollingTracker("vpnServerConfigurationsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &vpnServerConfigurationPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*VpnServerConfigurationResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *vpnServerConfigurationsOperations) ResumeCreateOrUpdate(token string) (VpnServerConfigurationPoller, error) {
	pt, err := resumePollingTracker("vpnServerConfigurationsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vpnServerConfigurationPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *vpnServerConfigurationsOperations) createOrUpdateCreateRequest(resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VpnServerConfiguration) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(vpnServerConfigurationParameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *vpnServerConfigurationsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*VpnServerConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := VpnServerConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnServerConfiguration)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *vpnServerConfigurationsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes a VpnServerConfiguration.
func (client *vpnServerConfigurationsOperations) BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, vpnServerConfigurationName)
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
	pt, err := createPollingTracker("vpnServerConfigurationsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *vpnServerConfigurationsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("vpnServerConfigurationsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *vpnServerConfigurationsOperations) deleteCreateRequest(resourceGroupName string, vpnServerConfigurationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
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
func (client *vpnServerConfigurationsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *vpnServerConfigurationsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieves the details of a VpnServerConfiguration.
func (client *vpnServerConfigurationsOperations) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string) (*VpnServerConfigurationResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, vpnServerConfigurationName)
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
func (client *vpnServerConfigurationsOperations) getCreateRequest(resourceGroupName string, vpnServerConfigurationName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
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
func (client *vpnServerConfigurationsOperations) getHandleResponse(resp *azcore.Response) (*VpnServerConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := VpnServerConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnServerConfiguration)
}

// getHandleError handles the Get error response.
func (client *vpnServerConfigurationsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all the VpnServerConfigurations in a subscription.
func (client *vpnServerConfigurationsOperations) List() (ListVpnServerConfigurationsResultPager, error) {
	req, err := client.listCreateRequest()
	if err != nil {
		return nil, err
	}
	return &listVpnServerConfigurationsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ListVpnServerConfigurationsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVpnServerConfigurationsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVpnServerConfigurationsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *vpnServerConfigurationsOperations) listCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations"
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
func (client *vpnServerConfigurationsOperations) listHandleResponse(resp *azcore.Response) (*ListVpnServerConfigurationsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ListVpnServerConfigurationsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnServerConfigurationsResult)
}

// listHandleError handles the List error response.
func (client *vpnServerConfigurationsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Lists all the vpnServerConfigurations in a resource group.
func (client *vpnServerConfigurationsOperations) ListByResourceGroup(resourceGroupName string) (ListVpnServerConfigurationsResultPager, error) {
	req, err := client.listByResourceGroupCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &listVpnServerConfigurationsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByResourceGroupHandleResponse,
		advancer: func(resp *ListVpnServerConfigurationsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ListVpnServerConfigurationsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ListVpnServerConfigurationsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *vpnServerConfigurationsOperations) listByResourceGroupCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *vpnServerConfigurationsOperations) listByResourceGroupHandleResponse(resp *azcore.Response) (*ListVpnServerConfigurationsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByResourceGroupHandleError(resp)
	}
	result := ListVpnServerConfigurationsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ListVpnServerConfigurationsResult)
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *vpnServerConfigurationsOperations) listByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates VpnServerConfiguration tags.
func (client *vpnServerConfigurationsOperations) UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject) (*VpnServerConfigurationResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *vpnServerConfigurationsOperations) updateTagsCreateRequest(resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(vpnServerConfigurationParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *vpnServerConfigurationsOperations) updateTagsHandleResponse(resp *azcore.Response) (*VpnServerConfigurationResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := VpnServerConfigurationResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnServerConfiguration)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *vpnServerConfigurationsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
