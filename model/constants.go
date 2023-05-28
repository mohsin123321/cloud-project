package model

// mime types of the responses and requests
const (
	MimeTypeJSON = "application/json"
)

const TokenHeader string = "X-Auth-Token"

// CtxKey represents a key (string) for retrieving struct saved in request context.
type CtxKey struct {
	Key string
}

func GetCtxKeyID() CtxKey {
	return CtxKey{Key: "id"}
}
