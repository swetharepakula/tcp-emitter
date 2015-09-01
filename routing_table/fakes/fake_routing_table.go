// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/receptor"
	"github.com/cloudfoundry-incubator/tcp-emitter/routing_table"
)

type FakeRoutingTable struct {
	RouteCountStub        func() int
	routeCountMutex       sync.RWMutex
	routeCountArgsForCall []struct{}
	routeCountReturns struct {
		result1 int
	}
	SetRoutesStub        func(desiredLRP receptor.DesiredLRPResponse) routing_table.RoutingEvents
	setRoutesMutex       sync.RWMutex
	setRoutesArgsForCall []struct {
		desiredLRP receptor.DesiredLRPResponse
	}
	setRoutesReturns struct {
		result1 routing_table.RoutingEvents
	}
	AddEndpointStub        func(actualLRP receptor.ActualLRPResponse) routing_table.RoutingEvents
	addEndpointMutex       sync.RWMutex
	addEndpointArgsForCall []struct {
		actualLRP receptor.ActualLRPResponse
	}
	addEndpointReturns struct {
		result1 routing_table.RoutingEvents
	}
	RemoveEndpointStub        func(actualLRP receptor.ActualLRPResponse) routing_table.RoutingEvents
	removeEndpointMutex       sync.RWMutex
	removeEndpointArgsForCall []struct {
		actualLRP receptor.ActualLRPResponse
	}
	removeEndpointReturns struct {
		result1 routing_table.RoutingEvents
	}
	SwapStub        func(t routing_table.RoutingTable) routing_table.RoutingEvents
	swapMutex       sync.RWMutex
	swapArgsForCall []struct {
		t routing_table.RoutingTable
	}
	swapReturns struct {
		result1 routing_table.RoutingEvents
	}
	GetRoutingEventsStub        func() routing_table.RoutingEvents
	getRoutingEventsMutex       sync.RWMutex
	getRoutingEventsArgsForCall []struct{}
	getRoutingEventsReturns struct {
		result1 routing_table.RoutingEvents
	}
}

func (fake *FakeRoutingTable) RouteCount() int {
	fake.routeCountMutex.Lock()
	fake.routeCountArgsForCall = append(fake.routeCountArgsForCall, struct{}{})
	fake.routeCountMutex.Unlock()
	if fake.RouteCountStub != nil {
		return fake.RouteCountStub()
	} else {
		return fake.routeCountReturns.result1
	}
}

func (fake *FakeRoutingTable) RouteCountCallCount() int {
	fake.routeCountMutex.RLock()
	defer fake.routeCountMutex.RUnlock()
	return len(fake.routeCountArgsForCall)
}

func (fake *FakeRoutingTable) RouteCountReturns(result1 int) {
	fake.RouteCountStub = nil
	fake.routeCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRoutingTable) SetRoutes(desiredLRP receptor.DesiredLRPResponse) routing_table.RoutingEvents {
	fake.setRoutesMutex.Lock()
	fake.setRoutesArgsForCall = append(fake.setRoutesArgsForCall, struct {
		desiredLRP receptor.DesiredLRPResponse
	}{desiredLRP})
	fake.setRoutesMutex.Unlock()
	if fake.SetRoutesStub != nil {
		return fake.SetRoutesStub(desiredLRP)
	} else {
		return fake.setRoutesReturns.result1
	}
}

func (fake *FakeRoutingTable) SetRoutesCallCount() int {
	fake.setRoutesMutex.RLock()
	defer fake.setRoutesMutex.RUnlock()
	return len(fake.setRoutesArgsForCall)
}

func (fake *FakeRoutingTable) SetRoutesArgsForCall(i int) receptor.DesiredLRPResponse {
	fake.setRoutesMutex.RLock()
	defer fake.setRoutesMutex.RUnlock()
	return fake.setRoutesArgsForCall[i].desiredLRP
}

func (fake *FakeRoutingTable) SetRoutesReturns(result1 routing_table.RoutingEvents) {
	fake.SetRoutesStub = nil
	fake.setRoutesReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) AddEndpoint(actualLRP receptor.ActualLRPResponse) routing_table.RoutingEvents {
	fake.addEndpointMutex.Lock()
	fake.addEndpointArgsForCall = append(fake.addEndpointArgsForCall, struct {
		actualLRP receptor.ActualLRPResponse
	}{actualLRP})
	fake.addEndpointMutex.Unlock()
	if fake.AddEndpointStub != nil {
		return fake.AddEndpointStub(actualLRP)
	} else {
		return fake.addEndpointReturns.result1
	}
}

func (fake *FakeRoutingTable) AddEndpointCallCount() int {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return len(fake.addEndpointArgsForCall)
}

func (fake *FakeRoutingTable) AddEndpointArgsForCall(i int) receptor.ActualLRPResponse {
	fake.addEndpointMutex.RLock()
	defer fake.addEndpointMutex.RUnlock()
	return fake.addEndpointArgsForCall[i].actualLRP
}

func (fake *FakeRoutingTable) AddEndpointReturns(result1 routing_table.RoutingEvents) {
	fake.AddEndpointStub = nil
	fake.addEndpointReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) RemoveEndpoint(actualLRP receptor.ActualLRPResponse) routing_table.RoutingEvents {
	fake.removeEndpointMutex.Lock()
	fake.removeEndpointArgsForCall = append(fake.removeEndpointArgsForCall, struct {
		actualLRP receptor.ActualLRPResponse
	}{actualLRP})
	fake.removeEndpointMutex.Unlock()
	if fake.RemoveEndpointStub != nil {
		return fake.RemoveEndpointStub(actualLRP)
	} else {
		return fake.removeEndpointReturns.result1
	}
}

func (fake *FakeRoutingTable) RemoveEndpointCallCount() int {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return len(fake.removeEndpointArgsForCall)
}

func (fake *FakeRoutingTable) RemoveEndpointArgsForCall(i int) receptor.ActualLRPResponse {
	fake.removeEndpointMutex.RLock()
	defer fake.removeEndpointMutex.RUnlock()
	return fake.removeEndpointArgsForCall[i].actualLRP
}

func (fake *FakeRoutingTable) RemoveEndpointReturns(result1 routing_table.RoutingEvents) {
	fake.RemoveEndpointStub = nil
	fake.removeEndpointReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) Swap(t routing_table.RoutingTable) routing_table.RoutingEvents {
	fake.swapMutex.Lock()
	fake.swapArgsForCall = append(fake.swapArgsForCall, struct {
		t routing_table.RoutingTable
	}{t})
	fake.swapMutex.Unlock()
	if fake.SwapStub != nil {
		return fake.SwapStub(t)
	} else {
		return fake.swapReturns.result1
	}
}

func (fake *FakeRoutingTable) SwapCallCount() int {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return len(fake.swapArgsForCall)
}

func (fake *FakeRoutingTable) SwapArgsForCall(i int) routing_table.RoutingTable {
	fake.swapMutex.RLock()
	defer fake.swapMutex.RUnlock()
	return fake.swapArgsForCall[i].t
}

func (fake *FakeRoutingTable) SwapReturns(result1 routing_table.RoutingEvents) {
	fake.SwapStub = nil
	fake.swapReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

func (fake *FakeRoutingTable) GetRoutingEvents() routing_table.RoutingEvents {
	fake.getRoutingEventsMutex.Lock()
	fake.getRoutingEventsArgsForCall = append(fake.getRoutingEventsArgsForCall, struct{}{})
	fake.getRoutingEventsMutex.Unlock()
	if fake.GetRoutingEventsStub != nil {
		return fake.GetRoutingEventsStub()
	} else {
		return fake.getRoutingEventsReturns.result1
	}
}

func (fake *FakeRoutingTable) GetRoutingEventsCallCount() int {
	fake.getRoutingEventsMutex.RLock()
	defer fake.getRoutingEventsMutex.RUnlock()
	return len(fake.getRoutingEventsArgsForCall)
}

func (fake *FakeRoutingTable) GetRoutingEventsReturns(result1 routing_table.RoutingEvents) {
	fake.GetRoutingEventsStub = nil
	fake.getRoutingEventsReturns = struct {
		result1 routing_table.RoutingEvents
	}{result1}
}

var _ routing_table.RoutingTable = new(FakeRoutingTable)
