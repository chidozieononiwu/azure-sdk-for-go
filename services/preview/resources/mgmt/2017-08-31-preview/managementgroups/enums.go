package managementgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ChildType enumerates the values for child type.
type ChildType string

const (
	// Account ...
	Account ChildType = "Account"
	// Department ...
	Department ChildType = "Department"
	// Enrollment ...
	Enrollment ChildType = "Enrollment"
	// Subscription ...
	Subscription ChildType = "Subscription"
)

// PossibleChildTypeValues returns an array of possible values for the ChildType const type.
func PossibleChildTypeValues() []ChildType {
	return []ChildType{Account, Department, Enrollment, Subscription}
}

// ChildType1 enumerates the values for child type 1.
type ChildType1 string

const (
	// ChildType1Account ...
	ChildType1Account ChildType1 = "Account"
	// ChildType1Department ...
	ChildType1Department ChildType1 = "Department"
	// ChildType1Enrollment ...
	ChildType1Enrollment ChildType1 = "Enrollment"
	// ChildType1Subscription ...
	ChildType1Subscription ChildType1 = "Subscription"
)

// PossibleChildType1Values returns an array of possible values for the ChildType1 const type.
func PossibleChildType1Values() []ChildType1 {
	return []ChildType1{ChildType1Account, ChildType1Department, ChildType1Enrollment, ChildType1Subscription}
}

// ManagementGroupType enumerates the values for management group type.
type ManagementGroupType string

const (
	// ManagementGroupTypeAccount ...
	ManagementGroupTypeAccount ManagementGroupType = "Account"
	// ManagementGroupTypeDepartment ...
	ManagementGroupTypeDepartment ManagementGroupType = "Department"
	// ManagementGroupTypeEnrollment ...
	ManagementGroupTypeEnrollment ManagementGroupType = "Enrollment"
	// ManagementGroupTypeSubscription ...
	ManagementGroupTypeSubscription ManagementGroupType = "Subscription"
)

// PossibleManagementGroupTypeValues returns an array of possible values for the ManagementGroupType const type.
func PossibleManagementGroupTypeValues() []ManagementGroupType {
	return []ManagementGroupType{ManagementGroupTypeAccount, ManagementGroupTypeDepartment, ManagementGroupTypeEnrollment, ManagementGroupTypeSubscription}
}
