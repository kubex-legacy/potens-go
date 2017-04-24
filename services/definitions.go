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
