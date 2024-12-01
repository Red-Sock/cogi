package cogi

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testFolderForCloning = "./testData/"
)

func Test_GithubClone(t *testing.T) {
	ghManager := newGithub()
	t.Parallel()

	type testCase struct {
		Origin Origin
	}

	testCases := map[string]testCase{
		"OK_ORGANISATION": {
			Origin: Origin{
				Dir:  "Red-Sock",
				Repo: "*",
			},
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			tempDir := testFolderForCloning + t.Name()

			require.NoError(t, os.RemoveAll(tempDir))
			defer func() { require.NoError(t, os.RemoveAll(tempDir)) }()

			cr, err := ghManager.Clone(tempDir, tc.Origin)
			require.NoError(t, err)
			_ = cr
		})
	}
}
