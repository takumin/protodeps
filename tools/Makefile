.PHONY: install
install: download tidy
	go list --tags tools -f '{{ join .Imports "\n" }}' | xargs go install

.PHONY: download
download:
	go mod download

.PHONY: tidy
tidy:
	go mod tidy
