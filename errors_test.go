package errors_test

import (
	"fmt"
	"testing"

	"github.com/cloudokyo/errors"
)

// go clean -testcache && go test -v -run ^TestErrorWrap
func TestErrorWrap(t *testing.T) {
	fmt.Println("--> Test wrapping error with single")
	{
		err1 := errors.New("level 1")
		err2 := errors.Wrap(err1, "level 2")
		err3 := errors.Wrap(err2, "level 3")

		fmt.Println("wrap(err1):", err2)
		fmt.Println("wrap(err2):", err3)

		fmt.Println("unwrap(err2):", errors.Unwrap(err2))
		fmt.Println("unwrap(err3):", errors.Unwrap(err3))
	}

	fmt.Println("--> Test wrapping error with format")
	{
		err1 := errors.New("level 1")
		err2 := errors.Wrapf(err1, "%s %s", "level", "2")
		err3 := errors.Wrapf(err2, "%s %s", "level", "3")

		fmt.Println("wrapf(err1):", err2)
		fmt.Println("wrapf(err2):", err3)

		fmt.Println("unwrap(err2):", errors.Unwrap(err2))
		fmt.Println("unwrap(err3):", errors.Unwrap(err3))
	}

	fmt.Println("--> Test wrapping error with format and custom separator")
	{
		err1 := errors.New("level 1")
		err2 := errors.Wrapf(err1, "%s %s", "level", "2")
		err3 := errors.Wrapf(err2, "%s %s", "level", "3")

		fmt.Println("wrapf(err1):", err2)
		fmt.Println("wrapf(err2):", err3)

		fmt.Println("unwrap(err2):", errors.Unwrap(err2))
		fmt.Println("unwrap(err3):", errors.Unwrap(err3))
	}
}

func TestErrorStackTrace(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.Wrap(err1, "error 2")
	err3 := errors.Wrap(err1, "error 3")

	t.Run("Test error wrap", func(t *testing.T) {
		fmt.Println("wrap(err1):", err1)
		fmt.Println("wrap(err2):", err2)
		fmt.Println("wrap(err3):", err3)
	})

	t.Run("Test error stack trace", func(t *testing.T) {
		fmt.Println(errors.String(err1, true))
		fmt.Println(errors.String(err2, true))
		fmt.Println(errors.String(err3, true))
	})

	t.Run("Test print error", func(t *testing.T) {
		errors.Print(err1)
		errors.Print(err2)
		errors.Print(err3)
	})

	t.Run("Test print stack error", func(t *testing.T) {
		errors.PrintStack(err1)
		errors.PrintStack(err2)
		errors.PrintStack(err3)
	})
}
