// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cherami

import cherami "github.com/uber/cherami-client-go/client/cherami"
import gocherami "github.com/uber/cherami-thrift/.generated/go/cherami"
import mock "github.com/stretchr/testify/mock"

// Delivery is an autogenerated mock type for the Delivery type
type Delivery struct {
	mock.Mock
}

// Ack provides a mock function with given fields:
func (_m *Delivery) Ack() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeliveryToken provides a mock function with given fields:
func (_m *Delivery) GetDeliveryToken() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetMessage provides a mock function with given fields:
func (_m *Delivery) GetMessage() *gocherami.ConsumerMessage {
	ret := _m.Called()

	var r0 *gocherami.ConsumerMessage
	if rf, ok := ret.Get(0).(func() *gocherami.ConsumerMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gocherami.ConsumerMessage)
		}
	}

	return r0
}

// Nack provides a mock function with given fields:
func (_m *Delivery) Nack() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyChecksum provides a mock function with given fields:
func (_m *Delivery) VerifyChecksum() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

var _ cherami.Delivery = (*Delivery)(nil)