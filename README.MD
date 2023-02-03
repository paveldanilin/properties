# Properties util

```go
props := properties.New("name", "Batman", "age", "35")

// Common
props.SetProperty("key", "value")
props.GetProperty("key")               // (string, error)
props.HasProperty("key")               // bool
props.RemoveProperty("keys")           // bool
props.Keys()                           // []string
props.KeysWithPrefix("prefix")         // []string
props.GetWithPrefix("prefix")          // *Properties
props.All()                            // map[string]string
props.Contains([]string{"p1", "p2"})   // bool
props.ContainsAny([]string{"p1", "p2"})// bool

// Retusn value or panic
v := props.MustString("name")  // Returns string
v := props.MustInt("age")      // Returns int
v := props.MustBool("flag")    // Returns bool
v := props.MustFloat("x")      // Returns float
v := props.MustUint("x")       // Returns uint

// Returns value or error (property not found or parsing error)
v, err := props.GetInt("v")
v, err := props.GetBool("v")
v, err := props.GetFloat("v")
v, err := props.GetUint("v")

// Returns value/default value or error (parsing error)
v, err := props.Int("v", 123)
v, err := props.Float("v", 0.2)
v, err := props.Bool("v", false)
v, err := props.Uint("v", 2)
```


```go
// See ./test/test1.ini
// ...
// [server]
// # Protocol (http or https)
// protocol = http
// ...

props := properties.LoadIni(IniOptions{
    Filename:           "./test/test1.ini",
    ConcatSectionName:  true, // Concatenates section name and property name
})

// Get property or default value
protocol := props.String("server.protocol", "http")
```