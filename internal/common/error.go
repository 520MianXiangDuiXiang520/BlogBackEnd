package common

import (
	"strconv"
)

type ErrorCode int

func (e ErrorCode) Error() string {
	return strconv.Itoa(int(e))
}

// server error

const (
	SrvErrorUnKnow ErrorCode = iota + 50000
	SrvErrorDb
	SrvErrorCfg
)

// user error

const (
	UserErrorUnKnow ErrorCode = iota + 20000
	UserErrorBadPasswordOrUserName
)
