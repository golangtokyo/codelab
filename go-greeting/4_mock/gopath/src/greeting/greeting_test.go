package greeting_test

import (
	"bytes"
	"testing"
	"time"

	"greeting"
)

func TestGreeting_Do(t *testing.T) {
	g := greeting.Greeting{
		Clock: greeting.ClockFunc(func() time.Time {
			// 2018/08/31 06:00:00
			return time.Date(2018, 8, 31, 06, 0, 0, 0, time.Local)
		}),
	}

	var buf bytes.Buffer
	if err := g.Do(&buf); err != nil {
		t.Error("unexpected error:", err)
	}

	if expected, actual := "おはよう", buf.String(); expected != actual {
		t.Errorf("greeting message wont %s but got %s", expected, actual)
	}
}
