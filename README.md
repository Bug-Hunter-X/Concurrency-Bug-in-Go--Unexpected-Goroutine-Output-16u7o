# Concurrency Bug in Go: Unexpected Goroutine Output

This repository demonstrates a common concurrency bug in Go where the output of goroutines is unexpected due to improper handling of loop variables.  The `bug.go` file showcases the buggy code, while `bugSolution.go` provides the corrected version.

The bug arises because the loop variable `i` is shared across all goroutines.  By the time a goroutine gets scheduled to run, the loop might have already completed, leading to unexpected values of `i` being printed.

The solution uses a closure to capture the current value of `i` for each goroutine, ensuring that each goroutine operates on its own copy of the variable.

This example highlights the importance of carefully managing variables in concurrent Go programs to avoid race conditions and unexpected behavior.