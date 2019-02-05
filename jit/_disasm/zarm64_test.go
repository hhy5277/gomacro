/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * zarm_64_test.go
 *
 *  Created on Jan 23, 2019
 *      Author Massimiliano Ghilardi
 */

package disasm

import (
	"testing"

	. "github.com/cosmos72/gomacro/jit/arm64"
)

func TestArm64Sample(t *testing.T) {
	var asm Asm

	for id := RLo; id+2 <= RHi; id++ {
		asm.Init()
		if asm.RegIsUsed(id) || asm.RegIsUsed(id+1) || asm.RegIsUsed(id+2) {
			continue
		}
		r := MakeReg(id+0, Int64)
		s := MakeReg(id+1, Int64)
		t := MakeReg(id+2, Int64)
		c := ConstInt64(0xFFF000)
		asm.Asm(MOV, c, r, //
			ADD3, r, s, t, //
			SUB3, r, s, t, //
			AND3, r, s, t, //
			OR3, r, s, t, //
			XOR3, r, s, t, //
			SHL3, r, s, t, //
			SHR3, r, s, t, //
			ADD3, r, c, t, //
			SUB3, r, c, t, //
			AND3, r, c, t, //
			OR3, r, c, t, //
			XOR3, r, c, t, //
		)

		PrintDisasm(ARM64, asm.Code())
	}
}