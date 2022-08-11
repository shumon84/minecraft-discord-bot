package console

import (
	"reflect"
	"testing"
	"time"
)

func TestReadLog(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    *Log
		wantErr bool
	}{
		{
			name: "正常系の場合",
			args: args{
				src: "[10:11:12] [Server thread/INFO]: shumon_84 joined the game",
			},
			want: &Log{
				Timestamp:  time.Hour*10 + time.Minute*11 + time.Second*12,
				ThreadName: "Server thread",
				LogLevel:   LogLevelInfo,
				Message:    "shumon_84 joined the game",
			},
			wantErr: false,
		},
		{
			name: "改行を含む場合",
			args: args{
				src: `[18:56:55] [Server console handler/ERROR]: Exception handling console input
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
			want: &Log{
				Timestamp:  time.Hour*18 + time.Minute*56 + time.Second*55,
				ThreadName: "Server console handler",
				LogLevel:   LogLevelError,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadLog(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLog() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadLog() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog_String(t *testing.T) {
	type fields struct {
		Timestamp  time.Duration
		ThreadName string
		LogLevel   LogLevel
		Message    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "正常系のテスト",
			fields: fields{
				Timestamp:  time.Hour*12 + time.Minute*13 + time.Second*14,
				ThreadName: "Thread Name",
				LogLevel:   LogLevelInfo,
				Message:    "sample text",
			},
			want: "[12:13:14] [Thread Name/INFO]: sample text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &Log{
				Timestamp:  tt.fields.Timestamp,
				ThreadName: tt.fields.ThreadName,
				LogLevel:   tt.fields.LogLevel,
				Message:    tt.fields.Message,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
