//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogz_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logz/armlogz"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_List.json
func ExampleSubAccountTagRulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubAccountTagRulesClient().NewListPager("myResourceGroup", "myMonitor", "SubAccount1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.MonitoringTagRulesListResponse = armlogz.MonitoringTagRulesListResponse{
		// 	Value: []*armlogz.MonitoringTagRules{
		// 		{
		// 			Name: to.Ptr("default"),
		// 			Type: to.Ptr("Microsoft.Logz/monitors/tagRules"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Logz/monitors/myMonitor/accounts/subAccount1/tagRules/default"),
		// 			Properties: &armlogz.MonitoringTagRulesProperties{
		// 				LogRules: &armlogz.LogRules{
		// 					FilteringTags: []*armlogz.FilteringTag{
		// 						{
		// 							Name: to.Ptr("Environment"),
		// 							Action: to.Ptr(armlogz.TagActionInclude),
		// 							Value: to.Ptr("Prod"),
		// 						},
		// 						{
		// 							Name: to.Ptr("Environment"),
		// 							Action: to.Ptr(armlogz.TagActionExclude),
		// 							Value: to.Ptr("Dev"),
		// 					}},
		// 					SendAADLogs: to.Ptr(false),
		// 					SendActivityLogs: to.Ptr(true),
		// 					SendSubscriptionLogs: to.Ptr(true),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_CreateOrUpdate.json
func ExampleSubAccountTagRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubAccountTagRulesClient().CreateOrUpdate(ctx, "myResourceGroup", "myMonitor", "SubAccount1", "default", &armlogz.SubAccountTagRulesClientCreateOrUpdateOptions{Body: nil})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitoringTagRules = armlogz.MonitoringTagRules{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors/tagRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Logz/monitors/myMonitor/accounts/SubAccount1/tagRules/default"),
	// 	Properties: &armlogz.MonitoringTagRulesProperties{
	// 		LogRules: &armlogz.LogRules{
	// 			FilteringTags: []*armlogz.FilteringTag{
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armlogz.TagActionInclude),
	// 					Value: to.Ptr("Prod"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armlogz.TagActionExclude),
	// 					Value: to.Ptr("Dev"),
	// 			}},
	// 			SendAADLogs: to.Ptr(false),
	// 			SendActivityLogs: to.Ptr(true),
	// 			SendSubscriptionLogs: to.Ptr(true),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_Get.json
func ExampleSubAccountTagRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubAccountTagRulesClient().Get(ctx, "myResourceGroup", "myMonitor", "SubAccount1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.MonitoringTagRules = armlogz.MonitoringTagRules{
	// 	Name: to.Ptr("default"),
	// 	Type: to.Ptr("Microsoft.Logz/monitors/tagRules"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Logz/monitors/myMonitor/tagRules/default"),
	// 	Properties: &armlogz.MonitoringTagRulesProperties{
	// 		LogRules: &armlogz.LogRules{
	// 			FilteringTags: []*armlogz.FilteringTag{
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armlogz.TagActionInclude),
	// 					Value: to.Ptr("Prod"),
	// 				},
	// 				{
	// 					Name: to.Ptr("Environment"),
	// 					Action: to.Ptr(armlogz.TagActionExclude),
	// 					Value: to.Ptr("Dev"),
	// 			}},
	// 			SendAADLogs: to.Ptr(false),
	// 			SendActivityLogs: to.Ptr(true),
	// 			SendSubscriptionLogs: to.Ptr(true),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logz/resource-manager/Microsoft.Logz/stable/2020-10-01/examples/SubAccountTagRules_Delete.json
func ExampleSubAccountTagRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogz.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSubAccountTagRulesClient().Delete(ctx, "myResourceGroup", "myMonitor", "SubAccount1", "default", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}