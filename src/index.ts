export const jsonrpc = {"$schema":"https://meta.json-schema.tools","$id":"https://meta.jsonrpc.net/","title":"JSONRPC","description":"Either a (possibly batch of) JSON-RPC request(s) and response(s)","oneOf":[{"$ref":"#/definitions/JSONRPCRequest"},{"$ref":"#/definitions/JSONRPCResponse"},{"$ref":"#/definitions/JSONRPCBatchRequest"},{"$ref":"#/definitions/JSONRPCBatchResponse"}],"definitions":{"JSONRPCRequestIdAsString":{"title":"JSONRPCRequestIdAsString","description":"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.","type":"string"},"JSONRPCRequestIdAsNumber":{"title":"JSONRPCRequestIdAsNumber","description":"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.","type":"number"},"JSONRPCRequestIdAsNull":{"title":"JSONRPCRequestIdAsNull","description":"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.","type":"null"},"JSONRPCRequestId":{"title":"JSONRPCRequestId","description":"If it is not included it is assumed to be a notification. The Server MUST reply with the same value in the Response object if included. This member is used to correlate the context between the two objects.","oneOf":[{"$ref":"#/definitions/JSONRPCRequestIdAsString"},{"$ref":"#/definitions/JSONRPCRequestIdAsNumber"},{"$ref":"#/definitions/JSONRPCRequestIdAsNull"}]},"JSONRPCVersion":{"title":"JSONRPCVersion","description":"A String specifying the version of the JSON-RPC protocol. MUST be exactly \"2.0\".","type":"string","const":"2.0"},"JSONRPCMethod":{"title":"JSONRPCMethod","description":"A String containing the name of the method to be invoked. Method names that begin with the word rpc followed by a period character (U+002E or ASCII 46) are reserved for rpc-internal methods and extensions and MUST NOT be used for anything else.","type":"string"},"JSONRPCParamValue":{"title":"JSONRPCParamValue","description":"Parameters literal value as provided to the invocation of the method."},"JSONRPCParamsByPosition":{"title":"JSONRPCParamsByPosition","description":"by-position: params MUST be an Array, containing the values in the Server expected order.","type":"array","items":{"$ref":"#/definitions/JSONRPCParamValue"}},"JSONRPCParamsByName":{"title":"JSONRPCParamsByName","description":"by-name: params MUST be an Object, with member names that match the Server expected parameter names. The absence of expected names MAY result in an error being generated. The names MUST match exactly, including case, to the method's expected parameters.","type":"object","additionalProperties":{"$ref":"#/definitions/JSONRPCParamValue"}},"JSONRPCParams":{"title":"JSONRPCParams","description":"A Structured value that holds the parameter values to be used during the invocation of the method. This member MAY be omitted.","oneOf":[{"$ref":"#/definitions/JSONRPCParamsByPosition"},{"$ref":"#/definitions/JSONRPCParamsByName"}]},"JSONRPCRequest":{"title":"JSONRPCRequest","description":"A rpc call is represented by sending a Request object to a Server.","type":"object","required":["jsonrpc","method"],"properties":{"id":{"$ref":"#/definitions/JSONRPCRequestId"},"jsonrpc":{"$ref":"#/definitions/JSONRPCVersion"},"method":{"$ref":"#/definitions/JSONRPCMethod"},"params":{"$ref":"#/definitions/JSONRPCParams"}}},"JSONRPCErrorCode":{"title":"JSONRPCErrorCode","description":"A Number that indicates the error type that occurred. The error codes from and including -32768 to -32000 are reserved for pre-defined errors. Any code within this range, but not defined explicitly below is reserved for future use. The error codes are nearly the same as those suggested for XML-RPC at the [following url](http://xmlrpc-epi.sourceforge.net/specs/rfc.fault_codes.php)","type":"integer"},"JSONRPCErrorMessage":{"title":"JSONRPCErrorMessage","description":"A String providing a short description of the error. The message SHOULD be limited to a concise single sentence.","type":"string"},"JSONRPCErrorData":{"title":"JSONRPCErrorData","description":"A Primitive or Structured value that contains additional information about the error. This may be omitted. The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.)."},"JSONRPCError":{"title":"JSONRPCError","description":"When a rpc call encounters an error, the Response Object MUST contain the error member with a value that is a Object.","type":"object","properties":{"code":{"$ref":"#/definitions/JSONRPCErrorCode"},"message":{"$ref":"#/definitions/JSONRPCErrorMessage"},"data":{"$ref":"#/definitions/JSONRPCErrorData"}}},"JSONRPCResponseResult":{"title":"JSONRPCResponseResult","description":"This member is REQUIRED on success. This member MUST NOT exist if there was an error invoking the method. The value of this member is determined by the method invoked on the Server."},"JSONRPCResponse":{"title":"JSONRPCResponse","description":"When a rpc call is made, the Server MUST reply with a Response, except for in the case of Notifications. The Response is expressed as a single JSON Object,","type":"object","required":["id","jsonrpc"],"properties":{"id":{"$ref":"#/definitions/JSONRPCRequestId"},"jsonrpc":{"$ref":"#/definitions/JSONRPCVersion"},"error":{"$ref":"#/definitions/JSONRPCError"},"result":{"$ref":"#/definitions/JSONRPCResponseResult"}}},"JSONRPCBatchRequest":{"title":"JSONRPCBatchRequest","description":"To send several Request objects at the same time, the Client MAY send an Array filled with Request objects. The Client SHOULD match contexts between the set of Request objects and the resulting set of Response objects based on the id member within each Object.","type":"array","items":{"$ref":"#/definitions/JSONRPCRequest"}},"JSONRPCBatchResponse":{"title":"JSONRPCBatchResponse","description":"The Server should respond with an Array containing the corresponding Response objects, after all of the batch Request objects have been processed. A Response object SHOULD exist for each Request object, except that there SHOULD NOT be any Response objects for notifications. The Server MAY process a batch rpc call as a set of concurrent tasks, processing them in any order and with any width of parallelism.\n The Response objects being returned from a batch call MAY be returned in any order within the Array. \n If the batch rpc call itself fails to be recognized as an valid JSON or as an Array with at least one value, the response from the Server MUST be a single Response object. If there are no Response objects contained within the Response array as it is to be sent to the client, the server MUST NOT return an empty Array and should return nothing at all.","type":"array","items":{"$ref":"#/definitions/JSONRPCResponse"}}}};
export default jsonrpc