package lib

import (
	"net"

	"github.com/cloudfoundry-incubator/silk/config"
	"github.com/containernetworking/cni/pkg/ns"
	"github.com/vishvananda/netlink"
)

//go:generate counterfeiter -o fakes/linkOperations.go --fake-name LinkOperations . linkOperations
type linkOperations interface {
	DisableIPv6(deviceName string) error
	StaticNeighborNoARP(link netlink.Link, dstIP net.IP, mac net.HardwareAddr) error
	SetPointToPointAddress(link netlink.Link, localIPAddr, peerIPAddr net.IP) error
	RenameLink(oldName, newName string) error
	DeleteLinkByName(deviceName string) error
}

//go:generate counterfeiter -o fakes/common.go --fake-name Common . common
type common interface {
	BasicSetup(deviceName string, local, peer config.DualAddress) error
}

//go:generate counterfeiter -o fakes/namespaceAdapter.go --fake-name NamespaceAdapter . namespaceAdapter
type namespaceAdapter interface {
	GetNS(string) (ns.NetNS, error)
	GetCurrentNS() (ns.NetNS, error)
}

//go:generate counterfeiter -o fakes/netlinkAdapter.go --fake-name NetlinkAdapter . netlinkAdapter
type netlinkAdapter interface {
	LinkByName(string) (netlink.Link, error)
	ParseAddr(string) (*netlink.Addr, error)
	AddrAddScopeLink(netlink.Link, *netlink.Addr) error
	LinkSetHardwareAddr(netlink.Link, net.HardwareAddr) error
	// NeighAdd(*netlink.Neigh) error
	NeighAddPermanentIPv4(index int, destIP net.IP, hwAddr net.HardwareAddr) error
	LinkSetARPOff(netlink.Link) error
	LinkSetName(netlink.Link, string) error
	LinkSetUp(netlink.Link) error
	LinkDel(netlink.Link) error
	LinkAdd(netlink.Link) error
	LinkSetNsFd(netlink.Link, int) error
}

//go:generate counterfeiter -o fakes/netNS.go --fake-name NetNS . netNS
type netNS interface {
	ns.NetNS
}

func NetNsDoStub(f func(h ns.NetNS) error) error {
	return f(nil)
}

//go:generate counterfeiter -o fakes/sysctlAdapter.go --fake-name SysctlAdapter . sysctlAdapter
type sysctlAdapter interface {
	Sysctl(name string, params ...string) (string, error)
}
