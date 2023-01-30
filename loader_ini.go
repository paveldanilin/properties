package properties

import (
	"strings"

	"gopkg.in/ini.v1"
)

type IniOptions struct {
	Filename          string
	ConcatSectionName bool
}

func (opts *IniOptions) Options() map[string]interface{} {
	return map[string]interface{}{
		"filename":          opts.Filename,
		"concatSectionName": opts.ConcatSectionName,
	}
}

// LoadIni loads properties from INI file
func LoadIni(options IniOptions) (*Properties, error) {
	cfg, err := ini.Load(options.Filename)
	if err != nil {
		return nil, err
	}

	props := New()

	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			k := key.Name()
			if len(k) > 0 {
				if options.ConcatSectionName {
					k = section.Name() + "." + k
				}
				props.SetProperty(strings.ToLower(k), key.Value())
			}
		}
	}

	return props, nil
}
