package client

type CommonOsuClient interface {
	Initialize()
	Login(username string, password string)
	WaitForLoginSuccess()
	WaitForWelcomeToBancho()

	RunClient()
}
