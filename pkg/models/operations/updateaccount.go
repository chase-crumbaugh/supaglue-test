// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
)

type UpdateAccountPathParams struct {
	AccountID string `pathParam:"style=simple,explode=false,name=account_id"`
}

type UpdateAccountHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type UpdateAccountRequestBody struct {
	Model shared.CreateUpdateAccount `json:"model"`
}

type UpdateAccountRequest struct {
	PathParams UpdateAccountPathParams
	Headers    UpdateAccountHeaders
	Request    UpdateAccountRequestBody `request:"mediaType=application/json"`
}

// UpdateAccount200ApplicationJSON - Account updated
type UpdateAccount200ApplicationJSON struct {
	Errors   []shared.Errors   `json:"errors,omitempty"`
	Logs     []shared.Logs     `json:"logs,omitempty"`
	Model    *shared.Account   `json:"model,omitempty"`
	Warnings []shared.Warnings `json:"warnings,omitempty"`
}

type UpdateAccountResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Account updated
	UpdateAccount200ApplicationJSONObject *UpdateAccount200ApplicationJSON
}
