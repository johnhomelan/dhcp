package dhcpv4

import (
	"gitlab.ask4.net/jbrown/dhcp.git/interfaces"
)

// BindToInterface (deprecated) redirects to interfaces.BindToInterface
func BindToInterface(fd int, ifname string) error {
	return interfaces.BindToInterface(fd, ifname)
}
