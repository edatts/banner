set positional-arguments

build:
    go build -o bin/banner ./cmd/...

run +args="Hello World": build
    ./bin/banner --text "$@"