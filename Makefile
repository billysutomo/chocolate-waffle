BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

devweb:
	npm start --prefix web