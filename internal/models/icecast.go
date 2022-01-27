package models

import "github.com/Felyp-Henrique/icegen/pkg/templates"

type IcecastModel struct {
	numclients string
	numsources string
	queue      string
	clitimeout string
	hdrtimeout string
	srctimeout string
	burst      string
	srcpass    string
	admin      string
	adminpass  string
	host       string
	port       string
}

func NewIcecastModel() *IcecastModel {
	return &IcecastModel{}
}

func (i *IcecastModel) GetNumClients() string {
	return i.numclients
}

func (i *IcecastModel) SetNumClients(numClients string) {
	i.numclients = numClients
}

func (i *IcecastModel) GetNumSources() string {
	return i.numsources
}

func (i *IcecastModel) SetNumSources(numSources string) {
	i.numsources = numSources
}

func (i *IcecastModel) GetQueue() string {
	return i.queue
}

func (i *IcecastModel) SetQueue(queue string) {
	i.queue = queue
}

func (i *IcecastModel) GetCliTimeout() string {
	return i.clitimeout
}

func (i *IcecastModel) SetCliTimeout(cliTimeout string) {
	i.clitimeout = cliTimeout
}

func (i *IcecastModel) GetHdrTimeout() string {
	return i.hdrtimeout
}

func (i *IcecastModel) SetHdrTimeout(hdrTimeout string) {
	i.hdrtimeout = hdrTimeout
}

func (i *IcecastModel) GetSrcTimeout() string {
	return i.srctimeout
}

func (i *IcecastModel) SetSrcTimeout(srcTimeout string) {
	i.srctimeout = srcTimeout
}

func (i *IcecastModel) GetBurst() string {
	return i.burst
}

func (i *IcecastModel) SetBurst(burst string) {
	i.burst = burst
}

func (i *IcecastModel) GetSrcPass() string {
	return i.srcpass
}

func (i *IcecastModel) SetSrcPass(srcPass string) {
	i.srcpass = srcPass
}

func (i *IcecastModel) GetAdmin() string {
	return i.admin
}

func (i *IcecastModel) SetAdmin(admin string) {
	i.admin = admin
}

func (i *IcecastModel) GetAdminPass() string {
	return i.adminpass
}

func (i *IcecastModel) SetAdminPass(adminPass string) {
	i.adminpass = adminPass
}

func (i *IcecastModel) GetHost() string {
	return i.host
}

func (i *IcecastModel) SetHost(host string) {
	i.host = host
}

func (i *IcecastModel) GetPort() string {
	return i.port
}

func (i *IcecastModel) SetPort(port string) {
	i.port = port
}

func (i *IcecastModel) ToContext() templates.Context {
	return templates.Context{
		"numclients": i.numclients,
		"numsources": i.numsources,
		"queue":      i.queue,
		"clitimeout": i.clitimeout,
		"hdrtimeout": i.hdrtimeout,
		"srctimeout": i.srctimeout,
		"burst":      i.burst,
		"srcpass":    i.srcpass,
		"admin":      i.admin,
		"adminpass":  i.adminpass,
		"host":       i.host,
		"port":       i.port,
	}
}
