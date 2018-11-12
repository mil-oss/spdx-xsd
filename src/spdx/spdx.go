package main

import (
	"log"
	spdxdoc "spdx-doc"
	spdxlic "spdx-license"
	spdxsecism "spdx-sec-ism"
	spdxsec "spdx-security"
	"xsdprov"
)

var (
	// Config ...
	cfg xsdprov.Cfg
	// Config ...
	Config    = "config/spdx-xml-cfg.json"
	cfgs      []xsdprov.Cfg
	resources = map[string]string{}
	sources   = map[string]string{}
	tempdir   string
	temppath  string
	name      string
)

func main() {
	cfg = xsdprov.ReadConfig(Config)
	tempdir = cfg.Tempdir
	temppath = cfg.Temppath
	name = cfg.Project
	BuildIEP(name, Config)
	for i := range cfg.Implementations {
		log.Println(cfg.Implementations[i].Name)
		var c = xsdprov.ReadConfig(cfg.Implementations[i].Src)
		cfgs = append(cfgs, c)
		BuildIEP(cfg.Implementations[i].Name, cfg.Implementations[i].Src)
	}
	xsdprov.StartWeb(cfg, cfgs)
}

// BuildIEP ...
func BuildIEP(name string, configpath string) {
	if name == "spdx-xml" {
		xsdprov.InitXSDProv(configpath)
		xsdprov.GetSourceResources()
	}
	if name == "spdx-doc" {
		spdxdoc.BuildDocIEP(configpath)
	}
	if name == "spdx-license" {
		spdxlic.BuildLicenseIEP(configpath)
	}
	if name == "spdx-security" {
		spdxsec.BuildSecurityIEP(configpath)
	}
	if name == "spdx-sec-ism" {
		spdxsecism.SecurityIsmIEP(configpath)
	}
}
