# run the example usage
[group('pub')]
run:
    go run cmd/main.go

# build the example usage 
[group('pub')]
build-ex:
    @go build cmd/main.go

# run the tests
[group('dev')]
test:
    go test ./test/.
    @echo 'You probably need to make more tests'
