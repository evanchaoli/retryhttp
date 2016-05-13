// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"

	"github.com/concourse/retryhttp/roundtripper"
)

type FakeRoundTripper struct {
	RoundTripStub        func(request *http.Request) (*http.Response, error)
	roundTripMutex       sync.RWMutex
	roundTripArgsForCall []struct {
		request *http.Request
	}
	roundTripReturns struct {
		result1 *http.Response
		result2 error
	}
}

func (fake *FakeRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	fake.roundTripMutex.Lock()
	fake.roundTripArgsForCall = append(fake.roundTripArgsForCall, struct {
		request *http.Request
	}{request})
	fake.roundTripMutex.Unlock()
	if fake.RoundTripStub != nil {
		return fake.RoundTripStub(request)
	} else {
		return fake.roundTripReturns.result1, fake.roundTripReturns.result2
	}
}

func (fake *FakeRoundTripper) RoundTripCallCount() int {
	fake.roundTripMutex.RLock()
	defer fake.roundTripMutex.RUnlock()
	return len(fake.roundTripArgsForCall)
}

func (fake *FakeRoundTripper) RoundTripArgsForCall(i int) *http.Request {
	fake.roundTripMutex.RLock()
	defer fake.roundTripMutex.RUnlock()
	return fake.roundTripArgsForCall[i].request
}

func (fake *FakeRoundTripper) RoundTripReturns(result1 *http.Response, result2 error) {
	fake.RoundTripStub = nil
	fake.roundTripReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

var _ roundtripper.RoundTripper = new(FakeRoundTripper)