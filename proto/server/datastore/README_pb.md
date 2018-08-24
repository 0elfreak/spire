# Protocol Documentation
<a name="top"/>

## Table of Contents

- [wrappers.proto](#wrappers.proto)
    - [BoolValue](#google.protobuf.BoolValue)
    - [BytesValue](#google.protobuf.BytesValue)
    - [DoubleValue](#google.protobuf.DoubleValue)
    - [FloatValue](#google.protobuf.FloatValue)
    - [Int32Value](#google.protobuf.Int32Value)
    - [Int64Value](#google.protobuf.Int64Value)
    - [StringValue](#google.protobuf.StringValue)
    - [UInt32Value](#google.protobuf.UInt32Value)
    - [UInt64Value](#google.protobuf.UInt64Value)
  
  
  
  

- [plugin.proto](#plugin.proto)
    - [ConfigureRequest](#spire.common.plugin.ConfigureRequest)
    - [ConfigureResponse](#spire.common.plugin.ConfigureResponse)
    - [GetPluginInfoRequest](#spire.common.plugin.GetPluginInfoRequest)
    - [GetPluginInfoResponse](#spire.common.plugin.GetPluginInfoResponse)
  
  
  
  

- [common.proto](#common.proto)
    - [AttestationData](#spire.common.AttestationData)
    - [Empty](#spire.common.Empty)
    - [RegistrationEntries](#spire.common.RegistrationEntries)
    - [RegistrationEntry](#spire.common.RegistrationEntry)
    - [Selector](#spire.common.Selector)
    - [Selectors](#spire.common.Selectors)
  
  
  
  

- [datastore.proto](#datastore.proto)
    - [AppendBundleRequest](#spire.server.datastore.AppendBundleRequest)
    - [AppendBundleResponse](#spire.server.datastore.AppendBundleResponse)
    - [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry)
    - [Bundle](#spire.server.datastore.Bundle)
    - [BySelectors](#spire.server.datastore.BySelectors)
    - [CreateAttestedNodeEntryRequest](#spire.server.datastore.CreateAttestedNodeEntryRequest)
    - [CreateAttestedNodeEntryResponse](#spire.server.datastore.CreateAttestedNodeEntryResponse)
    - [CreateBundleRequest](#spire.server.datastore.CreateBundleRequest)
    - [CreateBundleResponse](#spire.server.datastore.CreateBundleResponse)
    - [CreateJoinTokenRequest](#spire.server.datastore.CreateJoinTokenRequest)
    - [CreateJoinTokenResponse](#spire.server.datastore.CreateJoinTokenResponse)
    - [CreateNodeResolverMapEntryRequest](#spire.server.datastore.CreateNodeResolverMapEntryRequest)
    - [CreateNodeResolverMapEntryResponse](#spire.server.datastore.CreateNodeResolverMapEntryResponse)
    - [CreateRegistrationEntryRequest](#spire.server.datastore.CreateRegistrationEntryRequest)
    - [CreateRegistrationEntryResponse](#spire.server.datastore.CreateRegistrationEntryResponse)
    - [DeleteAttestedNodeEntryRequest](#spire.server.datastore.DeleteAttestedNodeEntryRequest)
    - [DeleteAttestedNodeEntryResponse](#spire.server.datastore.DeleteAttestedNodeEntryResponse)
    - [DeleteBundleRequest](#spire.server.datastore.DeleteBundleRequest)
    - [DeleteBundleResponse](#spire.server.datastore.DeleteBundleResponse)
    - [DeleteJoinTokenRequest](#spire.server.datastore.DeleteJoinTokenRequest)
    - [DeleteJoinTokenResponse](#spire.server.datastore.DeleteJoinTokenResponse)
    - [DeleteNodeResolverMapEntryRequest](#spire.server.datastore.DeleteNodeResolverMapEntryRequest)
    - [DeleteNodeResolverMapEntryResponse](#spire.server.datastore.DeleteNodeResolverMapEntryResponse)
    - [DeleteRegistrationEntryRequest](#spire.server.datastore.DeleteRegistrationEntryRequest)
    - [DeleteRegistrationEntryResponse](#spire.server.datastore.DeleteRegistrationEntryResponse)
    - [FetchAttestedNodeEntryRequest](#spire.server.datastore.FetchAttestedNodeEntryRequest)
    - [FetchAttestedNodeEntryResponse](#spire.server.datastore.FetchAttestedNodeEntryResponse)
    - [FetchBundleRequest](#spire.server.datastore.FetchBundleRequest)
    - [FetchBundleResponse](#spire.server.datastore.FetchBundleResponse)
    - [FetchJoinTokenRequest](#spire.server.datastore.FetchJoinTokenRequest)
    - [FetchJoinTokenResponse](#spire.server.datastore.FetchJoinTokenResponse)
    - [FetchRegistrationEntryRequest](#spire.server.datastore.FetchRegistrationEntryRequest)
    - [FetchRegistrationEntryResponse](#spire.server.datastore.FetchRegistrationEntryResponse)
    - [JoinToken](#spire.server.datastore.JoinToken)
    - [ListAttestedNodeEntriesRequest](#spire.server.datastore.ListAttestedNodeEntriesRequest)
    - [ListAttestedNodeEntriesResponse](#spire.server.datastore.ListAttestedNodeEntriesResponse)
    - [ListBundlesRequest](#spire.server.datastore.ListBundlesRequest)
    - [ListBundlesResponse](#spire.server.datastore.ListBundlesResponse)
    - [ListNodeResolverMapEntriesRequest](#spire.server.datastore.ListNodeResolverMapEntriesRequest)
    - [ListNodeResolverMapEntriesResponse](#spire.server.datastore.ListNodeResolverMapEntriesResponse)
    - [ListRegistrationEntriesRequest](#spire.server.datastore.ListRegistrationEntriesRequest)
    - [ListRegistrationEntriesResponse](#spire.server.datastore.ListRegistrationEntriesResponse)
    - [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry)
    - [PruneJoinTokensRequest](#spire.server.datastore.PruneJoinTokensRequest)
    - [PruneJoinTokensResponse](#spire.server.datastore.PruneJoinTokensResponse)
    - [RectifyNodeResolverMapEntriesRequest](#spire.server.datastore.RectifyNodeResolverMapEntriesRequest)
    - [RectifyNodeResolverMapEntriesResponse](#spire.server.datastore.RectifyNodeResolverMapEntriesResponse)
    - [UpdateAttestedNodeEntryRequest](#spire.server.datastore.UpdateAttestedNodeEntryRequest)
    - [UpdateAttestedNodeEntryResponse](#spire.server.datastore.UpdateAttestedNodeEntryResponse)
    - [UpdateBundleRequest](#spire.server.datastore.UpdateBundleRequest)
    - [UpdateBundleResponse](#spire.server.datastore.UpdateBundleResponse)
    - [UpdateRegistrationEntryRequest](#spire.server.datastore.UpdateRegistrationEntryRequest)
    - [UpdateRegistrationEntryResponse](#spire.server.datastore.UpdateRegistrationEntryResponse)
  
  
  
    - [DataStore](#spire.server.datastore.DataStore)
  

- [Scalar Value Types](#scalar-value-types)



<a name="wrappers.proto"/>
<p align="right"><a href="#top">Top</a></p>

## wrappers.proto



<a name="google.protobuf.BoolValue"/>

### BoolValue
Wrapper message for `bool`.

The JSON representation for `BoolValue` is JSON `true` and `false`.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bool](#bool) |  | The bool value. |






<a name="google.protobuf.BytesValue"/>

### BytesValue
Wrapper message for `bytes`.

The JSON representation for `BytesValue` is JSON string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  | The bytes value. |






<a name="google.protobuf.DoubleValue"/>

### DoubleValue
Wrapper message for `double`.

The JSON representation for `DoubleValue` is JSON number.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [double](#double) |  | The double value. |






<a name="google.protobuf.FloatValue"/>

### FloatValue
Wrapper message for `float`.

The JSON representation for `FloatValue` is JSON number.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [float](#float) |  | The float value. |






<a name="google.protobuf.Int32Value"/>

### Int32Value
Wrapper message for `int32`.

The JSON representation for `Int32Value` is JSON number.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int32](#int32) |  | The int32 value. |






<a name="google.protobuf.Int64Value"/>

### Int64Value
Wrapper message for `int64`.

The JSON representation for `Int64Value` is JSON string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [int64](#int64) |  | The int64 value. |






<a name="google.protobuf.StringValue"/>

### StringValue
Wrapper message for `string`.

The JSON representation for `StringValue` is JSON string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | The string value. |






<a name="google.protobuf.UInt32Value"/>

### UInt32Value
Wrapper message for `uint32`.

The JSON representation for `UInt32Value` is JSON number.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [uint32](#uint32) |  | The uint32 value. |






<a name="google.protobuf.UInt64Value"/>

### UInt64Value
Wrapper message for `uint64`.

The JSON representation for `UInt64Value` is JSON string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [uint64](#uint64) |  | The uint64 value. |





 

 

 

 



<a name="plugin.proto"/>
<p align="right"><a href="#top">Top</a></p>

## plugin.proto



<a name="spire.common.plugin.ConfigureRequest"/>

### ConfigureRequest
Represents the plugin-specific configuration string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| configuration | [string](#string) |  | The configuration for the plugin. |






<a name="spire.common.plugin.ConfigureResponse"/>

### ConfigureResponse
Represents a list of configuration problems
found in the configuration string.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| errorList | [string](#string) | repeated | A list of errors |






<a name="spire.common.plugin.GetPluginInfoRequest"/>

### GetPluginInfoRequest
Represents an empty request.






<a name="spire.common.plugin.GetPluginInfoResponse"/>

### GetPluginInfoResponse
Represents the plugin metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| category | [string](#string) |  |  |
| type | [string](#string) |  |  |
| description | [string](#string) |  |  |
| dateCreated | [string](#string) |  |  |
| location | [string](#string) |  |  |
| version | [string](#string) |  |  |
| author | [string](#string) |  |  |
| company | [string](#string) |  |  |





 

 

 

 



<a name="common.proto"/>
<p align="right"><a href="#top">Top</a></p>

## common.proto



<a name="spire.common.AttestationData"/>

### AttestationData
A type which contains attestation data for specific platform.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | Type of attestation to perform. |
| data | [bytes](#bytes) |  | The attestation data. |






<a name="spire.common.Empty"/>

### Empty
Represents an empty message






<a name="spire.common.RegistrationEntries"/>

### RegistrationEntries
A list of registration entries.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [RegistrationEntry](#spire.common.RegistrationEntry) | repeated | A list of RegistrationEntry. |






<a name="spire.common.RegistrationEntry"/>

### RegistrationEntry
This is a curated record that the Server uses to set up and
manage the various registered nodes and workloads that are controlled by it.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selectors | [Selector](#spire.common.Selector) | repeated | A list of selectors. |
| parent_id | [string](#string) |  | The SPIFFE ID of an entity that is authorized to attest the validity of a selector |
| spiffe_id | [string](#string) |  | The SPIFFE ID is a structured string used to identify a resource or caller. It is defined as a URI comprising a “trust domain” and an associated path. |
| ttl | [int32](#int32) |  | Time to live. |
| fb_spiffe_ids | [string](#string) | repeated | A list of federated bundle spiffe ids. |
| entry_id | [string](#string) |  | Entry ID |






<a name="spire.common.Selector"/>

### Selector
A type which describes the conditions under which a registration
entry is matched.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | A selector type represents the type of attestation used in attesting the entity (Eg: AWS, K8). |
| value | [string](#string) |  | The value to be attested. |






<a name="spire.common.Selectors"/>

### Selectors
Represents a type with a list of Selector.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [Selector](#spire.common.Selector) | repeated | A list of Selector. |





 

 

 

 



<a name="datastore.proto"/>
<p align="right"><a href="#top">Top</a></p>

## datastore.proto



<a name="spire.server.datastore.AppendBundleRequest"/>

### AppendBundleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.AppendBundleResponse"/>

### AppendBundleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.AttestedNodeEntry"/>

### AttestedNodeEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  | Node SPIFFE ID |
| attestation_data_type | [string](#string) |  | Attestation data type |
| cert_serial_number | [string](#string) |  | Node certificate serial number |
| cert_not_after | [int64](#int64) |  | Node certificate not_after (seconds since unix epoch) |






<a name="spire.server.datastore.Bundle"/>

### Bundle



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trust_domain | [string](#string) |  | Trust domain SPIFFE ID |
| ca_certs | [bytes](#bytes) |  | CA Certificates (ASN.1 DER encoded) |






<a name="spire.server.datastore.BySelectors"/>

### BySelectors



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| selectors | [.spire.common.Selector](#spire.server.datastore..spire.common.Selector) | repeated |  |
| allow_any_combination | [bool](#bool) |  |  |






<a name="spire.server.datastore.CreateAttestedNodeEntryRequest"/>

### CreateAttestedNodeEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) |  |  |






<a name="spire.server.datastore.CreateAttestedNodeEntryResponse"/>

### CreateAttestedNodeEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) |  |  |






<a name="spire.server.datastore.CreateBundleRequest"/>

### CreateBundleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.CreateBundleResponse"/>

### CreateBundleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.CreateJoinTokenRequest"/>

### CreateJoinTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| join_token | [JoinToken](#spire.server.datastore.JoinToken) |  |  |






<a name="spire.server.datastore.CreateJoinTokenResponse"/>

### CreateJoinTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| join_token | [JoinToken](#spire.server.datastore.JoinToken) |  |  |






<a name="spire.server.datastore.CreateNodeResolverMapEntryRequest"/>

### CreateNodeResolverMapEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) |  |  |






<a name="spire.server.datastore.CreateNodeResolverMapEntryResponse"/>

### CreateNodeResolverMapEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) |  |  |






<a name="spire.server.datastore.CreateRegistrationEntryRequest"/>

### CreateRegistrationEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) |  |  |






<a name="spire.server.datastore.CreateRegistrationEntryResponse"/>

### CreateRegistrationEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry_id | [string](#string) |  |  |






<a name="spire.server.datastore.DeleteAttestedNodeEntryRequest"/>

### DeleteAttestedNodeEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  |  |






<a name="spire.server.datastore.DeleteAttestedNodeEntryResponse"/>

### DeleteAttestedNodeEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) |  |  |






<a name="spire.server.datastore.DeleteBundleRequest"/>

### DeleteBundleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trust_domain | [string](#string) |  |  |






<a name="spire.server.datastore.DeleteBundleResponse"/>

### DeleteBundleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.DeleteJoinTokenRequest"/>

### DeleteJoinTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="spire.server.datastore.DeleteJoinTokenResponse"/>

### DeleteJoinTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| join_token | [JoinToken](#spire.server.datastore.JoinToken) |  |  |






<a name="spire.server.datastore.DeleteNodeResolverMapEntryRequest"/>

### DeleteNodeResolverMapEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) |  |  |






<a name="spire.server.datastore.DeleteNodeResolverMapEntryResponse"/>

### DeleteNodeResolverMapEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) | repeated |  |






<a name="spire.server.datastore.DeleteRegistrationEntryRequest"/>

### DeleteRegistrationEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry_id | [string](#string) |  |  |






<a name="spire.server.datastore.DeleteRegistrationEntryResponse"/>

### DeleteRegistrationEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) |  |  |






<a name="spire.server.datastore.FetchAttestedNodeEntryRequest"/>

### FetchAttestedNodeEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  |  |






<a name="spire.server.datastore.FetchAttestedNodeEntryResponse"/>

### FetchAttestedNodeEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) |  |  |






<a name="spire.server.datastore.FetchBundleRequest"/>

### FetchBundleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| trust_domain | [string](#string) |  |  |






<a name="spire.server.datastore.FetchBundleResponse"/>

### FetchBundleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.FetchJoinTokenRequest"/>

### FetchJoinTokenRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="spire.server.datastore.FetchJoinTokenResponse"/>

### FetchJoinTokenResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| join_token | [JoinToken](#spire.server.datastore.JoinToken) |  |  |






<a name="spire.server.datastore.FetchRegistrationEntryRequest"/>

### FetchRegistrationEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry_id | [string](#string) |  |  |






<a name="spire.server.datastore.FetchRegistrationEntryResponse"/>

### FetchRegistrationEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) |  |  |






<a name="spire.server.datastore.JoinToken"/>

### JoinToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Token value |
| expiry | [int64](#int64) |  | Expiration in seconds since unix epoch |






<a name="spire.server.datastore.ListAttestedNodeEntriesRequest"/>

### ListAttestedNodeEntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| by_expires_before | [.google.protobuf.Int64Value](#spire.server.datastore..google.protobuf.Int64Value) |  |  |






<a name="spire.server.datastore.ListAttestedNodeEntriesResponse"/>

### ListAttestedNodeEntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) | repeated |  |






<a name="spire.server.datastore.ListBundlesRequest"/>

### ListBundlesRequest







<a name="spire.server.datastore.ListBundlesResponse"/>

### ListBundlesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundles | [Bundle](#spire.server.datastore.Bundle) | repeated |  |






<a name="spire.server.datastore.ListNodeResolverMapEntriesRequest"/>

### ListNodeResolverMapEntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  |  |






<a name="spire.server.datastore.ListNodeResolverMapEntriesResponse"/>

### ListNodeResolverMapEntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) | repeated |  |






<a name="spire.server.datastore.ListRegistrationEntriesRequest"/>

### ListRegistrationEntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| by_parent_id | [.google.protobuf.StringValue](#spire.server.datastore..google.protobuf.StringValue) |  |  |
| by_selectors | [BySelectors](#spire.server.datastore.BySelectors) |  |  |
| by_spiffe_id | [.google.protobuf.StringValue](#spire.server.datastore..google.protobuf.StringValue) |  |  |






<a name="spire.server.datastore.ListRegistrationEntriesResponse"/>

### ListRegistrationEntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) | repeated |  |






<a name="spire.server.datastore.NodeResolverMapEntry"/>

### NodeResolverMapEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  | Node SPIFFE ID |
| selector | [.spire.common.Selector](#spire.server.datastore..spire.common.Selector) |  | Node selector |






<a name="spire.server.datastore.PruneJoinTokensRequest"/>

### PruneJoinTokensRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| expires_before | [int64](#int64) |  |  |






<a name="spire.server.datastore.PruneJoinTokensResponse"/>

### PruneJoinTokensResponse







<a name="spire.server.datastore.RectifyNodeResolverMapEntriesRequest"/>

### RectifyNodeResolverMapEntriesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) | repeated |  |






<a name="spire.server.datastore.RectifyNodeResolverMapEntriesResponse"/>

### RectifyNodeResolverMapEntriesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entries | [NodeResolverMapEntry](#spire.server.datastore.NodeResolverMapEntry) | repeated |  |






<a name="spire.server.datastore.UpdateAttestedNodeEntryRequest"/>

### UpdateAttestedNodeEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| spiffe_id | [string](#string) |  |  |
| cert_serial_number | [string](#string) |  |  |
| cert_not_after | [int64](#int64) |  |  |






<a name="spire.server.datastore.UpdateAttestedNodeEntryResponse"/>

### UpdateAttestedNodeEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [AttestedNodeEntry](#spire.server.datastore.AttestedNodeEntry) |  |  |






<a name="spire.server.datastore.UpdateBundleRequest"/>

### UpdateBundleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.UpdateBundleResponse"/>

### UpdateBundleResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle | [Bundle](#spire.server.datastore.Bundle) |  |  |






<a name="spire.server.datastore.UpdateRegistrationEntryRequest"/>

### UpdateRegistrationEntryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) |  |  |






<a name="spire.server.datastore.UpdateRegistrationEntryResponse"/>

### UpdateRegistrationEntryResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| entry | [.spire.common.RegistrationEntry](#spire.server.datastore..spire.common.RegistrationEntry) |  |  |





 

 

 


<a name="spire.server.datastore.DataStore"/>

### DataStore


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBundle | [CreateBundleRequest](#spire.server.datastore.CreateBundleRequest) | [CreateBundleResponse](#spire.server.datastore.CreateBundleRequest) | Creates a bundle |
| FetchBundle | [FetchBundleRequest](#spire.server.datastore.FetchBundleRequest) | [FetchBundleResponse](#spire.server.datastore.FetchBundleRequest) | Fetches a specific bundle |
| ListBundles | [ListBundlesRequest](#spire.server.datastore.ListBundlesRequest) | [ListBundlesResponse](#spire.server.datastore.ListBundlesRequest) | Lists bundles (optionally filtered) |
| UpdateBundle | [UpdateBundleRequest](#spire.server.datastore.UpdateBundleRequest) | [UpdateBundleResponse](#spire.server.datastore.UpdateBundleRequest) | Updates a specific bundle, overwriting existing certs |
| AppendBundle | [AppendBundleRequest](#spire.server.datastore.AppendBundleRequest) | [AppendBundleResponse](#spire.server.datastore.AppendBundleRequest) | Appends the provided certs onto an existing bundle, creating a new bundle if one doesn&#39;t exist |
| DeleteBundle | [DeleteBundleRequest](#spire.server.datastore.DeleteBundleRequest) | [DeleteBundleResponse](#spire.server.datastore.DeleteBundleRequest) | Deletes a specific bundle |
| CreateAttestedNodeEntry | [CreateAttestedNodeEntryRequest](#spire.server.datastore.CreateAttestedNodeEntryRequest) | [CreateAttestedNodeEntryResponse](#spire.server.datastore.CreateAttestedNodeEntryRequest) | Creates an attested node entry |
| FetchAttestedNodeEntry | [FetchAttestedNodeEntryRequest](#spire.server.datastore.FetchAttestedNodeEntryRequest) | [FetchAttestedNodeEntryResponse](#spire.server.datastore.FetchAttestedNodeEntryRequest) | Fetches a specific attested node entry |
| ListAttestedNodeEntries | [ListAttestedNodeEntriesRequest](#spire.server.datastore.ListAttestedNodeEntriesRequest) | [ListAttestedNodeEntriesResponse](#spire.server.datastore.ListAttestedNodeEntriesRequest) | Lists attested node entries (optionally filtered) |
| UpdateAttestedNodeEntry | [UpdateAttestedNodeEntryRequest](#spire.server.datastore.UpdateAttestedNodeEntryRequest) | [UpdateAttestedNodeEntryResponse](#spire.server.datastore.UpdateAttestedNodeEntryRequest) | Updates a specific attested node entry |
| DeleteAttestedNodeEntry | [DeleteAttestedNodeEntryRequest](#spire.server.datastore.DeleteAttestedNodeEntryRequest) | [DeleteAttestedNodeEntryResponse](#spire.server.datastore.DeleteAttestedNodeEntryRequest) | Deletes a specific attested node entry |
| CreateNodeResolverMapEntry | [CreateNodeResolverMapEntryRequest](#spire.server.datastore.CreateNodeResolverMapEntryRequest) | [CreateNodeResolverMapEntryResponse](#spire.server.datastore.CreateNodeResolverMapEntryRequest) | Creates a node resolver map entry |
| ListNodeResolverMapEntries | [ListNodeResolverMapEntriesRequest](#spire.server.datastore.ListNodeResolverMapEntriesRequest) | [ListNodeResolverMapEntriesResponse](#spire.server.datastore.ListNodeResolverMapEntriesRequest) | Lists node resolver map entries for a specified SPIFFE ID |
| DeleteNodeResolverMapEntry | [DeleteNodeResolverMapEntryRequest](#spire.server.datastore.DeleteNodeResolverMapEntryRequest) | [DeleteNodeResolverMapEntryResponse](#spire.server.datastore.DeleteNodeResolverMapEntryRequest) | Deletes a specific node resolver map entry |
| RectifyNodeResolverMapEntries | [RectifyNodeResolverMapEntriesRequest](#spire.server.datastore.RectifyNodeResolverMapEntriesRequest) | [RectifyNodeResolverMapEntriesResponse](#spire.server.datastore.RectifyNodeResolverMapEntriesRequest) | Sets the list of node resolver map entries for the specified SPIFFE ID |
| CreateRegistrationEntry | [CreateRegistrationEntryRequest](#spire.server.datastore.CreateRegistrationEntryRequest) | [CreateRegistrationEntryResponse](#spire.server.datastore.CreateRegistrationEntryRequest) | Creates a registration entry |
| FetchRegistrationEntry | [FetchRegistrationEntryRequest](#spire.server.datastore.FetchRegistrationEntryRequest) | [FetchRegistrationEntryResponse](#spire.server.datastore.FetchRegistrationEntryRequest) | Fetches a specific registration entry |
| ListRegistrationEntries | [ListRegistrationEntriesRequest](#spire.server.datastore.ListRegistrationEntriesRequest) | [ListRegistrationEntriesResponse](#spire.server.datastore.ListRegistrationEntriesRequest) | Lists registration entries (optionally filtered) |
| UpdateRegistrationEntry | [UpdateRegistrationEntryRequest](#spire.server.datastore.UpdateRegistrationEntryRequest) | [UpdateRegistrationEntryResponse](#spire.server.datastore.UpdateRegistrationEntryRequest) | Updates a specific registration entry |
| DeleteRegistrationEntry | [DeleteRegistrationEntryRequest](#spire.server.datastore.DeleteRegistrationEntryRequest) | [DeleteRegistrationEntryResponse](#spire.server.datastore.DeleteRegistrationEntryRequest) | Deletes a specific registration entry |
| CreateJoinToken | [CreateJoinTokenRequest](#spire.server.datastore.CreateJoinTokenRequest) | [CreateJoinTokenResponse](#spire.server.datastore.CreateJoinTokenRequest) | Creates a join token |
| FetchJoinToken | [FetchJoinTokenRequest](#spire.server.datastore.FetchJoinTokenRequest) | [FetchJoinTokenResponse](#spire.server.datastore.FetchJoinTokenRequest) | Fetches a specific join token |
| DeleteJoinToken | [DeleteJoinTokenRequest](#spire.server.datastore.DeleteJoinTokenRequest) | [DeleteJoinTokenResponse](#spire.server.datastore.DeleteJoinTokenRequest) | Delete a specific join token |
| PruneJoinTokens | [PruneJoinTokensRequest](#spire.server.datastore.PruneJoinTokensRequest) | [PruneJoinTokensResponse](#spire.server.datastore.PruneJoinTokensRequest) | Prunes all join tokens that expire before the specified timestamp |
| Configure | [spire.common.plugin.ConfigureRequest](#spire.common.plugin.ConfigureRequest) | [spire.common.plugin.ConfigureResponse](#spire.common.plugin.ConfigureRequest) | Applies the plugin configuration |
| GetPluginInfo | [spire.common.plugin.GetPluginInfoRequest](#spire.common.plugin.GetPluginInfoRequest) | [spire.common.plugin.GetPluginInfoResponse](#spire.common.plugin.GetPluginInfoRequest) | Returns the version and related metadata of the installed plugin |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

