package mercadopago

import (
	"context"
	"errors"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/preapproval"
	"github.com/mercadopago/sdk-go/pkg/preapprovalplan"
)

var ErrEmptyClient = errors.New("failed to initialize mp client")

type client struct {
	pap preapprovalplan.Client
	pa  preapproval.Client
}

type Client interface {
	CreatePlan(context.Context, preapprovalplan.Request) (string, error)
	GetPlan(ctx context.Context, id string) (*preapprovalplan.Response, error)
	CreateSubscription(ctx context.Context, req preapproval.Request) (*preapproval.Response, error)
	GetSubscriptionById(ctx context.Context, id string) (*preapproval.Response, error)
	UpdateSubscription(ctx context.Context, id string, status string) error
}

func NewClient(cfg *config.Config) Client {
	return &client{
		pap: preapprovalplan.NewClient(cfg),
		pa:  preapproval.NewClient(cfg),
	}
}

func (c client) CreatePlan(ctx context.Context, req preapprovalplan.Request) (string, error) {
	res, err := c.pap.Create(ctx, req)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

func (c client) GetPlan(ctx context.Context, id string) (*preapprovalplan.Response, error) {
	res, err := c.pap.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c client) CreateSubscription(ctx context.Context, req preapproval.Request) (*preapproval.Response, error) {
	res, err := c.pa.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c client) GetSubscriptionById(ctx context.Context, id string) (*preapproval.Response, error) {
	res, err := c.pa.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c client) UpdateSubscription(ctx context.Context, id string, status string) error {
	_, err := c.pa.Update(ctx, id, preapproval.UpdateRequest{Status: status})
	if err != nil {
		return err
	}
	return nil
}
