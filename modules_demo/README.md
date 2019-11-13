# GOPATH and Modules

##  Motivation

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
* Ensure that each project has a **go.mod** file.  Create one using `go mod init` as follows:

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



## For more information, see: https://blog.golang.org/using-go-modules

The example in this folder was taken directly from this article.