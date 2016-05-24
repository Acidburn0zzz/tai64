package tai64n

//go:generate go run ./tools/generate.go -pkg $GOPACKAGE -output offsets.go

import (
	"fmt"
	"strconv"
	"time"
)

const tai64Epoch = 2 << 61

func getOffset(utime int64) int64 {
	// default offset is 10
	offset := int64(10)
	for i := tia64nSize - 1; i >= 0; i-- {
		if utime < tia64nDifferences[i].utime {
			continue
		} else {
			offset = tia64nDifferences[i].offset
			break
		}
	}
	return offset
}

func getInvOffset(utime int64) int64 {
	// default offset is 10
	offset := int64(10)
	for i := tia64nSize - 1; i >= 0; i-- {
		t := tia64nDifferences[i]
		if utime < (t.utime + t.offset) {
			continue
		} else {
			offset = t.offset
			break
		}
	}
	return offset
}

// Format formats a time.Time as a TAI64N timestamp
// returns a string TAI64N timestamps
func Format(t time.Time) string {
	t = t.UTC()
	u := t.Unix()

	if u < 0 {
		return fmt.Sprintf("@%016x%08x", (2<<61)+u+10, t.Nanosecond())
	}
	return fmt.Sprintf("@4%015x%08x", u+getOffset(u), t.Nanosecond())
}

// Parse parses a TAI64N timestamp
// returns a time.Time and an error.
func Parse(s string) (time.Time, error) {
	var seconds, nanoseconds int64
	if s[0] == '@' {
		s = s[1:]
	}

	if len(s) < 16 {
		return time.Time{}, fmt.Errorf("invalid tai64 time string")
	}

	i, err := strconv.ParseInt(s[0:16], 16, 64)
	if err != nil {
		return time.Time{}, err
	}
	seconds = i
	s = s[16:]

	if len(s) == 8 {
		i, err := strconv.ParseInt(s[0:8], 16, 64)
		if err != nil {
			return time.Time{}, err
		}
		nanoseconds = i
	}

	if seconds >= tai64Epoch {
		// fiddle with add/remove time
		unix := seconds - tai64Epoch
		offset := getInvOffset(unix)
		unix = unix - offset
		t := time.Unix(unix, nanoseconds).UTC()
		return t, nil
	}

	unix := -(tai64Epoch - seconds + 10)
	t := time.Unix(unix, nanoseconds).UTC()
	return t, nil
}
