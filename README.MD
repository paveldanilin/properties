# Properties util

```go
props := properties.New("name", "Batman", "age", "35")

name := props.MustString("name") // Returns string
age  := props.MustInt("age")     // Returns int
```


```go
// See ./test/test1.ini

props := properties.LoadIni(IniOptions{
	Filename:           "./test/test1.ini",
    ConcatSectionName:  true,
})

// Get property or default value
protocol := props.String("server.protocol", "http")
```