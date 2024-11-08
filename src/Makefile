.PHONY: run
run:
	templ generate
	go run cmd/web/main.go | jq .

dev:
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run cmd/web/main.go"