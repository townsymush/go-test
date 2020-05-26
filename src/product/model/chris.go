package product

import (
	"encoding/json"
	"strconv"
)

type Chris struct {
	Name string
	Age int
}

type ChrisInput struct {
	Name string
	Age string
}

func (c *Chris) UnmarshalJSON(data []byte) (err error) {
	var ci ChrisInput
	if err := json.Unmarshal(data, &ci); err != nil {
		return err
	}
	c.Name = ci.Name
	c.Age, err = strconv.Atoi(ci.Age)
	if err != nil {
		return
	}
	return
}