start-server:
	cd server; go run .

build-ui:
	docker run --rm -v $(shell pwd)/frontend/src:/app/frontend/src -v $(shell pwd)/server/static:/app/dist vitebuild
