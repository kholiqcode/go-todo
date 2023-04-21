package token

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/kholiqcode/go-todolist/utils"
)

type key string

const TOKEN_PAYLOAD key = "token-payload"

func AppendRequestCtx(r *http.Request, ctxKey key, input interface{}) context.Context {
	return context.WithValue(r.Context(), ctxKey, input)
}

func GetRequestCtx(r *http.Request, ctxKey key) *Payload {
	return r.Context().Value(ctxKey).(*Payload)
}

func CheckIsAuthorize(r *http.Request, accessId uuid.UUID) {
	tokenPayload := GetRequestCtx(r, TOKEN_PAYLOAD)

	if tokenPayload.UserId != accessId {
		utils.PanicIfError(utils.CustomError("not authorize to perform this operation", 403))
	}
}
