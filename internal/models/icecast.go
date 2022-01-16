package models

import "github.com/Felyp-Henrique/icegen/pkg/templates"

type IcecastModel struct {
	templates.Context
}

func NewIcecastModel() *IcecastModel {
	return &IcecastModel{}
}

func (i *IcecastModel) GetNumClients() string {
	return i.Context["numclients"].(string)
}

func (i *IcecastModel) SetNumClients(numClients string) {
	i.Context["numclients"] = numClients
}

func (i *IcecastModel) GetNumSources() string {
	return i.Context["numsources"].(string)
}

func (i *IcecastModel) SetNumSources(numSources string) {
	i.Context["numsources"] = numSources
}

func (i *IcecastModel) GetQueue() string {
	return i.Context["queue"].(string)
}

func (i *IcecastModel) SetQueue(queue string) {
	i.Context["queue"] = queue
}

func (i *IcecastModel) GetCliTimeout() string {
	return i.Context["clitimeout"].(string)
}

func (i *IcecastModel) SetCliTimeout(cliTimeout string) {
	i.Context["clitimeout"] = cliTimeout
}

func (i *IcecastModel) GetHdrTimeout() string {
	return i.Context["hdrtimeout"].(string)
}

func (i *IcecastModel) SetHdrTimeout(hdrTimeout string) {
	i.Context["hdrtimeout"] = hdrTimeout
}

func (i *IcecastModel) GetSrcTimeout() string {
	return i.Context["srctimeout"].(string)
}

func (i *IcecastModel) SetSrcTimeout(srcTimeout string) {
	i.Context["srctimeout"] = srcTimeout
}

func (i *IcecastModel) GetBurst() string {
	return i.Context["burst"].(string)
}

func (i *IcecastModel) SetBurst(burst string) {
	i.Context["burst"] = burst
}

func (i *IcecastModel) GetSrcPass() string {
	return i.Context["srcpass"].(string)
}

func (i *IcecastModel) SetSrcPass(srcPass string) {
	i.Context["srcpass"] = srcPass
}

func (i *IcecastModel) GetAdmin() string {
	return i.Context["admin"].(string)
}

func (i *IcecastModel) SetAdmin(admin string) {
	i.Context["admin"] = admin
}

func (i *IcecastModel) GetAdminPass() string {
	return i.Context["adminpass"].(string)
}

func (i *IcecastModel) SetAdminPass(adminPass string) {
	i.Context["adminpass"] = adminPass
}

func (i *IcecastModel) GetHost() string {
	return i.Context["host"].(string)
}

func (i *IcecastModel) SetHost(host string) {
	i.Context["host"] = host
}

func (i *IcecastModel) GetPort() string {
	return i.Context["port"].(string)
}

func (i *IcecastModel) SetPort(port string) {
	i.Context["port"] = port
}

func (i *IcecastModel) GetRelayHost() string {
	return i.Context["relayhost"].(string)
}

func (i *IcecastModel) SetRelayHost(relayHost string) {
	i.Context["relayhost"] = relayHost
}

func (i *IcecastModel) GetRelayUpdateInterval() string {
	return i.Context["relayupdateinterval"].(string)
}

func (i *IcecastModel) SetRelayUpdateInterval(relayUpdateInterval string) {
	i.Context["relayupdateinterval"] = relayUpdateInterval
}

func (i *IcecastModel) GetRelayUser() string {
	return i.Context["relayuser"].(string)
}

func (i *IcecastModel) SetRelayUser(relayUser string) {
	i.Context["relayuser"] = relayUser
}

func (i *IcecastModel) GetRelayPassword() string {
	return i.Context["relaypassword"].(string)
}

func (i *IcecastModel) SetRelayPassword(relayPassword string) {
	i.Context["relaypassword"] = relayPassword
}

func (i *IcecastModel) GetRelayDemand() string {
	return i.Context["relaydemand"].(string)
}

func (i *IcecastModel) SetRelayDemand(relayDemand string) {
	i.Context["relaydemand"] = relayDemand
}

func (i *IcecastModel) GetRelayPort() string {
	return i.Context["relayport"].(string)
}

func (i *IcecastModel) SetRelayPort(relayPort string) {
	i.Context["relayport"] = relayPort
}

func (i *IcecastModel) GetRelayMount() string {
	return i.Context["relaymount"].(string)
}

func (i *IcecastModel) SetRelayMount(relayMount string) {
	i.Context["relaymount"] = relayMount
}

func (i *IcecastModel) GetRelayMountLocal() string {
	return i.Context["relaymountlocal"].(string)
}

func (i *IcecastModel) SetRelayMountLocal(relayMountLocal string) {
	i.Context["relaymountlocal"] = relayMountLocal
}

func (i *IcecastModel) GetRelayOnDemand() string {
	return i.Context["relayondemand"].(string)
}

func (i *IcecastModel) SetRelayOnDemand(relayOnDemand string) {
	i.Context["relayondemand"] = relayOnDemand
}

func (i *IcecastModel) GetRelayShoutcast() string {
	return i.Context["relayshoutcast"].(string)
}

func (i *IcecastModel) SetRelayShoutcast(relayShoutcast string) {
	i.Context["relayshoutcast"] = relayShoutcast
}
