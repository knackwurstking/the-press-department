all: build

init:
	npm install

build:
	npx tsc --outDir public/ src/sw.ts
	npx vite build --config vite.config.ts

build-android:
	make build
	npx cap sync android
