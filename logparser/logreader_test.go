package logparser

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func readExactlyN(t *testing.T, ch <-chan LogEntry, n int) []LogEntry {
	const (
		readTimeout = 200 * time.Millisecond
		waitTimeout = 200 * time.Millisecond
	)

	entries := make([]LogEntry, 0, n)
	for i := 0; i < n; i++ {
		select {
		case entry := <-ch:
			entries = append(entries, entry)
		case <-time.After(readTimeout):
			t.Fatalf("timed out waiting for entry %d of %d", i+1, n)
		}
	}

	select {
	case entry := <-ch:
		t.Fatalf("unexpected extra entry received: %v", entry)
	case <-time.After(waitTimeout):
		// Success - no additional entries
	}

	return entries
}

func TestLogReader(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []LogEntry
	}{
		{
			name:  "single_entry",
			input: "2024-01-01 10:00:00 INFO Test message\n",
			expected: []LogEntry{
				{
					Timestamp: mustParseTime("2024-01-01 10:00:00"),
					Level:     LogLevelInfo,
					Message:   "Test message",
				},
			},
		},
		{
			name: "multiple_entries",
			input: `2024-01-01 10:00:00 DEBU Debug message
2024-01-01 10:00:01 INFO Info message
2024-01-01 10:00:02 WARN Warning message
2024-01-01 10:00:03 ERRO Error message
`,
			expected: []LogEntry{
				{
					Timestamp: mustParseTime("2024-01-01 10:00:00"),
					Level:     LogLevelDebug,
					Message:   "Debug message",
				},
				{
					Timestamp: mustParseTime("2024-01-01 10:00:01"),
					Level:     LogLevelInfo,
					Message:   "Info message",
				},
				{
					Timestamp: mustParseTime("2024-01-01 10:00:02"),
					Level:     LogLevelWarn,
					Message:   "Warning message",
				},
				{
					Timestamp: mustParseTime("2024-01-01 10:00:03"),
					Level:     LogLevelError,
					Message:   "Error message",
				},
			},
		},
		{
			name:     "empty_input",
			input:    "",
			expected: []LogEntry{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			reader := NewLogReader()
			reader.AddSource(strings.NewReader(tc.input))

			entries := readExactlyN(t, reader.Stream(), len(tc.expected))
			require.Equal(t, tc.expected, entries)
			reader.Close()
		})
	}
}

func TestMultipleSources(t *testing.T) {
	source1 := strings.NewReader("2024-01-01 10:00:00 INFO Source 1\n")
	source2 := strings.NewReader("2024-01-01 10:00:01 INFO Source 2\n")

	reader := NewLogReader()
	reader.AddSource(source1)
	time.Sleep(100 * time.Millisecond)
	reader.AddSource(source2)

	entries := readExactlyN(t, reader.Stream(), 2)
	require.Equal(t, "Source 1", entries[0].Message)
	require.Equal(t, "Source 2", entries[1].Message)
	reader.Close()
}

func TestBlockingReader(t *testing.T) {
	r, w := io.Pipe()

	reader := NewLogReader()
	reader.AddSource(r)

	go func() {
		_, err := w.Write([]byte("2024-01-01 10:00:00 INFO Test message\n"))
		require.NoError(t, err)
		err = w.Close()
		require.NoError(t, err)
	}()

	entries := readExactlyN(t, reader.Stream(), 1)
	require.Equal(t, "Test message", entries[0].Message)
	reader.Close()
}

func TestPartialReads(t *testing.T) {
	r := &chunkedReader{
		content: []byte("2024-01-01 10:00:00 INFO Test message\n"),
		chunks:  []int{10, 10, 10, 10},
	}

	reader := NewLogReader()
	reader.AddSource(r)

	entries := readExactlyN(t, reader.Stream(), 1)
	require.Equal(t, "Test message", entries[0].Message)
	reader.Close()
}

func TestCloseWithPendingReads(t *testing.T) {
	r := &chunkedReader{
		content: []byte("2024-01-01 10:00:00 INFO Never ending message\n"),
		chunks:  []int{1},
	}

	reader := NewLogReader()
	reader.AddSource(r)

	done := make(chan struct{})
	go func() {
		for range reader.Stream() {
		}
		close(done)
	}()

	time.Sleep(100 * time.Millisecond)
	reader.Close()

	select {
	case <-done:
		// Success
	case <-time.After(time.Second):
		t.Fatal("stream was not closed")
	}
}

type chunkedReader struct {
	content []byte
	chunks  []int
	pos     int
	chunk   int
}

func (r *chunkedReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.content) {
		return 0, io.EOF
	}
	if r.chunk >= len(r.chunks) {
		return 0, io.EOF
	}

	n = r.chunks[r.chunk]
	if r.pos+n > len(r.content) {
		n = len(r.content) - r.pos
	}
	copy(p, r.content[r.pos:r.pos+n])
	r.pos += n
	r.chunk++
	return n, nil
}

func mustParseTime(s string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		panic(err)
	}
	return t
}
