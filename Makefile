fmt:
	for i in `find . -type f -regex '.*\.go'`; do go fmt $$i; done
