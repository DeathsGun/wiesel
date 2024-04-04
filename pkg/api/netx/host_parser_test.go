package netx

import (
	"testing"
)

func TestParseSimpleIP(t *testing.T) {
	ips := ParseHosts("192.168.178.3")
	if len(ips.Hosts) != 1 {
		t.Errorf("Expected 1 IP, got %d", len(ips.Hosts))
	}
	if ips.Hosts[0].String() != "192.168.178.3" {
		t.Errorf("Expected 192.168.178.3 got %s", ips.Hosts[0])
	}
}

func TestParseCIDR(t *testing.T) {
	ips := ParseHosts("192.168.178.1/24")
	if len(ips.Hosts) != 254 {
		t.Errorf("Expected 254 IPs, got %d", len(ips.Hosts))
	}
	if ips.Hosts[0].String() != "192.168.178.1" {
		t.Errorf("Expected 192.168.178.1 got %s", ips.Hosts[0])
	}
	if ips.Hosts[253].String() != "192.168.178.254" {
		t.Errorf("Expected 192.168.178.254 got %s", ips.Hosts[253])
	}
}

func TestParseMultipleIPs(t *testing.T) {
	ips := ParseHosts("192.168.178.3,192.168.1.2")
	if len(ips.Hosts) != 2 {
		t.Errorf("Expected 2 IPs, got %d", len(ips.Hosts))
	}
	if ips.Hosts[0].String() != "192.168.178.3" {
		t.Errorf("Expected 192.168.178.3 got %s", ips.Hosts[0])
	}

	if ips.Hosts[1].String() != "192.168.1.2" {
		t.Errorf("Expected 192.168.1.2 got %s", ips.Hosts[1])
	}
}

func TestParseEmpty(t *testing.T) {
	ips := ParseHosts("")
	if ips != nil {
		t.Error("Expected nil")
	}
}

func TestParseInvalid(t *testing.T) {
	ips := ParseHosts("invalid")
	if ips != nil {
		t.Error("Expected nil")
	}
}

func TestParseInvalidCIDR(t *testing.T) {
	ips := ParseHosts("192.168.178.1/33")
	if ips != nil {
		t.Error("Expected nil")
	}
}

func TestParseDomain(t *testing.T) {
	hosts := ParseHosts("google.com")
	if len(hosts.Hosts) == 0 {
		t.Errorf("Expected minimal 1 host, got %d", len(hosts.Hosts))
	}
}
