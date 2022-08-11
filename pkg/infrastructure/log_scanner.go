package infrastructure

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/minecraft/console"
)

type LogScanner struct {
	sc         *bufio.Scanner
	scannedLog *console.Log
}

func NewLogScanner(r io.Reader) *LogScanner {
	sc := bufio.NewScanner(r)
	ls := &LogScanner{
		sc:         sc,
		scannedLog: nil,
	}
	ls.sc.Split(ls.scanLog)
	return ls
}

var logPrefix = regexp.MustCompile(`^\[[0-9]{2}:[0-9]{2}:[0-9]{2}] \[[^/]+/[^]]+]:`)

func (ls *LogScanner) scanLog(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 2 {
		return 0, nil, nil
	}
	if !logPrefix.MatchString(lines[0]) {
		return len([]byte(lines[0])) + 1, nil, nil
	}
	tokenString := lines[0]
	for _, line := range lines[1:] {
		if logPrefix.MatchString(line) {
			break
		}
		tokenString += "\n" + line
	}
	token = []byte(tokenString)
	if strings.HasSuffix(tokenString, "\n") {
		token = token[:len(token)-1]
	}
	scannedLog, err := console.ReadLog(string(token))
	if err != nil {
		panic(err)
	}
	ls.scannedLog = scannedLog
	return len(token) + 1, token, nil
}

func (ls *LogScanner) Scan() bool {
	return ls.sc.Scan()
}

func (ls *LogScanner) Err() error {
	return ls.sc.Err()
}

func (ls *LogScanner) Log() *console.Log {
	return ls.scannedLog
}
