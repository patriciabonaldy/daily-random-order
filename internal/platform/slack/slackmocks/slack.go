// Code generated by mockery v2.9.4. DO NOT EDIT.

package slackmocks

import mock "github.com/stretchr/testify/mock"

// Slack is an autogenerated mock type for the Slack type
type Slack struct {
	mock.Mock
}

// SendSlackNotification provides a mock function with given fields: msg
func (_m *Slack) SendSlackNotification(msg string) error {
	ret := _m.Called(msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
