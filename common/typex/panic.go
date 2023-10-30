package typex

import "go-zero-bookstore/common/logx"

func MustNil(err error) {
	if err != nil {
		logx.Errorf("MustNil panic: %s", err.Error())
		panic(err)
	}
}
