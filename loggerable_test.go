package tcb_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/deadcheat/tcb"
)

func TestSilentLogger(t *testing.T) {
	t.Log("=== Case 1. Just call functions do nothing")
	// Do nothing, just only call
	s := tcb.SilentLogger{}
	s.Log()
	s.Logf("%d", 0)
}

func TestDefaultLogger(t *testing.T) {
	t.Log("=== Case 1. Test Log to call Defaultlogger")
	l := tcb.NewDefaultActiveLogger()
	t.Log("==== Case 1.1. Test LogEnabled to call Defaultlogger")
	if !l.LogEnabled() {
		t.Error("NewDefaultActiveLogger should return log-enabled logger")
		t.Fail()
	}

	var messageBuffer bytes.Buffer
	l.Logger.SetOutput(&messageBuffer)
	message := "message"
	l.Log(message)
	if !strings.Contains(messageBuffer.String(), message) {
		t.Error("Logged Message may be different", messageBuffer.String(), message)
		t.Fail()
	}
	t.Log("=== Case 3. Test Logf to call Defaultlogger")
	messageBuffer.Reset()
	l.Logger.SetOutput(&messageBuffer)
	l.Logf("message is %s", message)
	if !strings.Contains(messageBuffer.String(), "message is "+message) {
		t.Error("Logged Message may be different", messageBuffer.String(), message)
		t.Fail()
	}
	t.Log("=== Case 4. Test LogEnabled to call DefaultDisabledLogger")
	l = tcb.NewDefaultDisabledLogger()
	if l.LogEnabled() {
		t.Error("NewDefaultActiveLogger should return log-disabled logger")
		t.Fail()
	}
	t.Log("=== Case 5. Test LogEnabled to call Nil logger")
	l = nil
	if l.LogEnabled() {
		t.Error("nil logger should return nothing, but LogEnabled return false")
		t.Fail()
	}
}
