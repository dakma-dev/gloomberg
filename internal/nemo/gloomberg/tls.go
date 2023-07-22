package gloomberg

import (
	"crypto/tls"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc/credentials"
)

func GetServerTLSConfig() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair(viper.GetString("tls.certificate"), viper.GetString("tls.key"))
	if err != nil {
		log.Error("TLS certificate (%s) not found, using insecure connection", viper.GetString("tls.certificate"))

		return &tls.Config{}, err //nolint:gosec
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   0,

		// disable mutual tls for now
		ClientAuth: tls.NoClientCert,
	}, nil
}

func GetTLSCredentialsWithoutClientAuth() (credentials.TransportCredentials, error) { //nolint:ireturn
	tlsConfig, err := GetServerTLSConfig()
	if err != nil {
		log.Error("TLS certificate not found, using insecure connection")

		return nil, err
	}

	return credentials.NewTLS(tlsConfig), nil
}

func GetTLSClientCredentials() credentials.TransportCredentials { //nolint:ireturn
	return credentials.NewTLS(&tls.Config{}) //nolint:gosec
}
