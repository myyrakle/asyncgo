package async

// Future is a type that represents a value that will be available in the future.
type Future[T any] interface {
	// Await blocks until the value is available and returns it.
	Await() T
	// Channel returns a channel that will contain the value when it is available.
	Channel() chan T
}

// RunTask runs a function asynchronously and returns a Future that will contain the result.
func JoinAll[T any](futures ...Future[T]) Future[[]T] {
	return RunTask(func(args ...any) []T {
		results := AwaitAll(futures...)

		return results
	})
}

// Await all futures and return it
func AwaitAll[T any](futures ...Future[T]) []T {
	results := make([]T, len(futures))

	// 모든 채널이 끝날 때까지 기다립니다.
	for i, future := range futures {
		results[i] = future.Await()
	}

	return results
}


type Futures []Future[any]

func (f Futures) Await() []any {
	return AwaitAll(f...)
}