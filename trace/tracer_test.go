package trace

import "testing"
import "bytes"

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New should not be nil")
	} else {
		tracer.Trace("Hello trace package.")

		if buf.String() != "Hello trace package.\n" {
			t.Error("Trace should not write '%s'.", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer = Off()
	silentTracer.Trace("something")
}
