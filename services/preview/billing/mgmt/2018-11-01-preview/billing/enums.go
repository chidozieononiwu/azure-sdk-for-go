package billing

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AccountType enumerates the values for account type.
type AccountType string

const (
	// AccountTypeEnrollment ...
	AccountTypeEnrollment AccountType = "Enrollment"
	// AccountTypeOrganization ...
	AccountTypeOrganization AccountType = "Organization"
)

// PossibleAccountTypeValues returns an array of possible values for the AccountType const type.
func PossibleAccountTypeValues() []AccountType {
	return []AccountType{AccountTypeEnrollment, AccountTypeOrganization}
}

// AddressValidationStatus enumerates the values for address validation status.
type AddressValidationStatus string

const (
	// Invalid ...
	Invalid AddressValidationStatus = "Invalid"
	// Valid ...
	Valid AddressValidationStatus = "Valid"
)

// PossibleAddressValidationStatusValues returns an array of possible values for the AddressValidationStatus const type.
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return []AddressValidationStatus{Invalid, Valid}
}

// EligibleProductType enumerates the values for eligible product type.
type EligibleProductType string

const (
	// AzureReservation ...
	AzureReservation EligibleProductType = "AzureReservation"
	// DevTestAzureSubscription ...
	DevTestAzureSubscription EligibleProductType = "DevTestAzureSubscription"
	// StandardAzureSubscription ...
	StandardAzureSubscription EligibleProductType = "StandardAzureSubscription"
)

// PossibleEligibleProductTypeValues returns an array of possible values for the EligibleProductType const type.
func PossibleEligibleProductTypeValues() []EligibleProductType {
	return []EligibleProductType{AzureReservation, DevTestAzureSubscription, StandardAzureSubscription}
}

// Frequency enumerates the values for frequency.
type Frequency string

const (
	// Monthly ...
	Monthly Frequency = "Monthly"
	// OneTime ...
	OneTime Frequency = "OneTime"
	// UsageBased ...
	UsageBased Frequency = "UsageBased"
)

// PossibleFrequencyValues returns an array of possible values for the Frequency const type.
func PossibleFrequencyValues() []Frequency {
	return []Frequency{Monthly, OneTime, UsageBased}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// CreditNote ...
	CreditNote Kind = "CreditNote"
	// Invoice ...
	Invoice Kind = "Invoice"
	// Receipt ...
	Receipt Kind = "Receipt"
	// VoidNote ...
	VoidNote Kind = "VoidNote"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{CreditNote, Invoice, Receipt, VoidNote}
}

// PaymentMethodType enumerates the values for payment method type.
type PaymentMethodType string

const (
	// ChequeWire ...
	ChequeWire PaymentMethodType = "ChequeWire"
	// Credits ...
	Credits PaymentMethodType = "Credits"
)

// PossiblePaymentMethodTypeValues returns an array of possible values for the PaymentMethodType const type.
func PossiblePaymentMethodTypeValues() []PaymentMethodType {
	return []PaymentMethodType{ChequeWire, Credits}
}

// ProductStatusType enumerates the values for product status type.
type ProductStatusType string

const (
	// Active ...
	Active ProductStatusType = "Active"
	// AutoRenew ...
	AutoRenew ProductStatusType = "AutoRenew"
	// Cancelled ...
	Cancelled ProductStatusType = "Cancelled"
	// Disabled ...
	Disabled ProductStatusType = "Disabled"
	// Expired ...
	Expired ProductStatusType = "Expired"
	// Expiring ...
	Expiring ProductStatusType = "Expiring"
	// Inactive ...
	Inactive ProductStatusType = "Inactive"
	// PastDue ...
	PastDue ProductStatusType = "PastDue"
)

// PossibleProductStatusTypeValues returns an array of possible values for the ProductStatusType const type.
func PossibleProductStatusTypeValues() []ProductStatusType {
	return []ProductStatusType{Active, AutoRenew, Cancelled, Disabled, Expired, Expiring, Inactive, PastDue}
}

// ProductTransferStatus enumerates the values for product transfer status.
type ProductTransferStatus string

const (
	// Completed ...
	Completed ProductTransferStatus = "Completed"
	// Failed ...
	Failed ProductTransferStatus = "Failed"
	// InProgress ...
	InProgress ProductTransferStatus = "InProgress"
	// NotStarted ...
	NotStarted ProductTransferStatus = "NotStarted"
)

// PossibleProductTransferStatusValues returns an array of possible values for the ProductTransferStatus const type.
func PossibleProductTransferStatusValues() []ProductTransferStatus {
	return []ProductTransferStatus{Completed, Failed, InProgress, NotStarted}
}

// ProductTransferValidationErrorCode enumerates the values for product transfer validation error code.
type ProductTransferValidationErrorCode string

const (
	// CrossBillingAccountNotAllowed ...
	CrossBillingAccountNotAllowed ProductTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	// DestinationBillingProfilePastDue ...
	DestinationBillingProfilePastDue ProductTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	// InsufficientPermissionOnDestination ...
	InsufficientPermissionOnDestination ProductTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	// InsufficientPermissionOnSource ...
	InsufficientPermissionOnSource ProductTransferValidationErrorCode = "InsufficientPermissionOnSource"
	// InvalidSource ...
	InvalidSource ProductTransferValidationErrorCode = "InvalidSource"
	// NotAvailableForDestinationMarket ...
	NotAvailableForDestinationMarket ProductTransferValidationErrorCode = "NotAvailableForDestinationMarket"
	// OneTimePurchaseProductTransferNotAllowed ...
	OneTimePurchaseProductTransferNotAllowed ProductTransferValidationErrorCode = "OneTimePurchaseProductTransferNotAllowed"
	// ProductNotActive ...
	ProductNotActive ProductTransferValidationErrorCode = "ProductNotActive"
	// ProductTypeNotSupported ...
	ProductTypeNotSupported ProductTransferValidationErrorCode = "ProductTypeNotSupported"
)

// PossibleProductTransferValidationErrorCodeValues returns an array of possible values for the ProductTransferValidationErrorCode const type.
func PossibleProductTransferValidationErrorCodeValues() []ProductTransferValidationErrorCode {
	return []ProductTransferValidationErrorCode{CrossBillingAccountNotAllowed, DestinationBillingProfilePastDue, InsufficientPermissionOnDestination, InsufficientPermissionOnSource, InvalidSource, NotAvailableForDestinationMarket, OneTimePurchaseProductTransferNotAllowed, ProductNotActive, ProductTypeNotSupported}
}

// ProductType enumerates the values for product type.
type ProductType string

const (
	// ProductTypeAzureReservation ...
	ProductTypeAzureReservation ProductType = "AzureReservation"
	// ProductTypeAzureSubscription ...
	ProductTypeAzureSubscription ProductType = "AzureSubscription"
)

// PossibleProductTypeValues returns an array of possible values for the ProductType const type.
func PossibleProductTypeValues() []ProductType {
	return []ProductType{ProductTypeAzureReservation, ProductTypeAzureSubscription}
}

// ReservationType enumerates the values for reservation type.
type ReservationType string

const (
	// Purchase ...
	Purchase ReservationType = "Purchase"
	// UsageCharge ...
	UsageCharge ReservationType = "Usage Charge"
)

// PossibleReservationTypeValues returns an array of possible values for the ReservationType const type.
func PossibleReservationTypeValues() []ReservationType {
	return []ReservationType{Purchase, UsageCharge}
}

// Status enumerates the values for status.
type Status string

const (
	// Approved ...
	Approved Status = "Approved"
	// Rejected ...
	Rejected Status = "Rejected"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Approved, Rejected}
}

// Status1 enumerates the values for status 1.
type Status1 string

const (
	// Status1Due ...
	Status1Due Status1 = "Due"
	// Status1Paid ...
	Status1Paid Status1 = "Paid"
	// Status1PastDue ...
	Status1PastDue Status1 = "PastDue"
	// Status1Void ...
	Status1Void Status1 = "Void"
)

// PossibleStatus1Values returns an array of possible values for the Status1 const type.
func PossibleStatus1Values() []Status1 {
	return []Status1{Status1Due, Status1Paid, Status1PastDue, Status1Void}
}

// SubscriptionStatusType enumerates the values for subscription status type.
type SubscriptionStatusType string

const (
	// SubscriptionStatusTypeAbandoned ...
	SubscriptionStatusTypeAbandoned SubscriptionStatusType = "Abandoned"
	// SubscriptionStatusTypeActive ...
	SubscriptionStatusTypeActive SubscriptionStatusType = "Active"
	// SubscriptionStatusTypeDeleted ...
	SubscriptionStatusTypeDeleted SubscriptionStatusType = "Deleted"
	// SubscriptionStatusTypeInactive ...
	SubscriptionStatusTypeInactive SubscriptionStatusType = "Inactive"
	// SubscriptionStatusTypeWarning ...
	SubscriptionStatusTypeWarning SubscriptionStatusType = "Warning"
)

// PossibleSubscriptionStatusTypeValues returns an array of possible values for the SubscriptionStatusType const type.
func PossibleSubscriptionStatusTypeValues() []SubscriptionStatusType {
	return []SubscriptionStatusType{SubscriptionStatusTypeAbandoned, SubscriptionStatusTypeActive, SubscriptionStatusTypeDeleted, SubscriptionStatusTypeInactive, SubscriptionStatusTypeWarning}
}

// SubscriptionTransferValidationErrorCode enumerates the values for subscription transfer validation error
// code.
type SubscriptionTransferValidationErrorCode string

const (
	// SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed ...
	SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed SubscriptionTransferValidationErrorCode = "CrossBillingAccountNotAllowed"
	// SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue ...
	SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue SubscriptionTransferValidationErrorCode = "DestinationBillingProfilePastDue"
	// SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination ...
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnDestination"
	// SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource ...
	SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource SubscriptionTransferValidationErrorCode = "InsufficientPermissionOnSource"
	// SubscriptionTransferValidationErrorCodeInvalidSource ...
	SubscriptionTransferValidationErrorCodeInvalidSource SubscriptionTransferValidationErrorCode = "InvalidSource"
	// SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket ...
	SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket SubscriptionTransferValidationErrorCode = "NotAvailableForDestinationMarket"
	// SubscriptionTransferValidationErrorCodeSubscriptionNotActive ...
	SubscriptionTransferValidationErrorCodeSubscriptionNotActive SubscriptionTransferValidationErrorCode = "SubscriptionNotActive"
	// SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported ...
	SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported SubscriptionTransferValidationErrorCode = "SubscriptionTypeNotSupported"
)

// PossibleSubscriptionTransferValidationErrorCodeValues returns an array of possible values for the SubscriptionTransferValidationErrorCode const type.
func PossibleSubscriptionTransferValidationErrorCodeValues() []SubscriptionTransferValidationErrorCode {
	return []SubscriptionTransferValidationErrorCode{SubscriptionTransferValidationErrorCodeCrossBillingAccountNotAllowed, SubscriptionTransferValidationErrorCodeDestinationBillingProfilePastDue, SubscriptionTransferValidationErrorCodeInsufficientPermissionOnDestination, SubscriptionTransferValidationErrorCodeInsufficientPermissionOnSource, SubscriptionTransferValidationErrorCodeInvalidSource, SubscriptionTransferValidationErrorCodeNotAvailableForDestinationMarket, SubscriptionTransferValidationErrorCodeSubscriptionNotActive, SubscriptionTransferValidationErrorCodeSubscriptionTypeNotSupported}
}

// TransactionTypeKind enumerates the values for transaction type kind.
type TransactionTypeKind string

const (
	// All ...
	All TransactionTypeKind = "all"
	// Reservation ...
	Reservation TransactionTypeKind = "reservation"
)

// PossibleTransactionTypeKindValues returns an array of possible values for the TransactionTypeKind const type.
func PossibleTransactionTypeKindValues() []TransactionTypeKind {
	return []TransactionTypeKind{All, Reservation}
}

// TransferStatus enumerates the values for transfer status.
type TransferStatus string

const (
	// TransferStatusCanceled ...
	TransferStatusCanceled TransferStatus = "Canceled"
	// TransferStatusCompleted ...
	TransferStatusCompleted TransferStatus = "Completed"
	// TransferStatusCompletedWithErrors ...
	TransferStatusCompletedWithErrors TransferStatus = "CompletedWithErrors"
	// TransferStatusDeclined ...
	TransferStatusDeclined TransferStatus = "Declined"
	// TransferStatusFailed ...
	TransferStatusFailed TransferStatus = "Failed"
	// TransferStatusInProgress ...
	TransferStatusInProgress TransferStatus = "InProgress"
	// TransferStatusPending ...
	TransferStatusPending TransferStatus = "Pending"
)

// PossibleTransferStatusValues returns an array of possible values for the TransferStatus const type.
func PossibleTransferStatusValues() []TransferStatus {
	return []TransferStatus{TransferStatusCanceled, TransferStatusCompleted, TransferStatusCompletedWithErrors, TransferStatusDeclined, TransferStatusFailed, TransferStatusInProgress, TransferStatusPending}
}

// UpdateAutoRenew enumerates the values for update auto renew.
type UpdateAutoRenew string

const (
	// False ...
	False UpdateAutoRenew = "false"
	// True ...
	True UpdateAutoRenew = "true"
)

// PossibleUpdateAutoRenewValues returns an array of possible values for the UpdateAutoRenew const type.
func PossibleUpdateAutoRenewValues() []UpdateAutoRenew {
	return []UpdateAutoRenew{False, True}
}
