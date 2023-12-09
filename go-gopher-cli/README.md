### Cobra
Cobra is both a library for creating powerful modern CLI
applications and a program for generating applications and
batch files.

```bash
# Install the Cobra librasy and its dependencies
go get -u github.com/spf13/cobra@latest

# Install the Cobra CLI
go install github.com/spf13/cobra-cli@latest
```
Generate the CLI application

```bash
cobra-cli init
```

```bash
# Install the Viper dependency
go get github.com/spf13/viper@v1.8.1
```

In order to explain the goal of the ClI, we want to display:
  - a short description
  - a long description
  - using our app
by modifiyng the `cmd/root.go` file.

.
├── LICENSE
├── bin
├── cmd
│   └── root.go
├── go.mod
├── go.sum
└── main.go
