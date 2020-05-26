package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChris_UnmarshalJSON(t *testing.T) {
	file, _ := ioutil.ReadFile("../../test/data/chris.json")

	c := Chris{}

	err := json.Unmarshal([]byte(file), &c)
	if err != nil {
		t.Error(err)
	}


	assert.EqualValues(t, "Chris", c.Name, fmt.Sprintf("%s should be equal", "name"))
	assert.EqualValues(t, 32, c.Age, fmt.Sprintf("%s should be equal", "age"))
}
