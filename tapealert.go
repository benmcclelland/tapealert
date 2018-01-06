package tapealert

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

// TapeAlert holds fields of the tapealert info
type TapeAlert struct {
	Flag     int64
	Name     string
	Severity string
	Message  string
	Cause    string
	sevnum   int
}

// Session stores info about current session
type Session struct {
	standalone bool
	alerts     map[uint]TapeAlert
}

// New initializes alerts
func New(opts ...Option) (*Session, error) {
	s := &Session{}
	for _, opt := range opts {
		opt(s)
	}

	var csvdata string
	if s.standalone {
		csvdata = sacsv
	} else {
		csvdata = tpcsv
	}

	r := csv.NewReader(strings.NewReader(csvdata))
	lineCount := uint(0)
	s.alerts = make(map[uint]TapeAlert)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		if lineCount == 0 {
			// skip header
			lineCount++
			continue
		}
		if len(record) != 5 {
			return nil, fmt.Errorf("%s %d",
				"invalid csv format: expected 5 fields got", len(record))
		}
		var sev int
		switch record[2] {
		case "I":
			sev = 0
		case "W":
			sev = 1
		case "C":
			sev = 2
		default:
			return nil, fmt.Errorf("%s %s",
				"invalid csv format: expected severity [IWC] got", record[2])
		}
		s.alerts[lineCount] = TapeAlert{
			Flag:     1 << (lineCount - 1),
			Name:     record[1],
			Severity: record[2],
			Message:  record[3],
			Cause:    record[4],
			sevnum:   sev,
		}
		lineCount++
	}

	return s, nil
}

// Option sets various options for New
type Option func(*Session)

// WithStandalone initiales a standalone changer
func WithStandalone() Option {
	return func(s *Session) { s.standalone = true }
}

// FromFlags returns slice of TapeAlerts matching passed in flags
// with a min severity of:
// 0 = Info (all alerts)
// 1 = Warning
// 2 = Critical
func (s *Session) FromFlags(flags int64, severity int) []TapeAlert {
	var ta []TapeAlert
	for i := uint(0); i < 64; i++ {
		if flags&(1<<i) != 0 {
			if s.alerts[i+1].sevnum >= severity {
				ta = append(ta, s.alerts[i+1])
			}
		}
	}

	return ta
}
