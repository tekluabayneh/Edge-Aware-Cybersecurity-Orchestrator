package utils

// more files can be added to check here
var FilesNameToCheckAgainst = map[string]bool{
	"exe":  true,
	"dll":  true,
	"sh":   true,
	"bash": true,
	"bin":  true,
}

// here also if cretical fiels can be excluded
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
