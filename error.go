package rpc_error

import (
	"fmt"
	"encoding/json"
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
	CodeStdParseError     Code = -32700
	CodeStdInvalidRequest Code = -32600
	CodeStdMethodNotFound Code = -32601
	CodeStdInvalidParams  Code = -32602
	CodeStdInternalError  Code = -32603

	CodeUnexpectedError Code = -32000
	CodeUnauthorized    Code = -32001
	CodeForbidden       Code = -32002
	CodeDuplicateValue  Code = -32003
	CodeBadCredentials  Code = -32004
	CodeNotFound        Code = -32005
)

var messages = map[Code]string{
	CodeStdParseError:     "Parse errror",
	CodeStdInvalidRequest: "Invalid Request",
	CodeStdMethodNotFound: "Method not found",
	CodeStdInvalidParams:  "Invalid params",
	CodeStdInternalError:  "Internal error",

	CodeUnexpectedError: "Unexpected error",
	CodeUnauthorized:    "Unauthorized",
	CodeForbidden:       "Forbidden",
	CodeDuplicateValue:  "Duplicate value",
	CodeBadCredentials:  "Bad credentials",
	CodeNotFound:        "Not found",
}

type Error struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (e *Error) String() string {
	var stringData string
	var ok bool
	if stringData, ok = e.Data.(string); !ok {
		b, err := json.Marshal(e.Data)
		if err != nil {
			panic(err)
		}
		stringData = string(b)
	}
	return fmt.Sprintf("%d: %s (%s)", e.Code, e.Message, stringData)
}

func (e *Error) Error() string {
	return e.String()
}

func New(code Code, data interface{}) *Error {
	return &Error{
		Code:    code,
		Message: messages[code],
		Data:    data,
	}
}

func Unexpected(data interface{}) *Error {
	return &Error{
		Code:    CodeUnexpectedError,
		Message: messages[CodeUnexpectedError],
		Data:    data,
	}
}
