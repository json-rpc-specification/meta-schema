package jsonrpc


import "encoding/json"
import "errors"
// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
type JSONRPCRequestIdAsString string
// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
type JSONRPCRequestIdAsNumber float64
// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
type JSONRPCRequestIdAsNull interface{}
// If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.
type JSONRPCRequestId struct {
	JSONRPCRequestIdAsString *JSONRPCRequestIdAsString
	JSONRPCRequestIdAsNumber *JSONRPCRequestIdAsNumber
	JSONRPCRequestIdAsNull   *JSONRPCRequestIdAsNull
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONRPCRequestId) UnmarshalJSON(bytes []byte) error {
	var myJSONRPCRequestIdAsString JSONRPCRequestIdAsString
	if err := json.Unmarshal(bytes, &myJSONRPCRequestIdAsString); err == nil {
		o.JSONRPCRequestIdAsString = &myJSONRPCRequestIdAsString
		return nil
	}
	var myJSONRPCRequestIdAsNumber JSONRPCRequestIdAsNumber
	if err := json.Unmarshal(bytes, &myJSONRPCRequestIdAsNumber); err == nil {
		o.JSONRPCRequestIdAsNumber = &myJSONRPCRequestIdAsNumber
		return nil
	}
	var myJSONRPCRequestIdAsNull JSONRPCRequestIdAsNull
	if err := json.Unmarshal(bytes, &myJSONRPCRequestIdAsNull); err == nil {
		o.JSONRPCRequestIdAsNull = &myJSONRPCRequestIdAsNull
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONRPCRequestId) MarshalJSON() ([]byte, error) {
	if o.JSONRPCRequestIdAsString != nil {
		return json.Marshal(o.JSONRPCRequestIdAsString)
	}
	if o.JSONRPCRequestIdAsNumber != nil {
		return json.Marshal(o.JSONRPCRequestIdAsNumber)
	}
	if o.JSONRPCRequestIdAsNull != nil {
		return json.Marshal(o.JSONRPCRequestIdAsNull)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A String specifying the version of the JSON-RPC protocol. MUST be exactly "2.0".
type JSONRPCVersion string
// A String containing the name of the method to be invoked. Method names that begin with the word rpc followed by a period character (U+002E or ASCII 46) are reserved for rpc-internal methods and extensions and MUST NOT be used for anything else.
type JSONRPCMethod string
// Parameters literal value as provided to the invocation of the method.
type JSONRPCParamValue interface{}
// by-position: params MUST be an Array, containing the values in the Server expected order.
type JSONRPCParamsByPosition []JSONRPCParamValue
// by-name: params MUST be an Object, with member names that match the Server expected parameter names. The absence of expected names MAY result in an error being generated. The names MUST match exactly, including case, to the method's expected parameters.
type JSONRPCParamsByName map[string]interface{}
// A Structured value that holds the parameter values to be used during the invocation of the method. This member MAY be omitted.
type JSONRPCParams struct {
	JSONRPCParamsByPosition *JSONRPCParamsByPosition
	JSONRPCParamsByName     *JSONRPCParamsByName
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONRPCParams) UnmarshalJSON(bytes []byte) error {
	var myJSONRPCParamsByPosition JSONRPCParamsByPosition
	if err := json.Unmarshal(bytes, &myJSONRPCParamsByPosition); err == nil {
		o.JSONRPCParamsByPosition = &myJSONRPCParamsByPosition
		return nil
	}
	var myJSONRPCParamsByName JSONRPCParamsByName
	if err := json.Unmarshal(bytes, &myJSONRPCParamsByName); err == nil {
		o.JSONRPCParamsByName = &myJSONRPCParamsByName
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONRPCParams) MarshalJSON() ([]byte, error) {
	if o.JSONRPCParamsByPosition != nil {
		return json.Marshal(o.JSONRPCParamsByPosition)
	}
	if o.JSONRPCParamsByName != nil {
		return json.Marshal(o.JSONRPCParamsByName)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}
// A rpc call is represented by sending a Request object to a Server.
type JSONRPCRequest struct {
	Id      *JSONRPCRequestId `json:"id,omitempty"`
	Jsonrpc *JSONRPCVersion   `json:"jsonrpc"`
	Method  *JSONRPCMethod    `json:"method"`
	Params  *JSONRPCParams    `json:"params,omitempty"`
}
// A Number that indicates the error type that occurred. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. Any code within this range, but not defined explicitly below is reserved for future use. The error codes are nearly the same as those suggested for XML-RPC at the [following url](http://xmlrpc-epi.sourceforge.net/specs/rfc.fault_codes.php)
type JSONRPCErrorCode int64
// A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.
type JSONRPCErrorMessage string
// A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
type JSONRPCErrorData interface{}
// When a rpc call encounters an error, the Response Object MUST contain the error member with a value that is a Object.
type JSONRPCError struct {
	Code    *JSONRPCErrorCode    `json:"code,omitempty"`
	Message *JSONRPCErrorMessage `json:"message,omitempty"`
	Data    *JSONRPCErrorData    `json:"data,omitempty"`
}
// This member is REQUIRED on success. This member MUST NOT exist if there was an error invoking the method. The value of this member is determined by the method invoked on the Server.
type JSONRPCResponseResult interface{}
// When a rpc call is made, the Server MUST reply with a Response, except for in the case of Notifications. The Response is expressed as a single JSON Object,
type JSONRPCResponse struct {
	Id      *JSONRPCRequestId      `json:"id"`
	Jsonrpc *JSONRPCVersion        `json:"jsonrpc"`
	Error   *JSONRPCError          `json:"error,omitempty"`
	Result  *JSONRPCResponseResult `json:"result,omitempty"`
}
// To send several Request objects at the same time, the Client MAY send an Array filled with Request objects. The Client SHOULD match contexts between the set of Request objects and the resulting set of Response objects based on the id member within each Object.
type JSONRPCBatchRequest []JSONRPCRequest
// The Server should respond with an Array containing the corresponding Response objects, after all of the batch Request objects have been processed. A Response object SHOULD exist for each Request object, except that there SHOULD NOT be any Response objects for notifications. The Server MAY process a batch rpc call as a set of concurrent tasks, processing them in any order and with any width of parallelism.
 The Response objects being returned from a batch call MAY be returned in any order within the Array. 
 If the batch rpc call itself fails to be recognized as an valid JSON or as an Array with at least one value, the response from the Server MUST be a single Response object. If there are no Response objects contained within the Response array as it is to be sent to the client, the server MUST NOT return an empty Array and should return nothing at all.
type JSONRPCBatchResponse []JSONRPCResponse
// Either a (possibly batch of) JSON-RPC request(s) and response(s)
type JSONRPC struct {
	JSONRPCRequest       *JSONRPCRequest
	JSONRPCResponse      *JSONRPCResponse
	JSONRPCBatchRequest  *JSONRPCBatchRequest
	JSONRPCBatchResponse *JSONRPCBatchResponse
}
// UnmarshalJSON implements the json Unmarshaler interface.
// This implementation DOES NOT assert that ONE AND ONLY ONE
// of the simple properties is satisfied; it lazily uses the first one that is satisfied.
// Ergo, it will not return an error if more than one property is valid.
func (o *JSONRPC) UnmarshalJSON(bytes []byte) error {
	var myJSONRPCRequest JSONRPCRequest
	if err := json.Unmarshal(bytes, &myJSONRPCRequest); err == nil {
		o.JSONRPCRequest = &myJSONRPCRequest
		return nil
	}
	var myJSONRPCResponse JSONRPCResponse
	if err := json.Unmarshal(bytes, &myJSONRPCResponse); err == nil {
		o.JSONRPCResponse = &myJSONRPCResponse
		return nil
	}
	var myJSONRPCBatchRequest JSONRPCBatchRequest
	if err := json.Unmarshal(bytes, &myJSONRPCBatchRequest); err == nil {
		o.JSONRPCBatchRequest = &myJSONRPCBatchRequest
		return nil
	}
	var myJSONRPCBatchResponse JSONRPCBatchResponse
	if err := json.Unmarshal(bytes, &myJSONRPCBatchResponse); err == nil {
		o.JSONRPCBatchResponse = &myJSONRPCBatchResponse
		return nil
	}
	return errors.New("failed to unmarshal one of the object properties")
}
func (o JSONRPC) MarshalJSON() ([]byte, error) {
	if o.JSONRPCRequest != nil {
		return json.Marshal(o.JSONRPCRequest)
	}
	if o.JSONRPCResponse != nil {
		return json.Marshal(o.JSONRPCResponse)
	}
	if o.JSONRPCBatchRequest != nil {
		return json.Marshal(o.JSONRPCBatchRequest)
	}
	if o.JSONRPCBatchResponse != nil {
		return json.Marshal(o.JSONRPCBatchResponse)
	}
	return nil, errors.New("failed to marshal any one of the object properties")
}

const RawJsonrpc = "{\"$schema\":\"https://meta.json-schema.tools\",\"$id\":\"https://meta.jsonrpc.net/\",\"title\":\"JSONRPC\",\"description\":\"Either a (possibly batch of) JSON-RPC request(s) and response(s)\",\"oneOf\":[{\"$ref\":\"#/definitions/JSONRPCRequest\"},{\"$ref\":\"#/definitions/JSONRPCResponse\"},{\"$ref\":\"#/definitions/JSONRPCBatchRequest\"},{\"$ref\":\"#/definitions/JSONRPCBatchResponse\"}],\"definitions\":{\"JSONRPCRequestIdAsString\":{\"title\":\"JSONRPCRequestIdAsString\",\"description\":\"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.\",\"type\":\"string\"},\"JSONRPCRequestIdAsNumber\":{\"title\":\"JSONRPCRequestIdAsNumber\",\"description\":\"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.\",\"type\":\"number\"},\"JSONRPCRequestIdAsNull\":{\"title\":\"JSONRPCRequestIdAsNull\",\"description\":\"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.\",\"type\":\"null\"},\"JSONRPCRequestId\":{\"title\":\"JSONRPCRequestId\",\"description\":\"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.\",\"oneOf\":[{\"$ref\":\"#/definitions/JSONRPCRequestIdAsString\"},{\"$ref\":\"#/definitions/JSONRPCRequestIdAsNumber\"},{\"$ref\":\"#/definitions/JSONRPCRequestIdAsNull\"}]},\"JSONRPCVersion\":{\"title\":\"JSONRPCVersion\",\"description\":\"A String specifying the version of the JSON-RPC protocol. MUST be exactly \\"2.0\\".\",\"type\":\"string\",\"const\":\"2.0\"},\"JSONRPCMethod\":{\"title\":\"JSONRPCMethod\",\"description\":\"A String containing the name of the method to be invoked. Method names that begin with the word rpc followed by a period character (U+002E or ASCII 46) are reserved for rpc-internal methods and extensions and MUST NOT be used for anything else.\",\"type\":\"string\"},\"JSONRPCParamValue\":{\"title\":\"JSONRPCParamValue\",\"description\":\"Parameters literal value as provided to the invocation of the method.\"},\"JSONRPCParamsByPosition\":{\"title\":\"JSONRPCParamsByPosition\",\"description\":\"by-position: params MUST be an Array, containing the values in the Server expected order.\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/JSONRPCParamValue\"}},\"JSONRPCParamsByName\":{\"title\":\"JSONRPCParamsByName\",\"description\":\"by-name: params MUST be an Object, with member names that match the Server expected parameter names. The absence of expected names MAY result in an error being generated. The names MUST match exactly, including case, to the method's expected parameters.\",\"type\":\"object\",\"additionalProperties\":{\"$ref\":\"#/definitions/JSONRPCParamValue\"}},\"JSONRPCParams\":{\"title\":\"JSONRPCParams\",\"description\":\"A Structured value that holds the parameter values to be used during the invocation of the method. This member MAY be omitted.\",\"oneOf\":[{\"$ref\":\"#/definitions/JSONRPCParamsByPosition\"},{\"$ref\":\"#/definitions/JSONRPCParamsByName\"}]},\"JSONRPCRequest\":{\"title\":\"JSONRPCRequest\",\"description\":\"A rpc call is represented by sending a Request object to a Server.\",\"type\":\"object\",\"required\":[\"jsonrpc\",\"method\"],\"properties\":{\"id\":{\"$ref\":\"#/definitions/JSONRPCRequestId\"},\"jsonrpc\":{\"$ref\":\"#/definitions/JSONRPCVersion\"},\"method\":{\"$ref\":\"#/definitions/JSONRPCMethod\"},\"params\":{\"$ref\":\"#/definitions/JSONRPCParams\"}}},\"JSONRPCErrorCode\":{\"title\":\"JSONRPCErrorCode\",\"description\":\"A Number that indicates the error type that occurred. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. Any code within this range, but not defined explicitly below is reserved for future use. The error codes are nearly the same as those suggested for XML-RPC at the [following url](http://xmlrpc-epi.sourceforge.net/specs/rfc.fault_codes.php)\",\"type\":\"integer\"},\"JSONRPCErrorMessage\":{\"title\":\"JSONRPCErrorMessage\",\"description\":\"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.\",\"type\":\"string\"},\"JSONRPCErrorData\":{\"title\":\"JSONRPCErrorData\",\"description\":\"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).\"},\"JSONRPCError\":{\"title\":\"JSONRPCError\",\"description\":\"When a rpc call encounters an error, the Response Object MUST contain the error member with a value that is a Object.\",\"type\":\"object\",\"properties\":{\"code\":{\"$ref\":\"#/definitions/JSONRPCErrorCode\"},\"message\":{\"$ref\":\"#/definitions/JSONRPCErrorMessage\"},\"data\":{\"$ref\":\"#/definitions/JSONRPCErrorData\"}}},\"JSONRPCResponseResult\":{\"title\":\"JSONRPCResponseResult\",\"description\":\"This member is REQUIRED on success. This member MUST NOT exist if there was an error invoking the method. The value of this member is determined by the method invoked on the Server.\"},\"JSONRPCResponse\":{\"title\":\"JSONRPCResponse\",\"description\":\"When a rpc call is made, the Server MUST reply with a Response, except for in the case of Notifications. The Response is expressed as a single JSON Object,\",\"type\":\"object\",\"required\":[\"id\",\"jsonrpc\"],\"properties\":{\"id\":{\"$ref\":\"#/definitions/JSONRPCRequestId\"},\"jsonrpc\":{\"$ref\":\"#/definitions/JSONRPCVersion\"},\"error\":{\"$ref\":\"#/definitions/JSONRPCError\"},\"result\":{\"$ref\":\"#/definitions/JSONRPCResponseResult\"}}},\"JSONRPCBatchRequest\":{\"title\":\"JSONRPCBatchRequest\",\"description\":\"To send several Request objects at the same time, the Client MAY send an Array filled with Request objects. The Client SHOULD match contexts between the set of Request objects and the resulting set of Response objects based on the id member within each Object.\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/JSONRPCRequest\"}},\"JSONRPCBatchResponse\":{\"title\":\"JSONRPCBatchResponse\",\"description\":\"The Server should respond with an Array containing the corresponding Response objects, after all of the batch Request objects have been processed. A Response object SHOULD exist for each Request object, except that there SHOULD NOT be any Response objects for notifications. The Server MAY process a batch rpc call as a set of concurrent tasks, processing them in any order and with any width of parallelism.\n The Response objects being returned from a batch call MAY be returned in any order within the Array. \n If the batch rpc call itself fails to be recognized as an valid JSON or as an Array with at least one value, the response from the Server MUST be a single Response object. If there are no Response objects contained within the Response array as it is to be sent to the client, the server MUST NOT return an empty Array and should return nothing at all.\",\"type\":\"array\",\"items\":{\"$ref\":\"#/definitions/JSONRPCResponse\"}}}}"