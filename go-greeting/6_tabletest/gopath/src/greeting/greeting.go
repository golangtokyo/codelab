package greeting

import (
	"fmt"
	"io"
	"time"
)

type Clock interface {
	Now() time.Time
}

type ClockFunc func() time.Time

func (f ClockFunc) Now() time.Time {
	return f()
}

type Greeting struct {
	Clock Clock
}

func (g *Greeting) now() time.Time {
	if g.Clock == nil {
		return time.Now()
	}
	return g.Clock.Now()
}

func (g *Greeting) Do(w io.Writer) error {
	h := g.now().Hour()
	var msg string
	switch {
	case h >= 4 && h <= 9:
		msg = "おはよう"
	case h >= 10 && h <= 16:
		msg = "こんにちは"
	default:
		msg = "こんばんは"
	}

	_, err := fmt.Fprint(w, msg)
	if err != nil {
		return err
	}

	return nil
}
