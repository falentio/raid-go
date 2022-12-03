package raid

import (
	"testing"
	"time"
)

func TestRaid(t *testing.T) {
	t.Run("parse", func(t *testing.T) {
		r := NewRaid().WithMessage(0xff77)
		time.Sleep(time.Second)
		str := r.String()
		rr, err := RaidFromString(str)
		if err != nil {
			t.Errorf("unknown error: %v", err)
		}
		if rr.String() != r.String() {
			t.Errorf("raid changed from %q, parsed as %q", r.String(), rr.String())
		}
		if rr.Time() != r.Time() {
			t.Errorf("time changed from %v, parsed as %v", r.Time(), rr.Time())
		}
		if rr.Message() != r.Message() {
			t.Errorf("message changed from %d, parsed as %d", r.Message(), rr.Message())
		}

		r2 := NewRaid()
		if !r.Less(r2) {
			t.Errorf("less method false return true")
		}
	})
}

func BenchmarkRaid(b *testing.B) {
	l := ""
	for i := 0; i < b.N; i++ {
		l = NewRaid().String()
	}
	_ = l
	b.ReportAllocs()
}
