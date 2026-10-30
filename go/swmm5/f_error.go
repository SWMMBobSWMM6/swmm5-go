// Generated from EPA SWMM 5.2.4 (error.c) by tools/c2x/emit_go.py. Do not edit by hand;
// fix the translator, go/crt or overrides.go instead.

package main

import (
	R "swmm5go/crt"
)

var ErrString = R.Buf(256)

// error.c:30
func error_getMsg(errCode int32, msg R.CPtr) R.CPtr {
	switch errCode {
	case 101:
		R.Strcpy(msg, L46)
		break
	case 103:
		R.Strcpy(msg, L47)
		break
	case 105:
		R.Strcpy(msg, L48)
		break
	case 107:
		R.Strcpy(msg, L49)
		break
	case 108:
		R.Strcpy(msg, L50)
		break
	case 109:
		R.Strcpy(msg, L51)
		break
	case 110:
		R.Strcpy(msg, L52)
		break
	case 111:
		R.Strcpy(msg, L53)
		break
	case 112:
		R.Strcpy(msg, L54)
		break
	case 113:
		R.Strcpy(msg, L55)
		break
	case 114:
		R.Strcpy(msg, L56)
		break
	case 115:
		R.Strcpy(msg, L57)
		break
	case 117:
		R.Strcpy(msg, L58)
		break
	case 119:
		R.Strcpy(msg, L59)
		break
	case 121:
		R.Strcpy(msg, L60)
		break
	case 122:
		R.Strcpy(msg, L61)
		break
	case 131:
		R.Strcpy(msg, L62)
		break
	case 133:
		R.Strcpy(msg, L63)
		break
	case 134:
		R.Strcpy(msg, L64)
		break
	case 135:
		R.Strcpy(msg, L65)
		break
	case 136:
		R.Strcpy(msg, L66)
		break
	case 137:
		R.Strcpy(msg, L67)
		break
	case 138:
		R.Strcpy(msg, L68)
		break
	case 139:
		R.Strcpy(msg, L69)
		break
	case 140:
		R.Strcpy(msg, L70)
		break
	case 141:
		R.Strcpy(msg, L71)
		break
	case 143:
		R.Strcpy(msg, L72)
		break
	case 145:
		R.Strcpy(msg, L73)
		break
	case 151:
		R.Strcpy(msg, L74)
		break
	case 153:
		R.Strcpy(msg, L75)
		break
	case 155:
		R.Strcpy(msg, L76)
		break
	case 156:
		R.Strcpy(msg, L77)
		break
	case 157:
		R.Strcpy(msg, L78)
		break
	case 158:
		R.Strcpy(msg, L79)
		break
	case 159:
		R.Strcpy(msg, L80)
		break
	case 161:
		R.Strcpy(msg, L81)
		break
	case 171:
		R.Strcpy(msg, L82)
		break
	case 173:
		R.Strcpy(msg, L83)
		break
	case 181:
		R.Strcpy(msg, L84)
		break
	case 182:
		R.Strcpy(msg, L85)
		break
	case 183:
		R.Strcpy(msg, L86)
		break
	case 184:
		R.Strcpy(msg, L87)
		break
	case 185:
		R.Strcpy(msg, L88)
		break
	case 187:
		R.Strcpy(msg, L89)
		break
	case 188:
		R.Strcpy(msg, L90)
		break
	case 191:
		R.Strcpy(msg, L91)
		break
	case 193:
		R.Strcpy(msg, L92)
		break
	case 195:
		R.Strcpy(msg, L93)
		break
	case 200:
		R.Strcpy(msg, L94)
		break
	case 201:
		R.Strcpy(msg, L95)
		break
	case 203:
		R.Strcpy(msg, L96)
		break
	case 205:
		R.Strcpy(msg, L97)
		break
	case 207:
		R.Strcpy(msg, L98)
		break
	case 209:
		R.Strcpy(msg, L99)
		break
	case 211:
		R.Strcpy(msg, L100)
		break
	case 213:
		R.Strcpy(msg, L101)
		break
	case 217:
		R.Strcpy(msg, L102)
		break
	case 219:
		R.Strcpy(msg, L103)
		break
	case 221:
		R.Strcpy(msg, L104)
		break
	case 223:
		R.Strcpy(msg, L105)
		break
	case 225:
		R.Strcpy(msg, L106)
		break
	case 227:
		R.Strcpy(msg, L107)
		break
	case 229:
		R.Strcpy(msg, L108)
		break
	case 231:
		R.Strcpy(msg, L109)
		break
	case 233:
		R.Strcpy(msg, L110)
		break
	case 235:
		R.Strcpy(msg, L111)
		break
	case 301:
		R.Strcpy(msg, L112)
		break
	case 303:
		R.Strcpy(msg, L113)
		break
	case 305:
		R.Strcpy(msg, L114)
		break
	case 307:
		R.Strcpy(msg, L115)
		break
	case 308:
		R.Strcpy(msg, L116)
		break
	case 309:
		R.Strcpy(msg, L117)
		break
	case 311:
		R.Strcpy(msg, L118)
		break
	case 313:
		R.Strcpy(msg, L119)
		break
	case 315:
		R.Strcpy(msg, L120)
		break
	case 317:
		R.Strcpy(msg, L121)
		break
	case 318:
		R.Strcpy(msg, L122)
		break
	case 319:
		R.Strcpy(msg, L123)
		break
	case 320:
		R.Strcpy(msg, L124)
		break
	case 321:
		R.Strcpy(msg, L125)
		break
	case 323:
		R.Strcpy(msg, L126)
		break
	case 325:
		R.Strcpy(msg, L127)
		break
	case 327:
		R.Strcpy(msg, L128)
		break
	case 329:
		R.Strcpy(msg, L129)
		break
	case 331:
		R.Strcpy(msg, L130)
		break
	case 333:
		R.Strcpy(msg, L131)
		break
	case 335:
		R.Strcpy(msg, L132)
		break
	case 336:
		R.Strcpy(msg, L133)
		break
	case 337:
		R.Strcpy(msg, L134)
		break
	case 338:
		R.Strcpy(msg, L135)
		break
	case 339:
		R.Strcpy(msg, L136)
		break
	case 341:
		R.Strcpy(msg, L137)
		break
	case 343:
		R.Strcpy(msg, L138)
		break
	case 345:
		R.Strcpy(msg, L139)
		break
	case 351:
		R.Strcpy(msg, L140)
		break
	case 353:
		R.Strcpy(msg, L141)
		break
	case 355:
		R.Strcpy(msg, L142)
		break
	case 357:
		R.Strcpy(msg, L143)
		break
	case 361:
		R.Strcpy(msg, L144)
		break
	case 363:
		R.Strcpy(msg, L145)
		break
	case 500:
		R.Strcpy(msg, L146)
		break
	case 501:
		R.Strcpy(msg, L147)
		break
	case 502:
		R.Strcpy(msg, L148)
		break
	case 503:
		R.Strcpy(msg, L149)
		break
	case 504:
		R.Strcpy(msg, L150)
		break
	case 505:
		R.Strcpy(msg, L151)
		break
	case 506:
		R.Strcpy(msg, L152)
		break
	case 507:
		R.Strcpy(msg, L153)
		break
	case 508:
		R.Strcpy(msg, L154)
		break
	case 509:
		R.Strcpy(msg, L155)
		break
	default:
		R.Strcpy(msg, L0)
	}
	return (msg)
}

// error.c:45
func error_setInpError(errcode int32, s R.CPtr) int32 {
	R.Strcpy(ErrString, s)
	return errcode
}
