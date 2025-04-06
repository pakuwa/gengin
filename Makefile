deploy:
	@LATEST_TAG=$$(git describe --tags `git rev-list --tags --max-count=1`); \
	if [ -z "$$LATEST_TAG" ]; then \
		NEW_TAG="v0.0.1"; \
	else \
		MAJOR=$$(echo $$LATEST_TAG | sed 's/^v//' | cut -d. -f1); \
		MINOR=$$(echo $$LATEST_TAG | sed 's/^v//' | cut -d. -f2); \
		PATCH=$$(echo $$LATEST_TAG | sed 's/^v//' | cut -d. -f3); \
		NEW_PATCH=$$((PATCH + 1)); \
		NEW_TAG="v$$MAJOR.$$MINOR.$$NEW_PATCH"; \
	fi; \
	git tag $$NEW_TAG; \
	git push origin $$NEW_TAG
