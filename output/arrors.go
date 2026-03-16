package output

import (
	"github.com/fatih/color"
)

func PrintError(value any) {
	color.Red(value)
}
