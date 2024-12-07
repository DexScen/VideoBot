package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	ssov1 "github.com/j0n1que/sso-protos/gen/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Client struct {
	api ssov1.AuthClient
	log *slog.Logger
}

func New(ctx context.Context, log *slog.Logger, addr string, timeout time.Duration, retriesCount int) (*Client, error) {
	const op = "grpc.New"

	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithMax(uint(retriesCount)),
		grpcretry.WithPerRetryTimeout(timeout),
	}

	/*logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}*/

	cc, err := grpc.DialContext(ctx, addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			//grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...),
		))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	grpcClient := ssov1.NewAuthClient(cc)

	return &Client{
		api: grpcClient,
		log: log,
	}, nil
}

func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

func (c *Client) RegisterNewUser(ctx context.Context, login, password, telegramLogin string) error {
	const op = "grpc.RegisterNewUser"

	_, err := c.api.RegisterNewUser(ctx, &ssov1.RegisterRequest{
		Login:         login,
		Password:      password,
		TelegramLogin: telegramLogin,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Client) AuthorizeUser(ctx context.Context, login, password string) (string, error) {
	const op = "grpc.AuthorizeUser"

	resp, err := c.api.AuthorizeUser(ctx, &ssov1.AutohrizeRequest{
		Login:    login,
		Password: password,
	})

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.Token, nil
}

func (c *Client) IsAdmin(ctx context.Context, userID int64) (bool, error) {
	const op = "grpc.IsAdmin"

	resp, err := c.api.IsAdmin(ctx, &ssov1.IsAdminRequest{
		UserId: userID,
	})

	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return resp.IsAdmin, nil
}

func (c *Client) ChangePassword(ctx context.Context, userID int64, newPassword string) error {
	const op = "grpc.ChangePassword"

	_, err := c.api.ChangePassword(ctx, &ssov1.ChangePasswordRequest{
		UserId:      userID,
		NewPassword: newPassword,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Client) GetAllUsers(ctx context.Context) ([]*ssov1.User, error) {
	const op = "grpc.GetAllUsers"

	resp, err := c.api.GetAllUsers(ctx, &emptypb.Empty{})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp.Users, nil
}

func (c *Client) GetUserByTelegram(ctx context.Context, telegramLogin string) ([]*ssov1.User, error) {
	const op = "grpc.GetUserByTelegram"

	resp, err := c.api.GetUserByTelegram(ctx, &ssov1.GetUserByTelegramRequest{
		TelegramLogin: telegramLogin,
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return resp.Users, nil
}

func (c *Client) MakeAdmin(ctx context.Context, userID int64) error {
	const op = "grpc.MakeAdmin"

	_, err := c.api.MakeAdmin(ctx, &ssov1.MakeAdminRequest{
		UserId: userID,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (c *Client) GetJWT(ctx context.Context, userID int64) (string, error) {
	const op = "grpc.GetJWT"

	resp, err := c.api.GetJWT(ctx, &ssov1.GetJWTRequest{
		UserId: userID,
	})

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.Token, nil
}

func (c *Client) DeleteJWT(ctx context.Context, userID int64) error {
	const op = "grpc.DeleteJWT"

	_, err := c.api.DeleteJWT(ctx, &ssov1.DeleteJWTRequest{
		UserId: userID,
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
