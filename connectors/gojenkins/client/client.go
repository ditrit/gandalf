package client

import (
	"context"

	"github.com/bndr/gojenkins"
)

func ClientWithAuthentication(addr, login, password string) *gojenkins.Jenkins {
	ctx := context.Background()
	jenkins, _ := gojenkins.CreateJenkins(nil, addr, login, password).Init(ctx)

	return jenkins
}

func ClientWithoutAuthentication(addr string) *gojenkins.Jenkins {
	ctx := context.Background()
	jenkins, _ := gojenkins.CreateJenkins(nil, addr).Init(ctx)

	return jenkins
}

