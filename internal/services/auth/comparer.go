package auth

import "context"

type Comparer interface {
	Compare(ctx context.Context, hashed string, plaintext string) error
}
