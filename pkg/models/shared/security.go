// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SchemeAPIKeyAuth struct {
	APIKey string `security:"name=x-api-key"`
}

type Security struct {
	APIKeyAuth SchemeAPIKeyAuth `security:"scheme,type=apiKey,subtype=header"`
}
