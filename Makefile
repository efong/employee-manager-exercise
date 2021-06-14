all:
	go build -o testapp
	go test -count=1
