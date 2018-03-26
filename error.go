package rpc_error

import "github.com/gorilla/rpc/v2/json2"

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
	EParse      = json2.E_PARSE
	EInvalidReq = json2.E_INVALID_REQ
	ENoMethod   = json2.E_NO_METHOD
	EBadParams  = json2.E_BAD_PARAMS
	EInternal   = json2.E_INTERNAL
	EServer     = json2.E_SERVER

	EUnauthorised   json2.ErrorCode = -32001
	EForbidden      json2.ErrorCode = -32002
	EDuplicate      json2.ErrorCode = -32003
	EBadCredentials json2.ErrorCode = -32004
	ENotFound       json2.ErrorCode = -32005
)

var Messages = map[json2.ErrorCode]string{
	EParse:      "Parse error",
	EInvalidReq: "Invalid Request",
	ENoMethod:   "Method not found",
	EBadParams:  "Invalid params",
	EInternal:   "Internal error",
	EServer:     "Server error",

	EUnauthorised:   "Unauthorized",
	EForbidden:      "Forbidden",
	EDuplicate:      "Duplicate value",
	EBadCredentials: "Bad credentials",
	ENotFound:       "Not found",
}

func New(code json2.ErrorCode, data interface{}) *json2.Error {
	return &json2.Error{
		Code:    code,
		Message: Messages[code],
		Data:    data,
	}
}
