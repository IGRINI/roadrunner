package http

import (
	"github.com/pkg/errors"
	"github.com/spiral/roadrunner/util"
)

type rpcServer struct{ svc *Service }

// WorkerList contains list of workers.
type WorkerList struct {
	// Workers is list of workers.
	Workers []*util.State `json:"workers"`
}

// Reset resets underlying RR worker pool and restarts all of it's workers.
func (rpc *rpcServer) Reset(args []string, r *string) error {
	if rpc.svc == nil || rpc.svc.handler == nil {
		return errors.New("http server is not running")
	}

	// run args set
	rpc.svc.runArgs = args
	*r = "OK"
	return rpc.svc.Server().Reset()
}

// Workers returns list of active workers and their stats.
func (rpc *rpcServer) Workers(list bool, r *WorkerList) (err error) {
	if rpc.svc == nil || rpc.svc.handler == nil {
		return errors.New("http server is not running")
	}

	r.Workers, err = util.ServerState(rpc.svc.Server())
	return err
}
