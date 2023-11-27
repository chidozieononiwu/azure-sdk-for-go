//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CommunicationsServer is a fake server for instances of the armsupport.CommunicationsClient type.
type CommunicationsServer struct {
	// CheckNameAvailability is the fake for method CommunicationsClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, supportTicketName string, checkNameAvailabilityInput armsupport.CheckNameAvailabilityInput, options *armsupport.CommunicationsClientCheckNameAvailabilityOptions) (resp azfake.Responder[armsupport.CommunicationsClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method CommunicationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, supportTicketName string, communicationName string, createCommunicationParameters armsupport.CommunicationDetails, options *armsupport.CommunicationsClientBeginCreateOptions) (resp azfake.PollerResponder[armsupport.CommunicationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CommunicationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, supportTicketName string, communicationName string, options *armsupport.CommunicationsClientGetOptions) (resp azfake.Responder[armsupport.CommunicationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CommunicationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(supportTicketName string, options *armsupport.CommunicationsClientListOptions) (resp azfake.PagerResponder[armsupport.CommunicationsClientListResponse])
}

// NewCommunicationsServerTransport creates a new instance of CommunicationsServerTransport with the provided implementation.
// The returned CommunicationsServerTransport instance is connected to an instance of armsupport.CommunicationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCommunicationsServerTransport(srv *CommunicationsServer) *CommunicationsServerTransport {
	return &CommunicationsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armsupport.CommunicationsClientCreateResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armsupport.CommunicationsClientListResponse]](),
	}
}

// CommunicationsServerTransport connects instances of armsupport.CommunicationsClient to instances of CommunicationsServer.
// Don't use this type directly, use NewCommunicationsServerTransport instead.
type CommunicationsServerTransport struct {
	srv          *CommunicationsServer
	beginCreate  *tracker[azfake.PollerResponder[armsupport.CommunicationsClientCreateResponse]]
	newListPager *tracker[azfake.PagerResponder[armsupport.CommunicationsClientListResponse]]
}

// Do implements the policy.Transporter interface for CommunicationsServerTransport.
func (c *CommunicationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CommunicationsClient.CheckNameAvailability":
		resp, err = c.dispatchCheckNameAvailability(req)
	case "CommunicationsClient.BeginCreate":
		resp, err = c.dispatchBeginCreate(req)
	case "CommunicationsClient.Get":
		resp, err = c.dispatchGet(req)
	case "CommunicationsClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CommunicationsServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if c.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/supportTickets/(?P<supportTicketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsupport.CheckNameAvailabilityInput](req)
	if err != nil {
		return nil, err
	}
	supportTicketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("supportTicketName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CheckNameAvailability(req.Context(), supportTicketNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CommunicationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/supportTickets/(?P<supportTicketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communications/(?P<communicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsupport.CommunicationDetails](req)
		if err != nil {
			return nil, err
		}
		supportTicketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("supportTicketName")])
		if err != nil {
			return nil, err
		}
		communicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), supportTicketNameParam, communicationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

func (c *CommunicationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/supportTickets/(?P<supportTicketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communications/(?P<communicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	supportTicketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("supportTicketName")])
	if err != nil {
		return nil, err
	}
	communicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("communicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), supportTicketNameParam, communicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CommunicationDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CommunicationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Support/supportTickets/(?P<supportTicketName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/communications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		supportTicketNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("supportTicketName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armsupport.CommunicationsClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armsupport.CommunicationsClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := c.srv.NewListPager(supportTicketNameParam, options)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsupport.CommunicationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}
