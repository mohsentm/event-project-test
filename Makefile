#Check which operating system we are having and use specified command to build it
OSFLAG :=
BUILDSTART :=
ifeq ($(OS),Windows_NT)
		BUILDSTART =
		CLEAN = rd /s /q bin
else
	BUILDSTART = env GOOS=linux CGO_ENABLED=0
	CLEAN = rm -rf ./bin
endif

build:
	$(BUILDSTART) go build -o ./bin/api ./cmd/api
clean:
	$(CLEAN)

.PHONY: build clean