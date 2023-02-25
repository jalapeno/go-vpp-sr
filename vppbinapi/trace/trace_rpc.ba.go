// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package trace

import (
	"context"

	api "git.fd.io/govpp.git/api"
)

// RPCService defines RPC service trace.
type RPCService interface {
	TraceProfileAdd(ctx context.Context, in *TraceProfileAdd) (*TraceProfileAddReply, error)
	TraceProfileDel(ctx context.Context, in *TraceProfileDel) (*TraceProfileDelReply, error)
	TraceProfileShowConfig(ctx context.Context, in *TraceProfileShowConfig) (*TraceProfileShowConfigReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) TraceProfileAdd(ctx context.Context, in *TraceProfileAdd) (*TraceProfileAddReply, error) {
	out := new(TraceProfileAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceProfileDel(ctx context.Context, in *TraceProfileDel) (*TraceProfileDelReply, error) {
	out := new(TraceProfileDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) TraceProfileShowConfig(ctx context.Context, in *TraceProfileShowConfig) (*TraceProfileShowConfigReply, error) {
	out := new(TraceProfileShowConfigReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}