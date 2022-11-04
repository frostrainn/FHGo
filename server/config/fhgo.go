package config

type FHGo struct {
	Port string
	Pprof
}

type Pprof struct {
	Enabled bool
	Port    string
}
