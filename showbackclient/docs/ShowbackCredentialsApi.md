# \ShowbackCredentialsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowbackcredentialsCreate**](ShowbackCredentialsApi.md#ShowbackcredentialsCreate) | **Post** /showback/v1/showbackcredentials/create | Create showback credential
[**ShowbackcredentialsDelete**](ShowbackCredentialsApi.md#ShowbackcredentialsDelete) | **Delete** /showback/v1/showbackcredentials/{id} | Delete showback credential
[**ShowbackcredentialsDropdown**](ShowbackCredentialsApi.md#ShowbackcredentialsDropdown) | **Get** /showback/v1/showbackcredentials/list | Retrieve showback credentials by organization id
[**ShowbackcredentialsList**](ShowbackCredentialsApi.md#ShowbackcredentialsList) | **Get** /showback/v1/showbackcredentials | Retrieve all showback credentials
[**ShowbackcredentialsLockManagement**](ShowbackCredentialsApi.md#ShowbackcredentialsLockManagement) | **Post** /showback/v1/showbackcredentials/lockmanager | Lock management for showback credential



## ShowbackcredentialsCreate

> ApiResponse ShowbackcredentialsCreate(ctx).CreateShowbackCredentialCommand(createShowbackCredentialCommand).Execute()

Create showback credential

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
    createShowbackCredentialCommand := *openapiclient.NewCreateShowbackCredentialCommand() // CreateShowbackCredentialCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackCredentialsApi.ShowbackcredentialsCreate(context.Background()).CreateShowbackCredentialCommand(createShowbackCredentialCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackCredentialsApi.ShowbackcredentialsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackcredentialsCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ShowbackCredentialsApi.ShowbackcredentialsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackcredentialsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShowbackCredentialCommand** | [**CreateShowbackCredentialCommand**](CreateShowbackCredentialCommand.md) |  | 

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


## ShowbackcredentialsDelete

> ShowbackcredentialsDelete(ctx, id).Execute()

Delete showback credential

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackCredentialsApi.ShowbackcredentialsDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackCredentialsApi.ShowbackcredentialsDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiShowbackcredentialsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackcredentialsDropdown

> []ShowbackCredentialsDetailsDto ShowbackcredentialsDropdown(ctx).OrganizationId(organizationId).Execute()

Retrieve showback credentials by organization id

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
    resp, r, err := apiClient.ShowbackCredentialsApi.ShowbackcredentialsDropdown(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackCredentialsApi.ShowbackcredentialsDropdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackcredentialsDropdown`: []ShowbackCredentialsDetailsDto
    fmt.Fprintf(os.Stdout, "Response from `ShowbackCredentialsApi.ShowbackcredentialsDropdown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackcredentialsDropdownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]ShowbackCredentialsDetailsDto**](ShowbackCredentialsDetailsDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackcredentialsList

> ShowbackCredentialsList ShowbackcredentialsList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all showback credentials

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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackCredentialsApi.ShowbackcredentialsList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackCredentialsApi.ShowbackcredentialsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackcredentialsList`: ShowbackCredentialsList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackCredentialsApi.ShowbackcredentialsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackcredentialsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**ShowbackCredentialsList**](ShowbackCredentialsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackcredentialsLockManagement

> ShowbackcredentialsLockManagement(ctx).ShowbackCredentialLockCommand(showbackCredentialLockCommand).Execute()

Lock management for showback credential

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
    showbackCredentialLockCommand := *openapiclient.NewShowbackCredentialLockCommand() // ShowbackCredentialLockCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackCredentialsApi.ShowbackcredentialsLockManagement(context.Background()).ShowbackCredentialLockCommand(showbackCredentialLockCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackCredentialsApi.ShowbackcredentialsLockManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackcredentialsLockManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showbackCredentialLockCommand** | [**ShowbackCredentialLockCommand**](ShowbackCredentialLockCommand.md) |  | 

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

