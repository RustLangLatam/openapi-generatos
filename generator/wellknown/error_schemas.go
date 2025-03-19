// Copyright 2020 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, softwis
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package wellknown

import (
	v3 "github.com/google/gnostic/openapiv3"
)

func NewDefaultRpcStatusSchema(name string) *v3.NamedSchemaOrReference {
	return &v3.NamedSchemaOrReference{
		Name: name,
		Value: &v3.SchemaOrReference{
			Oneof: &v3.SchemaOrReference_Schema{
				Schema: &v3.Schema{
					Required:    []string{"message", "code", "error"},
					Type:        "object",
					Description: "Human-readable error response message in response to an invalid request.",
					Properties: &v3.Properties{
						AdditionalProperties: []*v3.NamedSchemaOrReference{
							{
								Name: "code",
								Value: &v3.SchemaOrReference{
									Oneof: &v3.SchemaOrReference_Schema{
										Schema: &v3.Schema{
											Type:        "string",
											Description: "The status code, which should be an enum value of MetaXar code errors",
										},
									},
								},
							},
							{
								Name: "message",
								Value: &v3.SchemaOrReference{
									Oneof: &v3.SchemaOrReference_Schema{
										Schema: &v3.Schema{
											Type:        "string",
											Description: "Human-readable message describing the error.",
										},
									},
								},
							},
							{
								Name: "error",
								Value: &v3.SchemaOrReference{
									Oneof: &v3.SchemaOrReference_Schema{
										Schema: &v3.Schema{
											Type:        "string",
											Description: "Detailed description of the error.",
										},
									},
								},
							},
						},
					},
					Example: &v3.Any{Yaml: "{\"message\": \"the user already has 1 player\", \"error\": \"UserAlreadyHasPlayer\", \"code\": \"InvalidArgument\"}"},
				},
			},
		},
	}
}
