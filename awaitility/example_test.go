package awaitility

import "time"

func ExampleAwait() {
	err := Await(100*time.Millisecond, 1000*time.Millisecond, func() bool {
		// do a real check here, e.g. some kind of isConnected()
		return false
	})

	if err != nil {
		// the await condition did not become true in time
	}
}
