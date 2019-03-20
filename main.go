package main

import (
	"fmt"
	"github.com/heptiolabs/healthcheck"
	"github.com/loyalsys/error"
	"github.com/loyalsys/google-cloud"
	"github.com/loyalsys/log"
	"github.com/loyalsys/lscorrid"
	"github.com/loyalsys/lsenv"
	"github.com/loyalsys/voucher-datastore-service/datastore"
	voucherGrpc "github.com/loyalsys/voucher-datastore-service/grpc"
	"github.com/loyalsys/voucher-datastore-service/grpc/proto"
	"github.com/oklog/run"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

var version string

const (
	envVarGoogleProjectId        lsenv.EnvVar = "GOOGLE_PROJECT_ID"
	envVarVoucherDsServerAddress lsenv.EnvVar = "VOUCHER_DATASTORE_SERVER_ADDRESS"
	envVarHttpHealthServerPort   lsenv.EnvVar = "HTTP_HEALTH_SERVER_PORT"
)

var envVarToRead = []lsenv.EnvVar{
	envVarGoogleProjectId,
	envVarVoucherDsServerAddress,
	envVarHttpHealthServerPort,
}

var envInfo lsenv.EnvInfo

var mServer *voucherGrpc.Service
var g run.Group

func main() {
	fmt.Println("main start - version: " + version)

	// load environment variables
	fPath, _ := filepath.Abs("./")
	err := lsenv.LoadEnvironmentVariables(fPath)
	if err != nil {
		log.Panicf("Can't continue without ALL ENV variables. %v", lserr.WithStack(err))
	}

	// generate google project credential
	err = lsgc.GenerateGoogleCredential()
	if err != nil {
		log.Panicf("failed to generate google credential. %v", lserr.WithStack(err))
	}

	// read environment variables
	envInfo, err = lsenv.ReadEnvironmentVariables(envVarToRead)
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to read env vars.")
	}

	// handle server stop
	gracefullStop()

	// get datastore
	ds, err := datastore.CreateInstance()
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to connect to datastore.")
	}
	lslog.Infof(context.Background(), "success to connect to datastore.").Write()

	// start health listener server
	lslog.Infof(context.Background(), "start health listener.").Write()
	err = runHealthListener(ds)
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to start health listener.")
	}

	// start grpc server
	lslog.Infof(context.Background(), "start gRPC server.").Write()
	err = startGrpcServer(ds)
	if err != nil {
		lslog.WritefAndExit(context.Background(), lslog.ErrorLevel, err, "failed to start/run gRPC server.")
	}

	// all is ok - write `server is up log`
	lslog.Infof(context.TODO(), "server is up. running version: %s", version).Write()

	// run the goroutine group and write log after exit
	lslog.WritefAndExit(context.Background(), lslog.InfoLevel, g.Run(), "server is down.")
}

func gracefullStop() {
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		lslog.WritefAndExit(context.Background(), lslog.InfoLevel, lserr.NewErrf(""), "server been taking down due to stop interrupt with sig: %#v", sig)
	}()
}

func startGrpcServer(ds *datastore.Datastore) (err error) {
	lis, err := net.Listen("tcp", envInfo[envVarVoucherDsServerAddress])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	interceptor := grpc.UnaryInterceptor(lscorrid.GrpcInterceptor)
	srv := grpc.NewServer(interceptor)

	// register our service with the gRPC server
	mServer = &voucherGrpc.Service{Ds: ds, AppVersion: version}
	lsvoucherds.RegisterGrpcServer(srv, mServer)

	// add listener to goroutine group handler
	g.Add(func() error {
		err := srv.Serve(lis)
		return lserr.WrapErrf(err, "failed to serve grpc.")
	}, func(err error) {
		lslog.WritefAndExit(context.Background(), lslog.InfoLevel, err, "grpc server is down.")
	})

	return err
}

func runHealthListener(ds *datastore.Datastore) (err error) {
	m := http.NewServeMux()
	health := healthcheck.NewHandler()

	health.AddReadinessCheck("/ready", func() error { return healthReadyCheck(ds) })
	m.HandleFunc("/ready", health.ReadyEndpoint)

	// add listener to goroutine group handler
	g.Add(func() error {
		err := http.ListenAndServe(envInfo[envVarHttpHealthServerPort], http.Handler(m))
		return lserr.WrapErrf(err, "failed to listen and serve health listener")
	}, func(err error) {
		lslog.WritefAndExit(context.Background(), lslog.InfoLevel, err, "health server is down.")
	})

	return err
}

func healthReadyCheck(ds *datastore.Datastore) error {
	for _, r := range ds.Voucher {
		err := r.Ping()
		if err != nil {
			err = lserr.WrapErrf(err, "Health return non healthy state.")
			log.Printf("%v", err)
			return err
		}
	}

	return nil
}
