package logger

import (
    "io"
    "os"
    "log"
);

type Log interface {
    Debug() *log.Logger
    Info() *log.Logger
    Warn() *log.Logger
    Error() *log.Logger
}

type Logging struct {
    debug *log.Logger
    info *log.Logger
    warn *log.Logger
    error *log.Logger
}

func Init() Log {
    return InitCustom(os.Stdout)
}

func InitCustom(writer io.Writer) Log {
    debug := log.New(writer, "Debug: ", log.Ldate|log.Ltime|log.Lshortfile)
    info := log.New(writer, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
    warn := log.New(writer, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
    error := log.New(writer, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)

    return &Logging{debug: debug, info: info, warn: warn, error: error}
}

func (l *Logging) Debug() *log.Logger {
    return l.debug
}

func (l *Logging) Info() *log.Logger {
    return l.info
}

func (l *Logging) Warn() *log.Logger {
    return l.warn
}

func (l *Logging) Error() *log.Logger {
    return l.error
}
