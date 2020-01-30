package internal

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"unsafe"

	"go.octolab.org/toolkit/github/internal/entity"
)

var repositoryURN = regexp.MustCompile(`^[^/]+/[^/]+$`)

// Repository represents a GitHub repository.
type Repository struct {
	ID     int64
	URN    RepositoryURN
	Labels []entity.Label
}

// RepositoryURN represents a full name of repository in the format "<owner>/<name>".
type RepositoryURN string

// ID returns the full name of repository with the URN scheme.
func (urn RepositoryURN) ID() string {
	return "urn:repo:" + string(urn)
}

// IsValid returns true if the full name of repository is satisfied with the pattern.
func (urn RepositoryURN) IsValid() bool {
	return repositoryURN.MatchString(string(urn))
}

// Name returns a repository name.
// It raises a panic if the full name of repository is not satisfied with the pattern.
func (urn RepositoryURN) Name() string {
	_, name := urn.Split()
	return name
}

// Owner returns a repository owner.
// It raises a panic if the full name of repository is not satisfied with the pattern.
func (urn RepositoryURN) Owner() string {
	owner, _ := urn.Split()
	return owner
}

// Split splits the full name of repository by name and owner.
// It raises a panic if the full name of repository is not satisfied with the pattern.
func (urn RepositoryURN) Split() (string, string) {
	if !urn.IsValid() {
		panic(fmt.Sprintf("%q doesn't satisfy the pattern %q", urn, repositoryURN))
	}
	parts := strings.Split(string(urn), "/")
	return parts[0], parts[1]
}

// String returns the string representation of the full name of repository.
func (urn RepositoryURN) String() string {
	return string(urn)
}

// RepositoryURNFromStrings converts a list of string
// to a list of RepositoryURN but not guarantee validity of the list.
func RepositoryURNFromStrings(list []string) []RepositoryURN {
	header := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	converted := (*[]RepositoryURN)(unsafe.Pointer(header))
	return *converted
}
