# \ShowbackRulesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowbackrulesCreate**](ShowbackRulesApi.md#ShowbackrulesCreate) | **Post** /showback/v1/showbackrules/create | Create showback rule
[**ShowbackrulesDelete**](ShowbackRulesApi.md#ShowbackrulesDelete) | **Delete** /showback/v1/showbackrules/{id} | Delete showback rule
[**ShowbackrulesDeleteAll**](ShowbackRulesApi.md#ShowbackrulesDeleteAll) | **Post** /showback/v1/showbackrules/delete | Delete multiple showback rule
[**ShowbackrulesList**](ShowbackRulesApi.md#ShowbackrulesList) | **Get** /showback/v1/showbackrules | Retrieve all showback rules
[**ShowbackrulesUpdate**](ShowbackRulesApi.md#ShowbackrulesUpdate) | **Post** /showback/v1/showbackrules/update | Create showback rule



## ShowbackrulesCreate

> ApiResponse ShowbackrulesCreate(ctx).CreateShowbackRuleCommand(createShowbackRuleCommand).Execute()

Create showback rule

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
    createShowbackRuleCommand := *openapiclient.NewCreateShowbackRuleCommand() // CreateShowbackRuleCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackRulesApi.ShowbackrulesCreate(context.Background()).CreateShowbackRuleCommand(createShowbackRuleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackRulesApi.ShowbackrulesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackrulesCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `ShowbackRulesApi.ShowbackrulesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackrulesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShowbackRuleCommand** | [**CreateShowbackRuleCommand**](CreateShowbackRuleCommand.md) |  | 

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


## ShowbackrulesDelete

> ShowbackrulesDelete(ctx, id).Execute()

Delete showback rule

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
    r, err := apiClient.ShowbackRulesApi.ShowbackrulesDelete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackRulesApi.ShowbackrulesDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiShowbackrulesDeleteRequest struct via the builder pattern


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


## ShowbackrulesDeleteAll

> ShowbackrulesDeleteAll(ctx).DeleteRulesCommand(deleteRulesCommand).Execute()

Delete multiple showback rule

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
    deleteRulesCommand := *openapiclient.NewDeleteRulesCommand() // DeleteRulesCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackRulesApi.ShowbackrulesDeleteAll(context.Background()).DeleteRulesCommand(deleteRulesCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackRulesApi.ShowbackrulesDeleteAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackrulesDeleteAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteRulesCommand** | [**DeleteRulesCommand**](DeleteRulesCommand.md) |  | 

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


## ShowbackrulesList

> ShowbackRuleList ShowbackrulesList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()

Retrieve all showback rules

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
    resp, r, err := apiClient.ShowbackRulesApi.ShowbackrulesList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackRulesApi.ShowbackrulesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackrulesList`: ShowbackRuleList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackRulesApi.ShowbackrulesList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackrulesListRequest struct via the builder pattern


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

[**ShowbackRuleList**](ShowbackRuleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackrulesUpdate

> ShowbackrulesUpdate(ctx).UpdateShowbackRuleCommand(updateShowbackRuleCommand).Execute()

Create showback rule

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
    updateShowbackRuleCommand := *openapiclient.NewUpdateShowbackRuleCommand() // UpdateShowbackRuleCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackRulesApi.ShowbackrulesUpdate(context.Background()).UpdateShowbackRuleCommand(updateShowbackRuleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackRulesApi.ShowbackrulesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbackrulesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateShowbackRuleCommand** | [**UpdateShowbackRuleCommand**](UpdateShowbackRuleCommand.md) |  | 

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

