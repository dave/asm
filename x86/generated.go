package x86

import "github.com/dave/asm/generator/unsafe-stub"

// AAA: ASCII adjust AL after addition.
// https://golang.org/s/x86manual#page=120
func AAA() {
	unsafe.Asm("AAA", nil)
}

// AAD: ASCII adjust AX before division.
// https://golang.org/s/x86manual#page=122
func AAD() {
	unsafe.Asm("AAD", nil)
}

// AAD_imm8: Adjust AX before division to number base imm8.
// imm8: imm8
// https://golang.org/s/x86manual#page=122
func AAD_imm8(imm8 interface{}) {
	unsafe.Asm("AAD", imm8)
}

// AAM: ASCII adjust AX after multiply.
// https://golang.org/s/x86manual#page=124
func AAM() {
	unsafe.Asm("AAM", nil)
}

// AAM_imm8: Adjust AX after multiply to number base imm8.
// imm8: imm8
// https://golang.org/s/x86manual#page=124
func AAM_imm8(imm8 interface{}) {
	unsafe.Asm("AAM", imm8)
}

// AAS: ASCII adjust AL after subtraction.
// https://golang.org/s/x86manual#page=126
func AAS() {
	unsafe.Asm("AAS", nil)
}

// ADCB_imm8_AL: Add with carry imm8 to AL.
// AL: AL/AX/EAX/RAX
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCB_imm8_AL(AL, imm8 interface{}) {
	unsafe.Asm("ADCB", AL, imm8)
}

// ADCW_imm16_AX: Add with carry imm16 to AX.
// AX: AL/AX/EAX/RAX
// imm16: imm8
// https://golang.org/s/x86manual#page=128
func ADCW_imm16_AX(AX, imm16 interface{}) {
	unsafe.Asm("ADCW", AX, imm16)
}

// ADCL_imm32_EAX: Add with carry imm32 to EAX.
// EAX: AL/AX/EAX/RAX
// imm32: imm8
// https://golang.org/s/x86manual#page=128
func ADCL_imm32_EAX(EAX, imm32 interface{}) {
	unsafe.Asm("ADCL", EAX, imm32)
}

// ADCQ_imm32_RAX: Add with carry imm32 sign extended to 64- bits to RAX.
// RAX: AL/AX/EAX/RAX
// imm32: imm8
// https://golang.org/s/x86manual#page=128
func ADCQ_imm32_RAX(RAX, imm32 interface{}) {
	unsafe.Asm("ADCQ", RAX, imm32)
}

// ADCW_imm16_r_m16: Add with carry imm16 to r/m16.
// r_m16: ModRM:r/m (r, w)
// imm16: imm8
// https://golang.org/s/x86manual#page=128
func ADCW_imm16_r_m16(r_m16, imm16 interface{}) {
	unsafe.Asm("ADCW", r_m16, imm16)
}

// ADCW_imm8_r_m16: Add with CF sign-extended imm8 to r/m16.
// r_m16: ModRM:r/m (r, w)
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCW_imm8_r_m16(r_m16, imm8 interface{}) {
	unsafe.Asm("ADCW", r_m16, imm8)
}

// ADCW_r16_r_m16: Add with carry r16 to r/m16.
// r_m16: ModRM:r/m (r, w)
// r16: ModRM:reg (r)
// https://golang.org/s/x86manual#page=128
func ADCW_r16_r_m16(r_m16, r16 interface{}) {
	unsafe.Asm("ADCW", r_m16, r16)
}

// ADCL_imm32_r_m32: Add with CF imm32 to r/m32.
// r_m32: ModRM:r/m (r, w)
// imm32: imm8
// https://golang.org/s/x86manual#page=128
func ADCL_imm32_r_m32(r_m32, imm32 interface{}) {
	unsafe.Asm("ADCL", r_m32, imm32)
}

// ADCL_imm8_r_m32: Add with CF sign-extended imm8 into r/m32.
// r_m32: ModRM:r/m (r, w)
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCL_imm8_r_m32(r_m32, imm8 interface{}) {
	unsafe.Asm("ADCL", r_m32, imm8)
}

// ADCL_r32_r_m32: Add with CF r32 to r/m32.
// r_m32: ModRM:r/m (r, w)
// r32: ModRM:reg (r)
// https://golang.org/s/x86manual#page=128
func ADCL_r32_r_m32(r_m32, r32 interface{}) {
	unsafe.Asm("ADCL", r_m32, r32)
}

// ADCQ_imm32_r_m64: Add with CF imm32 sign extended to 64-bits to r/m64.
// r_m64: ModRM:r/m (r, w)
// imm32: imm8
// https://golang.org/s/x86manual#page=128
func ADCQ_imm32_r_m64(r_m64, imm32 interface{}) {
	unsafe.Asm("ADCQ", r_m64, imm32)
}

// ADCQ_imm8_r_m64: Add with CF sign-extended imm8 into r/m64.
// r_m64: ModRM:r/m (r, w)
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCQ_imm8_r_m64(r_m64, imm8 interface{}) {
	unsafe.Asm("ADCQ", r_m64, imm8)
}

// ADCQ_r64_r_m64: Add with CF r64 to r/m64.
// r_m64: ModRM:r/m (r, w)
// r64: ModRM:reg (r)
// https://golang.org/s/x86manual#page=128
func ADCQ_r64_r_m64(r_m64, r64 interface{}) {
	unsafe.Asm("ADCQ", r_m64, r64)
}

// ADCB_imm8_r_m8: Add with carry imm8 to r/m8.
// r_m8: ModRM:r/m (r, w)
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCB_imm8_r_m8(r_m8, imm8 interface{}) {
	unsafe.Asm("ADCB", r_m8, imm8)
}

// ADCB_imm8_r_m8: Add with carry imm8 to r/m8.
// r_m8: ModRM:r/m (r, w)
// imm8: imm8
// https://golang.org/s/x86manual#page=128
func ADCB_imm8_r_m8(r_m8, imm8 interface{}) {
	unsafe.Asm("ADCB", r_m8, imm8)
}

// ADCB_r8_r_m8: Add with carry byte register to r/m8.
// r_m8: ModRM:r/m (r, w)
// r8: ModRM:reg (r)
// https://golang.org/s/x86manual#page=128
func ADCB_r8_r_m8(r_m8, r8 interface{}) {
	unsafe.Asm("ADCB", r_m8, r8)
}

// ADCB_r8_r_m8: Add with carry byte register to r/m64.
// r_m8: ModRM:r/m (r, w)
// r8: ModRM:reg (r)
// https://golang.org/s/x86manual#page=128
func ADCB_r8_r_m8(r_m8, r8 interface{}) {
	unsafe.Asm("ADCB", r_m8, r8)
}

// ADCW_r_m16_r16: Add with carry r/m16 to r16.
// r16: ModRM:reg (r, w)
// r_m16: ModRM:r/m (r)
// https://golang.org/s/x86manual#page=128
func ADCW_r_m16_r16(r16, r_m16 interface{}) {
	unsafe.Asm("ADCW", r16, r_m16)
}

// ADCL_r_m32_r32: Add with CF r/m32 to r32.
// r32: ModRM:reg (r, w)
// r_m32: ModRM:r/m (r)
// https://golang.org/s/x86manual#page=128
func ADCL_r_m32_r32(r32, r_m32 interface{}) {
	unsafe.Asm("ADCL", r32, r_m32)
}

// ADCQ_r_m64_r64: Add with CF r/m64 to r64.
// r64: ModRM:reg (r, w)
// r_m64: ModRM:r/m (r)
// https://golang.org/s/x86manual#page=128
func ADCQ_r_m64_r64(r64, r_m64 interface{}) {
	unsafe.Asm("ADCQ", r64, r_m64)
}

// ADCB_r_m8_r8: Add with carry r/m8 to byte register.
// r8: ModRM:reg (r, w)
// r_m8: ModRM:r/m (r)
// https://golang.org/s/x86manual#page=128
func ADCB_r_m8_r8(r8, r_m8 interface{}) {
	unsafe.Asm("ADCB", r8, r_m8)
}

// ADCB_r_m8_r8: Add with carry r/m64 to byte register.
// r8: ModRM:reg (r, w)
// r_m8: ModRM:r/m (r)
// https://golang.org/s/x86manual#page=128
func ADCB_r_m8_r8(r8, r_m8 interface{}) {
	unsafe.Asm("ADCB", r8, r_m8)
}
