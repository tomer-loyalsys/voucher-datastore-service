package grpc

import (
	"context"
	"github.com/loyalsys/error"
	"github.com/loyalsys/log"
)

func RecoverAndResponse(ctx context.Context, err *error) {
	if rec := recover(); rec != nil {
		*err = lserr.WrapErrf(rec.(error), "panic.")
		lslog.Errorf(ctx, *err, "panic.", rec).Write()
	}
}
