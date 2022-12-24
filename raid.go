package raid

import (
	"bytes"
	"crypto/rand"
	"database/sql/driver"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"
)

var (
	NilRaid     = Raid{}
	defaultRaid = Raid{}
	timeOffset  = time.Date(2022, time.January, 0, 0, 0, 0, 0, time.Local)
	bufferPool  = sync.Pool{
		New: func() any {
			b := make([]byte, 12)
			return &b
		},
	}
)

var (
	ErrInvalidId = errors.New("raid: invalid raid id")
)

func init() {
	defaultRaid = NilRaid.
		WithRandom().
		WithPrefix("axf").
		WithMessage(0x7ff0)
}

// Raid is 20 byte identifier.
// 2 byte prefix -- enough for 3 characters in base32.
// 4 byte unix seconds timestamp -- make id became sortable.
// 2 byte message -- 16 bit message, masked with 15th and 16th byte.
// 12 byte randoms -- 96 bit randoms, make id un guessable
type Raid [20]byte

func NewRaid() Raid {
	return defaultRaid.WithRandom().WithTimestampNow()
}

func RaidFromString(str string) (Raid, error) {
	rr := &Raid{}
	if err := rr.UnmarshallText(stringToBytes(str)); err != nil {
		return NilRaid, err
	}
	return *rr, nil
}

func (rr Raid) WithPrefix(prefix string) Raid {
	return rr.WithPrefixByte(stringToBytes(prefix))
}

func (rr Raid) WithPrefixByte(prefix []byte) Raid {
	rr.setPrefixByte(prefix)
	return rr
}

func (rr Raid) WithMessage(msg uint16) Raid {
	rr.setMessage(msg)
	return rr
}

func (rr Raid) WithRandom() Raid {
	rr.updateRandoms()
	return rr
}

func (rr Raid) WithTimestamp(t time.Time) Raid {
	rr.setTimestamp(t)
	return rr
}

func (rr Raid) WithTimestampNow() Raid {
	return rr.WithTimestamp(time.Now())
}

func (rr *Raid) updateRandoms() {
	bptr := bufferPool.Get().(*[]byte)
	defer bufferPool.Put(bptr)
	b := *bptr
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		panic(err.Error())
	}
	rr[19] = b[11]
	rr[18] = b[10]
	rr[17] = b[9]
	rr[16] = b[8]
	rr[15] = b[7]
	rr[14] = b[6]
	rr[13] = b[5]
	rr[12] = b[4]
	rr[11] = b[3]
	rr[10] = b[2]
	rr[9] = b[1]
	rr[8] = b[0]
}

func (rr *Raid) setTimestamp(t time.Time) {
	s := t.Sub(timeOffset)
	binary.BigEndian.PutUint32(rr[2:6], uint32(s.Seconds()))
}

func (rr *Raid) setMessage(i uint16) {
	binary.BigEndian.PutUint16(rr[6:8], i)
}

func (rr *Raid) setPrefixByte(prefix []byte) {
	if len(prefix) < 3 {
		prefix = append(prefix, 0, 0, 0, 0)
	}
	p := make([]byte, 3)
	decodePrefix(p, prefix)
	rr[2] = p[2]
	rr[1] = p[1]
	rr[0] = p[0]
}

func (rr Raid) MarshalText() ([]byte, error) {
	b := make([]byte, 32)
	encodeRaid(b, rr[:])
	return b, nil
}

func (rr *Raid) UnmarshalText(b []byte) error {
	if len(b) != 32 {
		return ErrInvalidId
	}
	for _, c := range b {
		if dec[c] == 0xff {
			return ErrInvalidId
		}
	}
	decodeRaid(rr[:], b)
	return nil
}

func (rr Raid) MarshalJSON() ([]byte, error) {
	if rr == NilRaid {
		return stringToBytes("null"), nil
	}
	b := make([]byte, 34)
	b[33] = '"'
	b[0] = '"'
	encodeRaid(b[1:33], rr[:])
	return b, nil
}

func (rr *Raid) UnmarshalJSON(b []byte) error {
	if bytesToString(b) == "null" {
		*rr = NilRaid
		return nil
	}
	if len(b) != 34 {
		return ErrInvalidId
	}
	return rr.UnmarshallText(b[1:33])
}

func (rr *Raid) Scan(value any) error {
	switch v := value.(type) {
	case string:
		rr.UnmarshallText(stringToBytes(v))
	case []byte:
		rr.UnmarshallText(v)
	case nil:
		*rr = NilRaid
	default:
		return fmt.Errorf("raid: scanning unsupported type %T", value)
	}
	return nil
}

func (rr Raid) Value() (driver.Value, error) {
	if rr == NewRaid() {
		return nil, nil
	}
	b := make([]byte, 32)
	encodeRaid(b, rr[:])
	return bytesToString(b), nil
}

func (rr Raid) Bytes() []byte {
	return rr[:]
}

func (rr Raid) Encode(dst []byte) {
	encodeRaid(dst, rr[:])
}

func (rr Raid) String() string {
	b := make([]byte, 32)
	encodeRaid(b, rr[:])
	return bytesToString(b)
}

func (rr Raid) Prefix() string {
	p := make([]byte, 3)
	encodePrefix(p, rr[0:2])
	return bytesToString(p)
}

func (rr Raid) Time() time.Time {
	s := binary.BigEndian.Uint32(rr[2:6])
	return time.Unix(int64(s) + timeOffset.Unix(), 0)
}

func (rr Raid) Message() uint16 {
	c := binary.BigEndian.Uint16(rr[6:8])
	return uint16(c)
}

func (rr Raid) Compare(r Raid) int {
	return bytes.Compare(rr[:], r[:])
}

func (rr Raid) Less(r Raid) bool {
	return rr.Compare(r) == -1
}
