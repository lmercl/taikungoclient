# \SlackApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlackCreate**](SlackApi.md#SlackCreate) | **Post** /api/v1/slack/create | Create slack configuration
[**SlackDeleteMultiple**](SlackApi.md#SlackDeleteMultiple) | **Post** /api/v1/slack/delete-multiple | Delete slack configuration(s)
[**SlackDropdown**](SlackApi.md#SlackDropdown) | **Get** /api/v1/slack/list | Retrieve all slack configs for organization
[**SlackList**](SlackApi.md#SlackList) | **Get** /api/v1/slack | Retrieve all slack configs
[**SlackUpdate**](SlackApi.md#SlackUpdate) | **Put** /api/v1/slack/update/{id} | Update slack configuration
[**SlackVerify**](SlackApi.md#SlackVerify) | **Post** /api/v1/slack/verify | Verify slack configuration



## SlackCreate

> ApiResponse SlackCreate(ctx).CreateSlackConfigurationCommand(createSlackConfigurationCommand).Execute()

Create slack configuration

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
    createSlackConfigurationCommand := *openapiclient.NewCreateSlackConfigurationCommand("Name_example", "Url_example", "Channel_example") // CreateSlackConfigurationCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackApi.SlackCreate(context.Background()).CreateSlackConfigurationCommand(createSlackConfigurationCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlackCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `SlackApi.SlackCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSlackConfigurationCommand** | [**CreateSlackConfigurationCommand**](CreateSlackConfigurationCommand.md) |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackDeleteMultiple

> SlackDeleteMultiple(ctx).DeleteSlackConfigCommand(deleteSlackConfigCommand).Execute()

Delete slack configuration(s)

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
    deleteSlackConfigCommand := *openapiclient.NewDeleteSlackConfigCommand() // DeleteSlackConfigCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlackApi.SlackDeleteMultiple(context.Background()).DeleteSlackConfigCommand(deleteSlackConfigCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackDeleteMultiple``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackDeleteMultipleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSlackConfigCommand** | [**DeleteSlackConfigCommand**](DeleteSlackConfigCommand.md) |  | 

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


## SlackDropdown

> []CommonDropdownDto SlackDropdown(ctx).OrganizationId(organizationId).Search(search).Execute()

Retrieve all slack configs for organization

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
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackApi.SlackDropdown(context.Background()).OrganizationId(organizationId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlackDropdown`: []CommonDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `SlackApi.SlackDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 

### Return type

[**[]CommonDropdownDto**](CommonDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackList

> SlackConfigurationList SlackList(ctx).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all slack configs

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SlackApi.SlackList(context.Background()).OrganizationId(organizationId).Limit(limit).Offset(offset).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlackList`: SlackConfigurationList
    fmt.Fprintf(os.Stdout, "Response from `SlackApi.SlackList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**SlackConfigurationList**](SlackConfigurationList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlackUpdate

> SlackUpdate(ctx, id).UpdateSlackConfigurationDto(updateSlackConfigurationDto).Execute()

Update slack configuration

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
    id := int32(56) // int32 | 
    updateSlackConfigurationDto := *openapiclient.NewUpdateSlackConfigurationDto() // UpdateSlackConfigurationDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlackApi.SlackUpdate(context.Background(), id).UpdateSlackConfigurationDto(updateSlackConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlackUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSlackConfigurationDto** | [**UpdateSlackConfigurationDto**](UpdateSlackConfigurationDto.md) |  | 

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


## SlackVerify

> SlackVerify(ctx).VerifySlackCredentialsCommand(verifySlackCredentialsCommand).Execute()

Verify slack configuration

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
    verifySlackCredentialsCommand := *openapiclient.NewVerifySlackCredentialsCommand() // VerifySlackCredentialsCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SlackApi.SlackVerify(context.Background()).VerifySlackCredentialsCommand(verifySlackCredentialsCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlackApi.SlackVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlackVerifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySlackCredentialsCommand** | [**VerifySlackCredentialsCommand**](VerifySlackCredentialsCommand.md) |  | 

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

