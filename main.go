// SPDX-License-Identifier: BSD-3-Clause
// Copyright (c) 2024, Alexander Jung.
// Licensed under the BSD-3-Clause License (the "License").
// You may not use this file except in compliance with the License.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
	"kraftkit.sh/kconfig"
)

type CLI struct {
	Vars []string `kong:"short=v,long=env,help='KConfig variables used when evaluating.'"`
	Root string   `kong:"arg,name='ROOT',type=existingfile,help='Path to root KConfig file.'"`
}

func main() {
	cli := CLI{}

	_ = kong.Parse(&cli)

	var kvs []*kconfig.KeyValue

	for _, v := range cli.Vars {
		_, kv := kconfig.NewKeyValue(v)
		if kv == nil {
			continue
		}

		kvs = append(kvs, kv)
	}

	root, err := filepath.Abs(cli.Root)
	if err != nil {
		fmt.Printf("could not get absolute path: %v", err)
		os.Exit(1)
	}

	rootDir := filepath.Dir(root)

	tree, err := kconfig.Parse(root, kvs...)
	if err != nil {
		fmt.Printf("could not parse KConfig file: %v", err)
		os.Exit(1)
	}

	if err := tree.Walk(func(menu *kconfig.KConfigMenu) error {
		fmt.Printf("- name: %s\n", menu.Name)
		fmt.Printf("  kind: %s\n", menu.Kind)
		fmt.Printf("  type: %s\n", menu.Type)
		if menu.Prompt.Text != "" {
			prompt := strings.ReplaceAll(menu.Prompt.Text, "\"", "'")
			fmt.Printf("  prompt: |\n   %s\n", prompt)
		}
		if menu.Prompt.Condition != nil {
			cond := menu.Prompt.Condition.String()
			cond = strings.ReplaceAll(cond, "\"$(", "$(")
			cond = strings.ReplaceAll(cond, ")\"", ")")
			fmt.Printf("  when: %s\n", cond)
		}
		fmt.Printf("  file: %s\n", strings.TrimPrefix(menu.Source, rootDir)[1:])
		if menu.Default.Value != nil {
			fmt.Printf("  default: \n")
			value := menu.Default.Value.String()
			value = strings.ReplaceAll(value, "\"$(", "$(")
			value = strings.ReplaceAll(value, ")\"", ")")
			value = strings.ReplaceAll(value, "\"", "\\\"")
			fmt.Printf("    value: \"%s\"\n", value)
			if menu.Default.Condition != nil {
				fmt.Printf("    when: \"%s\"\n", strings.ReplaceAll(menu.Default.Condition.String(), "\"", "\\\""))
			}
		}

		fmt.Println()

		return nil
	}); err != nil {
		fmt.Printf("could not walk KConfig tree: %v", err)
		os.Exit(1)
	}

	os.Exit(1)
}
