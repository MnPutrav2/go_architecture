package middleware

import (
	"context"

	jwtEnc "github.com/MnPutrav2/go_architecture/pkg/auth/jwt"
)

func UserClaims(ctx context.Context) (*jwtEnc.Claims, bool) {
	ct, ok := ctx.Value(UserClaimsKey).(*jwtEnc.Claims)
	return ct, ok
}
