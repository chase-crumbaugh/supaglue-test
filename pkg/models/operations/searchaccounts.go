// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
)

type SearchAccountsQueryParams struct {
	// The pagination cursor value
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Number of results to return per page
	PageSize *string `queryParam:"style=form,explode=true,name=page_size"`
}

type SearchAccountsHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type SearchAccountsRequestBodyFilters struct {
	Website *shared.Filter `json:"website,omitempty"`
}

type SearchAccountsRequestBody struct {
	Filters SearchAccountsRequestBodyFilters `json:"filters"`
}

type SearchAccountsRequest struct {
	QueryParams SearchAccountsQueryParams
	Headers     SearchAccountsHeaders
	Request     SearchAccountsRequestBody `request:"mediaType=application/json"`
}

// SearchAccounts200ApplicationJSON - Accounts
type SearchAccounts200ApplicationJSON struct {
	Next     *string          `json:"next,omitempty"`
	Previous *string          `json:"previous,omitempty"`
	Results  []shared.Account `json:"results,omitempty"`
}

type SearchAccountsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Accounts
	SearchAccounts200ApplicationJSONObject *SearchAccounts200ApplicationJSON
}