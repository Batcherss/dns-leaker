package resolver

import (
	"context"
	"net"
	"time"
)
// DANGER! DANGER! RADIATION!
// realy danger
func Resolve(domain, server string) ([]string, error) {
	dialer := &net.Dialer{
		Timeout: time.Second * 5,
	}
	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, "udp", server)
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ips, err := resolver.LookupHost(ctx, domain)
	if err != nil {
		return nil, err
	}
	return ips, nil
}
