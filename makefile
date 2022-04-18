all: i386-linux amd64-linux arm-linux arm64-linux i386-windows amd64-windows amd64-darwin arm64-darwin

i386-linux:
	GOARCH=386 GOOS=linux go build -ldflags="-s -w" -o prg/i386/linux/doc src/doc.go

amd64-linux:
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o prg/amd64/linux/doc src/doc.go

arm-linux:
	GOARCH=arm GOOS=linux go build -ldflags="-s -w" -o prg/arm/linux/doc src/doc.go

arm64-linux:
	GOARCH=arm64 GOOS=linux go build -ldflags="-s -w" -o prg/arm64/linux/doc src/doc.go

i386-windows:
	GOARCH=386 GOOS=windows go build -ldflags="-s -w" -o prg/i386/windows/doc.exe src/doc.go

amd64-windows:
	GOARCH=amd64 GOOS=windows go build -ldflags="-s -w" -o prg/amd64/windows/doc.exe src/doc.go

amd64-darwin:
	GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o prg/amd64/darwin/doc src/doc.go

arm64-darwin:
	GOARCH=arm64 GOOS=darwin go build -ldflags="-s -w" -o prg/arm64/darwin/doc src/doc.go
