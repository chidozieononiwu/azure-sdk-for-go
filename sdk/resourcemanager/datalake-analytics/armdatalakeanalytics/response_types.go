//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatalakeanalytics

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	NameAvailabilityInformation
}

// AccountsClientCreateResponse contains the response from method AccountsClient.BeginCreate.
type AccountsClientCreateResponse struct {
	Account
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	Account
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	AccountListResult
}

// AccountsClientListResponse contains the response from method AccountsClient.NewListPager.
type AccountsClientListResponse struct {
	AccountListResult
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	Account
}

// ComputePoliciesClientCreateOrUpdateResponse contains the response from method ComputePoliciesClient.CreateOrUpdate.
type ComputePoliciesClientCreateOrUpdateResponse struct {
	ComputePolicy
}

// ComputePoliciesClientDeleteResponse contains the response from method ComputePoliciesClient.Delete.
type ComputePoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// ComputePoliciesClientGetResponse contains the response from method ComputePoliciesClient.Get.
type ComputePoliciesClientGetResponse struct {
	ComputePolicy
}

// ComputePoliciesClientListByAccountResponse contains the response from method ComputePoliciesClient.NewListByAccountPager.
type ComputePoliciesClientListByAccountResponse struct {
	ComputePolicyListResult
}

// ComputePoliciesClientUpdateResponse contains the response from method ComputePoliciesClient.Update.
type ComputePoliciesClientUpdateResponse struct {
	ComputePolicy
}

// DataLakeStoreAccountsClientAddResponse contains the response from method DataLakeStoreAccountsClient.Add.
type DataLakeStoreAccountsClientAddResponse struct {
	// placeholder for future response values
}

// DataLakeStoreAccountsClientDeleteResponse contains the response from method DataLakeStoreAccountsClient.Delete.
type DataLakeStoreAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// DataLakeStoreAccountsClientGetResponse contains the response from method DataLakeStoreAccountsClient.Get.
type DataLakeStoreAccountsClientGetResponse struct {
	DataLakeStoreAccountInformation
}

// DataLakeStoreAccountsClientListByAccountResponse contains the response from method DataLakeStoreAccountsClient.NewListByAccountPager.
type DataLakeStoreAccountsClientListByAccountResponse struct {
	DataLakeStoreAccountInformationListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.CreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.Delete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByAccountResponse contains the response from method FirewallRulesClient.NewListByAccountPager.
type FirewallRulesClientListByAccountResponse struct {
	FirewallRuleListResult
}

// FirewallRulesClientUpdateResponse contains the response from method FirewallRulesClient.Update.
type FirewallRulesClientUpdateResponse struct {
	FirewallRule
}

// LocationsClientGetCapabilityResponse contains the response from method LocationsClient.GetCapability.
type LocationsClientGetCapabilityResponse struct {
	CapabilityInformation
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// StorageAccountsClientAddResponse contains the response from method StorageAccountsClient.Add.
type StorageAccountsClientAddResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientDeleteResponse contains the response from method StorageAccountsClient.Delete.
type StorageAccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// StorageAccountsClientGetResponse contains the response from method StorageAccountsClient.Get.
type StorageAccountsClientGetResponse struct {
	StorageAccountInformation
}

// StorageAccountsClientGetStorageContainerResponse contains the response from method StorageAccountsClient.GetStorageContainer.
type StorageAccountsClientGetStorageContainerResponse struct {
	StorageContainer
}

// StorageAccountsClientListByAccountResponse contains the response from method StorageAccountsClient.NewListByAccountPager.
type StorageAccountsClientListByAccountResponse struct {
	StorageAccountInformationListResult
}

// StorageAccountsClientListSasTokensResponse contains the response from method StorageAccountsClient.NewListSasTokensPager.
type StorageAccountsClientListSasTokensResponse struct {
	SasTokenInformationListResult
}

// StorageAccountsClientListStorageContainersResponse contains the response from method StorageAccountsClient.NewListStorageContainersPager.
type StorageAccountsClientListStorageContainersResponse struct {
	StorageContainerListResult
}

// StorageAccountsClientUpdateResponse contains the response from method StorageAccountsClient.Update.
type StorageAccountsClientUpdateResponse struct {
	// placeholder for future response values
}