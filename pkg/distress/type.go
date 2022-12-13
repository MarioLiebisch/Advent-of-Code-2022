package distress

type Packages []*Package

type Package struct {
	value    int
	contents Packages
}

type Message struct {
	packages Packages
}
