package service

import (
	"github.com/deta/deta-go/deta"
)

var Deta *deta.Deta

func init() {
	var err error
	Deta, err = deta.New()
	if err != nil {
		panic(err)
	}
}
