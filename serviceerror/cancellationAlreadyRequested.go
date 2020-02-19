// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/temporal-proto/errordetails"
)

type (
	// CancellationAlreadyRequested represents cancellation already requested error.
	CancellationAlreadyRequested struct {
		Message        string
		st             *status.Status
	}
)

// NewCancellationAlreadyRequested returns new CancellationAlreadyRequested error.
func NewCancellationAlreadyRequested(message string) *CancellationAlreadyRequested {
	return &CancellationAlreadyRequested{
		Message:        message,
	}
}

// Error returns string message.
func (e *CancellationAlreadyRequested) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *CancellationAlreadyRequested) GRPCStatus() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&errordetails.CancellationAlreadyRequestedFailure{},
	)
	return st
}

func newCancellationAlreadyRequested(st *status.Status) *CancellationAlreadyRequested {
	return &CancellationAlreadyRequested{
		Message: st.Message(),
		st:      st,
	}
}