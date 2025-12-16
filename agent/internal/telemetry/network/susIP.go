package network

import (
	"fmt"
	"net"
)

// get the machine all ips
// check them against suspiouse ip if so report them as suspiouse

func FilterSusIp() map[string]struct{} {
	susIPs := map[string]struct{}{
		"192.168.1.10":   {},
		"10.0.0.5":       {},
		"172.16.0.7":     {},
		"10.255.255.254": {},
	}

	collectSusIp := map[string]struct{}{}
	iface, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range iface {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP

			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			fmt.Println("ip:=>", ip)
			if ip == nil {
				continue
			}
			ipstr := ip.String()

			if _, exists := susIPs[ipstr]; exists {
				collectSusIp[ipstr] = struct{}{}
			}
		}
	}

	return collectSusIp
}
