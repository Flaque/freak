package cmds

import (
	"io/ioutil"
	"strings"
	"testing"

	filet "github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
)

func TestRecord(t *testing.T) {
	defer filet.CleanUp(t)
	path := filet.TmpFile(t, "", "time, intensity\n").Name()

	// Record something like "<date>, 2"
	Record(2, path)

	dat, err := ioutil.ReadFile(path)
	assert.NoError(t, err) // sanity

	lines := strings.Split(string(dat), "\n")

	// We should get 3 lines
	assert.Equal(t, len(lines), 3) // 3 including extra new line

	// Ends with 2 as the value
	lastCharacter := string(lines[1][len(lines[1])-1])
	assert.Equal(t, "2", lastCharacter, lines[1])
}
