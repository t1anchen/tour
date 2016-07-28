fmt:
	for i in `find . -type f -regex '.*\.go'`; do go fmt $$i; done

run:
	go run exercise-loops-and-functions.go  # flowcontrol/8
	go run exercise-slices.go								# moretypes/18
	go run exercise-maps.go									# moretypes/23
	go run exercise-fibonacci-closure.go    # moretypes/26
	go run exercise-stringer.go 						# methods/18
	go run exercise-errors.go								# methods/20
	go run exercise-reader.go								# methods/22
	go run exercise-rot-reader.go						# methods/23
	go run exercise-images.go								# methods/25
	go run exercise-equivalent-binary-trees.go # concurrency/8
	go run exercise-web-crawler.go					# concurrency/10
