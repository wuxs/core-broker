package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/pkg/errors"
	"github.com/tkeel-io/kit/log"
	tkeelTransutil "github.com/tkeel-io/kit/transport/http"
)

const _XtKeelAuthUserHeader = "x-tKeel-auth"

var (
	ErrNotFound = errors.New("authorization info not found")
)

type User struct {
	ID       string `json:"id"`
	TenantID string `json:"tenant_id"`
	Token    string `json:"token"`
}

func GetUser(ctx context.Context) (User, error) {
	u := User{}

	log.Debugf("get header from ctx: %v", tkeelTransutil.HeaderFromContext(ctx))

	authHTTPHeader, ok := tkeelTransutil.HeaderFromContext(ctx)[_XtKeelAuthUserHeader]
	if !ok {
		return u, ErrNotFound
	}
	authInfo := strings.Join(authHTTPHeader, "")
	authStrBytes, err := base64.StdEncoding.DecodeString(authInfo)
	if err != nil {
		err = errors.Wrap(err, "decode auth header error")
		return u, err
	}
	log.Debugf("Print Decode Auth Info: %s", string(authStrBytes))
	err = json.Unmarshal(authStrBytes, &u)
	if err != nil {
		err = errors.Wrap(err, "unmarshal auth header error")
		return u, err
	}
	return u, nil
}
