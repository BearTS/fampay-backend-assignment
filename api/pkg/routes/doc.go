package routes

//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config=oapi-codegen.yaml ./openapi-spec.yaml

import (
	// ensure that go:generate has oapi-codegen available for Rest API generation
	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
)
