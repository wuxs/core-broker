package pagination

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrInvalidPageNum      = errors.New("invalid page number")
	ErrInvalidPageSize     = errors.New("invalid page size")
	ErrInvalidOrderBy      = errors.New("invalid order")
	ErrInvalidKeyWords     = errors.New("invalid key words")
	ErrInvalidSearchKey    = errors.New("invalid search key")
	ErrInvalidIsDescending = errors.New("invalid is descending")
	ErrInvalidParseData    = errors.New("invalid data type parsing")
	ErrNoPageInfo          = errors.New("no page info")
	ErrInvalidResponse     = errors.New("invalid response")
)

type Page struct {
	Num              int32
	Size             int32
	OrderBy          string
	IsDescending     bool
	KeyWords         string
	SearchKey        string
	defaultSize      int32
	defaultSeparator string
}

func (p Page) Offset() uint32 {
	if p.Num <= 0 {
		return 0
	}
	return uint32((p.Num - 1) * p.Size)
}

func (p Page) Limit() uint32 {
	if p.Size != 0 {
		return uint32(p.Size)
	}

	return uint32(p.defaultSize)
}

func (p Page) SearchCondition() map[string]string {
	if p.KeyWords == "" {
		return nil
	}

	values := strings.Split(p.KeyWords, p.defaultSeparator)
	keys := strings.Split(p.SearchKey, p.defaultSeparator)

	cond := make(map[string]string, len(keys))

	for i := range keys {
		cond[keys[i]] = values[i]
	}

	return cond
}

func (p Page) Required() bool {
	return p.Num > 0 && p.Size > 0
}

func (p Page) FillResponse(resp interface{}, total int64) error {
	t := reflect.TypeOf(resp)
	v := reflect.ValueOf(resp)
	for t.Kind() != reflect.Struct {
		switch t.Kind() {
		case reflect.Ptr:
			v = v.Elem()
			t = t.Elem()
		default:
			return ErrInvalidResponse
		}
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			switch t.Field(i).Name {
			case "Total":
				v.Field(i).SetUint(uint64(total))
			case "PageNum":
				v.Field(i).SetUint(uint64(p.Num))
			case "LastPage":
				if p.Size == 0 {
					v.Field(i).SetUint(uint64(0))
					continue
				}
				lastPage := total / int64(p.Size)
				if total%int64(p.Size) == 0 {
					v.Field(i).SetUint(uint64(lastPage))
					continue
				}
				v.Field(i).SetUint(uint64(lastPage + 1))

			case "PageSize":
				v.Field(i).SetUint(uint64(p.Size))
			}
		}
	}
	return nil
}

type Option func(*Page) error

// Parse a struct which have defined Page fields.
func Parse(req interface{}, options ...Option) (Page, error) {
	q := Page{
		Num:              0,
		Size:             0,
		OrderBy:          "",
		IsDescending:     false,
		KeyWords:         "",
		SearchKey:        "",
		defaultSize:      15,
		defaultSeparator: ",",
	}
	v := reflect.ValueOf(req)
	t := reflect.TypeOf(req)
	for t.Kind() != reflect.Struct {
		switch t.Kind() {
		case reflect.Ptr:
			v = v.Elem()
			t = t.Elem()
		default:
			return q, ErrInvalidParseData
		}
	}

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() {
			switch t.Field(i).Name {
			case "PageNum":
				if val, ok := v.Field(i).Interface().(int32); ok {
					q.Num = val
				} else {
					return q, ErrInvalidPageNum
				}
			case "PageSize":
				if val, ok := v.Field(i).Interface().(int32); ok {
					q.Size = val
				} else {
					return q, ErrInvalidPageSize
				}
			case "OrderBy":
				if val, ok := v.Field(i).Interface().(string); ok {
					q.OrderBy = val
				} else {
					return q, ErrInvalidOrderBy
				}
			case "IsDescending":
				if val, ok := v.Field(i).Interface().(bool); ok {
					q.IsDescending = val
				} else {
					return q, ErrInvalidIsDescending
				}
			case "KeyWords":
				if val, ok := v.Field(i).Interface().(string); ok {
					q.KeyWords = val
				} else {
					return q, ErrInvalidKeyWords
				}
			case "SearchKey":
				if val, ok := v.Field(i).Interface().(string); ok {
					q.SearchKey = val
				} else {
					return q, ErrInvalidSearchKey
				}
			}
		}
	}

	if reflect.DeepEqual(q, Page{}) {
		return q, ErrNoPageInfo
	}

	for i := range options {
		if err := options[i](&q); err != nil {
			return q, err
		}
	}

	return q, nil
}

func SetDefaultSize(size int32) Option {
	return func(p *Page) error {
		p.defaultSize = size
		return nil
	}
}

func SetDefaultSeparator(separator string) Option {
	return func(p *Page) error {
		p.defaultSeparator = separator
		return nil
	}
}
