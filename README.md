# Overview

A validation library for GoLang structs and primitives. You define your validations as variable structures, then validate against them.

Uses interfaces for the emitters.

# Generating everything

```shell script
go run ./cmd/gen/generate.go -- --templatePath=templates --outputRootPath=.
```

# Todo

 * Floats!
 * Replace the pathing with a more agnostic pathing/pointer system. This is ROUGHLY using JSON pointers.
