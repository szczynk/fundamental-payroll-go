package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerologger zerolog.Logger
}

func New(isDebug bool) *Logger {
	logLvl := zerolog.InfoLevel
	if isDebug {
		logLvl = zerolog.TraceLevel
	}

	zerolog.SetGlobalLevel(logLvl)
	output := zerolog.ConsoleWriter{Out: os.Stderr, NoColor: false}
	zerologger := zerolog.New(output).With().Timestamp().Logger()

	logger := new(Logger)
	logger.zerologger = zerologger
	return logger
}

func (l *Logger) Err(err error) *zerolog.Event {
	return l.zerologger.Err(err)
}

// Trace starts a new message with trace level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Trace() *zerolog.Event {
	return l.zerologger.Trace()
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.zerologger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.zerologger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.zerologger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.zerologger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.zerologger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.zerologger.Panic()
}
