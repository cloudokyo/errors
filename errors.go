package errors

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/rotisserie/eris"
)

var (
	ErrNotFound = errors.New("not found")
)

var (
	// An alias of eris func
	Is = eris.Is

	// An alias of eris func
	As = eris.As

	// An alias of eris func
	New = eris.New

	// An alias of eris func
	Wrap = eris.Wrap

	// An alias of eris func
	Wrapf = eris.Wrapf

	// An alias of eris func
	Unwrap = eris.Unwrap

	// An alias of eris func
	Errorf = eris.Errorf

	// An alias of eris func
	Cause = eris.Cause

	// An alias of eris func
	ToJSON = eris.ToJSON

	// An alias of eris func
	ToString = eris.ToString

	// An alias of eris func
	ToCustomJSON = eris.ToCustomJSON

	// An alias of eris func
	ToCustomString = eris.ToCustomString
)

// Check if is not found error
func IsNotFound(err error) bool {
	if err == nil {
		return false
	} else {
		return err.Error() == ErrNotFound.Error()
	}
}

// Return an not found error
func IgnoreNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	} else {
		return err
	}
}

// Print prints error message with stack trace.
func Print(err error) {
	fmt.Println(ToString(err, true))
}

// Print prints error message with stack trace.
func PrintStack(err error) {
	fmt.Println(ToString(err, true))
}

// ToJSON returns a JSON formatted string for a given error.
//
// An error without trace will be formatted as follows:
//
//	{
//	  "root": [
//	    {
//	      "message": "Root error msg"
//	    }
//	  ],
//	  "wrap": {
//	    "message": "Wrap error msg"
//	  }
//	}
//
// An error with trace will be formatted as follows:
//
//	{
//	  "root": [
//	    {
//	      "message": "Root error msg",
//	      "stack": [
//	        "<Method2>:<File2>:<Line2>",
//	        "<Method1>:<File1>:<Line1>"
//	      ]
//	    }
//	  ],
//	  "wrap": {
//	    "message": "Wrap error msg",
//	    "stack": "<Method2>:<File2>:<Line2>"
//	  }
//	}
//
// Example
//
//	errors.String(err)
//	errors.String(err, false)
//	errors.String(err, true)
func String(err error, withTrace ...bool) string {
	// Get stack trace with default=true
	stackTrace := true
	if len(withTrace) > 0 {
		stackTrace = withTrace[0]
	}

	// Format error to JSON format
	output := ToJSON(err, stackTrace)

	// Convert json to string
	if data, err := json.Marshal(output); err == nil {
		return string(data)
	} else {
		return fmt.Sprintf("%#v", output)
	}
}
