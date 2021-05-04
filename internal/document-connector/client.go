package document_connector

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"time"
)

const (
	connectTimeout  = 10
	queryTimeout    = 30
)

type DocClientConfig struct {
	ConnnectionString string
	TLSConfig *tls.Config
	ConnectCtx context.Context
	QueryCtx context.Context
	Client *mongo.Client
}

func NewDocClientConfig(username string, password string, clusterEndpoint string, caFilePath string) DocClientConfig {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s/?tls=true&tlsCAFile=rds-combined-ca-bundle.pem&tlsAllowInvalidHostnames=true&connect=direct&replicaSet=rs0&readPreference=secondaryPreferred&retryWrites=false", username, password, clusterEndpoint)
	fmt.Printf("Connection String: %s\n", connectionString)

	tlsConfig, err := getCustomTLSConfig(caFilePath)
	if err != nil {
		log.Fatalf("Failed getting TLS configuration: %v", err)
	}

	connectCtx, _ := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	//defer cancel()

	queryCtx, _ := context.WithTimeout(context.Background(), queryTimeout*time.Second)
	//defer cancel()

	docClientConfig := DocClientConfig{
		ConnnectionString: connectionString,
		TLSConfig:         tlsConfig,
		ConnectCtx:        connectCtx,
		QueryCtx:          queryCtx,
	}

	docClientConfig.NewClient()
	return docClientConfig
}

func (config *DocClientConfig) NewClient() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.ConnnectionString).SetTLSConfig(config.TLSConfig))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	err = client.Connect(config.ConnectCtx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	config.Client = client
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	tlsConfig.InsecureSkipVerify = true

	if !ok {
		return tlsConfig, errors.New("Failed parsing pem file")
	}

	return tlsConfig, nil
}
