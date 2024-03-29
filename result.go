package async

// Error handling with the Result type.
// Result<T> is the type used for returning and propagating errors.
// It is an enum with the variants, Ok(T), representing success and containing a value, and Err, representing error and containing an error value.
type Result[T any] struct {
	ok  *T
	err error
}

// Creates a new Result<T> containing an Ok value.
func Ok[T any](value T) Result[T] {
	return Result[T]{ok: &value}
}

// Creates a new Result<T> containing an Err value.
func Err[T any](err error) Result[T] {
	return Result[T]{err: err}
}

// Returns true if the result is Ok.
func (self Result[T]) IsOk() bool {
	return self.ok != nil
}

// Returns true if the result is Err.
func (self Result[T]) IsErr() bool {
	return self.err != nil
}

// Maps a Result<T, E> to Result<U, E> by applying a function to a contained Ok value, leaving an Err value untouched.
// This function can be used to compose the results of two functions.
func (self Result[T]) Map(f func(T) T) Result[T] {
	if self.IsOk() {
		return Ok[T](f(*self.ok))
	} else {
		return self
	}
}

// Returns the provided default (if Err), or applies a function to the contained value (if Ok),
// Arguments passed to map_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use map_or_else, which is lazily evaluated.
func (self Result[T]) MapOr(defaultValue T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue
	}
}

// Maps a Result<T, E> to U by applying fallback function default to a contained Err value, or function f to a contained Ok value.
// This function can be used to unpack a successful result while handling an error.
func (self Result[T]) MapOrElse(defaultValue func() T, f func(T) T) T {
	if self.IsOk() {
		return f(*self.ok)
	} else {
		return defaultValue()
	}
}

// Maps a Result<T, E> to Result<T, F> by applying a function to a contained Err value, leaving an Ok value untouched.
// This function can be used to pass through a successful result while handling an error.
func (self Result[T]) MapErr(f func(error) error) Result[T] {
	if self.IsErr() {
		return Err[T](f(self.err))
	} else {
		return self
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
func (self Result[T]) Expect(message string) T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(message)
	}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
func (self Result[T]) Unwrap() T {
	if self.IsOk() {
		return *self.ok
	} else {
		panic(self.err)
	}
}

// Returns the contained Err value, consuming the self value.
func (self Result[T]) ExpectErr(message string) error {
	if self.IsErr() {
		return self.err
	} else {
		panic(message)
	}
}

// Returns the contained Err value, consuming the self value.
func (self Result[T]) UnwrapErr() error {
	if self.IsErr() {
		return self.err
	} else {
		panic(self.ok)
	}
}

// Returns res if the result is Ok, otherwise returns the Err value of self.
// Arguments passed to and are eagerly evaluated; if you are passing the result of a function call, it is recommended to use and_then, which is lazily evaluated.
func (self Result[T]) And(res Result[T]) Result[T] {
	if self.IsOk() {
		return res
	} else {
		return self
	}
}

// Calls op if the result is Ok, otherwise returns the Err value of self.
// This function can be used for control flow based on Result values.
func (self Result[T]) AndThen(op func(T) Result[T]) Result[T] {
	if self.IsOk() {
		return op(*self.ok)
	} else {
		return self
	}
}

// Returns res if the result is Err, otherwise returns the Ok value of self.
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use or_else, which is lazily evaluated.
func (self Result[T]) Or(res Result[T]) Result[T] {
	if self.IsErr() {
		return res
	} else {
		return self
	}
}

// Calls op if the result is Err, otherwise returns the Ok value of self.
// This function can be used for control flow based on result values.
func (self Result[T]) OrElse(op func(error) Result[T]) Result[T] {
	if self.IsErr() {
		return op(self.err)
	} else {
		return self
	}
}

// Returns the contained Ok value or a provided default.
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use unwrap_or_else, which is lazily evaluated.
func (self Result[T]) UnwrapOr(defaultValue T) T {
	if self.IsOk() {
		return *self.ok
	} else {
		return defaultValue
	}
}

// Returns the contained Ok value or computes it from a closure.
func (self Result[T]) UnwrapOrElse(f func() T) T {
	if self.IsOk() {
		return *self.ok
	} else {
		return f()
	}
}
