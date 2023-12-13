package profile

type Profile struct {
	Version string
}

var CurrentProfile = Profile{
	Version: "0.0.9",
}
