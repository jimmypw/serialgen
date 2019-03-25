BUILDVER := $(shell git describe --tags)

serialgen:
	go build

.PHONY: package
package: serialgen
	@echo Packaging version $(BUILDVER)
	fpm -s dir -t rpm -n serialgen -v $(BUILDVER) --prefix /usr/bin serialgen
