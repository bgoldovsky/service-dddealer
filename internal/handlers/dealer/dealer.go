package dealer

import (
	"context"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order/item"

	"github.com/bgoldovsky/dealer/service-dealer/internal/generated/rpc"

	"github.com/bgoldovsky/dealer/service-dealer/internal/app/models/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	service order.Service
	rpc.UnimplementedDealerServer
}

func New(service order.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateOrder(ctx context.Context, req *rpc.CreateOrderRequest) (*rpc.CreateOrderReply, error) {
	if len(req.Items) == 0 {
		return nil, status.Error(codes.InvalidArgument, "items not specified")
	}

	i, err := toItems(req.Items)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reply, err := h.service.Create(ctx, req.Workflow, i)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &rpc.CreateOrderReply{
		Id:      int64(reply.ID),
		Created: reply.Created.Unix(),
	}, nil
}

func (h *Handler) ApplyTransition(ctx context.Context, req *rpc.ApplyTransitionRequest) (*rpc.ApplyTransitionReply, error) {
	if req.Id == 0 {
		return nil, status.Error(codes.InvalidArgument, "items id not specified")
	}

	err := h.service.ApplyTransition(ctx, order.ID(req.Id), req.Transition)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &rpc.ApplyTransitionReply{Success: true}, nil
}

func toItems(req []*rpc.ItemCreateRequest) ([]item.Item, error) {
	items := make([]item.Item, len(req))

	for idx, val := range req {
		var discount *int64
		if val.Discount != nil {
			discount = &val.Discount.Value
		}

		i, err := item.New(item.ID(val.Id), val.SellerId, val.Price, discount, val.Name)
		if err != nil {
			return nil, err
		}

		items[idx] = *i
	}

	return items, nil
}

// &wrappers.StringValue{
