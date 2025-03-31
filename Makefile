APP = samsamoohooh

swag:
	@swag init -g cmd/samsamoohooh/samsamoohooh.go --output ./api

run: 
	@go run cmd/samsamoohooh/samsamoohooh.go