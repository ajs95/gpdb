package utils

import (
	"fmt"
)

//counterfeiter:generate . HostUtil
type HostUtil interface {
	VerifyHostIp(string) error
}

func NewHostUtil() HostUtil {
	return &hostUtil{}
}

type hostUtil struct{}

func (hostUtil) VerifyHostIp(ip string) error {
	netUtil := Dependency.NewNetworkUtil()
	netInterfaces, err := netUtil.Interfaces()
	if err != nil {
		return err
	}
	for _, netInterface := range netInterfaces {
		addrs, err := netInterface.Addrs()
		if err != nil {
			return err
		}
		for _, addr := range addrs {
			if addr.String() == ip {
				return nil
			}
		}
	}
	return fmt.Errorf("invalid host ip")
}
