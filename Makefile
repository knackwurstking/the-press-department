all: init build

APP_NAME := the-press-department

clean:
	git clean -fxd

init:
	cd ui && npm install

run:
	go run -v ./cmd/${APP_NAME}

build:
	env GOOS=js GOARCH=wasm go build -o ui/public/${APP_NAME}.wasm ./cmd/${APP_NAME}
	cd ui && make build
	go build -o ./bin/${APP_NAME} ./cmd/${APP_NAME}

build-android:
	cd ui && make build && make build-android
	
# NOTE: Standard rpi-server-project part

define SYSTEMD_SERVICE_FILE
[Unit]
Description=A interactive screensaver. No just for fun.
After=network.target

[Service]
ExecStart=${APP_NAME}

[Install]
WantedBy=default.target
endef

UNAME := $(shell uname)
check-linux:
ifneq ($(UNAME), Linux)
	@echo 'This won’t work here since you’re not on Linux.'
	@exit 1
endif

export SYSTEMD_SERVICE_FILE
install: check-linux
	echo "$$SYSTEMD_SERVICE_FILE" > ${HOME}/.config/systemd/user/${APP_NAME}.service 
	systemctl --user daemon-reload 
	echo "--> Created a service file @ ${HOME}/.config/systemd/user/${APP_NAME}.service"
	sudo cp ./bin/${APP_NAME} /usr/local/bin/

start: check-linux
	systemctl --user restart ${APP_NAME}

stop: check-linux
	systemctl --user stop ${APP_NAME}

log: check-linux
	journalctl --user -u ${APP_NAME} --follow --output cat
