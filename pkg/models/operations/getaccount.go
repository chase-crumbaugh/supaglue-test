// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
)

type GetAccountPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
}

type GetAccountQueryParams struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces
	Expand *string `queryParam:"style=form,explode=true,name=expand"`
}

type GetAccountHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type GetAccountRequest struct {
	PathParams  GetAccountPathParams
	QueryParams GetAccountQueryParams
	Headers     GetAccountHeaders
}

type GetAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Account
	Account *shared.Account
}
