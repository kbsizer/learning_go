// Sample test to show how to write a basic unit test in Golang.
//
// USING GO TEST
//       $ go test
// This will suppress informational messages when tests pass, but will automatically
// switch to verbose mode (see below) if there is a failure.
//       $ go test -v
// Force go test into verbose mode (writing out messages) whether tests are successful or
// not.
//       $ go test -run Down
// Run a subset of tests (-run takes a regular expression as its argument)
//
// HANDY TESTING METHODS
//   t.Log() / t.Logf()     -- add information; does not affect result of test
//   t.Error() / t.Errorf() -- log message and set status of test to FAILED
//   t.Fatal() / t.Fatalf() -- log message and end test *immediately* with status of FAILED
//
// PACKAGE NAMES IN TESTS
//
// Use the same package as the production code for whitebox testing (accessing unexported
// functions, methods and definitions); use production package plus "_test" to perform
// backbox testing (only exported definitions are visible to the test)
//
// Bill Kennedy-isms:
//   * Extra brackets are used to add readibility and demarcate scope
//   * These tests use the "given-when-then" framework (Bill Kennedy calls it "given-when-should")
//     see: https://martinfowler.com/bliki/GivenWhenThen.html
package example1
