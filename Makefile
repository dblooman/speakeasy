# Generate client from OpenAPI specification
generate:
	openapi-generator generate -i api/openapi.yaml -g go --git-repo-id speakeasy --git-user-id dblooman
