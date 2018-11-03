package cmds

import (
	"strconv"
	"time"

	"github.com/Flaque/freak/fs"
)

// Record records the current intensity in the file located at the
// specified path
func Record(intensity int, path string) error {
	data := []string{time.Now().Format(time.RFC1123), strconv.Itoa(intensity)}

	return fs.WriteToCSV(data, path)
}
