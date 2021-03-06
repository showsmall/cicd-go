// Copyright 2018 Drone.IO Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package secret

import (
	"context"

	"github.com/banzaicloud/cicd-go/cicd"
)

// V1 is version 1 of the secrets API
const V1 = "application/vnd.cicd.secret.v1+json"

type (
	// Request defines a secret request.
	Request struct {
		Name  string     `json:"name"`
		Repo  cicd.Repo  `json:"repo,omitempty"`
		Build cicd.Build `json:"build,omitempty"`
	}

	// Plugin responds to a secret request.
	Plugin interface {
		Find(context.Context, *Request) (*cicd.Secret, error)
	}
)
