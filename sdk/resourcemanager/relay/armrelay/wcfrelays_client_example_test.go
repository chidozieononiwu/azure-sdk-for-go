//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrelay_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/relay/armrelay"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleListAll.json
func ExampleWCFRelaysClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWCFRelaysClient().NewListAuthorizationRulesPager("resourcegroup", "example-RelayNamespace-01", "example-Relay-Wcf-01", nil)
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
		// page.AuthorizationRuleListResult = armrelay.AuthorizationRuleListResult{
		// 	Value: []*armrelay.AuthorizationRule{
		// 		{
		// 			Name: to.Ptr("example-RelayAuthRules-01"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelay/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01/AuthorizationRules/example-RelayAuthRules-01"),
		// 			Properties: &armrelay.AuthorizationRuleProperties{
		// 				Rights: []*armrelay.AccessRights{
		// 					to.Ptr(armrelay.AccessRightsListen),
		// 					to.Ptr(armrelay.AccessRightsSend)},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleCreate.json
func ExampleWCFRelaysClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().CreateOrUpdateAuthorizationRule(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", armrelay.AuthorizationRule{
		Properties: &armrelay.AuthorizationRuleProperties{
			Rights: []*armrelay.AccessRights{
				to.Ptr(armrelay.AccessRightsListen),
				to.Ptr(armrelay.AccessRightsSend)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armrelay.AuthorizationRule{
	// 	Name: to.Ptr("example-RelayAuthRules-01"),
	// 	Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelay/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01/AuthorizationRules/example-RelayAuthRules-01"),
	// 	Properties: &armrelay.AuthorizationRuleProperties{
	// 		Rights: []*armrelay.AccessRights{
	// 			to.Ptr(armrelay.AccessRightsListen),
	// 			to.Ptr(armrelay.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleDelete.json
func ExampleWCFRelaysClient_DeleteAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWCFRelaysClient().DeleteAuthorizationRule(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleGet.json
func ExampleWCFRelaysClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().GetAuthorizationRule(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armrelay.AuthorizationRule{
	// 	Name: to.Ptr("example-RelayAuthRules-01"),
	// 	Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelay/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01/AuthorizationRules/example-RelayAuthRules-01"),
	// 	Properties: &armrelay.AuthorizationRuleProperties{
	// 		Rights: []*armrelay.AccessRights{
	// 			to.Ptr(armrelay.AccessRightsListen)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleListKey.json
func ExampleWCFRelaysClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().ListKeys(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armrelay.AccessKeys{
	// 	KeyName: to.Ptr("example-RelayAuthRules-01"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayAuthorizationRuleRegenerateKey.json
func ExampleWCFRelaysClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().RegenerateKeys(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", "example-RelayAuthRules-01", armrelay.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armrelay.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armrelay.AccessKeys{
	// 	KeyName: to.Ptr("example-RelayAuthRules-01"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://example-Relaynamespace-01.servicebus.windows.net/;SharedAccessKeyName=example-RelayAuthRules-01;SharedAccessKey=############################################"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayListAll.json
func ExampleWCFRelaysClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWCFRelaysClient().NewListByNamespacePager("resourcegroup", "example-RelayNamespace-01", nil)
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
		// page.WcfRelaysListResult = armrelay.WcfRelaysListResult{
		// 	Value: []*armrelay.WcfRelay{
		// 		{
		// 			Name: to.Ptr("example-Relay-Wcf-01"),
		// 			Type: to.Ptr("Microsoft.Relay/Namespaces/WcfRelays"),
		// 			ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/RG1-eg/providers/Microsoft.Relay/namespaces/example-RelayNamespace-01/WcfRelays/example-Relay-Wcf-01"),
		// 			Properties: &armrelay.WcfRelayProperties{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T00:46:27.0049983Z"); return t}()),
		// 				IsDynamic: to.Ptr(false),
		// 				RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
		// 				RequiresClientAuthorization: to.Ptr(true),
		// 				RequiresTransportSecurity: to.Ptr(true),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-01-24T00:46:27.0049983Z"); return t}()),
		// 				UserMetadata: to.Ptr("usermetadata is a placeholder to store user-defined string data for the HybridConnection endpoint.e.g. it can be used to store  descriptive data, such as list of teams and their contact information also user-defined configuration settings can be stored"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayCreate.json
func ExampleWCFRelaysClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().CreateOrUpdate(ctx, "resourcegroup", "example-RelayNamespace-9953", "example-Relay-Wcf-1194", armrelay.WcfRelay{
		Properties: &armrelay.WcfRelayProperties{
			RelayType:                   to.Ptr(armrelay.RelaytypeNetTCP),
			RequiresClientAuthorization: to.Ptr(true),
			RequiresTransportSecurity:   to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WcfRelay = armrelay.WcfRelay{
	// 	Name: to.Ptr("example-Relay-Wcf-1194"),
	// 	Type: to.Ptr("Microsoft.Relay/WcfRelays"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/resourcegroup/providers/Microsoft.Relay/namespaces/example-RelayNamespace-9953/WcfRelays/example-Relay-Wcf-1194"),
	// 	Properties: &armrelay.WcfRelayProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 		IsDynamic: to.Ptr(false),
	// 		RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
	// 		RequiresClientAuthorization: to.Ptr(true),
	// 		RequiresTransportSecurity: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayDelete.json
func ExampleWCFRelaysClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewWCFRelaysClient().Delete(ctx, "resourcegroup", "example-RelayNamespace-01", "example-Relay-wcf-01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/relay/resource-manager/Microsoft.Relay/stable/2021-11-01/examples/Relay/RelayGet.json
func ExampleWCFRelaysClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armrelay.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewWCFRelaysClient().Get(ctx, "resourcegroup", "example-RelayNamespace-9953", "example-Relay-Wcf-1194", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WcfRelay = armrelay.WcfRelay{
	// 	Name: to.Ptr("example-Relay-Wcf-1194"),
	// 	Type: to.Ptr("Microsoft.Relay/WcfRelays"),
	// 	ID: to.Ptr("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/Default-ServiceBus-WestUS/providers/Microsoft.Relay/namespaces/example-RelayNamespace-9953/WcfRelays/example-Relay-Wcf-1194"),
	// 	Properties: &armrelay.WcfRelayProperties{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 		IsDynamic: to.Ptr(false),
	// 		ListenerCount: to.Ptr[int32](0),
	// 		RelayType: to.Ptr(armrelay.RelaytypeNetTCP),
	// 		RequiresClientAuthorization: to.Ptr(true),
	// 		RequiresTransportSecurity: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-16T00:26:17.5014661Z"); return t}()),
	// 	},
	// }
}