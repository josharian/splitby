package splitby

import (
	"bufio"
	"reflect"
	"regexp"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	const in = "abcabcabc"

	tests := []struct {
		sep  string
		want []string
	}{
		{sep: "a", want: []string{"", "bc", "bc", "bc"}},
		{sep: "b", want: []string{"a", "ca", "ca", "c"}},
		{sep: "c", want: []string{"ab", "ab", "ab"}},
		{sep: "ab", want: []string{"", "c", "c", "c"}},
		{sep: "bc", want: []string{"a", "a", "a"}},
		{sep: "abc", want: []string{"", "", ""}},
		{sep: "d", want: []string{"abcabcabc"}},
	}

	for _, tt := range tests {
		split := String(tt.sep)
		scanner := bufio.NewScanner(strings.NewReader(in))
		scanner.Split(split)
		var got []string
		for scanner.Scan() {
			got = append(got, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			t.Errorf("split %v using splitby.String(%v) error: %v", in, tt.sep, err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("split %v using splitby.String(%v) = %v, want %v", in, tt.sep, got, tt.want)
		}
	}
}

func TestRegexp(t *testing.T) {
	const in = "abcabcabc"

	tests := []struct {
		sep  string
		want []string
	}{
		{sep: "[a]", want: []string{"", "bc", "bc", "bc"}},
		{sep: "[b]", want: []string{"a", "ca", "ca", "c"}},
		{sep: "[c]", want: []string{"ab", "ab", "ab"}},
		{sep: "[ab]+", want: []string{"", "c", "c", "c"}},
		{sep: "[bc]+", want: []string{"a", "a", "a"}},
	}

	for _, tt := range tests {
		re := regexp.MustCompile(tt.sep)
		split := Regexp(re)
		scanner := bufio.NewScanner(strings.NewReader(in))
		scanner.Split(split)
		var got []string
		for scanner.Scan() {
			got = append(got, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			t.Errorf("split %v using splitby.String(%v) error: %v", in, tt.sep, err)
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("split %v using splitby.String(%v) = %v, want %v", in, tt.sep, got, tt.want)
		}
	}
}
