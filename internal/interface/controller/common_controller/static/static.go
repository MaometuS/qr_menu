package static

import (
	"embed"
)

//go:embed templates/*
var Templates embed.FS

//go:embed assets/*
var Assets embed.FS
