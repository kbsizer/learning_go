# GOPATH and Modules

## Motivation

Various tutorials and posts recommend defining GOPATH, not defining GOPATH, putting projects in GOPATH and not putting projects in GOPATH.  This confused/bugged me no end.

## Explanation

Go’s dependency management facilities have evolved a lot between 1.10 and 1.13. As a consequence, the best practices found in articles, posts and tutorials depends heavily on exactly *when* they were written.

Recent GOLANG versions:

| Version      | Dependency Management                                        |
| ------------ | ------------------------------------------------------------ |
| up to 1.10   | Projects are organized via a strict directory structure and the GOPATH environment variable |
| 1.11<br>1.12 | Added support for Go Modules (go.mod), but still defaulted to using GOPATH if current directory falls within $GOPATH/src. |
| 1.13         | Go Modules became default behavior (if there is a go.mod file present, then module semantics are used regardless of GOPATH) |

Using/advocating the use of modules during the 1.11-1.12 “transition period” gave rise to all the recommendations about where to put things relative to GOPATH, whether to use GO111MODULE, etc.

## Best Practice Recommendations (as of November 2019)

* Use Go version 1.13 or higher
* Don’t define/depend on GOPATH 
* Don’t define/depend on GO111MODULE
* Ensure that each project has a **go.mod** file.  Create one using `go mod init` 

### Commands for working with Go Modules

**`go mod init`** -- creates a new module, initializing the `go.mod` file that describes it.

**`go build`, `go test`,** *and other package-building commands* -- add new dependencies to `go.mod` as needed.****

**`go list -m all`** -- prints the current module’s dependencies.

**`go get`** -- changes the required version of a dependency (or adds a new dependency).

**`go mod tidy`** -- removes unused dependencies.

### Creating and using a go.mod file

```bash
$ cd ~/go/learning_go/modules_demo

$ go mod init modules_demo
go: creating new go.mod: module modules_demo

$ cat go.mod

module modules_demo

go 1.13

$ go test
PASS
ok      modules_demo    0.383s
```

### Updating a dependency version and introducing a breaking change

```bash
# Before update (all is goodness and light)
$ go test
PASS
ok      modules_demo    0.798s

$ go list -m all
modules_demo
golang.org/x/text v0.3.2
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0

# get newer version of a dependency
$ go get rsc.io/sampler
go: finding rsc.io/sampler v1.99.99
go: downloading rsc.io/sampler v1.99.99
go: extracting rsc.io/sampler v1.99.99

# rerun tests
$ go test
--- FAIL: TestHello (0.00s)
    hello_test.go:8: Hello() = "99 bottles of beer on the wall, 99 bottles of beer, ...", want "Hello, world."
FAIL
exit status 1
FAIL    modules_demo    0.372s

# Uh-oh!  Find an older version that will work better for us
$ go list -m -versions rsc.io/sampler
rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99

# Try v1.3.1
$ go get rsc.io/sampler@v1.3.1
go: finding rsc.io/sampler v1.3.1
go: downloading rsc.io/sampler v1.3.1
go: extracting rsc.io/sampler v1.3.1

$ go test                                 
PASS
ok      modules_demo    0.379s                # YAY! All happy again!
```

### Removing unused dependencies

Neither `go build` nor `go test` will remove unused modules.  For that, use `go mod tidy`:

```bash
$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote v1.5.2                   # we're no longer using this guy
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1

$ go mod tidy                         # exhaustively check all packages and remove unused modules

$ go list -m all
example.com/hello
golang.org/x/text v0.3.0
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.1
```



## For more information, see: https://blog.golang.org/using-go-modules

The example in this folder was taken directly from this article.