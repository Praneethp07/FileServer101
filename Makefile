ENCRYPTION_KEY="8a3f5c9d1b7e4f2a6c90d8b3e7f2410e"


run:
	@echo "running app"
	@ENCRYPTION_KEY=$(ENCRYPTION_KEY) go run main.go