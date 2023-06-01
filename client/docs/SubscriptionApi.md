# \SubscriptionApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionBind**](SubscriptionApi.md#SubscriptionBind) | **Post** /api/v1/subscription/bind | Bind subscription
[**SubscriptionBoundList**](SubscriptionApi.md#SubscriptionBoundList) | **Get** /api/v1/subscription/boundlist | Retrieve subscription for organization bound
[**SubscriptionDelete**](SubscriptionApi.md#SubscriptionDelete) | **Post** /api/v1/subscription/delete | Delete subscription package
[**SubscriptionList**](SubscriptionApi.md#SubscriptionList) | **Get** /api/v1/subscription | Retrieve private subscription list for partners
[**SubscriptionPublic**](SubscriptionApi.md#SubscriptionPublic) | **Get** /api/v1/subscription/public | Retrieve subscription for organization bound
[**SubscriptionSubscription**](SubscriptionApi.md#SubscriptionSubscription) | **Post** /api/v1/subscription/create | Add new subscription package
[**SubscriptionUpdate**](SubscriptionApi.md#SubscriptionUpdate) | **Post** /api/v1/subscription/update | Update subscription



## SubscriptionBind

> BindSubscriptionResponseDto SubscriptionBind(ctx).BindSubscriptionCommand(bindSubscriptionCommand).Execute()

Bind subscription

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
    bindSubscriptionCommand := *openapiclient.NewBindSubscriptionCommand(int32(123)) // BindSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.SubscriptionBind(context.Background()).BindSubscriptionCommand(bindSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionBind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionBind`: BindSubscriptionResponseDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.SubscriptionBind`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionBindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bindSubscriptionCommand** | [**BindSubscriptionCommand**](BindSubscriptionCommand.md) |  | 

### Return type

[**BindSubscriptionResponseDto**](BindSubscriptionResponseDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionBoundList

> []ListForOrganizationEditDto SubscriptionBoundList(ctx).Search(search).Execute()

Retrieve subscription for organization bound

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.SubscriptionBoundList(context.Background()).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionBoundList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionBoundList`: []ListForOrganizationEditDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.SubscriptionBoundList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionBoundListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** |  | 

### Return type

[**[]ListForOrganizationEditDto**](ListForOrganizationEditDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionDelete

> SubscriptionDelete(ctx).DeleteSubscriptionCommand(deleteSubscriptionCommand).Execute()

Delete subscription package

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
    deleteSubscriptionCommand := *openapiclient.NewDeleteSubscriptionCommand() // DeleteSubscriptionCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.SubscriptionDelete(context.Background()).DeleteSubscriptionCommand(deleteSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSubscriptionCommand** | [**DeleteSubscriptionCommand**](DeleteSubscriptionCommand.md) |  | 

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


## SubscriptionList

> map[string]interface{} SubscriptionList(ctx).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()

Retrieve private subscription list for partners

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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionApi.SubscriptionList(context.Background()).Offset(offset).Limit(limit).SortBy(sortBy).SortDirection(sortDirection).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionList`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.SubscriptionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionPublic

> []ListForLandingPageDto SubscriptionPublic(ctx).Execute()

Retrieve subscription for organization bound

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
    resp, r, err := apiClient.SubscriptionApi.SubscriptionPublic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionPublic`: []ListForLandingPageDto
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionApi.SubscriptionPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionPublicRequest struct via the builder pattern


### Return type

[**[]ListForLandingPageDto**](ListForLandingPageDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionSubscription

> SubscriptionSubscription(ctx).CreateSubscriptionCommand(createSubscriptionCommand).Execute()

Add new subscription package

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
    createSubscriptionCommand := *openapiclient.NewCreateSubscriptionCommand("Name_example") // CreateSubscriptionCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.SubscriptionSubscription(context.Background()).CreateSubscriptionCommand(createSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubscriptionCommand** | [**CreateSubscriptionCommand**](CreateSubscriptionCommand.md) |  | 

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


## SubscriptionUpdate

> SubscriptionUpdate(ctx).UpdateSubscriptionCommand(updateSubscriptionCommand).Execute()

Update subscription

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
    updateSubscriptionCommand := *openapiclient.NewUpdateSubscriptionCommand("Name_example") // UpdateSubscriptionCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubscriptionApi.SubscriptionUpdate(context.Background()).UpdateSubscriptionCommand(updateSubscriptionCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionApi.SubscriptionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSubscriptionCommand** | [**UpdateSubscriptionCommand**](UpdateSubscriptionCommand.md) |  | 

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

