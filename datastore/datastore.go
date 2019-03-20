package datastore

import (
	"github.com/loyalsys/couchbase"
	"github.com/loyalsys/error"
	"github.com/loyalsys/lsenv"
)

type rtDatastore map[lscb.Region]voucher

type Datastore struct {
	Voucher rtDatastore
}

func (d *rtDatastore) add(region lscb.Region, v voucher) {
	(*d)[region] = v
}

func (d *rtDatastore) Region(region lscb.Region) *voucher {
	if datastore, ok := (*d)[region]; ok {
		return &datastore
	}

	panic(lserr.NewErrf("datstore region %v not found.", region))
}

const (
	envVarEuropeHost     lsenv.EnvVar = "CB_EUROPE_HOST"
	envVarEuropeUserName lsenv.EnvVar = "CB_EUROPE_USER_NAME"
	envVarEuropePassword lsenv.EnvVar = "CB_EUROPE_PASSWORD"
)

var envVarToLoad = []lsenv.EnvVar{
	envVarEuropeHost,
	envVarEuropeUserName,
	envVarEuropePassword,
}

func CreateInstance() (*Datastore, error) {
	envInfo, err := lsenv.ReadEnvironmentVariables(envVarToLoad)
	if err != nil {
		return nil, lserr.WrapErrf(err, "failed to read environment variables for datastore.")
	}

	voucher := voucher{}
	err = lscb.CreateBucket(&voucher, envInfo[envVarEuropeHost], envInfo[envVarEuropeUserName], envInfo[envVarEuropePassword])
	if err != nil {
		return nil, lserr.WrapErrf(err, "failed to create main datastore bucket.")
	}

	rtds := rtDatastore{}
	rtds.add(lscb.RegionEurope, voucher)

	datastore := Datastore{Voucher: rtds}

	return &datastore, nil
}
