// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"supaglue/pkg/models/shared"
)

type CreateLeadHeaders struct {
	// The customer ID that uniquely identifies the customer in your application
	XCustomerID string `header:"style=simple,explode=false,name=x-customer-id"`
	// The provider name
	XProviderName string `header:"style=simple,explode=false,name=x-provider-name"`
}

type CreateLeadRequestBody struct {
	Model shared.CreateUpdateLead `json:"model"`
}

type CreateLeadRequest struct {
	Headers CreateLeadHeaders
	Request CreateLeadRequestBody `request:"mediaType=application/json"`
}

// CreateLead201ApplicationJSON - Lead created
type CreateLead201ApplicationJSON struct {
	Errors   []shared.Errors   `json:"errors,omitempty"`
	Logs     []shared.Logs     `json:"logs,omitempty"`
	Model    *shared.Lead      `json:"model,omitempty"`
	Warnings []shared.Warnings `json:"warnings,omitempty"`
}

type CreateLeadResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Lead created
	CreateLead201ApplicationJSONObject *CreateLead201ApplicationJSON
}