//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsupport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CheckNameAvailabilityForSupportTicketCommunication.json
func ExampleCommunicationsClient_CheckNameAvailability() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunicationsClient().CheckNameAvailability(ctx, "testticket", armsupport.CheckNameAvailabilityInput{
		Name: to.Ptr("sampleName"),
		Type: to.Ptr(armsupport.TypeMicrosoftSupportCommunications),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CheckNameAvailabilityOutput = armsupport.CheckNameAvailabilityOutput{
	// 	Message: to.Ptr("Name not available"),
	// 	NameAvailable: to.Ptr(false),
	// 	Reason: to.Ptr("Name is already in use"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListCommunicationsForSubscriptionSupportTicket.json
func ExampleCommunicationsClient_NewListPager_listCommunicationsForASubscriptionSupportTicket() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunicationsClient().NewListPager("testticket", &armsupport.CommunicationsClientListOptions{Top: nil,
		Filter: nil,
	})
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
		// page.CommunicationsListResult = armsupport.CommunicationsListResult{
		// 	Value: []*armsupport.CommunicationDetails{
		// 		{
		// 			Name: to.Ptr("testmessage1"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage1"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("this is a test message"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-24T20:18:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("this is a test message"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testmessage2"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage2"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("test"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-29T10:53:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("test"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListWebCommunicationsForSubscriptionSupportTicketCreatedOnOrAfter.json
func ExampleCommunicationsClient_NewListPager_listWebCommunicationCreatedOnOrAfterASpecificDateForASubscriptionSupportTicket() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunicationsClient().NewListPager("testticket", &armsupport.CommunicationsClientListOptions{Top: nil,
		Filter: to.Ptr("communicationType eq 'web' and createdDate ge 2020-03-10T22:08:51Z"),
	})
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
		// page.CommunicationsListResult = armsupport.CommunicationsListResult{
		// 	Value: []*armsupport.CommunicationDetails{
		// 		{
		// 			Name: to.Ptr("testmessage1"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage1"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("this is a test message"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T20:18:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("this is a test message"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testmessage2"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage2"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("test"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-12T10:53:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("test"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/ListWebCommunicationsForSubscriptionSupportTicket.json
func ExampleCommunicationsClient_NewListPager_listWebCommunicationsForASubscriptionSupportTicket() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCommunicationsClient().NewListPager("testticket", &armsupport.CommunicationsClientListOptions{Top: nil,
		Filter: to.Ptr("communicationType eq 'web'"),
	})
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
		// page.CommunicationsListResult = armsupport.CommunicationsListResult{
		// 	Value: []*armsupport.CommunicationDetails{
		// 		{
		// 			Name: to.Ptr("testmessage1"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage1"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("this is a test message"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("this is a test message"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("testmessage2"),
		// 			Type: to.Ptr("Microsoft.Support/communications"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage2"),
		// 			Properties: &armsupport.CommunicationDetailsProperties{
		// 				Body: to.Ptr("test"),
		// 				CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
		// 				CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
		// 				CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-11T10:53:19.000Z"); return t}()),
		// 				Sender: to.Ptr("user@contoso.com"),
		// 				Subject: to.Ptr("test"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/GetCommunicationDetailsForSubscriptionSupportTicket.json
func ExampleCommunicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCommunicationsClient().Get(ctx, "testticket", "testmessage", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunicationDetails = armsupport.CommunicationDetails{
	// 	Name: to.Ptr("testmessage"),
	// 	Type: to.Ptr("Microsoft.Support/communications"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testmessage"),
	// 	Properties: &armsupport.CommunicationDetailsProperties{
	// 		Body: to.Ptr("this is a test message"),
	// 		CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 		CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19.000Z"); return t}()),
	// 		Sender: to.Ptr("user@contoso.com"),
	// 		Subject: to.Ptr("this is a test message"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/support/resource-manager/Microsoft.Support/stable/2020-04-01/examples/CreateSupportTicketCommunication.json
func ExampleCommunicationsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsupport.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCommunicationsClient().BeginCreate(ctx, "testticket", "testcommunication", armsupport.CommunicationDetails{
		Properties: &armsupport.CommunicationDetailsProperties{
			Body:    to.Ptr("This is a test message from a customer!"),
			Sender:  to.Ptr("user@contoso.com"),
			Subject: to.Ptr("This is a test message from a customer!"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CommunicationDetails = armsupport.CommunicationDetails{
	// 	Name: to.Ptr("testcommunication"),
	// 	Type: to.Ptr("Microsoft.Support/communications"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.Support/supportTickets/testticket/communications/testcommunication"),
	// 	Properties: &armsupport.CommunicationDetailsProperties{
	// 		Body: to.Ptr("This is a test message from a customer!"),
	// 		CommunicationDirection: to.Ptr(armsupport.CommunicationDirectionOutbound),
	// 		CommunicationType: to.Ptr(armsupport.CommunicationTypeWeb),
	// 		CreatedDate: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-03-10T20:18:19.000Z"); return t}()),
	// 		Sender: to.Ptr("user@contoso.com"),
	// 		Subject: to.Ptr("This is a test message from a customer!"),
	// 	},
	// }
}
