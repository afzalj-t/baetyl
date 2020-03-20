package sync

import (
	"encoding/json"
	v1 "github.com/baetyl/baetyl-go/spec/v1"

	"github.com/baetyl/baetyl-core/common"
)

type ApplicationResource struct {
	Type    string         `yaml:"type" json:"type"`
	Name    string         `yaml:"name" json:"name"`
	Version string         `yaml:"version" json:"version"`
	Value   v1.Application `yaml:"value" json:"value"`
}

type ConfigurationResource struct {
	Type    string           `yaml:"type" json:"type"`
	Name    string           `yaml:"name" json:"name"`
	Version string           `yaml:"version" json:"version"`
	Value   v1.Configuration `yaml:"value" json:"value"`
}

type DesireRequest struct {
	Resources []*BaseResource `yaml:"resources" json:"resources"`
}

type DesireResponse struct {
	Resources []*Resource `yaml:"resources" json:"resources"`
}

type BaseResource struct {
	Type    common.Resource `yaml:"type,omitempty" json:"type,omitempty"`
	Name    string          `yaml:"name,omitempty" json:"name,omitempty"`
	Version string          `yaml:"version,omitempty" json:"version,omitempty"`
}

type Resource struct {
	BaseResource `yaml:",inline" json:",inline"`
	Data         []byte      `yaml:"data,omitempty" json:"data,omitempty"`
	Value        interface{} `yaml:"value,omitempty" json:"value,omitempty"`
}

type SecretResource struct {
	Type    string    `yaml:"type" json:"type"`
	Name    string    `yaml:"name" json:"name"`
	Version string    `yaml:"version" json:"version"`
	Value   v1.Secret `yaml:"value" json:"value"`
}

func (r *Resource) GetApplication() *v1.Application {
	if r.Type == common.Application {
		return r.Value.(*v1.Application)
	}
	return nil
}

func (r *Resource) GetConfiguration() *v1.Configuration {
	if r.Type == common.Configuration {
		return r.Value.(*v1.Configuration)
	}
	return nil
}

func (r *Resource) GetSecret() *v1.Secret {
	if r.Type == common.Secret {
		return r.Value.(*v1.Secret)
	}
	return nil
}

func (r *Resource) UnmarshalJSON(b []byte) error {
	var base BaseResource
	err := json.Unmarshal(b, &base)
	if err != nil {
		return err
	}
	switch base.Type {
	case common.Application:
		var app ApplicationResource
		err := json.Unmarshal(b, &app)
		if err != nil {
			return err
		}
		r.Value = &app.Value
	case common.Configuration:
		var config ConfigurationResource
		err := json.Unmarshal(b, &config)
		if err != nil {
			return err
		}
		r.Value = &config.Value
	case common.Secret:
		var secret SecretResource
		err := json.Unmarshal(b, &secret)
		if err != nil {
			return err
		}
		r.Value = &secret.Value
	}
	r.Data = b
	r.BaseResource = base
	return nil
}

type StorageObject struct {
	MD5         string `json:"md5,omitempty" yaml:"md5"`
	URL         string `json:"url,omitempty" yaml:"url"`
	Compression string `json:"compression,omitempty" yaml:"compression"`
}
