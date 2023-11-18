package ptr_test

import (
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/hightech-ninja/ptr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTo(t *testing.T) {
	t.Run("It should return pointer to the value", func(t *testing.T) {
		original := 42
		got := ptr.To(original)
		require.NotNil(t, got)
		require.Equal(t, original, *got)
	})

	t.Run("It should return pointer to the copy of object", func(t *testing.T) {
		original := 42
		got := ptr.To(original)
		assert.NotNil(t, got)
		assert.Equal(t, original, *got)

		original = 148
		require.NotEqual(t, original, *got)
	})
}

func TestToEmptyble(t *testing.T) {
	t.Run("It should return nil for zero values", func(t *testing.T) {
		var zeroTime time.Time
		require.Nil(t, ptr.ToEmptyble(zeroTime))

		var zeroStr string
		require.Nil(t, ptr.ToEmptyble(zeroStr))

		var zeroInt int
		require.Nil(t, ptr.ToEmptyble(zeroInt))
	})
	t.Run("It should return pointer for non-zero values", func(t *testing.T) {
		got1 := ptr.ToEmptyble(time.Now())
		require.NotNil(t, got1)

		got2 := ptr.ToEmptyble("str")
		require.Equal(t, "str", *got2)

		got3 := ptr.ToEmptyble(42)
		require.Equal(t, 42, *got3)
	})
}

func TestDeref(t *testing.T) {
	t.Run("It should dereference non-nil pointer", func(t *testing.T) {
		ptr1 := ptr.To(123)
		got1 := ptr.Deref(ptr1)
		require.Equal(t, *ptr1, got1)
	})

	t.Run("It should return zero-value if dereferencing nil", func(t *testing.T) {
		type user struct {
			Name  string
			Email string
		}
		var user1 *user
		got1 := ptr.Deref(user1)
		require.Equal(t, user{}, got1)
	})
}

func TestDerefOr(t *testing.T) {
	t.Run("It should dereference non-nil pointer", func(t *testing.T) {
		ptr1 := ptr.To(123)
		got1 := ptr.DerefOr(ptr1, 1)
		require.Equal(t, *ptr1, got1)
	})

	t.Run("It should return default value, if dereferencing nil", func(t *testing.T) {
		type user struct {
			Name  string
			Email string
		}
		defaultUser := user{
			Name:  "unknown",
			Email: "unknown",
		}

		got := ptr.DerefOr(nil, defaultUser)
		require.Equal(t, defaultUser, got)
	})
}

func TestReset(t *testing.T) {
	t.Run("It should not change value or panic, if Reset is called on nil", func(t *testing.T) {
		var nilPtr *string
		require.NotPanics(t, func() {
			ptr.Reset(nilPtr)
		})

		require.Nil(t, nilPtr)
	})

	t.Run("It should reset int pointer to zero", func(t *testing.T) {
		intPtr := ptr.To(1999)
		ptr.Reset(intPtr)
		require.Zero(t, *intPtr)
	})

	t.Run("It should reset string pointer to empty string", func(t *testing.T) {
		strPtr := ptr.To("string")
		ptr.Reset(strPtr)
		require.Zero(t, *strPtr)
	})
}

func TestResetTo(t *testing.T) {
	t.Run("It should not change value or panic, if ResetTo is called on nil", func(t *testing.T) {
		require.NotPanics(t, func() {
			ptr.ResetTo[string](nil, "default")
		})
	})

	t.Run("It should reset int pointer", func(t *testing.T) {
		maxConns := ptr.To(100)
		ptr.ResetTo(maxConns, 10)
		require.Equal(t, 10, *maxConns)
	})

	t.Run("It should reset string pointer", func(t *testing.T) {
		defaultName := "unknown"
		name := ptr.To("John")
		ptr.ResetTo(name, defaultName)
		require.Equal(t, defaultName, *name)
	})
}

func TestShallowCopy(t *testing.T) {
	t.Run("It should return nil and not panic, if coping nil", func(t *testing.T) {
		require.NotPanics(t, func() {
			got := ptr.ShallowCopy[int](nil)
			require.Nil(t, got)
		})
	})

	t.Run("It should return a copy of simple objects", func(t *testing.T) {
		t.Run("int", func(t *testing.T) {
			ptr1 := ptr.To(2024)
			ptr2 := ptr.ShallowCopy(ptr1)
			assert.NotEqual(t, uintptr(unsafe.Pointer(ptr1)), uintptr(unsafe.Pointer(ptr2)), "The address of the pointers must be different")
			require.Equal(t, *ptr1, *ptr2)
		})

		t.Run("string", func(t *testing.T) {
			ptr1 := ptr.To("string")
			ptr2 := ptr.ShallowCopy(ptr1)
			assert.NotEqual(t, uintptr(unsafe.Pointer(ptr1)), uintptr(unsafe.Pointer(ptr2)), "The address of the pointers must be different")
			require.Equal(t, *ptr1, *ptr2)
		})
	})
}

func TestCompare(t *testing.T) {
	t.Run("It should return true, if both parameters are nil", func(t *testing.T) {
		got := ptr.Compare[float64](nil, nil)
		require.True(t, got)
	})
	t.Run("It should return false, if exactly one parameter is nil", func(t *testing.T) {
		got := ptr.Compare(ptr.To(3.14), nil)
		require.False(t, got)

		got = ptr.Compare(nil, ptr.To(3.14))
		require.False(t, got)
	})
	t.Run("It should return true, if comparing same values", func(t *testing.T) {
		got := ptr.Compare(ptr.To(3.14), ptr.To(3.14))
		require.True(t, got)
	})
	t.Run("It should return false, if comparing different values", func(t *testing.T) {
		got := ptr.Compare(ptr.To("first"), ptr.To("second"))
		require.False(t, got)
	})
}

func TestMap(t *testing.T) {
	t.Run("It should convert *int to *string", func(t *testing.T) {
		ptr1 := ptr.To(42)
		ptr2 := ptr.Map(ptr1, strconv.Itoa)
		require.Equal(t, *ptr2, "42")
	})
	t.Run("It should convert *string to *float64", func(t *testing.T) {
		ptr1 := ptr.To("42.0")
		ptr2 := ptr.Map(ptr1, func(s string) float64 {
			v, _ := strconv.ParseFloat(s, 64)
			return v
		})
		require.Equal(t, *ptr2, 42.0)
	})
}
