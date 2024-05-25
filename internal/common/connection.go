package common

import (
	"crypto/tls"
	"crypto/x509"
	"strings"

	"github.com/userosettadev/rosetta-cli/internal/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func BuildGRPCConnection() (*grpc.ClientConn, error) {

	host := env.GetInstance().GetHome()
	opts := []grpc.DialOption{grpc.WithAuthority(host)}

	if strings.HasPrefix(host, "localhost:") {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{RootCAs: systemRoots})))
	}

	conn, err := grpc.NewClient(host, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
