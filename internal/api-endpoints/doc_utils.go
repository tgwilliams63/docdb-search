package api_endpoints

import (
	"fmt"
	document_connector "github.com/tgwilliams/simple-search-ui/internal/document-connector"
	"log"
	"os"
)

func GetDocConfig() document_connector.DocClientConfig {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	clusterEndpoint := os.Getenv("CLUSTER_ENDPOINT")
	caFilePath := "rds-combined-ca-bundle.pem"
	docClientConfig := document_connector.NewDocClientConfig(username, password, clusterEndpoint, caFilePath)

	// Force a connection to verify our connection string
	err := docClientConfig.Client.Ping(docClientConfig.ConnectCtx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	fmt.Println("Connected to DocumentDB!")
	return docClientConfig
}
