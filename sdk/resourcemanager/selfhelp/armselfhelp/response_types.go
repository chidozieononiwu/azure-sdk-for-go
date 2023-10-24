//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armselfhelp

// CheckNameAvailabilityClientPostResponse contains the response from method CheckNameAvailabilityClient.Post.
type CheckNameAvailabilityClientPostResponse struct {
	// Response for whether the requested resource name is available or not.
	CheckNameAvailabilityResponse
}

// DiagnosticsClientCreateResponse contains the response from method DiagnosticsClient.BeginCreate.
type DiagnosticsClientCreateResponse struct {
	// Diagnostic resource
	DiagnosticResource
}

// DiagnosticsClientGetResponse contains the response from method DiagnosticsClient.Get.
type DiagnosticsClientGetResponse struct {
	// Diagnostic resource
	DiagnosticResource
}

// DiscoverySolutionClientListResponse contains the response from method DiscoverySolutionClient.NewListPager.
type DiscoverySolutionClientListResponse struct {
	// Discovery response.
	DiscoveryResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// SolutionClientCreateResponse contains the response from method SolutionClient.BeginCreate.
type SolutionClientCreateResponse struct {
	// Solution response.
	SolutionResource
}

// SolutionClientGetResponse contains the response from method SolutionClient.Get.
type SolutionClientGetResponse struct {
	// Solution response.
	SolutionResource
}

// SolutionClientUpdateResponse contains the response from method SolutionClient.BeginUpdate.
type SolutionClientUpdateResponse struct {
	// Solution response.
	SolutionResource
}

// TroubleshootersClientContinueResponse contains the response from method TroubleshootersClient.Continue.
type TroubleshootersClientContinueResponse struct {
	// Location contains the information returned from the Location header response.
	Location *string
}

// TroubleshootersClientCreateResponse contains the response from method TroubleshootersClient.Create.
type TroubleshootersClientCreateResponse struct {
	// Troubleshooter response.
	TroubleshooterResource
}

// TroubleshootersClientEndResponse contains the response from method TroubleshootersClient.End.
type TroubleshootersClientEndResponse struct {
	// Location contains the information returned from the Location header response.
	Location *string
}

// TroubleshootersClientGetResponse contains the response from method TroubleshootersClient.Get.
type TroubleshootersClientGetResponse struct {
	// Troubleshooter response.
	TroubleshooterResource
}

// TroubleshootersClientRestartResponse contains the response from method TroubleshootersClient.Restart.
type TroubleshootersClientRestartResponse struct {
	// Troubleshooter restart response
	RestartTroubleshooterResponse

	// Location contains the information returned from the Location header response.
	Location *string
}
