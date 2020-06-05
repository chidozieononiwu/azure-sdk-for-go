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

// ApplicationSecurityGroupsOperations contains the methods for the ApplicationSecurityGroups group.
type ApplicationSecurityGroupsOperations interface {
	// BeginCreateOrUpdate - Creates or updates an application security group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup) (*ApplicationSecurityGroupResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (ApplicationSecurityGroupPoller, error)
	// BeginDelete - Deletes the specified application security group.
	BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified application security group.
	Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (*ApplicationSecurityGroupResponse, error)
	// List - Gets all the application security groups in a resource group.
	List(resourceGroupName string) (ApplicationSecurityGroupListResultPager, error)
	// ListAll - Gets all application security groups in a subscription.
	ListAll() (ApplicationSecurityGroupListResultPager, error)
	// UpdateTags - Updates an application security group's tags.
	UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject) (*ApplicationSecurityGroupResponse, error)
}

// applicationSecurityGroupsOperations implements the ApplicationSecurityGroupsOperations interface.
type applicationSecurityGroupsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Creates or updates an application security group.
func (client *applicationSecurityGroupsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup) (*ApplicationSecurityGroupResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, applicationSecurityGroupName, parameters)
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
	pt, err := createPollingTracker("applicationSecurityGroupsOperations.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &applicationSecurityGroupPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*ApplicationSecurityGroupResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *applicationSecurityGroupsOperations) ResumeCreateOrUpdate(token string) (ApplicationSecurityGroupPoller, error) {
	pt, err := resumePollingTracker("applicationSecurityGroupsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &applicationSecurityGroupPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *applicationSecurityGroupsOperations) createOrUpdateCreateRequest(resourceGroupName string, applicationSecurityGroupName string, parameters ApplicationSecurityGroup) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *applicationSecurityGroupsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*ApplicationSecurityGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *applicationSecurityGroupsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified application security group.
func (client *applicationSecurityGroupsOperations) BeginDelete(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, applicationSecurityGroupName)
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
	pt, err := createPollingTracker("applicationSecurityGroupsOperations.Delete", "location", resp, client.deleteHandleError)
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

func (client *applicationSecurityGroupsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("applicationSecurityGroupsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *applicationSecurityGroupsOperations) deleteCreateRequest(resourceGroupName string, applicationSecurityGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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
func (client *applicationSecurityGroupsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *applicationSecurityGroupsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets information about the specified application security group.
func (client *applicationSecurityGroupsOperations) Get(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string) (*ApplicationSecurityGroupResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, applicationSecurityGroupName)
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
func (client *applicationSecurityGroupsOperations) getCreateRequest(resourceGroupName string, applicationSecurityGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
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

// getHandleResponse handles the Get response.
func (client *applicationSecurityGroupsOperations) getHandleResponse(resp *azcore.Response) (*ApplicationSecurityGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
}

// getHandleError handles the Get error response.
func (client *applicationSecurityGroupsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Gets all the application security groups in a resource group.
func (client *applicationSecurityGroupsOperations) List(resourceGroupName string) (ApplicationSecurityGroupListResultPager, error) {
	req, err := client.listCreateRequest(resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &applicationSecurityGroupListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listHandleResponse,
		advancer: func(resp *ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ApplicationSecurityGroupListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ApplicationSecurityGroupListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listCreateRequest creates the List request.
func (client *applicationSecurityGroupsOperations) listCreateRequest(resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *applicationSecurityGroupsOperations) listHandleResponse(resp *azcore.Response) (*ApplicationSecurityGroupListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listHandleError(resp)
	}
	result := ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ApplicationSecurityGroupListResult)
}

// listHandleError handles the List error response.
func (client *applicationSecurityGroupsOperations) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all application security groups in a subscription.
func (client *applicationSecurityGroupsOperations) ListAll() (ApplicationSecurityGroupListResultPager, error) {
	req, err := client.listAllCreateRequest()
	if err != nil {
		return nil, err
	}
	return &applicationSecurityGroupListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listAllHandleResponse,
		advancer: func(resp *ApplicationSecurityGroupListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.ApplicationSecurityGroupListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.ApplicationSecurityGroupListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listAllCreateRequest creates the ListAll request.
func (client *applicationSecurityGroupsOperations) listAllCreateRequest() (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/applicationSecurityGroups"
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

// listAllHandleResponse handles the ListAll response.
func (client *applicationSecurityGroupsOperations) listAllHandleResponse(resp *azcore.Response) (*ApplicationSecurityGroupListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listAllHandleError(resp)
	}
	result := ApplicationSecurityGroupListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ApplicationSecurityGroupListResult)
}

// listAllHandleError handles the ListAll error response.
func (client *applicationSecurityGroupsOperations) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Updates an application security group's tags.
func (client *applicationSecurityGroupsOperations) UpdateTags(ctx context.Context, resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject) (*ApplicationSecurityGroupResponse, error) {
	req, err := client.updateTagsCreateRequest(resourceGroupName, applicationSecurityGroupName, parameters)
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
func (client *applicationSecurityGroupsOperations) updateTagsCreateRequest(resourceGroupName string, applicationSecurityGroupName string, parameters TagsObject) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationSecurityGroups/{applicationSecurityGroupName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{applicationSecurityGroupName}", url.PathEscape(applicationSecurityGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *applicationSecurityGroupsOperations) updateTagsHandleResponse(resp *azcore.Response) (*ApplicationSecurityGroupResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateTagsHandleError(resp)
	}
	result := ApplicationSecurityGroupResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.ApplicationSecurityGroup)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *applicationSecurityGroupsOperations) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
