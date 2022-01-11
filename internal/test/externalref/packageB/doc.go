package packageB

//go:generate go run github.com/tamj0rd2/oapi-codegen/cmd/oapi-codegen -generate types,skip-prune,spec --package=packageB -o externalref.gen.go spec.yaml
