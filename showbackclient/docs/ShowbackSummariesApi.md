# \ShowbackSummariesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ShowbackSummariesCreate**](ShowbackSummariesApi.md#ShowbackSummariesCreate) | **Post** /showback/v{v}/ShowbackSummaries/create | Create showback summary
[**ShowbackSummariesExportCsv**](ShowbackSummariesApi.md#ShowbackSummariesExportCsv) | **Get** /showback/v{v}/ShowbackSummaries/export | Export Csv
[**ShowbackSummariesGroupedByLabelList**](ShowbackSummariesApi.md#ShowbackSummariesGroupedByLabelList) | **Get** /showback/v{v}/ShowbackSummaries/grouped/byLabel | Retrieve a grouped list of by label showback summaries
[**ShowbackSummariesGroupedByProjectList**](ShowbackSummariesApi.md#ShowbackSummariesGroupedByProjectList) | **Get** /showback/v{v}/ShowbackSummaries/grouped/byProject | Retrieve a grouped list of by project showback summaries
[**ShowbackSummariesGroupedList**](ShowbackSummariesApi.md#ShowbackSummariesGroupedList) | **Get** /showback/v{v}/ShowbackSummaries/grouped | Retrieve a grouped list of showback summaries
[**ShowbackSummariesGroupedShowbackSummaryList**](ShowbackSummariesApi.md#ShowbackSummariesGroupedShowbackSummaryList) | **Get** /showback/v{v}/ShowbackSummaries/grouped/list | Retrieve grouped showback summary



## ShowbackSummariesCreate

> ShowbackSummariesCreate(ctx, v).Body(body).Execute()

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
    v := "v_example" // string | 
    body := *openapiclient.NewCreateShowbackSummaryCommand() // CreateShowbackSummaryCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesCreate(context.Background(), v).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateShowbackSummaryCommand**](CreateShowbackSummaryCommand.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/json, text/json, application/*+json
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackSummariesExportCsv

> ShowbackSummariesExportCsv(ctx, v).OrganizationId(organizationId).IsEmailEnabled(isEmailEnabled).Execute()

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
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)
    isEmailEnabled := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesExportCsv(context.Background(), v).OrganizationId(organizationId).IsEmailEnabled(isEmailEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesExportCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesExportCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **isEmailEnabled** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackSummariesGroupedByLabelList

> GroupedShowbackByLabelList ShowbackSummariesGroupedByLabelList(ctx, v).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Retrieve a grouped list of by label showback summaries

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)
    isDeleted := true // bool |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesGroupedByLabelList(context.Background(), v).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesGroupedByLabelList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackSummariesGroupedByLabelList`: GroupedShowbackByLabelList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbackSummariesGroupedByLabelList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesGroupedByLabelListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackByLabelList**](GroupedShowbackByLabelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackSummariesGroupedByProjectList

> GroupedShowbackByProjectList ShowbackSummariesGroupedByProjectList(ctx, v).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Retrieve a grouped list of by project showback summaries

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)
    isDeleted := true // bool |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesGroupedByProjectList(context.Background(), v).OrganizationId(organizationId).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesGroupedByProjectList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackSummariesGroupedByProjectList`: GroupedShowbackByProjectList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbackSummariesGroupedByProjectList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesGroupedByProjectListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackByProjectList**](GroupedShowbackByProjectList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackSummariesGroupedList

> GroupedShowbackList ShowbackSummariesGroupedList(ctx, v).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()

Retrieve a grouped list of showback summaries

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)
    periodDuration := "periodDuration_example" // string |  (optional)
    isDeleted := true // bool |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesGroupedList(context.Background(), v).OrganizationId(organizationId).PeriodDuration(periodDuration).IsDeleted(isDeleted).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesGroupedList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackSummariesGroupedList`: GroupedShowbackList
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbackSummariesGroupedList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesGroupedListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **periodDuration** | **string** |  | 
 **isDeleted** | **bool** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**GroupedShowbackList**](GroupedShowbackList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShowbackSummariesGroupedShowbackSummaryList

> []GroupedShowbackSummaryListDto ShowbackSummariesGroupedShowbackSummaryList(ctx, v).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()

Retrieve grouped showback summary

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/chnyda/taikungoclient"
)

func main() {
    v := "v_example" // string | 
    organizationId := int32(56) // int32 |  (optional)
    fromDate := time.Now() // time.Time |  (optional)
    toDate := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShowbackSummariesApi.ShowbackSummariesGroupedShowbackSummaryList(context.Background(), v).OrganizationId(organizationId).FromDate(fromDate).ToDate(toDate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShowbackSummariesApi.ShowbackSummariesGroupedShowbackSummaryList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShowbackSummariesGroupedShowbackSummaryList`: []GroupedShowbackSummaryListDto
    fmt.Fprintf(os.Stdout, "Response from `ShowbackSummariesApi.ShowbackSummariesGroupedShowbackSummaryList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShowbackSummariesGroupedShowbackSummaryListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationId** | **int32** |  | 
 **fromDate** | **time.Time** |  | 
 **toDate** | **time.Time** |  | 

### Return type

[**[]GroupedShowbackSummaryListDto**](GroupedShowbackSummaryListDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

