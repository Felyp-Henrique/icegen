package entities

type Icecast interface {
	GetNumClients() string
	SetNumClients(numClients string)

	GetNumSources() string
	SetNumSources(numSources string)

	GetQueue() string
	SetQueue(queue string)

	GetCliTimeout() string
	SetCliTimeout(cliTimeout string)

	GetHdrTimeout() string
	SetHdrTimeout(hdrTimeout string)

	GetSrcTimeout() string
	SetSrcTimeout(srcTimeout string)

	GetBurst() string
	SetBurst(burst string)

	GetSrcPass() string
	SetSrcPass(srcPass string)

	GetAdmin() string
	SetAdmin(admin string)

	GetAdminPass() string
	SetAdminPass(adminPass string)

	GetHost() string
	SetHost(host string)

	GetPort() string
	SetPort(port string)
}
