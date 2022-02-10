# Usage example

```go
import "github.com/tkeel.io/kit/pagination"


func (s *SubscribeService) ListSubscribeEntities(ctx context.Context, req *pb.ListSubscribeEntitiesRequest) (*pb.ListSubscribeEntitiesResponse, error) {
    page, err := pagination.Parse(req)

    // {Num:10 Size:50 OrderBy:test IsDescending:true KeyWords:key SearchKey:search}
    fmt.Println(page)
	
	if page.Requried {
		// Do paginate here
        query := "LIMIT " + strconv.Itoa(int(page.Limit()))
        query += "OFFSET " + strconv.Itoa(int(page.Offset()))

	} else {
		// Query all date here
    }
	
	count := db.Find(query).Count()
	
    page.FillResponse(resp, count)
}

```

## func Required
Used to determine if the paging request passed meets the paging needs

Judgement conditions：
```go
func (p Page) Required() bool {
	return p.Num > 0 && p.Size > 0
}
```

## func Limit
if no limit set this will return the default value.
```go
func (p Page) Limit() uint32 {
	if p.Size != 0 {
		return uint32(p.Size)
	}

	return uint32(p.defaultSize)
}
```

## func Offset
count the offset of the current page
```go
func (p Page) Offset() uint32 {
	if p.Num <= 0 {
		return 0
	}
	return uint32((p.Num - 1) * p.Size)
}
```

## func SearchCondition
return a `map[string]string` for the search condition 
```go
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
```

## func FillResponse
Automatic padding of paginated data to match paginated responsive design.
```go
func (p Page) FillResponse(resp interface{}, total int) error {
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
				lastPage := total / int(p.Size)
				if total%int(p.Size) == 0 {
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
```