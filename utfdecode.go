package utfdecode

import (
	"errors"
	"strconv"
	"unicode/utf16"
)

var ErrCodePositionStringToRune = errors.New("error: failed to parse string to code position")

type UtfDecode struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
}

func NewDecoder(s string) *UtfDecode {
	return &UtfDecode{
		input: []rune(s),
	}
}

func (ud *UtfDecode) Decode() (string, error) {
	result := ""
	codePoses := []rune{}
	for c := ud.next(); c != 0; c = ud.next() {
		switch c {
		case '\\':
			cNext := ud.next()
			codePos := ""
			if cNext == 'u' {
				for i := 0; i < 4; i++ {
					cc := ud.next()
					if isHexChar(cc) {
						codePos += string(cc)
					} else {
						return "", ErrCodePositionStringToRune
					}
				}
				codePosRune, err := strconv.ParseInt(codePos, 16, 32)
				if err != nil {
					return "", ErrCodePositionStringToRune
				}
				codePoses = append(codePoses, rune(codePosRune))
				if utf16.IsSurrogate(rune(codePosRune)) {
					if len(codePoses) == 2 {
						surrogate := utf16.DecodeRune(codePoses[0], codePoses[1])
						result += string(surrogate)
						codePoses = []rune{}
					}
				} else {
					result += string(codePoses[0])
					codePoses = []rune{}
				}
			}
		default:
			result += string(c)
		}
	}
	return result, nil
}

func (ud *UtfDecode) next() rune {
	if ud.readPosition >= len(ud.input) {
		ud.ch = 0
	} else {
		ud.ch = ud.input[ud.readPosition]
	}
	ud.position = ud.readPosition
	ud.readPosition += 1
	return ud.ch
}

func isHexChar(v rune) bool {
	if ('0' <= v && v <= '9') || ('a' <= v && v <= 'f') || ('A' <= v && v <= 'F') {
		return true
	}
	return false
}
