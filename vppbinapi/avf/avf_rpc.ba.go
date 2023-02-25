// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package avf

import (
	"context"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service avf.
type RPCService interface {
	AvfCreate(ctx context.Context, in *AvfCreate) (*AvfCreateReply, error)
	AvfDelete(ctx context.Context, in *AvfDelete) (*AvfDeleteReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) AvfCreate(ctx context.Context, in *AvfCreate) (*AvfCreateReply, error) {
	out := new(AvfCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) AvfDelete(ctx context.Context, in *AvfDelete) (*AvfDeleteReply, error) {
	out := new(AvfDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}