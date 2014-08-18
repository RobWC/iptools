package hostmanager

import (
	"bytes"
	"errors"
	"net"
	"strings"
)

//HostManager manages a list of IP address
// contains several helper functions to easily add and remove IPs, supports IPv4 and IPv6
type HostManager struct {
	Hosts    []net.IP
	position int
}

//NewHostManager creates an initialized host manager
// default position set to 0
func NewHostManager() HostManager {
	return HostManager{position: 0, Hosts: make([]net.IP, 0)}
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
func (hm *HostManager) RemoveHost(hosts ...string) ([]int, int) {
	hostsRemoved := 0
	var hostPositions []int
	for host := range hosts {
		pos, err := hm.removeHost(net.ParseIP(hosts[host]))
		if err != nil {
			//suppress error
		} else {
			hostsRemoved = hostsRemoved + 1
		}
		hostPositions = append(hostPositions, pos)
	}
	return hostPositions, hostsRemoved
}

//FindHost returns the position of a host or error if not found
func (hm *HostManager) FindHost(host net.IP) (int, error) {
	for checkHost := range hm.Hosts {
		if bytes.Compare(host, hm.Hosts[checkHost]) == 0 {
			return checkHost, nil
		}
	}
	return -1, errors.New("no host found")
}

//AddHostIP adds a host to the host manager. Accepts net.IP or []net.IP
func (hm *HostManager) AddHostIP(ips ...net.IP) int {
	var addedHosts int
	for item := range ips {
		hm.Hosts = append(hm.Hosts, ips[item])
		addedHosts = addedHosts + 1
	}
	return addedHosts
}

//RemoveHostIP removes a host from the most manager. Accepts net.IP or []net.IP
// variadic can use string or []string
func (hm *HostManager) RemoveHostIP(ips ...net.IP) ([]int, int) {
	hostsRemoved := 0
	var hostPositions []int
	for host := range ips {
		pos, err := hm.removeHost(ips[host])
		if err != nil {
			//suppress error
		} else {
			hostsRemoved = hostsRemoved + 1
		}
		hostPositions = append(hostPositions, pos)
	}
	return hostPositions, hostsRemoved
}

//NextHost returns the net.IP of the next host
func (hm *HostManager) NextHost() (net.IP, int) {
	if hm.position == len(hm.Hosts) {
		hm.position = 0
	} else {
		hm.position = hm.position + 1
	}
	return hm.Hosts[hm.position], hm.position
}

//PrevHost returns the previous net.IP of the previous host
func (hm *HostManager) PrevHost() (net.IP, int) {
	if hm.position == 0 {
		hm.position = len(hm.Hosts) - 1
	} else if hm.position >= 1 {
		hm.position = hm.position - 1
	}
	return hm.Hosts[hm.position], hm.position
}

//CurrentHost returns the current IP in the range. It also returns the current position
func (hm *HostManager) CurrentHost() (net.IP, int) {
	return hm.Hosts[hm.position], hm.position
}

//removeHost
// remove host internal function to simplify code
func (hm *HostManager) removeHost(host net.IP) (int, error) {
	hostPosition, err := hm.FindHost(host)
	if err != nil {
		return -1, err
	}
	hm.Hosts = append(hm.Hosts[:hostPosition], hm.Hosts[hostPosition+1:]...)
	return hostPosition, nil
}
