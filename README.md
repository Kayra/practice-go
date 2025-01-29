# practice-go

This is the code written while roughly following the [Go docs module creation tutorial](https://go.dev/doc/tutorial/create-module).

Technologies used:

- go: 1.23.5 darwin/arm64

## Run commands

Run the application:

```bash
cd hello; go run .
```

## Useful commands

Update hello module to use local greetings module:

```bash
go mod edit -replace example.com/greetings=../greetings
go mod tidy
```

Compile code into a binary executable:

```bash
cd hello; go build
```

Run the executable (unix):

```bash
./hello/hello
```

Discover go install path:

```bash
cd hello; go list -f '{{.Target}}'
```

Save install path to ENV (unix):

```bash
export PATH=$PATH:/path/to/your/install/directory
```

Compile and install package:

```bash
go install
```

Run package globally:

```bash
hello
```
