# \ShowbackSummariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowbacksummariesByLabel**](ShowbackSummariesApi.md#ShowbacksummariesByLabel) | **Get** /showback/v1/showbacksummaries/grouped/byLabel | Grouped showback summary by label
[**ShowbacksummariesByProject**](ShowbackSummariesApi.md#ShowbacksummariesByProject) | **Get** /showback/v1/showbacksummaries/grouped/byProject | Grouped showback summary by project
[**ShowbacksummariesCreate**](ShowbackSummariesApi.md#ShowbacksummariesCreate) | **Post** /showback/v1/showbacksummaries/create | Create showback summary
[**ShowbacksummariesExportCsv**](ShowbackSummariesApi.md#ShowbacksummariesExportCsv) | **Get** /showback/v1/showbacksummaries/export | Export Csv
[**ShowbacksummariesGrouped**](ShowbackSummariesApi.md#ShowbacksummariesGrouped) | **Get** /showback/v1/showbacksummaries/grouped/list | Grouped showback summary
[**ShowbacksummariesGroupedList**](ShowbackSummariesApi.md#ShowbacksummariesGroupedList) | **Get** /showback/v1/showbacksummaries/grouped | Grouped list of showback summary



## ShowbacksummariesByLabel

> GroupedShowbackByLabelList ShowbacksummariesByLabel(ctx).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Grouped showback summary by label

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
    isDeleted := true // bool |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesByLabel(context.Background()).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesByLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbacksummariesByLabel`: GroupedShowbackByLabelList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbacksummariesByLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesByLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**GroupedShowbackByLabelList**](GroupedShowbackByLabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesByProject

> GroupedShowbackByProjectList ShowbacksummariesByProject(ctx).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Grouped showback summary by project

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
    isDeleted := true // bool |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesByProject(context.Background()).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesByProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbacksummariesByProject`: GroupedShowbackByProjectList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbacksummariesByProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesByProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**GroupedShowbackByProjectList**](GroupedShowbackByProjectList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesCreate

> ShowbacksummariesCreate(ctx).CreateShowbackSummaryCommand(createShowbackSummaryCommand).Execute()

Create showback summary

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
    createShowbackSummaryCommand := *openapiclient.NewCreateShowbackSummaryCommand() // CreateShowbackSummaryCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesCreate(context.Background()).CreateShowbackSummaryCommand(createShowbackSummaryCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShowbackSummaryCommand** | [**CreateShowbackSummaryCommand**](CreateShowbackSummaryCommand.md) |  | 

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


## ShowbacksummariesExportCsv

> CsvExporter ShowbacksummariesExportCsv(ctx).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).Execute()

Export Csv

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
    isEmailEnabled := true // bool | 
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesExportCsv(context.Background()).IsEmailEnabled(isEmailEnabled).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesExportCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbacksummariesExportCsv`: CsvExporter
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbacksummariesExportCsv`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isEmailEnabled** | **bool** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**CsvExporter**](CsvExporter.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesGrouped

> []GroupedShowbackSummaryListDto ShowbacksummariesGrouped(ctx).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()

Grouped showback summary

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
    fromDate := "fromDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesGrouped(context.Background()).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesGrouped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbacksummariesGrouped`: []GroupedShowbackSummaryListDto
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbacksummariesGrouped`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesGroupedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**[]GroupedShowbackSummaryListDto**](GroupedShowbackSummaryListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbacksummariesGroupedList

> GroupedShowbackList ShowbacksummariesGroupedList(ctx).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Grouped list of showback summary

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
    periodDuration := "periodDuration_example" // string |  (optional)
    isDeleted := true // bool |  (optional)
    fromDate := "fromDate_example" // string |  (optional)
    toDate := "toDate_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbacksummariesGroupedList(context.Background()).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbacksummariesGroupedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbacksummariesGroupedList`: GroupedShowbackList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbacksummariesGroupedList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiShowbacksummariesGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 
 **periodDuration** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **string** |  | 
 **toDate** | **string** |  | 

### Return type

[**GroupedShowbackList**](GroupedShowbackList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

