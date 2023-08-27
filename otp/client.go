package otp

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/imroc/req/v3"
)

type Client struct {
	*req.Client
}

func NewClient(config Config) (client *Client) {
	client = &Client{
		Client: req.C().
			SetBaseURL(config.baseUrl).
			SetCommonErrorResult(&Error{}).
			SetUserAgent("AkouendyPay").
			SetCommonBasicAuth(config.application, config.secret).
			EnableDumpEachRequest().
			OnAfterResponse(func(client *req.Client, resp *req.Response) error {
				if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
					return nil
				}
				if errMsg, ok := resp.ErrorResult().(*Error); ok {
					resp.Err = errMsg // Convert api error into go error
					return nil
				}
				if !resp.IsSuccessState() {
					// Neither a success response nor a error response, record details to help troubleshooting
					resp.Err = fmt.Errorf("bad status: %s\nraw content:\n%s", resp.Status, resp.Dump())
				}
				return nil
			}),
	}
	if config.devMode {
		client.DevMode()
	}
	return client
}

func (c *Client) ListProviders(ctx context.Context) (response Response, err error) {
	_, err = c.R().
		SetContext(ctx).
		SetSuccessResult(&response).
		Get("/api/providers")

	return
}

func (c *Client) Init(ctx context.Context, init InitRequest) (response Response, err error) {
	data := make(map[string]string)
	data["to"] = init.Receiver
	data["provider"] = init.Provider
	_, err = c.R().
		SetContext(ctx).
		SetPathParam("id", init.ID.String()).
		SetFormData(data).
		SetSuccessResult(&response).
		Put("/api/otp/{id}")
	return
}

func (c *Client) Validate(ctx context.Context, validate ValidateRequest) (response Response, err error) {
	data := make(map[string]string)
	data["action"] = "check"
	data["otp"] = validate.Code
	data["skip_delete"] = strconv.FormatBool(validate.SkipDelete)
	_, err = c.R().
		SetContext(ctx).
		SetPathParam("id", validate.ID.String()).
		SetFormData(data).
		SetSuccessResult(&response).
		Post("/api/otp/{id}")
	return
}

func (c *Client) Check(ctx context.Context, ID uuid.UUID) (response Response, err error) {
	_, err = c.R().
		SetContext(ctx).
		SetPathParam("id", ID.String()).
		SetSuccessResult(&response).
		Post("/api/otp/{id}/status")
	return
}
