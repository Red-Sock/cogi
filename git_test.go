package cogi

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ParseOriginLink(t *testing.T) {
	type testCase struct {
		OriginString   string
		ExpectedOrigin *Origin
		ExpectedErr    error
	}

	testCases := map[string]testCase{
		"FOLDER": {
			OriginString: "github.com/Red-Sock",
			ExpectedErr:  ErrProtocolIsRequired,
		},
		"FOLDER_WITH_PROTOC": {
			OriginString: "https://github.com/Red-Sock",
			ExpectedOrigin: &Origin{
				Type: GithubOrigin,
				Url: url.URL{
					Scheme: "https",
					Host:   "github.com",
					Path:   "/Red-Sock",
				},
				Dir:  "Red-Sock",
				Repo: "*",
			},
		},

		"REPO": {
			OriginString: "github.com/Red-Sock/rscli",
			ExpectedErr:  ErrProtocolIsRequired,
		},
		"REPO_WITH_PROTOC": {
			OriginString: "https://github.com/Red-Sock/rscli",
			ExpectedOrigin: &Origin{
				Type: GithubOrigin,
				Url: url.URL{
					Scheme: "https",
					Host:   "github.com",
					Path:   "/Red-Sock/rscli",
				},
				Dir:  "Red-Sock",
				Repo: "rscli",
			},
		},

		"GITVERSE_REPO": {
			OriginString: "https://gitverse.ru/Red-Sock/rscli",
			ExpectedOrigin: &Origin{
				Type: GitVerseOrigin,
				Url: url.URL{
					Scheme: "https",
					Host:   "gitverse.ru",
					Path:   "/Red-Sock/rscli",
				},
				Dir:  "Red-Sock",
				Repo: "rscli",
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			origin, err := ParseGitSource(tc.OriginString)
			require.ErrorIs(t, err, tc.ExpectedErr)
			require.Equal(t, tc.ExpectedOrigin, origin)
		})
	}
}
