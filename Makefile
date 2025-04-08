clean:
	git clean -fxd

init:
	@cd www && npm install

run:
	@go run -v ./cmd/the-press-department

build-wasm:
	@env GOOS=js GOARCH=wasm go build -o www/public/the-press-department.wasm ./cmd/the-press-department
	@cd www && npm run build && npx cap sync android
