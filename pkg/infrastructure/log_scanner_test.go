package infrastructure

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/shumon84/minecraft-discord-bot/pkg/interface/gateway/minecraft/console"
)

func TestLogScanner_Log(t *testing.T) {
	tests := []struct {
		name    string
		scanner *LogScanner
		want    *console.Log
	}{
		{
			name:    "正常系のテスト",
			scanner: NewLogScanner(bytes.NewBufferString("[18:56:56] [Server thread/INFO]: This server is running CraftBukkit version 3545-Spigot-475f600-4230f8f (MC: 1.19) (Implementing API version 1.19-R0.1-SNAPSHOT)\n")),
			want: &console.Log{
				Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*56,
				ThreadName: "Server thread",
				LogLevel:   console.LogLevelInfo,
				Message:    "This server is running CraftBukkit version 3545-Spigot-475f600-4230f8f (MC: 1.19) (Implementing API version 1.19-R0.1-SNAPSHOT)",
			},
		},
		{
			name: "改行を含む場合",
			scanner: NewLogScanner(bytes.NewBufferString(`[18:56:55] [Server console handler/ERROR]: Exception handling console input
java.io.IOException: Bad file descriptor
        at java.io.FileInputStream.readBytes(Native Method) ~[?:?]
        at java.io.FileInputStream.read(FileInputStream.java:276) ~[?:?]
        at java.io.BufferedInputStream.fill(BufferedInputStream.java:244) ~[?:?]
        at java.io.BufferedInputStream.read(BufferedInputStream.java:263) ~[?:?]
        at jline.internal.NonBlockingInputStream.read(NonBlockingInputStream.java:248) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:261) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:198) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readCharacter(ConsoleReader.java:2145) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLineSimple(ConsoleReader.java:3183) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2333) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2269) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2257) ~[jline-2.12.1.jar:?]
        at net.minecraft.server.dedicated.DedicatedServer$1.run(DedicatedServer.java:124) [spigot-1.19-R0.1-SNAPSHOT.jar:3545-Spigot-475f600-4230f8f]
[18:56:56] [Server thread/INFO]: This server is running CraftBukkit version 3545-Spigot-475f600-4230f8f (MC: 1.19) (Implementing API version 1.19-R0.1-SNAPSH
OT)
`)),
			want: &console.Log{
				Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*55,
				ThreadName: "Server console handler",
				LogLevel:   console.LogLevelError,
				Message: `Exception handling console input
java.io.IOException: Bad file descriptor
        at java.io.FileInputStream.readBytes(Native Method) ~[?:?]
        at java.io.FileInputStream.read(FileInputStream.java:276) ~[?:?]
        at java.io.BufferedInputStream.fill(BufferedInputStream.java:244) ~[?:?]
        at java.io.BufferedInputStream.read(BufferedInputStream.java:263) ~[?:?]
        at jline.internal.NonBlockingInputStream.read(NonBlockingInputStream.java:248) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:261) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:198) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readCharacter(ConsoleReader.java:2145) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLineSimple(ConsoleReader.java:3183) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2333) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2269) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2257) ~[jline-2.12.1.jar:?]
        at net.minecraft.server.dedicated.DedicatedServer$1.run(DedicatedServer.java:124) [spigot-1.19-R0.1-SNAPSHOT.jar:3545-Spigot-475f600-4230f8f]`,
			},
		},
		{
			name: "log じゃない行を含む場合",
			scanner: NewLogScanner(bytes.NewBufferString(`not log line
[18:56:55] [Server console handler/ERROR]: Exception handling console input
`)),
			want: &console.Log{
				Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*55,
				ThreadName: "Server console handler",
				LogLevel:   console.LogLevelError,
				Message:    `Exception handling console input`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.scanner.Scan()
			if got := tt.scanner.Log(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogScanner_MultiLine(t *testing.T) {
	tests := []struct {
		name    string
		scanner *LogScanner
		want    []*console.Log
	}{
		{
			name: "正常系のテスト",
			scanner: NewLogScanner(bytes.NewBufferString(
				`[18:56:56] [Server thread/INFO]: This server is running CraftBukkit version 3545-Spigot-475f600-4230f8f (MC: 1.19) (Implementing API version 1.19-R0.1-SNAPSHOT)
[19:35:29] [Server thread/INFO]: Se_723_A has made the advancement [Monster Hunter]
[19:35:43] [Server thread/INFO]: Master_Tofu was slain by Zombie
[19:36:04] [Server thread/INFO]: shumon_84 tried to swim in lava
[18:56:55] [Server console handler/ERROR]: Exception handling console input
java.io.IOException: Bad file descriptor
        at java.io.FileInputStream.readBytes(Native Method) ~[?:?]
        at java.io.FileInputStream.read(FileInputStream.java:276) ~[?:?]
        at java.io.BufferedInputStream.fill(BufferedInputStream.java:244) ~[?:?]
        at java.io.BufferedInputStream.read(BufferedInputStream.java:263) ~[?:?]
        at jline.internal.NonBlockingInputStream.read(NonBlockingInputStream.java:248) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:261) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:198) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readCharacter(ConsoleReader.java:2145) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLineSimple(ConsoleReader.java:3183) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2333) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2269) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2257) ~[jline-2.12.1.jar:?]
        at net.minecraft.server.dedicated.DedicatedServer$1.run(DedicatedServer.java:124) [spigot-1.19-R0.1-SNAPSHOT.jar:3545-Spigot-475f600-4230f8f]
`)),
			want: []*console.Log{
				{
					Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*56,
					ThreadName: "Server thread",
					LogLevel:   console.LogLevelInfo,
					Message:    "This server is running CraftBukkit version 3545-Spigot-475f600-4230f8f (MC: 1.19) (Implementing API version 1.19-R0.1-SNAPSHOT)",
				},
				{
					Timestamp:  time.Hour*19 + time.Minute*35 + time.Second*29,
					ThreadName: "Server thread",
					LogLevel:   console.LogLevelInfo,
					Message:    "Se_723_A has made the advancement [Monster Hunter]",
				},
				{
					Timestamp:  time.Hour*19 + time.Minute*35 + time.Second*43,
					ThreadName: "Server thread",
					LogLevel:   console.LogLevelInfo,
					Message:    "Master_Tofu was slain by Zombie",
				},
				{
					Timestamp:  time.Hour*19 + time.Minute*36 + time.Second*4,
					ThreadName: "Server thread",
					LogLevel:   console.LogLevelInfo,
					Message:    "shumon_84 tried to swim in lava",
				},
				{
					Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*55,
					ThreadName: "Server console handler",
					LogLevel:   console.LogLevelError,
					Message: `Exception handling console input
java.io.IOException: Bad file descriptor
        at java.io.FileInputStream.readBytes(Native Method) ~[?:?]
        at java.io.FileInputStream.read(FileInputStream.java:276) ~[?:?]
        at java.io.BufferedInputStream.fill(BufferedInputStream.java:244) ~[?:?]
        at java.io.BufferedInputStream.read(BufferedInputStream.java:263) ~[?:?]
        at jline.internal.NonBlockingInputStream.read(NonBlockingInputStream.java:248) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:261) ~[jline-2.12.1.jar:?]
        at jline.internal.InputStreamReader.read(InputStreamReader.java:198) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readCharacter(ConsoleReader.java:2145) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLineSimple(ConsoleReader.java:3183) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2333) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2269) ~[jline-2.12.1.jar:?]
        at jline.console.ConsoleReader.readLine(ConsoleReader.java:2257) ~[jline-2.12.1.jar:?]
        at net.minecraft.server.dedicated.DedicatedServer$1.run(DedicatedServer.java:124) [spigot-1.19-R0.1-SNAPSHOT.jar:3545-Spigot-475f600-4230f8f]`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			count := 0
			for tt.scanner.Scan() {
				got := tt.scanner.Log()
				if !reflect.DeepEqual(got, tt.want[count]) {
					t.Errorf("Log(): got : %v", got.String())
					t.Errorf("       want: %v", tt.want[count].String())
				}
				count++
			}
			if err := tt.scanner.Err(); err != nil {
				t.Errorf("scan error: %v", err)
			}
			if len(tt.want) != count {
				t.Errorf("count = %v, want %v", count, len(tt.want))
			}
		})
	}
}
