# \TicketApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketArchive**](TicketApi.md#TicketArchive) | **Post** /api/v1/ticket/archive | Archive ticket
[**TicketClose**](TicketApi.md#TicketClose) | **Post** /api/v1/ticket/close | Close ticket
[**TicketCreate**](TicketApi.md#TicketCreate) | **Post** /api/v1/ticket/create | Create ticket
[**TicketDelete**](TicketApi.md#TicketDelete) | **Delete** /api/v1/ticket/delete/{ticketId} | 
[**TicketDeleteMessage**](TicketApi.md#TicketDeleteMessage) | **Delete** /api/v1/ticket/delete/message/{messageId} | 
[**TicketEdit**](TicketApi.md#TicketEdit) | **Post** /api/v1/ticket/edit | Edit ticket
[**TicketEditMessage**](TicketApi.md#TicketEditMessage) | **Post** /api/v1/ticket/edit/message | Edit ticket message
[**TicketList**](TicketApi.md#TicketList) | **Post** /api/v1/ticket/list | Retrieve list of tickets
[**TicketMessages**](TicketApi.md#TicketMessages) | **Post** /api/v1/ticket/{ticketId}/messages | 
[**TicketOpen**](TicketApi.md#TicketOpen) | **Post** /api/v1/ticket/open | Open ticket
[**TicketReply**](TicketApi.md#TicketReply) | **Post** /api/v1/ticket/reply | Reply message
[**TicketSetPriority**](TicketApi.md#TicketSetPriority) | **Post** /api/v1/ticket/set-priority | Set priority
[**TicketTransfer**](TicketApi.md#TicketTransfer) | **Post** /api/v1/ticket/transfer | Transfer ticket
[**TicketTransferList**](TicketApi.md#TicketTransferList) | **Post** /api/v1/ticket/transfer/list | Retrieve organization managers



## TicketArchive

> TicketArchive(ctx).ArchiveTicketCommand(archiveTicketCommand).Execute()

Archive ticket

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
    archiveTicketCommand := *openapiclient.NewArchiveTicketCommand() // ArchiveTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketArchive(context.Background()).ArchiveTicketCommand(archiveTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveTicketCommand** | [**ArchiveTicketCommand**](ArchiveTicketCommand.md) |  | 

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


## TicketClose

> TicketClose(ctx).CloseTicketCommand(closeTicketCommand).Execute()

Close ticket

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
    closeTicketCommand := *openapiclient.NewCloseTicketCommand() // CloseTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketClose(context.Background()).CloseTicketCommand(closeTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **closeTicketCommand** | [**CloseTicketCommand**](CloseTicketCommand.md) |  | 

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


## TicketCreate

> TicketCreate(ctx).CreateTicketCommand(createTicketCommand).Execute()

Create ticket

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
    createTicketCommand := *openapiclient.NewCreateTicketCommand() // CreateTicketCommand |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketCreate(context.Background()).CreateTicketCommand(createTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTicketCommand** | [**CreateTicketCommand**](CreateTicketCommand.md) |  | 

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


## TicketDelete

> TicketDelete(ctx, ticketId).Execute()



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
    ticketId := "ticketId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketDelete(context.Background(), ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteRequest struct via the builder pattern


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


## TicketDeleteMessage

> TicketDeleteMessage(ctx, messageId).Execute()



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
    messageId := "messageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketDeleteMessage(context.Background(), messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketDeleteMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketDeleteMessageRequest struct via the builder pattern


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


## TicketEdit

> TicketEdit(ctx).EditTicketCommand(editTicketCommand).Execute()

Edit ticket

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
    editTicketCommand := *openapiclient.NewEditTicketCommand() // EditTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketEdit(context.Background()).EditTicketCommand(editTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editTicketCommand** | [**EditTicketCommand**](EditTicketCommand.md) |  | 

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


## TicketEditMessage

> TicketEditMessage(ctx).EditArticleCommand(editArticleCommand).Execute()

Edit ticket message

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
    editArticleCommand := *openapiclient.NewEditArticleCommand() // EditArticleCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketEditMessage(context.Background()).EditArticleCommand(editArticleCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketEditMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketEditMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **editArticleCommand** | [**EditArticleCommand**](EditArticleCommand.md) |  | 

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


## TicketList

> AllTicketsList TicketList(ctx).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()

Retrieve list of tickets

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
    search := "search_example" // string |  (optional)
    startDate := "startDate_example" // string |  (optional)
    endDate := "endDate_example" // string |  (optional)
    ticketId := "ticketId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.TicketList(context.Background()).Limit(limit).Offset(offset).OrganizationId(organizationId).Search(search).StartDate(startDate).EndDate(endDate).TicketId(ticketId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketList`: AllTicketsList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 
 **organizationId** | **int32** |  | 
 **search** | **string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **ticketId** | **string** |  | 

### Return type

[**AllTicketsList**](AllTicketsList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketMessages

> ArticleList TicketMessages(ctx, ticketId).Limit(limit).Offset(offset).Execute()



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
    ticketId := "ticketId_example" // string | 
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TicketApi.TicketMessages(context.Background(), ticketId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketMessages`: ArticleList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ticketId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**ArticleList**](ArticleList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketOpen

> TicketOpen(ctx).OpenTicketCommand(openTicketCommand).Execute()

Open ticket

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
    openTicketCommand := *openapiclient.NewOpenTicketCommand() // OpenTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketOpen(context.Background()).OpenTicketCommand(openTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **openTicketCommand** | [**OpenTicketCommand**](OpenTicketCommand.md) |  | 

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


## TicketReply

> TicketReply(ctx).ReplyTicketCommand(replyTicketCommand).Execute()

Reply message

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
    replyTicketCommand := *openapiclient.NewReplyTicketCommand() // ReplyTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketReply(context.Background()).ReplyTicketCommand(replyTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **replyTicketCommand** | [**ReplyTicketCommand**](ReplyTicketCommand.md) |  | 

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


## TicketSetPriority

> TicketSetPriority(ctx).SetTicketPriorityCommand(setTicketPriorityCommand).Execute()

Set priority

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
    setTicketPriorityCommand := *openapiclient.NewSetTicketPriorityCommand() // SetTicketPriorityCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketSetPriority(context.Background()).SetTicketPriorityCommand(setTicketPriorityCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketSetPriority``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketSetPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setTicketPriorityCommand** | [**SetTicketPriorityCommand**](SetTicketPriorityCommand.md) |  | 

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


## TicketTransfer

> TicketTransfer(ctx).TransferTicketCommand(transferTicketCommand).Execute()

Transfer ticket

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
    transferTicketCommand := *openapiclient.NewTransferTicketCommand() // TransferTicketCommand | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TicketApi.TicketTransfer(context.Background()).TransferTicketCommand(transferTicketCommand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferTicketCommand** | [**TransferTicketCommand**](TransferTicketCommand.md) |  | 

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


## TicketTransferList

> []TransferList TicketTransferList(ctx).OrganizationId(organizationId).Execute()

Retrieve organization managers

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
    resp, r, err := apiClient.TicketApi.TicketTransferList(context.Background()).OrganizationId(organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketApi.TicketTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketTransferList`: []TransferList
    fmt.Fprintf(os.Stdout, "Response from `TicketApi.TicketTransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationId** | **int32** |  | 

### Return type

[**[]TransferList**](TransferList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

