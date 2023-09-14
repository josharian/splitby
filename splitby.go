package splitby

import (
	"bufio"
	"bytes"
	"errors"
)

// String returns a bufio.SplitFunc that splits on s.
// It is similar to bufio.ScanLines, but it splits on s instead of '\n' (and we pretend '\r' doesn't exist).
func String(s string) bufio.SplitFunc {
	if s == "" {
		return errorSplitFunc(errors.New("empty splitby.String separator"))
	}
	return Regexp(&bytesIndexFinder{sep: []byte(s)})
}

// Bytes returns a bufio.SplitFunc that splits on sep.
// It is similar to bufio.ScanLines, but it splits on sep instead of '\n' (and we pretend '\r' doesn't exist).
func Bytes(sep []byte) bufio.SplitFunc {
	if len(sep) == 0 {
		return errorSplitFunc(errors.New("empty splitby.Bytes separator"))
	}
	return Regexp(&bytesIndexFinder{sep: sep})
}

// Regexp returns a bufio.SplitFunc that splits using re.FindIndex.
func Regexp(re IndexFinder) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}
		index := re.FindIndex(data)
		if index != nil {
			// Avoid looping forever.
			if index[0] == index[1] {
				return 0, nil, errors.New("empty token")
			}
			// Matched full token.
			return index[1], data[:index[0]], nil
		}
		// Return the final token.
		if atEOF {
			return len(data), data, nil
		}
		// Request more data.
		return 0, nil, nil
	}
}

// IndexFinder corresponds to the regexp.IndexFunc method.
type IndexFinder interface {
	// FindIndex returns a two-element slice of integers defining the location of the leftmost match in data.
	// The match is at data[loc[0]:loc[1]].
	// A return value of nil indicates no match.
	FindIndex(data []byte) []int
}

type bytesIndexFinder struct {
	sep []byte
}

func (b *bytesIndexFinder) FindIndex(data []byte) []int {
	idx := bytes.Index(data, b.sep)
	if idx >= 0 {
		return []int{idx, idx + len(b.sep)}
	}
	return nil
}

func errorSplitFunc(err error) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		return 0, nil, err
	}
}
