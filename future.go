package async

// Future is a type that represents a value that will be available in the future.
type Future[T any] interface {
	// Await blocks until the value is available and returns it.
	Await() T
	// Channel returns a channel that will contain the value when it is available.
	Channel() chan T
}
