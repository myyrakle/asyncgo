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
		length := len(futures)

		// 채널을 전부 받아옵니다.
		channels := make([]chan T, length)
		for i, future := range futures {
			channels[i] = future.Channel()
		}
		results := make([]T, length)
		doneList := make([]bool, length)
		doneCount := 0

		// 모든 채널이 끝날 때까지 기다립니다.
		for doneCount < length {
			for i, ch := range channels {
				if !doneList[i] {
					select {
					case result := <-ch:
						results[i] = result
						doneList[i] = true
						doneCount++
					default:
						continue
					}
				}
			}
		}

		return results
	})
}
