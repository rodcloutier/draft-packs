// Copyright © 2017 Rodrigue Cloutier <rodcloutier@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/rodcloutier/draft-packs/pkg/draftpath"
	"github.com/spf13/cobra"
)

type packListCmd struct {
	home draftpath.Home
}

func init() {
	list := &packListCmd{
		home: draftpath.NewHome(os.ExpandEnv("$DRAFT_HOME")),
	}

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: "list packs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return list.run()
		},
	}

	RootCmd.AddCommand(cmd)
}

func (p *packListCmd) run() error {

	packHomeDir := p.home.Packs()
	files, err := ioutil.ReadDir(packHomeDir)
	if err != nil {
		return fmt.Errorf("there was an error reading %s: %v", packHomeDir, err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}

	return nil
}
