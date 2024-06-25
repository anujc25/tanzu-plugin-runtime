// Package testing exports a GraphQL Mock Server that facilitates
// the testing of client.
//
// NOTE: A portion of this file is adapted from github.com/getoutreach/goql
//
//nolint:goheader
package testing

import "context"

// Operation Type Constants
const (
	opQuery = iota + 1
	opMutation
	opSubscription
)

// Operation is a general type that encompasses the Operation type and Response which
// is of the same type, but with data.
type Operation struct {
	// opType denotes whether the operation is a query or a mutation, using the opQuery
	// and opMutation constants. This is unexported as it is set by the *Server.RegisterQuery
	// and *Server.RegisterMutation functions, respectively.
	opType int

	// Identifier helps identify the operation in a request when coming through the Server.
	// For example, if your operation looks like this:
	//
	//	query {
	//		myOperation(foo: $foo) {
	//			fieldOne
	//			fieldTwo
	//		}
	//	}
	//
	// Then this field should be set to myOperation. It can also be more specific, a simple
	// strings.Contains check occurs to match operations. A more specific example of a
	// valid Identifier for the same operation given above would be myOperation(foo: $foo).
	Identifier string

	// Variables represents the map of variables that should be passed along with the
	// operation whenever it is invoked on the Server.
	Variables map[string]interface{}

	// Response represents the response that should be returned whenever the server makes
	// a match on Operation.opType, Operation.Name, and Operation.Variables.
	Response interface{}

	// EventGenerator should generate mock events
	EventGenerator EventGenerator
}

// OperationError is a special type that brings together the properties that a
// response error can include.
type OperationError struct {
	// Identifier helps identify the operation error in a request when coming through the Server.
	// For example, if your operation looks like this:
	//
	//	error {
	//		myOperation(foo: $foo) {
	//			fieldOne
	//			fieldTwo
	//		}
	//	}
	//
	// Then this field should be set to myOperation. It can also be more specific, a simple
	// strings.Contains check occurs to match operations. A more specific example of a
	// valid Identifier for the same operation given above would be myOperation(foo: $foo).
	Identifier string

	// Status represents the http status code that should be returned in the response
	// whenever the server makes a match on OperationError.Identifier
	Status int

	// Error represents the error that should be returned in the response whenever
	// the server makes a match on OperationError.Identifier
	Error error

	// Extensions represents the object that should be returned in the response
	// as part of the api error whenever the server makes a match on OperationError.Extensions
	Extensions interface{}
}

// EventGenerator should implement a eventData generator for testing and
// send mock event response to the `eventData` channel. To suggest end of
// the event responses from server side, you can close the eventData channel
type EventGenerator func(ctx context.Context, eventData chan<- Response)
