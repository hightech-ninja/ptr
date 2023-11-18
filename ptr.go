package ptr

// To takes pointer to the value. Useful for creating pointers for literals or non-addressable values.
// Avoid using on large structs or arrays where passing the address of the original variable is more efficient.
func To[T any](value T) *T {
	return &value
}

func ToEmptyble[T comparable](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}

// Deref safely dereference pointer. If ptr is nil, it returns zero value of type T.
func Deref[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// DerefOr safely dereference pointer. If ptr is nil, it returns provided default value of tpye T.
func DerefOr[T any](ptr *T, def T) T {
	if ptr == nil {
		return def
	}
	return *ptr
}

// Reset sets the value of pointer to zero value of type T. It does nothing if the pointer is nil.
// Best use for reinitializing a pointer variable without reallocating memory.
func Reset[T any](ptr *T) {
	if ptr != nil {
		var zero T
		*ptr = zero
	}
}

// ResetTo sets the value of pointer to the given value of type T. It does nothing, if the pointer is nil.
func ResetTo[T any](ptr *T, to T) {
	if ptr != nil {
		*ptr = to
	}
}

// ShallowCopy creates a shallow copy of the value pointed to by *T and returns a pointer to the new value. Returns nil if the input pointer is nil.
// Useful when you need a separate instance of a value to modify without affecting the original. Not suitable for types containing pointers themselves.
func ShallowCopy[T any](ptr *T) *T {
	if ptr == nil {
		return nil
	}

	value := *ptr
	return &value
}

// Compare checks if two pointers of type *T point to equal values. Returns true if both are nil or if the values are equal, false otherwise.
func Compare[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return *a == *b
}

// Map applies a function to the value pointed to by *T and returns a pointer to the new value of type *U. Returns nil if the input pointer is nil.
// Useful for transforming pointer type.
func Map[T any, U any](ptr *T, f func(T) U) *U {
	if ptr == nil {
		return nil
	}

	result := f(*ptr)
	return &result
}
