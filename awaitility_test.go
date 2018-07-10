package golang_awaitility

import (
	"testing"
	"time"
	"errors"
)

func TestAwaitTrue(t *testing.T) {
	err := AwaitDefault(func() bool {
		return true
	})

	if err != nil {
		t.Errorf("Await ended with unexpected error: %s", err.Error())
	}
}

func TestAwaitFalseThenOk(t *testing.T) {

	testVar := false

	go func() {
		time.Sleep(150 * time.Millisecond)
		testVar = true
	}()

	err := Await(100*time.Millisecond, 300*time.Millisecond, func() bool {
		return testVar
	})

	if err != nil {
		t.Errorf("Await ended with unexpected error: %s", err.Error())
	}
}

func TestAwaitFalse(t *testing.T) {

	startTime := time.Now()

	err := Await(10*time.Millisecond, 20*time.Millisecond, func() bool {
		return false
	})

	delta := time.Now().Sub(startTime) / time.Millisecond

	if delta > 20 {
		t.Errorf("Took too long to cancel, took %dms, expected to be under 20ms", delta)
	}

	if err == nil {
		t.Errorf("Expected await to end with error, but ended ok")
	} else {
		if ! IsAwaitTimeoutError(err) {
			t.Errorf("Expected a Timeout Error, actual error is: %s", err.Error())
		}
	}
}

func TestIsAwaitTimeoutError(t *testing.T) {
	okErr := errors.New(TIMEOUT_ERROR)
	otherErr := errors.New("Other error")

	if !IsAwaitTimeoutError(okErr) {
		t.Errorf("Expected '%s' to be a timeout error, but wasn't", okErr.Error())
	}

	if IsAwaitTimeoutError(otherErr) {
		t.Errorf("Expected '%s' not to be a timeout error, but was", otherErr.Error())
	}
}

func TestPassTimeInFunc(t *testing.T) {

	startTime := time.Now()

	err := Await(10*time.Millisecond, 100*time.Millisecond, func() bool {
		time.Sleep(200 * time.Millisecond)
		return true
	})

	delta := time.Now().Sub(startTime) / time.Millisecond

	if delta > 100 {
		t.Errorf("Took too long to cancel, took %dms, expected to be under 100ms", delta)
	}

	if err == nil {
		t.Errorf("Expected await to cancel")
	}
}

func TestAwaitPanicFalse(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Logf("Catched PANIC successfully: %s", r)
		}
	}()

	AwaitPanic(10*time.Millisecond, 100*time.Millisecond, func() bool {
		return false
	})

	t.Error("AwaitPanic should panic and not continue until this statement")
}

func TestAwaitPanicTrue(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Catched unexpected PANIC: %s", r)
		}
	}()

	AwaitPanic(10*time.Millisecond, 100*time.Millisecond, func() bool {
		return true
	})
}

func TestAwaitPanicDefaultTrue(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Catched unexpected PANIC: %s", r)
		}
	}()

	AwaitPanicDefault(func() bool {
		return true
	})
}

func TestAwaitLimits(t *testing.T) {

	err := Await(0, 10*time.Millisecond, func() bool {
		return false
	})

	if err == nil {
		t.Errorf("Expected error when 0 poll interval but none received")
	} else {
		expected := "PollInterval cannot be 0 or below, got: 0"
	   if err.Error() != expected {
	   	t.Errorf("Expected message '%s' but got '%s'", expected, err.Error())
	   }
	}

	err = Await(10*time.Millisecond, 0, func() bool {
		return false
	})

	if err == nil {
		t.Errorf("Expected error when 0 poll interval but none received")
	} else {
		expected := "AtMost timeout cannot be 0 or below, got: 0"
		if err.Error() != expected {
			t.Errorf("Expected message '%s' but got '%s'", expected, err.Error())
		}
	}


	err = Await(20*time.Millisecond, 10*time.Millisecond, func() bool {
		return false
	})

	if err == nil {
		t.Errorf("Expected error when 0 poll interval but none received")
	} else {
		expected := "PollInterval must be smaller than atMost timeout, got: pollInterval=20000000, atMost=10000000"
		if err.Error() != expected {
			t.Errorf("Expected message '%s' but got '%s'", expected, err.Error())
		}
	}


}
