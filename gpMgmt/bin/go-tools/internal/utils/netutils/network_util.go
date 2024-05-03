package netutils

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . NetworkUtil
type NetworkUtil interface {
	IsValidIpv4(string) bool
	Interfaces() ([]net.Interface, error)
	GenerateIpv4List(string, string) ([]string, error)
	IsReachable(string) bool
}

func NewNetworkUtil() NetworkUtil {
	return &networkUtil{}
}

type networkUtil struct{}

func (networkUtil) Interfaces() ([]net.Interface, error) {
	return net.Interfaces()
}

func (networkUtil) isValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

func (n networkUtil) IsValidIpv4(ip string) bool {
	return n.isValidIp(ip) && strings.Count(ip, ":") == 0
}

func (n networkUtil) GenerateIpv4List(first, last string) ([]string, error) {
	if !n.IsValidIpv4(first) || !n.IsValidIpv4(last) {
		return nil, fmt.Errorf("ip range is invalid")
	}
	ipList := []string{}
	currentIp := net.ParseIP(first).To4()
	lastIp := net.ParseIP(last).To4()
	for currentIp.String() != lastIp.String() {
		ipList = append(ipList, currentIp.String())
		currentIp = n.nextIPv4Address(currentIp)
	}
	return append(ipList, lastIp.String()), nil
}

func (n networkUtil) nextIPv4Address(ip net.IP) net.IP {
	nextIp := ip.To4()
	nextIp[3]++
	if nextIp[3] == byte(0) {
		nextIp[2]++
		if nextIp[2] == byte(0) {
			nextIp[1]++
			if nextIp[1] == byte(0) {
				nextIp[0]++
			}
		}
	}
	return nextIp
}

func (n networkUtil) IsReachable(ip string) bool {
	if !n.isValidIp(ip) {
		return false
	}
	out, _ := exec.Command("ping", ip, "-c 5", "-i 3", "-w 10").Output()
	return !strings.Contains(string(out), "Destination Host Unreachable")
}
