package log

import (
	"bytes"
	"testing"

	"github.com/nbio/st"
	log "gopkg.in/Sirupsen/logrus.v0"
)

// Since current implementation acts just as proxy to
// Logrus API, tests are not testing too much here.
func TestLogger(t *testing.T) {
	var buffer bytes.Buffer
	SetOutput(&buffer)
	SetLevel(log.DebugLevel)
	SetFormatter(&log.TextFormatter{DisableColors: true, DisableTimestamp: true, DisableSorting: true})

	Info("foo bar")
	st.Expect(t, safeString(buffer), `level=info msg="foo bar"`)

	buffer.Reset()
	Infof("foo %s", "bar")
	st.Expect(t, safeString(buffer), `level=info msg="foo bar"`)

	buffer.Reset()
	Error("foo bar")
	st.Expect(t, safeString(buffer), `level=error msg="foo bar"`)

	buffer.Reset()
	Errorf("foo %s", "bar")
	st.Expect(t, safeString(buffer), `level=error msg="foo bar"`)

	buffer.Reset()
	Warn("foo bar")
	st.Expect(t, safeString(buffer), `level=warning msg="foo bar"`)

	buffer.Reset()
	Warnf("foo %s", "bar")
	st.Expect(t, safeString(buffer), `level=warning msg="foo bar"`)

	buffer.Reset()
	Debug("foo bar")
	st.Expect(t, safeString(buffer), `level=debug msg="foo bar"`)

	buffer.Reset()
	Debugf("foo %s", "bar")
	st.Expect(t, safeString(buffer), `level=debug msg="foo bar"`)
}

func safeString(buf bytes.Buffer) string {
	b := buf.Bytes()
	return string(b[:len(b)-2])
}
