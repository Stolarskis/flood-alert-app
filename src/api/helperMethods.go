package pb

import (
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func getIPEnv() string {
	ip := os.Getenv("KUBERNETES_IP")
	if ip == "" {
		return "localhost"
	}
	return ip
}
func CreateClient() (FloodAlertAppServiceClient, error) {
	ip := getIPEnv()
	conn, err := grpc.Dial(ip+":3000", grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create grpc client")
	}
	return NewFloodAlertAppServiceClient(conn), nil
}
