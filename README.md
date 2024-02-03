# asyncgo

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.1.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Boilerplate implementing async pattern

## Basic

You can create asynchronous tasks in the following form:

```
// Executes asynchronous tasks (goroutines) immediately.
// The returned task object is a channel wrapper that holds the execution context of the goroutine.
task := async.RunTask(func(args ...any) int {
    ...
    return 10;
})

taskResult := task.Await() // Waits until the task completes and returns a value.
fmt.Println(taskResult) // 10
```

Parameters can also be passed to Tasks. This allows you to clearly distinguish between task input and output.

```
// Executes asynchronous tasks (goroutines) immediately.
// The returned task object is a channel wrapper that holds the execution context of the goroutine.
task := async.RunTask(func(args ...any) int {
    lhs := args[0].(int)
    rhs := args[1].(int)
    return lhs + rhs;
}, 10, 20)

taskResult := task.Await()
fmt.Println(taskResult) // 30
```

If you absolutely need to recover when a panic occurs within a task, you can use `RunPanicableTask`.
In this case, you must use the Result type to handle panic propagation.

```
task := async.RunPanicableTask(func(args ...any) async.Result[int] {
    ... // If a panic occurs, async.Err is returned.
    return async.Ok(10);
})

taskResult := task.Await()
fmt.Println(taskResult)
```

If you want to wait for multiple tasks simultaneously, you can use the JoinAll function.

```
task1 := ...
task2 := ...

// Wait for both task1 and task2 to finish.
results := async.JoinAll(task1, task2).Await()

for _, result := results {
    ...
}
```
