.PHONY: spec.update
spec.update:
	curl -o swagger.json https://circleci.com/api/v2/openapi.json
	prettier --write swagger.json
