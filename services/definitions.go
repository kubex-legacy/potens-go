package services

func Discovery() Service {
	return Service{key: "discovery"}
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

func Groups() Service {
	return Service{key: "groups"}
}

func Apps() Service {
	return Service{key: "apps"}
}

func Vendors() Service {
	return Service{key: "vendors"}
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

func DataPipe() Service {
	return Service{key: "datapipe"}
}

func Config() Service {
	return Service{key: "config"}
}

func Credentials() Service {
	return Service{key: "credentials"}
}

func EventPipe() Service {
	return Service{key: "eventpipe"}
}

func Portcullis() Service {
	return Service{key: "portcullis"}
}
