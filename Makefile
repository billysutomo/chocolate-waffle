BINARY=engine
test: 
	go test -v -cover -covermode=atomic ./...

devweb:
	npm start --prefix web

devservice:
	go run main.go

mock-repo:
	mockgen -destination=internal/domain/mocks/element.go -package=mocks github.com/billysutomo/chocolate-waffle/internal/domain ElementRepository
	mockgen -destination=internal/domain/mocks/project.go -package=mocks github.com/billysutomo/chocolate-waffle/internal/domain ProjectRepository