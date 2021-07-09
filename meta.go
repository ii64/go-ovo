package ovo

var (
	MetaAppID      = ""
	MetaAppVersion = ""
	MetaOS         = ""
)

type Device struct {
	OS string
}

func NewDefaultDevice() *Device {
	return &Device{
		OS: MetaOS,
	}
}

type App struct {
	ID      string
	Version string

	*Device
}

func NewDefaultApp() *App {
	return &App{
		ID:      MetaAppID,
		Version: MetaAppVersion,
		Device:  NewDefaultDevice(),
	}
}
