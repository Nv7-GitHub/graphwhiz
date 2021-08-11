install:
	go install fyne.io/fyne/v2/cmd/fyne@latest
	go install github.com/fyne-io/fyne-cross@latest

build:
	go build -o graphwhiz
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

# Run "make cross" from sudo -E bash
cross:
	mkdir -p dist

	fyne-cross linux -app-id com.nv.graphwhiz
	mv fyne-cross/dist/linux-amd64/graphwhiz.tar.gz dist/graphwhiz_linux.tar.gz

	fyne-cross windows -app-id com.nv.graphwhiz
	mv fyne-cross/dist/windows-amd64/graphwhiz.exe.zip dist/graphwhiz_windows.exe.zip

# Due to license restrictions, can't cross-compile for MacOS