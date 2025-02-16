package conversation

import (
	"context"
)

type Device interface {
	GetState() []byte
	SetState(state []byte)
	GetHost() string
	RefreshToken(ctx context.Context) error
	GetToken() string
	GetCertificate() string
}
