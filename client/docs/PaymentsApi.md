# \PaymentsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentsBillingInfo**](PaymentsApi.md#PaymentsBillingInfo) | **Get** /api/v1/payments/billing-info | Get billing info for organization
[**PaymentsCardinfo**](PaymentsApi.md#PaymentsCardinfo) | **Get** /api/v1/payments/cardinfo | Get card information
[**PaymentsClear**](PaymentsApi.md#PaymentsClear) | **Post** /api/v1/payments/clear | Clear payment
[**PaymentsCreateCustomer**](PaymentsApi.md#PaymentsCreateCustomer) | **Post** /api/v1/payments/createcustomer | Create customer
[**PaymentsFinalPrice**](PaymentsApi.md#PaymentsFinalPrice) | **Post** /api/v1/payments/finalprice | Fetch final price
[**PaymentsGetStripeInvoices**](PaymentsApi.md#PaymentsGetStripeInvoices) | **Get** /api/v1/payments/stripeinvoices/{subscriptionId} | 
[**PaymentsPay**](PaymentsApi.md#PaymentsPay) | **Post** /api/v1/payments/pay | Pay invoice
[**PaymentsUpdateCard**](PaymentsApi.md#PaymentsUpdateCard) | **Post** /api/v1/payments/updatecard | Update payment card
[**PaymentsWebhook**](PaymentsApi.md#PaymentsWebhook) | **Post** /api/v1/payments/webhook | Listen to payment webhook



## PaymentsBillingInfo

> BillingInfoDto PaymentsBillingInfo(ctx).Execute()

Get billing info for organization

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PaymentsBillingInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsBillingInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsBillingInfo`: BillingInfoDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsBillingInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsBillingInfoRequest struct via the builder pattern


### Return type

[**BillingInfoDto**](BillingInfoDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsCardinfo

> CardInformationDto PaymentsCardinfo(ctx).OrganizationId(organizationId).Execute()

Get card information

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PaymentsCardinfo(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsCardinfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsCardinfo`: CardInformationDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsCardinfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsCardinfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**CardInformationDto**](CardInformationDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsClear

> PaymentsClear(ctx).Body(body).Execute()

Clear payment

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsApi.PaymentsClear(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsClear``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsClearRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsCreateCustomer

> PaymentsCreateCustomer(ctx).CreateStripeCustomerCommand(createStripeCustomerCommand).Execute()

Create customer

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    createStripeCustomerCommand := *openapiclient.NewCreateStripeCustomerCommand() // CreateStripeCustomerCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsApi.PaymentsCreateCustomer(context.Background()).CreateStripeCustomerCommand(createStripeCustomerCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsCreateCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsCreateCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStripeCustomerCommand** | [**CreateStripeCustomerCommand**](CreateStripeCustomerCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsFinalPrice

> FinalPriceDto PaymentsFinalPrice(ctx).FinalPriceCommand(finalPriceCommand).Execute()

Fetch final price

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    finalPriceCommand := *openapiclient.NewFinalPriceCommand() // FinalPriceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PaymentsFinalPrice(context.Background()).FinalPriceCommand(finalPriceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsFinalPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsFinalPrice`: FinalPriceDto
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsFinalPrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsFinalPriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **finalPriceCommand** | [**FinalPriceCommand**](FinalPriceCommand.md) |  | 

### Return type

[**FinalPriceDto**](FinalPriceDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsGetStripeInvoices

> StripeInvoices PaymentsGetStripeInvoices(ctx, subscriptionId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    subscriptionId := "subscriptionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PaymentsGetStripeInvoices(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsGetStripeInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsGetStripeInvoices`: StripeInvoices
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsGetStripeInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsGetStripeInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StripeInvoices**](StripeInvoices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsPay

> map[string]interface{} PaymentsPay(ctx).PayInvoiceCommand(payInvoiceCommand).Execute()

Pay invoice

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    payInvoiceCommand := *openapiclient.NewPayInvoiceCommand() // PayInvoiceCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentsApi.PaymentsPay(context.Background()).PayInvoiceCommand(payInvoiceCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentsPay`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.PaymentsPay`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payInvoiceCommand** | [**PayInvoiceCommand**](PayInvoiceCommand.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsUpdateCard

> PaymentsUpdateCard(ctx).ChangeCardCommand(changeCardCommand).Execute()

Update payment card

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    changeCardCommand := *openapiclient.NewChangeCardCommand() // ChangeCardCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsApi.PaymentsUpdateCard(context.Background()).ChangeCardCommand(changeCardCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsUpdateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsUpdateCardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeCardCommand** | [**ChangeCardCommand**](ChangeCardCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentsWebhook

> PaymentsWebhook(ctx).Body(body).Execute()

Listen to payment webhook

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PaymentsApi.PaymentsWebhook(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.PaymentsWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentsWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

