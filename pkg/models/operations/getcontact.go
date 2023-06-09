// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
)

type GetContactPathParams struct {
	ContactID string `pathParam:"style=simple,explode=false,name=contact_id"`
}

type GetContactQueryParams struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces
	Expand *string `queryParam:"style=form,explode=true,name=expand"`
}

type GetContactHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type GetContactRequest struct {
	PathParams  GetContactPathParams
	QueryParams GetContactQueryParams
	Headers     GetContactHeaders
}

type GetContactResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Contact
	Contact *shared.Contact
}
