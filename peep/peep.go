/*
 * (C) 2017-2018 Kenneth Ceyer <kennethan@nhpcw.com>
 * this is distributed under
 * Apache 2.0 <https://www.apache.org/licenses/LICENSE-2.0>
 */

package peep

import (
	log "github.com/sirupsen/logrus"
)

type Peep struct {
	Level log.Level
}

func New(level log.Level) *Peep {
	var peep = new(Peep)
	peep.Level = level
	log.SetLevel(level)
	return peep
}

func (p *Peep) Debug(args ...interface{}) {
	log.Debug(args...)
}

func (p *Peep) Info(args ...interface{}) {
	log.Info(args...)
}

func (p *Peep) Warn(args ...interface{}) {
	log.Warn(args...)
}

func (p *Peep) Error(args ...interface{}) {
	log.Error(args...)
}

func (p *Peep) Fatal(args ...interface{}) {
	log.Fatal(args...)
}