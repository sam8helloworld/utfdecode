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

func Decode(s string) (string, error) {
	ud := UtfDecode{
		input: []rune(s),
	}

	result := ""
	codePoses := []rune{}
	for c := ud.next().ch; c != 0; c = ud.next().ch {
		switch c {
		case '\\':
			cNext := ud.next().ch
			codePos := ""
			if cNext == 'u' {
				for i := 0; i < 5; i++ {
					cc := ud.next().ch
					if isHexChar(cc) {
						codePos += string(cc)
						continue
					}
					if i == 4 {
						ud.back()
						continue
					}
					return "", ErrCodePositionStringToRune
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

func (ud *UtfDecode) next() UtfDecode {
	if ud.readPosition >= len(ud.input) {
		ud.ch = 0
	} else {
		ud.ch = ud.input[ud.readPosition]
	}
	ud.position = ud.readPosition
	ud.readPosition += 1
	return *ud
}

func (ud *UtfDecode) back() UtfDecode {
	ud.readPosition -= 1
	if ud.readPosition >= len(ud.input) {
		ud.ch = 0
	} else {
		ud.ch = ud.input[ud.readPosition]
	}
	ud.position = ud.readPosition
	return *ud
}

func isHexChar(v rune) bool {
	if ('0' <= v && v <= '9') || ('a' <= v && v <= 'f') || ('A' <= v && v <= 'F') {
		return true
	}
	return false
}
