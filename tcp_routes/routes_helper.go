package tcp_routes

import (
	"encoding/json"

	"github.com/cloudfoundry-incubator/bbs/models"
)

const TCP_ROUTER = "tcp-router"

type TCPRoutes []TCPRoute

type TCPRoute struct {
	ExternalPort  uint32 `json:"external_port"`
	ContainerPort uint32 `json:"container_port"`
}

func (c TCPRoutes) RoutingInfo() *models.Routes {
	data, _ := json.Marshal(c)
	routingInfo := json.RawMessage(data)
	return &models.Routes{
		TCP_ROUTER: &routingInfo,
	}
}

func TCPRoutesFromRoutingInfo(routingInfoPtr *models.Routes) (TCPRoutes, error) {
	if routingInfoPtr == nil {
		return nil, nil
	}

	routingInfo := *routingInfoPtr

	data, found := routingInfo[TCP_ROUTER]
	if !found {
		return nil, nil
	}

	if data == nil {
		return nil, nil
	}

	routes := TCPRoutes{}
	err := json.Unmarshal(*data, &routes)

	return routes, err
}
