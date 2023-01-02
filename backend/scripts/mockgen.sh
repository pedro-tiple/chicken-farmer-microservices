#!/bin/sh

mockgen -package farm -source ./internal/farm/controller.go -destination ./internal/farm/controller_mock.go
mockgen -package farm -source ./internal/farm/service.go -destination ./internal/farm/service_mock.go
