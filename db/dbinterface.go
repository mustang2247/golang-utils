package db


type DBBase interface {
	Open(args ...interface{})
}
