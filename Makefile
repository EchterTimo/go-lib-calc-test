NORMALIZED_VERSION := v$(patsubst v%,%,$(VERSION))

all: help

.PHONY: help
help: ## Show this help message
	@ echo "Available targets:"
	@ awk ' \
		BEGIN { \
			FS = ":.*?## " \
		} \
		/^[a-zA-Z_-]+:.*?## / { \
			targets[$$1] = $$2; \
			if (length($$1) > max_len) { \
				max_len = length($$1); \
			} \
		} \
		END { \
			for (t in targets) { \
				printf "  \033[36m%-*s\033[0m %s\n", max_len, t, targets[t]; \
			} \
		}' $(MAKEFILE_LIST) | sort

# function to validate VERSION is passed; usage: $(call ensure_version_set)
ensure_version_set = $(if $(VERSION),,$(error ERROR: You must pass VERSION=x.y.z))

.PHONY: tag
tag: ## tag the given version
	@ $(call ensure_version_set)
	@ echo "Tagging version $(NORMALIZED_VERSION)"
	@ git tag $(NORMALIZED_VERSION)
	@ git push origin $(NORMALIZED_VERSION)



.PHONY: untag
untag: ## untag the given version
	@ $(call ensure_version_set)
	
	@ git push origin --delete $(NORMALIZED_VERSION)
	@ echo "Removed version $(NORMALIZED_VERSION) from remote..."

	@ git tag -d $(NORMALIZED_VERSION)
	@ echo "Removed version $(NORMALIZED_VERSION) locally..."