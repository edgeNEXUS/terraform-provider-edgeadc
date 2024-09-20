package provider

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func Retry(
	fn interface{}, args []interface{}, maxRetry int,
	startBackoff, maxBackoff time.Duration) ([]reflect.Value, error) {

	fnVal := reflect.ValueOf(fn)
	if fnVal.Kind() != reflect.Func {
		return nil, errors.New("retry: function type required")
	}

	argVals := make([]reflect.Value, len(args))
	for i, arg := range args {
		argVals[i] = reflect.ValueOf(arg)
	}

	for attempt := 0; attempt < maxRetry; attempt++ {
		result := fnVal.Call(argVals)
		errVal := result[len(result)-1]

		if errVal.IsNil() {
			return result, nil
		}
		if attempt == maxRetry-1 {
			return result, errVal.Interface().(error)
		}
		time.Sleep(startBackoff)
		if startBackoff < maxBackoff {
			startBackoff *= 2
		}
		fmt.Printf(
			"Retrying function call, attempt: %d, error: %v\n",
			attempt+1, errVal,
		)
	}
	return nil, fmt.Errorf("retry: max retries reached without success")
}
