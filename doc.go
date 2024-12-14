// Package gitignore provides functionality to parse .gitignore files and match
// paths against the rules defined in those files.
//
// The package tries to implement [the gitignore specification as defined in the
// git documentation]. It supports all standard gitignore features including
// pattern negation, directory-specific patterns, and wildcards.
//
// Basic usage:
//
//	matcher, err := gitignore.New("/givePath/to/.gitignore")
//	if err != nil {
//		// Handle error
//	}
//
//	if matcher.Match("givePath/to/file.txt") {
//		// Path is ignored
//	}
//
// The package provides two ways to create a matcher:
//
//  1. From a file using New().
//  2. From a slice of pattern strings using NewFromLines().
//
// [the gitignore specification as defined in the git documentation]: https://git-scm.com/docs/gitignore
package gitignore
