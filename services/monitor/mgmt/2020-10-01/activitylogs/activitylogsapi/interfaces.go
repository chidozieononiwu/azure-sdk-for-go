package activitylogsapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/monitor/mgmt/2020-10-01/activitylogs"
	"github.com/Azure/go-autorest/autorest"
)

// ClientAPI contains the set of methods on the Client type.
type ClientAPI interface {
	List(ctx context.Context, filter string, selectParameter string) (result activitylogs.EventDataCollectionPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string) (result activitylogs.EventDataCollectionIterator, err error)
}

var _ ClientAPI = (*activitylogs.Client)(nil)

// TenantActivityLogsClientAPI contains the set of methods on the TenantActivityLogsClient type.
type TenantActivityLogsClientAPI interface {
	List(ctx context.Context, filter string, selectParameter string) (result activitylogs.EventDataCollectionPage, err error)
	ListComplete(ctx context.Context, filter string, selectParameter string) (result activitylogs.EventDataCollectionIterator, err error)
}

var _ TenantActivityLogsClientAPI = (*activitylogs.TenantActivityLogsClient)(nil)

// AlertsClientAPI contains the set of methods on the AlertsClient type.
type AlertsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRule activitylogs.AlertResource) (result activitylogs.AlertResource, err error)
	Delete(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result autorest.Response, err error)
	Get(ctx context.Context, resourceGroupName string, activityLogAlertName string) (result activitylogs.AlertResource, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result activitylogs.AlertRuleListPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result activitylogs.AlertRuleListIterator, err error)
	ListBySubscriptionID(ctx context.Context) (result activitylogs.AlertRuleListPage, err error)
	ListBySubscriptionIDComplete(ctx context.Context) (result activitylogs.AlertRuleListIterator, err error)
	Update(ctx context.Context, resourceGroupName string, activityLogAlertName string, activityLogAlertRulePatch activitylogs.AlertRulePatchObject) (result activitylogs.AlertResource, err error)
}

var _ AlertsClientAPI = (*activitylogs.AlertsClient)(nil)