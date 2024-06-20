// Code generated by genpjrpc. DO NOT EDIT.
//  genpjrpc version: v0.2.0

package generated

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.com/pjrpc/pjrpc/v2/client"

	"henry.com/rinatusmanov/jsonrpc20/internal/pkg/types"
)

// List of the client JSON-RPC methods.
const (
	JSONRPCMethodGetBalance_Client          = "getBalance"
	JSONRPCMethodRollbackTransaction_Client = "withdrawAndDeposit"
	JSONRPCMethodWithdrawAndDeposit_Client  = "rollbackTransaction"
)

// SeamlessV2ServiceClient is an API client for SeamlessV2Service service.
type SeamlessV2ServiceClient interface {
	GetBalance(ctx context.Context, in *types.GetBalanceRequest, mods ...client.Mod) (*types.GetBalanceResponse, error)
	RollbackTransaction(ctx context.Context, in *types.RollbackTransactionRequest, mods ...client.Mod) (*types.RollbackTransactionResponse, error)
	WithdrawAndDeposit(ctx context.Context, in *types.WithdrawAndDepositRequest, mods ...client.Mod) (*types.WithdrawAndDepositResponse, error)
}

type implSeamlessV2ServiceClient struct {
	cl client.Invoker
}

// NewSeamlessV2ServiceClient returns new client implementation of the SeamlessV2Service service.
func NewSeamlessV2ServiceClient(cl client.Invoker) SeamlessV2ServiceClient {
	return &implSeamlessV2ServiceClient{cl: cl}
}

func (c *implSeamlessV2ServiceClient) GetBalance(ctx context.Context, in *types.GetBalanceRequest, mods ...client.Mod) (result *types.GetBalanceResponse, err error) {
	gen, err := uuid.NewUUID()
	if err != nil {
		return result, fmt.Errorf("failed to create uuid generator: %w", err)
	}

	err = c.cl.Invoke(ctx, gen.String(), JSONRPCMethodGetBalance_Client, in, result, mods...)
	if err != nil {
		return result, fmt.Errorf("failed to Invoke method %q: %w", JSONRPCMethodGetBalance_Client, err)
	}

	return result, nil
}

func (c *implSeamlessV2ServiceClient) RollbackTransaction(ctx context.Context, in *types.RollbackTransactionRequest, mods ...client.Mod) (result *types.RollbackTransactionResponse, err error) {
	gen, err := uuid.NewUUID()
	if err != nil {
		return result, fmt.Errorf("failed to create uuid generator: %w", err)
	}

	err = c.cl.Invoke(ctx, gen.String(), JSONRPCMethodRollbackTransaction_Client, in, result, mods...)
	if err != nil {
		return result, fmt.Errorf("failed to Invoke method %q: %w", JSONRPCMethodRollbackTransaction_Client, err)
	}

	return result, nil
}

func (c *implSeamlessV2ServiceClient) WithdrawAndDeposit(ctx context.Context, in *types.WithdrawAndDepositRequest, mods ...client.Mod) (result *types.WithdrawAndDepositResponse, err error) {
	gen, err := uuid.NewUUID()
	if err != nil {
		return result, fmt.Errorf("failed to create uuid generator: %w", err)
	}

	err = c.cl.Invoke(ctx, gen.String(), JSONRPCMethodWithdrawAndDeposit_Client, in, result, mods...)
	if err != nil {
		return result, fmt.Errorf("failed to Invoke method %q: %w", JSONRPCMethodWithdrawAndDeposit_Client, err)
	}

	return result, nil
}