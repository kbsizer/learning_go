# Building an Awesome CPI App in Go (OSCON 2017)

## Slides: https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/


## Local Path: `/c/Go_Projects/learning_go/cli_building_an_awesome_cli_OSCON2017`


## COBRA

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

```
Usage:
  cobra [command]

Available Commands:
  add         Add a command to a Cobra Application
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize a Cobra Application

Flags:
  -a, --author string    author name for copyright attribution (default "YOUR NAME")
      --config string    config file (default is $HOME/.cobra.yaml)
  -h, --help             help for cobra
  -l, --license string   name of license for the project
      --viper            use Viper for configuration (default true)

Use "cobra [command] --help" for more information about a command.
```



## Notes

### Demo uses GOPATH, not modules

### Deprecation notice from GOLANG tooling:

> installing executables with 'go get' in module mode is deprecated.
> To adjust and download dependencies of the current module, use 'go get -d'.
> To install using requirements of the current module, use 'go install'.
> To install ignoring the current module, use 'go install' with a version, like 'go install example.com/cmd@latest'.
> For more information, see https://golang.org/doc/go-get-install-deprecation or run 'go help get' or 'go help install'.

### Initializing COBRA

```
$ cobra init --pkg-name awesome.cli
Your Cobra application is ready at
C:\Go_Projects\learning_go\cli_building_an_awesome_cli_OSCON2017
```

### Constructors

GOLANG does not have constructors. If construction (i.e., initialization prior to use) is required, use a factory.

The convention is to have a function named `New_____()`

### Composite Literals

An expression that creates a new value each time it is evaluated.

Example: `[]todo.Item{}`

*slide 203*
