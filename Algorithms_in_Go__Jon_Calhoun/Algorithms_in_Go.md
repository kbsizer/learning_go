# Ken's Notes for *Algorithms in Go by Jon Calhoun*

[TOC]



## Running specific tests

Using straight up go tooling...
```bash
$ cd /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun/module01 (master)
$ go test -run=NumInList
PASS
ok      algo/module01   0.515s
```
Using the cute little colorizer...
```bash
$ gotest -run=NumInList
PASS
ok      algo/module01   0.511s

```





## Appendix: How to resolve `Cannot find module for path .`

### Error message

```bash
$ go test
build .: cannot find module for path .
```

### What it means

Since no package (directory) was specified, the tooling assumes the current working directory (thus the dot) is where it should look for test files.  In the case above, however, all `.go` files were in subdirectories *under* the current directory and the build tool found nothing to build.

***Resolution***

Either `cd` into one of the subdirectories with `.go` files or specify the subdirectory(ies) containing `.go` files you want to build (or test).

Example:

```bash
# option 1
$ cd module01
$ go test

# option 2
$ go test ./module01
```



### Some history

| &nbsp;&nbsp;Go&nbsp;Version&nbsp;&nbsp; | Dependency Management                                        |
| :-------------------------------------: | ------------------------------------------------------------ |
|               before 1.11               | `GOPATH` centralizes all packages into one common workspace; every project is a package under `$GOPATH/src`<br />(`GOPATH` should not be confused with `GOROOT`, which should point to the standard packages, programmers should not install anything in `GOROOT`, modify standard packages, etc.) |
|                  1.11                   | Module support introduced: [Modules](https://golang.org/doc/go1.11#modules) are defined as a collection of [packages](https://golang.org/ref/spec#Packages) stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module’s *module path*, which is also the import path used for the root directory, and its *dependency requirements*, which are the other modules needed for a successful build.<br>Behavior still defaults to using GOPATH if the working directory is within $GOPATH/src. [Modules](https://golang.org/doc/go1.11#modules) begin the journey to becoming Go’s [new dependency management system](https://blog.golang.org/versioning-proposal). |
|               1.13 and up               | Starting in Go 1.13, module mode becomes the default for all development. |

### Experimenting

The `go.mod` file when project first cloned

```
$ cat go.mod
module algo

go 1.13

require (
        github.com/fatih/color v1.9.0 // indirect
        github.com/rakyll/gotest v0.0.0-20191108192113-45d501058f2a // indirect
        golang.org/x/sys v0.0.0-20200107162124-548cf772de50 // indirect
)
```
Tidy things up and add test dependencies
```
Ken@KEN-HP-ENVY MINGW64 /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun (master)
$ go mod tidy
```
Resulting `go.mod`
```
Ken@KEN-HP-ENVY MINGW64 /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun (master)
$ cat go.mod
module algo

go 1.13
```

Not exactly what I was expecting.  

### Further Reading

[The Go Blog: Using Go Modules](https://blog.golang.org/using-go-modules)

[Go without GOPATH (introduction to modules)](https://dev.to/dizdarevic/golang-without-a-path-47jp)

## Appendix: How to resolve `go get: no install location`

### Error Message

```bash
$ go get -t
go get: no install location for directory C:\Go_Projects\learning_go\gin_web_server_module_demo outside GOPATH
        For more details see: 'go help gopath'
```

or

```bash
$ go get -u
package _/C_/Go_Projects/learning_go/gin_web_server_module_demo: unrecognized import path "_/C_/Go_Projects/learning_go/gin_web_server_module_demo" (import path does not begin with hostname)

```

Note that `go build` succeeds:

```bash
$ go build

$ ll
total 14630
-rwxr-xr-x 1 Ken 197609 14975488 Feb 12 23:52 gin_web_server_module_demo.exe*
-rw-r--r-- 1 Ken 197609      296 Feb 12 21:55 main.go
-rw-r--r-- 1 Ken 197609      215 Feb 12 23:32 README.md
```

### Resolution

Need to initialize the module:

```bash
$ go mod init learning/gin_web_server
go: creating new go.mod: module learning/gin_web_server

$ ll
total 14631
-rw-r--r-- 1 Ken 197609       40 Feb 12 23:57 go.mod                # <== New "go.mod" file created
-rw-r--r-- 1 Ken 197609      296 Feb 12 21:55 main.go
-rw-r--r-- 1 Ken 197609      215 Feb 12 23:32 README.md

$ cat go.mod                                                        # <== Inspecting contents of "go.mod"
module learning/gin_web_server

go 1.13

```

Now retry...

```bash
$ go get -u
go: downloading github.com/json-iterator/go v1.1.7
go: downloading golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a
go: extracting github.com/json-iterator/go v1.1.7
       :
       :
go: extracting github.com/leodido/go-urn v1.2.0
go: downloading github.com/go-playground/locales v0.13.0
go: extracting github.com/go-playground/locales v0.13.0
```



### More Information

***From the `get` docs***

**-t** instructs get to also download the packages required to build the tests for the specified packages.

**-u** instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

**-v**  enables verbose progress and debug output.

**GOPATH** When checking out a new package, get creates the target directory `GOPATH/src/<import-path>`. If the GOPATH contains multiple entries, get uses the first one. (For more details see: 'go help gopath'.)

Note: When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go. The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1". If no such version exists it retrieves the default branch of the package.

## Appendix: How to resolve `path ... is not a package in module rooted at ...`

### Error Message

```bash
$ go get -u
go get .: path C:\Go_Projects\learning_go\Algorithms_in_Go__Jon_Calhoun is not a package in module rooted at C:\Go_Projects\learning_go\Algorithms_in_Go__Jon_Calhoun
```

### Resolution

Use either `...` or `all`  to tell go to find all required packages recursively 

```bash
 $ go get -u -v ...
go: downloading github.com/rakyll/gotest v0.0.0-20191108192113-45d501058f2a
go: extracting github.com/rakyll/gotest v0.0.0-20191108192113-45d501058f2a
go: downloading golang.org/x/sys v0.0.0-20200107162124-548cf772de50
      :
      :
golang.org/x/sys/windows/svc/mgr
golang.org/x/sys/windows/svc/example
```

