# Generate client from OpenAPI specification
generate_client:
	openapi-generator generate -i api/openapi.yaml -g go --git-repo-id speakeasy/client --git-user-id dblooman -o client
