package mongodb

import (
	"astro_pro/api/models"
	mg "astro_pro/data/mongodb"
	"time"
)

func cityToDomain(c *mg.City) *models.City {
	return &models.City{Name: c.Name}
}

func segmentToDomain(s *mg.Segment) *models.Segment {
	return &models.Segment{Name: s.Name, City: s.City}
}

func controllerToDomain(c *mg.Controller) *models.Controller {
	return &models.Controller{
		Mac:       c.Mac,
		Segment:   c.Segment,
		City:      c.City,
		Power:     c.Power,
		Serial:    c.Serial,
		Imei:      c.Imei,
		Imsi:      c.Imsi,
		Ccid:      c.Ccid,
		Phone:     c.Phone,
		Connected: c.Connected,
		Master:    c.Master,
		Tlevel:    c.Tlevel,
		Rng:       c.Rng,
		Smac:      c.Smac,
		Dir:       c.Dir,
		Nid:       c.Nid,
		Group:     c.Group,
		Level:     c.Level,
		Mrssi:     c.Mrssi,
		Rssi:      c.Rssi,
		Devt:      c.Devt,
		Devm:      c.Devm,
	}
}

func complexProfileToDomain(cp *mg.ComplexProfile) *models.ComplexProfile {
	if cp == nil {
		return nil
	}

	return &models.ComplexProfile{
		Pwm0:  cp.Pwm0,
		Time1: cp.Time1,
		Pwm1:  cp.Pwm1,
		Time2: cp.Time2,
		Pwm2:  cp.Pwm2,
		Time3: cp.Time3,
		Pwm3:  cp.Pwm3,
		Time4: cp.Time4,
		Pwm4:  cp.Pwm4,
	}
}

func simpleProfileToDomain(sp *mg.SimpleProfile) *models.SimpleProfile {
	if sp == nil {
		return nil
	}

	return &models.SimpleProfile{
		D1: sp.D1,
		P1: sp.P1,
		D2: sp.D2,
		P2: sp.P2,
	}
}

func lampToDomain(l *mg.Lamp) *models.Lamp {

	cp := []*models.ComplexProfile{nil, nil, nil}

	if l.ComplexProfiles != nil {
		for i, p := range l.ComplexProfiles {
			cp[i] = complexProfileToDomain(p)
		}
	}

	sp := []*models.SimpleProfile{nil, nil, nil}

	if l.SimpleProfiles != nil {
		for i, p := range l.SimpleProfiles {
			sp[i] = simpleProfileToDomain(p)
		}
	}

	lDomain := &models.Lamp{
		Mac:             l.Mac,
		Source:          l.Source,
		Segment:         l.Segment,
		City:            l.City,
		Dir:             l.Dir,
		Level:           l.Level,
		Nid:             l.Nid,
		Group:           l.Group,
		Smac:            l.Smac,
		Rssi:            l.Rssi,
		Devt:            l.Devt,
		Devm:            l.Devm,
		Eblk:            l.Eblk,
		Cycles:          l.Cycles,
		Runh:            l.Runh,
		Nvsc:            l.Nvsc,
		Lpwm:            l.Lpwm,
		Cpwm:            l.Cpwm,
		Mrssi:           l.Mrssi,
		Rfch:            l.Rfch,
		Rfpwr:           l.Rfpwr,
		Pwm:             l.Pwm,
		Pwmct:           l.Pwmct,
		Pow:             l.Pow,
		Lux:             l.Lux,
		Temp:            l.Temp,
		Energy:          l.Energy,
		Rng:             l.Rng,
		Tlevel:          l.Tlevel,
		Date:            l.Date,
		Val:             l.Val,
		Rise:            l.Rise,
		Set:             l.Set,
		ID:              l.ProfileId,
		Scdtm:           l.Scdtm,
		Rfps:            l.Rfps,
		Twil:            l.Twil,
		Received:        l.Received.Time().Unix(),
		SimpleProfiles:  sp,
		ComplexProfiles: cp,
	}

	if l.Location != nil {
		lat := l.Location.Coordinates[0]
		lng := l.Location.Coordinates[1]
		lDomain.Lat = &lat
		lDomain.Lng = &lng
	}

	isLampBroken := lampIsBroken(l)
	if isLampBroken {
		lDomain.Status = models.LampStatusBroken
	} else {
		if l.Pwm == nil {
			lDomain.Status = models.LampStatusUnknown
		} else {
			if *l.Pwm > 0 {
				lDomain.Status = models.LampStatusOn
			} else {
				lDomain.Status = models.LampStatusOff
			}
		}
	}

	return lDomain
}

func lampIsBroken(l *mg.Lamp) bool {
	lampTime := l.Received.Time()
	currTime := time.Now()
	status := false
	if currTime.Year() > lampTime.Year() {
		status = true
	} else {
		if currTime.Month() > lampTime.Month() {
			status = true
		} else {
			diffDay := currTime.Day() - lampTime.Day()
			if diffDay > 4 {
				status = true
			}
		}
	}
	return status
}
