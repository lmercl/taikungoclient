# \GoogleApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GooglecloudBillingAccountList**](GoogleApi.md#GooglecloudBillingAccountList) | **Post** /api/v1/googlecloud/billing-accounts | Retrieve google billing accounts list
[**GooglecloudCreate**](GoogleApi.md#GooglecloudCreate) | **Post** /api/v1/googlecloud/create | Create google cloud credential
[**GooglecloudList**](GoogleApi.md#GooglecloudList) | **Get** /api/v1/googlecloud/list | Retrieve list of google cloud credentials
[**GooglecloudRegionList**](GoogleApi.md#GooglecloudRegionList) | **Post** /api/v1/googlecloud/regions | Retrieve google region list
[**GooglecloudZoneList**](GoogleApi.md#GooglecloudZoneList) | **Post** /api/v1/googlecloud/zones | Google zones list



## GooglecloudBillingAccountList

> []CommonStringBasedDropdownDto GooglecloudBillingAccountList(ctx).Config(config).Execute()

Retrieve google billing accounts list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleApi.GooglecloudBillingAccountList(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.GooglecloudBillingAccountList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudBillingAccountList`: []CommonStringBasedDropdownDto
    fmt.Fprintf(os.Stdout, "Response from `GoogleApi.GooglecloudBillingAccountList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudBillingAccountListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

[**[]CommonStringBasedDropdownDto**](CommonStringBasedDropdownDto.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudCreate

> ApiResponse GooglecloudCreate(ctx).Name(name).Config(config).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).Execute()

Create google cloud credential

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
    name := "name_example" // string |  (optional)
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)
    importProject := true // bool |  (optional)
    folderId := "folderId_example" // string |  (optional)
    billingAccountId := "billingAccountId_example" // string |  (optional)
    azCount := int32(56) // int32 |  (optional)
    region := "region_example" // string |  (optional)
    organizationId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleApi.GooglecloudCreate(context.Background()).Name(name).Config(config).ImportProject(importProject).FolderId(folderId).BillingAccountId(billingAccountId).AzCount(azCount).Region(region).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.GooglecloudCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudCreate`: ApiResponse
    fmt.Fprintf(os.Stdout, "Response from `GoogleApi.GooglecloudCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 
 **config** | ***os.File** |  | 
 **importProject** | **bool** |  | 
 **folderId** | **string** |  | 
 **billingAccountId** | **string** |  | 
 **azCount** | **int32** |  | 
 **region** | **string** |  | 
 **organizationId** | **int32** |  | 

### Return type

[**ApiResponse**](ApiResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudList

> GoogleCredentialList GooglecloudList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()

Retrieve list of google cloud credentials

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
    resp, r, err := apiClient.GoogleApi.GooglecloudList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).SortBy(sortBy).SortDirection(sortDirection).Search(search).SearchId(searchId).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.GooglecloudList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudList`: GoogleCredentialList
    fmt.Fprintf(os.Stdout, "Response from `GoogleApi.GooglecloudList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudListRequest struct via the builder pattern


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

[**GoogleCredentialList**](GoogleCredentialList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudRegionList

> []string GooglecloudRegionList(ctx).Config(config).Execute()

Retrieve google region list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleApi.GooglecloudRegionList(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.GooglecloudRegionList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudRegionList`: []string
    fmt.Fprintf(os.Stdout, "Response from `GoogleApi.GooglecloudRegionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudRegionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 

### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GooglecloudZoneList

> AzResult GooglecloudZoneList(ctx).Config(config).Region(region).CloudId(cloudId).Execute()

Google zones list

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
    config := os.NewFile(1234, "some_file") // *os.File |  (optional)
    region := "region_example" // string |  (optional)
    cloudId := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GoogleApi.GooglecloudZoneList(context.Background()).Config(config).Region(region).CloudId(cloudId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GoogleApi.GooglecloudZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GooglecloudZoneList`: AzResult
    fmt.Fprintf(os.Stdout, "Response from `GoogleApi.GooglecloudZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGooglecloudZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | ***os.File** |  | 
 **region** | **string** |  | 
 **cloudId** | **int32** |  | 

### Return type

[**AzResult**](AzResult.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

