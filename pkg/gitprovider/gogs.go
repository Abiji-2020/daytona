// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package gitprovider 

import (
	"fmt"
)

type GogsGitProvider struct {
	*AbstractGitProvider
	token string 
	baseApiUrl string
}

fun NewGogsGitProvider(token string, baseApiUrl string) *GogsGitProvider {
	gitProvider := &GogsGitProvider{
		token: token,
		baseApiUrl: baseApiUrl,
		AbstractGitProvider: &AbstractGitProvider{},
	}
	gitProvider.AbstractGitProvider.gitProvider = gitProvider
	return gitProvider
}

func (g *GogsGitProvider) GetNamespaces() (*[]GitNamespace, error){
	
}