# Unexpected Behavior of Defer Statement with Variable Capture in Go

This example demonstrates an uncommon and potentially confusing behavior of Go's `defer` statement related to variable capture.  The `defer` statement, while generally straightforward, can exhibit unintuitive behavior when dealing with variables that change their value between the point of the `defer` call and the function's actual termination.

The `bug.go` file contains the problematic code.  The `bugSolution.go` file illustrates a way to resolve the issue.

**Key Point:**  The `defer` statement captures the value of the variable *at the time of the `defer` call*, not at the time of execution. Therefore, any subsequent modifications to the variable will not be reflected in the deferred function.