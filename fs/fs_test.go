package fs

import (
	"io/ioutil"
	"testing"

	filet "github.com/Flaque/filet"
	"github.com/stretchr/testify/assert"
)

func TestWritingToCSVCreatesFile(t *testing.T) {
	defer filet.CleanUp(t)

	dir := filet.TmpDir(t, "")

	err := WriteToCSV([]string{"a", "b", "c"}, dir+"/example.csv")
	assert.NoError(t, err)

	// Creates the file
	assert.True(t, filet.Exists(t, dir+"/example.csv"))

	// Populates the file
	dat, err := ioutil.ReadFile(dir + "/example.csv")
	assert.NoError(t, err) // sanity check
	assert.Equal(t, "a,b,c\n", string(dat))
}
