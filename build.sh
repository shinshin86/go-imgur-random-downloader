#!/bin/sh

fname="go-imgur-random-downloader"

GOOS=linux GOARCH=amd64 go build -o ./bin/linux64/${fname}
GOOS=linux GOARCH=386 go build -o ./bin/linux386/${fname}

GOOS=windows GOARCH=386 go build -o ./bin/windows386/${fname}.exe
GOOS=windows GOARCH=amd64 go build -o ./bin/windows64/${fname}.exe

GOOS=darwin GOARCH=386 go build -o ./bin/darwin386/${fname}
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin64/${fname}

echo "==========> Build successful"

