// Copyright 2023 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

//go:build windows
// +build windows

// Package common contains common constants and functions needed to be accessed across ssm-setup-cli
package common

import (
	"strings"

	"github.com/aws/amazon-ssm-agent/agent/appconfig"
)

func setPlatformSpecificCommand(parts []string) []string {
	cmd := appconfig.PowerShellPluginCommandName + " -ExecutionPolicy unrestricted"
	return append(strings.Split(cmd, " "), parts...)
}
