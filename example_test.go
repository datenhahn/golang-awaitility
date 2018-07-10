package golang_awaitility

import "time"

func ExampleAwait() {
	Await(100 * time.Millisecond, 1000 * time.Millisecond, func() bool {
		// do a real check here, e.g. some kind of isConnected()
		return true
	})
}
