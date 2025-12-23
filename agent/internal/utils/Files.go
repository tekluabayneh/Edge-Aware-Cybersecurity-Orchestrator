package utils

var FilesNameToCheckAgainst = map[string]bool{
	"exe":  true,
	"dll":  true,
	"bat":  true,
	"ps1":  true,
	"sh":   true,
	"app":  true,
	"bin":  true,
	"docm": true,
	"xlsm": true,
	"pptm": true,
	"zip":  true,
	"rar":  true,
	"7z":   true,
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
