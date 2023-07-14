# \OpenshiftApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenshiftStorageClass**](OpenshiftApi.md#OpenshiftStorageClass) | **Post** /api/v1/openshift/storage-class | 



## OpenshiftStorageClass

> []string OpenshiftStorageClass(ctx).Config(config).Execute()



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
    resp, r, err := apiClient.OpenshiftApi.OpenshiftStorageClass(context.Background()).Config(config).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenshiftApi.OpenshiftStorageClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OpenshiftStorageClass`: []string
    fmt.Fprintf(os.Stdout, "Response from `OpenshiftApi.OpenshiftStorageClass`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOpenshiftStorageClassRequest struct via the builder pattern


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

