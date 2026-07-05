package contract

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const (
	defaultSpecRelPath    = "specs/001-ai-shopping-copilot/contracts/openapi.yaml"
	defaultSwaggerRelPath = "apps/server/docs/swagger.yaml"
)

type Spec struct {
	Paths map[string]map[string]operation `yaml:"paths"`
}

type operation struct {
	Responses map[string]response `yaml:"responses"`
}

type response struct {
	Description string `yaml:"description"`
}

func LoadSpec(t *testing.T) *Spec {
	t.Helper()

	path, err := resolveSpecPath()
	if err != nil {
		t.Fatalf("resolve spec path: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read spec %s: %v", path, err)
	}

	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		t.Fatalf("parse spec: %v", err)
	}
	if len(spec.Paths) == 0 {
		t.Fatal("spec has no paths")
	}
	return &spec
}

func AssertPathMethod(t *testing.T, spec *Spec, path, method string) {
	t.Helper()
	method = strings.ToLower(method)
	ops, ok := spec.Paths[path]
	if !ok {
		t.Fatalf("path %q not found in OpenAPI spec", path)
	}
	if _, ok := ops[method]; !ok {
		t.Fatalf("method %s %q not found in OpenAPI spec", strings.ToUpper(method), path)
	}
}

func AssertResponseCode(t *testing.T, spec *Spec, path, method string, status int) {
	t.Helper()
	method = strings.ToLower(method)
	ops, ok := spec.Paths[path]
	if !ok {
		t.Fatalf("path %q not found in OpenAPI spec", path)
	}
	op, ok := ops[method]
	if !ok {
		t.Fatalf("method %s %q not found in OpenAPI spec", strings.ToUpper(method), path)
	}
	code := fmt.Sprintf("%d", status)
	if _, ok := op.Responses[code]; !ok {
		t.Fatalf("response %s for %s %q not found in OpenAPI spec", code, strings.ToUpper(method), path)
	}
}

func LoadSwaggerSpec(t *testing.T) *Spec {
	t.Helper()

	path, err := resolveRepoRelativePath(defaultSwaggerRelPath)
	if err != nil {
		t.Fatalf("resolve swagger path: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read swagger %s: %v", path, err)
	}

	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		t.Fatalf("parse swagger: %v", err)
	}
	if len(spec.Paths) == 0 {
		t.Fatal("swagger has no paths")
	}
	return &spec
}

// AssertSwaggerDocumentsOpenAPIPaths ensures generated swagger is a superset of the
// Phase 1 product contract in specs/001-ai-shopping-copilot/contracts/openapi.yaml.
func AssertSwaggerDocumentsOpenAPIPaths(t *testing.T, openapi, swagger *Spec) {
	t.Helper()

	for path, methods := range openapi.Paths {
		swaggerMethods, ok := findSwaggerMethods(swagger, path)
		if !ok {
			t.Fatalf("swagger missing OpenAPI path %q", path)
		}
		for method := range methods {
			if _, ok := swaggerMethods[method]; !ok {
				t.Fatalf("swagger missing %s %q documented in OpenAPI", strings.ToUpper(method), path)
			}
		}
	}
}

func findSwaggerMethods(swagger *Spec, openapiPath string) (map[string]operation, bool) {
	if methods, ok := swagger.Paths[openapiPath]; ok {
		return methods, true
	}

	for path, methods := range swagger.Paths {
		if pathsEquivalent(openapiPath, path) {
			return methods, true
		}
	}
	return nil, false
}

func pathsEquivalent(openapiPath, swaggerPath string) bool {
	openapiParts := strings.Split(openapiPath, "/")
	swaggerParts := strings.Split(swaggerPath, "/")
	if len(openapiParts) != len(swaggerParts) {
		return false
	}

	for i := range openapiParts {
		openPart := openapiParts[i]
		swaggerPart := swaggerParts[i]
		if strings.HasPrefix(openPart, "{") && strings.HasSuffix(openPart, "}") {
			if !strings.HasPrefix(swaggerPart, "{") || !strings.HasSuffix(swaggerPart, "}") {
				return false
			}
			continue
		}
		if openPart != swaggerPart {
			return false
		}
	}
	return true
}

func resolveSpecPath() (string, error) {
	return resolveRepoRelativePath(defaultSpecRelPath)
}

func resolveRepoRelativePath(relPath string) (string, error) {
	if root := os.Getenv("TEST_REPO_ROOT"); root != "" {
		return filepath.Join(root, relPath), nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := wd
	for {
		candidate := filepath.Join(dir, relPath)
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find %s from %s", relPath, wd)
}
