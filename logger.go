package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

type severity int

// Severity levels.
const (
	sInfo severity = iota
	sWarning
	sError
	sFatal
)

// Severity tags.
const (
	tagInfo    = "INFO : "
	tagWarning = "WARN : "
	tagError   = "ERROR: "
	tagFatal   = "FATAL: "
)

//Log format flags
const (
	LDate         = 1 << iota                      // the date in the local time zone: 2009/01/23
	LTime                                          // the time in the local time zone: 01:23:23
	LMicroseconds                                  // microsecond resolution: 01:23:23.123123.  assumes LTime.
	LLongFile                                      // full file name and line number: /a/b/c/d.go:23
	LShortFile                                     // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                                           // if Ldate or Ltime is set, use UTC rather than the local time zone
	LStdFlags     = LDate | LTime | log.Lshortfile // initial values for the standard logger
)

var (
	logLock        sync.Mutex
	defaultLogFile io.Writer
	flag           int
)

// set default log file to stdout when moule init
func init() {
	logLock.Lock()
	defer logLock.Unlock()
	defaultLogFile = os.Stdout
	flag = LStdFlags
}

// Init sets up output to non default and should be called before log functions, usually in
// the caller's main(). Default log output is stdout
// If the logFile passed in also satisfies io.Closer, logFile.Close will be called
// when closing the logger.
func Init(logFile io.Writer, flags int) {
	logLock.Lock()
	defer logLock.Unlock()

	defaultLogFile = logFile
	flag = flags
}

// Close closes the default logger.
func Close() {
	logLock.Lock()
	defer logLock.Unlock()

	if defaultLogFile == nil {
		return
	}

	if c, ok := defaultLogFile.(io.Closer); ok && c != nil {
		_ = c.Close()
	}
}

// Info uses the default logger and logs with the Info severity.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	output(sInfo, 2, fmt.Sprint(v...))
}

// Infof uses the default logger and logs with the Info severity.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	output(sInfo, 2, fmt.Sprintf(format, v...))
}

// Warning uses the default logger and logs with the Warning severity.
// Arguments are handled in the manner of fmt.Print.
func Warning(v ...interface{}) {
	output(sWarning, 2, fmt.Sprint(v...))
}

// Warningf uses the default logger and logs with the Warning severity.
// Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, v ...interface{}) {
	output(sWarning, 2, fmt.Sprintf(format, v...))
}

// Error uses the default logger and logs with the Error severity.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	output(sError, 2, fmt.Sprint(v...))
}

// Errorf uses the default logger and logs with the Error severity.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	output(sError, 2, fmt.Sprintf(format, v...))
}

// Fatal uses the default logger, logs with the Fatal severity,
// and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {
	output(sFatal, 2, fmt.Sprint(v...))
	Close()
	os.Exit(1)
}

// Fatalf uses the default logger, logs with the Fatal severity,
// and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	output(sFatal, 2, fmt.Sprintf(format, v...))
	Close()
	os.Exit(1)
}

func output(level severity, callDepth int, message string) {
	now := time.Now() // get this early.
	var file string
	var line int
	logLock.Lock()
	defer logLock.Unlock()

	if flag&(LShortFile|LLongFile) != 0 {
		// Release lock while getting caller info - it's expensive.
		logLock.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		logLock.Lock()
	}

	buf := formatHeader(message, level, now, file, line)
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	_, _ = defaultLogFile.Write(buf)
}

func formatHeader(message string, level severity, t time.Time, file string, line int) []byte {
	buf := make([]byte, 0)

	if flag&(LDate|LTime|LMicroseconds) != 0 {
		if flag&LUTC != 0 {
			t = t.UTC()
		}
		if flag&LDate != 0 {
			buf = append(buf, []byte(t.Format("2006/01/02"))...)
		}
		if flag&(LTime|LMicroseconds) != 0 {
			buf = append(buf, []byte(t.Format(" 15:04:05"))...)
			if flag&LMicroseconds != 0 {
				buf = append(buf, []byte(t.Format(" .000000"))...)
			}
			buf = append(buf, ' ')
		}
	}

	if flag&(LShortFile|LLongFile) != 0 {
		if flag&LShortFile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		buf = append(buf, fmt.Sprintf("%s:%d ", file, line)...)
	}

	switch level {
	case sInfo:
		buf = append(buf, tagInfo...)
	case sWarning:
		buf = append(buf, tagWarning...)
	case sError:
		buf = append(buf, tagError...)
	case sFatal:
		buf = append(buf, tagFatal...)
	default:
		panic(fmt.Sprintln("unrecognized severity:", level))
	}

	return append(buf, message...)
}
