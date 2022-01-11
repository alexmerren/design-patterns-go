## I need to make a general build function because all the directories are the same layout

.PHONY: help
help:
	@echo "Print a help message"

.PHONY: adapter
adapter:
	@go run adapter/*.go

.PHONY: singleton
singleton:
	@go run singleton/*.go

.PHONY: observer
observer:
	@go run observer/*.go

.PHONY: factoryMethod
factoryMethod:
	@go run factoryMethod/*.go

.PHONY: decorator
decorator:
	@go run decorator/*.go

.PHONY: facade
facade:
	@go run facade/*.go

.PHONY: command
command:
	@go run command/*.go


