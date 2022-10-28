extern crate serde;
extern crate serde_json;
extern crate derive_builder;

use serde::{Serialize, Deserialize};
use std::collections::HashMap;
use derive_builder::Builder;
/// JSONRPCRequestIdAsString
///
/// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
///
pub type JSONRPCRequestIdAsString = String;
/// JSONRPCRequestIdAsNumber
///
/// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
///
pub type JSONRPCRequestIdAsNumber = f64;
/// JSONRPCRequestIdAsNull
///
/// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
///
pub type JSONRPCRequestIdAsNull = serde_json::Value;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum JSONRPCRequestId {
    JSONRPCRequestIdAsString(JSONRPCRequestIdAsString),
    JSONRPCRequestIdAsNumber(JSONRPCRequestIdAsNumber),
    JSONRPCRequestIdAsNull(JSONRPCRequestIdAsNull),
}
/// JSONRPCVersion
///
/// A String specifying the version of the JSON-RPC protocol. MUST be exactly "2.0".
///
#[derive(Serialize, Deserialize, Clone, Debug, Eq, PartialEq)]
pub enum JSONRPCVersion {
    #[serde(rename = "2.0")]
    TwoZero,
}
/// JSONRPCMethod
///
/// A String containing the name of the method to be invoked. Method names that begin with the word rpc followed by a period character (U+002E or ASCII 46) are reserved for rpc-internal methods and extensions and MUST NOT be used for anything else.
///
pub type JSONRPCMethod = String;
/// JSONRPCParamValue
///
/// Parameters literal value as provided to the invocation of the method.
///
pub type JSONRPCParamValue = serde_json::Value;
/// JSONRPCParamsByPosition
///
/// by-position: params MUST be an Array, containing the values in the Server expected order.
///
pub type JSONRPCParamsByPosition = Vec<JSONRPCParamValue>;
/// JSONRPCParamsByName
///
/// by-name: params MUST be an Object, with member names that match the Server expected parameter names. The absence of expected names MAY result in an error being generated. The names MUST match exactly, including case, to the method's expected parameters.
///
pub type JSONRPCParamsByName = HashMap<String, JSONRPCParamValue>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum JSONRPCParams {
    JSONRPCParamsByPosition(JSONRPCParamsByPosition),
    JSONRPCParamsByName(JSONRPCParamsByName),
}
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct JSONRPCRequest {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub id: Option<JSONRPCRequestId>,
    pub jsonrpc: JSONRPCVersion,
    pub method: JSONRPCMethod,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub params: Option<JSONRPCParams>,
}
/// JSONRPCErrorCode
///
/// A Number that indicates the error type that occurred. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. Any code within this range, but not defined explicitly below is reserved for future use. The error codes are nearly the same as those suggested for XML-RPC at the [following url](http://xmlrpc-epi.sourceforge.net/specs/rfc.fault_codes.php)
///
pub type JSONRPCErrorCode = i64;
/// JSONRPCErrorMessage
///
/// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
///
pub type JSONRPCErrorMessage = String;
/// JSONRPCErrorData
///
/// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
///
pub type JSONRPCErrorData = serde_json::Value;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct JSONRPCError {
    #[serde(skip_serializing_if = "Option::is_none")]
    pub code: Option<JSONRPCErrorCode>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub message: Option<JSONRPCErrorMessage>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<JSONRPCErrorData>,
}
/// JSONRPCResponseResult
///
/// This member is REQUIRED on success. This member MUST NOT exist if there was an error invoking the method. The value of this member is determined by the method invoked on the Server.
///
pub type JSONRPCResponseResult = serde_json::Value;
/// JSONRPCResponse
///
/// When a rpc call is made, the Server MUST reply with a Response, except for in the case of Notifications. The Response is expressed as a single JSON Object,
///
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq, Builder, Default)]
#[builder(setter(strip_option), default)]
#[serde(default)]
pub struct JSONRPCResponse {
    pub id: JSONRPCRequestId,
    pub jsonrpc: JSONRPCVersion,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub error: Option<JSONRPCError>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub result: Option<JSONRPCResponseResult>,
}
/// JSONRPCBatchRequest
///
/// To send several Request objects at the same time, the Client MAY send an Array filled with Request objects. The Client SHOULD match contexts between the set of Request objects and the resulting set of Response objects based on the id member within each Object.
///
pub type JSONRPCBatchRequest = Vec<JSONRPCRequest>;
/// JSONRPCBatchResponse
///
/// The Server should respond with an Array containing the corresponding Response objects, after all of the batch Request objects have been processed. A Response object SHOULD exist for each Request object, except that there SHOULD NOT be any Response objects for notifications. The Server MAY process a batch rpc call as a set of concurrent tasks, processing them in any order and with any width of parallelism.
 The Response objects being returned from a batch call MAY be returned in any order within the Array. 
 If the batch rpc call itself fails to be recognized as an valid JSON or as an Array with at least one value, the response from the Server MUST be a single Response object. If there are no Response objects contained within the Response array as it is to be sent to the client, the server MUST NOT return an empty Array and should return nothing at all.
///
pub type JSONRPCBatchResponse = Vec<JSONRPCResponse>;
#[derive(Serialize, Deserialize, Clone, Debug, PartialEq)]
#[serde(untagged)]
pub enum JSONRPC {
    JSONRPCRequest(JSONRPCRequest),
    JSONRPCResponse(JSONRPCResponse),
    JSONRPCBatchRequest(JSONRPCBatchRequest),
    JSONRPCBatchResponse(JSONRPCBatchResponse),
}