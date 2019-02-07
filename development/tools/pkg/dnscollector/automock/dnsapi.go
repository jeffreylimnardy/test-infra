// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import dns "google.golang.org/api/dns/v1"

import mock "github.com/stretchr/testify/mock"

// DNSAPI is an autogenerated mock type for the DNSAPI type
type DNSAPI struct {
	mock.Mock
}

// DeleteDNSRecord provides a mock function with given fields: project, managedZone, record
func (_m *DNSAPI) DeleteDNSRecord(project string, managedZone string, record *dns.ResourceRecordSet) error {
	ret := _m.Called(project, managedZone, record)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *dns.ResourceRecordSet) error); ok {
		r0 = rf(project, managedZone, record)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LookupDNSRecords provides a mock function with given fields: project, managedZone
func (_m *DNSAPI) LookupDNSRecords(project string, managedZone string) ([]*dns.ResourceRecordSet, error) {
	ret := _m.Called(project, managedZone)

	var r0 []*dns.ResourceRecordSet
	if rf, ok := ret.Get(0).(func(string, string) []*dns.ResourceRecordSet); ok {
		r0 = rf(project, managedZone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dns.ResourceRecordSet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(project, managedZone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
