package properties

type LoaderOptions interface {
	Options() map[string]interface{}
}

type Load func(options LoaderOptions) (*Properties, error)
