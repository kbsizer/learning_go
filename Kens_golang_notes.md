# Ken's GOLANG Notes

[TOC]



## General Info on Running Tests

### Using straight up go tooling...

```bash
$ cd /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun/module01 (master)

$ go test -run=NumInList
PASS
ok      algo/module01   0.515s
```

Note: the `-run` command takes a regular expression as its argument.  More `-run` examples...

```bash
$ go test -run Foo                    # Run top-level tests matching "Foo", such as "TestFooBar".
$ go test -run TestBubbleSortInt$     # Run "TestBubbleSortInt" but not "TestBubbleSortInterface"
$ go test -run Foo/A=                 # For top-level tests matching "Foo", run subtests matching "A=".
$ go test -run /A=1                   # For all top-level tests, run subtests matching "A=1".
```

Listing test methods...

```bash
$ go test -list .
Test_main
TestGetHello
TestHypotenuse
TestAverage
ok      _/C_/Go_Projects/learning_go/testing
```

Notes

* `go test` recompiles each package along with any files with names matching the file pattern `*_test.go`
* The `*_test.go` files can contain test functions, benchmark functions, and example functions. See `go help testfunc` for more.
* Each listed package causes the execution of a separate test binary.
* Files whose names begin with `_` (including `_test.go`) or `.` are ignored.
* Test files that declare a package with the suffix `_test` will be compiled as a separate package, and then linked and run with the main test binary.
* `go` will ignore a directory named `testdata`, making it available to hold ancillary data needed by the tests.

For more info

```bash
$ go help test
```

For more details, see: https://golang.org/pkg/testing/

and https://ieftimov.com/post/testing-in-go-go-test/

### Obtaining and using the `gotest` test colorizer

```bash
$ cd /c/Go_Projects/learning_go/testing
$ gotest
bash: gotest: command not found

$ go get -u github.com/rakyll/gotest     # download/install the 'gotest' tool

$ gotest                                 # run all tests found in current folder; show only failures

$ gotest -v                              # show successful tests in green; failing tests in red

$ gotest -run NumInList                  # run only tests whose name contains "NumInList"
```



## Common Golang Environment Problems and Solutions

###  `Cannot find module for path .`

#### Error message

```bash
$ go test
build .: cannot find module for path .
```

#### What it means

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

#### Some history

| &nbsp;&nbsp;Go&nbsp;Version&nbsp;&nbsp; | Dependency Management                                        |
| :-------------------------------------: | ------------------------------------------------------------ |
|               before 1.11               | `GOPATH` centralizes all packages into one common workspace; every project is a package under `$GOPATH/src`<br />(`GOPATH` should not be confused with `GOROOT`, which should point to the standard packages, programmers should not install anything in `GOROOT`, modify standard packages, etc.) |
|                  1.11                   | Module support introduced: [Modules](https://golang.org/doc/go1.11#modules) are defined as a collection of [packages](https://golang.org/ref/spec#Packages) stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module’s *module path*, which is also the import path used for the root directory, and its *dependency requirements*, which are the other modules needed for a successful build.<br>Behavior still defaults to using GOPATH if the working directory is within $GOPATH/src. [Modules](https://golang.org/doc/go1.11#modules) begin the journey to becoming Go’s [new dependency management system](https://blog.golang.org/versioning-proposal). |
|               1.13 and up               | Starting in Go 1.13, module mode becomes the default for all development. |

#### Experimenting

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

#### Further Reading

[The Go Blog: Using Go Modules](https://blog.golang.org/using-go-modules)

[Go without GOPATH (introduction to modules)](https://dev.to/dizdarevic/golang-without-a-path-47jp)

### `go get: no install location`

#### Error Message

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

#### Resolution

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

#### More Information

***From the `get` docs***

**-t** instructs get to also download the packages required to build the tests for the specified packages.

**-u** instructs get to use the network to update the named packages and their dependencies. By default, get uses the network to check out missing packages but does not use it to look for updates to existing packages.

**-v**  enables verbose progress and debug output.

**GOPATH** When checking out a new package, get creates the target directory `GOPATH/src/<import-path>`. If the GOPATH contains multiple entries, get uses the first one. (For more details see: 'go help gopath'.)

Note: When checking out or updating a package, get looks for a branch or tag that matches the locally installed version of Go. The most important rule is that if the local installation is running version "go1", get searches for a branch or tag named "go1". If no such version exists it retrieves the default branch of the package.

### `path ... is not a package in module rooted at ...`

#### Error Message

```bash
$ go get -u
go get .: path C:\Go_Projects\learning_go\Algorithms_in_Go__Jon_Calhoun is not a package in module rooted at C:\Go_Projects\learning_go\Algorithms_in_Go__Jon_Calhoun
```

#### Resolution

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

### `exec: "gcc": executable file not found in %PATH%`

#### Error Message

```bash
$ cd /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun (master)
$ go test ...
# runtime/cgo
exec: "gcc": executable file not found in %PATH%
ok      archive/tar     (cached)
ok      archive/zip     (cached)
ok      bufio   (cached)
ok      bytes   (cached)
ok      compress/bzip2  (cached)
ok      compress/flate  (cached)
ok      compress/gzip   (cached)
ok      compress/lzw    (cached)
          :
          :
--- FAIL: TestInternalLinkerCgoExec (9.88s)
    nm_test.go:111: building test executable failed: exit status 2 # command-line-arguments
        exec: "gcc": executable file not found in %PATH%
--- FAIL: TestExternalLinkerCgoExec (12.75s)
    nm_test.go:111: building test executable failed: exit status 2 # command-line-arguments
        exec: "gcc": executable file not found in %PATH%
--- FAIL: TestCgoLib (12.79s)
    nm_test.go:246: building test lib failed: exit status 2 # mylib
        exec: "gcc": executable file not found in %PATH%
          :
          :          
```

#### Resolution

1. Download and install the latest Go (1.13.8 as of 2020-02-14)
   https://golang.org/doc/install?download=go1.13.8.windows-amd64.msi

2. Explicitly define GOPATH

   ```bash
   $ vi ~/.bashrc
   # insert the following line
   export GOPATH=/c/Users/kesize/go
   ```

3. (Re)install gotest

   ```bash
   $ go get -u github.com/rakyll/gotest
   ```


#### Notes

* [cmd/go: 'go build' in module mode rebuilds vendored dependencies in GOROOT #27285](https://github.com/golang/go/issues/27285)

* [cmd/go: `go test` complains about missing gcc (even if no tests exist, and only for certain imports) #27303](https://github.com/golang/go/issues/27303)

* To install gcc, the GNU C Compiler, on Windows, use [TDM-GCC](http://tdm-gcc.tdragon.net/download). 







## [Notes from Algorithms in Go, Module 01](Algorithms_in_Go__Jon_Calhoun\Algorithms_in_Go_Module01.md)

## Using **Delve** to Debug Go Programs from the Command Line

Resources: 

* https://www.jamessturtevant.com/posts/Using-the-Go-Delve-Debugger-from-the-command-line/
* https://www.grant.pizza/blog/test-build-modes/

```bash
$ cd /c/Go_Projects/learning_go/testing
# Install Delve
$ go get -u github.com/derekparker/delve/cmd/dlv
go: downloading github.com/derekparker/delve v1.4.0
go: found github.com/derekparker/delve/cmd/dlv in github.com/derekparker/delve v1.4.0
go get: github.com/derekparker/delve@v1.4.0: parsing go.mod:
        module declares its path as: github.com/go-delve/delve
                but was required as: github.com/derekparker/delve
# Verify installation and version 
$ dlv version
Delve Debugger
Version: 1.5.1
Build: $Id: bca418ea7ae2a4dcda985e623625da727d4525b5 $
# Compile a test binary
$ go test -c
$ ll
total 3312
-rw-r--r-- 1 kbsiz 197611     999 Feb  6  2020 hello.go
-rw-r--r-- 1 kbsiz 197611    2339 Feb  6  2020 hello_test.go
-rwxr-xr-x 1 kbsiz 197611 3383296 Dec 30 09:55 testing.test.exe*
# Run that binary through the delve debugger
$ dlv exec testing.test.exe
Type 'help' for list of commands.
(dlv)
# Set a breakpoint for the unit test you’re trying to debug (TestGetHello, in this case)
(dlv) break main.TestGetHello
Breakpoint 1 set at 0x2bb09f for _/C_/Go_Projects/learning_go/testing.TestGetHello() c:/go_p
rojects/learning_go/testing/hello_test.go:28
# Print Help
(dlv) help
# Some of the most common commands
    break (alias: b) ------ Sets a breakpoint.
    continue (alias: c) --- Run until breakpoint or program termination.
    next (alias: n) ------- Step over to next source line.
    restart (alias: r) ---- Restart process.
    step (alias: s) ------- Single step through program.
    stepout (alias: so) --- Step out of the current function.
    args ----------------- Print function arguments.
    locals --------------- Print local variables.
    print (alias: p) ----- Evaluate an expression.
    set ------------------ Changes the value of a variable.
    vars ----------------- Print package variables.

```



##  The Triple Dot Operator

### Variadic function parameters

If the **last parameter** of a function has type `...T`, it can be called with any number of trailing arguments of type `T`. The actual type of `...T` inside the function is `[]T`.

This example function can be called with, for instance, `Sum(1, 2, 3)` or `Sum()`.

```go
func Sum(nums ...int) int {
    res := 0
    for _, n := range nums {
        res += n
    }
    return res
}
```
Usage example:

```go
fmt.Println(Sum())         // prints 0
fmt.Println(Sum(2))        // prints 2
fmt.Println(Sum(1,2,3,4))  // prints 10
```

### Arguments to variadic functions

You can pass a slice `s` directly to a variadic function if you unpack it with the `s...` notation. In this case no new slice is created.

In this example, we pass a slice to the `Sum` function.

```go
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // prints 17
```

### Array literals

In an array literal, the `...` notation specifies a length equal to the number of elements in the literal.

```
stooges := [...]string{"Moe", "Larry", "Curly"}    // len(stooges) == 3
```

### The go command

Last but not least, three dots are used by the [`go`](https://golang.org/cmd/go/) command as a wildcard when describing package lists.

This command tests all packages in the current directory and its subdirectories.

```
$ go test ./...
```

## [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks)

| Operation                                                    | Idiomatic Go                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| Append **x** to end of slice **s**                           | `s = append(s, x)`                                           |
| Append slice **s2** to the end of slice **s**                | `s = append(s, s2)`                                          |
| Create copy of slice **s**                                   | `newSlice = make([]string, s)`<br />`copy(newSlice, s)`      |
| Cut; remove elements **i** through **j** from **s**          | `s = append(s[:i], s[j:]...)`                                |
| Remove the **i**th element from **s**                        | `s = append(s[:i], s[i+1:]...)`                              |
| Insert **x** at position **i** (concise but less efficient)  | `s = append(s[:i], append([]T{x}, s[i:]...)...)`             |
| Insert **x** at position **i**<br />(avoids extra copy/garbage collection) | `s = append(s, 0 /* use the zero value of the element type */)`<br />`copy(s[i+1:], s[i:])`<br />`s[i] = x` |
| Insert slice **s2** into **s** at position **i**             | `s = append(s[:i], append(s2, s[i:]...)...)`                 |
| Reverse slice **s**                                          | `for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {`<br />` 	   s[left], s[right] = s[right], s[left] `<br />`}` |

Using a slice as a stack

| Operation                         | Idiomatic Go                       |
| --------------------------------- | ---------------------------------- |
| Push **x** onto stack **s**       | `s = append(s, x)`                 |
| Pop **x** from top of stack **s** | `x, s = s[len(s)-1], s[:len(s)-1]` |

