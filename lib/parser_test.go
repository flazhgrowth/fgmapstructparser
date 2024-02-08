package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StructTest struct {
	ID   int64  `tostruct:"id"`
	Name string `tostruct:"name"`
}

func TestParse(t *testing.T) {
	t.Run("TestParseStringMap", func(t *testing.T) {
		og := map[string]interface{}{
			"id":   1,
			"name": "Shani Indira Natio",
		}
		dest := &StructTest{}

		err := New("tostruct").Parse(og, dest)
		assert.Nil(t, err)
		assert.Equal(t, int64(1), dest.ID)
		assert.Equal(t, og["name"], dest.Name)
	})
	t.Run("TestParseSliceStringMap", func(t *testing.T) {
		og := []map[string]interface{}{
			{
				"id":   1,
				"name": "Shani Indira Natio",
			},
			{
				"id":   2,
				"name": "Shani Indira",
			},
		}
		var dest []StructTest

		err := New("tostruct").Parse(og, dest)
		assert.Nil(t, err)
		for i := 0; i < len(og); i += 1 {
			assert.Equal(t, int64(i+1), dest[i].ID)
			assert.Equal(t, og[i]["name"], dest[i].Name)
		}
	})
}
