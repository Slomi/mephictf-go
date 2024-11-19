package logparser

import (
	"io"
	"time"
)

type LogLevel int

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// LogEntry represents a single log entry.
// It contains the timestamp when the log was created,
// the log level (DEBU, INFO, WARN, ERRO), and
// the actual log message text.
type LogEntry struct {
	Timestamp time.Time // The time when the log entry was created
	Level     LogLevel  // Log level (DEBUG, INFO, WARNING, ERROR)
	Message   string    // The actual log message content
}

// LogReader provides functionality for reading and parsing log entries from multiple sources.
// It reads strings in the following format:
//
//	2024-01-01 10:00:00 DEBU This is a debug message
//	2024-01-01 10:00:01 INFO This is an info message
//	2024-01-01 10:00:02 WARN This is a warning message
//	2024-01-01 10:00:03 ERRO This is an error message
//
// The timestamp must be in "2006-01-02 15:04:05" format
// The log level must be one of: DEBU, INFO, WARN, ERRO
// The message can contain any text and is separated from the level by a space.
type LogReader struct {
}

// NewLogReader creates and returns a new instance of LogReader.
// The returned LogReader is ready to accept log sources through AddSource().
func NewLogReader() *LogReader {
	return nil
}

// AddSource adds a new log source to the reader.
// Multiple sources can be added and will be processed concurrently.
// The reader parameter should provide log entries in the expected format.
// Invalid log entries will be skipped.
func (lr *LogReader) AddSource(reader io.Reader) {
}

// Stream returns a receive-only channel of LogEntry.
// The channel will receive parsed log entries from all added sources
// in chronological order. The channel will be closed when Close() is called.
func (lr *LogReader) Stream() <-chan LogEntry {
	return nil
}

// Close stops all reading operations and closes the output channel.
// After calling Close(), no more log entries will be processed.
func (lr *LogReader) Close() {
}
