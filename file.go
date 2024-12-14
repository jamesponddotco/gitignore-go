package gitignore

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~jamesponddotco/gitignore-go/internal/pattern"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
)

// ErrRegexCompile is returned when an error occurs while compiling regular
// expressions when parsing a .gitignore file.
const ErrRegexCompile xerrors.Error = "failed to compile regex"

// File represents a .gitignore file and provides the functionality to match
// paths against its rules.
type File struct {
	patterns []*pattern.Pattern
}

// New creates a new File instance from a given .gitignore file.
func New(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer file.Close()

	patterns, err := pattern.Parse(file)
	if err != nil {
		if errors.Is(err, pattern.ErrInvalidRegex) {
			return nil, fmt.Errorf("%w: %w", ErrRegexCompile, err)
		}

		return nil, fmt.Errorf("%w", err)
	}

	return &File{
		patterns: patterns,
	}, nil
}

// NewFromLines creates a new File instance from a list of strings. Useful when
// patterns are available in memory rather than in a file or for testing.
func NewFromLines(lines []string) (*File, error) {
	r := strings.NewReader(xstrings.JoinWithSeparator("\n", lines...))

	patterns, err := pattern.Parse(r)
	if err != nil {
		if errors.Is(err, pattern.ErrInvalidRegex) {
			return nil, fmt.Errorf("%w: %w", ErrRegexCompile, err)
		}

		return nil, fmt.Errorf("%w", err)
	}

	return &File{
		patterns: patterns,
	}, nil
}

// Match checks if the given path matches any of the .gitignore rules, and
// return true if the path should be ignored according to the rules.
//
// The path is normalized to use forward slashes (/) regardless of the operating
// system.
func (f *File) Match(path string) bool {
	path = strings.ReplaceAll(path, string(os.PathSeparator), "/")

	var match bool

	for _, pat := range f.patterns {
		if pat.Regex.MatchString(path) {
			if pat.Negate {
				return false
			}

			match = true
		}
	}

	return match
}
