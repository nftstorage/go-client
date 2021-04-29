# \NFTStorageAPI

All URIs are relative to *https://api.nft.storage*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](NFTStorageAPI.md#Delete) | **Delete** /{cid} | Stop storing the content with the passed CID
[**List**](NFTStorageAPI.md#List) | **Get** / | List all stored files
[**Status**](NFTStorageAPI.md#Status) | **Get** /{cid} | Get information for the stored file CID
[**Store**](NFTStorageAPI.md#Store) | **Post** /upload | Store a file



## Delete

> DeleteResponse Delete(ctx, cid).Execute()

Stop storing the content with the passed CID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cid := "cid_example" // string | CID for the NFT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NFTStorageAPI.Delete(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: DeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | CID for the NFT | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteResponse**](DeleteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListResponse List(ctx).Before(before).Limit(limit).Execute()

List all stored files

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    before := time.Now() // time.Time | Return results created before provided timestamp (optional)
    limit := int32(56) // int32 | Max records to return (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NFTStorageAPI.List(context.Background()).Before(before).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **before** | **time.Time** | Return results created before provided timestamp | 
 **limit** | **int32** | Max records to return | [default to 10]

### Return type

[**ListResponse**](ListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Status

> GetResponse Status(ctx, cid).Execute()

Get information for the stored file CID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cid := "cid_example" // string | CID for the NFT

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NFTStorageAPI.Status(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.Status``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Status`: GetResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.Status`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **string** | CID for the NFT | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetResponse**](GetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Store

> UploadResponse Store(ctx).Body(body).Execute()

Store a file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := os.NewFile(1234, "some_file") // *os.File | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NFTStorageAPI.Store(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTStorageAPI.Store``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Store`: UploadResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTStorageAPI.Store`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

[**UploadResponse**](UploadResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: image/png, application/octet-stream, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

