package client

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/mkusaka/go-circleci-api/pkg/utils"
)

type CustomPolicy struct {
}

var _ policy.Policy = (*CustomPolicy)(nil)

func (c CustomPolicy) Do(req *policy.Request) (*http.Response, error) {
	req.Raw().Header.Set("authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(utils.CircleToken))))
	fmt.Println("auth header length", len(req.Raw().Header.Get("authorization")))
	return req.Next()
}
