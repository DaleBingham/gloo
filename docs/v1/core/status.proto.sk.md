<!-- Code generated by protoc-gen-solo-kit. DO NOT EDIT. -->

## Package:
core.solo.io

## Source File:
github.com/solo-io/solo-kit/api/v1/status.proto 

## Description:  

## Contents:
- Messages:  
	- [Status](#Status)  
	- [SubresourceStatusesEntry](#SubresourceStatusesEntry)

---
  
### <a name="Status">Status</a>

Description: *
Status indicates whether a resource has been (in)validated by a reporter in the system.
Statuses are meant to be read-only by users

```yaml
"state": ***TODO ENUMS***!
"reason": string
"reported_by": string
"subresource_statuses": map<string, .core.solo.io.Status>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| state | ***TODO ENUMS***! | State is the enum indicating the state of the resource |  |
| reason | string | Reason is a description of the error for Rejected resources. If the resource is pending or accepted, this field will be empty |  |
| reported_by | string | Reference to the reporter who wrote this status |  |
| subresource_statuses | [map<string, .core.solo.io.Status>](status.proto.sk.md#SubresourceStatusesEntry) | Reference to statuses (by resource-ref string: "Kind.Namespace.Name") of subresources of the parent resource |  |
  
### <a name="SubresourceStatusesEntry">SubresourceStatusesEntry</a>

Description: 

```yaml
"key": string
"value": .core.solo.io.Status

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| key | string |  |  |
| value | [.core.solo.io.Status](status.proto.sk.md#Status) |  |  |


<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->
