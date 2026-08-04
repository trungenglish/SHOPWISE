package contract_test

import (
	"testing"

	"shopwise/retail/internal/platform/testutil/contract"
)

func TestOpenAPISpecHasIdentityPaths(t *testing.T) {
	spec := contract.LoadSpec(t)

	contract.AssertPathMethod(t, spec, "/identity/register", "POST")
	contract.AssertResponseCode(t, spec, "/identity/register", "POST", 201)
	contract.AssertResponseCode(t, spec, "/identity/register", "POST", 409)

	contract.AssertPathMethod(t, spec, "/identity/login", "POST")
	contract.AssertResponseCode(t, spec, "/identity/login", "POST", 200)
	contract.AssertResponseCode(t, spec, "/identity/login", "POST", 401)

	contract.AssertPathMethod(t, spec, "/identity/refresh", "POST")
	contract.AssertPathMethod(t, spec, "/identity/google/start", "GET")
	contract.AssertPathMethod(t, spec, "/identity/google/callback", "GET")
	contract.AssertPathMethod(t, spec, "/identity/verify-email", "POST")

}

func TestSwaggerDocumentsAllOpenAPIPaths(t *testing.T) {
	openapi := contract.LoadSpec(t)
	swagger := contract.LoadSwaggerSpec(t)
	contract.AssertSwaggerDocumentsOpenAPIPaths(t, openapi, swagger)
}
