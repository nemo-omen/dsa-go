# Watch mode with Go Watch

Instead of invoking `run` the normal way: `go run blahblah.go`, run it with `gow run blahblah.go`.

Testing: `gow test`

## Steps
1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
   1. Commit
5. Refactor
   1. Amend commit with refactored code