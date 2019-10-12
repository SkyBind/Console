// +build windows

package proxy

import (
	"net"
	"net/http"

	"github.com/Microsoft/go-winio"
<<<<<<< HEAD
=======

>>>>>>> 80ad5079f7dcc2a26c0bf4fbc8b6db46fc38fa85
	portainer "github.com/portainer/portainer/api"
)

func (factory *proxyFactory) newLocalProxy(path string, endpoint *portainer.Endpoint) http.Handler {
	proxy := &localProxy{}
	transport := &proxyTransport{
		enableSignature:        false,
		ResourceControlService: factory.ResourceControlService,
		UserService:            factory.UserService,
		TeamMembershipService:  factory.TeamMembershipService,
		SettingsService:        factory.SettingsService,
		RegistryService:        factory.RegistryService,
		DockerHubService:       factory.DockerHubService,
		ReverseTunnelService:   factory.ReverseTunnelService,
		ExtensionService:       factory.ExtensionService,
		dockerTransport:        newNamedPipeTransport(path),
		endpointIdentifier:     endpoint.ID,
		endpointType:           endpoint.Type,
	}
	proxy.Transport = transport
	return proxy
}

func newNamedPipeTransport(namedPipePath string) *http.Transport {
	return &http.Transport{
		Dial: func(proto, addr string) (conn net.Conn, err error) {
			return winio.DialPipe(namedPipePath, nil)
		},
	}
}
