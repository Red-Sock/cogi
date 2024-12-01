package cogi

import (
	"net/url"
	"path"
	"strings"

	errors "github.com/Red-Sock/trace-errors"
)

var (
	ErrProtocolIsRequired = errors.New("protocol (http|https) is required")
)

type originType string

type Origin struct {
	Type originType
	Url  url.URL
	Dir  string
	Repo string
}

const (
	GithubOrigin     originType = "github.com"
	GitVerseOrigin   originType = "gitverse.ru"
	GitGenericOrigin originType = "generic"

	allRepoInOrganization = "*"
)

func ParseGitSource(urlIn string) (*Origin, error) {
	if !strings.HasPrefix(urlIn, "http") {
		return nil, errors.Wrap(ErrProtocolIsRequired)
	}

	link, err := url.Parse(urlIn)
	if err != nil {
		return nil, errors.Wrap(err, "error parsing url")
	}

	dir := path.Dir(link.Path)
	repo := path.Base(link.Path)
	if dir == "/" {
		dir = repo
	} else {
		dir = dir[1:]
	}

	o := &Origin{
		Type: getGitType(link.Host),
		Url:  *link,
		Dir:  dir,
	}

	if repo == dir {
		o.Repo = allRepoInOrganization
	} else {
		o.Repo = repo
	}

	return o, nil
}

func getGitType(s string) originType {
	switch originType(s) {
	case GithubOrigin:
		return GithubOrigin
	case GitVerseOrigin:
		return GitVerseOrigin
	default:
		return GitGenericOrigin
	}
}
