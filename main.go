package main

import (
	"fmt"
	"golang-question/config"
	"golang-question/errorx"
)

const (
	ErrCodeSecretTooShort = 1001
)

type Secret string

func (s Secret) Validate() errorx.Error {
	if len(s) < 8 {
		return errorx.Cf(ErrCodeSecretTooShort, "invalid secret %s", s)
	}
	return nil
}

type Config struct {
	Secret Secret `yaml:"secret" json:"secret"`
}

var conf = config.Local[Config]().Watch().InitData(Config{
	Secret: "hello",
})

func main() {
	s := conf.Get().Secret
	if err := s.Validate(); err != nil {
		fmt.Printf("validate error: %+v\n", err)
	}
	if err := conf.Update(Config{Secret: Secret("updated secret")}); err != nil {
		fmt.Printf("update error: %+v\n", err)
	}
}
