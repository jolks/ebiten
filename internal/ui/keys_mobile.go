// Copyright 2013 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by genkeys.go using 'go generate'. DO NOT EDIT.

//go:build (android || ios) && !nintendosdk && !ebitencbackend
// +build android ios
// +build !nintendosdk
// +build !ebitencbackend

package ui

import (
	"golang.org/x/mobile/event/key"
)

var gbuildKeyToUIKey = map[key.Code]Key{
	key.CodeA:                  KeyA,
	key.CodeB:                  KeyB,
	key.CodeC:                  KeyC,
	key.CodeD:                  KeyD,
	key.CodeE:                  KeyE,
	key.CodeF:                  KeyF,
	key.CodeG:                  KeyG,
	key.CodeH:                  KeyH,
	key.CodeI:                  KeyI,
	key.CodeJ:                  KeyJ,
	key.CodeK:                  KeyK,
	key.CodeL:                  KeyL,
	key.CodeM:                  KeyM,
	key.CodeN:                  KeyN,
	key.CodeO:                  KeyO,
	key.CodeP:                  KeyP,
	key.CodeQ:                  KeyQ,
	key.CodeR:                  KeyR,
	key.CodeS:                  KeyS,
	key.CodeT:                  KeyT,
	key.CodeU:                  KeyU,
	key.CodeV:                  KeyV,
	key.CodeW:                  KeyW,
	key.CodeX:                  KeyX,
	key.CodeY:                  KeyY,
	key.CodeZ:                  KeyZ,
	key.Code1:                  KeyDigit1,
	key.Code2:                  KeyDigit2,
	key.Code3:                  KeyDigit3,
	key.Code4:                  KeyDigit4,
	key.Code5:                  KeyDigit5,
	key.Code6:                  KeyDigit6,
	key.Code7:                  KeyDigit7,
	key.Code8:                  KeyDigit8,
	key.Code9:                  KeyDigit9,
	key.Code0:                  KeyDigit0,
	key.CodeReturnEnter:        KeyEnter,
	key.CodeEscape:             KeyEscape,
	key.CodeDeleteBackspace:    KeyBackspace,
	key.CodeTab:                KeyTab,
	key.CodeSpacebar:           KeySpace,
	key.CodeHyphenMinus:        KeyMinus,
	key.CodeEqualSign:          KeyEqual,
	key.CodeLeftSquareBracket:  KeyBracketLeft,
	key.CodeRightSquareBracket: KeyBracketRight,
	key.CodeBackslash:          KeyBackslash,
	key.CodeSemicolon:          KeySemicolon,
	key.CodeApostrophe:         KeyQuote,
	key.CodeGraveAccent:        KeyBackquote,
	key.CodeComma:              KeyComma,
	key.CodeFullStop:           KeyPeriod,
	key.CodeSlash:              KeySlash,
	key.CodeCapsLock:           KeyCapsLock,
	key.CodeF1:                 KeyF1,
	key.CodeF2:                 KeyF2,
	key.CodeF3:                 KeyF3,
	key.CodeF4:                 KeyF4,
	key.CodeF5:                 KeyF5,
	key.CodeF6:                 KeyF6,
	key.CodeF7:                 KeyF7,
	key.CodeF8:                 KeyF8,
	key.CodeF9:                 KeyF9,
	key.CodeF10:                KeyF10,
	key.CodeF11:                KeyF11,
	key.CodeF12:                KeyF12,
	key.CodePause:              KeyPause,
	key.CodeInsert:             KeyInsert,
	key.CodeHome:               KeyHome,
	key.CodePageUp:             KeyPageUp,
	key.CodeDeleteForward:      KeyDelete,
	key.CodeEnd:                KeyEnd,
	key.CodePageDown:           KeyPageDown,
	key.CodeRightArrow:         KeyArrowRight,
	key.CodeLeftArrow:          KeyArrowLeft,
	key.CodeDownArrow:          KeyArrowDown,
	key.CodeUpArrow:            KeyArrowUp,
	key.CodeKeypadNumLock:      KeyNumLock,
	key.CodeKeypadSlash:        KeyNumpadDivide,
	key.CodeKeypadAsterisk:     KeyNumpadMultiply,
	key.CodeKeypadHyphenMinus:  KeyNumpadSubtract,
	key.CodeKeypadPlusSign:     KeyNumpadAdd,
	key.CodeKeypadEnter:        KeyNumpadEnter,
	key.CodeKeypad1:            KeyNumpad1,
	key.CodeKeypad2:            KeyNumpad2,
	key.CodeKeypad3:            KeyNumpad3,
	key.CodeKeypad4:            KeyNumpad4,
	key.CodeKeypad5:            KeyNumpad5,
	key.CodeKeypad6:            KeyNumpad6,
	key.CodeKeypad7:            KeyNumpad7,
	key.CodeKeypad8:            KeyNumpad8,
	key.CodeKeypad9:            KeyNumpad9,
	key.CodeKeypad0:            KeyNumpad0,
	key.CodeKeypadFullStop:     KeyNumpadDecimal,
	key.CodeKeypadEqualSign:    KeyNumpadEqual,
	key.CodeLeftControl:        KeyControlLeft,
	key.CodeLeftShift:          KeyShiftLeft,
	key.CodeLeftAlt:            KeyAltLeft,
	key.CodeLeftGUI:            KeyMetaLeft,
	key.CodeRightControl:       KeyControlRight,
	key.CodeRightShift:         KeyShiftRight,
	key.CodeRightAlt:           KeyAltRight,
	key.CodeRightGUI:           KeyMetaRight,
}
