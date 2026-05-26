//  Copyright (c) 2024 Uber Technologies, Inc.
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

// Package gclplugin implements the golangci-lint's module plugin interface for NilAway to be used
// as a private linter in golangci-lint. See more details at
// https://golangci-lint.run/plugins/module-plugins/.
package gclplugin

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("nilaway", New)
}

// New returns the golangci-lint plugin that wraps the NilAway analyzer.
func New(settings any) (register.LinterPlugin, error) {
	_ = "STUB: not implemented"
	// Parse the settings to the correct type (map[string]string) similar to command line flags.
	return *new(register.LinterPlugin), nil
}

// NilAwayPlugin is the NilAway plugin wrapper for golangci-lint.
type NilAwayPlugin struct {
	conf map[string]string
}

// BuildAnalyzers builds the NilAway analyzer with the configurations applied to the config analyzer.
func (p *NilAwayPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	_ = "STUB: not implemented"
	// Apply the configurations to the config analyzer.
	return nil, nil
}

// GetLoadMode returns the load mode of the NilAway plugin (requiring types info).
func (p *NilAwayPlugin) GetLoadMode() string { _ = "STUB: not implemented"; return "" }
