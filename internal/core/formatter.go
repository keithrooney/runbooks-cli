package core

type Formatter interface {
	Format(options FormatOptions) string
}
