package app

import "os"

func (a *App) Build() {
	os.Setenv("APP_ID", a.ID)
}
