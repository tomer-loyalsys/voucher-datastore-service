package integration

import (
	"fmt"
	"github.com/loyalsys/error"
	"github.com/loyalsys/time"
	"github.com/loyalsys/voucher-datastore-service/grpc/proto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"testing"
	"time"
)

type rpcClientTestParams struct {
	region     string
	customerId string
	poolId     string
}

var grpcClient lsvoucherds.GrpcClient
var clientParams = rpcClientTestParams{region: "europe", customerId: "1", poolId: "343GDF655YH7TGS67"}

func init() {
	createClient()
}

func TestPingCient(t *testing.T) {

	out, err := grpcClient.Ping(context.Background(), &lsvoucherds.PingReq{})
	println(out)
	assert.NoError(t, err, "failed to ping client. err: %v", lserr.WithStack(err))
}

func TestUpload(t *testing.T) {
	var voucherBatch [][]string

	for i := 0; i < 10; i++ {
		var vouchersBatch []string
		for i := 0; i < 3000; i++ {
			vouchersBatch = append(vouchersBatch, fakeStringGenerator(10))
		}
		voucherBatch = append(voucherBatch, vouchersBatch)
	}

	in := lsvoucherds.UploadToPoolReq{Region: clientParams.region, CustomerId: clientParams.customerId, PoolId: clientParams.poolId, UploadId: fakeStringGenerator(5)}

	startTime := lstime.NowUTC()

	stream, err := grpcClient.UploadToPool(context.Background())
	if err != nil {
		log.Fatalf("failed to create stream.")
	}
	for _, batch := range voucherBatch {
		in.Vouchers = batch
		if err := stream.Send(&in); err != nil {
			if err == io.EOF {
				break
			}
			assert.NoError(t, err, "failed in stream send. err: %v", lserr.WithStack(err))
		}
	}
	reply, err := stream.CloseAndRecv()
	assert.NoError(t, err, "failed on pool upload test. err: %v", lserr.WithStack(err))

	println(fmt.Sprintf("elapsed time: %v", time.Since(startTime)))
	println(fmt.Sprintf("%+v", reply))
}

func TestGetPoolStatusRpc(t *testing.T) {
	in := lsvoucherds.GetPoolStatusReq{Region: clientParams.region, CustomerId: clientParams.customerId, PoolIds: []string{clientParams.poolId}}

	out, err := grpcClient.GetPoolStatus(context.Background(), &in)
	println(fmt.Sprintf("%+v", out))
	assert.NoError(t, err, "failed on get pool status test. err: %v", lserr.WithStack(err))
}

func TestDeletePoolRpc(t *testing.T) {
	in := lsvoucherds.DeletePoolReq{Region: clientParams.region, CustomerId: clientParams.customerId, PoolId: clientParams.poolId}

	out, err := grpcClient.DeletePool(context.Background(), &in)
	println(fmt.Sprintf("%+v", out))
	assert.NoError(t, err, "failed on delete pool test. err: %v", lserr.WithStack(err))
}

func createClient() {
	conn, _ := grpc.Dial(":9600", grpc.WithInsecure())
	grpcClient = lsvoucherds.NewGrpcClient(conn)
}
