package utils

var FilesNameToCheckAgainst = map[string]bool{
	"exe": true,
	"dll": true,
	"sh":  true,
	"bin": true,
	"app": true,
	"ps1": true,
}

var SkipDirs = map[string]bool{
	".git":                      true,
	"node_modules":              true,
	"vendor":                    true,
	".cache":                    true,
	"/proc":                     true,
	"/sys":                      true,
	"/dev":                      true,
	"/run":                      true,
	".local":                    true,
	".npm":                      true,
	".cargo":                    true,
	"$Recycle.Bin":              true,
	"System Volume Information": true,
	"Library":                   true,
}
