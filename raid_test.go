package raid

import (
	"testing"
	"time"
)

func TestRaid(t *testing.T) {
	t.Run("en_de_code", func(t *testing.T) {
		src := []byte{1, 2, 255, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		encRaid := make([]byte, 32)
		decRaid := make([]byte, 20)
		encodeRaid(encRaid, src)
		decodeRaid(decRaid, encRaid)
		encPrefix := make([]byte, 4)
		decPrefix := make([]byte, 2)
		encodePrefix(encPrefix, src)
		decodePrefix(decPrefix, encPrefix)

		for i := range decRaid {
			if decRaid[i] != src[i] {
				t.Error("missmatch decoded raid")
			}
		}
		for i := range decPrefix {
			if decPrefix[i] != src[i] {
				t.Error("missmatch decoded prefix")
			}
		}
		t.Logf("%#+v\n", src)
		t.Logf("%#+v\n", decRaid)
		t.Logf("%#+v\n", encRaid)
		t.Logf("%#+v\n", decPrefix)
		t.Logf("%#+v\n", encPrefix)
	})
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
		if rr.Prefix() != r.Prefix() {
			t.Errorf("prefix changed from %s, parsed as %s", r.Prefix(), rr.Prefix())
		}
		t.Log(r.Time().String())

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
