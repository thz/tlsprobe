// Copyright 2020 Tobias Hintze
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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func usage() {
	fmt.Fprintf(os.Stderr, `Usage: probe endpoint-fqdn.example.com.:12345

Probe a given endpoint and be verbose about it.

Check DNS, TCP, TLS.
`)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "probe",
		Short: `A tool to probe TLS/TCP endpoints`,
	}

	rootCmd.AddCommand(probeCmd())
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
