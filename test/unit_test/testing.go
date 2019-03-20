package unit_test

import (
	"context"
	"github.com/loyalsys/google-cloud"
	"github.com/loyalsys/log"
	"github.com/loyalsys/lsenv"
	"github.com/loyalsys/voucher-datastore-service/datastore"
	"github.com/loyalsys/voucher-datastore-service/grpc"
	"log"
	"path/filepath"
)

var s *grpc.Service

func init() {
	const (
		envVarGoogleProjectId         lsenv.EnvVar = "GOOGLE_PROJECT_ID"
		envVarCustomerDsServerAddress lsenv.EnvVar = "POINTS_DATASTORE_SERVER_ADDRESS"
		envVarHttpHealthServerPort    lsenv.EnvVar = "HTTP_HEALTH_SERVER_PORT"
	)

	var envVarToRead = []lsenv.EnvVar{
		envVarGoogleProjectId,
		envVarCustomerDsServerAddress,
		envVarHttpHealthServerPort,
	}
	// load environment variables
	fPath, _ := filepath.Abs("../")
	err := lsenv.LoadEnvironmentVariables(fPath)
	if err != nil {
		log.Panicf("Can't continue without ALL ENV variables. %v", err)
	}

	// generate google project credential
	lsgc.GenerateGoogleCredential()

	// read environment variables
	_, err = lsenv.ReadEnvironmentVariables(envVarToRead)
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to read env vars.")
	}

	// get datastore
	ds, err := datastore.CreateInstance()
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to connect to datastore.")
	}

	s = &grpc.Service{Ds: ds}
}
