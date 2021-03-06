package customerlockbox

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Decision enumerates the values for decision.
type Decision string

const (
	// Approve ...
	Approve Decision = "Approve"
	// Deny ...
	Deny Decision = "Deny"
)

// PossibleDecisionValues returns an array of possible values for the Decision const type.
func PossibleDecisionValues() []Decision {
	return []Decision{Approve, Deny}
}

// Status enumerates the values for status.
type Status string

const (
	// Approved ...
	Approved Status = "Approved"
	// Approving ...
	Approving Status = "Approving"
	// Completed ...
	Completed Status = "Completed"
	// Completing ...
	Completing Status = "Completing"
	// Denied ...
	Denied Status = "Denied"
	// Denying ...
	Denying Status = "Denying"
	// Error ...
	Error Status = "Error"
	// Expired ...
	Expired Status = "Expired"
	// Initializing ...
	Initializing Status = "Initializing"
	// Pending ...
	Pending Status = "Pending"
	// Revoked ...
	Revoked Status = "Revoked"
	// Revoking ...
	Revoking Status = "Revoking"
	// Unknown ...
	Unknown Status = "Unknown"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{Approved, Approving, Completed, Completing, Denied, Denying, Error, Expired, Initializing, Pending, Revoked, Revoking, Unknown}
}
