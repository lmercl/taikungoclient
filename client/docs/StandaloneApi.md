# \StandaloneApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandaloneCommit**](StandaloneApi.md#StandaloneCommit) | **Post** /api/v1/standalone/commit | Commit vm
[**StandaloneCreate**](StandaloneApi.md#StandaloneCreate) | **Post** /api/v1/standalone/create | Create a new vm in the given project.
[**StandaloneDelete**](StandaloneApi.md#StandaloneDelete) | **Post** /api/v1/standalone/delete | Delete vm
[**StandaloneDetails**](StandaloneApi.md#StandaloneDetails) | **Get** /api/v1/standalone/{projectId} | Retrieve a list of standalone vm with detailed info
[**StandaloneForPoller**](StandaloneApi.md#StandaloneForPoller) | **Get** /api/v1/standalone/forpoller | List all StandaloneVms for poller
[**StandaloneIpManagement**](StandaloneApi.md#StandaloneIpManagement) | **Post** /api/v1/standalone/ip/management | Enable/Disable stand alone public ip
[**StandaloneList**](StandaloneApi.md#StandaloneList) | **Get** /api/v1/standalone | Retrieve all vms
[**StandaloneProjectDetails**](StandaloneApi.md#StandaloneProjectDetails) | **Get** /api/v1/standalone/project/{projectId} | Retrieve details of the project by Id
[**StandalonePurge**](StandaloneApi.md#StandalonePurge) | **Post** /api/v1/standalone/purge | Purge vm
[**StandaloneRepair**](StandaloneApi.md#StandaloneRepair) | **Post** /api/v1/standalone/repair | Repair vm
[**StandaloneReset**](StandaloneApi.md#StandaloneReset) | **Post** /api/v1/standalone/reset | Reset vm status
[**StandaloneUpdate**](StandaloneApi.md#StandaloneUpdate) | **Post** /api/v1/standalone/update | Update vm
[**StandaloneUpdateFlavor**](StandaloneApi.md#StandaloneUpdateFlavor) | **Post** /api/v1/standalone/update/flavor | Update standalone vm flavor



## StandaloneCommit

> StandaloneCommit(ctx).CommitStandAloneVmCommand(commitStandAloneVmCommand).Execute()

Commit vm

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
    commitStandAloneVmCommand := *openapiclient.NewCommitStandAloneVmCommand() // CommitStandAloneVmCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneCommit(context.Background()).CommitStandAloneVmCommand(commitStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commitStandAloneVmCommand** | [**CommitStandAloneVmCommand**](CommitStandAloneVmCommand.md) |  | 

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


## StandaloneCreate

> ApiResponse StandaloneCreate(ctx).CreateStandAloneVmCommand(createStandAloneVmCommand).Execute()

Create a new vm in the given project.

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
    createStandAloneVmCommand := *openapiclient.NewCreateStandAloneVmCommand("Name_example", "FlavorName_example", "Image_example", int32(123), int32(123)) // CreateStandAloneVmCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneApi.StandaloneCreate(context.Background()).CreateStandAloneVmCommand(createStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `StandaloneApi.StandaloneCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandAloneVmCommand** | [**CreateStandAloneVmCommand**](CreateStandAloneVmCommand.md) |  | 

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


## StandaloneDelete

> StandaloneDelete(ctx).DeleteStandAloneVmCommand(deleteStandAloneVmCommand).Execute()

Delete vm

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
    deleteStandAloneVmCommand := *openapiclient.NewDeleteStandAloneVmCommand() // DeleteStandAloneVmCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneDelete(context.Background()).DeleteStandAloneVmCommand(deleteStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteStandAloneVmCommand** | [**DeleteStandAloneVmCommand**](DeleteStandAloneVmCommand.md) |  | 

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


## StandaloneDetails

> StandAloneVmListForDetails StandaloneDetails(ctx, projectId).SortBy(sortBy).SortDirection(sortDirection).Id(id).Execute()

Retrieve a list of standalone vm with detailed info

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
    projectId := int32(56) // int32 | 
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    id := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneApi.StandaloneDetails(context.Background(), projectId).SortBy(sortBy).SortDirection(sortDirection).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneDetails`: StandAloneVmListForDetails
    fmt.Fprintf(os.Stdout, "Response from `StandaloneApi.StandaloneDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **id** | **int32** |  | 

### Return type

[**StandAloneVmListForDetails**](StandAloneVmListForDetails.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneForPoller

> StandaloneVmsListForPoller StandaloneForPoller(ctx).Limit(limit).Offset(offset).ProjectId(projectId).Execute()

List all StandaloneVms for poller

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
    projectId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneApi.StandaloneForPoller(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneForPoller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneForPoller`: StandaloneVmsListForPoller
    fmt.Fprintf(os.Stdout, "Response from `StandaloneApi.StandaloneForPoller`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneForPollerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 

### Return type

[**StandaloneVmsListForPoller**](StandaloneVmsListForPoller.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneIpManagement

> StandaloneIpManagement(ctx).StandAloneVmIpManagementCommand(standAloneVmIpManagementCommand).Execute()

Enable/Disable stand alone public ip

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
    standAloneVmIpManagementCommand := *openapiclient.NewStandAloneVmIpManagementCommand(int32(123), "Mode_example") // StandAloneVmIpManagementCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneIpManagement(context.Background()).StandAloneVmIpManagementCommand(standAloneVmIpManagementCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneIpManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneIpManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **standAloneVmIpManagementCommand** | [**StandAloneVmIpManagementCommand**](StandAloneVmIpManagementCommand.md) |  | 

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


## StandaloneList

> StandaloneVmsList StandaloneList(ctx).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartVolumeSize(startVolumeSize).EndVolumeSize(endVolumeSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).SearchId(searchId).FilterBy(filterBy).Execute()

Retrieve all vms

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
    projectId := int32(56) // int32 |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    sortDirection := "sortDirection_example" // string |  (optional)
    search := "search_example" // string |  (optional)
    startRam := "startRam_example" // string |  (optional)
    endRam := "endRam_example" // string |  (optional)
    startVolumeSize := int64(789) // int64 |  (optional)
    endVolumeSize := int64(789) // int64 |  (optional)
    startCpu := int32(56) // int32 |  (optional)
    endCpu := int32(56) // int32 |  (optional)
    organizationId := int32(56) // int32 |  (optional)
    id := int32(56) // int32 |  (optional)
    searchId := "searchId_example" // string |  (optional)
    filterBy := "filterBy_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneApi.StandaloneList(context.Background()).Limit(limit).Offset(offset).ProjectId(projectId).SortBy(sortBy).SortDirection(sortDirection).Search(search).StartRam(startRam).EndRam(endRam).StartVolumeSize(startVolumeSize).EndVolumeSize(endVolumeSize).StartCpu(startCpu).EndCpu(endCpu).OrganizationId(organizationId).Id(id).SearchId(searchId).FilterBy(filterBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneList`: StandaloneVmsList
    fmt.Fprintf(os.Stdout, "Response from `StandaloneApi.StandaloneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **projectId** | **int32** |  | 
 **sortBy** | **string** |  | 
 **sortDirection** | **string** |  | 
 **search** | **string** |  | 
 **startRam** | **string** |  | 
 **endRam** | **string** |  | 
 **startVolumeSize** | **int64** |  | 
 **endVolumeSize** | **int64** |  | 
 **startCpu** | **int32** |  | 
 **endCpu** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **id** | **int32** |  | 
 **searchId** | **string** |  | 
 **filterBy** | **string** |  | 

### Return type

[**StandaloneVmsList**](StandaloneVmsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandaloneProjectDetails

> ProjectFullListDto StandaloneProjectDetails(ctx, projectId).Execute()

Retrieve details of the project by Id

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
    projectId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StandaloneApi.StandaloneProjectDetails(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneProjectDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StandaloneProjectDetails`: ProjectFullListDto
    fmt.Fprintf(os.Stdout, "Response from `StandaloneApi.StandaloneProjectDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneProjectDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectFullListDto**](ProjectFullListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandalonePurge

> StandalonePurge(ctx).PurgeStandAloneCommand(purgeStandAloneCommand).Execute()

Purge vm

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
    purgeStandAloneCommand := *openapiclient.NewPurgeStandAloneCommand() // PurgeStandAloneCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandalonePurge(context.Background()).PurgeStandAloneCommand(purgeStandAloneCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandalonePurge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandalonePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **purgeStandAloneCommand** | [**PurgeStandAloneCommand**](PurgeStandAloneCommand.md) |  | 

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


## StandaloneRepair

> StandaloneRepair(ctx).RepairStandAloneVmCommand(repairStandAloneVmCommand).Execute()

Repair vm

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
    repairStandAloneVmCommand := *openapiclient.NewRepairStandAloneVmCommand() // RepairStandAloneVmCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneRepair(context.Background()).RepairStandAloneVmCommand(repairStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneRepair``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneRepairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repairStandAloneVmCommand** | [**RepairStandAloneVmCommand**](RepairStandAloneVmCommand.md) |  | 

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


## StandaloneReset

> StandaloneReset(ctx).ResetStandAloneVmStatusCommand(resetStandAloneVmStatusCommand).Execute()

Reset vm status

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
    resetStandAloneVmStatusCommand := *openapiclient.NewResetStandAloneVmStatusCommand(int32(123)) // ResetStandAloneVmStatusCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneReset(context.Background()).ResetStandAloneVmStatusCommand(resetStandAloneVmStatusCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resetStandAloneVmStatusCommand** | [**ResetStandAloneVmStatusCommand**](ResetStandAloneVmStatusCommand.md) |  | 

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


## StandaloneUpdate

> StandaloneUpdate(ctx).UpdateStandAloneVmCommand(updateStandAloneVmCommand).Execute()

Update vm

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
    updateStandAloneVmCommand := *openapiclient.NewUpdateStandAloneVmCommand(int32(123)) // UpdateStandAloneVmCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneUpdate(context.Background()).UpdateStandAloneVmCommand(updateStandAloneVmCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandAloneVmCommand** | [**UpdateStandAloneVmCommand**](UpdateStandAloneVmCommand.md) |  | 

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


## StandaloneUpdateFlavor

> StandaloneUpdateFlavor(ctx).UpdateStandAloneVmFlavorCommand(updateStandAloneVmFlavorCommand).Execute()

Update standalone vm flavor

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
    updateStandAloneVmFlavorCommand := *openapiclient.NewUpdateStandAloneVmFlavorCommand(int32(123), "Flavor_example") // UpdateStandAloneVmFlavorCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StandaloneApi.StandaloneUpdateFlavor(context.Background()).UpdateStandAloneVmFlavorCommand(updateStandAloneVmFlavorCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandaloneApi.StandaloneUpdateFlavor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandaloneUpdateFlavorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateStandAloneVmFlavorCommand** | [**UpdateStandAloneVmFlavorCommand**](UpdateStandAloneVmFlavorCommand.md) |  | 

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

