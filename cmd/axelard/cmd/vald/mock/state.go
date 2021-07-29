// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/cmd/axelard/cmd/vald"
	"sync"
)

// Ensure, that ReadWriterMock does implement vald.ReadWriter.
// If this is not the case, regenerate this file with moq.
var _ vald.ReadWriter = &ReadWriterMock{}

// ReadWriterMock is a mock implementation of vald.ReadWriter.
//
// 	func TestSomethingThatUsesReadWriter(t *testing.T) {
//
// 		// make and configure a mocked vald.ReadWriter
// 		mockedReadWriter := &ReadWriterMock{
// 			ReadAllFunc: func() ([]byte, error) {
// 				panic("mock out the ReadAll method")
// 			},
// 			WriteAllFunc: func(bytes []byte) error {
// 				panic("mock out the WriteAll method")
// 			},
// 		}
//
// 		// use mockedReadWriter in code that requires vald.ReadWriter
// 		// and then make assertions.
//
// 	}
type ReadWriterMock struct {
	// ReadAllFunc mocks the ReadAll method.
	ReadAllFunc func() ([]byte, error)

	// WriteAllFunc mocks the WriteAll method.
	WriteAllFunc func(bytes []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// ReadAll holds details about calls to the ReadAll method.
		ReadAll []struct {
		}
		// WriteAll holds details about calls to the WriteAll method.
		WriteAll []struct {
			// Bytes is the bytes argument value.
			Bytes []byte
		}
	}
	lockReadAll  sync.RWMutex
	lockWriteAll sync.RWMutex
}

// ReadAll calls ReadAllFunc.
func (mock *ReadWriterMock) ReadAll() ([]byte, error) {
	if mock.ReadAllFunc == nil {
		panic("ReadWriterMock.ReadAllFunc: method is nil but ReadWriter.ReadAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockReadAll.Lock()
	mock.calls.ReadAll = append(mock.calls.ReadAll, callInfo)
	mock.lockReadAll.Unlock()
	return mock.ReadAllFunc()
}

// ReadAllCalls gets all the calls that were made to ReadAll.
// Check the length with:
//     len(mockedReadWriter.ReadAllCalls())
func (mock *ReadWriterMock) ReadAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockReadAll.RLock()
	calls = mock.calls.ReadAll
	mock.lockReadAll.RUnlock()
	return calls
}

// WriteAll calls WriteAllFunc.
func (mock *ReadWriterMock) WriteAll(bytes []byte) error {
	if mock.WriteAllFunc == nil {
		panic("ReadWriterMock.WriteAllFunc: method is nil but ReadWriter.WriteAll was just called")
	}
	callInfo := struct {
		Bytes []byte
	}{
		Bytes: bytes,
	}
	mock.lockWriteAll.Lock()
	mock.calls.WriteAll = append(mock.calls.WriteAll, callInfo)
	mock.lockWriteAll.Unlock()
	return mock.WriteAllFunc(bytes)
}

// WriteAllCalls gets all the calls that were made to WriteAll.
// Check the length with:
//     len(mockedReadWriter.WriteAllCalls())
func (mock *ReadWriterMock) WriteAllCalls() []struct {
	Bytes []byte
} {
	var calls []struct {
		Bytes []byte
	}
	mock.lockWriteAll.RLock()
	calls = mock.calls.WriteAll
	mock.lockWriteAll.RUnlock()
	return calls
}
