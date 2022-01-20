//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// ClientResourcesHistoryOptions contains the optional parameters for the Client.ResourcesHistory method.
type ClientResourcesHistoryOptions struct {
	// placeholder for future optional parameters
}

// ClientResourcesOptions contains the optional parameters for the Client.Resources method.
type ClientResourcesOptions struct {
	// placeholder for future optional parameters
}

// DateTimeInterval - An interval in time specifying the date and time for the inclusive start and exclusive end, i.e. [start,
// end).
type DateTimeInterval struct {
	// REQUIRED; A datetime indicating the exclusive/open end of the time interval, i.e. [start,end). Specifying an end that occurs
	// chronologically before start will result in an error.
	End *time.Time `json:"end,omitempty"`

	// REQUIRED; A datetime indicating the inclusive/closed start of the time interval, i.e. [start, end). Specifying a start
	// that occurs chronologically after end will result in an error.
	Start *time.Time `json:"start,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type DateTimeInterval.
func (d DateTimeInterval) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "end", d.End)
	populateTimeRFC3339(objectMap, "start", d.Start)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DateTimeInterval.
func (d *DateTimeInterval) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "end":
			err = unpopulateTimeRFC3339(val, &d.End)
			delete(rawMsg, key)
		case "start":
			err = unpopulateTimeRFC3339(val, &d.Start)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ErrorDetails - Error details.
type ErrorDetails struct {
	// REQUIRED; Error code identifying the specific error.
	Code *string `json:"code,omitempty"`

	// REQUIRED; A human readable error message.
	Message *string `json:"message,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
}

// MarshalJSON implements the json.Marshaller interface for type ErrorDetails.
func (e ErrorDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	if e.AdditionalProperties != nil {
		for key, val := range e.AdditionalProperties {
			objectMap[key] = val
		}
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorDetails.
func (e *ErrorDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, &e.Message)
			delete(rawMsg, key)
		default:
			if e.AdditionalProperties == nil {
				e.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				e.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FacetClassification provides polymorphic access to related types.
// Call the interface's GetFacet() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Facet, *FacetError, *FacetResult
type FacetClassification interface {
	// GetFacet returns the Facet content of the underlying type.
	GetFacet() *Facet
}

// Facet - A facet containing additional statistics on the response of a query. Can be either FacetResult or FacetError.
type Facet struct {
	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`

	// REQUIRED; Result type
	ResultType *string `json:"resultType,omitempty"`
}

// GetFacet implements the FacetClassification interface for type Facet.
func (f *Facet) GetFacet() *Facet { return f }

// FacetError - A facet whose execution resulted in an error.
type FacetError struct {
	// REQUIRED; An array containing detected facet errors with details.
	Errors []*ErrorDetails `json:"errors,omitempty"`

	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`

	// REQUIRED; Result type
	ResultType *string `json:"resultType,omitempty"`
}

// GetFacet implements the FacetClassification interface for type FacetError.
func (f *FacetError) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type FacetError.
func (f FacetError) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errors", f.Errors)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetError"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetError.
func (f *FacetError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errors":
			err = unpopulate(val, &f.Errors)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, &f.ResultType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// FacetRequest - A request to compute additional statistics (facets) over the query results.
type FacetRequest struct {
	// REQUIRED; The column or list of columns to summarize by
	Expression *string `json:"expression,omitempty"`

	// The options for facet evaluation
	Options *FacetRequestOptions `json:"options,omitempty"`
}

// FacetRequestOptions - The options for facet evaluation
type FacetRequestOptions struct {
	// Specifies the filter condition for the 'where' clause which will be run on main query's result, just before the actual
	// faceting.
	Filter *string `json:"filter,omitempty"`

	// The column name or query expression to sort on. Defaults to count if not present.
	SortBy *string `json:"sortBy,omitempty"`

	// The sorting order by the selected column (count by default).
	SortOrder *FacetSortOrder `json:"sortOrder,omitempty"`

	// The maximum number of facet rows that should be returned.
	Top *int32 `json:"$top,omitempty"`
}

// FacetResult - Successfully executed facet containing additional statistics on the response of a query.
type FacetResult struct {
	// REQUIRED; Number of records returned in the facet response.
	Count *int32 `json:"count,omitempty"`

	// REQUIRED; A JObject array or Table containing the desired facets. Only present if the facet is valid.
	Data interface{} `json:"data,omitempty"`

	// REQUIRED; Facet expression, same as in the corresponding facet request.
	Expression *string `json:"expression,omitempty"`

	// REQUIRED; Result type
	ResultType *string `json:"resultType,omitempty"`

	// REQUIRED; Number of total records in the facet results.
	TotalRecords *int64 `json:"totalRecords,omitempty"`
}

// GetFacet implements the FacetClassification interface for type FacetResult.
func (f *FacetResult) GetFacet() *Facet {
	return &Facet{
		Expression: f.Expression,
		ResultType: f.ResultType,
	}
}

// MarshalJSON implements the json.Marshaller interface for type FacetResult.
func (f FacetResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", f.Count)
	populate(objectMap, "data", &f.Data)
	populate(objectMap, "expression", f.Expression)
	objectMap["resultType"] = "FacetResult"
	populate(objectMap, "totalRecords", f.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetResult.
func (f *FacetResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &f.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, &f.Data)
			delete(rawMsg, key)
		case "expression":
			err = unpopulate(val, &f.Expression)
			delete(rawMsg, key)
		case "resultType":
			err = unpopulate(val, &f.ResultType)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, &f.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// Operation - Resource Graph REST API operation definition.
type Operation struct {
	// Display metadata associated with the operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// Operation name: {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`

	// The origin of operations.
	Origin *string `json:"origin,omitempty"`
}

// OperationDisplay - Display metadata associated with the operation.
type OperationDisplay struct {
	// Description for the operation.
	Description *string `json:"description,omitempty"`

	// Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Service provider: Microsoft Resource Graph.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed etc.
	Resource *string `json:"resource,omitempty"`
}

// OperationListResult - Result of the request to list Resource Graph operations. It contains a list of operations and a URL
// link to get the next set of results.
type OperationListResult struct {
	// List of Resource Graph operations supported by the Resource Graph resource provider.
	Value []*Operation `json:"value,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// QueryRequest - Describes a query to be executed.
type QueryRequest struct {
	// REQUIRED; The resources query.
	Query *string `json:"query,omitempty"`

	// An array of facet requests to be computed against the query result.
	Facets []*FacetRequest `json:"facets,omitempty"`

	// Azure management groups against which to execute the query. Example: [ 'mg1', 'mg2' ]
	ManagementGroups []*string `json:"managementGroups,omitempty"`

	// The query evaluation options
	Options *QueryRequestOptions `json:"options,omitempty"`

	// Azure subscriptions against which to execute the query.
	Subscriptions []*string `json:"subscriptions,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type QueryRequest.
func (q QueryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "managementGroups", q.ManagementGroups)
	populate(objectMap, "options", q.Options)
	populate(objectMap, "query", q.Query)
	populate(objectMap, "subscriptions", q.Subscriptions)
	return json.Marshal(objectMap)
}

// QueryRequestOptions - The options for query evaluation
type QueryRequestOptions struct {
	// Only applicable for tenant and management group level queries to decide whether to allow partial scopes for result in case
	// the number of subscriptions exceed allowed limits.
	AllowPartialScopes *bool `json:"allowPartialScopes,omitempty"`

	// Defines what level of authorization resources should be returned based on the which subscriptions and management groups
	// are passed as scopes.
	AuthorizationScopeFilter *AuthorizationScopeFilter `json:"authorizationScopeFilter,omitempty"`

	// Defines in which format query result returned.
	ResultFormat *ResultFormat `json:"resultFormat,omitempty"`

	// The number of rows to skip from the beginning of the results. Overrides the next page offset when $skipToken property is
	// present.
	Skip *int32 `json:"$skip,omitempty"`

	// Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string `json:"$skipToken,omitempty"`

	// The maximum number of rows that the query should return. Overrides the page size when $skipToken property is present.
	Top *int32 `json:"$top,omitempty"`
}

// QueryResponse - Query result.
type QueryResponse struct {
	// REQUIRED; Number of records returned in the current response. In the case of paging, this is the number of records in the
	// current page.
	Count *int64 `json:"count,omitempty"`

	// REQUIRED; Query output in JObject array or Table format.
	Data interface{} `json:"data,omitempty"`

	// REQUIRED; Indicates whether the query results are truncated.
	ResultTruncated *ResultTruncated `json:"resultTruncated,omitempty"`

	// REQUIRED; Number of total records matching the query.
	TotalRecords *int64 `json:"totalRecords,omitempty"`

	// Query facets.
	Facets []FacetClassification `json:"facets,omitempty"`

	// When present, the value can be passed to a subsequent query call (together with the same query and scopes used in the current
	// request) to retrieve the next page of data.
	SkipToken *string `json:"$skipToken,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type QueryResponse.
func (q QueryResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", q.Count)
	populate(objectMap, "data", &q.Data)
	populate(objectMap, "facets", q.Facets)
	populate(objectMap, "resultTruncated", q.ResultTruncated)
	populate(objectMap, "$skipToken", q.SkipToken)
	populate(objectMap, "totalRecords", q.TotalRecords)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type QueryResponse.
func (q *QueryResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, &q.Count)
			delete(rawMsg, key)
		case "data":
			err = unpopulate(val, &q.Data)
			delete(rawMsg, key)
		case "facets":
			q.Facets, err = unmarshalFacetClassificationArray(val)
			delete(rawMsg, key)
		case "resultTruncated":
			err = unpopulate(val, &q.ResultTruncated)
			delete(rawMsg, key)
		case "$skipToken":
			err = unpopulate(val, &q.SkipToken)
			delete(rawMsg, key)
		case "totalRecords":
			err = unpopulate(val, &q.TotalRecords)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// ResourcesHistoryRequest - Describes a history request to be executed.
type ResourcesHistoryRequest struct {
	// Azure management groups against which to execute the query. Example: [ 'mg1', 'mg2' ]
	ManagementGroups []*string `json:"managementGroups,omitempty"`

	// The history request evaluation options
	Options *ResourcesHistoryRequestOptions `json:"options,omitempty"`

	// The resources query.
	Query *string `json:"query,omitempty"`

	// Azure subscriptions against which to execute the query.
	Subscriptions []*string `json:"subscriptions,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type ResourcesHistoryRequest.
func (r ResourcesHistoryRequest) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "managementGroups", r.ManagementGroups)
	populate(objectMap, "options", r.Options)
	populate(objectMap, "query", r.Query)
	populate(objectMap, "subscriptions", r.Subscriptions)
	return json.Marshal(objectMap)
}

// ResourcesHistoryRequestOptions - The options for history request evaluation
type ResourcesHistoryRequestOptions struct {
	// The time interval used to fetch history.
	Interval *DateTimeInterval `json:"interval,omitempty"`

	// Defines in which format query result returned.
	ResultFormat *ResultFormat `json:"resultFormat,omitempty"`

	// The number of rows to skip from the beginning of the results. Overrides the next page offset when $skipToken property is
	// present.
	Skip *int32 `json:"$skip,omitempty"`

	// Continuation token for pagination, capturing the next page size and offset, as well as the context of the query.
	SkipToken *string `json:"$skipToken,omitempty"`

	// The maximum number of rows that the query should return. Overrides the page size when $skipToken property is present.
	Top *int32 `json:"$top,omitempty"`
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
