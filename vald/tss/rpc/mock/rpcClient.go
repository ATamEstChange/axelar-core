// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/axelarnetwork/axelar-core/vald/tss/rpc"
	"github.com/axelarnetwork/axelar-core/x/tss/tofnd"
	"google.golang.org/grpc"
	"sync"
)

// Ensure, that MultiSigClientMock does implement rpc.MultiSigClient.
// If this is not the case, regenerate this file with moq.
var _ rpc.MultiSigClient = &MultiSigClientMock{}

// MultiSigClientMock is a mock implementation of rpc.MultiSigClient.
//
// 	func TestSomethingThatUsesMultiSigClient(t *testing.T) {
//
// 		// make and configure a mocked rpc.MultiSigClient
// 		mockedMultiSigClient := &MultiSigClientMock{
// 			KeyPresenceFunc: func(ctx context.Context, in *tofnd.KeyPresenceRequest, opts ...grpc.CallOption) (*tofnd.KeyPresenceResponse, error) {
// 				panic("mock out the KeyPresence method")
// 			},
// 			KeygenFunc: func(ctx context.Context, in *tofnd.KeygenRequest, opts ...grpc.CallOption) (*tofnd.KeygenResponse, error) {
// 				panic("mock out the Keygen method")
// 			},
// 			SignFunc: func(ctx context.Context, in *tofnd.SignRequest, opts ...grpc.CallOption) (*tofnd.SignResponse, error) {
// 				panic("mock out the Sign method")
// 			},
// 		}
//
// 		// use mockedMultiSigClient in code that requires rpc.MultiSigClient
// 		// and then make assertions.
//
// 	}
type MultiSigClientMock struct {
	// KeyPresenceFunc mocks the KeyPresence method.
	KeyPresenceFunc func(ctx context.Context, in *tofnd.KeyPresenceRequest, opts ...grpc.CallOption) (*tofnd.KeyPresenceResponse, error)

	// KeygenFunc mocks the Keygen method.
	KeygenFunc func(ctx context.Context, in *tofnd.KeygenRequest, opts ...grpc.CallOption) (*tofnd.KeygenResponse, error)

	// SignFunc mocks the Sign method.
	SignFunc func(ctx context.Context, in *tofnd.SignRequest, opts ...grpc.CallOption) (*tofnd.SignResponse, error)

	// calls tracks calls to the methods.
	calls struct {
		// KeyPresence holds details about calls to the KeyPresence method.
		KeyPresence []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *tofnd.KeyPresenceRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Keygen holds details about calls to the Keygen method.
		Keygen []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *tofnd.KeygenRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Sign holds details about calls to the Sign method.
		Sign []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *tofnd.SignRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
	lockKeyPresence sync.RWMutex
	lockKeygen      sync.RWMutex
	lockSign        sync.RWMutex
}

// KeyPresence calls KeyPresenceFunc.
func (mock *MultiSigClientMock) KeyPresence(ctx context.Context, in *tofnd.KeyPresenceRequest, opts ...grpc.CallOption) (*tofnd.KeyPresenceResponse, error) {
	if mock.KeyPresenceFunc == nil {
		panic("MultiSigClientMock.KeyPresenceFunc: method is nil but MultiSigClient.KeyPresence was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *tofnd.KeyPresenceRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockKeyPresence.Lock()
	mock.calls.KeyPresence = append(mock.calls.KeyPresence, callInfo)
	mock.lockKeyPresence.Unlock()
	return mock.KeyPresenceFunc(ctx, in, opts...)
}

// KeyPresenceCalls gets all the calls that were made to KeyPresence.
// Check the length with:
//     len(mockedMultiSigClient.KeyPresenceCalls())
func (mock *MultiSigClientMock) KeyPresenceCalls() []struct {
	Ctx  context.Context
	In   *tofnd.KeyPresenceRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *tofnd.KeyPresenceRequest
		Opts []grpc.CallOption
	}
	mock.lockKeyPresence.RLock()
	calls = mock.calls.KeyPresence
	mock.lockKeyPresence.RUnlock()
	return calls
}

// Keygen calls KeygenFunc.
func (mock *MultiSigClientMock) Keygen(ctx context.Context, in *tofnd.KeygenRequest, opts ...grpc.CallOption) (*tofnd.KeygenResponse, error) {
	if mock.KeygenFunc == nil {
		panic("MultiSigClientMock.KeygenFunc: method is nil but MultiSigClient.Keygen was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *tofnd.KeygenRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockKeygen.Lock()
	mock.calls.Keygen = append(mock.calls.Keygen, callInfo)
	mock.lockKeygen.Unlock()
	return mock.KeygenFunc(ctx, in, opts...)
}

// KeygenCalls gets all the calls that were made to Keygen.
// Check the length with:
//     len(mockedMultiSigClient.KeygenCalls())
func (mock *MultiSigClientMock) KeygenCalls() []struct {
	Ctx  context.Context
	In   *tofnd.KeygenRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *tofnd.KeygenRequest
		Opts []grpc.CallOption
	}
	mock.lockKeygen.RLock()
	calls = mock.calls.Keygen
	mock.lockKeygen.RUnlock()
	return calls
}

// Sign calls SignFunc.
func (mock *MultiSigClientMock) Sign(ctx context.Context, in *tofnd.SignRequest, opts ...grpc.CallOption) (*tofnd.SignResponse, error) {
	if mock.SignFunc == nil {
		panic("MultiSigClientMock.SignFunc: method is nil but MultiSigClient.Sign was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		In   *tofnd.SignRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockSign.Lock()
	mock.calls.Sign = append(mock.calls.Sign, callInfo)
	mock.lockSign.Unlock()
	return mock.SignFunc(ctx, in, opts...)
}

// SignCalls gets all the calls that were made to Sign.
// Check the length with:
//     len(mockedMultiSigClient.SignCalls())
func (mock *MultiSigClientMock) SignCalls() []struct {
	Ctx  context.Context
	In   *tofnd.SignRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *tofnd.SignRequest
		Opts []grpc.CallOption
	}
	mock.lockSign.RLock()
	calls = mock.calls.Sign
	mock.lockSign.RUnlock()
	return calls
}
