package hostmanager

import (
	"net"
	"strings"
)

type HostManager struct {
	Hosts    []net.IP
	position int
}

//AddHost adds a host to the host manager. Accepts IP or hostname strings.
// variadic can use string or []string
func (hm *HostManager) AddHost(hosts ...string) int {
	var addedHosts int
	for item := range hosts {
		hm.Hosts = append(hm.Hosts, net.ParseIP(hosts[item]))
		addedHosts = addedHosts + 1
	}
	return addedHosts
}

//AddHostSep adds a string of hosts which are delimited by sep.
func (hm *HostManager) AddHostSep(hosts string, sep string) int {
	sepHosts := strings.Split(hosts, sep)
	var addedHosts int
	for item := range sepHosts {
		hm.Hosts = append(hm.Hosts, net.ParseIP(sepHosts[item]))
		addedHosts = addedHosts + 1
	}
	return addedHosts
}

//RemoveHost removes a host from the most manager. Accepts IP or hostname strings.
// variadic can use string or []string
func (hm *HostManager) RemoveHost(hosts ...string) int {
	return 0
}

//AddHostIP adds a host to the host manager. Accepts net.IP or []net.IP
func (hm *HostManager) AddHostIP(ips ...net.IP) {

}

//RemoveHostIP removes a host from the most manager. Accepts net.IP or []net.IP
// variadic can use string or []string
func (hm *HostManager) RemoveHostIP(ips ...net.IP) {

}

//NextHost returns the net.IP of the next host
func (hm *HostManager) NextHost() (net.IP, int) {
	return net.IP{}, 0
}

//PrevHost returns the previous net.IP of the previous host
func (hm *HostManager) PrevHost() (net.IP, int) {
return net.IP{}, 0

}

//CurrentHost returns the current IP in the range. It also returns the current position
func (hm *HostManager) CurrentHost() (net.IP, int) {
return net.IP{}, 0
}

func (hm *HostManager) removeHost() {}
