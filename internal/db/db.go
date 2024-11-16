package db

import (
	"sync"
)

type IDbOperation interface {
	IIdGenerateOperation
	IArticleDbOperation
	ITagDbOperation
	ITokenDbOperation
}

var (
	_db             IDbOperation
	_dbRegisterOnce = sync.Once{}
)

func InitDatabase() (err error) {
	_dbRegisterOnce.Do(func() {
		op, newErr := NewDbLogic()
		if newErr != nil {
			err = newErr
			return
		}
		_db = op
	})
	return
}

func Db() IDbOperation {
	return _db
}
