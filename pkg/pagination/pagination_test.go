package pagination

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	PageNum      int32
	PageSize     int32
	OrderBy      string
	IsDescending bool
	KeyWords     string
	SearchKey    string
	CustomField  string
}

var (
	pbData = ListRequest{
		PageNum:      10,
		PageSize:     50,
		OrderBy:      "test",
		IsDescending: true,
		KeyWords:     "key",
		SearchKey:    "search",
		PacketId:     0,
	}
	customData = testRequest{
		PageNum:      10,
		PageSize:     50,
		OrderBy:      "test",
		IsDescending: true,
		KeyWords:     "key",
		SearchKey:    "search",
		CustomField:  "my data",
	}
	targetPage = Page{
		Num:          10,
		Size:         50,
		OrderBy:      "test",
		IsDescending: true,
		KeyWords:     "key",
		SearchKey:    "search",
	}
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		excepted Page
	}{
		{
			name:     "test pb struct",
			data:     pbData,
			excepted: targetPage,
		},
		{
			name:     "test custom data",
			data:     customData,
			excepted: targetPage,
		},
		{
			name:     "test pb ptr",
			data:     &pbData,
			excepted: targetPage,
		},
		{
			name:     "test custom ptr",
			data:     &customData,
			excepted: targetPage,
		},
	}
	for _, test := range tests {
		page, err := Parse(test.data)
		assert.NoError(t, err)
		assert.Equal(t, targetPage, page)
		fmt.Printf("%+v\n", page)
	}
}

func TestInvalidDataTypeParse(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		excepted error
	}{
		{
			name:     "map type",
			data:     make(map[string]interface{}),
			excepted: ErrInvalidParseData,
		},
		{
			name:     "string",
			data:     "string",
			excepted: ErrInvalidParseData,
		},
		{
			name:     "int",
			data:     int(32),
			excepted: ErrInvalidParseData,
		},
		{
			name:     "float",
			data:     float32(32.0),
			excepted: ErrInvalidParseData,
		},
		{
			name:     "bool",
			data:     true,
			excepted: ErrInvalidParseData,
		},
		{
			name:     "struct",
			data:     struct{ invalidField string }{"test"},
			excepted: ErrNoPageInfo,
		},
		{
			name:     "mismatch struct type",
			data:     struct{ PageNum string }{"PageNum"},
			excepted: ErrInvalidPageNum,
		},
	}
	for _, test := range tests {
		_, err := Parse(test.data)
		assert.Equal(t, test.excepted, err)
	}
}
