
include .env
test:
	mgo_url=${mgo_url}  go  test *.go