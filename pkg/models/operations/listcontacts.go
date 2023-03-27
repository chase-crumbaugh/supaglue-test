// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
	"time"
)

type ListContactsQueryParams struct {
	// If provided, will only return objects created after this datetime
	CreatedAfter *time.Time `queryParam:"style=form,explode=true,name=created_after"`
	// If provided, will only return objects created before this datetime
	CreatedBefore *time.Time `queryParam:"style=form,explode=true,name=created_before"`
	// The pagination cursor value
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces
	Expand *string `queryParam:"style=form,explode=true,name=expand"`
	// Number of results to return per page
	PageSize *string `queryParam:"style=form,explode=true,name=page_size"`
	// If provided, will only return objects modified after this datetime
	UpdatedAfter *time.Time `queryParam:"style=form,explode=true,name=updated_after"`
	// If provided, will only return objects modified before this datetime
	UpdatedBefore *time.Time `queryParam:"style=form,explode=true,name=updated_before"`
}

type ListContactsHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type ListContactsRequest struct {
	QueryParams ListContactsQueryParams
	Headers     ListContactsHeaders
}

// ListContacts200ApplicationJSON - Contacts
type ListContacts200ApplicationJSON struct {
	Next     *string          `json:"next,omitempty"`
	Previous *string          `json:"previous,omitempty"`
	Results  []shared.Contact `json:"results,omitempty"`
}

type ListContactsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Contacts
	ListContacts200ApplicationJSONObject *ListContacts200ApplicationJSON
}
