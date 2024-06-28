package version

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/pascaliske/magicmirror/logger"
)

// build information
var Name string = "MagicMirror"
var Version string
var GitCommit string
var BuildTime string

func GetBanner() string {
	return figure.NewFigure(Name, "graffiti", true).String()
}

func GetVersion() string {
	return fmt.Sprintf("Version %s @ %s (%s)", color.CyanString(Version), color.CyanString(GitCommit), color.CyanString(BuildTime))
}

func PrintBanner() {
	logger.Raw(GetBanner())
}

func PrintVersion() {
	logger.Raw(GetVersion())
}

func PrintBannerWithVersion() {
	logger.Raw(GetBanner())
	logger.Raw(GetVersion())
}
