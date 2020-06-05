// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azeventhub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ArmDisasterRecoveryListResultPager provides iteration over ArmDisasterRecoveryListResult pages.
type ArmDisasterRecoveryListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ArmDisasterRecoveryListResultResponse.
	PageResponse() *ArmDisasterRecoveryListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type armDisasterRecoveryListResultHandleResponse func(*azcore.Response) (*ArmDisasterRecoveryListResultResponse, error)

type armDisasterRecoveryListResultAdvancePage func(*ArmDisasterRecoveryListResultResponse) (*azcore.Request, error)

type armDisasterRecoveryListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder armDisasterRecoveryListResultHandleResponse
	// callback for advancing to the next page
	advancer armDisasterRecoveryListResultAdvancePage
	// contains the current response
	current *ArmDisasterRecoveryListResultResponse
	// any error encountered
	err error
}

func (p *armDisasterRecoveryListResultPager) Err() error {
	return p.err
}

func (p *armDisasterRecoveryListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.ArmDisasterRecoveryListResult.NextLink == nil || len(*p.current.ArmDisasterRecoveryListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *armDisasterRecoveryListResultPager) PageResponse() *ArmDisasterRecoveryListResultResponse {
	return p.current
}

// AuthorizationRuleListResultPager provides iteration over AuthorizationRuleListResult pages.
type AuthorizationRuleListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current AuthorizationRuleListResultResponse.
	PageResponse() *AuthorizationRuleListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type authorizationRuleListResultHandleResponse func(*azcore.Response) (*AuthorizationRuleListResultResponse, error)

type authorizationRuleListResultAdvancePage func(*AuthorizationRuleListResultResponse) (*azcore.Request, error)

type authorizationRuleListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder authorizationRuleListResultHandleResponse
	// callback for advancing to the next page
	advancer authorizationRuleListResultAdvancePage
	// contains the current response
	current *AuthorizationRuleListResultResponse
	// any error encountered
	err error
}

func (p *authorizationRuleListResultPager) Err() error {
	return p.err
}

func (p *authorizationRuleListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.AuthorizationRuleListResult.NextLink == nil || len(*p.current.AuthorizationRuleListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *authorizationRuleListResultPager) PageResponse() *AuthorizationRuleListResultResponse {
	return p.current
}

// ConsumerGroupListResultPager provides iteration over ConsumerGroupListResult pages.
type ConsumerGroupListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current ConsumerGroupListResultResponse.
	PageResponse() *ConsumerGroupListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type consumerGroupListResultHandleResponse func(*azcore.Response) (*ConsumerGroupListResultResponse, error)

type consumerGroupListResultAdvancePage func(*ConsumerGroupListResultResponse) (*azcore.Request, error)

type consumerGroupListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder consumerGroupListResultHandleResponse
	// callback for advancing to the next page
	advancer consumerGroupListResultAdvancePage
	// contains the current response
	current *ConsumerGroupListResultResponse
	// any error encountered
	err error
}

func (p *consumerGroupListResultPager) Err() error {
	return p.err
}

func (p *consumerGroupListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.ConsumerGroupListResult.NextLink == nil || len(*p.current.ConsumerGroupListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *consumerGroupListResultPager) PageResponse() *ConsumerGroupListResultResponse {
	return p.current
}

// EhNamespaceListResultPager provides iteration over EhNamespaceListResult pages.
type EhNamespaceListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current EhNamespaceListResultResponse.
	PageResponse() *EhNamespaceListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type ehNamespaceListResultHandleResponse func(*azcore.Response) (*EhNamespaceListResultResponse, error)

type ehNamespaceListResultAdvancePage func(*EhNamespaceListResultResponse) (*azcore.Request, error)

type ehNamespaceListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder ehNamespaceListResultHandleResponse
	// callback for advancing to the next page
	advancer ehNamespaceListResultAdvancePage
	// contains the current response
	current *EhNamespaceListResultResponse
	// any error encountered
	err error
}

func (p *ehNamespaceListResultPager) Err() error {
	return p.err
}

func (p *ehNamespaceListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.EhNamespaceListResult.NextLink == nil || len(*p.current.EhNamespaceListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *ehNamespaceListResultPager) PageResponse() *EhNamespaceListResultResponse {
	return p.current
}

// EventHubListResultPager provides iteration over EventHubListResult pages.
type EventHubListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current EventHubListResultResponse.
	PageResponse() *EventHubListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type eventHubListResultHandleResponse func(*azcore.Response) (*EventHubListResultResponse, error)

type eventHubListResultAdvancePage func(*EventHubListResultResponse) (*azcore.Request, error)

type eventHubListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder eventHubListResultHandleResponse
	// callback for advancing to the next page
	advancer eventHubListResultAdvancePage
	// contains the current response
	current *EventHubListResultResponse
	// any error encountered
	err error
}

func (p *eventHubListResultPager) Err() error {
	return p.err
}

func (p *eventHubListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.EventHubListResult.NextLink == nil || len(*p.current.EventHubListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *eventHubListResultPager) PageResponse() *EventHubListResultResponse {
	return p.current
}

// MessagingRegionsListResultPager provides iteration over MessagingRegionsListResult pages.
type MessagingRegionsListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current MessagingRegionsListResultResponse.
	PageResponse() *MessagingRegionsListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type messagingRegionsListResultHandleResponse func(*azcore.Response) (*MessagingRegionsListResultResponse, error)

type messagingRegionsListResultAdvancePage func(*MessagingRegionsListResultResponse) (*azcore.Request, error)

type messagingRegionsListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder messagingRegionsListResultHandleResponse
	// callback for advancing to the next page
	advancer messagingRegionsListResultAdvancePage
	// contains the current response
	current *MessagingRegionsListResultResponse
	// any error encountered
	err error
}

func (p *messagingRegionsListResultPager) Err() error {
	return p.err
}

func (p *messagingRegionsListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.MessagingRegionsListResult.NextLink == nil || len(*p.current.MessagingRegionsListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *messagingRegionsListResultPager) PageResponse() *MessagingRegionsListResultResponse {
	return p.current
}

// NetworkRuleSetListResultPager provides iteration over NetworkRuleSetListResult pages.
type NetworkRuleSetListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current NetworkRuleSetListResultResponse.
	PageResponse() *NetworkRuleSetListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type networkRuleSetListResultHandleResponse func(*azcore.Response) (*NetworkRuleSetListResultResponse, error)

type networkRuleSetListResultAdvancePage func(*NetworkRuleSetListResultResponse) (*azcore.Request, error)

type networkRuleSetListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder networkRuleSetListResultHandleResponse
	// callback for advancing to the next page
	advancer networkRuleSetListResultAdvancePage
	// contains the current response
	current *NetworkRuleSetListResultResponse
	// any error encountered
	err error
}

func (p *networkRuleSetListResultPager) Err() error {
	return p.err
}

func (p *networkRuleSetListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.NetworkRuleSetListResult.NextLink == nil || len(*p.current.NetworkRuleSetListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *networkRuleSetListResultPager) PageResponse() *NetworkRuleSetListResultResponse {
	return p.current
}

// OperationListResultPager provides iteration over OperationListResult pages.
type OperationListResultPager interface {
	// NextPage returns true if the pager advanced to the next page.
	// Returns false if there are no more pages or an error occurred.
	NextPage(context.Context) bool

	// Page returns the current OperationListResultResponse.
	PageResponse() *OperationListResultResponse

	// Err returns the last error encountered while paging.
	Err() error
}

type operationListResultHandleResponse func(*azcore.Response) (*OperationListResultResponse, error)

type operationListResultAdvancePage func(*OperationListResultResponse) (*azcore.Request, error)

type operationListResultPager struct {
	// the pipeline for making the request
	pipeline azcore.Pipeline
	// contains the pending request
	request *azcore.Request
	// callback for handling the HTTP response
	responder operationListResultHandleResponse
	// callback for advancing to the next page
	advancer operationListResultAdvancePage
	// contains the current response
	current *OperationListResultResponse
	// any error encountered
	err error
}

func (p *operationListResultPager) Err() error {
	return p.err
}

func (p *operationListResultPager) NextPage(ctx context.Context) bool {
	if p.current != nil {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err := p.advancer(p.current)
		if err != nil {
			p.err = err
			return false
		}
		p.request = req
	}
	resp, err := p.pipeline.Do(ctx, p.request)
	if err != nil {
		p.err = err
		return false
	}
	result, err := p.responder(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *operationListResultPager) PageResponse() *OperationListResultResponse {
	return p.current
}
