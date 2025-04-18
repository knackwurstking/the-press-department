all: build

init:
	npm install

build:
	npx vite build --config vite.main.config.ts
	npx vite build --config vite.sw.config.ts

build-android:
	make build
	npx cap sync android
