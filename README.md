# Go API client for msxsdk

MSX SDK client.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.5
- Package version: 1.0.5
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./msxsdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditingGenericEventsApi* | [**CreateGenericEvent**](docs/AuditingGenericEventsApi.md#creategenericevent) | **Post** /auditing/api/v8/genericevents | Create Generic Event
*DeviceTemplatesApi* | [**CreateDeviceTemplate**](docs/DeviceTemplatesApi.md#createdevicetemplate) | **Post** /manage/api/v8/devicetemplates | Creates a device template.
*DeviceTemplatesApi* | [**CreateDeviceTemplateVersion**](docs/DeviceTemplatesApi.md#createdevicetemplateversion) | **Post** /manage/api/v8/devicetemplates/versions | Creates a new version of an existing device template.
*DeviceTemplatesApi* | [**DeleteDeviceTemplate**](docs/DeviceTemplatesApi.md#deletedevicetemplate) | **Delete** /manage/api/v8/devicetemplates/{id} | Deletes a device template.
*DeviceTemplatesApi* | [**GetDeviceTemplate**](docs/DeviceTemplatesApi.md#getdevicetemplate) | **Get** /manage/api/v8/devicetemplates/{id} | Returns a device template.
*DeviceTemplatesApi* | [**GetDeviceTemplatesList**](docs/DeviceTemplatesApi.md#getdevicetemplateslist) | **Get** /manage/api/v8/devicetemplates/list | Returns a list of device templates.
*DeviceTemplatesApi* | [**ScanDeviceTemplateParameters**](docs/DeviceTemplatesApi.md#scandevicetemplateparameters) | **Post** /manage/api/v8/devicetemplates/parameters/scan | API to scan parameters from the device template XML.
*DeviceTemplatesApi* | [**UpdateDeviceTemplateAccess**](docs/DeviceTemplatesApi.md#updatedevicetemplateaccess) | **Put** /manage/api/v8/devicetemplates/{id}/access | Updates device template access.
*DevicesApi* | [**AttachDeviceTemplates**](docs/DevicesApi.md#attachdevicetemplates) | **Post** /manage/api/v8/devices/{id}/templates | Attaches one or more device templates to a device instance.
*DevicesApi* | [**CreateDevice**](docs/DevicesApi.md#createdevice) | **Post** /manage/api/v8/devices | Creates a device.
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /manage/api/v8/devices/{id} | Deletes a device.
*DevicesApi* | [**DetachDeviceTemplate**](docs/DevicesApi.md#detachdevicetemplate) | **Delete** /manage/api/v8/devices/{id}/templates/{templateId} | Detaches a template from a device.
*DevicesApi* | [**DetachDeviceTemplates**](docs/DevicesApi.md#detachdevicetemplates) | **Delete** /manage/api/v8/devices/{id}/templates | Detach device templates that are already attached to a device.
*DevicesApi* | [**GetDevice**](docs/DevicesApi.md#getdevice) | **Get** /manage/api/v8/devices/{id} | Returns a device.
*DevicesApi* | [**GetDeviceConfig**](docs/DevicesApi.md#getdeviceconfig) | **Get** /manage/api/v8/devices/{id}/config | Returns the running configuration for a device.
*DevicesApi* | [**GetDeviceTemplateHistory**](docs/DevicesApi.md#getdevicetemplatehistory) | **Get** /manage/api/v8/devices/{id}/templates | Returns device template history.
*DevicesApi* | [**GetDevicesPage**](docs/DevicesApi.md#getdevicespage) | **Get** /manage/api/v8/devices | Returns a page of devices.
*DevicesApi* | [**PatchDevice**](docs/DevicesApi.md#patchdevice) | **Patch** /manage/api/v8/devices/{id} | Update a device.
*DevicesApi* | [**RedeployDevice**](docs/DevicesApi.md#redeploydevice) | **Post** /manage/api/v8/devices/{id}/redeploy | Dedeploys a device.
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Put** /manage/api/v8/devices/{id} | Update a device.
*DevicesApi* | [**UpdateDeviceTemplates**](docs/DevicesApi.md#updatedevicetemplates) | **Put** /manage/api/v8/devices/{id}/templates | Update device templates that are already attached to a device.
*HealthApi* | [**GetDevicesHealthList**](docs/HealthApi.md#getdeviceshealthlist) | **Get** /monitor/api/v8/health/devices/list | 
*HealthApi* | [**GetServicesHealthList**](docs/HealthApi.md#getserviceshealthlist) | **Get** /monitor/api/v8/health/services/list | 
*OffersApi* | [**CreateOffer**](docs/OffersApi.md#createoffer) | **Post** /consume/api/v8/offers | Creates a product offer.
*OffersApi* | [**DeleteOffer**](docs/OffersApi.md#deleteoffer) | **Delete** /consume/api/v8/offers/{id} | Deletes a product offer
*OffersApi* | [**GetOffer**](docs/OffersApi.md#getoffer) | **Get** /consume/api/v8/offers/{id} | Returns a product offer.
*OffersApi* | [**GetOfferAssignmentsList**](docs/OffersApi.md#getofferassignmentslist) | **Get** /consume/api/v8/offers/{id}/assignments/list | Returns a list of tenant assignments for a product offer.
*OffersApi* | [**GetOffersCount**](docs/OffersApi.md#getofferscount) | **Get** /consume/api/v8/offers/count | Returns the number of product offers.
*OffersApi* | [**GetOffersPage**](docs/OffersApi.md#getofferspage) | **Get** /consume/api/v8/offers | Returns a page of product offers.
*OffersApi* | [**UpdateOffer**](docs/OffersApi.md#updateoffer) | **Put** /consume/api/v8/offers/{id} | Updates a product offer.
*OffersApi* | [**UpdateOfferAssignments**](docs/OffersApi.md#updateofferassignments) | **Put** /consume/api/v8/offers/{id}/assignments | Updates the tenant assignemnts for a product offer.
*ProductsApi* | [**CreateProduct**](docs/ProductsApi.md#createproduct) | **Post** /consume/api/v8/products | Creates a product.
*ProductsApi* | [**DeleteProduct**](docs/ProductsApi.md#deleteproduct) | **Delete** /consume/api/v8/products/{id} | Deletes a product.
*ProductsApi* | [**GetProduct**](docs/ProductsApi.md#getproduct) | **Get** /consume/api/v8/products/{id} | Returns a product.
*ProductsApi* | [**GetProductAssignmentsList**](docs/ProductsApi.md#getproductassignmentslist) | **Get** /consume/api/v8/products/{id}/assignments/list | Returns a list of tenant assignments for a product .
*ProductsApi* | [**GetProductsCount**](docs/ProductsApi.md#getproductscount) | **Get** /consume/api/v8/products/count | Returns the number of products.
*ProductsApi* | [**GetProductsPage**](docs/ProductsApi.md#getproductspage) | **Get** /consume/api/v8/products | Returns a page of products.
*ProductsApi* | [**UpdateProduct**](docs/ProductsApi.md#updateproduct) | **Put** /consume/api/v8/products/{id} | Updates a product.
*ProductsApi* | [**UpdateProductAssignments**](docs/ProductsApi.md#updateproductassignments) | **Put** /consume/api/v8/products/{id}/assignments | Updates the tenant assignments for a product.
*RolesApi* | [**GetRoleByName**](docs/RolesApi.md#getrolebyname) | **Get** /idm/api/v8/roles/name/{name} | Returns a role by name.
*RolesApi* | [**GetRolesList**](docs/RolesApi.md#getroleslist) | **Get** /idm/api/v8/roles/list | Returns a list of roles.
*SecurityApi* | [**GetAccessToken**](docs/SecurityApi.md#getaccesstoken) | **Post** /idm/v2/token | Returns an access token.
*ServicesApi* | [**DeleteService**](docs/ServicesApi.md#deleteservice) | **Delete** /manage/api/v8/services/{id} | Deletes a service.
*ServicesApi* | [**GetService**](docs/ServicesApi.md#getservice) | **Get** /manage/api/v8/services/{id} | Returns a service.
*ServicesApi* | [**GetServicesPage**](docs/ServicesApi.md#getservicespage) | **Get** /manage/api/v8/services | Returns a page of services.
*ServicesApi* | [**SubmitOrder**](docs/ServicesApi.md#submitorder) | **Post** /manage/api/v8/services | Submits an order.
*ServicesApi* | [**UpdateOrder**](docs/ServicesApi.md#updateorder) | **Put** /manage/api/v8/services | Updates an order.
*ServicesApi* | [**UpdateService**](docs/ServicesApi.md#updateservice) | **Put** /manage/api/v8/services/{id} | Updates a service.
*SitesApi* | [**AddDevicesToSite**](docs/SitesApi.md#adddevicestosite) | **Post** /manage/api/v8/sites/{id}/devices/add | Add devices to a site.
*SitesApi* | [**AddServicesToSite**](docs/SitesApi.md#addservicestosite) | **Post** /manage/api/v8/sites/{id}/services/add | Add services to a site.
*SitesApi* | [**CreateSite**](docs/SitesApi.md#createsite) | **Post** /manage/api/v8/sites | Creates a new site.
*SitesApi* | [**DeleteSite**](docs/SitesApi.md#deletesite) | **Delete** /manage/api/v8/sites/{id} | Deletes a site.
*SitesApi* | [**GetSite**](docs/SitesApi.md#getsite) | **Get** /manage/api/v8/sites/{id} | Returns a site.
*SitesApi* | [**GetSitesPage**](docs/SitesApi.md#getsitespage) | **Get** /manage/api/v8/sites | Returns a page of Sites. Only one filter is supported at a time.
*SitesApi* | [**RemoveDevicesFromSite**](docs/SitesApi.md#removedevicesfromsite) | **Post** /manage/api/v8/sites/{id}/devices/remove | Removes devices from a site.
*SitesApi* | [**RemoveServicesFromSite**](docs/SitesApi.md#removeservicesfromsite) | **Post** /manage/api/v8/sites/{id}/services/remove | Remove services from a site.
*SitesApi* | [**UpdateSite**](docs/SitesApi.md#updatesite) | **Put** /manage/api/v8/sites/{id} | Updates a site.
*TenantsApi* | [**CreateTenant**](docs/TenantsApi.md#createtenant) | **Post** /idm/api/v8/tenants | Creates a new tenant.
*TenantsApi* | [**DeleteTenant**](docs/TenantsApi.md#deletetenant) | **Delete** /idm/api/v8/tenants/{id} | Deletes a tenant by id.
*TenantsApi* | [**GetTenant**](docs/TenantsApi.md#gettenant) | **Get** /idm/api/v8/tenants/{id} | Returns a tenant by id.
*TenantsApi* | [**GetTenantsList**](docs/TenantsApi.md#gettenantslist) | **Get** /idm/api/v8/tenants/list | Returns a list of tenants.
*TenantsApi* | [**GetTenantsPage**](docs/TenantsApi.md#gettenantspage) | **Get** /idm/api/v8/tenants | Returns a page of tenants.
*TenantsApi* | [**UpdateTenant**](docs/TenantsApi.md#updatetenant) | **Put** /idm/api/v8/tenants/{id} | Updates a tenant by id.
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /idm/api/v8/users | Creates a new user.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /idm/api/v8/users/{id} | Deletes a user by id.
*UsersApi* | [**GetCurrentUser**](docs/UsersApi.md#getcurrentuser) | **Get** /idm/api/v8/users/current | Returns the current user.
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /idm/api/v8/users/{id} | Returns an existing user.
*UsersApi* | [**GetUsersPage**](docs/UsersApi.md#getuserspage) | **Get** /idm/api/v8/users | Returns a page of users.
*UsersApi* | [**UpdateUser**](docs/UsersApi.md#updateuser) | **Put** /idm/api/v8/users/{id} | Updates an existing user.
*UsersApi* | [**UpdateUserPassword**](docs/UsersApi.md#updateuserpassword) | **Put** /idm/api/v8/users/updatepassword | Update a user password.
*WorkflowCategoriesApi* | [**CreateWorkflowCategory**](docs/WorkflowCategoriesApi.md#createworkflowcategory) | **Post** /workflow/api/v8/categories | Creates a new workflow category.
*WorkflowCategoriesApi* | [**DeleteWorkflowCategory**](docs/WorkflowCategoriesApi.md#deleteworkflowcategory) | **Delete** /workflow/api/v8/categories/{id} | Deletes a workflow category.
*WorkflowCategoriesApi* | [**GetWorkflowCategoriesList**](docs/WorkflowCategoriesApi.md#getworkflowcategorieslist) | **Get** /workflow/api/v8/categories/list | Returns a list of workflow categories.
*WorkflowCategoriesApi* | [**GetWorkflowCategory**](docs/WorkflowCategoriesApi.md#getworkflowcategory) | **Get** /workflow/api/v8/categories/{id} | Returns a workflow category.
*WorkflowCategoriesApi* | [**UpdateWorkflowCategory**](docs/WorkflowCategoriesApi.md#updateworkflowcategory) | **Put** /workflow/api/v8/categories/{id} | Updates a workflow category.
*WorkflowEventsApi* | [**CreateWorkflowEvent**](docs/WorkflowEventsApi.md#createworkflowevent) | **Post** /workflow/api/v8/events | Creates a new workflow event.
*WorkflowEventsApi* | [**DeleteWorkflowEvent**](docs/WorkflowEventsApi.md#deleteworkflowevent) | **Delete** /workflow/api/v8/events/{id} | Deletes a workflow event.
*WorkflowEventsApi* | [**GetWorkflowEvent**](docs/WorkflowEventsApi.md#getworkflowevent) | **Get** /workflow/api/v8/events/{id} | Returns a workflow event.
*WorkflowEventsApi* | [**GetWorkflowEventsList**](docs/WorkflowEventsApi.md#getworkfloweventslist) | **Get** /workflow/api/v8/events/list | Returns a list of workflow events.
*WorkflowEventsApi* | [**UpdateWorkflowEvent**](docs/WorkflowEventsApi.md#updateworkflowevent) | **Put** /workflow/api/v8/events/{id} | Updates a workflow event.
*WorkflowInstancesApi* | [**CancelWorkflowInstance**](docs/WorkflowInstancesApi.md#cancelworkflowinstance) | **Post** /workflow/api/v8/workflows/instances/{id}/cancel | Cancels a workflow instance.
*WorkflowInstancesApi* | [**DeleteWorkflowInstance**](docs/WorkflowInstancesApi.md#deleteworkflowinstance) | **Delete** /workflow/api/v8/workflows/instances/{id} | Deletes a workflow instance.
*WorkflowInstancesApi* | [**GetWorkflowInstance**](docs/WorkflowInstancesApi.md#getworkflowinstance) | **Get** /workflow/api/v8/workflows/instances/{id} | Returns a workflow instance.
*WorkflowInstancesApi* | [**GetWorkflowInstanceAction**](docs/WorkflowInstancesApi.md#getworkflowinstanceaction) | **Get** /workflow/api/v8/workflows/instances/{id}/actions/{actionId} | Returns a workflow instance action.
*WorkflowInstancesApi* | [**GetWorkflowInstancesList**](docs/WorkflowInstancesApi.md#getworkflowinstanceslist) | **Get** /workflow/api/v8/workflows/{id}/instances/list | Returns a list of workflow instances.
*WorkflowSchemasApi* | [**GetWorkflowSchema**](docs/WorkflowSchemasApi.md#getworkflowschema) | **Get** /workflow/api/v8/schemas/{id} | Returns a workflow schema.
*WorkflowSchemasApi* | [**GetWorkflowSchemasList**](docs/WorkflowSchemasApi.md#getworkflowschemaslist) | **Get** /workflow/api/v8/schemas/list | Returns a list of workflow schemas.
*WorkflowTargetsApi* | [**CreateWorkflowTarget**](docs/WorkflowTargetsApi.md#createworkflowtarget) | **Post** /workflow/api/v8/targets | Creates a new workflow target.
*WorkflowTargetsApi* | [**DeleteWorkflowTarget**](docs/WorkflowTargetsApi.md#deleteworkflowtarget) | **Delete** /workflow/api/v8/targets/{id} | Deletes a workflow target.
*WorkflowTargetsApi* | [**GetWorkflowTarget**](docs/WorkflowTargetsApi.md#getworkflowtarget) | **Get** /workflow/api/v8/targets/{id} | Returns a workflow target.
*WorkflowTargetsApi* | [**GetWorkflowTargetsList**](docs/WorkflowTargetsApi.md#getworkflowtargetslist) | **Get** /workflow/api/v8/targets/list | Returns a list of workflow targets.
*WorkflowTargetsApi* | [**UpdateWorkflowTarget**](docs/WorkflowTargetsApi.md#updateworkflowtarget) | **Put** /workflow/api/v8/targets/{id} | Updates a workflow target.
*WorkflowsApi* | [**DeleteWorkflow**](docs/WorkflowsApi.md#deleteworkflow) | **Delete** /workflow/api/v8/workflows/{id} | Delete a workflow.
*WorkflowsApi* | [**ExportWorkflow**](docs/WorkflowsApi.md#exportworkflow) | **Get** /workflow/api/v8/workflows/{id}/export | Exports a workflow.
*WorkflowsApi* | [**GetWorkflow**](docs/WorkflowsApi.md#getworkflow) | **Get** /workflow/api/v8/workflows/{id} | Returns a workflow.
*WorkflowsApi* | [**GetWorkflowStartConfig**](docs/WorkflowsApi.md#getworkflowstartconfig) | **Get** /workflow/api/v8/workflows/{id}/startconfig | Returns a workflow start config.
*WorkflowsApi* | [**GetWorkflowsList**](docs/WorkflowsApi.md#getworkflowslist) | **Get** /workflow/api/v8/workflows/list | Returns a list of workflows.
*WorkflowsApi* | [**ImportWorkflow**](docs/WorkflowsApi.md#importworkflow) | **Post** /workflow/api/v8/workflows | Imports a workflow.
*WorkflowsApi* | [**StartWorkflow**](docs/WorkflowsApi.md#startworkflow) | **Post** /workflow/api/v8/workflows/{id}/start | Starts a workflow.
*WorkflowsApi* | [**UpdateWorkflow**](docs/WorkflowsApi.md#updateworkflow) | **Put** /workflow/api/v8/workflows/{id} | Updates a workflow.
*WorkflowsApi* | [**ValidateWorkflow**](docs/WorkflowsApi.md#validateworkflow) | **Post** /workflow/api/v8/workflows/{id}/validate | Validates a workflow.


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [CatalogAssignment](docs/CatalogAssignment.md)
 - [Device](docs/Device.md)
 - [DeviceAllOf](docs/DeviceAllOf.md)
 - [DeviceComplianceState](docs/DeviceComplianceState.md)
 - [DeviceCreate](docs/DeviceCreate.md)
 - [DeviceCreateAllOf](docs/DeviceCreateAllOf.md)
 - [DevicePatch](docs/DevicePatch.md)
 - [DeviceSummary](docs/DeviceSummary.md)
 - [DeviceTemplate](docs/DeviceTemplate.md)
 - [DeviceTemplateAccess](docs/DeviceTemplateAccess.md)
 - [DeviceTemplateAccessResponse](docs/DeviceTemplateAccessResponse.md)
 - [DeviceTemplateAttachRequest](docs/DeviceTemplateAttachRequest.md)
 - [DeviceTemplateCreate](docs/DeviceTemplateCreate.md)
 - [DeviceTemplateDetails](docs/DeviceTemplateDetails.md)
 - [DeviceTemplateHistory](docs/DeviceTemplateHistory.md)
 - [DeviceTemplateUpdateDetails](docs/DeviceTemplateUpdateDetails.md)
 - [DeviceTemplateUpdateRequest](docs/DeviceTemplateUpdateRequest.md)
 - [DeviceTemplateVersionCreate](docs/DeviceTemplateVersionCreate.md)
 - [DeviceUpdate](docs/DeviceUpdate.md)
 - [DeviceUpdateAllOf](docs/DeviceUpdateAllOf.md)
 - [DeviceVulnerabilityState](docs/DeviceVulnerabilityState.md)
 - [DevicesPage](docs/DevicesPage.md)
 - [DevicesPageAllOf](docs/DevicesPageAllOf.md)
 - [Error](docs/Error.md)
 - [GenericEvent](docs/GenericEvent.md)
 - [GenericEventAllOf](docs/GenericEventAllOf.md)
 - [GenericEventCreate](docs/GenericEventCreate.md)
 - [GenericEventSecurity](docs/GenericEventSecurity.md)
 - [GenericEventSeverity](docs/GenericEventSeverity.md)
 - [GenericEventTrace](docs/GenericEventTrace.md)
 - [LegacyAbsoluteConfig](docs/LegacyAbsoluteConfig.md)
 - [LegacyAddress](docs/LegacyAddress.md)
 - [LegacyNsoResponseTypes](docs/LegacyNsoResponseTypes.md)
 - [LegacyRelativeConfig](docs/LegacyRelativeConfig.md)
 - [LegacyScheduleConfig](docs/LegacyScheduleConfig.md)
 - [LegacyServiceOrder](docs/LegacyServiceOrder.md)
 - [LegacyServiceOrderDetail](docs/LegacyServiceOrderDetail.md)
 - [LegacyServiceOrderResponse](docs/LegacyServiceOrderResponse.md)
 - [LegacySite](docs/LegacySite.md)
 - [LegacySiteDevice](docs/LegacySiteDevice.md)
 - [LegacySiteDeviceOnboard](docs/LegacySiteDeviceOnboard.md)
 - [LegacySubscriptionDetail](docs/LegacySubscriptionDetail.md)
 - [NSOConfigDataXPath](docs/NSOConfigDataXPath.md)
 - [NameValue](docs/NameValue.md)
 - [Offer](docs/Offer.md)
 - [OfferAllOf](docs/OfferAllOf.md)
 - [OfferCreate](docs/OfferCreate.md)
 - [OfferUpdate](docs/OfferUpdate.md)
 - [OffersPage](docs/OffersPage.md)
 - [OffersPageAllOf](docs/OffersPageAllOf.md)
 - [PageHeader](docs/PageHeader.md)
 - [Product](docs/Product.md)
 - [ProductAllOf](docs/ProductAllOf.md)
 - [ProductCreate](docs/ProductCreate.md)
 - [ProductUpdate](docs/ProductUpdate.md)
 - [ProductsPage](docs/ProductsPage.md)
 - [ProductsPageAllOf](docs/ProductsPageAllOf.md)
 - [ResourceHealth](docs/ResourceHealth.md)
 - [ResourceStatus](docs/ResourceStatus.md)
 - [ResourceType](docs/ResourceType.md)
 - [Role](docs/Role.md)
 - [Service](docs/Service.md)
 - [ServiceAllOf](docs/ServiceAllOf.md)
 - [ServiceElement](docs/ServiceElement.md)
 - [ServiceElementPrice](docs/ServiceElementPrice.md)
 - [ServiceUIConfig](docs/ServiceUIConfig.md)
 - [ServiceUILink](docs/ServiceUILink.md)
 - [ServiceUpdate](docs/ServiceUpdate.md)
 - [ServicesPage](docs/ServicesPage.md)
 - [ServicesPageAllOf](docs/ServicesPageAllOf.md)
 - [Site](docs/Site.md)
 - [SiteAddress](docs/SiteAddress.md)
 - [SiteContact](docs/SiteContact.md)
 - [SiteCreate](docs/SiteCreate.md)
 - [SiteCreateAllOf](docs/SiteCreateAllOf.md)
 - [SiteLocation](docs/SiteLocation.md)
 - [SiteStatus](docs/SiteStatus.md)
 - [SiteUpdate](docs/SiteUpdate.md)
 - [SitesPage](docs/SitesPage.md)
 - [SitesPageAllOf](docs/SitesPageAllOf.md)
 - [StartWorkflowResponse](docs/StartWorkflowResponse.md)
 - [TemplateParameterValidator](docs/TemplateParameterValidator.md)
 - [Tenant](docs/Tenant.md)
 - [TenantAllOf](docs/TenantAllOf.md)
 - [TenantCreate](docs/TenantCreate.md)
 - [TenantCreateAllOf](docs/TenantCreateAllOf.md)
 - [TenantUpdate](docs/TenantUpdate.md)
 - [TenantsPage](docs/TenantsPage.md)
 - [TenantsPageAllOf](docs/TenantsPageAllOf.md)
 - [UpdatePassword](docs/UpdatePassword.md)
 - [User](docs/User.md)
 - [UserAllOf](docs/UserAllOf.md)
 - [UserCreate](docs/UserCreate.md)
 - [UserCreateAllOf](docs/UserCreateAllOf.md)
 - [UserUpdate](docs/UserUpdate.md)
 - [UsersPage](docs/UsersPage.md)
 - [UsersPageAllOf](docs/UsersPageAllOf.md)
 - [ValidateWorkflowResponse](docs/ValidateWorkflowResponse.md)
 - [Workflow](docs/Workflow.md)
 - [WorkflowAccessMeta](docs/WorkflowAccessMeta.md)
 - [WorkflowAccessMetaType](docs/WorkflowAccessMetaType.md)
 - [WorkflowAction](docs/WorkflowAction.md)
 - [WorkflowActionBlock](docs/WorkflowActionBlock.md)
 - [WorkflowAllOf](docs/WorkflowAllOf.md)
 - [WorkflowCategory](docs/WorkflowCategory.md)
 - [WorkflowCategoryAllOf](docs/WorkflowCategoryAllOf.md)
 - [WorkflowCategoryCreate](docs/WorkflowCategoryCreate.md)
 - [WorkflowCategoryUpdate](docs/WorkflowCategoryUpdate.md)
 - [WorkflowDefAccessMeta](docs/WorkflowDefAccessMeta.md)
 - [WorkflowEvent](docs/WorkflowEvent.md)
 - [WorkflowEventAllOf](docs/WorkflowEventAllOf.md)
 - [WorkflowEventCreate](docs/WorkflowEventCreate.md)
 - [WorkflowEventUpdate](docs/WorkflowEventUpdate.md)
 - [WorkflowFooter](docs/WorkflowFooter.md)
 - [WorkflowInstance](docs/WorkflowInstance.md)
 - [WorkflowInstanceAllOf](docs/WorkflowInstanceAllOf.md)
 - [WorkflowInstanceDeleteResponse](docs/WorkflowInstanceDeleteResponse.md)
 - [WorkflowMapping](docs/WorkflowMapping.md)
 - [WorkflowMetadata](docs/WorkflowMetadata.md)
 - [WorkflowMetadataGitInfo](docs/WorkflowMetadataGitInfo.md)
 - [WorkflowSchema](docs/WorkflowSchema.md)
 - [WorkflowSchemaAllOf](docs/WorkflowSchemaAllOf.md)
 - [WorkflowSchemaByTypeResponse](docs/WorkflowSchemaByTypeResponse.md)
 - [WorkflowStartConfig](docs/WorkflowStartConfig.md)
 - [WorkflowTarget](docs/WorkflowTarget.md)
 - [WorkflowTargetAllOf](docs/WorkflowTargetAllOf.md)
 - [WorkflowTargetCreate](docs/WorkflowTargetCreate.md)
 - [WorkflowTargetUpdate](docs/WorkflowTargetUpdate.md)
 - [WorkflowVariable](docs/WorkflowVariable.md)
 - [WorkflowVariableAllOf](docs/WorkflowVariableAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



