# Ken's Notes for *Algorithms in Go by Jon Calhoun*

[TOC]



## [Find a number in a list of numbers](https://courses.calhoun.io/lessons/les_algo_m01_05)

```bash
$ cd /c/Go_Projects/learning_go/Algorithms_in_Go__Jon_Calhoun/module01

$ ll num_in_list*
-rw-r--r-- 1 Ken 197609  236 Feb 11 22:21 num_in_list.go
-rw-r--r-- 1 Ken 197609 1879 Feb 11 21:02 num_in_list_test.go

$ grep TestNumInList *.go
num_in_list_test.go:8:func TestNumInList(t *testing.T) {

$ gotest -run=NumInList
PASS
ok      algo/module01   0.686s

$ gotest -run=NumInList -v
=== RUN   TestNumInList
=== RUN   TestNumInList/([1_2_3_4_5],1)
=== RUN   TestNumInList/([1_2_3_4_5],2)
      :
      :
    --- PASS: TestNumInList/([8_2_5_4_1_8_9_3_0_88_23_44_123],6) (0.00s)
    --- PASS: TestNumInList/([-1_-1_-1_-1_-1_-1_-1_-1],-1) (0.00s)
    --- PASS: TestNumInList/([-1_-1_-1_-1_-1_-1_-1_-1],1) (0.00s)
PASS
ok      algo/module01   0.614s      
```

***Understanding the `for` loop in Go***

```go
// Prints out integers from 0 to len(mySlice)-1
for i := range mySlice {
    fmt.Println(i)
}
```

...is equivalent to...

```go
// Prints out integers from 0 to len(mySlice)-1
for i := 0; i < len(mySlice); i++ {
    fmt.Println(i)
}
```

...and not to be confused with...

```go
// Prints out the values in mySlice
for _, value := range mySlice {
    fmt.Println(value)
}
```



## [Sum a list of numbers](https://courses.calhoun.io/lessons/les_algo_m01_06)

Naïve attempt to test:

```bash
 $ gotest -run=Sum
--- FAIL: TestFindTwoThatSum (0.00s)
    --- FAIL: TestFindTwoThatSum/[1_2_3_4]_with_sum_7 (0.00s)
```

***What happened?***

The argument passed to  `-run` is a regular expression.  In this case, any tests with a name containing "Sum" are run and that includes the tests for "Sum" and the tests for "TestFindTwoThatSum".

Correct way to execute tests for "Sum" only:

```bash
$ gotest -run=TestSum
PASS
ok      algo/module01   0.541s
```

## [Reverse a string](https://courses.calhoun.io/lessons/les_algo_m01_07)

This exercise gave us a change to play with

* Runes
* **For loop syntax** and **commas**
* Typed and untyped **nil**
* Adding names to **table-driven tests**

Some things we learned...

* When manipulating strings in Go, working with **runes** is a good idea.

* `range` over a `string` will iterate rune by rune (not byte-by-byte)



***VSCode note: Format on Save***

* Using the [Go plugin]( https://github.com/Microsoft/vscode-go), set ` "go.formatOnSave": true` 

Another thing you should look at is changing "go.formatTool" from "go fmt" to "goimports" (and installing goimports). Having it manage your imports is very handy (though use with caution, double-check it's importing the right thing)

To install goimports: ` go get golang.org/x/tools/cmd/goimports`

Note: If there are errors, neither format nor import cleanup happens

## [The classic FizzBuzz problem](https://courses.calhoun.io/lessons/les_algo_m01_08)

**TIP**: `fizz_buzz_test.go` illustrates how to build a test that checks what is written to STDOUT.

```go
osStdout := os.Stdout // keep backup of the real stdout
os.Stdout = writer
defer func() {
	// Undo what we changed when this test is done.
	os.Stdout = osStdout
}()

FizzBuzz(tc.n)
writer.Close()

var buf bytes.Buffer
io.Copy(&buf, testStdout)
got := buf.String()
if got != tc.want {
   :
   :
```

**TIP**: How to execute a single case within a set of table-driven tests

```bash
$ gotest -run=FizzBuzz/N=5 -v
=== RUN   TestFizzBuzz
=== RUN   TestFizzBuzz/N=5
--- PASS: TestFizzBuzz (0.00s)
    --- PASS: TestFizzBuzz/N=5 (0.00s)
PASS
ok      algo/module01   0.396s
```



## Convert a decimal to any base (2-16)

## Convert a number in any base (2-16) to Decimal

## Convert a number from any base to any other base (2-16)

## Find two numbers in a list that sum to a given amount

## Factor a number

## Fibonacci numbers

## Greatest common divisor (GCD)

## stdin and stdout



------------------

## Module 2: Sorting Algorithms

## Bubble Sort

## Insertion Sort

## x





------------------------------

## Appendix: Running specific tests

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

More `-run` examples...

```go test -run &#39;&#39;      # Run all tests.
go test -run ''      # Run all tests.
go test -run Foo     # Run top-level tests matching "Foo", such as "TestFooBar".
go test -run Foo/A=  # For top-level tests matching "Foo", run subtests matching "A=".
go test -run /A=1    # For all top-level tests, run subtests matching "A=1".
```

For more details, see: https://golang.org/pkg/testing/

and https://ieftimov.com/post/testing-in-go-go-test/

## Appendix: Common Golang Environment Problems and Solutions

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






