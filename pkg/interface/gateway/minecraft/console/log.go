package console

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var layout = regexp.MustCompile(`^\[([0-9]{2}):([0-9]{2}):([0-9]{2})] \[([^/]+)/([^]]+)]: ([\s\S]*)$`)

var (
	ErrInvalidLogLayout = errors.New("invalid log layout error")
	ErrInvalidLogLevel  = errors.New("invalid log level error")
)

type Log struct {
	Timestamp  time.Duration
	ThreadName string
	LogLevel   LogLevel
	Message    string
}

func ReadLog(src string) (*Log, error) {
	subMatch := layout.FindStringSubmatch(src)
	if len(subMatch) != 7 {
		return nil, fmt.Errorf("%w : %s", ErrInvalidLogLayout, src)
	}
	hours, err := strconv.Atoi(subMatch[1])
	if err != nil {
		return nil, err
	}
	minutes, err := strconv.Atoi(subMatch[2])
	if err != nil {
		return nil, err
	}
	seconds, err := strconv.Atoi(subMatch[3])
	if err != nil {
		return nil, err
	}
	timestamp := time.Hour*time.Duration(hours) + time.Minute*time.Duration(minutes) + time.Second*time.Duration(seconds)
	threadName := subMatch[4]
	level := ToLevel(subMatch[5])
	if level == LogLevelUndefined {
		return nil, fmt.Errorf("%w : %s", ErrInvalidLogLevel, src)
	}
	message := subMatch[6]
	return &Log{
		Timestamp:  timestamp,
		ThreadName: threadName,
		LogLevel:   level,
		Message:    message,
	}, nil
}

func (l *Log) String() string {
	return fmt.Sprintf("[%02d:%02d:%02d] [%s/%s]: %s",
		int(l.Timestamp.Hours())%60,
		int(l.Timestamp.Minutes())%60,
		int(l.Timestamp.Seconds())%60,
		l.ThreadName,
		l.LogLevel,
		l.Message,
	)
}
