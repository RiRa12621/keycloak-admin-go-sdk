# \WorkflowsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRealmsRealmWorkflowsGet**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsGet) | **Get** /admin/realms/{realm}/workflows | List workflows
[**AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost) | **Post** /admin/realms/{realm}/workflows/{id}/activate/{type}/{resourceId} | Activate workflow for resource
[**AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost) | **Post** /admin/realms/{realm}/workflows/{id}/deactivate/{type}/{resourceId} | Deactivate workflow for resource
[**AdminRealmsRealmWorkflowsIdDelete**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsIdDelete) | **Delete** /admin/realms/{realm}/workflows/{id} | Delete workflow
[**AdminRealmsRealmWorkflowsIdGet**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsIdGet) | **Get** /admin/realms/{realm}/workflows/{id} | Get workflow
[**AdminRealmsRealmWorkflowsIdPut**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsIdPut) | **Put** /admin/realms/{realm}/workflows/{id} | Update workflow
[**AdminRealmsRealmWorkflowsPost**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsPost) | **Post** /admin/realms/{realm}/workflows | Create workflow
[**AdminRealmsRealmWorkflowsScheduledResourceIdGet**](WorkflowsAPI.md#AdminRealmsRealmWorkflowsScheduledResourceIdGet) | **Get** /admin/realms/{realm}/workflows/scheduled/{resource-id} | List scheduled workflows for resource



## AdminRealmsRealmWorkflowsGet

> WorkflowRepresentation AdminRealmsRealmWorkflowsGet(ctx, realm).Exact(exact).First(first).Max(max).Search(search).Execute()

List workflows



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	exact := true // bool | Boolean which defines whether the param 'search' must match exactly or not (optional)
	first := int32(56) // int32 | The position of the first result to be processed (pagination offset) (optional)
	max := int32(56) // int32 | The maximum number of results to be returned - defaults to 10 (optional)
	search := "search_example" // string | A String representing the workflow name - either partial or exact (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsGet(context.Background(), realm).Exact(exact).First(first).Max(max).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmWorkflowsGet`: WorkflowRepresentation
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AdminRealmsRealmWorkflowsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exact** | **bool** | Boolean which defines whether the param &#39;search&#39; must match exactly or not | 
 **first** | **int32** | The position of the first result to be processed (pagination offset) | 
 **max** | **int32** | The maximum number of results to be returned - defaults to 10 | 
 **search** | **string** | A String representing the workflow name - either partial or exact | 

### Return type

[**WorkflowRepresentation**](WorkflowRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost

> AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost(ctx, realm, id, resourceId, type_).NotBefore(notBefore).Execute()

Activate workflow for resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | Workflow identifier
	resourceId := "resourceId_example" // string | Resource identifier
	type_ := map[string]interface{}{ ... } // map[string]interface{} | Resource type
	notBefore := "notBefore_example" // string | Optional value representing the time to schedule the first workflow step. The value is either an integer representing the seconds from now, an integer followed by 'ms' representing milliseconds from now, or an ISO-8601 date string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost(context.Background(), realm, id, resourceId, type_).NotBefore(notBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsIdActivateTypeResourceIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Workflow identifier | 
**resourceId** | **string** | Resource identifier | 
**type_** | [**map[string]interface{}**](.md) | Resource type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsIdActivateTypeResourceIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **notBefore** | **string** | Optional value representing the time to schedule the first workflow step. The value is either an integer representing the seconds from now, an integer followed by &#39;ms&#39; representing milliseconds from now, or an ISO-8601 date string. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost

> AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost(ctx, realm, id, resourceId, type_).Execute()

Deactivate workflow for resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | Workflow identifier
	resourceId := "resourceId_example" // string | Resource identifier
	type_ := map[string]interface{}{ ... } // map[string]interface{} | Resource type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost(context.Background(), realm, id, resourceId, type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Workflow identifier | 
**resourceId** | **string** | Resource identifier | 
**type_** | [**map[string]interface{}**](.md) | Resource type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsIdDeactivateTypeResourceIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsIdDelete

> AdminRealmsRealmWorkflowsIdDelete(ctx, realm, id).Execute()

Delete workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | Workflow identifier

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsIdDelete(context.Background(), realm, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Workflow identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsIdGet

> WorkflowRepresentation AdminRealmsRealmWorkflowsIdGet(ctx, realm, id).IncludeId(includeId).Execute()

Get workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | Workflow identifier
	includeId := true // bool | Indicates whether the workflow id should be included in the representation or not - defaults to true (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsIdGet(context.Background(), realm, id).IncludeId(includeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmWorkflowsIdGet`: WorkflowRepresentation
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AdminRealmsRealmWorkflowsIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Workflow identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeId** | **bool** | Indicates whether the workflow id should be included in the representation or not - defaults to true | 

### Return type

[**WorkflowRepresentation**](WorkflowRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsIdPut

> AdminRealmsRealmWorkflowsIdPut(ctx, realm, id).WorkflowRepresentation(workflowRepresentation).Execute()

Update workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	id := "id_example" // string | Workflow identifier
	workflowRepresentation := *openapiclient.NewWorkflowRepresentation() // WorkflowRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsIdPut(context.Background(), realm, id).WorkflowRepresentation(workflowRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**id** | **string** | Workflow identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **workflowRepresentation** | [**WorkflowRepresentation**](WorkflowRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/yaml, application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsPost

> AdminRealmsRealmWorkflowsPost(ctx, realm).WorkflowRepresentation(workflowRepresentation).Execute()

Create workflow



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	workflowRepresentation := *openapiclient.NewWorkflowRepresentation() // WorkflowRepresentation |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsPost(context.Background(), realm).WorkflowRepresentation(workflowRepresentation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowRepresentation** | [**WorkflowRepresentation**](WorkflowRepresentation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/yaml, application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminRealmsRealmWorkflowsScheduledResourceIdGet

> WorkflowRepresentation AdminRealmsRealmWorkflowsScheduledResourceIdGet(ctx, realm, resourceId).Execute()

List scheduled workflows for resource



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/rira12621/keycloak-admin-go-sdk"
)

func main() {
	realm := "realm_example" // string | realm name (not id!)
	resourceId := "resourceId_example" // string | Identifier of the resource associated with the scheduled workflows

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.AdminRealmsRealmWorkflowsScheduledResourceIdGet(context.Background(), realm, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.AdminRealmsRealmWorkflowsScheduledResourceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminRealmsRealmWorkflowsScheduledResourceIdGet`: WorkflowRepresentation
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.AdminRealmsRealmWorkflowsScheduledResourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**realm** | **string** | realm name (not id!) | 
**resourceId** | **string** | Identifier of the resource associated with the scheduled workflows | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminRealmsRealmWorkflowsScheduledResourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WorkflowRepresentation**](WorkflowRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

