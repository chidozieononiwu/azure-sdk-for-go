// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

type AccessPolicyUpdateKind string

const (
	AccessPolicyUpdateKindAdd     AccessPolicyUpdateKind = "add"
	AccessPolicyUpdateKindReplace AccessPolicyUpdateKind = "replace"
	AccessPolicyUpdateKindRemove  AccessPolicyUpdateKind = "remove"
)

func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return []AccessPolicyUpdateKind{
		AccessPolicyUpdateKindAdd,
		AccessPolicyUpdateKindReplace,
		AccessPolicyUpdateKindRemove,
	}
}

func (c AccessPolicyUpdateKind) ToPtr() *AccessPolicyUpdateKind {
	return &c
}

type CertificatePermissions string

const (
	CertificatePermissionsBackup         CertificatePermissions = "backup"
	CertificatePermissionsCreate         CertificatePermissions = "create"
	CertificatePermissionsDelete         CertificatePermissions = "delete"
	CertificatePermissionsDeleteissuers  CertificatePermissions = "deleteissuers"
	CertificatePermissionsGet            CertificatePermissions = "get"
	CertificatePermissionsGetissuers     CertificatePermissions = "getissuers"
	CertificatePermissionsImport         CertificatePermissions = "import"
	CertificatePermissionsList           CertificatePermissions = "list"
	CertificatePermissionsListissuers    CertificatePermissions = "listissuers"
	CertificatePermissionsManagecontacts CertificatePermissions = "managecontacts"
	CertificatePermissionsManageissuers  CertificatePermissions = "manageissuers"
	CertificatePermissionsPurge          CertificatePermissions = "purge"
	CertificatePermissionsRecover        CertificatePermissions = "recover"
	CertificatePermissionsRestore        CertificatePermissions = "restore"
	CertificatePermissionsSetissuers     CertificatePermissions = "setissuers"
	CertificatePermissionsUpdate         CertificatePermissions = "update"
)

func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return []CertificatePermissions{
		CertificatePermissionsBackup,
		CertificatePermissionsCreate,
		CertificatePermissionsDelete,
		CertificatePermissionsDeleteissuers,
		CertificatePermissionsGet,
		CertificatePermissionsGetissuers,
		CertificatePermissionsImport,
		CertificatePermissionsList,
		CertificatePermissionsListissuers,
		CertificatePermissionsManagecontacts,
		CertificatePermissionsManageissuers,
		CertificatePermissionsPurge,
		CertificatePermissionsRecover,
		CertificatePermissionsRestore,
		CertificatePermissionsSetissuers,
		CertificatePermissionsUpdate,
	}
}

func (c CertificatePermissions) ToPtr() *CertificatePermissions {
	return &c
}

// CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not.
type CreateMode string

const (
	CreateModeRecover CreateMode = "recover"
	CreateModeDefault CreateMode = "default"
)

func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeRecover,
		CreateModeDefault,
	}
}

func (c CreateMode) ToPtr() *CreateMode {
	return &c
}

type KeyPermissions string

const (
	KeyPermissionsBackup    KeyPermissions = "backup"
	KeyPermissionsCreate    KeyPermissions = "create"
	KeyPermissionsDecrypt   KeyPermissions = "decrypt"
	KeyPermissionsDelete    KeyPermissions = "delete"
	KeyPermissionsEncrypt   KeyPermissions = "encrypt"
	KeyPermissionsGet       KeyPermissions = "get"
	KeyPermissionsImport    KeyPermissions = "import"
	KeyPermissionsList      KeyPermissions = "list"
	KeyPermissionsPurge     KeyPermissions = "purge"
	KeyPermissionsRecover   KeyPermissions = "recover"
	KeyPermissionsRestore   KeyPermissions = "restore"
	KeyPermissionsSign      KeyPermissions = "sign"
	KeyPermissionsUnwrapKey KeyPermissions = "unwrapKey"
	KeyPermissionsUpdate    KeyPermissions = "update"
	KeyPermissionsVerify    KeyPermissions = "verify"
	KeyPermissionsWrapKey   KeyPermissions = "wrapKey"
)

func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{
		KeyPermissionsBackup,
		KeyPermissionsCreate,
		KeyPermissionsDecrypt,
		KeyPermissionsDelete,
		KeyPermissionsEncrypt,
		KeyPermissionsGet,
		KeyPermissionsImport,
		KeyPermissionsList,
		KeyPermissionsPurge,
		KeyPermissionsRecover,
		KeyPermissionsRestore,
		KeyPermissionsSign,
		KeyPermissionsUnwrapKey,
		KeyPermissionsUpdate,
		KeyPermissionsVerify,
		KeyPermissionsWrapKey,
	}
}

func (c KeyPermissions) ToPtr() *KeyPermissions {
	return &c
}

// NetworkRuleAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used after the bypass property has been evaluated.
type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{
		NetworkRuleActionAllow,
		NetworkRuleActionDeny,
	}
}

func (c NetworkRuleAction) ToPtr() *NetworkRuleAction {
	return &c
}

// NetworkRuleBypassOptions - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'.  If not specified the default is 'AzureServices'.
type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{
		NetworkRuleBypassOptionsAzureServices,
		NetworkRuleBypassOptionsNone,
	}
}

func (c NetworkRuleBypassOptions) ToPtr() *NetworkRuleBypassOptions {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateDisconnected,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusDisconnected,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

func (c PrivateEndpointServiceConnectionStatus) ToPtr() *PrivateEndpointServiceConnectionStatus {
	return &c
}

// Reason - The reason that a vault name could not be used. The Reason element is only returned if NameAvailable is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

func (c Reason) ToPtr() *Reason {
	return &c
}

type SecretPermissions string

const (
	SecretPermissionsBackup  SecretPermissions = "backup"
	SecretPermissionsDelete  SecretPermissions = "delete"
	SecretPermissionsGet     SecretPermissions = "get"
	SecretPermissionsList    SecretPermissions = "list"
	SecretPermissionsPurge   SecretPermissions = "purge"
	SecretPermissionsRecover SecretPermissions = "recover"
	SecretPermissionsRestore SecretPermissions = "restore"
	SecretPermissionsSet     SecretPermissions = "set"
)

func PossibleSecretPermissionsValues() []SecretPermissions {
	return []SecretPermissions{
		SecretPermissionsBackup,
		SecretPermissionsDelete,
		SecretPermissionsGet,
		SecretPermissionsList,
		SecretPermissionsPurge,
		SecretPermissionsRecover,
		SecretPermissionsRestore,
		SecretPermissionsSet,
	}
}

func (c SecretPermissions) ToPtr() *SecretPermissions {
	return &c
}

// SkuName - SKU name to specify whether the key vault is a standard vault or a premium vault.
type SkuName string

const (
	SkuNameStandard SkuName = "standard"
	SkuNamePremium  SkuName = "premium"
)

func PossibleSkuNameValues() []SkuName {
	return []SkuName{
		SkuNameStandard,
		SkuNamePremium,
	}
}

func (c SkuName) ToPtr() *SkuName {
	return &c
}

type StoragePermissions string

const (
	StoragePermissionsBackup        StoragePermissions = "backup"
	StoragePermissionsDelete        StoragePermissions = "delete"
	StoragePermissionsDeletesas     StoragePermissions = "deletesas"
	StoragePermissionsGet           StoragePermissions = "get"
	StoragePermissionsGetsas        StoragePermissions = "getsas"
	StoragePermissionsList          StoragePermissions = "list"
	StoragePermissionsListsas       StoragePermissions = "listsas"
	StoragePermissionsPurge         StoragePermissions = "purge"
	StoragePermissionsRecover       StoragePermissions = "recover"
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	StoragePermissionsRestore       StoragePermissions = "restore"
	StoragePermissionsSet           StoragePermissions = "set"
	StoragePermissionsSetsas        StoragePermissions = "setsas"
	StoragePermissionsUpdate        StoragePermissions = "update"
)

func PossibleStoragePermissionsValues() []StoragePermissions {
	return []StoragePermissions{
		StoragePermissionsBackup,
		StoragePermissionsDelete,
		StoragePermissionsDeletesas,
		StoragePermissionsGet,
		StoragePermissionsGetsas,
		StoragePermissionsList,
		StoragePermissionsListsas,
		StoragePermissionsPurge,
		StoragePermissionsRecover,
		StoragePermissionsRegeneratekey,
		StoragePermissionsRestore,
		StoragePermissionsSet,
		StoragePermissionsSetsas,
		StoragePermissionsUpdate,
	}
}

func (c StoragePermissions) ToPtr() *StoragePermissions {
	return &c
}
