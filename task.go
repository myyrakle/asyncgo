package async

type runTaskFuture[T any] struct {
	ch chan T
}

func (self *runTaskFuture[T]) Await() T {
	return <-self.ch
}

func (self *runTaskFuture[T]) Channel() chan T {
	return self.ch
}

// RunTask runs a function asynchronously and returns a Future that will contain the result.
func RunTask[T any](f func(args ...any) T, args ...any) Future[T] {
	ch := make(chan T)

	future := runTaskFuture[T]{ch: ch}

	go func() {
		defer close(ch)
		ch <- f(args...)
	}()

	return &future
}
