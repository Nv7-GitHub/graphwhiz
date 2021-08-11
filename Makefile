install:
	go install fyne.io/fyne/v2/cmd/fyne@latest

build:
	go build -ldflags="-s -w" -o graphwhiz
ifeq "$(OS)" "Windows_NT"
	fyne package -os darwin -executable graphwhiz
else
ifeq "$(shell uname -s)" "Linux"
	fyne package -os linux -executable graphwhiz
endif

ifeq "$(shell uname -s)" "Darwin"
	fyne package -os darwin -executable graphwhiz
endif
endif

	rm graphwhiz