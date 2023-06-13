# \SshUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SshusersCreate**](SshUsersApi.md#SshusersCreate) | **Post** /api/v1/sshusers/create | Create access profile ssh user
[**SshusersDelete**](SshUsersApi.md#SshusersDelete) | **Post** /api/v1/sshusers/delete | Delete access profile ssh user
[**SshusersEdit**](SshUsersApi.md#SshusersEdit) | **Post** /api/v1/sshusers/edit | Edit access profile ssh user
[**SshusersList**](SshUsersApi.md#SshusersList) | **Get** /api/v1/sshusers/list/{accessProfileId} | List ssh user by access profile id



## SshusersCreate

> ApiResponse SshusersCreate(ctx).CreateSshUserCommand(createSshUserCommand).Execute()

Create access profile ssh user

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
    createSshUserCommand := *openapiclient.NewCreateSshUserCommand() // CreateSshUserCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshUsersApi.SshusersCreate(context.Background()).CreateSshUserCommand(createSshUserCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshUsersApi.SshusersCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshusersCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `SshUsersApi.SshusersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshusersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSshUserCommand** | [**CreateSshUserCommand**](CreateSshUserCommand.md) |  | 

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


## SshusersDelete

> SshusersDelete(ctx).DeleteSshUserCommand(deleteSshUserCommand).Execute()

Delete access profile ssh user

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
    deleteSshUserCommand := *openapiclient.NewDeleteSshUserCommand() // DeleteSshUserCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SshUsersApi.SshusersDelete(context.Background()).DeleteSshUserCommand(deleteSshUserCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshUsersApi.SshusersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshusersDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteSshUserCommand** | [**DeleteSshUserCommand**](DeleteSshUserCommand.md) |  | 

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


## SshusersEdit

> SshusersEdit(ctx).EditSshUserCommand(editSshUserCommand).Execute()

Edit access profile ssh user

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
    editSshUserCommand := *openapiclient.NewEditSshUserCommand() // EditSshUserCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SshUsersApi.SshusersEdit(context.Background()).EditSshUserCommand(editSshUserCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshUsersApi.SshusersEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSshusersEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editSshUserCommand** | [**EditSshUserCommand**](EditSshUserCommand.md) |  | 

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


## SshusersList

> []SshUsersListDto SshusersList(ctx, accessProfileId).Search(search).Execute()

List ssh user by access profile id

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
    accessProfileId := int32(56) // int32 | 
    search := "search_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SshUsersApi.SshusersList(context.Background(), accessProfileId).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SshUsersApi.SshusersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SshusersList`: []SshUsersListDto
    fmt.Fprintf(os.Stdout, "Response from `SshUsersApi.SshusersList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessProfileId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSshusersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **search** | **string** |  | 

### Return type

[**[]SshUsersListDto**](SshUsersListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

