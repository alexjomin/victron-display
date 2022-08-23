package vedirect

import (
	"errors"
)

var (
	cariageReturn       = byte('\r')
	linefeed            = byte('\n')
	tab                 = byte('\t')
	hexmarker           = byte(':')
	ErrCheckSumNotValid = errors.New("checksum is not valid")
)

const (
	stateHex = iota
	stateWaitHeader
	stateKey
	stateValue
	stateChecksum
)

type KeyValue map[string]string

type Parser struct {
	currentState int
	checksum     int
	key          string
	value        string
	KV           KeyValue
	Ready        bool
}

func NewParser() (Parser, error) {
	return Parser{
		currentState: stateWaitHeader,
		KV:           KeyValue{},
		Ready:        false,
	}, nil
}

func (p Parser) GetKV() (data KeyValue, err error) {
	if p.Ready {
		data = p.KV
		p.KV = KeyValue{}
		return data, nil
	}
	return data, errors.New("parser is not ready")

}

func (p Parser) ParseByte(b byte) (Parser, error) {
	if b == hexmarker && p.currentState != stateChecksum {
		p.currentState = stateHex
	}

	switch p.currentState {

	case stateWaitHeader:
		p.checksum += int(b)
		switch b {
		case cariageReturn:
			p.currentState = stateWaitHeader
		case linefeed:
			p.currentState = stateKey
		}
		return p, nil

	case stateKey:
		p.checksum += int(b)
		if b == tab {
			if p.key == Checksum {
				p.currentState = stateChecksum
			} else {
				p.currentState = stateValue
			}
			return p, nil
		}
		p.key += string(b)
		return p, nil

	case stateValue:
		p.checksum += int(b)
		if b == cariageReturn {
			p.currentState = stateWaitHeader
			err := IsKnownLabel(p.key)
			if err != nil {
				p.key = ""
				p.value = ""
				return p, nil
			}
			p.KV[p.key] = p.value
			p.key = ""
			p.value = ""
			return p, nil
		}

		p.value += string(b)
		return p, nil

	case stateChecksum:
		p.checksum += int(b)
		p.key = ""
		p.value = ""
		p.currentState = stateWaitHeader
		if p.checksum%256 != 0 {
			p.checksum = 0
			return p, ErrCheckSumNotValid
		}
		p.Ready = true
		p.checksum = 0

	case stateHex:
		p.checksum = 0
		if b == linefeed {
			p.currentState = stateWaitHeader
		}
	}
	return p, nil
}
