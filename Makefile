#! /usr/bin/make -f

mockgen:
	mockgen -destination mock/tcb_mock.go -package mock github.com/deadcheat/tcb CouchBaseAdapter,Configurer,Operator,Logger
