# \TanzuApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TanzuCreate**](TanzuApi.md#TanzuCreate) | **Post** /api/v1/tanzu/create | Create tanzu credentials
[**TanzuKubernetesVersions**](TanzuApi.md#TanzuKubernetesVersions) | **Post** /api/v1/tanzu/kubernetes-versions/{cloudId} | Tanzu available k8s version list
[**TanzuList**](TanzuApi.md#TanzuList) | **Post** /api/v1/tanzu/list | Retrieve list of tanzu cloud credentials
[**TanzuStorageList**](TanzuApi.md#TanzuStorageList) | **Post** /api/v1/tanzu/storage-list | Tanzu storage list
[**TanzuUpdate**](TanzuApi.md#TanzuUpdate) | **Post** /api/v1/tanzu/update | Update tanzu credentials



## TanzuCreate

> ApiResponse TanzuCreate(ctx).CreateTanzuCommand(createTanzuCommand).Execute()

Create tanzu credentials

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
    createTanzuCommand := *openapiclient.NewCreateTanzuCommand("Name_example", "Username_example", "Url_example", "Password_example", "VolumeType_example", "Namespace_example") // CreateTanzuCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TanzuApi.TanzuCreate(context.Background()).CreateTanzuCommand(createTanzuCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TanzuApi.TanzuCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TanzuCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `TanzuApi.TanzuCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTanzuCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTanzuCommand** | [**CreateTanzuCommand**](CreateTanzuCommand.md) |  | 

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


## TanzuKubernetesVersions

> []string TanzuKubernetesVersions(ctx, cloudId).Execute()

Tanzu available k8s version list

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
    cloudId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TanzuApi.TanzuKubernetesVersions(context.Background(), cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TanzuApi.TanzuKubernetesVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TanzuKubernetesVersions`: []string
    fmt.Fprintf(os.Stdout, "Response from `TanzuApi.TanzuKubernetesVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTanzuKubernetesVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TanzuList

> TanzuCredentialsList TanzuList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of tanzu cloud credentials

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
    searchId := "searchId_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TanzuApi.TanzuList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TanzuApi.TanzuList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TanzuList`: TanzuCredentialsList
    fmt.Fprintf(os.Stdout, "Response from `TanzuApi.TanzuList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTanzuListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **searchId** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**TanzuCredentialsList**](TanzuCredentialsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TanzuStorageList

> []string TanzuStorageList(ctx).TanzuStorageListCommand(tanzuStorageListCommand).Execute()

Tanzu storage list

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
    tanzuStorageListCommand := *openapiclient.NewTanzuStorageListCommand() // TanzuStorageListCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TanzuApi.TanzuStorageList(context.Background()).TanzuStorageListCommand(tanzuStorageListCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TanzuApi.TanzuStorageList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TanzuStorageList`: []string
    fmt.Fprintf(os.Stdout, "Response from `TanzuApi.TanzuStorageList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTanzuStorageListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tanzuStorageListCommand** | [**TanzuStorageListCommand**](TanzuStorageListCommand.md) |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TanzuUpdate

> TanzuUpdate(ctx).UpdateTanzuCommand(updateTanzuCommand).Execute()

Update tanzu credentials

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
    updateTanzuCommand := *openapiclient.NewUpdateTanzuCommand(int32(123), "Name_example", "Password_example") // UpdateTanzuCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TanzuApi.TanzuUpdate(context.Background()).UpdateTanzuCommand(updateTanzuCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TanzuApi.TanzuUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTanzuUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTanzuCommand** | [**UpdateTanzuCommand**](UpdateTanzuCommand.md) |  | 

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

