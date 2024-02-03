package async

/// Future is a type that represents a value that will be available in the future.
type Future[T any] interface {
	/// Await blocks until the value is available and returns it.
	Await() T

	Channel() chan T
}
