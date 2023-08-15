package netrc

type Data struct {
	Machines []Machine
}

type Machine struct {
	Name     string
	Login    string
	Password string
}
