// Package rtutil implements runtime utilities.
package rtutil

import (
	"fmt"
	"runtime"
)

func RecoverError(r interface{}) error {
	buf := make([]byte, 1<<18)
	n := runtime.Stack(buf, false)
	return fmt.Errorf("Panic recovered. Reason: %v. Stack: %s", r, buf[0:n])
}
