// Copyright 2018-2019 VMware, Inc. 
// SPDX-License-Identifier: Apache-2.0

package util

// Describes how to match a given string in HTTP headers. Match is
// case-sensitive.
type StringMatch struct {
	// Specified exactly one of the fields below.

	// exact string match
	Exact string `json:"exact,omitempty"`

	// prefix-based match
	Prefix string `json:"prefix,omitempty"`

	// suffix-based match.
	Suffix string `json:"prefix,omitempty"`

	// ECMAscript style regex-based match
	Regex string `json:"regex,omitempty"`
}
