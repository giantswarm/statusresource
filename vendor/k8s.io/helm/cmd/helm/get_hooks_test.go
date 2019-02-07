/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"io"
	"testing"

	"github.com/spf13/cobra"

	"k8s.io/helm/pkg/helm"
	"k8s.io/helm/pkg/proto/hapi/release"
)

func TestGetHooks(t *testing.T) {
	tests := []releaseCase{
		{
			name:     "get hooks with release",
			args:     []string{"aeneas"},
			expected: fmt.Sprintf("---\n# %s\n%s\n", "pre-install-hook", helm.MockHookTemplate),
			resp:     helm.ReleaseMock(&helm.MockReleaseOptions{Name: "aeneas"}),
			rels:     []*release.Release{helm.ReleaseMock(&helm.MockReleaseOptions{Name: "aeneas"})},
		},
		{
			name: "get hooks without args",
			args: []string{},
			err:  true,
		},
	}
	runReleaseCases(t, tests, func(c *helm.FakeClient, out io.Writer) *cobra.Command {
		return newGetHooksCmd(c, out)
	})
}
