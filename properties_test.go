package properties

import (
	"strings"
	"testing"
)

func TestMustString(t *testing.T) {
	props := New()
	props.SetProperty("day", "Monday")

	name := props.MustString("day")
	if name != "Monday" {
		t.Errorf("expected value `%s`, but got `%s`", "Monday", name)
	}
}

func TestString(t *testing.T) {
	props := New("name", "Pavel")

	name := props.String("name", "Batman")

	if name != "Pavel" {
		t.Errorf("expected value `%s`, but got `%s`", "Pavel", name)
	}

	user := props.String("user", "Batman")

	if user != "Batman" {
		t.Errorf("expected value `%s`, but got `%s`", "Batman", user)
	}
}

func TestMustBool(t *testing.T) {
	props := New()
	props.SetProperty("flag", "true")

	flag := props.MustBool("flag")
	if !flag {
		t.Errorf("expected value `%v`, but got `%v`", true, flag)
	}
}

func TestMustInt(t *testing.T) {
	props := New()
	props.SetProperty("age", "35")

	age := props.MustInt("age")
	if age != 35 {
		t.Errorf("expected value %d, but got %d", 35, age)
	}
}

func TestMustFloat(t *testing.T) {
	props := New("posx", "4.3", "posy", "0.21")

	posx := props.MustFloat("posx")
	posy := props.MustFloat("posy")

	if posx != 4.3 {
		t.Errorf("expected value %f, but got %f", 4.3, posx)
	}

	if posy != 0.21 {
		t.Errorf("expected value %f, but got %f", 0.21, posy)
	}
}

func TestLoaderIni(t *testing.T) {
	props, err := LoadIni(IniOptions{
		Filename: "./test/test1.ini",
	})
	if err != nil {
		t.Error(err)
	}

	protocol := props.MustString("protocol")

	if protocol != "http" {
		t.Errorf("expected protocol value %s, but got %s", "http", protocol)
	}
}

func TestLoaderIniConcatSection(t *testing.T) {
	props, err := LoadIni(IniOptions{
		Filename:          "./test/test1.ini",
		ConcatSectionName: true,
	})
	if err != nil {
		t.Error(err)
	}

	protocol := props.MustString("server.protocol")

	if protocol != "http" {
		t.Errorf("expected protocol value %s, but got %s", "http", protocol)
	}
}

func TestKeysWithPrefix(t *testing.T) {
	props := New("aa.a", "1", "bb.b", "2", "aa.c", "3", "bb.d", "4")

	keys := props.KeysWithPrefix("aa.")

	if len(keys) != 2 {
		t.Errorf("expected two keys with a prefix `aa.`, but got %d", len(keys))
	}
}

func TestNewFromMap(t *testing.T) {
	props := NewFromMap(map[string]string{
		"message": "Hello!",
	})

	if props.MustString("message") != "Hello!" {
		t.Errorf("expected value `%s`, but got `%s`", "Hello!", props.MustString("message"))
	}
}

func TestRemoveProperty(t *testing.T) {

	props := New()

	props.SetProperty("a", "11")
	props.SetProperty("b", "22")
	props.SetProperty("c", "33")

	props.RemoveProperty("b")
	props.RemoveProperty("unknown")

	if props.Size() != 2 {
		t.Errorf("expected size is 2, but got %d", props.Size())
	}
}

func TestInt(t *testing.T) {
	props := NewFromMap(map[string]string{
		"a": "2",
		"b": "3",
	})

	var a, b int

	if v, err := props.Int("a", 0); err == nil {
		a = v
	} else {
		t.Error(err)
	}

	if v, err := props.Int("b", 0); err == nil {
		b = v
	} else {
		t.Error(err)
	}

	if a+b != 5 {
		t.Errorf("expected value is 5, but got %d", a+b)
	}
}

func TestGetWithPrefix(t *testing.T) {
	props := NewFromMap(map[string]string{
		"some.property": "abcd",
		"log.level":     "trace",
		"log.format":    "txt",
		"log.layout":    "{{time}} {{msg}}\n",
		"some.value":    "331",
	})

	props = props.GetWithPrefix("log.")

	if props.Size() != 3 {
		t.Errorf("expected size is 3, but got %d", props.Size())
	}

	requiredKeys := []string{"log.level", "log.layout", "log.format"}

	if !props.Contains(requiredKeys) {
		t.Errorf("expected value %v, but got %v", requiredKeys, props.Keys())
	}
}

func TestContainsAny(t *testing.T) {
	props := New()

	props.SetProperty("prop.one", "one")
	props.SetProperty("prop.two", "two")
	props.SetProperty("prop.three", "three")

	if !props.ContainsAny([]string{"a", "prop.two", "b"}) {
		t.Errorf("not found expected key")
	}
}

func TestMerge(t *testing.T) {
	props1 := New("a", "1")
	props2 := New("b", "2")
	props3 := props1.Merge(props2, true)

	if props3.Size() != 2 {
		t.Errorf("expected size is 2, but got %d", props3.Size())
	}

	if !props3.Contains([]string{"a", "b"}) {
		t.Errorf("must contains [a,b], but contains %v", props3.Keys())
	}
}

func TestMergeKeepSameKeyValue(t *testing.T) {
	props := New("a", "1").Merge(New("a", "2"), false)

	a := props.String("a", "")
	if a != "1" {
		t.Errorf("expected value `1`, but got `%s`", a)
	}
}

func TestRenameKeys(t *testing.T) {
	props := NewFromMap(map[string]string{
		"prefix_1.a": "1",
		"prefix_1.b": "2",
		"prefix_2.a": "3",
	})

	prefix := "prefix_1."

	rename := func(key string) string {
		return strings.TrimPrefix(key, prefix)
	}

	propsNew := props.RenameKeys(rename)

	if !propsNew.Contains([]string{"a", "b", "prefix_2.a"}) {
		t.Errorf("expected keys [], but got %v", propsNew.Keys())
	}
}
