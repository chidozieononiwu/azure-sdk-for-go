// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azcompute

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// GalleryImageVersionsOperations contains the methods for the GalleryImageVersions group.
type GalleryImageVersionsOperations interface {
	// BeginCreateOrUpdate - Create or update a gallery Image Version.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion) (*GalleryImageVersionResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (GalleryImageVersionPoller, error)
	// BeginDelete - Delete a gallery Image Version.
	BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*HTTPResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieves information about a gallery Image Version.
	Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*GalleryImageVersionResponse, error)
	// ListByGalleryImage - List gallery Image Versions in a gallery Image Definition.
	ListByGalleryImage(resourceGroupName string, galleryName string, galleryImageName string) (GalleryImageVersionListPager, error)
	// BeginUpdate - Update a gallery Image Version.
	BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate) (*GalleryImageVersionResponse, error)
	// ResumeUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeUpdate(token string) (GalleryImageVersionPoller, error)
}

// galleryImageVersionsOperations implements the GalleryImageVersionsOperations interface.
type galleryImageVersionsOperations struct {
	*Client
	subscriptionID string
}

// CreateOrUpdate - Create or update a gallery Image Version.
func (client *galleryImageVersionsOperations) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion) (*GalleryImageVersionResponse, error) {
	req, err := client.createOrUpdateCreateRequest(resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion)
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
	pt, err := createPollingTracker("galleryImageVersionsOperations.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryImageVersionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryImageVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *galleryImageVersionsOperations) ResumeCreateOrUpdate(token string) (GalleryImageVersionPoller, error) {
	pt, err := resumePollingTracker("galleryImageVersionsOperations.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryImageVersionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *galleryImageVersionsOperations) createOrUpdateCreateRequest(resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersion) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(galleryImageVersion)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *galleryImageVersionsOperations) createOrUpdateHandleResponse(resp *azcore.Response) (*GalleryImageVersionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	result := GalleryImageVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImageVersion)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *galleryImageVersionsOperations) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Delete a gallery Image Version.
func (client *galleryImageVersionsOperations) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*HTTPResponse, error) {
	req, err := client.deleteCreateRequest(resourceGroupName, galleryName, galleryImageName, galleryImageVersionName)
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
	pt, err := createPollingTracker("galleryImageVersionsOperations.Delete", "", resp, client.deleteHandleError)
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

func (client *galleryImageVersionsOperations) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := resumePollingTracker("galleryImageVersionsOperations.Delete", token, client.deleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *galleryImageVersionsOperations) deleteCreateRequest(resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *galleryImageVersionsOperations) deleteHandleResponse(resp *azcore.Response) (*HTTPResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return &HTTPResponse{RawResponse: resp.Response}, nil
}

// deleteHandleError handles the Delete error response.
func (client *galleryImageVersionsOperations) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieves information about a gallery Image Version.
func (client *galleryImageVersionsOperations) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*GalleryImageVersionResponse, error) {
	req, err := client.getCreateRequest(resourceGroupName, galleryName, galleryImageName, galleryImageVersionName)
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
func (client *galleryImageVersionsOperations) getCreateRequest(resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *galleryImageVersionsOperations) getHandleResponse(resp *azcore.Response) (*GalleryImageVersionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getHandleError(resp)
	}
	result := GalleryImageVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImageVersion)
}

// getHandleError handles the Get error response.
func (client *galleryImageVersionsOperations) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByGalleryImage - List gallery Image Versions in a gallery Image Definition.
func (client *galleryImageVersionsOperations) ListByGalleryImage(resourceGroupName string, galleryName string, galleryImageName string) (GalleryImageVersionListPager, error) {
	req, err := client.listByGalleryImageCreateRequest(resourceGroupName, galleryName, galleryImageName)
	if err != nil {
		return nil, err
	}
	return &galleryImageVersionListPager{
		pipeline:  client.p,
		request:   req,
		responder: client.listByGalleryImageHandleResponse,
		advancer: func(resp *GalleryImageVersionListResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.GalleryImageVersionList.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.GalleryImageVersionList.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// listByGalleryImageCreateRequest creates the ListByGalleryImage request.
func (client *galleryImageVersionsOperations) listByGalleryImageCreateRequest(resourceGroupName string, galleryName string, galleryImageName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// listByGalleryImageHandleResponse handles the ListByGalleryImage response.
func (client *galleryImageVersionsOperations) listByGalleryImageHandleResponse(resp *azcore.Response) (*GalleryImageVersionListResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.listByGalleryImageHandleError(resp)
	}
	result := GalleryImageVersionListResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImageVersionList)
}

// listByGalleryImageHandleError handles the ListByGalleryImage error response.
func (client *galleryImageVersionsOperations) listByGalleryImageHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Update - Update a gallery Image Version.
func (client *galleryImageVersionsOperations) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate) (*GalleryImageVersionResponse, error) {
	req, err := client.updateCreateRequest(resourceGroupName, galleryName, galleryImageName, galleryImageVersionName, galleryImageVersion)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := createPollingTracker("galleryImageVersionsOperations.Update", "", resp, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &galleryImageVersionPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*GalleryImageVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *galleryImageVersionsOperations) ResumeUpdate(token string) (GalleryImageVersionPoller, error) {
	pt, err := resumePollingTracker("galleryImageVersionsOperations.Update", token, client.updateHandleError)
	if err != nil {
		return nil, err
	}
	return &galleryImageVersionPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// updateCreateRequest creates the Update request.
func (client *galleryImageVersionsOperations) updateCreateRequest(resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion GalleryImageVersionUpdate) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/images/{galleryImageName}/versions/{galleryImageVersionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageVersionName}", url.PathEscape(galleryImageVersionName))
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2019-12-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(galleryImageVersion)
}

// updateHandleResponse handles the Update response.
func (client *galleryImageVersionsOperations) updateHandleResponse(resp *azcore.Response) (*GalleryImageVersionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.updateHandleError(resp)
	}
	result := GalleryImageVersionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.GalleryImageVersion)
}

// updateHandleError handles the Update error response.
func (client *galleryImageVersionsOperations) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
