package rpc_error

import (
	"github.com/gorilla/rpc/v2/json2"
)

// From: http://www.jsonrpc.org/specification#error_object
//
// code
//   A Number that indicates the error type that occurred.
//   This MUST be an integer.
// message
//   A String providing a short description of the error.
//   The message SHOULD be limited to a concise single sentence.
// data
//   A Primitive or Structured value that contains additional information about the error.
//   This may be omitted.
//   The value of this member is defined by the Server (e.g. detailed error information, nested errors etc.).
//
// code 	        message 	        meaning
// -32700 	        Parse error 	    Invalid JSON was received by the server.
//                                      An error occurred on the server while parsing the JSON text.
// -32600 	        Invalid Request 	The JSON sent is not a valid Request object.
// -32601           Method not found 	The method does not exist / is not available.
// -32602 	        Invalid params 	    Invalid method parameter(s).
// -32603 	        Internal error 	    Internal JSON-RPC error.
// -32000 to -32099 Server error 	    Reserved for implementation-defined server-errors.

type Code int

const (
	E_UNAUTHORISED    json2.ErrorCode = -32001
	E_FORBIDDEN       json2.ErrorCode = -32002
	E_DUPLICATE       json2.ErrorCode = -32003
	E_BAD_CREDENTIALS json2.ErrorCode = -32004
	E_NOT_FOUND       json2.ErrorCode = -32005
)

var messages = map[json2.ErrorCode]string{
	json2.E_PARSE:       "Parse error",
	json2.E_INVALID_REQ: "Invalid Request",
	json2.E_NO_METHOD:   "Method not found",
	json2.E_BAD_PARAMS:  "Invalid params",
	json2.E_INTERNAL:    "Internal error",
	json2.E_SERVER:      "Server error",

	E_UNAUTHORISED:    "Unauthorized",
	E_FORBIDDEN:       "Forbidden",
	E_DUPLICATE:       "Duplicate value",
	E_BAD_CREDENTIALS: "Bad credentials",
	E_NOT_FOUND:       "Not found",
}

func New(code json2.ErrorCode, data interface{}) *json2.Error {
	return &json2.Error{
		Code:    code,
		Message: messages[code],
		Data:    data,
	}
}

func Unexpected(data interface{}) *json2.Error {
	return &json2.Error{
		Code:    json2.E_SERVER,
		Message: messages[json2.E_SERVER],
		Data:    data,
	}
}
