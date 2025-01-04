package commonMock

import "github.com/stretchr/testify/mock"

type BaseMock struct {
	mock.Mock
}

func NewBaseMock() *BaseMock {
	return &BaseMock{}
}

func (m *BaseMock) Reset() {
	m.Mock.ExpectedCalls = nil
	m.Mock.Calls = nil
}

func (m *BaseMock) ProcessMethod(method string, args ...interface{}) (interface{}, error) {
	argsMock := m.MethodCalled(method, args...)

	numReturns := len(argsMock)

	if numReturns == 0 {
		return nil, nil
	}

	var err error
	lastIndex := numReturns - 1

	if e, ok := argsMock.Get(lastIndex).(error); ok {
		err = e
		numReturns--
	}

	if numReturns == 0 {
		return nil, err
	}

	return argsMock.Get(0), err
}
