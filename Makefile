build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/hooknode hooknode/main.go

deploy:
	make build
	serverless deploy -v

deploy-prod:
	make build
	serverless deploy -v  --stage production
