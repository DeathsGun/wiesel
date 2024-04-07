package netx

import (
	"net"
	"net/netip"
	"strings"
)

type HostList struct {
	Hosts     []netip.Addr
	shortList string
}

func (h *HostList) GoString() string {
	return h.shortList
}

func (h *HostList) UnmarshalText(text []byte) error {
	hostList := ParseHosts(string(text))
	h.Hosts = hostList.Hosts
	h.shortList = hostList.shortList
	return nil
}

func ParseHosts(input string) *HostList {
	if input == "" {
		return &HostList{Hosts: make([]netip.Addr, 0)}
	}
	if ip, err := netip.ParseAddr(input); err == nil {
		if ip.IsUnspecified() {
			return &HostList{Hosts: make([]netip.Addr, 0)}
		}
		return &HostList{Hosts: []netip.Addr{ip}, shortList: input}
	}

	if prefix, err := netip.ParsePrefix(input); err == nil {
		var ips []netip.Addr
		for addr := prefix.Addr(); prefix.Contains(addr); addr = addr.Next() {
			if addr.IsUnspecified() {
				continue
			}
			ips = append(ips, addr)
		}
		// remove network address and broadcast address
		return &HostList{Hosts: uniqueHosts(ips[1 : len(ips)-1]), shortList: input}
	}

	if strings.Contains(input, ",") {
		var shortList []string
		var result []netip.Addr
		for _, host := range strings.Split(input, ",") {
			hostList := ParseHosts(host)
			if hostList.shortList != "" {
				shortList = append(shortList, hostList.shortList)
			}
			result = append(result, hostList.Hosts...)
		}
		return &HostList{Hosts: uniqueHosts(result), shortList: strings.Join(shortList, ",")}
	}

	hostIps, err := net.LookupIP(input)
	if err != nil {
		return &HostList{Hosts: make([]netip.Addr, 0)}
	}
	hosts := make([]netip.Addr, 0)
	for _, host := range hostIps {
		ip, ok := netip.AddrFromSlice(host)
		if !ok || ip.IsUnspecified() {
			continue
		}
		hosts = append(hosts, ip)
	}

	return &HostList{Hosts: uniqueHosts(hosts), shortList: input}
}

func uniqueHosts(hosts []netip.Addr) []netip.Addr {
	unique := make([]netip.Addr, 0)
	for _, host := range hosts {
		if contains(unique, host) {
			continue
		}
		unique = append(unique, host)
	}
	return unique
}

func contains(hosts []netip.Addr, host netip.Addr) bool {
	for _, h := range hosts {
		if h.Compare(host) == 0 {
			return true
		}
	}
	return false
}
