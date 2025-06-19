package maybe

// Steps to solve the problem in lesson 3 of the workshop:
// Define a Maybe type that can hold a value of type T or be invalid
// Define a function that returns an invalid Maybe (it can be None, Nothing,  or any other name you prefer).
// Define an OrElse method that returns the value if valid, or a default value if not valid.
// Go to the pkg/pnp/engine/tview package and fix the code that uses the AsciiArt method.
// Remove the AsciiArt method from the Player interface and implement it in each concrete player type instead.

// Maybe represents a generic type that can hold a value or be invalid.
type Maybe[T any] struct {
	value   T
	valid   bool
}

// None returns an invalid Maybe instance.
func None[T any]() Maybe[T] {
	return Maybe[T]{valid: false}
}

func This[T any](value T) Maybe[T] {
	return Maybe[T]{value: value, valid: true}
}

func If[T any](check bool) Maybe[T] {
	return Maybe[T]{valid: check}
}

func (m Maybe[T]) If(check bool) Maybe[T] {
	m.valid = check

	return  m
}

func (m Maybe[T]) Then(value T) Maybe[T] {
	m.value = value

	return m
}

// Else returns the value if valid, or the provided default value if invalid.
func (m Maybe[T]) Else(defaultValue T) T {
	if m.valid {
		return m.value
	}

	return defaultValue
}

func (m Maybe[T]) Or(other Maybe[T]) Maybe[T] {
	if m.valid {
		return m
	}

	return other
}

// Happy solving!
