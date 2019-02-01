# GUI Notes

Keep a lookout for https://www.packtpub.com/application-development/hands-gui-application-development-go

https://github.com/fyne-io/fyne/wiki/Cross-Compiling:

* brew install mingw-w64
* which x86_64-w64-mingw32-gcc
* export CC=/usr/local/bin/x86_64-w64-mingw32-gcc
* CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build

See also https://github.com/fyne-io/fyne/wiki/Packaging