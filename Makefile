build:
	@cd network/client && go build -ldflags="-s -w" . && GOOS=windows GOOS=windows go build -ldflags "-H windowsgui -s -w" .
	@go build .
	@echo Success