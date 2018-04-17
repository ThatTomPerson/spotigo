clean:
	rm -rf dist
.PHONY: clean

release: clean
	goreleaser
.PHONY: clean