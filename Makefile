APP = samsamoohooh

run:
	@go run cmd/samsamoohooh/samsamoohooh.go

swag:
	@go install github.com/swaggo/swag/cmd/swag@v1.16.4
	@swag init -g cmd/samsamoohooh/samsamoohooh.go --output ./api

package:
	@go mod tidy

lint:
	@golangci-lint run
