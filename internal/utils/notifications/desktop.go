package notifications

import (
	"regexp"
	"runtime"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	"github.com/gen2brain/beeep"
	"github.com/spf13/viper"
)

const notificationImage = "assets/warning.png"

const ansi = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:[a-zA-Z\\d]*(?:;[a-zA-Z\\d]*)*)?\u0007)|(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PRZcf-ntqry=><~]))"

var pattern = regexp.MustCompile(ansi)

// stripANSI removes ANSI escape sequences from a string. From https://github.com/acarl005/stripansi
func stripANSI(str string) string {
	return pattern.ReplaceAllString(str, "")
}

// SendNotification sends a desktop notification without sound.
func SendNotification(title string, message string) {
	if viper.GetBool("notifications") {
		if runtime.GOOS == "darwin" {
			if err := beeep.Notify(stripANSI(title), stripANSI(message), notificationImage); err != nil {
				gbl.Log.Error(err)
			}
		}
	}
}

//// SendAlert sends a desktop notification with sound.
// func SendAlert(title string, message string, force bool) {
//	if viper.GetBool("notifications") || force {
//		if runtime.GOOS == "darwin" {
//			if err := beeep.Alert(stripANSI(title), stripANSI(message), notificationImage); err != nil {
//				gbl.Log.Error(err)
//			}
//		}
//	}
//}
