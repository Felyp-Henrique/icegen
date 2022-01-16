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

	GetRelayHost() string
	SetRelayHost(relayHost string)

	GetRelayUpdateInterval() string
	SetRelayUpdateInterval(relayUpdateInterval string)

	GetRelayUser() string
	SetRelayUser(relayUser string)

	GetRelayPassword() string
	SetRelayPassword(relayPassword string)

	GetRelayDemand() string
	SetRelayDemand(relayDemand string)

	GetRelayPort() string
	SetRelayPort(relayPort string)

	GetRelayMount() string
	SetRelayMount(relayMount string)

	GetRelayMountLocal() string
	SetRelayMountLocal(relayMountLocal string)

	GetRelayOnDemand() string
	SetRelayOnDemand(relayOnDemand string)

	GetRelayShoutcast() string
	SetRelayShoutcast(relayShoutcast string)
}
