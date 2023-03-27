<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "supaglue"
    "supaglue/pkg/models/shared"
    "supaglue/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKeyAuth: shared.SchemeAPIKeyAuth{
                APIKey: "YOUR_API_KEY_HERE",
            },
        }),
    )

    req := operations.CreateAccountRequest{
        Headers: operations.CreateAccountHeaders{
            XCustomerID: "my-customer-1",
            XProviderName: "salesforce",
        },
        Request: operations.CreateAccountRequestBody{
            Model: shared.CreateUpdateAccount{
                Addresses: []shared.Addresses{
                    shared.Addresses{
                        AddressType: "other",
                        City: "San Francisco",
                        Country: "USA",
                        PostalCode: "94107",
                        State: "CA",
                        Street1: "525 Brannan",
                        Street2: "null",
                    },
                    shared.Addresses{
                        AddressType: "billing",
                        City: "San Francisco",
                        Country: "USA",
                        PostalCode: "94107",
                        State: "CA",
                        Street1: "525 Brannan",
                        Street2: "null",
                    },
                    shared.Addresses{
                        AddressType: "shipping",
                        City: "San Francisco",
                        Country: "USA",
                        PostalCode: "94107",
                        State: "CA",
                        Street1: "525 Brannan",
                        Street2: "null",
                    },
                },
                CustomFields: map[string]interface{}{
                    "vero": "perspiciatis",
                    "nulla": "nihil",
                    "fuga": "facilis",
                },
                Description: "Integration API",
                Industry: "API's",
                Name: "Sample Customer",
                NumberOfEmployees: 276000,
                OwnerID: "9f3e97fd-4d5d-4efc-959d-bbebfac079f5",
                PhoneNumbers: []shared.PhoneNumbers{
                    shared.PhoneNumbers{
                        PhoneNumber: "+14151234567",
                        PhoneNumberType: "mobile",
                    },
                    shared.PhoneNumbers{
                        PhoneNumber: "+14151234567",
                        PhoneNumberType: "primary",
                    },
                },
                Website: "https://supaglue.com/",
            },
        },
    }

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateAccount201ApplicationJSONObject != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->