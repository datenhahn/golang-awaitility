# golang-awaitility

![Golang Awaitility Logo](golang-awaitility-logo.png)

 Package golang_awaitility provides a simple mechanism to poll for conditions with a general timeout.

It is inspired by the great jvm lib "awaitility" (see https://github.com/awaitility/awaitility)

## Example

```go
package example

import (
	"github.com/ecodia/golang-awaitility/awaitility"
	"testing"
	"time"
	)

func TestSomething(t *testing.T)  {
   	err := awaitility.Await(100 * time.Millisecond, 1000 * time.Millisecond, func() bool {
   		// do a real check here, e.g. some kind of isConnected()
   		return true
   	})
   	
   	if err != nil {
   		t.Errorf("Unexpected error during await: %s", err)
   	}
 
}
```