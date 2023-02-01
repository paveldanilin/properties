package properties

import (
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
