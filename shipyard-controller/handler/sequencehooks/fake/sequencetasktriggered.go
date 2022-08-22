// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	apimodels "github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// ISequenceTaskTriggeredHookMock is a mock implementation of sequencehooks.ISequenceTaskTriggeredHook.
//
// 	func TestSomethingThatUsesISequenceTaskTriggeredHook(t *testing.T) {
//
// 		// make and configure a mocked sequencehooks.ISequenceTaskTriggeredHook
// 		mockedISequenceTaskTriggeredHook := &ISequenceTaskTriggeredHookMock{
// 			OnSequenceTaskTriggeredFunc: func(event apimodels.KeptnContextExtendedCE)  {
// 				panic("mock out the OnSequenceTaskTriggered method")
// 			},
// 		}
//
// 		// use mockedISequenceTaskTriggeredHook in code that requires sequencehooks.ISequenceTaskTriggeredHook
// 		// and then make assertions.
//
// 	}
type ISequenceTaskTriggeredHookMock struct {
	// OnSequenceTaskTriggeredFunc mocks the OnSequenceTaskTriggered method.
	OnSequenceTaskTriggeredFunc func(event apimodels.KeptnContextExtendedCE)

	// calls tracks calls to the methods.
	calls struct {
		// OnSequenceTaskTriggered holds details about calls to the OnSequenceTaskTriggered method.
		OnSequenceTaskTriggered []struct {
			//models.KeptnContextExtendedCEis the event argument value.
			Event apimodels.KeptnContextExtendedCE
		}
	}
	lockOnSequenceTaskTriggered sync.RWMutex
}

// OnSequenceTaskTriggered calls OnSequenceTaskTriggeredFunc.
func (mock *ISequenceTaskTriggeredHookMock) OnSequenceTaskTriggered(event apimodels.KeptnContextExtendedCE) {
	if mock.OnSequenceTaskTriggeredFunc == nil {
		panic("ISequenceTaskTriggeredHookMock.OnSequenceTaskTriggeredFunc: method is nil but ISequenceTaskTriggeredHook.OnSequenceTaskTriggered was just called")
	}
	callInfo := struct {
		Event apimodels.KeptnContextExtendedCE
	}{
		Event: event,
	}
	mock.lockOnSequenceTaskTriggered.Lock()
	mock.calls.OnSequenceTaskTriggered = append(mock.calls.OnSequenceTaskTriggered, callInfo)
	mock.lockOnSequenceTaskTriggered.Unlock()
	mock.OnSequenceTaskTriggeredFunc(event)
}

// OnSequenceTaskTriggeredCalls gets all the calls that were made to OnSequenceTaskTriggered.
// Check the length with:
//     len(mockedISequenceTaskTriggeredHook.OnSequenceTaskTriggeredCalls())
func (mock *ISequenceTaskTriggeredHookMock) OnSequenceTaskTriggeredCalls() []struct {
	Event apimodels.KeptnContextExtendedCE
} {
	var calls []struct {
		Event apimodels.KeptnContextExtendedCE
	}
	mock.lockOnSequenceTaskTriggered.RLock()
	calls = mock.calls.OnSequenceTaskTriggered
	mock.lockOnSequenceTaskTriggered.RUnlock()
	return calls
}