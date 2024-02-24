package github

import "errors"

var (
	ErrInvalidGithubURL    = errors.New("invalid github url")
	ErrNoOrganizationFound = errors.New("no organization found")
)
