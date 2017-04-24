package services

func Discovery() Service {
	return Service{key: "discovery"}
}

func Imperium() Service {
	return Service{key: "imperium"}
}

func Socket() Service {
	return Service{key: "socket"}
}

func Project() Service {
	return Service{key: "project"}
}

func Registry() Service {
	return Service{key: "registry"}
}

func Apps() Service {
	return Service{key: "apps"}
}

func Vendor() Service {
	return Service{key: "vendor"}
}

func Schema() Service {
	return Service{key: "schema"}
}

func Preference() Service {
	return Service{key: "preference"}
}

func Notify() Service {
	return Service{key: "notify"}
}

func Impart() Service {
	return Service{key: "impart"}
}

func Config() Service {
	return Service{key: "config"}
}
