package x86

import "github.com/dave/asm/generator/unsafe-stub"

// AAA
// ASCII adjust AL after addition.
//
// Documentation: https://golang.org/s/x86manual#page=120
func AAA() {
	unsafe.Asm("AAA", nil)
}

// AAD
// ASCII adjust AX before division.
// Adjust AX before division to number base imm8.
//
// Documentation: https://golang.org/s/x86manual#page=122
func AAD() {
	unsafe.Asm("AAD", nil)
}

// AAM
// ASCII adjust AX after multiply.
// Adjust AX after multiply to number base imm8.
//
// Documentation: https://golang.org/s/x86manual#page=124
func AAM() {
	unsafe.Asm("AAM", nil)
}

// AAS
// ASCII adjust AL after subtraction.
//
// Documentation: https://golang.org/s/x86manual#page=126
func AAS() {
	unsafe.Asm("AAS", nil)
}

// ADC_I
// Add with carry imm8 to AL.
// Add with carry imm16 to AX.
// Add with carry imm32 to EAX.
// Add with carry imm32 sign extended to 64- bits to RAX.
//
// al: AL/AX/EAX/RAX
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=128
func ADC_I(al, imm interface{}) {
	unsafe.Asm("ADC", al, imm)
}

// ADC_MI
// Add with carry imm16 to r/m16.
// Add with CF sign-extended imm8 to r/m16.
// Add with CF imm32 to r/m32.
// Add with CF sign-extended imm8 into r/m32.
// Add with CF imm32 sign extended to 64-bits to r/m64.
// Add with CF sign-extended imm8 into r/m64.
// Add with carry imm8 to r/m8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=128
func ADC_MI(rm, imm interface{}) {
	unsafe.Asm("ADC", rm, imm)
}

// ADC_MR
// Add with carry r16 to r/m16.
// Add with CF r32 to r/m32.
// Add with CF r64 to r/m64.
// Add with carry byte register to r/m8.
// Add with carry byte register to r/m64.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=128
func ADC_MR(rm, reg interface{}) {
	unsafe.Asm("ADC", rm, reg)
}

// ADC_RM
// Add with carry r/m16 to r16.
// Add with CF r/m32 to r32.
// Add with CF r/m64 to r64.
// Add with carry r/m8 to byte register.
// Add with carry r/m64 to byte register.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=128
func ADC_RM(reg, rm interface{}) {
	unsafe.Asm("ADC", reg, rm)
}

// ADCX
// Unsigned addition of r32 with CF, r/m32 to r32, writes CF.
// Unsigned addition of r64 with CF, r/m64 to r64, writes CF.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=131
func ADCX(reg, rm interface{}) {
	unsafe.Asm("ADCX", reg, rm)
}

// ADD_I
// Add imm8 to AL.
// Add imm16 to AX.
// Add imm32 to EAX.
// Add imm32 sign-extended to 64-bits to RAX.
//
// al: AL/AX/EAX/RAX
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=133
func ADD_I(al, imm interface{}) {
	unsafe.Asm("ADD", al, imm)
}

// ADD_MI
// Add imm16 to r/m16.
// Add sign-extended imm8 to r/m16.
// Add imm32 to r/m32.
// Add sign-extended imm8 to r/m32.
// Add imm32 sign-extended to 64-bits to r/m64.
// Add sign-extended imm8 to r/m64.
// Add imm8 to r/m8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=133
func ADD_MI(rm, imm interface{}) {
	unsafe.Asm("ADD", rm, imm)
}

// ADD_MR
// Add r16 to r/m16.
// Add r32 to r/m32.
// Add r64 to r/m64.
// Add r8 to r/m8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=133
func ADD_MR(rm, reg interface{}) {
	unsafe.Asm("ADD", rm, reg)
}

// ADD_RM
// Add r/m16 to r16.
// Add r/m32 to r32.
// Add r/m64 to r64.
// Add r/m8 to r8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=133
func ADD_RM(reg, rm interface{}) {
	unsafe.Asm("ADD", reg, rm)
}

// ADDPD
// Add packed double-precision floating-point values from xmm2/mem to xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=135
func ADDPD(reg, rm interface{}) {
	unsafe.Asm("ADDPD", reg, rm)
}

// ADDPS
// Add packed single-precision floating-point values from xmm2/m128 to xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=138
func ADDPS(reg, rm interface{}) {
	unsafe.Asm("ADDPS", reg, rm)
}

// ADDSD
// Add the low double-precision floating-point value from xmm2/mem to xmm1 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=141
func ADDSD(reg, rm interface{}) {
	unsafe.Asm("ADDSD", reg, rm)
}

// ADDSS
// Add the low single-precision floating-point value from xmm2/mem to xmm1 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=143
func ADDSS(reg, rm interface{}) {
	unsafe.Asm("ADDSS", reg, rm)
}

// ADDSUBPD
// Add/subtract double-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=145
func ADDSUBPD(reg, rm interface{}) {
	unsafe.Asm("ADDSUBPD", reg, rm)
}

// ADDSUBPS
// Add/subtract single-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=147
func ADDSUBPS(reg, rm interface{}) {
	unsafe.Asm("ADDSUBPS", reg, rm)
}

// ADOX
// Unsigned addition of r32 with OF, r/m32 to r32, writes OF.
// Unsigned addition of r64 with OF, r/m64 to r64, writes OF.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=150
func ADOX(reg, rm interface{}) {
	unsafe.Asm("ADOX", reg, rm)
}

// AESDEC
// Perform one round of an AES decryption flow, using the Equivalent Inverse Cipher, operating on a 128-bit data (state) from xmm1 with a 128-bit round key from xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=152
func AESDEC(reg, rm interface{}) {
	unsafe.Asm("AESDEC", reg, rm)
}

// AESDECLAST
// Perform the last round of an AES decryption flow, using the Equivalent Inverse Cipher, operating on a 128-bit data (state) from xmm1 with a 128-bit round key from xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=154
func AESDECLAST(reg, rm interface{}) {
	unsafe.Asm("AESDECLAST", reg, rm)
}

// AESENC
// Perform one round of an AES encryption flow, operating on a 128-bit data (state) from xmm1 with a 128-bit round key from xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=156
func AESENC(reg, rm interface{}) {
	unsafe.Asm("AESENC", reg, rm)
}

// AESENCLAST
// Perform the last round of an AES encryption flow, operating on a 128-bit data (state) from xmm1 with a 128-bit round key from xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=158
func AESENCLAST(reg, rm interface{}) {
	unsafe.Asm("AESENCLAST", reg, rm)
}

// AESIMC
// Perform the InvMixColumn transformation on a 128-bit round key from xmm2/m128 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=160
func AESIMC(reg, rm interface{}) {
	unsafe.Asm("AESIMC", reg, rm)
}

// AESKEYGENASSIST
// Assist in AES round key generation using an 8 bits Round Constant (RCON) specified in the immediate byte, operating on 128 bits of data specified in xmm2/m128 and stores the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=161
func AESKEYGENASSIST(reg, rm, imm interface{}) {
	unsafe.Asm("AESKEYGENASSIST", reg, rm, imm)
}

// AND_I
// AL AND imm8.
// AX AND imm16.
// EAX AND imm32.
// RAX AND imm32 sign-extended to 64-bits.
//
// al: AL/AX/EAX/RAX
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=163
func AND_I(al, imm interface{}) {
	unsafe.Asm("AND", al, imm)
}

// AND_MI
// r/m16 AND imm16.
// r/m16 AND imm8 (sign-extended).
// r/m32 AND imm32.
// r/m32 AND imm8 (sign-extended).
// r/m64 AND imm32 sign extended to 64-bits.
// r/m64 AND imm8 (sign-extended).
// r/m8 AND imm8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=163
func AND_MI(rm, imm interface{}) {
	unsafe.Asm("AND", rm, imm)
}

// AND_MR
// r/m16 AND r16.
// r/m32 AND r32.
// r/m64 AND r32.
// r/m8 AND r8.
// r/m64 AND r8 (sign-extended).
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=163
func AND_MR(rm, reg interface{}) {
	unsafe.Asm("AND", rm, reg)
}

// AND_RM
// r16 AND r/m16.
// r32 AND r/m32.
// r64 AND r/m64.
// r8 AND r/m8.
// r/m64 AND r8 (sign-extended).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=163
func AND_RM(reg, rm interface{}) {
	unsafe.Asm("AND", reg, rm)
}

// ANDN
// Bitwise AND of inverted r32b with r/m32, store result in r32a.
// Bitwise AND of inverted r64b with r/m64, store result in r64a.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=165
func ANDN(reg, vex, rm interface{}) {
	unsafe.Asm("ANDN", reg, vex, rm)
}

// ANDNPD
// Return the bitwise logical AND NOT of packed double- precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=172
func ANDNPD(reg, rm interface{}) {
	unsafe.Asm("ANDNPD", reg, rm)
}

// ANDNPS
// Return the bitwise logical AND NOT of packed single-precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=175
func ANDNPS(reg, rm interface{}) {
	unsafe.Asm("ANDNPS", reg, rm)
}

// ANDPD
// Return the bitwise logical AND of packed double- precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=166
func ANDPD(reg, rm interface{}) {
	unsafe.Asm("ANDPD", reg, rm)
}

// ANDPS
// Return the bitwise logical AND of packed single-precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=169
func ANDPS(reg, rm interface{}) {
	unsafe.Asm("ANDPS", reg, rm)
}

// ARPL
// Adjust RPL of r/m16 to not less than RPL of r16.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=178
func ARPL(rm, reg interface{}) {
	unsafe.Asm("ARPL", rm, reg)
}

// BEXTR
// Contiguous bitwise extract from r/m32 using r32b as control; store result in r32a.
// Contiguous bitwise extract from r/m64 using r64b as control; store result in r64a
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// vex: VEX.vvvv (r)
//
// Documentation: https://golang.org/s/x86manual#page=182
func BEXTR(reg, rm, vex interface{}) {
	unsafe.Asm("BEXTR", reg, rm, vex)
}

// BLENDPD
// Select packed DP-FP values from xmm1 and xmm2/m128 from mask specified in imm8 and store the values into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=180
func BLENDPD(reg, rm, imm interface{}) {
	unsafe.Asm("BLENDPD", reg, rm, imm)
}

// BLENDPS
// Select packed single precision floating-point values from xmm1 and xmm2/m128 from mask specified in imm8 and store the values into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=183
func BLENDPS(reg, rm, imm interface{}) {
	unsafe.Asm("BLENDPS", reg, rm, imm)
}

// BLENDVPD
// Select packed DP FP values from xmm1 and xmm2 from mask specified in XMM0 and store the values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// implicit: implicit XMM0
//
// Documentation: https://golang.org/s/x86manual#page=185
func BLENDVPD(reg, rm, implicit interface{}) {
	unsafe.Asm("BLENDVPD", reg, rm, implicit)
}

// BLENDVPS
// Select packed single precision floating-point values from xmm1 and xmm2/m128 from mask specified in XMM0 and store the values into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// implicit: implicit XMM0
//
// Documentation: https://golang.org/s/x86manual#page=187
func BLENDVPS(reg, rm, implicit interface{}) {
	unsafe.Asm("BLENDVPS", reg, rm, implicit)
}

// BLSI
// Extract lowest set bit from r/m32 and set that bit in r32.
// Extract lowest set bit from r/m64, and set that bit in r64.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=190
func BLSI(vex, rm interface{}) {
	unsafe.Asm("BLSI", vex, rm)
}

// BLSMSK
// Set all lower bits in r32 to “1” starting from bit 0 to lowest set bit in r/m32.
// Set all lower bits in r64 to “1” starting from bit 0 to lowest set bit in r/m64.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=191
func BLSMSK(vex, rm interface{}) {
	unsafe.Asm("BLSMSK", vex, rm)
}

// BLSR
// Reset lowest set bit of r/m32, keep all other bits of r/m32 and write result to r32.
// Reset lowest set bit of r/m64, keep all other bits of r/m64 and write result to r64.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=192
func BLSR(vex, rm interface{}) {
	unsafe.Asm("BLSR", vex, rm)
}

// BNDCL
// Generate a #BR if the address in r/m32 is lower than the lower bound in bnd.LB.
// Generate a #BR if the address in r/m64 is lower than the lower bound in bnd.LB.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=193
func BNDCL(reg, rm interface{}) {
	unsafe.Asm("BNDCL", reg, rm)
}

// BNDCN
// Generate a #BR if the address in r/m32 is higher than the upper bound in bnd.UB (bnb.UB not in 1's complement form).
// Generate a #BR if the address in r/m64 is higher than the upper bound in bnd.UB (bnb.UB not in 1's complement form).
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=195
func BNDCN(reg, rm interface{}) {
	unsafe.Asm("BNDCN", reg, rm)
}

// BNDCU
// Generate a #BR if the address in r/m32 is higher than the upper bound in bnd.UB (bnb.UB in 1's complement form).
// Generate a #BR if the address in r/m64 is higher than the upper bound in bnd.UB (bnb.UB in 1's complement form).
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=195
func BNDCU(reg, rm interface{}) {
	unsafe.Asm("BNDCU", reg, rm)
}

// BNDLDX
// Load the bounds stored in a bound table entry (BTE) into bnd with address translation using the base of mib and conditional on the index of mib matching the pointer value in the BTE.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=197
func BNDLDX(reg, rm interface{}) {
	unsafe.Asm("BNDLDX", reg, rm)
}

// BNDMK
// Make lower and upper bounds from m32 and store them in bnd.
// Make lower and upper bounds from m64 and store them in bnd.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=200
func BNDMK(reg, rm interface{}) {
	unsafe.Asm("BNDMK", reg, rm)
}

// BNDMOV_MR
// Move lower and upper bound from bnd2 to bound register bnd1/m128.
// Move lower and upper bound from bnd2 to bnd1/m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=202
func BNDMOV_MR(rm, reg interface{}) {
	unsafe.Asm("BNDMOV", rm, reg)
}

// BNDMOV_RM
// Move lower and upper bound from bnd2/m128 to bound register bnd1.
// Move lower and upper bound from bnd2/m64 to bound register bnd1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=202
func BNDMOV_RM(reg, rm interface{}) {
	unsafe.Asm("BNDMOV", reg, rm)
}

// BNDSTX
// Store the bounds in bnd and the pointer value in the index regis- ter of mib to a bound table entry (BTE) with address translation using the base of mib.
//
// rm: ModRM:r/m (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=205
func BNDSTX(rm, reg interface{}) {
	unsafe.Asm("BNDSTX", rm, reg)
}

// BOUND
// Check if r16 (array index) is within bounds specified by m16&16.
// Check if r32 (array index) is within bounds specified by m32&32.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=208
func BOUND(reg, rm interface{}) {
	unsafe.Asm("BOUND", reg, rm)
}

// BSF
// Bit scan forward on r/m16.
// Bit scan forward on r/m32.
// Bit scan forward on r/m64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=210
func BSF(reg, rm interface{}) {
	unsafe.Asm("BSF", reg, rm)
}

// BSR
// Bit scan reverse on r/m16.
// Bit scan reverse on r/m32.
// Bit scan reverse on r/m64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=212
func BSR(reg, rm interface{}) {
	unsafe.Asm("BSR", reg, rm)
}

// BSWAP
// Reverses the byte order of a 32-bit register.
// Reverses the byte order of a 64-bit register.
//
// opcode: opcode + rd (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=214
func BSWAP(opcode interface{}) {
	unsafe.Asm("BSWAP", opcode)
}

// BT_MI
// Store selected bit in CF flag.
//
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=215
func BT_MI(rm, imm interface{}) {
	unsafe.Asm("BT", rm, imm)
}

// BT_MR
// Store selected bit in CF flag.
//
// rm: ModRM:r/m (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=215
func BT_MR(rm, reg interface{}) {
	unsafe.Asm("BT", rm, reg)
}

// BTC_MI
// Store selected bit in CF flag and complement.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=217
func BTC_MI(rm, imm interface{}) {
	unsafe.Asm("BTC", rm, imm)
}

// BTC_MR
// Store selected bit in CF flag and complement.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=217
func BTC_MR(rm, reg interface{}) {
	unsafe.Asm("BTC", rm, reg)
}

// BTR_MI
// Store selected bit in CF flag and clear.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=219
func BTR_MI(rm, imm interface{}) {
	unsafe.Asm("BTR", rm, imm)
}

// BTR_MR
// Store selected bit in CF flag and clear.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=219
func BTR_MR(rm, reg interface{}) {
	unsafe.Asm("BTR", rm, reg)
}

// BTS_MI
// Store selected bit in CF flag and set.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=221
func BTS_MI(rm, imm interface{}) {
	unsafe.Asm("BTS", rm, imm)
}

// BTS_MR
// Store selected bit in CF flag and set.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=221
func BTS_MR(rm, reg interface{}) {
	unsafe.Asm("BTS", rm, reg)
}

// BZHI
// Zero bits in r/m32 starting with the position in r32b, write result to r32a.
// Zero bits in r/m64 starting with the position in r64b, write result to r64a.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// vex: VEX.vvvv (r)
//
// Documentation: https://golang.org/s/x86manual#page=223
func BZHI(reg, rm, vex interface{}) {
	unsafe.Asm("BZHI", reg, rm, vex)
}

// CALL_D
// Call far, absolute, address given in operand.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=224
func CALL_D(offset interface{}) {
	unsafe.Asm("CALL", offset)
}

// CALL_M
// Call near, absolute indirect, address given in r/m16.
// Call near, absolute indirect, address given in r/m32.
// Call near, absolute indirect, address given in r/m64.
// Call near, relative, displacement relative to next instruction.
// Call near, relative, displacement relative to next instruction. 32-bit displacement sign extended to 64-bits in 64-bit mode.
// Call far, absolute indirect address given in m16:16. In 32-bit mode: if selector points to a gate, then RIP = 32-bit zero extended displacement taken from gate; else RIP = zero extended 16-bit offset from far pointer referenced in the instruction.
// In 64-bit mode: If selector points to a gate, then RIP = 64-bit displacement taken from gate; else RIP = zero extended 32-bit offset from far pointer referenced in the instruction.
// In 64-bit mode: If selector points to a gate, then RIP = 64-bit displacement taken from gate; else RIP = 64-bit offset from far pointer referenced in the instruction.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=224
func CALL_M(rm interface{}) {
	unsafe.Asm("CALL", rm)
}

// CBW
// AX ← sign-extend of AL.
//
// Documentation: https://golang.org/s/x86manual#page=237
func CBW() {
	unsafe.Asm("CBW", nil)
}

// CDQ
// EDX:EAX ← sign-extend of EAX.
//
// Documentation: https://golang.org/s/x86manual#page=380
func CDQ() {
	unsafe.Asm("CDQ", nil)
}

// CDQE
// RAX ← sign-extend of EAX.
//
// Documentation: https://golang.org/s/x86manual#page=237
func CDQE() {
	unsafe.Asm("CDQE", nil)
}

// CLAC
// Clear the AC flag in the EFLAGS register.
//
// Documentation: https://golang.org/s/x86manual#page=238
func CLAC() {
	unsafe.Asm("CLAC", nil)
}

// CLC
// Clear CF flag.
//
// Documentation: https://golang.org/s/x86manual#page=239
func CLC() {
	unsafe.Asm("CLC", nil)
}

// CLD
// Clear DF flag.
//
// Documentation: https://golang.org/s/x86manual#page=240
func CLD() {
	unsafe.Asm("CLD", nil)
}

// CLFLUSH
// Flushes cache line containing m8.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=241
func CLFLUSH(rm interface{}) {
	unsafe.Asm("CLFLUSH", rm)
}

// CLFLUSHOPT
// Flushes cache line containing m8.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=243
func CLFLUSHOPT(rm interface{}) {
	unsafe.Asm("CLFLUSHOPT", rm)
}

// CLI
// Clear interrupt flag; interrupts disabled when interrupt flag cleared.
//
// Documentation: https://golang.org/s/x86manual#page=245
func CLI() {
	unsafe.Asm("CLI", nil)
}

// CLTS
// Clears TS flag in CR0.
//
// Documentation: https://golang.org/s/x86manual#page=247
func CLTS() {
	unsafe.Asm("CLTS", nil)
}

// CMC
// Complement CF flag.
//
// Documentation: https://golang.org/s/x86manual#page=250
func CMC() {
	unsafe.Asm("CMC", nil)
}

// CMOVA
// Move if above (CF=0 and ZF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVA(reg, rm interface{}) {
	unsafe.Asm("CMOVA", reg, rm)
}

// CMOVAE
// Move if above or equal (CF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVAE(reg, rm interface{}) {
	unsafe.Asm("CMOVAE", reg, rm)
}

// CMOVB
// Move if below (CF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVB(reg, rm interface{}) {
	unsafe.Asm("CMOVB", reg, rm)
}

// CMOVBE
// Move if below or equal (CF=1 or ZF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVBE(reg, rm interface{}) {
	unsafe.Asm("CMOVBE", reg, rm)
}

// CMOVC
// Move if carry (CF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVC(reg, rm interface{}) {
	unsafe.Asm("CMOVC", reg, rm)
}

// CMOVE
// Move if equal (ZF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVE(reg, rm interface{}) {
	unsafe.Asm("CMOVE", reg, rm)
}

// CMOVG
// Move if greater (ZF=0 and SF=OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVG(reg, rm interface{}) {
	unsafe.Asm("CMOVG", reg, rm)
}

// CMOVGE
// Move if greater or equal (SF=OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVGE(reg, rm interface{}) {
	unsafe.Asm("CMOVGE", reg, rm)
}

// CMOVL
// Move if less (SF≠ OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVL(reg, rm interface{}) {
	unsafe.Asm("CMOVL", reg, rm)
}

// CMOVLE
// Move if less or equal (ZF=1 or SF≠ OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVLE(reg, rm interface{}) {
	unsafe.Asm("CMOVLE", reg, rm)
}

// CMOVNA
// Move if not above (CF=1 or ZF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNA(reg, rm interface{}) {
	unsafe.Asm("CMOVNA", reg, rm)
}

// CMOVNAE
// Move if not above or equal (CF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNAE(reg, rm interface{}) {
	unsafe.Asm("CMOVNAE", reg, rm)
}

// CMOVNB
// Move if not below (CF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNB(reg, rm interface{}) {
	unsafe.Asm("CMOVNB", reg, rm)
}

// CMOVNBE
// Move if not below or equal (CF=0 and ZF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNBE(reg, rm interface{}) {
	unsafe.Asm("CMOVNBE", reg, rm)
}

// CMOVNC
// Move if not carry (CF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNC(reg, rm interface{}) {
	unsafe.Asm("CMOVNC", reg, rm)
}

// CMOVNE
// Move if not equal (ZF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNE(reg, rm interface{}) {
	unsafe.Asm("CMOVNE", reg, rm)
}

// CMOVNG
// Move if not greater (ZF=1 or SF≠ OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNG(reg, rm interface{}) {
	unsafe.Asm("CMOVNG", reg, rm)
}

// CMOVNGE
// Move if not greater or equal (SF≠ OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNGE(reg, rm interface{}) {
	unsafe.Asm("CMOVNGE", reg, rm)
}

// CMOVNL
// Move if not less (SF=OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNL(reg, rm interface{}) {
	unsafe.Asm("CMOVNL", reg, rm)
}

// CMOVNLE
// Move if not less or equal (ZF=0 and SF=OF).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNLE(reg, rm interface{}) {
	unsafe.Asm("CMOVNLE", reg, rm)
}

// CMOVNO
// Move if not overflow (OF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNO(reg, rm interface{}) {
	unsafe.Asm("CMOVNO", reg, rm)
}

// CMOVNP
// Move if not parity (PF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNP(reg, rm interface{}) {
	unsafe.Asm("CMOVNP", reg, rm)
}

// CMOVNS
// Move if not sign (SF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNS(reg, rm interface{}) {
	unsafe.Asm("CMOVNS", reg, rm)
}

// CMOVNZ
// Move if not zero (ZF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVNZ(reg, rm interface{}) {
	unsafe.Asm("CMOVNZ", reg, rm)
}

// CMOVO
// Move if overflow (OF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVO(reg, rm interface{}) {
	unsafe.Asm("CMOVO", reg, rm)
}

// CMOVP
// Move if parity (PF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVP(reg, rm interface{}) {
	unsafe.Asm("CMOVP", reg, rm)
}

// CMOVPE
// Move if parity even (PF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVPE(reg, rm interface{}) {
	unsafe.Asm("CMOVPE", reg, rm)
}

// CMOVPO
// Move if parity odd (PF=0).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVPO(reg, rm interface{}) {
	unsafe.Asm("CMOVPO", reg, rm)
}

// CMOVS
// Move if sign (SF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVS(reg, rm interface{}) {
	unsafe.Asm("CMOVS", reg, rm)
}

// CMOVZ
// Move if zero (ZF=1).
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=251
func CMOVZ(reg, rm interface{}) {
	unsafe.Asm("CMOVZ", reg, rm)
}

// CMP_I
// Compare imm8 with AL.
// Compare imm16 with AX.
// Compare imm32 with EAX.
// Compare imm32 sign-extended to 64-bits with RAX.
//
// al: AL/AX/EAX/RAX (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=255
func CMP_I(al, imm interface{}) {
	unsafe.Asm("CMP", al, imm)
}

// CMP_MI
// Compare imm16 with r/m16.
// Compare imm8 with r/m16.
// Compare imm32 with r/m32.
// Compare imm8 with r/m32.
// Compare imm32 sign-extended to 64-bits with r/m64.
// Compare imm8 with r/m64.
// Compare imm8 with r/m8.
//
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=255
func CMP_MI(rm, imm interface{}) {
	unsafe.Asm("CMP", rm, imm)
}

// CMP_MR
// Compare r16 with r/m16.
// Compare r32 with r/m32.
// Compare r64 with r/m64.
// Compare r8 with r/m8.
//
// rm: ModRM:r/m (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=255
func CMP_MR(rm, reg interface{}) {
	unsafe.Asm("CMP", rm, reg)
}

// CMP_RM
// Compare r/m16 with r16.
// Compare r/m32 with r32.
// Compare r/m64 with r64.
// Compare r/m8 with r8.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=255
func CMP_RM(reg, rm interface{}) {
	unsafe.Asm("CMP", reg, rm)
}

// CMPPD
// Compare packed double-precision floating-point values in xmm2/m128 and xmm1 using bits 2:0 of imm8 as a comparison predicate.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=257
func CMPPD(reg, rm, imm interface{}) {
	unsafe.Asm("CMPPD", reg, rm, imm)
}

// CMPPS
// Compare packed single-precision floating-point values in xmm2/m128 and xmm1 using bits 2:0 of imm8 as a comparison predicate.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=264
func CMPPS(reg, rm, imm interface{}) {
	unsafe.Asm("CMPPS", reg, rm, imm)
}

// CMPSB
// For legacy mode, compare byte at address DS:(E)SI with byte at address ES:(E)DI; For 64-bit mode compare byte at address (R|E)SI with byte at address (R|E)DI. The status flags are set accordingly.
//
// Documentation: https://golang.org/s/x86manual#page=271
func CMPSB() {
	unsafe.Asm("CMPSB", nil)
}

// CMPSD_NP
// For legacy mode, compare dword at address DS:(E)SI with dword at address ES:(E)DI; For 64-bit mode compare dword at address (R|E)SI with dword at address (R|E)DI. The status flags are set accordingly.
//
// Documentation: https://golang.org/s/x86manual#page=271
func CMPSD_NP() {
	unsafe.Asm("CMPSD", nil)
}

// CMPSD_RMI
// Compare low double-precision floating-point value in xmm2/m64 and xmm1 using bits 2:0 of imm8 as comparison predicate.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=275
func CMPSD_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("CMPSD", reg, rm, imm)
}

// CMPSQ
// Compares quadword at address (R|E)SI with quadword at address (R|E)DI and sets the status flags accordingly.
//
// Documentation: https://golang.org/s/x86manual#page=271
func CMPSQ() {
	unsafe.Asm("CMPSQ", nil)
}

// CMPSS
// Compare low single-precision floating-point value in xmm2/m32 and xmm1 using bits 2:0 of imm8 as comparison predicate.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=279
func CMPSS(reg, rm, imm interface{}) {
	unsafe.Asm("CMPSS", reg, rm, imm)
}

// CMPSW
// For legacy mode, compare word at address DS:(E)SI with word at address ES:(E)DI; For 64-bit mode compare word at address (R|E)SI with word at address (R|E)DI. The status flags are set accordingly.
//
// Documentation: https://golang.org/s/x86manual#page=271
func CMPSW() {
	unsafe.Asm("CMPSW", nil)
}

// CMPXCHG
// Compare AX with r/m16. If equal, ZF is set and r16 is loaded into r/m16. Else, clear ZF and load r/m16 into AX.
// Compare EAX with r/m32. If equal, ZF is set and r32 is loaded into r/m32. Else, clear ZF and load r/m32 into EAX.
// Compare RAX with r/m64. If equal, ZF is set and r64 is loaded into r/m64. Else, clear ZF and load r/m64 into RAX.
// Compare AL with r/m8. If equal, ZF is set and r8 is loaded into r/m8. Else, clear ZF and load r/m8 into AL.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=283
func CMPXCHG(rm, reg interface{}) {
	unsafe.Asm("CMPXCHG", rm, reg)
}

// CMPXCHG16B
// Compare RDX:RAX with m128. If equal, set ZF and load RCX:RBX into m128. Else, clear ZF and load m128 into RDX:RAX.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=285
func CMPXCHG16B(rm interface{}) {
	unsafe.Asm("CMPXCHG16B", rm)
}

// CMPXCHG8B
// Compare EDX:EAX with m64. If equal, set ZF and load ECX:EBX into m64. Else, clear ZF and load m64 into EDX:EAX.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=285
func CMPXCHG8B(rm interface{}) {
	unsafe.Asm("CMPXCHG8B", rm)
}

// COMISD
// Compare low double-precision floating-point values in xmm1 and xmm2/mem64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=288
func COMISD(reg, rm interface{}) {
	unsafe.Asm("COMISD", reg, rm)
}

// COMISS
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=290
func COMISS(reg, rm interface{}) {
	unsafe.Asm("COMISS", reg, rm)
}

// CPUID
// Returns processor identification and feature information to the EAX, EBX, ECX, and EDX registers, as determined by input entered in EAX (in some cases, ECX as well).
//
// Documentation: https://golang.org/s/x86manual#page=292
func CPUID() {
	unsafe.Asm("CPUID", nil)
}

// CQO
// RDX:RAX← sign-extend of RAX.
//
// Documentation: https://golang.org/s/x86manual#page=380
func CQO() {
	unsafe.Asm("CQO", nil)
}

// CRC32
// Accumulate CRC32 on r/m16.
// Accumulate CRC32 on r/m32.
// Accumulate CRC32 on r/m8.
// Accumulate CRC32 on r/m64.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=327
func CRC32(reg, rm interface{}) {
	unsafe.Asm("CRC32", reg, rm)
}

// CVTDQ2PD
// Convert two packed signed doubleword integers from xmm2/mem to two packed double-precision floating- point values in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=330
func CVTDQ2PD(reg, rm interface{}) {
	unsafe.Asm("CVTDQ2PD", reg, rm)
}

// CVTPD2PI
// Convert two packed double-precision floating- point values from xmm/m128 to two packed signed doubleword integers in mm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=341
func CVTPD2PI(reg, rm interface{}) {
	unsafe.Asm("CVTPD2PI", reg, rm)
}

// CVTPD2PS
// Convert two packed double-precision floating-point values in xmm2/mem to two single-precision floating-point values in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=342
func CVTPD2PS(reg, rm interface{}) {
	unsafe.Asm("CVTPD2PS", reg, rm)
}

// CVTPI2PD
// Convert two packed signed doubleword integers from mm/mem64 to two packed double-precision floating-point values in xmm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=346
func CVTPI2PD(reg, rm interface{}) {
	unsafe.Asm("CVTPI2PD", reg, rm)
}

// CVTPI2PS
// Convert two signed doubleword integers from mm/m64 to two single-precision floating-point values in xmm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=347
func CVTPI2PS(reg, rm interface{}) {
	unsafe.Asm("CVTPI2PS", reg, rm)
}

// CVTPS2DQ
// Convert four packed single-precision floating-point values from xmm2/mem to four packed signed doubleword values in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=348
func CVTPS2DQ(reg, rm interface{}) {
	unsafe.Asm("CVTPS2DQ", reg, rm)
}

// CVTPS2PD
// Convert two packed single-precision floating-point values in xmm2/m64 to two packed double-precision floating-point values in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=351
func CVTPS2PD(reg, rm interface{}) {
	unsafe.Asm("CVTPS2PD", reg, rm)
}

// CVTPS2PI
// Convert two packed single-precision floating- point values from xmm/m64 to two packed signed doubleword integers in mm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=354
func CVTPS2PI(reg, rm interface{}) {
	unsafe.Asm("CVTPS2PI", reg, rm)
}

// CVTSD2SI
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer r32.
// Convert one double-precision floating-point value from xmm1/m64 to one signed quadword integer sign- extended into r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=355
func CVTSD2SI(reg, rm interface{}) {
	unsafe.Asm("CVTSD2SI", reg, rm)
}

// CVTSD2SS
// Convert one double-precision floating-point value in xmm2/m64 to one single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=357
func CVTSD2SS(reg, rm interface{}) {
	unsafe.Asm("CVTSD2SS", reg, rm)
}

// CVTSI2SD
// Convert one signed doubleword integer from r32/m32 to one double-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=359
func CVTSI2SD(reg, rm interface{}) {
	unsafe.Asm("CVTSI2SD", reg, rm)
}

// CVTSI2SS
// Convert one signed doubleword integer from r/m32 to one single-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=361
func CVTSI2SS(reg, rm interface{}) {
	unsafe.Asm("CVTSI2SS", reg, rm)
}

// CVTSS2SD
// Convert one single-precision floating-point value in xmm2/m32 to one double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=363
func CVTSS2SD(reg, rm interface{}) {
	unsafe.Asm("CVTSS2SD", reg, rm)
}

// CVTSS2SI
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=365
func CVTSS2SI(reg, rm interface{}) {
	unsafe.Asm("CVTSS2SI", reg, rm)
}

// CVTTPD2DQ
// Convert two packed double-precision floating-point values in xmm2/mem to two signed doubleword integers in xmm1 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=367
func CVTTPD2DQ(reg, rm interface{}) {
	unsafe.Asm("CVTTPD2DQ", reg, rm)
}

// CVTTPD2PI
// Convert two packer double-precision floating- point values from xmm/m128 to two packed signed doubleword integers in mm using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=371
func CVTTPD2PI(reg, rm interface{}) {
	unsafe.Asm("CVTTPD2PI", reg, rm)
}

// CVTTPS2DQ
// Convert four packed single-precision floating-point values from xmm2/mem to four packed signed doubleword values in xmm1 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=372
func CVTTPS2DQ(reg, rm interface{}) {
	unsafe.Asm("CVTTPS2DQ", reg, rm)
}

// CVTTPS2PI
// Convert two single-precision floating-point values from xmm/m64 to two signed doubleword signed integers in mm using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=375
func CVTTPS2PI(reg, rm interface{}) {
	unsafe.Asm("CVTTPS2PI", reg, rm)
}

// CVTTSD2SI
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer in r32 using truncation.
// Convert one double-precision floating-point value from xmm1/m64 to one signed quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=376
func CVTTSD2SI(reg, rm interface{}) {
	unsafe.Asm("CVTTSD2SI", reg, rm)
}

// CVTTSS2SI
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32 using truncation.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=378
func CVTTSS2SI(reg, rm interface{}) {
	unsafe.Asm("CVTTSS2SI", reg, rm)
}

// CWD
// DX:AX ← sign-extend of AX.
//
// Documentation: https://golang.org/s/x86manual#page=380
func CWD() {
	unsafe.Asm("CWD", nil)
}

// CWDE
// EAX ← sign-extend of AX.
//
// Documentation: https://golang.org/s/x86manual#page=237
func CWDE() {
	unsafe.Asm("CWDE", nil)
}

// DAA
// Decimal adjust AL after addition.
//
// Documentation: https://golang.org/s/x86manual#page=381
func DAA() {
	unsafe.Asm("DAA", nil)
}

// DAS
// Decimal adjust AL after subtraction.
//
// Documentation: https://golang.org/s/x86manual#page=383
func DAS() {
	unsafe.Asm("DAS", nil)
}

// DEC_M
// Decrement r/m16 by 1.
// Decrement r/m32 by 1.
// Decrement r/m64 by 1.
// Decrement r/m8 by 1.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=385
func DEC_M(rm interface{}) {
	unsafe.Asm("DEC", rm)
}

// DEC_O
// Decrement r16 by 1.
// Decrement r32 by 1.
//
// opcode: opcode + rd (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=385
func DEC_O(opcode interface{}) {
	unsafe.Asm("DEC", opcode)
}

// DIV
// Unsigned divide DX:AX by r/m16, with result stored in AX ← Quotient, DX ← Remainder.
// Unsigned divide EDX:EAX by r/m32, with result stored in EAX ← Quotient, EDX ← Remainder.
// Unsigned divide RDX:RAX by r/m64, with result stored in RAX ← Quotient, RDX ← Remainder.
// Unsigned divide AX by r/m8, with result stored in AL ← Quotient, AH ← Remainder.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=387
func DIV(rm interface{}) {
	unsafe.Asm("DIV", rm)
}

// DIVPD
// Divide packed double-precision floating-point values in xmm1 by packed double-precision floating-point values in xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=390
func DIVPD(reg, rm interface{}) {
	unsafe.Asm("DIVPD", reg, rm)
}

// DIVPS
// Divide packed single-precision floating-point values in xmm1 by packed single-precision floating-point values in xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=393
func DIVPS(reg, rm interface{}) {
	unsafe.Asm("DIVPS", reg, rm)
}

// DIVSD
// Divide low double-precision floating-point value in xmm1 by low double-precision floating-point value in xmm2/m64.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=396
func DIVSD(reg, rm interface{}) {
	unsafe.Asm("DIVSD", reg, rm)
}

// DIVSS
// Divide low single-precision floating-point value in xmm1 by low single-precision floating-point value in xmm2/m32.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=398
func DIVSS(reg, rm interface{}) {
	unsafe.Asm("DIVSS", reg, rm)
}

// DPPD
// Selectively multiply packed DP floating-point values from xmm1 with packed DP floating- point values from xmm2, add and selectively store the packed DP floating-point values to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=400
func DPPD(reg, rm, imm interface{}) {
	unsafe.Asm("DPPD", reg, rm, imm)
}

// DPPS
// Selectively multiply packed SP floating-point values from xmm1 with packed SP floating- point values from xmm2, add and selectively store the packed SP floating-point values or zero values to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=402
func DPPS(reg, rm, imm interface{}) {
	unsafe.Asm("DPPS", reg, rm, imm)
}

// EMMS
// Set the x87 FPU tag word to empty.
//
// Documentation: https://golang.org/s/x86manual#page=405
func EMMS() {
	unsafe.Asm("EMMS", nil)
}

// ENTER
// Create a stack frame for a procedure.
// Create a stack frame with a nested pointer for a procedure.
// Create a stack frame with nested pointers for a procedure.
//
// iw: iw
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=406
func ENTER(iw, imm interface{}) {
	unsafe.Asm("ENTER", iw, imm)
}

// F2XM1
// Replace ST(0) with (2 – 1).
//
// Documentation: https://golang.org/s/x86manual#page=411
func F2XM1() {
	unsafe.Asm("F2XM1", nil)
}

// FABS
// Replace ST with its absolute value.
//
// Documentation: https://golang.org/s/x86manual#page=413
func FABS() {
	unsafe.Asm("FABS", nil)
}

// FADD
// Add ST(0) to ST(i) and store result in ST(0).
// Add ST(i) to ST(0) and store result in ST(i).
// Add m32fp to ST(0) and store result in ST(0).
// Add m64fp to ST(0) and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=414
func FADD(st0, sti interface{}) {
	unsafe.Asm("FADD", st0, sti)
}

// FADDP
// Add ST(0) to ST(1), store result in ST(1), and pop the register stack.
// Add ST(0) to ST(i), store result in ST(i), and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=414
func FADDP() {
	unsafe.Asm("FADDP", nil)
}

// FBLD
// Convert BCD value to floating-point and push onto the FPU stack.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=417
func FBLD(rm interface{}) {
	unsafe.Asm("FBLD", rm)
}

// FBSTP
// Store ST(0) in m80bcd and pop ST(0).
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=419
func FBSTP(rm interface{}) {
	unsafe.Asm("FBSTP", rm)
}

// FCHS
// Complements sign of ST(0).
//
// Documentation: https://golang.org/s/x86manual#page=421
func FCHS() {
	unsafe.Asm("FCHS", nil)
}

// FCLEX
// Clear floating-point exception flags after checking for pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=423
func FCLEX() {
	unsafe.Asm("FCLEX", nil)
}

// FCMOVB
// Move if below (CF=1).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVB(st0, sti interface{}) {
	unsafe.Asm("FCMOVB", st0, sti)
}

// FCMOVBE
// Move if below or equal (CF=1 or ZF=1).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVBE(st0, sti interface{}) {
	unsafe.Asm("FCMOVBE", st0, sti)
}

// FCMOVE
// Move if equal (ZF=1).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVE(st0, sti interface{}) {
	unsafe.Asm("FCMOVE", st0, sti)
}

// FCMOVNB
// Move if not below (CF=0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVNB(st0, sti interface{}) {
	unsafe.Asm("FCMOVNB", st0, sti)
}

// FCMOVNBE
// Move if not below or equal (CF=0 and ZF=0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVNBE(st0, sti interface{}) {
	unsafe.Asm("FCMOVNBE", st0, sti)
}

// FCMOVNE
// Move if not equal (ZF=0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVNE(st0, sti interface{}) {
	unsafe.Asm("FCMOVNE", st0, sti)
}

// FCMOVNU
// Move if not unordered (PF=0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVNU(st0, sti interface{}) {
	unsafe.Asm("FCMOVNU", st0, sti)
}

// FCMOVU
// Move if unordered (PF=1).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=425
func FCMOVU(st0, sti interface{}) {
	unsafe.Asm("FCMOVU", st0, sti)
}

// FCOM
// Compare ST(0) with ST(1).
// Compare ST(0) with ST(i).
// Compare ST(0) with m32fp.
// Compare ST(0) with m64fp.
//
// Documentation: https://golang.org/s/x86manual#page=427
func FCOM() {
	unsafe.Asm("FCOM", nil)
}

// FCOMI
// Compare ST(0) with ST(i) and set status flags accordingly.
//
// st0: ST(0) (r)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=430
func FCOMI(st0, sti interface{}) {
	unsafe.Asm("FCOMI", st0, sti)
}

// FCOMIP
// Compare ST(0) with ST(i), set status flags accordingly, and pop register stack.
//
// st0: ST(0) (r)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=430
func FCOMIP(st0, sti interface{}) {
	unsafe.Asm("FCOMIP", st0, sti)
}

// FCOMP
// Compare ST(0) with ST(1) and pop register stack.
// Compare ST(0) with ST(i) and pop register stack.
// Compare ST(0) with m32fp and pop register stack.
// Compare ST(0) with m64fp and pop register stack.
//
// Documentation: https://golang.org/s/x86manual#page=427
func FCOMP() {
	unsafe.Asm("FCOMP", nil)
}

// FCOMPP
// Compare ST(0) with ST(1) and pop register stack twice.
//
// Documentation: https://golang.org/s/x86manual#page=427
func FCOMPP() {
	unsafe.Asm("FCOMPP", nil)
}

// FCOS
// Replace ST(0) with its approximate cosine.
//
// Documentation: https://golang.org/s/x86manual#page=433
func FCOS() {
	unsafe.Asm("FCOS", nil)
}

// FDECSTP
// Decrement TOP field in FPU status word.
//
// Documentation: https://golang.org/s/x86manual#page=435
func FDECSTP() {
	unsafe.Asm("FDECSTP", nil)
}

// FDIV
// Divide ST(0) by ST(i) and store result in ST(0).
// Divide ST(i) by ST(0) and store result in ST(i).
// Divide ST(0) by m32fp and store result in ST(0).
// Divide ST(0) by m64fp and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=436
func FDIV(st0, sti interface{}) {
	unsafe.Asm("FDIV", st0, sti)
}

// FDIVP
// Divide ST(1) by ST(0), store result in ST(1), and pop the register stack.
// Divide ST(i) by ST(0), store result in ST(i), and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=436
func FDIVP() {
	unsafe.Asm("FDIVP", nil)
}

// FDIVR
// Divide ST(i) by ST(0) and store result in ST(0).
// Divide ST(0) by ST(i) and store result in ST(i).
// Divide m32fp by ST(0) and store result in ST(0).
// Divide m64fp by ST(0) and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=439
func FDIVR(st0, sti interface{}) {
	unsafe.Asm("FDIVR", st0, sti)
}

// FDIVRP
// Divide ST(0) by ST(1), store result in ST(1), and pop the register stack.
// Divide ST(0) by ST(i), store result in ST(i), and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=439
func FDIVRP() {
	unsafe.Asm("FDIVRP", nil)
}

// FFREE
// Sets tag for ST(i) to empty.
//
// sti: ST(i) (w)
//
// Documentation: https://golang.org/s/x86manual#page=442
func FFREE(sti interface{}) {
	unsafe.Asm("FFREE", sti)
}

// FIADD
// Add m16int to ST(0) and store result in ST(0).
// Add m32int to ST(0) and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=414
func FIADD(rm interface{}) {
	unsafe.Asm("FIADD", rm)
}

// FICOM
// Compare ST(0) with m16int.
// Compare ST(0) with m32int.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=443
func FICOM(rm interface{}) {
	unsafe.Asm("FICOM", rm)
}

// FICOMP
// Compare ST(0) with m16int and pop stack register.
// Compare ST(0) with m32int and pop stack register.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=443
func FICOMP(rm interface{}) {
	unsafe.Asm("FICOMP", rm)
}

// FIDIV
// Divide ST(0) by m16int and store result in ST(0).
// Divide ST(0) by m32int and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=436
func FIDIV(rm interface{}) {
	unsafe.Asm("FIDIV", rm)
}

// FIDIVR
// Divide m16int by ST(0) and store result in ST(0).
// Divide m32int by ST(0) and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=439
func FIDIVR(rm interface{}) {
	unsafe.Asm("FIDIVR", rm)
}

// FILD
// Push m16int onto the FPU register stack.
// Push m32int onto the FPU register stack.
// Push m64int onto the FPU register stack.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=445
func FILD(rm interface{}) {
	unsafe.Asm("FILD", rm)
}

// FIMUL
// Multiply ST(0) by m16int and store result in ST(0).
// Multiply ST(0) by m32int and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=463
func FIMUL(rm interface{}) {
	unsafe.Asm("FIMUL", rm)
}

// FINCSTP
// Increment the TOP field in the FPU status register.
//
// Documentation: https://golang.org/s/x86manual#page=447
func FINCSTP() {
	unsafe.Asm("FINCSTP", nil)
}

// FINIT
// Initialize FPU after checking for pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=448
func FINIT() {
	unsafe.Asm("FINIT", nil)
}

// FIST
// Store ST(0) in m16int.
// Store ST(0) in m32int.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=450
func FIST(rm interface{}) {
	unsafe.Asm("FIST", rm)
}

// FISTP
// Store ST(0) in m16int and pop register stack.
// Store ST(0) in m32int and pop register stack.
// Store ST(0) in m64int and pop register stack.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=450
func FISTP(rm interface{}) {
	unsafe.Asm("FISTP", rm)
}

// FISTTP
// Store ST(0) in m16int with truncation.
// Store ST(0) in m32int with truncation.
// Store ST(0) in m64int with truncation.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=453
func FISTTP(rm interface{}) {
	unsafe.Asm("FISTTP", rm)
}

// FISUB
// Subtract m16int from ST(0) and store result in ST(0).
// Subtract m32int from ST(0) and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=497
func FISUB(rm interface{}) {
	unsafe.Asm("FISUB", rm)
}

// FISUBR
// Subtract ST(0) from m16int and store result in ST(0).
// Subtract ST(0) from m32int and store result in ST(0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=500
func FISUBR(rm interface{}) {
	unsafe.Asm("FISUBR", rm)
}

// FLD
// Push ST(i) onto the FPU register stack.
// Push m32fp onto the FPU register stack.
// Push m64fp onto the FPU register stack.
// Push m80fp onto the FPU register stack.
//
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=455
func FLD(sti interface{}) {
	unsafe.Asm("FLD", sti)
}

// FLD1
// Push +1.0 onto the FPU register stack.
//
// Documentation: https://golang.org/s/x86manual#page=457
func FLD1() {
	unsafe.Asm("FLD1", nil)
}

// FLDCW
// Load FPU control word from m2byte.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=459
func FLDCW(rm interface{}) {
	unsafe.Asm("FLDCW", rm)
}

// FLDENV
// Load FPU environment from m14byte or m28byte.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=461
func FLDENV(rm interface{}) {
	unsafe.Asm("FLDENV", rm)
}

// FLDL2E
// Push log e onto the FPU register stack.
//
// Documentation: https://golang.org/s/x86manual#page=457
func FLDL2E() {
	unsafe.Asm("FLDL2E", nil)
}

// FLDL2T
// Push log 10 onto the FPU register stack.
//
// Documentation: https://golang.org/s/x86manual#page=457
func FLDL2T() {
	unsafe.Asm("FLDL2T", nil)
}

// FLDLG2
// Push log 2 onto the FPU register stack.
//
// Documentation: https://golang.org/s/x86manual#page=457
func FLDLG2() {
	unsafe.Asm("FLDLG2", nil)
}

// FLDPI
// Push π onto the FPU register stack.
//
// Documentation: https://golang.org/s/x86manual#page=457
func FLDPI() {
	unsafe.Asm("FLDPI", nil)
}

// FMUL
// Multiply ST(0) by ST(i) and store result in ST(0).
// Multiply ST(i) by ST(0) and store result in ST(i).
// Multiply ST(0) by m32fp and store result in ST(0).
// Multiply ST(0) by m64fp and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=463
func FMUL(st0, sti interface{}) {
	unsafe.Asm("FMUL", st0, sti)
}

// FMULP
// Multiply ST(1) by ST(0), store result in ST(1), and pop the register stack.
// Multiply ST(i) by ST(0), store result in ST(i), and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=463
func FMULP() {
	unsafe.Asm("FMULP", nil)
}

// FNCLEX
// Clear floating-point exception flags without checking for pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=423
func FNCLEX() {
	unsafe.Asm("FNCLEX", nil)
}

// FNINIT
// Initialize FPU without checking for pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=448
func FNINIT() {
	unsafe.Asm("FNINIT", nil)
}

// FNOP
// No operation is performed.
//
// Documentation: https://golang.org/s/x86manual#page=466
func FNOP() {
	unsafe.Asm("FNOP", nil)
}

// FNSAVE
// Store FPU environment to m94byte or m108byte without checking for pending unmasked floating- point exceptions. Then re-initialize the FPU.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=478
func FNSAVE(rm interface{}) {
	unsafe.Asm("FNSAVE", rm)
}

// FNSTCW
// Store FPU control word to m2byte without checking for pending unmasked floating-point exceptions.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=491
func FNSTCW(rm interface{}) {
	unsafe.Asm("FNSTCW", rm)
}

// FNSTENV
// Store FPU environment to m14byte or m28byte without checking for pending unmasked floating- point exceptions. Then mask all floating- point exceptions.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=493
func FNSTENV(rm interface{}) {
	unsafe.Asm("FNSTENV", rm)
}

// FNSTSW
// Store FPU status word in AX register without checking for pending unmasked floating-point exceptions.
// Store FPU status word at m2byte without checking for pending unmasked floating-point exceptions.
//
// ax: AX (w)
//
// Documentation: https://golang.org/s/x86manual#page=495
func FNSTSW(ax interface{}) {
	unsafe.Asm("FNSTSW", ax)
}

// FPATAN
// Replace ST(1) with arctan(ST(1)/ST(0)) and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=467
func FPATAN() {
	unsafe.Asm("FPATAN", nil)
}

// FPREM
// Replace ST(0) with the remainder obtained from dividing ST(0) by ST(1).
//
// Documentation: https://golang.org/s/x86manual#page=469
func FPREM() {
	unsafe.Asm("FPREM", nil)
}

// FPREM1
// Replace ST(0) with the IEEE remainder obtained from dividing ST(0) by ST(1).
//
// Documentation: https://golang.org/s/x86manual#page=471
func FPREM1() {
	unsafe.Asm("FPREM1", nil)
}

// FPTAN
// Replace ST(0) with its approximate tangent and push 1 onto the FPU stack.
//
// Documentation: https://golang.org/s/x86manual#page=473
func FPTAN() {
	unsafe.Asm("FPTAN", nil)
}

// FRNDINT
// Round ST(0) to an integer.
//
// Documentation: https://golang.org/s/x86manual#page=475
func FRNDINT() {
	unsafe.Asm("FRNDINT", nil)
}

// FRSTOR
// Load FPU state from m94byte or m108byte.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=476
func FRSTOR(rm interface{}) {
	unsafe.Asm("FRSTOR", rm)
}

// FSAVE
// Store FPU state to m94byte or m108byte after checking for pending unmasked floating-point exceptions. Then re-initialize the FPU.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=478
func FSAVE(rm interface{}) {
	unsafe.Asm("FSAVE", rm)
}

// FSCALE
// Scale ST(0) by ST(1).
//
// Documentation: https://golang.org/s/x86manual#page=481
func FSCALE() {
	unsafe.Asm("FSCALE", nil)
}

// FSIN
// Replace ST(0) with the approximate of its sine.
//
// Documentation: https://golang.org/s/x86manual#page=483
func FSIN() {
	unsafe.Asm("FSIN", nil)
}

// FSINCOS
// Compute the sine and cosine of ST(0); replace ST(0) with the approximate sine, and push the approximate cosine onto the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=485
func FSINCOS() {
	unsafe.Asm("FSINCOS", nil)
}

// FSQRT
// Computes square root of ST(0) and stores the result in ST(0).
//
// Documentation: https://golang.org/s/x86manual#page=487
func FSQRT() {
	unsafe.Asm("FSQRT", nil)
}

// FST
// Copy ST(0) to ST(i).
// Copy ST(0) to m32fp.
// Copy ST(0) to m64fp.
//
// sti: ST(i) (w)
//
// Documentation: https://golang.org/s/x86manual#page=489
func FST(sti interface{}) {
	unsafe.Asm("FST", sti)
}

// FSTCW
// Store FPU control word to m2byte after checking for pending unmasked floating-point exceptions.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=491
func FSTCW(rm interface{}) {
	unsafe.Asm("FSTCW", rm)
}

// FSTENV
// Store FPU environment to m14byte or m28byte after checking for pending unmasked floating-point exceptions. Then mask all floating-point exceptions.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=493
func FSTENV(rm interface{}) {
	unsafe.Asm("FSTENV", rm)
}

// FSTP
// Copy ST(0) to ST(i) and pop register stack.
// Copy ST(0) to m32fp and pop register stack.
// Copy ST(0) to m64fp and pop register stack.
// Copy ST(0) to m80fp and pop register stack.
//
// sti: ST(i) (w)
//
// Documentation: https://golang.org/s/x86manual#page=489
func FSTP(sti interface{}) {
	unsafe.Asm("FSTP", sti)
}

// FSTSW
// Store FPU status word in AX register after checking for pending unmasked floating-point exceptions.
// Store FPU status word at m2byte after checking for pending unmasked floating-point exceptions.
//
// ax: AX (w)
//
// Documentation: https://golang.org/s/x86manual#page=495
func FSTSW(ax interface{}) {
	unsafe.Asm("FSTSW", ax)
}

// FSUB
// Subtract ST(i) from ST(0) and store result in ST(0).
// Subtract ST(0) from ST(i) and store result in ST(i).
// Subtract m32fp from ST(0) and store result in ST(0).
// Subtract m64fp from ST(0) and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=497
func FSUB(st0, sti interface{}) {
	unsafe.Asm("FSUB", st0, sti)
}

// FSUBP
// Subtract ST(0) from ST(1), store result in ST(1), and pop register stack.
// Subtract ST(0) from ST(i), store result in ST(i), and pop register stack.
//
// Documentation: https://golang.org/s/x86manual#page=497
func FSUBP() {
	unsafe.Asm("FSUBP", nil)
}

// FSUBR
// Subtract ST(0) from ST(i) and store result in ST(0).
// Subtract ST(i) from ST(0) and store result in ST(i).
// Subtract ST(0) from m32fp and store result in ST(0).
// Subtract ST(0) from m64fp and store result in ST(0).
//
// st0: ST(0) (r, w)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=500
func FSUBR(st0, sti interface{}) {
	unsafe.Asm("FSUBR", st0, sti)
}

// FSUBRP
// Subtract ST(1) from ST(0), store result in ST(1), and pop register stack.
// Subtract ST(i) from ST(0), store result in ST(i), and pop register stack.
//
// Documentation: https://golang.org/s/x86manual#page=500
func FSUBRP() {
	unsafe.Asm("FSUBRP", nil)
}

// FTST
// Compare ST(0) with 0.0.
//
// Documentation: https://golang.org/s/x86manual#page=503
func FTST() {
	unsafe.Asm("FTST", nil)
}

// FUCOM
// Compare ST(0) with ST(1).
// Compare ST(0) with ST(i).
//
// Documentation: https://golang.org/s/x86manual#page=505
func FUCOM() {
	unsafe.Asm("FUCOM", nil)
}

// FUCOMI
// Compare ST(0) with ST(i), check for ordered values, and set status flags accordingly.
//
// st0: ST(0) (r)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=430
func FUCOMI(st0, sti interface{}) {
	unsafe.Asm("FUCOMI", st0, sti)
}

// FUCOMIP
// Compare ST(0) with ST(i), check for ordered values, set status flags accordingly, and pop register stack.
//
// st0: ST(0) (r)
// sti: ST(i) (r)
//
// Documentation: https://golang.org/s/x86manual#page=430
func FUCOMIP(st0, sti interface{}) {
	unsafe.Asm("FUCOMIP", st0, sti)
}

// FUCOMP
// Compare ST(0) with ST(1) and pop register stack.
// Compare ST(0) with ST(i) and pop register stack.
//
// Documentation: https://golang.org/s/x86manual#page=505
func FUCOMP() {
	unsafe.Asm("FUCOMP", nil)
}

// FUCOMPP
// Compare ST(0) with ST(1) and pop register stack twice.
//
// Documentation: https://golang.org/s/x86manual#page=505
func FUCOMPP() {
	unsafe.Asm("FUCOMPP", nil)
}

// FWAIT
// Check pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=1923
func FWAIT() {
	unsafe.Asm("FWAIT", nil)
}

// FXAM
// Classify value or number in ST(0).
//
// Documentation: https://golang.org/s/x86manual#page=508
func FXAM() {
	unsafe.Asm("FXAM", nil)
}

// FXCH
// Exchange the contents of ST(0) and ST(1).
// Exchange the contents of ST(0) and ST(i).
//
// Documentation: https://golang.org/s/x86manual#page=510
func FXCH() {
	unsafe.Asm("FXCH", nil)
}

// FXRSTOR
// Restore the x87 FPU, MMX, XMM, and MXCSR register state from m512byte.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=512
func FXRSTOR(rm interface{}) {
	unsafe.Asm("FXRSTOR", rm)
}

// FXRSTOR64
// Restore the x87 FPU, MMX, XMM, and MXCSR register state from m512byte.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=512
func FXRSTOR64(rm interface{}) {
	unsafe.Asm("FXRSTOR64", rm)
}

// FXSAVE
// Save the x87 FPU, MMX, XMM, and MXCSR register state to m512byte.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=515
func FXSAVE(rm interface{}) {
	unsafe.Asm("FXSAVE", rm)
}

// FXSAVE64
// Save the x87 FPU, MMX, XMM, and MXCSR register state to m512byte.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=515
func FXSAVE64(rm interface{}) {
	unsafe.Asm("FXSAVE64", rm)
}

// FXTRACT
// Separate value in ST(0) into exponent and significand, store exponent in ST(0), and push the significand onto the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=523
func FXTRACT() {
	unsafe.Asm("FXTRACT", nil)
}

// FYL2X
// Replace ST(1) with (ST(1) ∗ log ST(0)) and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=525
func FYL2X() {
	unsafe.Asm("FYL2X", nil)
}

// FYL2XP1
// Replace ST(1) with ST(1) ∗ log (ST(0) + 1.0) and pop the register stack.
//
// Documentation: https://golang.org/s/x86manual#page=527
func FYL2XP1() {
	unsafe.Asm("FYL2XP1", nil)
}

// HADDPD
// Horizontal add packed double-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=529
func HADDPD(reg, rm interface{}) {
	unsafe.Asm("HADDPD", reg, rm)
}

// HADDPS
// Horizontal add packed single-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=532
func HADDPS(reg, rm interface{}) {
	unsafe.Asm("HADDPS", reg, rm)
}

// HLT
// Halt
//
// Documentation: https://golang.org/s/x86manual#page=535
func HLT() {
	unsafe.Asm("HLT", nil)
}

// HSUBPD
// Horizontal subtract packed double-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=536
func HSUBPD(reg, rm interface{}) {
	unsafe.Asm("HSUBPD", reg, rm)
}

// HSUBPS
// Horizontal subtract packed single-precision floating-point values from xmm2/m128 to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=539
func HSUBPS(reg, rm interface{}) {
	unsafe.Asm("HSUBPS", reg, rm)
}

// IDIV
// Signed divide DX:AX by r/m16, with result stored in AX ← Quotient, DX ← Remainder.
// Signed divide EDX:EAX by r/m32, with result stored in EAX ← Quotient, EDX ← Remainder.
// Signed divide RDX:RAX by r/m64, with result stored in RAX ← Quotient, RDX ← Remainder.
// Signed divide AX by r/m8, with result stored in: AL ← Quotient, AH ← Remainder.
// Signed divide AX by r/m8, with result stored in AL ← Quotient, AH ← Remainder.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=542
func IDIV(rm interface{}) {
	unsafe.Asm("IDIV", rm)
}

// IMUL_M
// DX:AX ← AX ∗ r/m word.
// EDX:EAX ← EAX ∗ r/m32.
// RDX:RAX ← RAX ∗ r/m64.
// AX← AL ∗ r/m byte.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=545
func IMUL_M(rm interface{}) {
	unsafe.Asm("IMUL", rm)
}

// IMUL_RM
// word register ← word register ∗ r/m16.
// doubleword register ← doubleword register ∗ r/m32.
// Quadword register ← Quadword register ∗ r/m64.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=545
func IMUL_RM(reg, rm interface{}) {
	unsafe.Asm("IMUL", reg, rm)
}

// IMUL_RMI
// word register ← r/m16 ∗ immediate word.
// word register ← r/m16 ∗ sign-extended immediate byte.
// doubleword register ← r/m32 ∗ immediate doubleword.
// doubleword register ← r/m32 ∗ sign- extended immediate byte.
// Quadword register ← r/m64 ∗ immediate doubleword.
// Quadword register ← r/m64 ∗ sign-extended immediate byte.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=545
func IMUL_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("IMUL", reg, rm, imm)
}

// IN_I
// Input byte from imm8 I/O port address into AL.
// Input word from imm8 I/O port address into AX.
// Input dword from imm8 I/O port address into EAX.
//
// al: AL (w)
// imm: imm8 (r)
//
// Documentation: https://golang.org/s/x86manual#page=549
func IN_I(al, imm interface{}) {
	unsafe.Asm("IN", al, imm)
}

// IN_NP
// Input byte from I/O port in DX into AL.
// Input word from I/O port in DX into AX.
// Input doubleword from I/O port in DX into EAX.
//
// al: AL (w)
// dx: DX (r)
//
// Documentation: https://golang.org/s/x86manual#page=549
func IN_NP(al, dx interface{}) {
	unsafe.Asm("IN", al, dx)
}

// INC_M
// Increment r/m word by 1.
// Increment r/m doubleword by 1.
// Increment r/m quadword by 1.
// Increment r/m byte by 1.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=551
func INC_M(rm interface{}) {
	unsafe.Asm("INC", rm)
}

// INC_O
// Increment word register by 1.
// Increment doubleword register by 1.
//
// opcode: opcode + rd (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=551
func INC_O(opcode interface{}) {
	unsafe.Asm("INC", opcode)
}

// INSB
// Input byte from I/O port specified in DX into memory location specified with ES:(E)DI or RDI.
//
// Documentation: https://golang.org/s/x86manual#page=553
func INSB() {
	unsafe.Asm("INSB", nil)
}

// INSD
// Input doubleword from I/O port specified in DX into memory location specified in ES:(E)DI or RDI.
//
// Documentation: https://golang.org/s/x86manual#page=553
func INSD() {
	unsafe.Asm("INSD", nil)
}

// INSERTPS
// Insert a single-precision floating-point value selected by imm8 from xmm2/m32 into xmm1 at the specified destination element specified by imm8 and zero out destination elements in xmm1 as indicated in imm8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=556
func INSERTPS(reg, rm, imm interface{}) {
	unsafe.Asm("INSERTPS", reg, rm, imm)
}

// INSW
// Input word from I/O port specified in DX into memory location specified in ES:(E)DI or RDI.
//
// Documentation: https://golang.org/s/x86manual#page=553
func INSW() {
	unsafe.Asm("INSW", nil)
}

// INT_I
// Interrupt vector specified by immediate byte.
//
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=559
func INT_I(imm interface{}) {
	unsafe.Asm("INT", imm)
}

// INT_NP
// Interrupt 3—trap to debugger.
//
// v3: 3 (r)
//
// Documentation: https://golang.org/s/x86manual#page=559
func INT_NP(v3 interface{}) {
	unsafe.Asm("INT", v3)
}

// INTO
// Interrupt 4—if overflow flag is 1.
//
// Documentation: https://golang.org/s/x86manual#page=559
func INTO() {
	unsafe.Asm("INTO", nil)
}

// INVD
// Flush internal caches; initiate flushing of external caches.
//
// Documentation: https://golang.org/s/x86manual#page=571
func INVD() {
	unsafe.Asm("INVD", nil)
}

// INVLPG
// Invalidate TLB entries for page containing m.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=573
func INVLPG(rm interface{}) {
	unsafe.Asm("INVLPG", rm)
}

// INVPCID
// Invalidates entries in the TLBs and paging-structure caches based on invalidation type in r32 and descrip- tor in m128.
// Invalidates entries in the TLBs and paging-structure caches based on invalidation type in r64 and descrip- tor in m128.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=575
func INVPCID(reg, rm interface{}) {
	unsafe.Asm("INVPCID", reg, rm)
}

// IRET
// Interrupt return (16-bit operand size).
//
// Documentation: https://golang.org/s/x86manual#page=578
func IRET() {
	unsafe.Asm("IRET", nil)
}

// IRETD
// Interrupt return (32-bit operand size).
//
// Documentation: https://golang.org/s/x86manual#page=578
func IRETD() {
	unsafe.Asm("IRETD", nil)
}

// IRETQ
// Interrupt return (64-bit operand size).
//
// Documentation: https://golang.org/s/x86manual#page=578
func IRETQ() {
	unsafe.Asm("IRETQ", nil)
}

// JA
// Jump near if above (CF=0 and ZF=0). Not supported in 64-bit mode.
// Jump near if above (CF=0 and ZF=0).
// Jump short if above (CF=0 and ZF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JA(offset interface{}) {
	unsafe.Asm("JA", offset)
}

// JAE
// Jump near if above or equal (CF=0). Not supported in 64-bit mode.
// Jump near if above or equal (CF=0).
// Jump short if above or equal (CF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JAE(offset interface{}) {
	unsafe.Asm("JAE", offset)
}

// JB
// Jump near if below (CF=1). Not supported in 64-bit mode.
// Jump near if below (CF=1).
// Jump short if below (CF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JB(offset interface{}) {
	unsafe.Asm("JB", offset)
}

// JBE
// Jump near if below or equal (CF=1 or ZF=1). Not supported in 64-bit mode.
// Jump near if below or equal (CF=1 or ZF=1).
// Jump short if below or equal (CF=1 or ZF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JBE(offset interface{}) {
	unsafe.Asm("JBE", offset)
}

// JC
// Jump near if carry (CF=1). Not supported in 64-bit mode.
// Jump near if carry (CF=1).
// Jump short if carry (CF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JC(offset interface{}) {
	unsafe.Asm("JC", offset)
}

// JCXZ
// Jump short if CX register is 0.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JCXZ(offset interface{}) {
	unsafe.Asm("JCXZ", offset)
}

// JE
// Jump near if equal (ZF=1). Not supported in 64-bit mode.
// Jump near if equal (ZF=1).
// Jump short if equal (ZF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JE(offset interface{}) {
	unsafe.Asm("JE", offset)
}

// JECXZ
// Jump short if ECX register is 0.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JECXZ(offset interface{}) {
	unsafe.Asm("JECXZ", offset)
}

// JG
// Jump near if greater (ZF=0 and SF=OF). Not supported in 64-bit mode.
// Jump near if greater (ZF=0 and SF=OF).
// Jump short if greater (ZF=0 and SF=OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JG(offset interface{}) {
	unsafe.Asm("JG", offset)
}

// JGE
// Jump near if greater or equal (SF=OF). Not supported in 64-bit mode.
// Jump near if greater or equal (SF=OF).
// Jump short if greater or equal (SF=OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JGE(offset interface{}) {
	unsafe.Asm("JGE", offset)
}

// JL
// Jump near if less (SF≠ OF). Not supported in 64-bit mode.
// Jump near if less (SF≠ OF).
// Jump short if less (SF≠ OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JL(offset interface{}) {
	unsafe.Asm("JL", offset)
}

// JLE
// Jump near if less or equal (ZF=1 or SF≠ OF). Not supported in 64-bit mode.
// Jump near if less or equal (ZF=1 or SF≠ OF).
// Jump short if less or equal (ZF=1 or SF≠ OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JLE(offset interface{}) {
	unsafe.Asm("JLE", offset)
}

// JMP_D
// Jump near, relative, displacement relative to next instruction. Not supported in 64-bit mode.
// Jump near, relative, RIP = RIP + 32-bit displacement sign extended to 64-bits
// Jump short, RIP = RIP + 8-bit displacement sign extended to 64-bits
// Jump far, absolute indirect, address given in m16:16
// Jump far, absolute indirect, address given in m16:32.
// Jump far, absolute indirect, address given in m16:64.
// Jump far, absolute, address given in operand
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=590
func JMP_D(offset interface{}) {
	unsafe.Asm("JMP", offset)
}

// JMP_M
// Jump near, absolute indirect, address = zero- extended r/m16. Not supported in 64-bit mode.
// Jump near, absolute indirect, address given in r/m32. Not supported in 64-bit mode.
// Jump near, absolute indirect, RIP = 64-Bit offset from register or memory
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=590
func JMP_M(rm interface{}) {
	unsafe.Asm("JMP", rm)
}

// JNA
// Jump near if not above (CF=1 or ZF=1). Not supported in 64-bit mode.
// Jump near if not above (CF=1 or ZF=1).
// Jump short if not above (CF=1 or ZF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNA(offset interface{}) {
	unsafe.Asm("JNA", offset)
}

// JNAE
// Jump near if not above or equal (CF=1). Not supported in 64-bit mode.
// Jump near if not above or equal (CF=1).
// Jump short if not above or equal (CF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNAE(offset interface{}) {
	unsafe.Asm("JNAE", offset)
}

// JNB
// Jump near if not below (CF=0). Not supported in 64-bit mode.
// Jump near if not below (CF=0).
// Jump short if not below (CF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNB(offset interface{}) {
	unsafe.Asm("JNB", offset)
}

// JNBE
// Jump near if not below or equal (CF=0 and ZF=0). Not supported in 64-bit mode.
// Jump near if not below or equal (CF=0 and ZF=0).
// Jump short if not below or equal (CF=0 and ZF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNBE(offset interface{}) {
	unsafe.Asm("JNBE", offset)
}

// JNC
// Jump near if not carry (CF=0). Not supported in 64-bit mode.
// Jump near if not carry (CF=0).
// Jump short if not carry (CF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNC(offset interface{}) {
	unsafe.Asm("JNC", offset)
}

// JNE
// Jump near if not equal (ZF=0). Not supported in 64-bit mode.
// Jump near if not equal (ZF=0).
// Jump short if not equal (ZF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNE(offset interface{}) {
	unsafe.Asm("JNE", offset)
}

// JNG
// Jump near if not greater (ZF=1 or SF≠ OF). Not supported in 64-bit mode.
// Jump near if not greater (ZF=1 or SF≠ OF).
// Jump short if not greater (ZF=1 or SF≠ OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNG(offset interface{}) {
	unsafe.Asm("JNG", offset)
}

// JNGE
// Jump near if not greater or equal (SF≠ OF). Not supported in 64-bit mode.
// Jump near if not greater or equal (SF≠ OF).
// Jump short if not greater or equal (SF≠ OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNGE(offset interface{}) {
	unsafe.Asm("JNGE", offset)
}

// JNL
// Jump near if not less (SF=OF). Not supported in 64-bit mode.
// Jump near if not less (SF=OF).
// Jump short if not less (SF=OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNL(offset interface{}) {
	unsafe.Asm("JNL", offset)
}

// JNLE
// Jump near if not less or equal (ZF=0 and SF=OF). Not supported in 64-bit mode.
// Jump near if not less or equal (ZF=0 and SF=OF).
// Jump short if not less or equal (ZF=0 and SF=OF).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNLE(offset interface{}) {
	unsafe.Asm("JNLE", offset)
}

// JNO
// Jump near if not overflow (OF=0). Not supported in 64-bit mode.
// Jump near if not overflow (OF=0).
// Jump short if not overflow (OF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNO(offset interface{}) {
	unsafe.Asm("JNO", offset)
}

// JNP
// Jump near if not parity (PF=0). Not supported in 64-bit mode.
// Jump near if not parity (PF=0).
// Jump short if not parity (PF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNP(offset interface{}) {
	unsafe.Asm("JNP", offset)
}

// JNS
// Jump near if not sign (SF=0). Not supported in 64-bit mode.
// Jump near if not sign (SF=0).
// Jump short if not sign (SF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNS(offset interface{}) {
	unsafe.Asm("JNS", offset)
}

// JNZ
// Jump near if not zero (ZF=0). Not supported in 64-bit mode.
// Jump near if not zero (ZF=0).
// Jump short if not zero (ZF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JNZ(offset interface{}) {
	unsafe.Asm("JNZ", offset)
}

// JO
// Jump near if overflow (OF=1). Not supported in 64-bit mode.
// Jump near if overflow (OF=1).
// Jump short if overflow (OF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JO(offset interface{}) {
	unsafe.Asm("JO", offset)
}

// JP
// Jump near if parity (PF=1). Not supported in 64-bit mode.
// Jump near if parity (PF=1).
// Jump short if parity (PF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JP(offset interface{}) {
	unsafe.Asm("JP", offset)
}

// JPE
// Jump near if parity even (PF=1). Not supported in 64-bit mode.
// Jump near if parity even (PF=1).
// Jump short if parity even (PF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JPE(offset interface{}) {
	unsafe.Asm("JPE", offset)
}

// JPO
// Jump near if parity odd (PF=0). Not supported in 64-bit mode.
// Jump near if parity odd (PF=0).
// Jump short if parity odd (PF=0).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JPO(offset interface{}) {
	unsafe.Asm("JPO", offset)
}

// JRCXZ
// Jump short if RCX register is 0.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JRCXZ(offset interface{}) {
	unsafe.Asm("JRCXZ", offset)
}

// JS
// Jump near if sign (SF=1). Not supported in 64- bit mode.
// Jump near if sign (SF=1).
// Jump short if sign (SF=1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JS(offset interface{}) {
	unsafe.Asm("JS", offset)
}

// JZ
// Jump near if 0 (ZF=1). Not supported in 64-bit mode.
// Jump near if 0 (ZF=1).
// Jump short if zero (ZF = 1).
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=585
func JZ(offset interface{}) {
	unsafe.Asm("JZ", offset)
}

// KADDB
// Add 8 bits masks in k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=598
func KADDB(reg, vex interface{}) {
	unsafe.Asm("KADDB", reg, vex)
}

// KADDD
// Add 32 bits masks in k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=598
func KADDD(reg, vex interface{}) {
	unsafe.Asm("KADDD", reg, vex)
}

// KADDQ
// Add 64 bits masks in k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=598
func KADDQ(reg, vex interface{}) {
	unsafe.Asm("KADDQ", reg, vex)
}

// KADDW
// Add 16 bits masks in k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=598
func KADDW(reg, vex interface{}) {
	unsafe.Asm("KADDW", reg, vex)
}

// KANDB
// Bitwise AND 8 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=599
func KANDB(reg, vex interface{}) {
	unsafe.Asm("KANDB", reg, vex)
}

// KANDD
// Bitwise AND 32 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=599
func KANDD(reg, vex interface{}) {
	unsafe.Asm("KANDD", reg, vex)
}

// KANDNB
// Bitwise AND NOT 8 bits masks k1 and k2 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=600
func KANDNB(reg, vex interface{}) {
	unsafe.Asm("KANDNB", reg, vex)
}

// KANDND
// Bitwise AND NOT 32 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=600
func KANDND(reg, vex interface{}) {
	unsafe.Asm("KANDND", reg, vex)
}

// KANDNQ
// Bitwise AND NOT 64 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=600
func KANDNQ(reg, vex interface{}) {
	unsafe.Asm("KANDNQ", reg, vex)
}

// KANDNW
// Bitwise AND NOT 16 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=600
func KANDNW(reg, vex interface{}) {
	unsafe.Asm("KANDNW", reg, vex)
}

// KANDQ
// Bitwise AND 64 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=599
func KANDQ(reg, vex interface{}) {
	unsafe.Asm("KANDQ", reg, vex)
}

// KANDW
// Bitwise AND 16 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=599
func KANDW(reg, vex interface{}) {
	unsafe.Asm("KANDW", reg, vex)
}

// KMOVB_MR
// Move 8 bits mask from k1 and store the result in m8.
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVB_MR() {
	unsafe.Asm("KMOVB", nil)
}

// KMOVB_RM
// Move 8 bits mask from k2/m8 and store the result in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVB_RM(reg, rm interface{}) {
	unsafe.Asm("KMOVB", reg, rm)
}

// KMOVB_RR
// Move 8 bits mask from r32 to k1.
// Move 8 bits mask from k1 to r32.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVB_RR(reg interface{}) {
	unsafe.Asm("KMOVB", reg)
}

// KMOVD_MR
// Move 32 bits mask from k1 and store the result in m32.
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVD_MR() {
	unsafe.Asm("KMOVD", nil)
}

// KMOVD_RM
// Move 32 bits mask from k2/m32 and store the result in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVD_RM(reg, rm interface{}) {
	unsafe.Asm("KMOVD", reg, rm)
}

// KMOVD_RR
// Move 32 bits mask from r32 to k1.
// Move 32 bits mask from k1 to r32.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVD_RR(reg interface{}) {
	unsafe.Asm("KMOVD", reg)
}

// KMOVQ_MR
// Move 64 bits mask from k1 and store the result in m64.
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVQ_MR() {
	unsafe.Asm("KMOVQ", nil)
}

// KMOVQ_RM
// Move 64 bits mask from k2/m64 and store the result in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVQ_RM(reg, rm interface{}) {
	unsafe.Asm("KMOVQ", reg, rm)
}

// KMOVQ_RR
// Move 64 bits mask from r64 to k1.
// Move 64 bits mask from k1 to r64.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVQ_RR(reg interface{}) {
	unsafe.Asm("KMOVQ", reg)
}

// KMOVW_MR
// Move 16 bits mask from k1 and store the result in m16.
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVW_MR() {
	unsafe.Asm("KMOVW", nil)
}

// KMOVW_RM
// Move 16 bits mask from k2/m16 and store the result in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVW_RM(reg, rm interface{}) {
	unsafe.Asm("KMOVW", reg, rm)
}

// KMOVW_RR
// Move 16 bits mask from r32 to k1.
// Move 16 bits mask from k1 to r32.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=601
func KMOVW_RR(reg interface{}) {
	unsafe.Asm("KMOVW", reg)
}

// KNOTB
// Bitwise NOT of 8 bits mask k2.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=603
func KNOTB(reg interface{}) {
	unsafe.Asm("KNOTB", reg)
}

// KNOTD
// Bitwise NOT of 32 bits mask k2.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=603
func KNOTD(reg interface{}) {
	unsafe.Asm("KNOTD", reg)
}

// KNOTQ
// Bitwise NOT of 64 bits mask k2.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=603
func KNOTQ(reg interface{}) {
	unsafe.Asm("KNOTQ", reg)
}

// KNOTW
// Bitwise NOT of 16 bits mask k2.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=603
func KNOTW(reg interface{}) {
	unsafe.Asm("KNOTW", reg)
}

// KORB
// Bitwise OR 8 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=604
func KORB(reg, vex interface{}) {
	unsafe.Asm("KORB", reg, vex)
}

// KORD
// Bitwise OR 32 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=604
func KORD(reg, vex interface{}) {
	unsafe.Asm("KORD", reg, vex)
}

// KORQ
// Bitwise OR 64 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=604
func KORQ(reg, vex interface{}) {
	unsafe.Asm("KORQ", reg, vex)
}

// KORTESTB
// Bitwise OR 8 bits masks k1 and k2 and update ZF and CF accordingly.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=605
func KORTESTB(reg interface{}) {
	unsafe.Asm("KORTESTB", reg)
}

// KORTESTD
// Bitwise OR 32 bits masks k1 and k2 and update ZF and CF accordingly.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=605
func KORTESTD(reg interface{}) {
	unsafe.Asm("KORTESTD", reg)
}

// KORTESTQ
// Bitwise OR 64 bits masks k1 and k2 and update ZF and CF accordingly.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=605
func KORTESTQ(reg interface{}) {
	unsafe.Asm("KORTESTQ", reg)
}

// KORTESTW
// Bitwise OR 16 bits masks k1 and k2 and update ZF and CF accordingly.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=605
func KORTESTW(reg interface{}) {
	unsafe.Asm("KORTESTW", reg)
}

// KORW
// Bitwise OR 16 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=604
func KORW(reg, vex interface{}) {
	unsafe.Asm("KORW", reg, vex)
}

// KSHIFTLB
// Shift left 8 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=607
func KSHIFTLB(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTLB", reg, arg, imm)
}

// KSHIFTLD
// Shift left 32 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=607
func KSHIFTLD(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTLD", reg, arg, imm)
}

// KSHIFTLQ
// Shift left 64 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=607
func KSHIFTLQ(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTLQ", reg, arg, imm)
}

// KSHIFTLW
// Shift left 16 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=607
func KSHIFTLW(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTLW", reg, arg, imm)
}

// KSHIFTRB
// Shift right 8 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=609
func KSHIFTRB(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTRB", reg, arg, imm)
}

// KSHIFTRD
// Shift right 32 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=609
func KSHIFTRD(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTRD", reg, arg, imm)
}

// KSHIFTRQ
// Shift right 64 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=609
func KSHIFTRQ(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTRQ", reg, arg, imm)
}

// KSHIFTRW
// Shift right 16 bits in k2 by immediate and write result in k1.
//
// reg: ModRM:reg (w) ModRM:r/m (r, ModRM:[7:6] must be 11b)
// arg:
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=609
func KSHIFTRW(reg, arg, imm interface{}) {
	unsafe.Asm("KSHIFTRW", reg, arg, imm)
}

// KUNPCKBW
// Unpack and interleave 8 bits masks in k2 and k3 and write word result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=613
func KUNPCKBW(reg, vex interface{}) {
	unsafe.Asm("KUNPCKBW", reg, vex)
}

// KUNPCKDQ
// Unpack and interleave 32 bits masks in k2 and k3 and write quadword result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=613
func KUNPCKDQ(reg, vex interface{}) {
	unsafe.Asm("KUNPCKDQ", reg, vex)
}

// KUNPCKWD
// Unpack and interleave 16 bits in k2 and k3 and write double- word result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=613
func KUNPCKWD(reg, vex interface{}) {
	unsafe.Asm("KUNPCKWD", reg, vex)
}

// KXNORB
// Bitwise XNOR 8 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=614
func KXNORB(reg, vex interface{}) {
	unsafe.Asm("KXNORB", reg, vex)
}

// KXNORD
// Bitwise XNOR 32 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=614
func KXNORD(reg, vex interface{}) {
	unsafe.Asm("KXNORD", reg, vex)
}

// KXNORQ
// Bitwise XNOR 64 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=614
func KXNORQ(reg, vex interface{}) {
	unsafe.Asm("KXNORQ", reg, vex)
}

// KXNORW
// Bitwise XNOR 16 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=614
func KXNORW(reg, vex interface{}) {
	unsafe.Asm("KXNORW", reg, vex)
}

// KXORB
// Bitwise XOR 8 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=615
func KXORB(reg, vex interface{}) {
	unsafe.Asm("KXORB", reg, vex)
}

// KXORD
// Bitwise XOR 32 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=615
func KXORD(reg, vex interface{}) {
	unsafe.Asm("KXORD", reg, vex)
}

// KXORQ
// Bitwise XOR 64 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=615
func KXORQ(reg, vex interface{}) {
	unsafe.Asm("KXORQ", reg, vex)
}

// KXORW
// Bitwise XOR 16 bits masks k2 and k3 and place result in k1.
//
// reg: ModRM:reg (w)
// vex: VEX.1vvv (r) ModRM:r/m (r, ModRM:[7:6] must be 11b)
//
// Documentation: https://golang.org/s/x86manual#page=615
func KXORW(reg, vex interface{}) {
	unsafe.Asm("KXORW", reg, vex)
}

// LAHF
// Load: AH ← EFLAGS(SF:ZF:0:AF:0:PF:1:CF).
//
// Documentation: https://golang.org/s/x86manual#page=616
func LAHF() {
	unsafe.Asm("LAHF", nil)
}

// LAR
// r16 ← access rights referenced by r16/m16
// reg ← access rights referenced by r32/m16
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=617
func LAR(reg, rm interface{}) {
	unsafe.Asm("LAR", reg, rm)
}

// LDDQU
// Load unaligned data from mem and return double quadword in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=620
func LDDQU(reg, rm interface{}) {
	unsafe.Asm("LDDQU", reg, rm)
}

// LDMXCSR
// Load MXCSR register from m32.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=622
func LDMXCSR(rm interface{}) {
	unsafe.Asm("LDMXCSR", rm)
}

// LDS
// Load DS:r16 with far pointer from memory.
// Load DS:r32 with far pointer from memory.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=623
func LDS(reg, rm interface{}) {
	unsafe.Asm("LDS", reg, rm)
}

// LEA
// Store effective address for m in register r16.
// Store effective address for m in register r32.
// Store effective address for m in register r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=627
func LEA(reg, rm interface{}) {
	unsafe.Asm("LEA", reg, rm)
}

// LEAVE
// Set ESP to EBP, then pop EBP.
// Set RSP to RBP, then pop RBP.
// Set SP to BP, then pop BP.
//
// Documentation: https://golang.org/s/x86manual#page=629
func LEAVE() {
	unsafe.Asm("LEAVE", nil)
}

// LES
// Load ES:r16 with far pointer from memory.
// Load ES:r32 with far pointer from memory.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=623
func LES(reg, rm interface{}) {
	unsafe.Asm("LES", reg, rm)
}

// LFENCE
// Serializes load operations.
//
// Documentation: https://golang.org/s/x86manual#page=631
func LFENCE() {
	unsafe.Asm("LFENCE", nil)
}

// LFS
// Load FS:r16 with far pointer from memory.
// Load FS:r32 with far pointer from memory.
// Load FS:r64 with far pointer from memory.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=623
func LFS(reg, rm interface{}) {
	unsafe.Asm("LFS", reg, rm)
}

// LGDT
// Load m into GDTR.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=632
func LGDT(rm interface{}) {
	unsafe.Asm("LGDT", rm)
}

// LGS
// Load GS:r16 with far pointer from memory.
// Load GS:r32 with far pointer from memory.
// Load GS:r64 with far pointer from memory.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=623
func LGS(reg, rm interface{}) {
	unsafe.Asm("LGS", reg, rm)
}

// LIDT
// Load m into IDTR.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=632
func LIDT(rm interface{}) {
	unsafe.Asm("LIDT", rm)
}

// LLDT
// Load segment selector r/m16 into LDTR.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=635
func LLDT(rm interface{}) {
	unsafe.Asm("LLDT", rm)
}

// LMSW
// Loads r/m16 in machine status word of CR0.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=637
func LMSW(rm interface{}) {
	unsafe.Asm("LMSW", rm)
}

// LOCK
// Asserts LOCK# signal for duration of the accompanying instruction.
//
// Documentation: https://golang.org/s/x86manual#page=639
func LOCK() {
	unsafe.Asm("LOCK", nil)
}

// LODSB
// For legacy mode, Load byte at address DS:(E)SI into AL. For 64-bit mode load byte at address (R)SI into AL.
//
// Documentation: https://golang.org/s/x86manual#page=641
func LODSB() {
	unsafe.Asm("LODSB", nil)
}

// LODSD
// For legacy mode, Load dword at address DS:(E)SI into EAX. For 64-bit mode load dword at address (R)SI into EAX.
//
// Documentation: https://golang.org/s/x86manual#page=641
func LODSD() {
	unsafe.Asm("LODSD", nil)
}

// LODSQ
// Load qword at address (R)SI into RAX.
//
// Documentation: https://golang.org/s/x86manual#page=641
func LODSQ() {
	unsafe.Asm("LODSQ", nil)
}

// LODSW
// For legacy mode, Load word at address DS:(E)SI into AX. For 64-bit mode load word at address (R)SI into AX.
//
// Documentation: https://golang.org/s/x86manual#page=641
func LODSW() {
	unsafe.Asm("LODSW", nil)
}

// LOOP
// Decrement count; jump short if count ≠ 0.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=644
func LOOP(offset interface{}) {
	unsafe.Asm("LOOP", offset)
}

// LOOPE
// Decrement count; jump short if count ≠ 0 and ZF = 1.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=644
func LOOPE(offset interface{}) {
	unsafe.Asm("LOOPE", offset)
}

// LOOPNE
// Decrement count; jump short if count ≠ 0 and ZF = 0.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=644
func LOOPNE(offset interface{}) {
	unsafe.Asm("LOOPNE", offset)
}

// LSL
// Load: r16 ← segment limit, selector r16/m16.
// Load: r32 ← segment limit, selector r32/m16.
// Load: r64 ← segment limit, selector r32/m16
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=646
func LSL(reg, rm interface{}) {
	unsafe.Asm("LSL", reg, rm)
}

// LSS
// Load SS:r16 with far pointer from memory.
// Load SS:r32 with far pointer from memory.
// Load SS:r64 with far pointer from memory.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=623
func LSS(reg, rm interface{}) {
	unsafe.Asm("LSS", reg, rm)
}

// LTR
// Load r/m16 into task register.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=649
func LTR(rm interface{}) {
	unsafe.Asm("LTR", rm)
}

// LZCNT
// Count the number of leading zero bits in r/m16, return result in r16.
// Count the number of leading zero bits in r/m32, return result in r32.
// Count the number of leading zero bits in r/m64, return result in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=651
func LZCNT(reg, rm interface{}) {
	unsafe.Asm("LZCNT", reg, rm)
}

// MASKMOVDQU
// Selectively write bytes from xmm1 to memory location using the byte mask in xmm2. The default memory location is specified by DS:DI/EDI/RDI.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=660
func MASKMOVDQU(reg, rm interface{}) {
	unsafe.Asm("MASKMOVDQU", reg, rm)
}

// MASKMOVQ
// Selectively write bytes from mm1 to memory location using the byte mask in mm2. The default memory location is specified by DS:DI/EDI/RDI.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=662
func MASKMOVQ(reg, rm interface{}) {
	unsafe.Asm("MASKMOVQ", reg, rm)
}

// MAXPD
// Return the maximum double-precision floating-point values between xmm1 and xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=664
func MAXPD(reg, rm interface{}) {
	unsafe.Asm("MAXPD", reg, rm)
}

// MAXPS
// Return the maximum single-precision floating-point values between xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=667
func MAXPS(reg, rm interface{}) {
	unsafe.Asm("MAXPS", reg, rm)
}

// MAXSD
// Return the maximum scalar double-precision floating-point value between xmm2/m64 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=670
func MAXSD(reg, rm interface{}) {
	unsafe.Asm("MAXSD", reg, rm)
}

// MAXSS
// Return the maximum scalar single-precision floating-point value between xmm2/m32 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=672
func MAXSS(reg, rm interface{}) {
	unsafe.Asm("MAXSS", reg, rm)
}

// MFENCE
// Serializes load and store operations.
//
// Documentation: https://golang.org/s/x86manual#page=674
func MFENCE() {
	unsafe.Asm("MFENCE", nil)
}

// MINPD
// Return the minimum double-precision floating-point values between xmm1 and xmm2/mem
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=675
func MINPD(reg, rm interface{}) {
	unsafe.Asm("MINPD", reg, rm)
}

// MINPS
// Return the minimum single-precision floating-point values between xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=678
func MINPS(reg, rm interface{}) {
	unsafe.Asm("MINPS", reg, rm)
}

// MINSD
// Return the minimum scalar double-precision floating- point value between xmm2/m64 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=681
func MINSD(reg, rm interface{}) {
	unsafe.Asm("MINSD", reg, rm)
}

// MINSS
// Return the minimum scalar single-precision floating- point value between xmm2/m32 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=683
func MINSS(reg, rm interface{}) {
	unsafe.Asm("MINSS", reg, rm)
}

// MONITOR
// Sets up a linear address range to be monitored by hardware and activates the monitor. The address range should be a write- back memory caching type. The address is DS:EAX (DS:RAX in 64-bit mode).
//
// Documentation: https://golang.org/s/x86manual#page=685
func MONITOR() {
	unsafe.Asm("MONITOR", nil)
}

// MOV_FD
// Move byte at (seg:offset) to AL.
// Move byte at (offset) to AL.
// Move word at (seg:offset) to AX.
// Move doubleword at (seg:offset) to EAX.
// Move quadword at (offset) to RAX.
//
// al: AL/AX/EAX/RAX
// moffs: Moffs
//
// Documentation: https://golang.org/s/x86manual#page=687
func MOV_FD(al, moffs interface{}) {
	unsafe.Asm("MOV", al, moffs)
}

// MOV_MI
// Move imm16 to r/m16.
// Move imm32 to r/m32.
// Move imm32 sign extended to 64-bits to r/m64.
// Move imm8 to r/m8.
//
// rm: ModRM:r/m (w)
// imm: imm8/16/32/64
//
// Documentation: https://golang.org/s/x86manual#page=687
func MOV_MI(rm, imm interface{}) {
	unsafe.Asm("MOV", rm, imm)
}

// MOV_MR
// Move segment register to r/m16.
// Move zero extended 16-bit segment register to r/m64.
// Move r16 to r/m16.
// Move r32 to r/m32.
// Move r64 to r/m64.
// Move r8 to r/m8.
// Move control register to r32.
// Move debug register to r32.
// Move extended control register to r64.
// Move extended CR8 to r64.
// Move extended debug register to r64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=687
func MOV_MR(rm, reg interface{}) {
	unsafe.Asm("MOV", rm, reg)
}

// MOV_OI
// Move imm16 to r16.
// Move imm32 to r32.
// Move imm64 to r64.
// Move imm8 to r8.
//
// opcode: opcode + rd (w)
// imm: imm8/16/32/64
//
// Documentation: https://golang.org/s/x86manual#page=687
func MOV_OI(opcode, imm interface{}) {
	unsafe.Asm("MOV", opcode, imm)
}

// MOV_RM
// Move r32 to control register1.
// Move r64 to extended control register.
// Move r64 to extended CR8.
// Move r32 to debug register.
// Move r64 to extended debug register.
// Move r/m16 to segment register.
// Move lower 16 bits of r/m64 to segment register.
// Move r/m16 to r16.
// Move r/m32 to r32.
// Move r/m64 to r64.
// Move r/m8 to r8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=692
func MOV_RM(reg, rm interface{}) {
	unsafe.Asm("MOV", reg, rm)
}

// MOV_TD
// Move AX to (seg:offset).
// Move EAX to (seg:offset).
// Move RAX to (offset).
// Move AL to (seg:offset).
// Move AL to (offset).
//
// moffs: Moffs (w)
// al: AL/AX/EAX/RAX
//
// Documentation: https://golang.org/s/x86manual#page=687
func MOV_TD(moffs, al interface{}) {
	unsafe.Asm("MOV", moffs, al)
}

// MOVAPD_MR
// Move aligned packed double-precision floating- point values from xmm1 to xmm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func MOVAPD_MR(rm, reg interface{}) {
	unsafe.Asm("MOVAPD", rm, reg)
}

// MOVAPD_RM
// Move aligned packed double-precision floating- point values from xmm2/mem to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func MOVAPD_RM(reg, rm interface{}) {
	unsafe.Asm("MOVAPD", reg, rm)
}

// MOVAPS_MR
// Move aligned packed single-precision floating-point values from xmm1 to xmm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func MOVAPS_MR(rm, reg interface{}) {
	unsafe.Asm("MOVAPS", rm, reg)
}

// MOVAPS_RM
// Move aligned packed single-precision floating-point values from xmm2/mem to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func MOVAPS_RM(reg, rm interface{}) {
	unsafe.Asm("MOVAPS", reg, rm)
}

// MOVBE_MR
// Reverse byte order in r16 and move to m16.
// Reverse byte order in r32 and move to m32.
// Reverse byte order in r64 and move to m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=705
func MOVBE_MR(rm, reg interface{}) {
	unsafe.Asm("MOVBE", rm, reg)
}

// MOVBE_RM
// Reverse byte order in m16 and move to r16.
// Reverse byte order in m32 and move to r32.
// Reverse byte order in m64 and move to r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=705
func MOVBE_RM(reg, rm interface{}) {
	unsafe.Asm("MOVBE", reg, rm)
}

// MOVD_MR
// Move doubleword from mm to r/m32.
// Move doubleword from xmm register to r/m32.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=707
func MOVD_MR(rm, reg interface{}) {
	unsafe.Asm("MOVD", rm, reg)
}

// MOVD_RM
// Move doubleword from r/m32 to mm.
// Move doubleword from r/m32 to xmm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=707
func MOVD_RM(reg, rm interface{}) {
	unsafe.Asm("MOVD", reg, rm)
}

// MOVD_RVM
// Conditionally load dword values from m128 using mask in xmm2 and store in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1753
func MOVD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("MOVD", reg, vex, rm)
}

// MOVDQ2Q
// Move low quadword from xmm to mmx register.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=727
func MOVDQ2Q(reg, rm interface{}) {
	unsafe.Asm("MOVDQ2Q", reg, rm)
}

// MOVDQA_MR
// Move aligned packed integer values from xmm1 to xmm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func MOVDQA_MR(rm, reg interface{}) {
	unsafe.Asm("MOVDQA", rm, reg)
}

// MOVDQA_RM
// Move aligned packed integer values from xmm2/mem to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func MOVDQA_RM(reg, rm interface{}) {
	unsafe.Asm("MOVDQA", reg, rm)
}

// MOVHLPS
// Move two packed single-precision floating-point values from high quadword of xmm2 to low quadword of xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=728
func MOVHLPS(reg, rm interface{}) {
	unsafe.Asm("MOVHLPS", reg, rm)
}

// MOVLHPS
// Move two packed single-precision floating-point values from low quadword of xmm2 to high quadword of xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=734
func MOVLHPS(reg, rm interface{}) {
	unsafe.Asm("MOVLHPS", reg, rm)
}

// MOVMSKPD
// Extract 2-bit sign mask from xmm and store in reg. The upper bits of r32 or r64 are filled with zeros.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=740
func MOVMSKPD(reg, rm interface{}) {
	unsafe.Asm("MOVMSKPD", reg, rm)
}

// MOVMSKPS
// Extract 4-bit sign mask from xmm and store in reg. The upper bits of r32 or r64 are filled with zeros.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=742
func MOVMSKPS(reg, rm interface{}) {
	unsafe.Asm("MOVMSKPS", reg, rm)
}

// MOVNTDQ
// Move packed integer values in xmm1 to m128 using non- temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=746
func MOVNTDQ(rm, reg interface{}) {
	unsafe.Asm("MOVNTDQ", rm, reg)
}

// MOVNTDQA
// Move double quadword from m128 to xmm1 using non- temporal hint if WC memory type.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=744
func MOVNTDQA(reg, rm interface{}) {
	unsafe.Asm("MOVNTDQA", reg, rm)
}

// MOVNTI
// Move doubleword from r32 to m32 using non- temporal hint.
// Move quadword from r64 to m64 using non- temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=748
func MOVNTI(rm, reg interface{}) {
	unsafe.Asm("MOVNTI", rm, reg)
}

// MOVNTPD
// Move packed double-precision values in xmm1 to m128 using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=750
func MOVNTPD(rm, reg interface{}) {
	unsafe.Asm("MOVNTPD", rm, reg)
}

// MOVNTPS
// Move packed single-precision values xmm1 to mem using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=752
func MOVNTPS(rm, reg interface{}) {
	unsafe.Asm("MOVNTPS", rm, reg)
}

// MOVNTQ
// Move quadword from mm to m64 using non- temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=754
func MOVNTQ(rm, reg interface{}) {
	unsafe.Asm("MOVNTQ", rm, reg)
}

// MOVQ_MR
// Move quadword from mm to mm/m64.
// Move quadword from mm to r/m64.
// Move quadword from xmm register to r/m64.
// Move quadword from xmm1 to xmm2/mem64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func MOVQ_MR(rm, reg interface{}) {
	unsafe.Asm("MOVQ", rm, reg)
}

// MOVQ_RM
// Move quadword from mm/m64 to mm.
// Move quadword from r/m64 to mm.
// Move quadword from r/m64 to xmm.
// Move quadword from xmm2/mem64 to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func MOVQ_RM(reg, rm interface{}) {
	unsafe.Asm("MOVQ", reg, rm)
}

// MOVQ2DQ
// Move quadword from mmx to low quadword of xmm.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=758
func MOVQ2DQ(reg, rm interface{}) {
	unsafe.Asm("MOVQ2DQ", reg, rm)
}

// MOVSB
// For legacy mode, Move byte from address DS:(E)SI to ES:(E)DI. For 64-bit mode move byte from address (R|E)SI to (R|E)DI.
//
// Documentation: https://golang.org/s/x86manual#page=759
func MOVSB() {
	unsafe.Asm("MOVSB", nil)
}

// MOVSD
// For legacy mode, move dword from address DS:(E)SI to ES:(E)DI. For 64-bit mode move dword from address (R|E)SI to (R|E)DI.
//
// Documentation: https://golang.org/s/x86manual#page=759
func MOVSD() {
	unsafe.Asm("MOVSD", nil)
}

// MOVSHDUP
// Move odd index single-precision floating-point values from xmm2/mem and duplicate each element into xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=766
func MOVSHDUP(reg, rm interface{}) {
	unsafe.Asm("MOVSHDUP", reg, rm)
}

// MOVSQ
// Move qword from address (R|E)SI to (R|E)DI.
//
// Documentation: https://golang.org/s/x86manual#page=759
func MOVSQ() {
	unsafe.Asm("MOVSQ", nil)
}

// MOVSW
// For legacy mode, move word from address DS:(E)SI to ES:(E)DI. For 64-bit mode move word at address (R|E)SI to (R|E)DI.
//
// Documentation: https://golang.org/s/x86manual#page=759
func MOVSW() {
	unsafe.Asm("MOVSW", nil)
}

// MOVSX
// Move byte to word with sign-extension.
// Move word to doubleword, with sign- extension.
// Move byte to doubleword with sign- extension.
// Move word to quadword with sign-extension.
// Move byte to quadword with sign-extension.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=776
func MOVSX(reg, rm interface{}) {
	unsafe.Asm("MOVSX", reg, rm)
}

// MOVSXD
// Move doubleword to quadword with sign- extension.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=776
func MOVSXD(reg, rm interface{}) {
	unsafe.Asm("MOVSXD", reg, rm)
}

// MOVZX
// Move byte to word with zero-extension.
// Move word to doubleword, zero-extension.
// Move byte to doubleword, zero-extension.
// Move word to quadword, zero-extension.
// Move byte to quadword, zero-extension.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=786
func MOVZX(reg, rm interface{}) {
	unsafe.Asm("MOVZX", reg, rm)
}

// MPSADBW
// Sums absolute 8-bit integer difference of adjacent groups of 4 byte integers in xmm1 and xmm2/m128 and writes the results in xmm1. Starting offsets within xmm1 and xmm2/m128 are determined by imm8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=788
func MPSADBW(reg, rm, imm interface{}) {
	unsafe.Asm("MPSADBW", reg, rm, imm)
}

// MUL
// Unsigned multiply (DX:AX ← AX ∗ r/m16).
// Unsigned multiply (EDX:EAX ← EAX ∗ r/m32).
// Unsigned multiply (RDX:RAX ← RAX ∗ r/m64).
// Unsigned multiply (AX ← AL ∗ r/m8).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=796
func MUL(rm interface{}) {
	unsafe.Asm("MUL", rm)
}

// MULPD
// Multiply packed double-precision floating-point values in xmm2/m128 with xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=798
func MULPD(reg, rm interface{}) {
	unsafe.Asm("MULPD", reg, rm)
}

// MULPS
// Multiply packed single-precision floating-point values in xmm2/m128 with xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=801
func MULPS(reg, rm interface{}) {
	unsafe.Asm("MULPS", reg, rm)
}

// MULSD
// Multiply the low double-precision floating-point value in xmm2/m64 by low double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=804
func MULSD(reg, rm interface{}) {
	unsafe.Asm("MULSD", reg, rm)
}

// MULSS
// Multiply the low single-precision floating-point value in xmm2/m32 by the low single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=806
func MULSS(reg, rm interface{}) {
	unsafe.Asm("MULSS", reg, rm)
}

// MULX
// Unsigned multiply of r/m32 with EDX without affecting arithmetic flags.
// Unsigned multiply of r/m64 with RDX without affecting arithmetic flags.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=808
func MULX(reg, vex, rm interface{}) {
	unsafe.Asm("MULX", reg, vex, rm)
}

// MWAIT
// A hint that allow the processor to stop instruction execution and enter an implementation-dependent optimized state until occurrence of a class of events.
//
// Documentation: https://golang.org/s/x86manual#page=810
func MWAIT() {
	unsafe.Asm("MWAIT", nil)
}

// NEG
// Two's complement negate r/m16.
// Two's complement negate r/m32.
// Two's complement negate r/m64.
// Two's complement negate r/m8.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=813
func NEG(rm interface{}) {
	unsafe.Asm("NEG", rm)
}

// NOP_M
// Multi-byte no-operation instruction.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=815
func NOP_M(rm interface{}) {
	unsafe.Asm("NOP", rm)
}

// NOP_NP
// One byte no-operation instruction.
//
// Documentation: https://golang.org/s/x86manual#page=815
func NOP_NP() {
	unsafe.Asm("NOP", nil)
}

// NOT
// Reverse each bit of r/m16.
// Reverse each bit of r/m32.
// Reverse each bit of r/m64.
// Reverse each bit of r/m8.
//
// rm: ModRM:r/m (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=816
func NOT(rm interface{}) {
	unsafe.Asm("NOT", rm)
}

// OR_I
// AL OR imm8.
// AX OR imm16.
// EAX OR imm32.
// RAX OR imm32 (sign-extended).
//
// al: AL/AX/EAX/RAX
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=818
func OR_I(al, imm interface{}) {
	unsafe.Asm("OR", al, imm)
}

// OR_MI
// r/m16 OR imm16.
// r/m16 OR imm8 (sign-extended).
// r/m32 OR imm32.
// r/m32 OR imm8 (sign-extended).
// r/m64 OR imm32 (sign-extended).
// r/m64 OR imm8 (sign-extended).
// r/m8 OR imm8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=818
func OR_MI(rm, imm interface{}) {
	unsafe.Asm("OR", rm, imm)
}

// OR_MR
// r/m16 OR r16.
// r/m32 OR r32.
// r/m64 OR r64.
// r/m8 OR r8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=818
func OR_MR(rm, reg interface{}) {
	unsafe.Asm("OR", rm, reg)
}

// OR_RM
// r16 OR r/m16.
// r32 OR r/m32.
// r64 OR r/m64.
// r8 OR r/m8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=818
func OR_RM(reg, rm interface{}) {
	unsafe.Asm("OR", reg, rm)
}

// ORPD
// Return the bitwise logical OR of packed double-precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=820
func ORPD(reg, rm interface{}) {
	unsafe.Asm("ORPD", reg, rm)
}

// ORPS
// Return the bitwise logical OR of packed single-precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=823
func ORPS(reg, rm interface{}) {
	unsafe.Asm("ORPS", reg, rm)
}

// OUT_I
// Output byte in AL to I/O port address imm8.
// Output word in AX to I/O port address imm8.
// Output doubleword in EAX to I/O port address imm8.
//
// imm: imm8 (r)
// al: AL (r)
//
// Documentation: https://golang.org/s/x86manual#page=826
func OUT_I(imm, al interface{}) {
	unsafe.Asm("OUT", imm, al)
}

// OUT_NP
// Output byte in AL to I/O port address in DX.
// Output word in AX to I/O port address in DX.
// Output doubleword in EAX to I/O port address in DX.
//
// dx: DX (r)
// al: AL (r)
//
// Documentation: https://golang.org/s/x86manual#page=826
func OUT_NP(dx, al interface{}) {
	unsafe.Asm("OUT", dx, al)
}

// OUTSB
// Output byte from memory location specified in DS:(E)SI or RSI to I/O port specified in DX**.
//
// Documentation: https://golang.org/s/x86manual#page=828
func OUTSB() {
	unsafe.Asm("OUTSB", nil)
}

// OUTSD
// Output doubleword from memory location specified in DS:(E)SI or RSI to I/O port specified in DX**.
//
// Documentation: https://golang.org/s/x86manual#page=828
func OUTSD() {
	unsafe.Asm("OUTSD", nil)
}

// OUTSW
// Output word from memory location specified in DS:(E)SI or RSI to I/O port specified in DX**.
//
// Documentation: https://golang.org/s/x86manual#page=828
func OUTSW() {
	unsafe.Asm("OUTSW", nil)
}

// PABSB
// Compute the absolute value of bytes in mm2/m64 and store UNSIGNED result in mm1.
// Compute the absolute value of bytes in xmm2/m128 and store UNSIGNED result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func PABSB(reg, rm interface{}) {
	unsafe.Asm("PABSB", reg, rm)
}

// PABSD
// Compute the absolute value of 32-bit integers in mm2/m64 and store UNSIGNED result in mm1.
// Compute the absolute value of 32-bit integers in xmm2/m128 and store UNSIGNED result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func PABSD(reg, rm interface{}) {
	unsafe.Asm("PABSD", reg, rm)
}

// PABSW
// Compute the absolute value of 16-bit integers in mm2/m64 and store UNSIGNED result in mm1.
// Compute the absolute value of 16-bit integers in xmm2/m128 and store UNSIGNED result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func PABSW(reg, rm interface{}) {
	unsafe.Asm("PABSW", reg, rm)
}

// PACKSSDW
// Converts 2 packed signed doubleword integers from mm1 and from mm2/m64 into 4 packed signed word integers in mm1 using signed saturation.
// Converts 4 packed signed doubleword integers from xmm1 and from xxm2/m128 into 8 packed signed word integers in xxm1 using signed saturation.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func PACKSSDW(reg, rm interface{}) {
	unsafe.Asm("PACKSSDW", reg, rm)
}

// PACKSSWB
// Converts 4 packed signed word integers from mm1 and from mm2/m64 into 8 packed signed byte integers in mm1 using signed saturation.
// Converts 8 packed signed word integers from xmm1 and from xxm2/m128 into 16 packed signed byte integers in xxm1 using signed saturation.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func PACKSSWB(reg, rm interface{}) {
	unsafe.Asm("PACKSSWB", reg, rm)
}

// PACKUSDW
// Convert 4 packed signed doubleword integers from xmm1 and 4 packed signed doubleword integers from xmm2/m128 into 8 packed unsigned word integers in xmm1 using unsigned saturation.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=846
func PACKUSDW(reg, rm interface{}) {
	unsafe.Asm("PACKUSDW", reg, rm)
}

// PACKUSWB
// Converts 4 signed word integers from mm and 4 signed word integers from mm/m64 into 8 unsigned byte integers in mm using unsigned saturation.
// Converts 8 signed word integers from xmm1 and 8 signed word integers from xmm2/m128 into 16 unsigned byte integers in xmm1 using unsigned saturation.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=851
func PACKUSWB(reg, rm interface{}) {
	unsafe.Asm("PACKUSWB", reg, rm)
}

// PADDSB
// Add packed signed byte integers from mm/m64 and mm and saturate the results.
// Add packed signed byte integers from xmm2/m128 and xmm1 saturate the results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func PADDSB(reg, rm interface{}) {
	unsafe.Asm("PADDSB", reg, rm)
}

// PADDSW
// Add packed signed word integers from mm/m64 and mm and saturate the results.
// Add packed signed word integers from xmm2/m128 and xmm1 and saturate the results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func PADDSW(reg, rm interface{}) {
	unsafe.Asm("PADDSW", reg, rm)
}

// PADDUSB
// Add packed unsigned byte integers from mm/m64 and mm and saturate the results.
// Add packed unsigned byte integers from xmm2/m128 and xmm1 saturate the results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func PADDUSB(reg, rm interface{}) {
	unsafe.Asm("PADDUSB", reg, rm)
}

// PADDUSW
// Add packed unsigned word integers from mm/m64 and mm and saturate the results.
// Add packed unsigned word integers from xmm2/m128 to xmm1 and saturate the results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func PADDUSW(reg, rm interface{}) {
	unsafe.Asm("PADDUSW", reg, rm)
}

// PALIGNR
// Concatenate destination and source operands, extract byte-aligned result shifted to the right by constant value in imm8 into mm1.
// Concatenate destination and source operands, extract byte-aligned result shifted to the right by constant value in imm8 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=871
func PALIGNR(reg, rm, imm interface{}) {
	unsafe.Asm("PALIGNR", reg, rm, imm)
}

// PAND
// Bitwise AND mm/m64 and mm.
// Bitwise AND of xmm2/m128 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=875
func PAND(reg, rm interface{}) {
	unsafe.Asm("PAND", reg, rm)
}

// PANDN
// Bitwise AND NOT of mm/m64 and mm.
// Bitwise AND NOT of xmm2/m128 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=878
func PANDN(reg, rm interface{}) {
	unsafe.Asm("PANDN", reg, rm)
}

// PAUSE
// Gives hint to processor that improves performance of spin-wait loops.
//
// Documentation: https://golang.org/s/x86manual#page=881
func PAUSE() {
	unsafe.Asm("PAUSE", nil)
}

// PAVGB
// Average packed unsigned byte integers from mm2/m64 and mm1 with rounding.
// Average packed unsigned byte integers from xmm2/m128 and xmm1 with rounding.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=882
func PAVGB(reg, rm interface{}) {
	unsafe.Asm("PAVGB", reg, rm)
}

// PAVGW
// Average packed unsigned word integers from mm2/m64 and mm1 with rounding.
// Average packed unsigned word integers from xmm2/m128 and xmm1 with rounding.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=882
func PAVGW(reg, rm interface{}) {
	unsafe.Asm("PAVGW", reg, rm)
}

// PBLENDVB
// Select byte values from xmm1 and xmm2/m128 from mask specified in the high bit of each byte in XMM0 and store the values into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// xmm: <XMM0>
//
// Documentation: https://golang.org/s/x86manual#page=886
func PBLENDVB(reg, rm, xmm interface{}) {
	unsafe.Asm("PBLENDVB", reg, rm, xmm)
}

// PBLENDW
// Select words from xmm1 and xmm2/m128 from mask specified in imm8 and store the values into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=890
func PBLENDW(reg, rm, imm interface{}) {
	unsafe.Asm("PBLENDW", reg, rm, imm)
}

// PCLMULQDQ
// Carry-less multiplication of one quadword of xmm1 by one quadword of xmm2/m128, stores the 128-bit result in xmm1. The imme- diate is used to determine which quadwords of xmm1 and xmm2/m128 should be used.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=893
func PCLMULQDQ(reg, rm, imm interface{}) {
	unsafe.Asm("PCLMULQDQ", reg, rm, imm)
}

// PCMPEQB
// Compare packed bytes in mm/m64 and mm for equality.
// Compare packed bytes in xmm2/m128 and xmm1 for equality.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func PCMPEQB(reg, rm interface{}) {
	unsafe.Asm("PCMPEQB", reg, rm)
}

// PCMPEQD
// Compare packed doublewords in mm/m64 and mm for equality.
// Compare packed doublewords in xmm2/m128 and xmm1 for equality.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func PCMPEQD(reg, rm interface{}) {
	unsafe.Asm("PCMPEQD", reg, rm)
}

// PCMPEQQ
// Compare packed qwords in xmm2/m128 and xmm1 for equality.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=902
func PCMPEQQ(reg, rm interface{}) {
	unsafe.Asm("PCMPEQQ", reg, rm)
}

// PCMPEQW
// Compare packed words in mm/m64 and mm for equality.
// Compare packed words in xmm2/m128 and xmm1 for equality.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func PCMPEQW(reg, rm interface{}) {
	unsafe.Asm("PCMPEQW", reg, rm)
}

// PCMPESTRI
// Perform a packed comparison of string data with explicit lengths, generating an index, and storing the result in ECX.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=905
func PCMPESTRI(reg, rm, imm interface{}) {
	unsafe.Asm("PCMPESTRI", reg, rm, imm)
}

// PCMPESTRM
// Perform a packed comparison of string data with explicit lengths, generating a mask, and storing the result in XMM0.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=907
func PCMPESTRM(reg, rm, imm interface{}) {
	unsafe.Asm("PCMPESTRM", reg, rm, imm)
}

// PCMPGTB
// Compare packed signed byte integers in mm and mm/m64 for greater than.
// Compare packed signed byte integers in xmm1 and xmm2/m128 for greater than.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func PCMPGTB(reg, rm interface{}) {
	unsafe.Asm("PCMPGTB", reg, rm)
}

// PCMPGTD
// Compare packed signed doubleword integers in mm and mm/m64 for greater than.
// Compare packed signed doubleword integers in xmm1 and xmm2/m128 for greater than.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func PCMPGTD(reg, rm interface{}) {
	unsafe.Asm("PCMPGTD", reg, rm)
}

// PCMPGTQ
// Compare packed signed qwords in xmm2/m128 and xmm1 for greater than.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=915
func PCMPGTQ(reg, rm interface{}) {
	unsafe.Asm("PCMPGTQ", reg, rm)
}

// PCMPGTW
// Compare packed signed word integers in mm and mm/m64 for greater than.
// Compare packed signed word integers in xmm1 and xmm2/m128 for greater than.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func PCMPGTW(reg, rm interface{}) {
	unsafe.Asm("PCMPGTW", reg, rm)
}

// PCMPISTRI
// Perform a packed comparison of string data with implicit lengths, generating an index, and storing the result in ECX.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=918
func PCMPISTRI(reg, rm, imm interface{}) {
	unsafe.Asm("PCMPISTRI", reg, rm, imm)
}

// PCMPISTRM
// Perform a packed comparison of string data with implicit lengths, generating a mask, and storing the result in XMM0.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=920
func PCMPISTRM(reg, rm, imm interface{}) {
	unsafe.Asm("PCMPISTRM", reg, rm, imm)
}

// PDEP
// Parallel deposit of bits from r32b using mask in r/m32, result is writ- ten to r32a.
// Parallel deposit of bits from r64b using mask in r/m64, result is writ- ten to r64a.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=922
func PDEP(reg, vex, rm interface{}) {
	unsafe.Asm("PDEP", reg, vex, rm)
}

// PEXT
// Parallel extract of bits from r32b using mask in r/m32, result is writ- ten to r32a.
// Parallel extract of bits from r64b using mask in r/m64, result is writ- ten to r64a.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=924
func PEXT(reg, vex, rm interface{}) {
	unsafe.Asm("PEXT", reg, vex, rm)
}

// PEXTRB
// Extract a byte integer value from xmm2 at the source byte offset specified by imm8 into reg or m8. The upper bits of r32 or r64 are zeroed.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func PEXTRB(rm, reg, imm interface{}) {
	unsafe.Asm("PEXTRB", rm, reg, imm)
}

// PEXTRD
// Extract a dword integer value from xmm2 at the source dword offset specified by imm8 into r/m32.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func PEXTRD(rm, reg, imm interface{}) {
	unsafe.Asm("PEXTRD", rm, reg, imm)
}

// PEXTRQ
// Extract a qword integer value from xmm2 at the source qword offset specified by imm8 into r/m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func PEXTRQ(rm, reg, imm interface{}) {
	unsafe.Asm("PEXTRQ", rm, reg, imm)
}

// PEXTRW_MRI
// Extract the word specified by imm8 from xmm and copy it to lowest 16 bits of reg or m16. Zero-extend the result in the destination, r32 or r64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=929
func PEXTRW_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("PEXTRW", rm, reg, imm)
}

// PEXTRW_RMI
// Extract the word specified by imm8 from mm and move it to reg, bits 15-0. The upper bits of r32 or r64 is zeroed.
// Extract the word specified by imm8 from xmm and move it to reg, bits 15-0. The upper bits of r32 or r64 is zeroed.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=929
func PEXTRW_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("PEXTRW", reg, rm, imm)
}

// PHADDD
// Add 32-bit integers horizontally, pack to mm1.
// Add 32-bit integers horizontally, pack to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=932
func PHADDD(reg, rm interface{}) {
	unsafe.Asm("PHADDD", reg, rm)
}

// PHADDSW
// Add 16-bit signed integers horizontally, pack saturated integers to mm1.
// Add 16-bit signed integers horizontally, pack saturated integers to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=936
func PHADDSW(reg, rm interface{}) {
	unsafe.Asm("PHADDSW", reg, rm)
}

// PHADDW
// Add 16-bit integers horizontally, pack to mm1.
// Add 16-bit integers horizontally, pack to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=932
func PHADDW(reg, rm interface{}) {
	unsafe.Asm("PHADDW", reg, rm)
}

// PHMINPOSUW
// Find the minimum unsigned word in xmm2/m128 and place its value in the low word of xmm1 and its index in the second- lowest word of xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=938
func PHMINPOSUW(reg, rm interface{}) {
	unsafe.Asm("PHMINPOSUW", reg, rm)
}

// PHSUBD
// Subtract 32-bit signed integers horizontally, pack to mm1.
// Subtract 32-bit signed integers horizontally, pack to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=940
func PHSUBD(reg, rm interface{}) {
	unsafe.Asm("PHSUBD", reg, rm)
}

// PHSUBSW
// Subtract 16-bit signed integer horizontally, pack saturated integers to mm1.
// Subtract 16-bit signed integer horizontally, pack saturated integers to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=943
func PHSUBSW(reg, rm interface{}) {
	unsafe.Asm("PHSUBSW", reg, rm)
}

// PHSUBW
// Subtract 16-bit signed integers horizontally, pack to mm1.
// Subtract 16-bit signed integers horizontally, pack to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=940
func PHSUBW(reg, rm interface{}) {
	unsafe.Asm("PHSUBW", reg, rm)
}

// PINSRB
// Insert a byte integer value from r32/m8 into xmm1 at the destination element in xmm1 specified by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func PINSRB(reg, rm, imm interface{}) {
	unsafe.Asm("PINSRB", reg, rm, imm)
}

// PINSRD
// Insert a dword integer value from r/m32 into the xmm1 at the destination element specified by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func PINSRD(reg, rm, imm interface{}) {
	unsafe.Asm("PINSRD", reg, rm, imm)
}

// PINSRQ
// Insert a qword integer value from r/m64 into the xmm1 at the destination element specified by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func PINSRQ(reg, rm, imm interface{}) {
	unsafe.Asm("PINSRQ", reg, rm, imm)
}

// PINSRW
// Insert the low word from r32 or from m16 into mm at the word position specified by imm8.
// Move the low word of r32 or from m16 into xmm at the word position specified by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=948
func PINSRW(reg, rm, imm interface{}) {
	unsafe.Asm("PINSRW", reg, rm, imm)
}

// PMADDUBSW
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to mm1.
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=950
func PMADDUBSW(reg, rm interface{}) {
	unsafe.Asm("PMADDUBSW", reg, rm)
}

// PMADDWD
// Multiply the packed words in mm by the packed words in mm/m64, add adjacent doubleword results, and store in mm.
// Multiply the packed word integers in xmm1 by the packed word integers in xmm2/m128, add adjacent doubleword results, and store in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=953
func PMADDWD(reg, rm interface{}) {
	unsafe.Asm("PMADDWD", reg, rm)
}

// PMAXSB
// Compare packed signed byte integers in xmm1 and xmm2/m128 and store packed maximum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func PMAXSB(reg, rm interface{}) {
	unsafe.Asm("PMAXSB", reg, rm)
}

// PMAXSD
// Compare packed signed dword integers in xmm1 and xmm2/m128 and store packed maximum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func PMAXSD(reg, rm interface{}) {
	unsafe.Asm("PMAXSD", reg, rm)
}

// PMAXSW
// Compare signed word integers in mm2/m64 and mm1 and return maximum values.
// Compare packed signed word integers in xmm2/m128 and xmm1 and stores maximum packed values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func PMAXSW(reg, rm interface{}) {
	unsafe.Asm("PMAXSW", reg, rm)
}

// PMAXUB
// Compare unsigned byte integers in mm2/m64 and mm1 and returns maximum values.
// Compare packed unsigned byte integers in xmm1 and xmm2/m128 and store packed maximum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func PMAXUB(reg, rm interface{}) {
	unsafe.Asm("PMAXUB", reg, rm)
}

// PMAXUW
// Compare packed unsigned word integers in xmm2/m128 and xmm1 and stores maximum packed values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func PMAXUW(reg, rm interface{}) {
	unsafe.Asm("PMAXUW", reg, rm)
}

// PMINSB
// Compare packed signed byte integers in xmm1 and xmm2/m128 and store packed minimum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func PMINSB(reg, rm interface{}) {
	unsafe.Asm("PMINSB", reg, rm)
}

// PMINSW
// Compare signed word integers in mm2/m64 and mm1 and return minimum values.
// Compare packed signed word integers in xmm2/m128 and xmm1 and store packed minimum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func PMINSW(reg, rm interface{}) {
	unsafe.Asm("PMINSW", reg, rm)
}

// PMINUB
// Compare unsigned byte integers in mm2/m64 and mm1 and returns minimum values.
// Compare packed unsigned byte integers in xmm1 and xmm2/m128 and store packed minimum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func PMINUB(reg, rm interface{}) {
	unsafe.Asm("PMINUB", reg, rm)
}

// PMINUD
// Compare packed unsigned dword integers in xmm1 and xmm2/m128 and store packed minimum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=986
func PMINUD(reg, rm interface{}) {
	unsafe.Asm("PMINUD", reg, rm)
}

// PMINUW
// Compare packed unsigned word integers in xmm2/m128 and xmm1 and store packed minimum values in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func PMINUW(reg, rm interface{}) {
	unsafe.Asm("PMINUW", reg, rm)
}

// PMOVMSKB
// Move a byte mask of mm to reg. The upper bits of r32 or r64 are zeroed
// Move a byte mask of xmm to reg. The upper bits of r32 or r64 are zeroed
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=990
func PMOVMSKB(reg, rm interface{}) {
	unsafe.Asm("PMOVMSKB", reg, rm)
}

// PMOVSXBD
// Sign extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXBD(reg, rm interface{}) {
	unsafe.Asm("PMOVSXBD", reg, rm)
}

// PMOVSXBQ
// Sign extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXBQ(reg, rm interface{}) {
	unsafe.Asm("PMOVSXBQ", reg, rm)
}

// PMOVSXBW
// Sign extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 16-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXBW(reg, rm interface{}) {
	unsafe.Asm("PMOVSXBW", reg, rm)
}

// PMOVSXDQ
// Sign extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXDQ(reg, rm interface{}) {
	unsafe.Asm("PMOVSXDQ", reg, rm)
}

// PMOVSXWD
// Sign extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 32-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXWD(reg, rm interface{}) {
	unsafe.Asm("PMOVSXWD", reg, rm)
}

// PMOVSXWQ
// Sign extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func PMOVSXWQ(reg, rm interface{}) {
	unsafe.Asm("PMOVSXWQ", reg, rm)
}

// PMOVZXBD
// Zero extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXBD(reg, rm interface{}) {
	unsafe.Asm("PMOVZXBD", reg, rm)
}

// PMOVZXBQ
// Zero extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXBQ(reg, rm interface{}) {
	unsafe.Asm("PMOVZXBQ", reg, rm)
}

// PMOVZXBW
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 16-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXBW(reg, rm interface{}) {
	unsafe.Asm("PMOVZXBW", reg, rm)
}

// PMOVZXDQ
// Zero extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXDQ(reg, rm interface{}) {
	unsafe.Asm("PMOVZXDQ", reg, rm)
}

// PMOVZXWD
// Zero extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 32-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXWD(reg, rm interface{}) {
	unsafe.Asm("PMOVZXWD", reg, rm)
}

// PMOVZXWQ
// Zero extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func PMOVZXWQ(reg, rm interface{}) {
	unsafe.Asm("PMOVZXWQ", reg, rm)
}

// PMULDQ
// Multiply packed signed doubleword integers in xmm1 by packed signed doubleword integers in xmm2/m128, and store the quadword results in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1011
func PMULDQ(reg, rm interface{}) {
	unsafe.Asm("PMULDQ", reg, rm)
}

// PMULHRSW
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to mm1.
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1014
func PMULHRSW(reg, rm interface{}) {
	unsafe.Asm("PMULHRSW", reg, rm)
}

// PMULHUW
// Multiply the packed unsigned word integers in mm1 register and mm2/m64, and store the high 16 bits of the results in mm1.
// Multiply the packed unsigned word integers in xmm1 and xmm2/m128, and store the high 16 bits of the results in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1018
func PMULHUW(reg, rm interface{}) {
	unsafe.Asm("PMULHUW", reg, rm)
}

// PMULHW
// Multiply the packed signed word integers in mm1 register and mm2/m64, and store the high 16 bits of the results in mm1.
// Multiply the packed signed word integers in xmm1 and xmm2/m128, and store the high 16 bits of the results in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1022
func PMULHW(reg, rm interface{}) {
	unsafe.Asm("PMULHW", reg, rm)
}

// PMULLW
// Multiply the packed signed word integers in mm1 register and mm2/m64, and store the low 16 bits of the results in mm1.
// Multiply the packed signed word integers in xmm1 and xmm2/m128, and store the low 16 bits of the results in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1030
func PMULLW(reg, rm interface{}) {
	unsafe.Asm("PMULLW", reg, rm)
}

// PMULUDQ
// Multiply unsigned doubleword integer in mm1 by unsigned doubleword integer in mm2/m64, and store the quadword result in mm1.
// Multiply packed unsigned doubleword integers in xmm1 by packed unsigned doubleword integers in xmm2/m128, and store the quadword results in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1034
func PMULUDQ(reg, rm interface{}) {
	unsafe.Asm("PMULUDQ", reg, rm)
}

// POP_M
// Pop top of stack into m16; increment stack pointer.
// Pop top of stack into m32; increment stack pointer.
// Pop top of stack into m64; increment stack pointer. Cannot encode 32-bit operand size.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1037
func POP_M(rm interface{}) {
	unsafe.Asm("POP", rm)
}

// POP_NP
// Pop top of stack into DS; increment stack pointer.
// Pop top of stack into ES; increment stack pointer.
// Pop top of stack into FS; increment stack pointer by 16 bits.
// Pop top of stack into FS; increment stack pointer by 64 bits.
// Pop top of stack into FS; increment stack pointer by 32 bits.
// Pop top of stack into GS; increment stack pointer by 32 bits.
// Pop top of stack into GS; increment stack pointer by 64 bits.
// Pop top of stack into GS; increment stack pointer by 16 bits.
// Pop top of stack into SS; increment stack pointer.
//
// ds: DS (w)
//
// Documentation: https://golang.org/s/x86manual#page=1037
func POP_NP(ds interface{}) {
	unsafe.Asm("POP", ds)
}

// POP_O
// Pop top of stack into r16; increment stack pointer.
// Pop top of stack into r32; increment stack pointer.
// Pop top of stack into r64; increment stack pointer. Cannot encode 32-bit operand size.
//
// opcode: opcode + rd (w)
//
// Documentation: https://golang.org/s/x86manual#page=1037
func POP_O(opcode interface{}) {
	unsafe.Asm("POP", opcode)
}

// POPA
// Pop DI, SI, BP, BX, DX, CX, and AX.
//
// Documentation: https://golang.org/s/x86manual#page=1042
func POPA() {
	unsafe.Asm("POPA", nil)
}

// POPAD
// Pop EDI, ESI, EBP, EBX, EDX, ECX, and EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1042
func POPAD() {
	unsafe.Asm("POPAD", nil)
}

// POPCNT
// POPCNT on r/m16
// POPCNT on r/m32
// POPCNT on r/m64
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1044
func POPCNT(reg, rm interface{}) {
	unsafe.Asm("POPCNT", reg, rm)
}

// POPF
// Pop top of stack into lower 16 bits of EFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1046
func POPF() {
	unsafe.Asm("POPF", nil)
}

// POPFD
// Pop top of stack into EFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1046
func POPFD() {
	unsafe.Asm("POPFD", nil)
}

// POPFQ
// Pop top of stack and zero-extend into RFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1046
func POPFQ() {
	unsafe.Asm("POPFQ", nil)
}

// POR
// Bitwise OR of mm/m64 and mm.
// Bitwise OR of xmm2/m128 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1051
func POR(reg, rm interface{}) {
	unsafe.Asm("POR", reg, rm)
}

// PREFETCHNTA
// Move data from m8 closer to the processor using NTA hint.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1054
func PREFETCHNTA(rm interface{}) {
	unsafe.Asm("PREFETCHNTA", rm)
}

// PREFETCHT0
// Move data from m8 closer to the processor using T0 hint.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1054
func PREFETCHT0(rm interface{}) {
	unsafe.Asm("PREFETCHT0", rm)
}

// PREFETCHT1
// Move data from m8 closer to the processor using T1 hint.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1054
func PREFETCHT1(rm interface{}) {
	unsafe.Asm("PREFETCHT1", rm)
}

// PREFETCHT2
// Move data from m8 closer to the processor using T2 hint.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1054
func PREFETCHT2(rm interface{}) {
	unsafe.Asm("PREFETCHT2", rm)
}

// PREFETCHW
// Move data from m8 closer to the processor in anticipation of a write.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1056
func PREFETCHW(rm interface{}) {
	unsafe.Asm("PREFETCHW", rm)
}

// PREFETCHWT1
// Move data from m8 closer to the processor using T1 hint with intent to write.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1058
func PREFETCHWT1(rm interface{}) {
	unsafe.Asm("PREFETCHWT1", rm)
}

// PSHUFB
// Shuffle bytes in mm1 according to contents of mm2/m64.
// Shuffle bytes in xmm1 according to contents of xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1064
func PSHUFB(reg, rm interface{}) {
	unsafe.Asm("PSHUFB", reg, rm)
}

// PSHUFD
// Shuffle the doublewords in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1068
func PSHUFD(reg, rm, imm interface{}) {
	unsafe.Asm("PSHUFD", reg, rm, imm)
}

// PSHUFHW
// Shuffle the high words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1072
func PSHUFHW(reg, rm, imm interface{}) {
	unsafe.Asm("PSHUFHW", reg, rm, imm)
}

// PSHUFLW
// Shuffle the low words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1075
func PSHUFLW(reg, rm, imm interface{}) {
	unsafe.Asm("PSHUFLW", reg, rm, imm)
}

// PSHUFW
// Shuffle the words in mm2/m64 based on the encoding in imm8 and store the result in mm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1078
func PSHUFW(reg, rm, imm interface{}) {
	unsafe.Asm("PSHUFW", reg, rm, imm)
}

// PSIGNB
// Negate/zero/preserve packed byte integers in mm1 depending on the corresponding sign in mm2/m64.
// Negate/zero/preserve packed byte integers in xmm1 depending on the corresponding sign in xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func PSIGNB(reg, rm interface{}) {
	unsafe.Asm("PSIGNB", reg, rm)
}

// PSIGND
// Negate/zero/preserve packed doubleword integers in mm1 depending on the corresponding sign in mm2/m128.
// Negate/zero/preserve packed doubleword integers in xmm1 depending on the corresponding sign in xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func PSIGND(reg, rm interface{}) {
	unsafe.Asm("PSIGND", reg, rm)
}

// PSIGNW
// Negate/zero/preserve packed word integers in mm1 depending on the corresponding sign in mm2/m128.
// Negate/zero/preserve packed word integers in xmm1 depending on the corresponding sign in xmm2/m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func PSIGNW(reg, rm interface{}) {
	unsafe.Asm("PSIGNW", reg, rm)
}

// PSLLD_MI
// Shift doublewords in mm left by imm8 while shifting in 0s.
// Shift doublewords in xmm1 left by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLD_MI(rm, imm interface{}) {
	unsafe.Asm("PSLLD", rm, imm)
}

// PSLLD_RM
// Shift doublewords in mm left by mm/m64 while shifting in 0s.
// Shift doublewords in xmm1 left by xmm2/m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLD_RM(reg, rm interface{}) {
	unsafe.Asm("PSLLD", reg, rm)
}

// PSLLDQ
// Shift xmm1 left by imm8 bytes while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1083
func PSLLDQ(rm, imm interface{}) {
	unsafe.Asm("PSLLDQ", rm, imm)
}

// PSLLQ_MI
// Shift quadword in mm left by imm8 while shifting in 0s.
// Shift quadwords in xmm1 left by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLQ_MI(rm, imm interface{}) {
	unsafe.Asm("PSLLQ", rm, imm)
}

// PSLLQ_RM
// Shift quadword in mm left by mm/m64 while shifting in 0s.
// Shift quadwords in xmm1 left by xmm2/m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLQ_RM(reg, rm interface{}) {
	unsafe.Asm("PSLLQ", reg, rm)
}

// PSLLW_MI
// Shift words in mm left by imm8 while shifting in 0s.
// Shift words in xmm1 left by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLW_MI(rm, imm interface{}) {
	unsafe.Asm("PSLLW", rm, imm)
}

// PSLLW_RM
// Shift words in mm left mm/m64 while shifting in 0s.
// Shift words in xmm1 left by xmm2/m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func PSLLW_RM(reg, rm interface{}) {
	unsafe.Asm("PSLLW", reg, rm)
}

// PSRAD_MI
// Shift doublewords in mm right by imm8 while shifting in sign bits.
// Shift doublewords in xmm1 right by imm8 while shifting in sign bits.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1097
func PSRAD_MI(rm, imm interface{}) {
	unsafe.Asm("PSRAD", rm, imm)
}

// PSRAD_RM
// Shift doublewords in mm right by mm/m64 while shifting in sign bits.
// Shift doubleword in xmm1 right by xmm2 /m128 while shifting in sign bits.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1097
func PSRAD_RM(reg, rm interface{}) {
	unsafe.Asm("PSRAD", reg, rm)
}

// PSRAW_MI
// Shift words in mm right by imm8 while shifting in sign bits
// Shift words in xmm1 right by imm8 while shifting in sign bits
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1097
func PSRAW_MI(rm, imm interface{}) {
	unsafe.Asm("PSRAW", rm, imm)
}

// PSRAW_RM
// Shift words in mm right by mm/m64 while shifting in sign bits.
// Shift words in xmm1 right by xmm2/m128 while shifting in sign bits.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1097
func PSRAW_RM(reg, rm interface{}) {
	unsafe.Asm("PSRAW", reg, rm)
}

// PSRLD_MI
// Shift doublewords in mm right by imm8 while shifting in 0s.
// Shift doublewords in xmm1 right by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLD_MI(rm, imm interface{}) {
	unsafe.Asm("PSRLD", rm, imm)
}

// PSRLD_RM
// Shift doublewords in mm right by amount specified in mm/m64 while shifting in 0s.
// Shift doublewords in xmm1 right by amount specified in xmm2 /m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLD_RM(reg, rm interface{}) {
	unsafe.Asm("PSRLD", reg, rm)
}

// PSRLDQ
// Shift xmm1 right by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1107
func PSRLDQ(rm, imm interface{}) {
	unsafe.Asm("PSRLDQ", rm, imm)
}

// PSRLQ_MI
// Shift mm right by imm8 while shifting in 0s.
// Shift quadwords in xmm1 right by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLQ_MI(rm, imm interface{}) {
	unsafe.Asm("PSRLQ", rm, imm)
}

// PSRLQ_RM
// Shift mm right by amount specified in mm/m64 while shifting in 0s.
// Shift quadwords in xmm1 right by amount specified in xmm2/m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLQ_RM(reg, rm interface{}) {
	unsafe.Asm("PSRLQ", reg, rm)
}

// PSRLW_MI
// Shift words in mm right by imm8 while shifting in 0s.
// Shift words in xmm1 right by imm8 while shifting in 0s.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLW_MI(rm, imm interface{}) {
	unsafe.Asm("PSRLW", rm, imm)
}

// PSRLW_RM
// Shift words in mm right by amount specified in mm/m64 while shifting in 0s.
// Shift words in xmm1 right by amount specified in xmm2/m128 while shifting in 0s.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func PSRLW_RM(reg, rm interface{}) {
	unsafe.Asm("PSRLW", reg, rm)
}

// PSUBB
// Subtract packed byte integers in mm/m64 from packed byte integers in mm.
// Subtract packed byte integers in xmm2/m128 from packed byte integers in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func PSUBB(reg, rm interface{}) {
	unsafe.Asm("PSUBB", reg, rm)
}

// PSUBD
// Subtract packed doubleword integers in mm/m64 from packed doubleword integers in mm.
// Subtract packed doubleword integers in xmm2/mem128 from packed doubleword integers in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func PSUBD(reg, rm interface{}) {
	unsafe.Asm("PSUBD", reg, rm)
}

// PSUBQ
// Subtract quadword integer in mm1 from mm2 /m64.
// Subtract packed quadword integers in xmm1 from xmm2 /m128.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1128
func PSUBQ(reg, rm interface{}) {
	unsafe.Asm("PSUBQ", reg, rm)
}

// PSUBSB
// Subtract signed packed bytes in mm/m64 from signed packed bytes in mm and saturate results.
// Subtract packed signed byte integers in xmm2/m128 from packed signed byte integers in xmm1 and saturate results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func PSUBSB(reg, rm interface{}) {
	unsafe.Asm("PSUBSB", reg, rm)
}

// PSUBSW
// Subtract signed packed words in mm/m64 from signed packed words in mm and saturate results.
// Subtract packed signed word integers in xmm2/m128 from packed signed word integers in xmm1 and saturate results.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func PSUBSW(reg, rm interface{}) {
	unsafe.Asm("PSUBSW", reg, rm)
}

// PSUBUSB
// Subtract unsigned packed bytes in mm/m64 from unsigned packed bytes in mm and saturate result.
// Subtract packed unsigned byte integers in xmm2/m128 from packed unsigned byte integers in xmm1 and saturate result.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func PSUBUSB(reg, rm interface{}) {
	unsafe.Asm("PSUBUSB", reg, rm)
}

// PSUBUSW
// Subtract unsigned packed words in mm/m64 from unsigned packed words in mm and saturate result.
// Subtract packed unsigned word integers in xmm2/m128 from packed unsigned word integers in xmm1 and saturate result.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func PSUBUSW(reg, rm interface{}) {
	unsafe.Asm("PSUBUSW", reg, rm)
}

// PSUBW
// Subtract packed word integers in mm/m64 from packed word integers in mm.
// Subtract packed word integers in xmm2/m128 from packed word integers in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func PSUBW(reg, rm interface{}) {
	unsafe.Asm("PSUBW", reg, rm)
}

// PTEST
// Set ZF if xmm2/m128 AND xmm1 result is all 0s. Set CF if xmm2/m128 AND NOT xmm1 result is all 0s.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1139
func PTEST(reg, rm interface{}) {
	unsafe.Asm("PTEST", reg, rm)
}

// PUNPCKHBW
// Unpack and interleave high-order bytes from mm and mm/m64 into mm.
// Unpack and interleave high-order bytes from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func PUNPCKHBW(reg, rm interface{}) {
	unsafe.Asm("PUNPCKHBW", reg, rm)
}

// PUNPCKHDQ
// Unpack and interleave high-order doublewords from mm and mm/m64 into mm.
// Unpack and interleave high-order doublewords from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func PUNPCKHDQ(reg, rm interface{}) {
	unsafe.Asm("PUNPCKHDQ", reg, rm)
}

// PUNPCKHQDQ
// Unpack and interleave high-order quadwords from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func PUNPCKHQDQ(reg, rm interface{}) {
	unsafe.Asm("PUNPCKHQDQ", reg, rm)
}

// PUNPCKHWD
// Unpack and interleave high-order words from mm and mm/m64 into mm.
// Unpack and interleave high-order words from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func PUNPCKHWD(reg, rm interface{}) {
	unsafe.Asm("PUNPCKHWD", reg, rm)
}

// PUNPCKLBW
// Interleave low-order bytes from mm and mm/m32 into mm.
// Interleave low-order bytes from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func PUNPCKLBW(reg, rm interface{}) {
	unsafe.Asm("PUNPCKLBW", reg, rm)
}

// PUNPCKLDQ
// Interleave low-order doublewords from mm and mm/m32 into mm.
// Interleave low-order doublewords from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func PUNPCKLDQ(reg, rm interface{}) {
	unsafe.Asm("PUNPCKLDQ", reg, rm)
}

// PUNPCKLQDQ
// Interleave low-order quadword from xmm1 and xmm2/m128 into xmm1 register.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func PUNPCKLQDQ(reg, rm interface{}) {
	unsafe.Asm("PUNPCKLQDQ", reg, rm)
}

// PUNPCKLWD
// Interleave low-order words from mm and mm/m32 into mm.
// Interleave low-order words from xmm1 and xmm2/m128 into xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func PUNPCKLWD(reg, rm interface{}) {
	unsafe.Asm("PUNPCKLWD", reg, rm)
}

// PUSH_I
// Push imm16.
// Push imm32.
// Push imm8.
//
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1163
func PUSH_I(imm interface{}) {
	unsafe.Asm("PUSH", imm)
}

// PUSH_M
// Push r/m16.
// Push r/m32.
// Push r/m64.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1163
func PUSH_M(rm interface{}) {
	unsafe.Asm("PUSH", rm)
}

// PUSH_NP
// Push CS.
// Push DS.
// Push ES.
// Push FS.
// Push GS.
// Push SS.
//
// cs: CS (r)
//
// Documentation: https://golang.org/s/x86manual#page=1163
func PUSH_NP(cs interface{}) {
	unsafe.Asm("PUSH", cs)
}

// PUSH_O
// Push r16.
// Push r32.
// Push r64.
//
// opcode: opcode + rd (r)
//
// Documentation: https://golang.org/s/x86manual#page=1163
func PUSH_O(opcode interface{}) {
	unsafe.Asm("PUSH", opcode)
}

// PUSHA
// Push AX, CX, DX, BX, original SP, BP, SI, and DI.
//
// Documentation: https://golang.org/s/x86manual#page=1166
func PUSHA() {
	unsafe.Asm("PUSHA", nil)
}

// PUSHAD
// Push EAX, ECX, EDX, EBX, original ESP, EBP, ESI, and EDI.
//
// Documentation: https://golang.org/s/x86manual#page=1166
func PUSHAD() {
	unsafe.Asm("PUSHAD", nil)
}

// PUSHF
// Push lower 16 bits of EFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1168
func PUSHF() {
	unsafe.Asm("PUSHF", nil)
}

// PUSHFD
// Push EFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1168
func PUSHFD() {
	unsafe.Asm("PUSHFD", nil)
}

// PUSHFQ
// Push RFLAGS.
//
// Documentation: https://golang.org/s/x86manual#page=1168
func PUSHFQ() {
	unsafe.Asm("PUSHFQ", nil)
}

// PXOR
// Bitwise XOR of mm/m64 and mm.
// Bitwise XOR of xmm2/m128 and xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1170
func PXOR(reg, rm interface{}) {
	unsafe.Asm("PXOR", reg, rm)
}

// RCL_M1
// Rotate 17 bits (CF, r/m16) left once.
// Rotate 33 bits (CF, r/m32) left once.
// Rotate 65 bits (CF, r/m64) left once. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) left once.
//
// rm: ModRM:r/m (w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCL_M1(rm, v1 interface{}) {
	unsafe.Asm("RCL", rm, v1)
}

// RCL_MC
// Rotate 17 bits (CF, r/m16) left CL times.
// Rotate 33 bits (CF, r/m32) left CL times.
// Rotate 65 bits (CF, r/m64) left CL times. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) left CL times.
//
// rm: ModRM:r/m (w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCL_MC(rm, cl interface{}) {
	unsafe.Asm("RCL", rm, cl)
}

// RCL_MI
// Rotate 17 bits (CF, r/m16) left imm8 times.
// Rotate 33 bits (CF, r/m32) left imm8 times.
// Rotate 65 bits (CF, r/m64) left imm8 times. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) left imm8 times.
//
// rm: ModRM:r/m (w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCL_MI(rm, imm interface{}) {
	unsafe.Asm("RCL", rm, imm)
}

// RCPPS
// Computes the approximate reciprocals of the packed single-precision floating-point values in xmm2/m128 and stores the results in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1178
func RCPPS(reg, rm interface{}) {
	unsafe.Asm("RCPPS", reg, rm)
}

// RCPSS
// Computes the approximate reciprocal of the scalar single-precision floating-point value in xmm2/m32 and stores the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1180
func RCPSS(reg, rm interface{}) {
	unsafe.Asm("RCPSS", reg, rm)
}

// RCR_M1
// Rotate 17 bits (CF, r/m16) right once.
// Rotate 33 bits (CF, r/m32) right once. Uses a 6 bit count.
// Rotate 65 bits (CF, r/m64) right once. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) right once.
//
// rm: ModRM:r/m (w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCR_M1(rm, v1 interface{}) {
	unsafe.Asm("RCR", rm, v1)
}

// RCR_MC
// Rotate 17 bits (CF, r/m16) right CL times.
// Rotate 33 bits (CF, r/m32) right CL times.
// Rotate 65 bits (CF, r/m64) right CL times. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) right CL times.
//
// rm: ModRM:r/m (w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCR_MC(rm, cl interface{}) {
	unsafe.Asm("RCR", rm, cl)
}

// RCR_MI
// Rotate 17 bits (CF, r/m16) right imm8 times.
// Rotate 33 bits (CF, r/m32) right imm8 times.
// Rotate 65 bits (CF, r/m64) right imm8 times. Uses a 6 bit count.
// Rotate 9 bits (CF, r/m8) right imm8 times.
//
// rm: ModRM:r/m (w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1173
func RCR_MI(rm, imm interface{}) {
	unsafe.Asm("RCR", rm, imm)
}

// RDFSBASE
// Load the 32-bit destination register with the FS base address.
// Load the 64-bit destination register with the FS base address.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1182
func RDFSBASE(rm interface{}) {
	unsafe.Asm("RDFSBASE", rm)
}

// RDGSBASE
// Load the 32-bit destination register with the GS base address.
// Load the 64-bit destination register with the GS base address.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1182
func RDGSBASE(rm interface{}) {
	unsafe.Asm("RDGSBASE", rm)
}

// RDMSR
// Read MSR specified by ECX into EDX:EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1184
func RDMSR() {
	unsafe.Asm("RDMSR", nil)
}

// RDPID
// Read IA32_TSC_AUX into r32.
// Read IA32_TSC_AUX into r64.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1186
func RDPID(rm interface{}) {
	unsafe.Asm("RDPID", rm)
}

// RDPKRU
// Reads PKRU into EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1187
func RDPKRU() {
	unsafe.Asm("RDPKRU", nil)
}

// RDPMC
// Read performance-monitoring counter specified by ECX into EDX:EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1189
func RDPMC() {
	unsafe.Asm("RDPMC", nil)
}

// RDRAND
// Read a 16-bit random number and store in the destination register.
// Read a 32-bit random number and store in the destination register.
// Read a 64-bit random number and store in the destination register.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1193
func RDRAND(rm interface{}) {
	unsafe.Asm("RDRAND", rm)
}

// RDSEED
// Read a 16-bit NIST SP800-90B & C compliant random value and store in the destination register.
// Read a 32-bit NIST SP800-90B & C compliant random value and store in the destination register.
// Read a 64-bit NIST SP800-90B & C compliant random value and store in the destination register.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1195
func RDSEED(rm interface{}) {
	unsafe.Asm("RDSEED", rm)
}

// RDTSC
// Read time-stamp counter into EDX:EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1197
func RDTSC() {
	unsafe.Asm("RDTSC", nil)
}

// RDTSCP
// Read 64-bit time-stamp counter and IA32_TSC_AUX value into EDX:EAX and ECX.
//
// Documentation: https://golang.org/s/x86manual#page=1199
func RDTSCP() {
	unsafe.Asm("RDTSCP", nil)
}

// RET_I
// Near return to calling procedure and pop imm16 bytes from stack.
// Far return to calling procedure and pop imm16 bytes from stack.
//
// imm: imm16
//
// Documentation: https://golang.org/s/x86manual#page=1205
func RET_I(imm interface{}) {
	unsafe.Asm("RET", imm)
}

// RET_NP
// Near return to calling procedure.
// Far return to calling procedure.
//
// Documentation: https://golang.org/s/x86manual#page=1205
func RET_NP() {
	unsafe.Asm("RET", nil)
}

// ROL_M1
// Rotate 16 bits r/m16 left once.
// Rotate 32 bits r/m32 left once.
// Rotate 64 bits r/m64 left once. Uses a 6 bit count.
// Rotate 8 bits r/m8 left once.
// Rotate 8 bits r/m8 left once
//
// rm: ModRM:r/m (w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROL_M1(rm, v1 interface{}) {
	unsafe.Asm("ROL", rm, v1)
}

// ROL_MC
// Rotate 16 bits r/m16 left CL times.
// Rotate 32 bits r/m32 left CL times.
// Rotate 64 bits r/m64 left CL times. Uses a 6 bit count.
// Rotate 8 bits r/m8 left CL times.
//
// rm: ModRM:r/m (w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROL_MC(rm, cl interface{}) {
	unsafe.Asm("ROL", rm, cl)
}

// ROL_MI
// Rotate 16 bits r/m16 left imm8 times.
// Rotate 32 bits r/m32 left imm8 times.
// Rotate 64 bits r/m64 left imm8 times. Uses a 6 bit count.
// Rotate 8 bits r/m8 left imm8 times.
//
// rm: ModRM:r/m (w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROL_MI(rm, imm interface{}) {
	unsafe.Asm("ROL", rm, imm)
}

// ROR_M1
// Rotate 16 bits r/m16 right once.
// Rotate 32 bits r/m32 right once.
// Rotate 64 bits r/m64 right once. Uses a 6 bit count.
// Rotate 8 bits r/m8 right once.
//
// rm: ModRM:r/m (w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROR_M1(rm, v1 interface{}) {
	unsafe.Asm("ROR", rm, v1)
}

// ROR_MC
// Rotate 16 bits r/m16 right CL times.
// Rotate 32 bits r/m32 right CL times.
// Rotate 64 bits r/m64 right CL times. Uses a 6 bit count.
// Rotate 8 bits r/m8 right CL times.
//
// rm: ModRM:r/m (w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROR_MC(rm, cl interface{}) {
	unsafe.Asm("ROR", rm, cl)
}

// ROR_MI
// Rotate 16 bits r/m16 right imm8 times.
// Rotate 32 bits r/m32 right imm8 times.
// Rotate 64 bits r/m64 right imm8 times. Uses a 6 bit count.
// Rotate 8 bits r/m16 right imm8 times.
//
// rm: ModRM:r/m (w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1173
func ROR_MI(rm, imm interface{}) {
	unsafe.Asm("ROR", rm, imm)
}

// RORX
// Rotate 32-bit r/m32 right imm8 times without affecting arithmetic flags.
// Rotate 64-bit r/m64 right imm8 times without affecting arithmetic flags.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1215
func RORX(reg, rm, imm interface{}) {
	unsafe.Asm("RORX", reg, rm, imm)
}

// ROUNDPD
// Round packed double precision floating-point values in xmm2/m128 and place the result in xmm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1216
func ROUNDPD(reg, rm, imm interface{}) {
	unsafe.Asm("ROUNDPD", reg, rm, imm)
}

// ROUNDPS
// Round packed single precision floating-point values in xmm2/m128 and place the result in xmm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1219
func ROUNDPS(reg, rm, imm interface{}) {
	unsafe.Asm("ROUNDPS", reg, rm, imm)
}

// ROUNDSD
// Round the low packed double precision floating-point value in xmm2/m64 and place the result in xmm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1222
func ROUNDSD(reg, rm, imm interface{}) {
	unsafe.Asm("ROUNDSD", reg, rm, imm)
}

// ROUNDSS
// Round the low packed single precision floating-point value in xmm2/m32 and place the result in xmm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1224
func ROUNDSS(reg, rm, imm interface{}) {
	unsafe.Asm("ROUNDSS", reg, rm, imm)
}

// RSM
// Resume operation of interrupted program.
//
// Documentation: https://golang.org/s/x86manual#page=1226
func RSM() {
	unsafe.Asm("RSM", nil)
}

// RSQRTPS
// Computes the approximate reciprocals of the square roots of the packed single-precision floating-point values in xmm2/m128 and stores the results in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1228
func RSQRTPS(reg, rm interface{}) {
	unsafe.Asm("RSQRTPS", reg, rm)
}

// RSQRTSS
// Computes the approximate reciprocal of the square root of the low single-precision floating-point value in xmm2/m32 and stores the results in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1230
func RSQRTSS(reg, rm interface{}) {
	unsafe.Asm("RSQRTSS", reg, rm)
}

// SAHF
// Loads SF, ZF, AF, PF, and CF from AH into EFLAGS register.
//
// Documentation: https://golang.org/s/x86manual#page=1232
func SAHF() {
	unsafe.Asm("SAHF", nil)
}

// SAL_M1
// Multiply r/m16 by 2, once.
// Multiply r/m32 by 2, once.
// Multiply r/m64 by 2, once.
// Multiply r/m8 by 2, once.
//
// rm: ModRM:r/m (r, w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAL_M1(rm, v1 interface{}) {
	unsafe.Asm("SAL", rm, v1)
}

// SAL_MC
// Multiply r/m16 by 2, CL times.
// Multiply r/m32 by 2, CL times.
// Multiply r/m64 by 2, CL times.
// Multiply r/m8 by 2, CL times.
//
// rm: ModRM:r/m (r, w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAL_MC(rm, cl interface{}) {
	unsafe.Asm("SAL", rm, cl)
}

// SAL_MI
// Multiply r/m16 by 2, imm8 times.
// Multiply r/m32 by 2, imm8 times.
// Multiply r/m64 by 2, imm8 times.
// Multiply r/m8 by 2, imm8 times.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAL_MI(rm, imm interface{}) {
	unsafe.Asm("SAL", rm, imm)
}

// SAR_M1
// Signed divide* r/m16 by 2, once.
// Signed divide* r/m32 by 2, once.
// Signed divide* r/m64 by 2, once.
// Signed divide* r/m8 by 2, once.
//
// rm: ModRM:r/m (r, w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAR_M1(rm, v1 interface{}) {
	unsafe.Asm("SAR", rm, v1)
}

// SAR_MC
// Signed divide* r/m16 by 2, CL times.
// Signed divide* r/m32 by 2, CL times.
// Signed divide* r/m64 by 2, CL times.
// Signed divide* r/m8 by 2, CL times.
//
// rm: ModRM:r/m (r, w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAR_MC(rm, cl interface{}) {
	unsafe.Asm("SAR", rm, cl)
}

// SAR_MI
// Signed divide* r/m16 by 2, imm8 times.
// Signed divide* r/m32 by 2, imm8 times.
// Signed divide* r/m64 by 2, imm8 times
// Signed divide* r/m8 by 2, imm8 time.
// Signed divide* r/m8 by 2, imm8 times.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SAR_MI(rm, imm interface{}) {
	unsafe.Asm("SAR", rm, imm)
}

// SARX
// Shift r/m32 arithmetically right with count specified in r32b.
// Shift r/m64 arithmetically right with count specified in r64b.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// vex: VEX.vvvv (r)
//
// Documentation: https://golang.org/s/x86manual#page=1239
func SARX(reg, rm, vex interface{}) {
	unsafe.Asm("SARX", reg, rm, vex)
}

// SBB_I
// Subtract with borrow imm8 from AL.
// Subtract with borrow imm16 from AX.
// Subtract with borrow imm32 from EAX.
// Subtract with borrow sign-extended imm.32 to 64-bits from RAX.
//
// al: AL/AX/EAX/RAX
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1241
func SBB_I(al, imm interface{}) {
	unsafe.Asm("SBB", al, imm)
}

// SBB_MI
// Subtract with borrow imm16 from r/m16.
// Subtract with borrow sign-extended imm8 from r/m16.
// Subtract with borrow imm32 from r/m32.
// Subtract with borrow sign-extended imm8 from r/m32.
// Subtract with borrow sign-extended imm32 to 64-bits from r/m64.
// Subtract with borrow sign-extended imm8 from r/m64.
// Subtract with borrow imm8 from r/m8.
//
// rm: ModRM:r/m (w)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1241
func SBB_MI(rm, imm interface{}) {
	unsafe.Asm("SBB", rm, imm)
}

// SBB_MR
// Subtract with borrow r16 from r/m16.
// Subtract with borrow r32 from r/m32.
// Subtract with borrow r64 from r/m64.
// Subtract with borrow r8 from r/m8.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1241
func SBB_MR(rm, reg interface{}) {
	unsafe.Asm("SBB", rm, reg)
}

// SBB_RM
// Subtract with borrow r/m16 from r16.
// Subtract with borrow r/m32 from r32.
// Subtract with borrow r/m64 from r64.
// Subtract with borrow r/m8 from r8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1241
func SBB_RM(reg, rm interface{}) {
	unsafe.Asm("SBB", reg, rm)
}

// SCASB
// Compare AL with byte at ES:(E)DI or RDI then set status flags.
//
// Documentation: https://golang.org/s/x86manual#page=1244
func SCASB() {
	unsafe.Asm("SCASB", nil)
}

// SCASD
// Compare EAX w*ith doubleword at ES:(E)DI or RDI then set status flags.
//
// Documentation: https://golang.org/s/x86manual#page=1244
func SCASD() {
	unsafe.Asm("SCASD", nil)
}

// SCASQ
// Compare RAX with quadw*ord at RDI or EDI then set status flags.
//
// Documentation: https://golang.org/s/x86manual#page=1244
func SCASQ() {
	unsafe.Asm("SCASQ", nil)
}

// SCASW
// Compare AX wit*h word at ES:(E)DI or RDI then set status flags.
//
// Documentation: https://golang.org/s/x86manual#page=1244
func SCASW() {
	unsafe.Asm("SCASW", nil)
}

// SETA
// Set byte if above (CF=0 and ZF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETA(rm interface{}) {
	unsafe.Asm("SETA", rm)
}

// SETAE
// Set byte if above or equal (CF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETAE(rm interface{}) {
	unsafe.Asm("SETAE", rm)
}

// SETB
// Set byte if below (CF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETB(rm interface{}) {
	unsafe.Asm("SETB", rm)
}

// SETBE
// Set byte if below or equal (CF=1 or ZF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETBE(rm interface{}) {
	unsafe.Asm("SETBE", rm)
}

// SETC
// Set byte if carry (CF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETC(rm interface{}) {
	unsafe.Asm("SETC", rm)
}

// SETE
// Set byte if equal (ZF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETE(rm interface{}) {
	unsafe.Asm("SETE", rm)
}

// SETG
// Set byte if greater (ZF=0 and SF=OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETG(rm interface{}) {
	unsafe.Asm("SETG", rm)
}

// SETGE
// Set byte if greater or equal (SF=OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETGE(rm interface{}) {
	unsafe.Asm("SETGE", rm)
}

// SETL
// Set byte if less (SF≠ OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETL(rm interface{}) {
	unsafe.Asm("SETL", rm)
}

// SETLE
// Set byte if less or equal (ZF=1 or SF≠ OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETLE(rm interface{}) {
	unsafe.Asm("SETLE", rm)
}

// SETNA
// Set byte if not above (CF=1 or ZF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNA(rm interface{}) {
	unsafe.Asm("SETNA", rm)
}

// SETNAE
// Set byte if not above or equal (CF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNAE(rm interface{}) {
	unsafe.Asm("SETNAE", rm)
}

// SETNB
// Set byte if not below (CF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNB(rm interface{}) {
	unsafe.Asm("SETNB", rm)
}

// SETNBE
// Set byte if not below or equal (CF=0 and ZF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNBE(rm interface{}) {
	unsafe.Asm("SETNBE", rm)
}

// SETNC
// Set byte if not carry (CF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNC(rm interface{}) {
	unsafe.Asm("SETNC", rm)
}

// SETNE
// Set byte if not equal (ZF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNE(rm interface{}) {
	unsafe.Asm("SETNE", rm)
}

// SETNG
// Set byte if not greater (ZF=1 or SF≠ OF)
// Set byte if not greater (ZF=1 or SF≠ OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNG(rm interface{}) {
	unsafe.Asm("SETNG", rm)
}

// SETNGE
// Set byte if not greater or equal (SF≠ OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNGE(rm interface{}) {
	unsafe.Asm("SETNGE", rm)
}

// SETNL
// Set byte if not less (SF=OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNL(rm interface{}) {
	unsafe.Asm("SETNL", rm)
}

// SETNLE
// Set byte if not less or equal (ZF=0 and SF=OF).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNLE(rm interface{}) {
	unsafe.Asm("SETNLE", rm)
}

// SETNO
// Set byte if not overflow (OF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNO(rm interface{}) {
	unsafe.Asm("SETNO", rm)
}

// SETNP
// Set byte if not parity (PF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNP(rm interface{}) {
	unsafe.Asm("SETNP", rm)
}

// SETNS
// Set byte if not sign (SF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNS(rm interface{}) {
	unsafe.Asm("SETNS", rm)
}

// SETNZ
// Set byte if not zero (ZF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETNZ(rm interface{}) {
	unsafe.Asm("SETNZ", rm)
}

// SETO
// Set byte if overflow (OF=1)
// Set byte if overflow (OF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETO(rm interface{}) {
	unsafe.Asm("SETO", rm)
}

// SETP
// Set byte if parity (PF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETP(rm interface{}) {
	unsafe.Asm("SETP", rm)
}

// SETPE
// Set byte if parity even (PF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETPE(rm interface{}) {
	unsafe.Asm("SETPE", rm)
}

// SETPO
// Set byte if parity odd (PF=0).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETPO(rm interface{}) {
	unsafe.Asm("SETPO", rm)
}

// SETS
// Set byte if sign (SF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETS(rm interface{}) {
	unsafe.Asm("SETS", rm)
}

// SETZ
// Set byte if zero (ZF=1).
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1248
func SETZ(rm interface{}) {
	unsafe.Asm("SETZ", rm)
}

// SFENCE
// Serializes store operations.
//
// Documentation: https://golang.org/s/x86manual#page=1251
func SFENCE() {
	unsafe.Asm("SFENCE", nil)
}

// SGDT
// Store GDTR to m.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1252
func SGDT(rm interface{}) {
	unsafe.Asm("SGDT", rm)
}

// SHA1MSG1
// Performs an intermediate calculation for the next four SHA1 message dwords using previous message dwords from xmm1 and xmm2/m128, storing the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1257
func SHA1MSG1(reg, rm interface{}) {
	unsafe.Asm("SHA1MSG1", reg, rm)
}

// SHA1MSG2
// Performs the final calculation for the next four SHA1 message dwords using intermediate results from xmm1 and the previous message dwords from xmm2/m128, storing the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1258
func SHA1MSG2(reg, rm interface{}) {
	unsafe.Asm("SHA1MSG2", reg, rm)
}

// SHA1NEXTE
// Calculates SHA1 state variable E after four rounds of operation from the current SHA1 state variable A in xmm1. The calculated value of the SHA1 state variable E is added to the scheduled dwords in xmm2/m128, and stored with some of the scheduled dwords in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1256
func SHA1NEXTE(reg, rm interface{}) {
	unsafe.Asm("SHA1NEXTE", reg, rm)
}

// SHA1RNDS4
// Performs four rounds of SHA1 operation operating on SHA1 state (A,B,C,D) from xmm1, with a pre-computed sum of the next 4 round message dwords and state variable E from xmm2/m128. The immediate byte controls logic functions and round constants.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1254
func SHA1RNDS4(reg, rm, imm interface{}) {
	unsafe.Asm("SHA1RNDS4", reg, rm, imm)
}

// SHA256MSG1
// Performs an intermediate calculation for the next four SHA256 message dwords using previous message dwords from xmm1 and xmm2/m128, storing the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1261
func SHA256MSG1(reg, rm interface{}) {
	unsafe.Asm("SHA256MSG1", reg, rm)
}

// SHA256MSG2
// Performs the final calculation for the next four SHA256 message dwords using previous message dwords from xmm1 and xmm2/m128, storing the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1262
func SHA256MSG2(reg, rm interface{}) {
	unsafe.Asm("SHA256MSG2", reg, rm)
}

// SHA256RNDS2
// Perform 2 rounds of SHA256 operation using an initial SHA256 state (C,D,G,H) from xmm1, an initial SHA256 state (A,B,E,F) from xmm2/m128, and a pre-computed sum of the next 2 round mes- sage dwords and the corresponding round constants from the implicit operand XMM0, storing the updated SHA256 state (A,B,E,F) result in xmm1.
//
// Documentation: https://golang.org/s/x86manual#page=1259
func SHA256RNDS2() {
	unsafe.Asm("SHA256RNDS2", nil)
}

// SHL_M1
// Multiply r/m16 by 2, once.
// Multiply r/m32 by 2, once.
// Multiply r/m64 by 2, once.
// Multiply r/m8 by 2, once.
//
// rm: ModRM:r/m (r, w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHL_M1(rm, v1 interface{}) {
	unsafe.Asm("SHL", rm, v1)
}

// SHL_MC
// Multiply r/m16 by 2, CL times.
// Multiply r/m32 by 2, CL times.
// Multiply r/m64 by 2, CL times.
// Multiply r/m8 by 2, CL times.
//
// rm: ModRM:r/m (r, w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHL_MC(rm, cl interface{}) {
	unsafe.Asm("SHL", rm, cl)
}

// SHL_MI
// Multiply r/m16 by 2, imm8 times.
// Multiply r/m32 by 2, imm8 times.
// Multiply r/m64 by 2, imm8 times.
// Multiply r/m8 by 2, imm8 times.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHL_MI(rm, imm interface{}) {
	unsafe.Asm("SHL", rm, imm)
}

// SHLD_MRC
// Shift r/m16 to left CL places while shifting bits from r16 in from the right.
// Shift r/m32 to left CL places while shifting bits from r32 in from the right.
// Shift r/m64 to left CL places while shifting bits from r64 in from the right.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1263
func SHLD_MRC(rm, reg, cl interface{}) {
	unsafe.Asm("SHLD", rm, reg, cl)
}

// SHLD_MRI
// Shift r/m16 to left imm8 places while shifting bits from r16 in from the right.
// Shift r/m32 to left imm8 places while shifting bits from r32 in from the right.
// Shift r/m64 to left imm8 places while shifting bits from r64 in from the right.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1263
func SHLD_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("SHLD", rm, reg, imm)
}

// SHLX
// Shift r/m32 logically left with count specified in r32b.
// Shift r/m64 logically left with count specified in r64b.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// vex: VEX.vvvv (r)
//
// Documentation: https://golang.org/s/x86manual#page=1239
func SHLX(reg, rm, vex interface{}) {
	unsafe.Asm("SHLX", reg, rm, vex)
}

// SHR_M1
// Unsigned divide r/m16 by 2, once.
// Unsigned divide r/m32 by 2, once.
// Unsigned divide r/m64 by 2, once.
// Unsigned divide r/m8 by 2, once.
//
// rm: ModRM:r/m (r, w)
// v1: 1
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHR_M1(rm, v1 interface{}) {
	unsafe.Asm("SHR", rm, v1)
}

// SHR_MC
// Unsigned divide r/m16 by 2, CL times
// Unsigned divide r/m32 by 2, CL times.
// Unsigned divide r/m64 by 2, CL times.
// Unsigned divide r/m8 by 2, CL times.
//
// rm: ModRM:r/m (r, w)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHR_MC(rm, cl interface{}) {
	unsafe.Asm("SHR", rm, cl)
}

// SHR_MI
// Unsigned divide r/m16 by 2, imm8 times.
// Unsigned divide r/m32 by 2, imm8 times.
// Unsigned divide r/m64 by 2, imm8 times.
// Unsigned divide r/m8 by 2, imm8 times.
//
// rm: ModRM:r/m (r, w)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1234
func SHR_MI(rm, imm interface{}) {
	unsafe.Asm("SHR", rm, imm)
}

// SHRD_MRC
// Shift r/m16 to right CL places while shifting bits from r16 in from the left.
// Shift r/m32 to right CL places while shifting bits from r32 in from the left.
// Shift r/m64 to right CL places while shifting bits from r64 in from the left.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// cl: CL
//
// Documentation: https://golang.org/s/x86manual#page=1266
func SHRD_MRC(rm, reg, cl interface{}) {
	unsafe.Asm("SHRD", rm, reg, cl)
}

// SHRD_MRI
// Shift r/m16 to right imm8 places while shifting bits from r16 in from the left.
// Shift r/m32 to right imm8 places while shifting bits from r32 in from the left.
// Shift r/m64 to right imm8 places while shifting bits from r64 in from the left.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1266
func SHRD_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("SHRD", rm, reg, imm)
}

// SHRX
// Shift r/m32 logically right with count specified in r32b.
// Shift r/m64 logically right with count specified in r64b.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// vex: VEX.vvvv (r)
//
// Documentation: https://golang.org/s/x86manual#page=1239
func SHRX(reg, rm, vex interface{}) {
	unsafe.Asm("SHRX", reg, rm, vex)
}

// SHUFPD
// Shuffle two pairs of double-precision floating-point values from xmm1 and xmm2/m128 using imm8 to select from each pair, interleaved result is stored in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1269
func SHUFPD(reg, rm, imm interface{}) {
	unsafe.Asm("SHUFPD", reg, rm, imm)
}

// SHUFPS
// Select from quadruplet of single-precision floating- point values in xmm1 and xmm2/m128 using imm8, interleaved result pairs are stored in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1274
func SHUFPS(reg, rm, imm interface{}) {
	unsafe.Asm("SHUFPS", reg, rm, imm)
}

// SIDT
// Store IDTR to m.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1278
func SIDT(rm interface{}) {
	unsafe.Asm("SIDT", rm)
}

// SLDT
// Stores segment selector from LDTR in r/m16.
// Stores segment selector from LDTR in r64/m16.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1280
func SLDT(rm interface{}) {
	unsafe.Asm("SLDT", rm)
}

// SMSW
// Store machine status word to r/m16.
// Store machine status word in low-order 16 bits of r32/m16; high-order 16 bits of r32 are undefined.
// Store machine status word in low-order 16 bits of r64/m16; high-order 16 bits of r32 are undefined.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1282
func SMSW(rm interface{}) {
	unsafe.Asm("SMSW", rm)
}

// SQRTPD
// Computes Square Roots of the packed double-precision floating-point values in xmm2/m128 and stores the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1284
func SQRTPD(reg, rm interface{}) {
	unsafe.Asm("SQRTPD", reg, rm)
}

// SQRTPS
// Computes Square Roots of the packed single-precision floating-point values in xmm2/m128 and stores the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1287
func SQRTPS(reg, rm interface{}) {
	unsafe.Asm("SQRTPS", reg, rm)
}

// SQRTSD
// Computes square root of the low double-precision floating- point value in xmm2/m64 and stores the results in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1290
func SQRTSD(reg, rm interface{}) {
	unsafe.Asm("SQRTSD", reg, rm)
}

// SQRTSS
// Computes square root of the low single-precision floating-point value in xmm2/m32 and stores the results in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1292
func SQRTSS(reg, rm interface{}) {
	unsafe.Asm("SQRTSS", reg, rm)
}

// STAC
// Set the AC flag in the EFLAGS register.
//
// Documentation: https://golang.org/s/x86manual#page=1294
func STAC() {
	unsafe.Asm("STAC", nil)
}

// STC
// Set CF flag.
//
// Documentation: https://golang.org/s/x86manual#page=1295
func STC() {
	unsafe.Asm("STC", nil)
}

// STD
// Set DF flag.
//
// Documentation: https://golang.org/s/x86manual#page=1296
func STD() {
	unsafe.Asm("STD", nil)
}

// STI
// Set interrupt flag; external, maskable interrupts enabled at the end of the next instruction.
//
// Documentation: https://golang.org/s/x86manual#page=1297
func STI() {
	unsafe.Asm("STI", nil)
}

// STMXCSR
// Store contents of MXCSR register to m32.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1299
func STMXCSR(rm interface{}) {
	unsafe.Asm("STMXCSR", rm)
}

// STOSB
// For legacy mode, store AL at address ES:(E)DI; For 64-bit mode store AL at address RDI or EDI.
//
// Documentation: https://golang.org/s/x86manual#page=1300
func STOSB() {
	unsafe.Asm("STOSB", nil)
}

// STOSD
// For legacy mode, store EAX at address ES:(E)DI; For 64-bit mode store EAX at address RDI or EDI.
//
// Documentation: https://golang.org/s/x86manual#page=1300
func STOSD() {
	unsafe.Asm("STOSD", nil)
}

// STOSQ
// Store RAX at address RDI or EDI.
//
// Documentation: https://golang.org/s/x86manual#page=1300
func STOSQ() {
	unsafe.Asm("STOSQ", nil)
}

// STOSW
// For legacy mode, store AX at address ES:(E)DI; For 64-bit mode store AX at address RDI or EDI.
//
// Documentation: https://golang.org/s/x86manual#page=1300
func STOSW() {
	unsafe.Asm("STOSW", nil)
}

// STR
// Stores segment selector from TR in r/m16.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1304
func STR(rm interface{}) {
	unsafe.Asm("STR", rm)
}

// SUB_I
// Subtract imm8 from AL.
// Subtract imm16 from AX.
// Subtract imm32 from EAX.
// Subtract imm32 sign-extended to 64-bits from RAX.
//
// al: AL/AX/EAX/RAX
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1306
func SUB_I(al, imm interface{}) {
	unsafe.Asm("SUB", al, imm)
}

// SUB_MI
// Subtract imm16 from r/m16.
// Subtract sign-extended imm8 from r/m16.
// Subtract imm32 from r/m32.
// Subtract sign-extended imm8 from r/m32.
// Subtract imm32 sign-extended to 64-bits from r/m64.
// Subtract sign-extended imm8 from r/m64.
// Subtract imm8 from r/m8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1306
func SUB_MI(rm, imm interface{}) {
	unsafe.Asm("SUB", rm, imm)
}

// SUB_MR
// Subtract r16 from r/m16.
// Subtract r32 from r/m32.
// Subtract r64 from r/m64.
// Subtract r8 from r/m8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1306
func SUB_MR(rm, reg interface{}) {
	unsafe.Asm("SUB", rm, reg)
}

// SUB_RM
// Subtract r/m16 from r16.
// Subtract r/m32 from r32.
// Subtract r/m64 from r64.
// Subtract r/m8 from r8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1306
func SUB_RM(reg, rm interface{}) {
	unsafe.Asm("SUB", reg, rm)
}

// SUBPD
// Subtract packed double-precision floating-point values in xmm2/mem from xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1308
func SUBPD(reg, rm interface{}) {
	unsafe.Asm("SUBPD", reg, rm)
}

// SUBPS
// Subtract packed single-precision floating-point values in xmm2/mem from xmm1 and store result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1311
func SUBPS(reg, rm interface{}) {
	unsafe.Asm("SUBPS", reg, rm)
}

// SUBSD
// Subtract the low double-precision floating-point value in xmm2/m64 from xmm1 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1314
func SUBSD(reg, rm interface{}) {
	unsafe.Asm("SUBSD", reg, rm)
}

// SUBSS
// Subtract the low single-precision floating-point value in xmm2/m32 from xmm1 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1316
func SUBSS(reg, rm interface{}) {
	unsafe.Asm("SUBSS", reg, rm)
}

// SWAPGS
// Exchanges the current GS base register value with the value contained in MSR address C0000102H.
//
// Documentation: https://golang.org/s/x86manual#page=1318
func SWAPGS() {
	unsafe.Asm("SWAPGS", nil)
}

// SYSCALL
// Fast call to privilege level 0 system procedures.
//
// Documentation: https://golang.org/s/x86manual#page=1320
func SYSCALL() {
	unsafe.Asm("SYSCALL", nil)
}

// SYSENTER
// Fast call to privilege level 0 system procedures.
//
// Documentation: https://golang.org/s/x86manual#page=1322
func SYSENTER() {
	unsafe.Asm("SYSENTER", nil)
}

// SYSEXIT
// Fast return to privilege level 3 user code.
// Fast return to 64-bit mode privilege level 3 user code.
//
// Documentation: https://golang.org/s/x86manual#page=1325
func SYSEXIT() {
	unsafe.Asm("SYSEXIT", nil)
}

// SYSRET
// Return to compatibility mode from fast system call
// Return to 64-bit mode from fast system call
//
// Documentation: https://golang.org/s/x86manual#page=1328
func SYSRET() {
	unsafe.Asm("SYSRET", nil)
}

// TEST_I
// AND imm8 with AL; set SF, ZF, PF according to result.
// AND imm16 with AX; set SF, ZF, PF according to result.
// AND imm32 with EAX; set SF, ZF, PF according to result.
// AND imm32 sign-extended to 64-bits with RAX; set SF, ZF, PF according to result.
//
// al: AL/AX/EAX/RAX
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1331
func TEST_I(al, imm interface{}) {
	unsafe.Asm("TEST", al, imm)
}

// TEST_MI
// AND imm16 with r/m16; set SF, ZF, PF according to result.
// AND imm32 with r/m32; set SF, ZF, PF according to result.
// AND imm32 sign-extended to 64-bits with r/m64; set SF, ZF, PF according to result.
// AND imm8 with r/m8; set SF, ZF, PF according to result.
//
// rm: ModRM:r/m (r)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1331
func TEST_MI(rm, imm interface{}) {
	unsafe.Asm("TEST", rm, imm)
}

// TEST_MR
// AND r16 with r/m16; set SF, ZF, PF according to result.
// AND r32 with r/m32; set SF, ZF, PF according to result.
// AND r64 with r/m64; set SF, ZF, PF according to result.
// AND r8 with r/m8; set SF, ZF, PF according to result.
//
// rm: ModRM:r/m (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1331
func TEST_MR(rm, reg interface{}) {
	unsafe.Asm("TEST", rm, reg)
}

// TZCNT
// Count the number of trailing zero bits in r/m16, return result in r16.
// Count the number of trailing zero bits in r/m32, return result in r32.
// Count the number of trailing zero bits in r/m64, return result in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1333
func TZCNT(reg, rm interface{}) {
	unsafe.Asm("TZCNT", reg, rm)
}

// UCOMISD
// Compare low double-precision floating-point values in xmm1 and xmm2/mem64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1335
func UCOMISD(reg, rm interface{}) {
	unsafe.Asm("UCOMISD", reg, rm)
}

// UCOMISS
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1337
func UCOMISS(reg, rm interface{}) {
	unsafe.Asm("UCOMISS", reg, rm)
}

// UD2
// Raise invalid opcode exception.
//
// Documentation: https://golang.org/s/x86manual#page=1339
func UD2() {
	unsafe.Asm("UD2", nil)
}

// UNPCKHPD
// Unpacks and Interleaves double-precision floating-point values from high quadwords of xmm1 and xmm2/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1340
func UNPCKHPD() {
	unsafe.Asm("UNPCKHPD", nil)
}

// UNPCKHPS
// Unpacks and Interleaves single-precision floating-point values from high quadwords of xmm1 and xmm2/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1344
func UNPCKHPS() {
	unsafe.Asm("UNPCKHPS", nil)
}

// UNPCKLPD
// Unpacks and Interleaves double-precision floating-point values from low quadwords of xmm1 and xmm2/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1348
func UNPCKLPD() {
	unsafe.Asm("UNPCKLPD", nil)
}

// UNPCKLPS
// Unpacks and Interleaves single-precision floating-point values from low quadwords of xmm1 and xmm2/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1352
func UNPCKLPS() {
	unsafe.Asm("UNPCKLPS", nil)
}

// VADDPD_FV
// Add packed double-precision floating-point values from xmm3/m128/m64bcst to xmm2 and store result in xmm1 with writemask k1.
// Add packed double-precision floating-point values from ymm3/m256/m64bcst to ymm2 and store result in ymm1 with writemask k1.
// Add packed double-precision floating-point values from zmm3/m512/m64bcst to zmm2 and store result in zmm1 with writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=135
func VADDPD_FV() {
	unsafe.Asm("VADDPD", nil)
}

// VADDPD_RVM
// Add packed double-precision floating-point values from xmm3/mem to xmm2 and store result in xmm1.
// Add packed double-precision floating-point values from ymm3/mem to ymm2 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=135
func VADDPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VADDPD", reg, vex, rm)
}

// VADDPS_FV
// Add packed single-precision floating-point values from xmm3/m128/m32bcst to xmm2 and store result in xmm1 with writemask k1.
// Add packed single-precision floating-point values from ymm3/m256/m32bcst to ymm2 and store result in ymm1 with writemask k1.
// Add packed single-precision floating-point values from zmm3/m512/m32bcst to zmm2 and store result in zmm1 with writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=138
func VADDPS_FV() {
	unsafe.Asm("VADDPS", nil)
}

// VADDPS_RVM
// Add packed single-precision floating-point values from xmm3/m128 to xmm2 and store result in xmm1.
// Add packed single-precision floating-point values from ymm3/m256 to ymm2 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=138
func VADDPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VADDPS", reg, vex, rm)
}

// VADDSD_RVM
// Add the low double-precision floating-point value from xmm3/mem to xmm2 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=141
func VADDSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VADDSD", reg, vex, rm)
}

// VADDSD_T1S
// Add the low double-precision floating-point value from xmm3/m64 to xmm2 and store the result in xmm1 with writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=141
func VADDSD_T1S() {
	unsafe.Asm("VADDSD", nil)
}

// VADDSS_RVM
// Add the low single-precision floating-point value from xmm3/mem to xmm2 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=143
func VADDSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VADDSS", reg, vex, rm)
}

// VADDSS_T1S
// Add the low single-precision floating-point value from xmm3/m32 to xmm2 and store the result in xmm1with writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=143
func VADDSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VADDSS", reg, evex, rm)
}

// VADDSUBPD
// Add/subtract packed double-precision floating-point values from xmm3/mem to xmm2 and stores result in xmm1.
// Add / subtract packed double-precision floating-point values from ymm3/mem to ymm2 and stores result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=145
func VADDSUBPD(reg, vex, rm interface{}) {
	unsafe.Asm("VADDSUBPD", reg, vex, rm)
}

// VADDSUBPS
// Add/subtract single-precision floating-point values from xmm3/mem to xmm2 and stores result in xmm1.
// Add / subtract single-precision floating-point values from ymm3/mem to ymm2 and stores result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=147
func VADDSUBPS(reg, vex, rm interface{}) {
	unsafe.Asm("VADDSUBPS", reg, vex, rm)
}

// VAESDEC
// Perform one round of an AES decryption flow, using the Equivalent Inverse Cipher, operating on a 128-bit data (state) from xmm2 with a 128-bit round key from xmm3/m128; store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=152
func VAESDEC(reg, vex, rm interface{}) {
	unsafe.Asm("VAESDEC", reg, vex, rm)
}

// VAESDECLAST
// Perform the last round of an AES decryption flow, using the Equivalent Inverse Cipher, operating on a 128-bit data (state) from xmm2 with a 128-bit round key from xmm3/m128; store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=154
func VAESDECLAST(reg, vex, rm interface{}) {
	unsafe.Asm("VAESDECLAST", reg, vex, rm)
}

// VAESENC
// Perform one round of an AES encryption flow, operating on a 128-bit data (state) from xmm2 with a 128-bit round key from the xmm3/m128; store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=156
func VAESENC(reg, vex, rm interface{}) {
	unsafe.Asm("VAESENC", reg, vex, rm)
}

// VAESENCLAST
// Perform the last round of an AES encryption flow, operating on a 128-bit data (state) from xmm2 with a 128 bit round key from xmm3/m128; store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=158
func VAESENCLAST(reg, vex, rm interface{}) {
	unsafe.Asm("VAESENCLAST", reg, vex, rm)
}

// VAESIMC
// Perform the InvMixColumn transformation on a 128-bit round key from xmm2/m128 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=160
func VAESIMC(reg, rm interface{}) {
	unsafe.Asm("VAESIMC", reg, rm)
}

// VAESKEYGENASSIST
// Assist in AES round key generation using 8 bits Round Constant (RCON) specified in the immediate byte, operating on 128 bits of data specified in xmm2/m128 and stores the result in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=161
func VAESKEYGENASSIST(reg, rm, imm interface{}) {
	unsafe.Asm("VAESKEYGENASSIST", reg, rm, imm)
}

// VANDNPD_FV
// Return the bitwise logical AND NOT of packed double- precision floating-point values in xmm2 and xmm3/m128/m64bcst subject to writemask k1.
// Return the bitwise logical AND NOT of packed double- precision floating-point values in ymm2 and ymm3/m256/m64bcst subject to writemask k1.
// Return the bitwise logical AND NOT of packed double- precision floating-point values in zmm2 and zmm3/m512/m64bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=172
func VANDNPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VANDNPD", reg, evex, rm)
}

// VANDNPD_RVM
// Return the bitwise logical AND NOT of packed double- precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical AND NOT of packed double- precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=172
func VANDNPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VANDNPD", reg, vex, rm)
}

// VANDNPS_FV
// Return the bitwise logical AND of packed single-precision floating-point values in xmm2 and xmm3/m128/m32bcst subject to writemask k1.
// Return the bitwise logical AND of packed single-precision floating-point values in ymm2 and ymm3/m256/m32bcst subject to writemask k1.
// Return the bitwise logical AND of packed single-precision floating-point values in zmm2 and zmm3/m512/m32bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=175
func VANDNPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VANDNPS", reg, evex, rm)
}

// VANDNPS_RVM
// Return the bitwise logical AND NOT of packed single-precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical AND NOT of packed single-precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=175
func VANDNPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VANDNPS", reg, vex, rm)
}

// VANDPD_FV
// Return the bitwise logical AND of packed double- precision floating-point values in xmm2 and xmm3/m128/m64bcst subject to writemask k1.
// Return the bitwise logical AND of packed double- precision floating-point values in ymm2 and ymm3/m256/m64bcst subject to writemask k1.
// Return the bitwise logical AND of packed double- precision floating-point values in zmm2 and zmm3/m512/m64bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=166
func VANDPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VANDPD", reg, evex, rm)
}

// VANDPD_RVM
// Return the bitwise logical AND of packed double- precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical AND of packed double- precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=166
func VANDPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VANDPD", reg, vex, rm)
}

// VANDPS_FV
// Return the bitwise logical AND of packed single-precision floating-point values in xmm2 and xmm3/m128/m32bcst subject to writemask k1.
// Return the bitwise logical AND of packed single-precision floating-point values in ymm2 and ymm3/m256/m32bcst subject to writemask k1.
// Return the bitwise logical AND of packed single-precision floating-point values in zmm2 and zmm3/m512/m32bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=169
func VANDPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VANDPS", reg, evex, rm)
}

// VANDPS_RVM
// Return the bitwise logical AND of packed single-precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical AND of packed single-precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=169
func VANDPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VANDPS", reg, vex, rm)
}

// VBLENDMPD
// Blend double-precision vector xmm2 and double-precision vector xmm3/m128/m64bcst and store the result in xmm1, under control mask.
// Blend double-precision vector ymm2 and double-precision vector ymm3/m256/m64bcst and store the result in ymm1, under control mask.
// Blend double-precision vector zmm2 and double-precision vector zmm3/m512/m64bcst and store the result in zmm1, under control mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1365
func VBLENDMPD(reg, evex, rm interface{}) {
	unsafe.Asm("VBLENDMPD", reg, evex, rm)
}

// VBLENDMPS
// Blend single-precision vector xmm2 and single-precision vector xmm3/m128/m32bcst and store the result in xmm1, under control mask.
// Blend single-precision vector ymm2 and single-precision vector ymm3/m256/m32bcst and store the result in ymm1, under control mask.
// Blend single-precision vector zmm2 and single-precision vector zmm3/m512/m32bcst using k1 as select control and store the result in zmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1365
func VBLENDMPS(reg, evex, rm interface{}) {
	unsafe.Asm("VBLENDMPS", reg, evex, rm)
}

// VBLENDPS
// Select packed single-precision floating-point values from xmm2 and xmm3/m128 from mask in imm8 and store the values in xmm1.
// Select packed single-precision floating-point values from ymm2 and ymm3/m256 from mask in imm8 and store the values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=183
func VBLENDPS(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VBLENDPS", reg, vex, rm, imm)
}

// VBLENDVPD
// Conditionally copy double-precision floating- point values from xmm2 or xmm3/m128 to xmm1, based on mask bits in the mask operand, xmm4.
// Conditionally copy double-precision floating- point values from ymm2 or ymm3/m256 to ymm1, based on mask bits in the mask operand, ymm4.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8[7:4]
//
// Documentation: https://golang.org/s/x86manual#page=185
func VBLENDVPD(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VBLENDVPD", reg, vex, rm, imm)
}

// VBLENDVPS
// Conditionally copy single-precision floating- point values from xmm2 or xmm3/m128 to xmm1, based on mask bits in the specified mask operand, xmm4.
// Conditionally copy single-precision floating- point values from ymm2 or ymm3/m256 to ymm1, based on mask bits in the specified mask register, ymm4.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8[7:4]
//
// Documentation: https://golang.org/s/x86manual#page=187
func VBLENDVPS(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VBLENDVPS", reg, vex, rm, imm)
}

// VBROADCASTI128
// Broadcast 128 bits of integer data in mem to low and high 128-bits in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI128(reg, rm interface{}) {
	unsafe.Asm("VBROADCASTI128", reg, rm)
}

// VBROADCASTI32X4
// Broadcast 128 bits of 4 doubleword integer data in mem to locations in ymm1 using writemask k1.
// Broadcast 128 bits of 4 doubleword integer data in mem to locations in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI32X4() {
	unsafe.Asm("VBROADCASTI32X4", nil)
}

// VBROADCASTI32X8
// Broadcast 256 bits of 8 doubleword integer data in mem to locations in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI32X8() {
	unsafe.Asm("VBROADCASTI32X8", nil)
}

// VBROADCASTI32x2
// Broadcast two dword elements in source operand to locations in xmm1 subject to writemask k1.
// Broadcast two dword elements in source operand to locations in ymm1 subject to writemask k1.
// Broadcast two dword elements in source operand to locations in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI32x2() {
	unsafe.Asm("VBROADCASTI32x2", nil)
}

// VBROADCASTI64X2
// Broadcast 128 bits of 2 quadword integer data in mem to locations in ymm1 using writemask k1.
// Broadcast 128 bits of 2 quadword integer data in mem to locations in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI64X2() {
	unsafe.Asm("VBROADCASTI64X2", nil)
}

// VBROADCASTI64X4
// Broadcast 256 bits of 4 quadword integer data in mem to locations in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VBROADCASTI64X4() {
	unsafe.Asm("VBROADCASTI64X4", nil)
}

// VCMPPD_FV
// Compare packed double-precision floating-point values in xmm3/m128/m64bcst and xmm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed double-precision floating-point values in ymm3/m256/m64bcst and ymm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed double-precision floating-point values in zmm3/m512/m64bcst and zmm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=257
func VCMPPD_FV(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VCMPPD", reg, evex, rm, imm)
}

// VCMPPD_RVMI
// Compare packed double-precision floating-point values in xmm3/m128 and xmm2 using bits 4:0 of imm8 as a comparison predicate.
// Compare packed double-precision floating-point values in ymm3/m256 and ymm2 using bits 4:0 of imm8 as a comparison predicate.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=257
func VCMPPD_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VCMPPD", reg, vex, rm, imm)
}

// VCMPPS_FV
// Compare packed single-precision floating-point values in xmm3/m128/m32bcst and xmm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed single-precision floating-point values in ymm3/m256/m32bcst and ymm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed single-precision floating-point values in zmm3/m512/m32bcst and zmm2 using bits 4:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=264
func VCMPPS_FV(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VCMPPS", reg, evex, rm, imm)
}

// VCMPPS_RVMI
// Compare packed single-precision floating-point values in xmm3/m128 and xmm2 using bits 4:0 of imm8 as a comparison predicate.
// Compare packed single-precision floating-point values in ymm3/m256 and ymm2 using bits 4:0 of imm8 as a comparison predicate.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=264
func VCMPPS_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VCMPPS", reg, vex, rm, imm)
}

// VCMPSD_RVMI
// Compare low double-precision floating-point value in xmm3/m64 and xmm2 using bits 4:0 of imm8 as comparison predicate.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=275
func VCMPSD_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VCMPSD", reg, vex, rm, imm)
}

// VCMPSD_T1S
// Compare low double-precision floating-point value in xmm3/m64 and xmm2 using bits 4:0 of imm8 as comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=275
func VCMPSD_T1S(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VCMPSD", reg, evex, rm, imm)
}

// VCMPSS_RVMI
// Compare low single-precision floating-point value in xmm3/m32 and xmm2 using bits 4:0 of imm8 as comparison predicate.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=279
func VCMPSS_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VCMPSS", reg, vex, rm, imm)
}

// VCMPSS_T1S
// Compare low single-precision floating-point value in xmm3/m32 and xmm2 using bits 4:0 of imm8 as comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=279
func VCMPSS_T1S(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VCMPSS", reg, evex, rm, imm)
}

// VCOMISD_RM
// Compare low double-precision floating-point values in xmm1 and xmm2/mem64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=288
func VCOMISD_RM(reg, rm interface{}) {
	unsafe.Asm("VCOMISD", reg, rm)
}

// VCOMISD_T1S
// Compare low double-precision floating-point values in xmm1 and xmm2/mem64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=288
func VCOMISD_T1S(reg, rm interface{}) {
	unsafe.Asm("VCOMISD", reg, rm)
}

// VCOMISS_RM
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=290
func VCOMISS_RM(reg, rm interface{}) {
	unsafe.Asm("VCOMISS", reg, rm)
}

// VCOMISS_T1S
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=290
func VCOMISS_T1S(reg, rm interface{}) {
	unsafe.Asm("VCOMISS", reg, rm)
}

// VCOMPRESSPD
// Compress packed double-precision floating-point values from xmm2 to xmm1/m128 using writemask k1.
// Compress packed double-precision floating-point values from ymm2 to ymm1/m256 using writemask k1.
// Compress packed double-precision floating-point values from zmm2 using control mask k1 to zmm1/m512.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1377
func VCOMPRESSPD(rm, reg interface{}) {
	unsafe.Asm("VCOMPRESSPD", rm, reg)
}

// VCOMPRESSPS
// Compress packed single-precision floating-point values from xmm2 to xmm1/m128 using writemask k1.
// Compress packed single-precision floating-point values from ymm2 to ymm1/m256 using writemask k1.
// Compress packed single-precision floating-point values from zmm2 using control mask k1 to zmm1/m512.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1379
func VCOMPRESSPS(rm, reg interface{}) {
	unsafe.Asm("VCOMPRESSPS", rm, reg)
}

// VCVTDQ2PD_HV
// Convert 2 packed signed doubleword integers from xmm2/m128/m32bcst to eight packed double-precision floating-point values in xmm1 with writemask k1.
// Convert 4 packed signed doubleword integers from xmm2/m128/m32bcst to 4 packed double-precision floating-point values in ymm1 with writemask k1.
// Convert eight packed signed doubleword integers from ymm2/m256/m32bcst to eight packed double-precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=330
func VCVTDQ2PD_HV(reg, rm interface{}) {
	unsafe.Asm("VCVTDQ2PD", reg, rm)
}

// VCVTDQ2PD_RM
// Convert two packed signed doubleword integers from xmm2/mem to two packed double-precision floating- point values in xmm1.
// Convert four packed signed doubleword integers from xmm2/mem to four packed double-precision floating- point values in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=330
func VCVTDQ2PD_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTDQ2PD", reg, rm)
}

// VCVTPD2PS_FV
// Convert two packed double-precision floating-point values in xmm2/m128/m64bcst to two single- precision floating-point values in xmm1with writemask k1.
// Convert four packed double-precision floating-point values in ymm2/m256/m64bcst to four single- precision floating-point values in xmm1with writemask k1.
// Convert eight packed double-precision floating-point values in zmm2/m512/m64bcst to eight single- precision floating-point values in ymm1with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=342
func VCVTPD2PS_FV(reg, rm interface{}) {
	unsafe.Asm("VCVTPD2PS", reg, rm)
}

// VCVTPD2PS_RM
// Convert two packed double-precision floating-point values in xmm2/mem to two single-precision floating-point values in xmm1.
// Convert four packed double-precision floating-point values in ymm2/mem to four single-precision floating-point values in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=342
func VCVTPD2PS_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTPD2PS", reg, rm)
}

// VCVTPD2QQ
// Convert two packed double-precision floating-point values from xmm2/m128/m64bcst to two packed quadword integers in xmm1 with writemask k1.
// Convert four packed double-precision floating-point values from ymm2/m256/m64bcst to four packed quadword integers in ymm1 with writemask k1.
// Convert eight packed double-precision floating-point values from zmm2/m512/m64bcst to eight packed quadword integers in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1381
func VCVTPD2QQ(reg, rm interface{}) {
	unsafe.Asm("VCVTPD2QQ", reg, rm)
}

// VCVTPD2UQQ
// Convert two packed double-precision floating-point values from xmm2/mem to two packed unsigned quadword integers in xmm1 with writemask k1.
// Convert fourth packed double-precision floating-point values from ymm2/mem to four packed unsigned quadword integers in ymm1 with writemask k1.
// Convert eight packed double-precision floating-point values from zmm2/mem to eight packed unsigned quadword integers in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1387
func VCVTPD2UQQ(reg, rm interface{}) {
	unsafe.Asm("VCVTPD2UQQ", reg, rm)
}

// VCVTPH2PS_HVM
// Convert four packed half precision (16-bit) floating- point values in xmm2/m64 to packed single-precision floating-point values in xmm1.
// Convert eight packed half precision (16-bit) floating- point values in xmm2/m128 to packed single- precision floating-point values in ymm1.
// Convert sixteen packed half precision (16-bit) floating-point values in ymm2/m256 to packed single-precision floating-point values in zmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1390
func VCVTPH2PS_HVM(reg, rm interface{}) {
	unsafe.Asm("VCVTPH2PS", reg, rm)
}

// VCVTPH2PS_RM
// Convert four packed half precision (16-bit) floating- point values in xmm2/m64 to packed single-precision floating-point value in xmm1.
// Convert eight packed half precision (16-bit) floating- point values in xmm2/m128 to packed single- precision floating-point value in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1390
func VCVTPH2PS_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTPH2PS", reg, rm)
}

// VCVTPS2DQ_FV
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed signed doubleword values in xmm1 subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed signed doubleword values in ymm1 subject to writemask k1.
// Convert sixteen packed single-precision floating-point values from zmm2/m512/m32bcst to sixteen packed signed doubleword values in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=348
func VCVTPS2DQ_FV(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2DQ", reg, rm)
}

// VCVTPS2DQ_RM
// Convert four packed single-precision floating-point values from xmm2/mem to four packed signed doubleword values in xmm1.
// Convert eight packed single-precision floating-point values from ymm2/mem to eight packed signed doubleword values in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=348
func VCVTPS2DQ_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2DQ", reg, rm)
}

// VCVTPS2PD_HV
// Convert two packed single-precision floating-point values in xmm2/m64/m32bcst to packed double-precision floating- point values in xmm1 with writemask k1.
// Convert four packed single-precision floating-point values in xmm2/m128/m32bcst to packed double-precision floating-point values in ymm1 with writemask k1.
// Convert eight packed single-precision floating-point values in ymm2/m256/b32bcst to eight packed double-precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=351
func VCVTPS2PD_HV(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2PD", reg, rm)
}

// VCVTPS2PD_RM
// Convert two packed single-precision floating-point values in xmm2/m64 to two packed double-precision floating-point values in xmm1.
// Convert four packed single-precision floating-point values in xmm2/m128 to four packed double-precision floating- point values in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=351
func VCVTPS2PD_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2PD", reg, rm)
}

// VCVTPS2QQ
// Convert two packed single precision floating-point values from xmm2/m64/m32bcst to two packed signed quadword values in xmm1 subject to writemask k1.
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed signed quadword values in ymm1 subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed signed quadword values in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1400
func VCVTPS2QQ(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2QQ", reg, rm)
}

// VCVTPS2UDQ
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed unsigned doubleword values in xmm1 subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed unsigned doubleword values in ymm1 subject to writemask k1.
// Convert sixteen packed single-precision floating-point values from zmm2/m512/m32bcst to sixteen packed unsigned doubleword values in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1397
func VCVTPS2UDQ(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2UDQ", reg, rm)
}

// VCVTPS2UQQ
// Convert two packed single precision floating-point values from zmm2/m64/m32bcst to two packed unsigned quadword values in zmm1 subject to writemask k1.
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed unsigned quadword values in ymm1 subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed unsigned quadword values in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1403
func VCVTPS2UQQ(reg, rm interface{}) {
	unsafe.Asm("VCVTPS2UQQ", reg, rm)
}

// VCVTQQ2PD
// Convert two packed quadword integers from xmm2/m128/m64bcst to packed double-precision floating- point values in xmm1 with writemask k1.
// Convert four packed quadword integers from ymm2/m256/m64bcst to packed double-precision floating- point values in ymm1 with writemask k1.
// Convert eight packed quadword integers from zmm2/m512/m64bcst to eight packed double-precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1406
func VCVTQQ2PD(reg, rm interface{}) {
	unsafe.Asm("VCVTQQ2PD", reg, rm)
}

// VCVTQQ2PS
// Convert two packed quadword integers from xmm2/mem to packed single-precision floating-point values in xmm1 with writemask k1.
// Convert four packed quadword integers from ymm2/mem to packed single-precision floating-point values in xmm1 with writemask k1.
// Convert eight packed quadword integers from zmm2/mem to eight packed single-precision floating-point values in ymm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1408
func VCVTQQ2PS(reg, rm interface{}) {
	unsafe.Asm("VCVTQQ2PS", reg, rm)
}

// VCVTSD2SI_RM
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer r32.
// Convert one double-precision floating-point value from xmm1/m64 to one signed quadword integer sign- extended into r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=355
func VCVTSD2SI_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTSD2SI", reg, rm)
}

// VCVTSD2SI_T1F
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer r32.
// Convert one double-precision floating-point value from xmm1/m64 to one signed quadword integer sign- extended into r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=355
func VCVTSD2SI_T1F(reg, rm interface{}) {
	unsafe.Asm("VCVTSD2SI", reg, rm)
}

// VCVTSD2SS_RVM
// Convert one double-precision floating-point value in xmm3/m64 to one single-precision floating-point value and merge with high bits in xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=357
func VCVTSD2SS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VCVTSD2SS", reg, vex, rm)
}

// VCVTSD2SS_T1S
// Convert one double-precision floating-point value in xmm3/m64 to one single-precision floating-point value and merge with high bits in xmm2 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=357
func VCVTSD2SS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VCVTSD2SS", reg, evex, rm)
}

// VCVTSD2USI
// Convert one double-precision floating-point value from xmm1/m64 to one unsigned doubleword integer r32.
// Convert one double-precision floating-point value from xmm1/m64 to one unsigned quadword integer zero- extended into r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1410
func VCVTSD2USI(reg, rm interface{}) {
	unsafe.Asm("VCVTSD2USI", reg, rm)
}

// VCVTSI2SD_RVM
// Convert one signed doubleword integer from r/m32 to one double-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=359
func VCVTSI2SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VCVTSI2SD", reg, vex, rm)
}

// VCVTSI2SD_T1S
// Convert one signed doubleword integer from r/m32 to one double-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=359
func VCVTSI2SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VCVTSI2SD", reg, evex, rm)
}

// VCVTSI2SS_RVM
// Convert one signed doubleword integer from r/m32 to one single-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=361
func VCVTSI2SS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VCVTSI2SS", reg, vex, rm)
}

// VCVTSI2SS_T1S
// Convert one signed doubleword integer from r/m32 to one single-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=361
func VCVTSI2SS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VCVTSI2SS", reg, evex, rm)
}

// VCVTSS2SD_RVM
// Convert one single-precision floating-point value in xmm3/m32 to one double-precision floating-point value and merge with high bits of xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=363
func VCVTSS2SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VCVTSS2SD", reg, vex, rm)
}

// VCVTSS2SD_T1S
// Convert one single-precision floating-point value in xmm3/m32 to one double-precision floating-point value and merge with high bits of xmm2 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=363
func VCVTSS2SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VCVTSS2SD", reg, evex, rm)
}

// VCVTSS2SI_RM
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=365
func VCVTSS2SI_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTSS2SI", reg, rm)
}

// VCVTSS2SI_T1F
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=365
func VCVTSS2SI_T1F(reg, rm interface{}) {
	unsafe.Asm("VCVTSS2SI", reg, rm)
}

// VCVTSS2USI
// Convert one single-precision floating-point value from xmm1/m32 to one unsigned doubleword integer in r32.
// Convert one single-precision floating-point value from xmm1/m32 to one unsigned quadword integer in r64.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1411
func VCVTSS2USI(reg, rm interface{}) {
	unsafe.Asm("VCVTSS2USI", reg, rm)
}

// VCVTTPD2DQ_FV
// Convert two packed double-precision floating-point values in xmm2/m128/m64bcst to two signed doubleword integers in xmm1 using truncation subject to writemask k1.
// Convert four packed double-precision floating-point values in ymm2/m256/m64bcst to four signed doubleword integers in xmm1 using truncation subject to writemask k1.
// Convert eight packed double-precision floating-point values in zmm2/m512/m64bcst to eight signed doubleword integers in ymm1 using truncation subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=367
func VCVTTPD2DQ_FV(reg, rm interface{}) {
	unsafe.Asm("VCVTTPD2DQ", reg, rm)
}

// VCVTTPD2DQ_RM
// Convert two packed double-precision floating-point values in xmm2/mem to two signed doubleword integers in xmm1 using truncation.
// Convert four packed double-precision floating-point values in ymm2/mem to four signed doubleword integers in xmm1 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=367
func VCVTTPD2DQ_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTTPD2DQ", reg, rm)
}

// VCVTTPD2QQ
// Convert two packed double-precision floating-point values from zmm2/m128/m64bcst to two packed quadword integers in zmm1 using truncation with writemask k1.
// Convert four packed double-precision floating-point values from ymm2/m256/m64bcst to four packed quadword integers in ymm1 using truncation with writemask k1.
// Convert eight packed double-precision floating-point values from zmm2/m512 to eight packed quadword integers in zmm1 using truncation with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1413
func VCVTTPD2QQ(reg, rm interface{}) {
	unsafe.Asm("VCVTTPD2QQ", reg, rm)
}

// VCVTTPD2UQQ
// Convert two packed double-precision floating-point values from xmm2/m128/m64bcst to two packed unsigned quadword integers in xmm1 using truncation with writemask k1.
// Convert four packed double-precision floating-point values from ymm2/m256/m64bcst to four packed unsigned quadword integers in ymm1 using truncation with writemask k1.
// Convert eight packed double-precision floating-point values from zmm2/mem to eight packed unsigned quadword integers in zmm1 using truncation with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1418
func VCVTTPD2UQQ(reg, rm interface{}) {
	unsafe.Asm("VCVTTPD2UQQ", reg, rm)
}

// VCVTTPS2DQ_FV
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed signed doubleword values in xmm1 using truncation subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed signed doubleword values in ymm1 using truncation subject to writemask k1.
// Convert sixteen packed single-precision floating-point values from zmm2/m512/m32bcst to sixteen packed signed doubleword values in zmm1 using truncation subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=372
func VCVTTPS2DQ_FV(reg, rm interface{}) {
	unsafe.Asm("VCVTTPS2DQ", reg, rm)
}

// VCVTTPS2DQ_RM
// Convert four packed single-precision floating-point values from xmm2/mem to four packed signed doubleword values in xmm1 using truncation.
// Convert eight packed single-precision floating-point values from ymm2/mem to eight packed signed doubleword values in ymm1 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=372
func VCVTTPS2DQ_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTTPS2DQ", reg, rm)
}

// VCVTTPS2QQ
// Convert two packed single precision floating-point values from xmm2/m64/m32bcst to two packed signed quadword values in xmm1 using truncation subject to writemask k1.
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed signed quadword values in ymm1 using truncation subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed signed quadword values in zmm1 using truncation subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1422
func VCVTTPS2QQ(reg, rm interface{}) {
	unsafe.Asm("VCVTTPS2QQ", reg, rm)
}

// VCVTTPS2UDQ
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed unsigned doubleword values in xmm1 using truncation subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed unsigned doubleword values in ymm1 using truncation subject to writemask k1.
// Convert sixteen packed single-precision floating- point values from zmm2/m512/m32bcst to sixteen packed unsigned doubleword values in zmm1 using truncation subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1420
func VCVTTPS2UDQ(reg, rm interface{}) {
	unsafe.Asm("VCVTTPS2UDQ", reg, rm)
}

// VCVTTPS2UQQ
// Convert two packed single precision floating-point values from xmm2/m64/m32bcst to two packed unsigned quadword values in xmm1 using truncation subject to writemask k1.
// Convert four packed single precision floating-point values from xmm2/m128/m32bcst to four packed unsigned quadword values in ymm1 using truncation subject to writemask k1.
// Convert eight packed single precision floating-point values from ymm2/m256/m32bcst to eight packed unsigned quadword values in zmm1 using truncation subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1424
func VCVTTPS2UQQ(reg, rm interface{}) {
	unsafe.Asm("VCVTTPS2UQQ", reg, rm)
}

// VCVTTSD2SI_RM
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer in r32 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=376
func VCVTTSD2SI_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTTSD2SI", reg, rm)
}

// VCVTTSD2SI_T1F
// Convert one double-precision floating-point value from xmm1/m64 to one signed doubleword integer in r32 using truncation.
// Convert one double-precision floating-point value from xmm1/m64 to one signed quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=376
func VCVTTSD2SI_T1F(reg, rm interface{}) {
	unsafe.Asm("VCVTTSD2SI", reg, rm)
}

// VCVTTSD2USI
// Convert one double-precision floating-point value from xmm1/m64 to one unsigned doubleword integer r32 using truncation.
// Convert one double-precision floating-point value from xmm1/m64 to one unsigned quadword integer zero- extended into r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1426
func VCVTTSD2USI(reg, rm interface{}) {
	unsafe.Asm("VCVTTSD2USI", reg, rm)
}

// VCVTTSS2SI_RM
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32 using truncation.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=378
func VCVTTSS2SI_RM(reg, rm interface{}) {
	unsafe.Asm("VCVTTSS2SI", reg, rm)
}

// VCVTTSS2SI_T1F
// Convert one single-precision floating-point value from xmm1/m32 to one signed doubleword integer in r32 using truncation.
// Convert one single-precision floating-point value from xmm1/m32 to one signed quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=378
func VCVTTSS2SI_T1F(reg, rm interface{}) {
	unsafe.Asm("VCVTTSS2SI", reg, rm)
}

// VCVTTSS2USI
// Convert one single-precision floating-point value from xmm1/m32 to one unsigned doubleword integer in r32 using truncation.
// Convert one single-precision floating-point value from xmm1/m32 to one unsigned quadword integer in r64 using truncation.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1427
func VCVTTSS2USI(reg, rm interface{}) {
	unsafe.Asm("VCVTTSS2USI", reg, rm)
}

// VCVTUDQ2PD
// Convert two packed unsigned doubleword integers from ymm2/m64/m32bcst to packed double-precision floating-point values in zmm1 with writemask k1.
// Convert four packed unsigned doubleword integers from xmm2/m128/m32bcst to packed double- precision floating-point values in zmm1 with writemask k1.
// Convert eight packed unsigned doubleword integers from ymm2/m256/m32bcst to eight packed double- precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1429
func VCVTUDQ2PD(reg, rm interface{}) {
	unsafe.Asm("VCVTUDQ2PD", reg, rm)
}

// VCVTUDQ2PS
// Convert four packed unsigned doubleword integers from xmm2/m128/m32bcst to packed single-precision floating-point values in xmm1 with writemask k1.
// Convert eight packed unsigned doubleword integers from ymm2/m256/m32bcst to packed single-precision floating-point values in zmm1 with writemask k1.
// Convert sixteen packed unsigned doubleword integers from zmm2/m512/m32bcst to sixteen packed single- precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1431
func VCVTUDQ2PS(reg, rm interface{}) {
	unsafe.Asm("VCVTUDQ2PS", reg, rm)
}

// VCVTUQQ2PD
// Convert two packed unsigned quadword integers from xmm2/m128/m64bcst to two packed double-precision floating-point values in xmm1 with writemask k1.
// Convert four packed unsigned quadword integers from ymm2/m256/m64bcst to packed double-precision floating- point values in ymm1 with writemask k1.
// Convert eight packed unsigned quadword integers from zmm2/m512/m64bcst to eight packed double-precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1433
func VCVTUQQ2PD(reg, rm interface{}) {
	unsafe.Asm("VCVTUQQ2PD", reg, rm)
}

// VCVTUQQ2PS
// Convert two packed unsigned quadword integers from xmm2/m128/m64bcst to packed single-precision floating- point values in zmm1 with writemask k1.
// Convert four packed unsigned quadword integers from ymm2/m256/m64bcst to packed single-precision floating- point values in xmm1 with writemask k1.
// Convert eight packed unsigned quadword integers from zmm2/m512/m64bcst to eight packed single-precision floating-point values in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1435
func VCVTUQQ2PS(reg, rm interface{}) {
	unsafe.Asm("VCVTUQQ2PS", reg, rm)
}

// VCVTUSI2SD
// Convert one unsigned doubleword integer from r/m32 to one double-precision floating-point value in xmm1.
// Convert one unsigned quadword integer from r/m64 to one double-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1437
func VCVTUSI2SD(reg, evex, rm interface{}) {
	unsafe.Asm("VCVTUSI2SD", reg, evex, rm)
}

// VCVTUSI2SS
// Convert one signed doubleword integer from r/m32 to one single-precision floating-point value in xmm1.
// Convert one signed quadword integer from r/m64 to one single-precision floating-point value in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1439
func VCVTUSI2SS(reg, vex, rm interface{}) {
	unsafe.Asm("VCVTUSI2SS", reg, vex, rm)
}

// VDIVPD_FV
// Divide packed double-precision floating-point values in xmm2 by packed double-precision floating-point values in xmm3/m128/m64bcst and write results to xmm1 subject to writemask k1.
// Divide packed double-precision floating-point values in ymm2 by packed double-precision floating-point values in ymm3/m256/m64bcst and write results to ymm1 subject to writemask k1.
// Divide packed double-precision floating-point values in zmm2 by packed double-precision FP values in zmm3/m512/m64bcst and write results to zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=390
func VDIVPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VDIVPD", reg, evex, rm)
}

// VDIVPD_RVM
// Divide packed double-precision floating-point values in xmm2 by packed double-precision floating-point values in xmm3/mem.
// Divide packed double-precision floating-point values in ymm2 by packed double-precision floating-point values in ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=390
func VDIVPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VDIVPD", reg, vex, rm)
}

// VDIVPS_FV
// Divide packed single-precision floating-point values in xmm2 by packed single-precision floating-point values in xmm3/m128/m32bcst and write results to xmm1 subject to writemask k1.
// Divide packed single-precision floating-point values in ymm2 by packed single-precision floating-point values in ymm3/m256/m32bcst and write results to ymm1 subject to writemask k1.
// Divide packed single-precision floating-point values in zmm2 by packed single-precision floating-point values in zmm3/m512/m32bcst and write results to zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=393
func VDIVPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VDIVPS", reg, evex, rm)
}

// VDIVPS_RVM
// Divide packed single-precision floating-point values in xmm2 by packed single-precision floating-point values in xmm3/mem.
// Divide packed single-precision floating-point values in ymm2 by packed single-precision floating-point values in ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=393
func VDIVPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VDIVPS", reg, vex, rm)
}

// VDIVSD_RVM
// Divide low double-precision floating-point value in xmm2 by low double-precision floating-point value in xmm3/m64.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=396
func VDIVSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VDIVSD", reg, vex, rm)
}

// VDIVSD_T1S
// Divide low double-precision floating-point value in xmm2 by low double-precision floating-point value in xmm3/m64.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=396
func VDIVSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VDIVSD", reg, evex, rm)
}

// VDIVSS_RVM
// Divide low single-precision floating-point value in xmm2 by low single-precision floating-point value in xmm3/m32.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=398
func VDIVSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VDIVSS", reg, vex, rm)
}

// VDIVSS_T1S
// Divide low single-precision floating-point value in xmm2 by low single-precision floating-point value in xmm3/m32.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=398
func VDIVSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VDIVSS", reg, evex, rm)
}

// VDPPD
// Selectively multiply packed DP floating-point values from xmm2 with packed DP floating- point values from xmm3, add and selectively store the packed DP floating-point values to xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=400
func VDPPD(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VDPPD", reg, vex, rm, imm)
}

// VDPPS
// Multiply packed SP floating point values from xmm1 with packed SP floating point values from xmm2/mem selectively add and store to xmm1.
// Multiply packed single-precision floating-point values from ymm2 with packed SP floating point values from ymm3/mem, selectively add pairs of elements and store to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=402
func VDPPS(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VDPPS", reg, vex, rm, imm)
}

// VERR
// Set ZF=1 if segment specified with r/m16 can be read.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1449
func VERR(rm interface{}) {
	unsafe.Asm("VERR", rm)
}

// VERW
// Set ZF=1 if segment specified with r/m16 can be written.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1449
func VERW(rm interface{}) {
	unsafe.Asm("VERW", rm)
}

// VEXP2PD
// Computes approximations to the exponential 2^x (with less than 2^-23 of maximum relative error) of the packed double- precision floating-point values from zmm2/m512/m64bcst and stores the floating-point result in zmm1with writemask k1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1451
func VEXP2PD(reg, rm interface{}) {
	unsafe.Asm("VEXP2PD", reg, rm)
}

// VEXP2PS
// Computes approximations to the exponential 2^x (with less than 2^-23 of maximum relative error) of the packed single- precision floating-point values from zmm2/m512/m32bcst and stores the floating-point result in zmm1with writemask k1.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1453
func VEXP2PS(reg, rm interface{}) {
	unsafe.Asm("VEXP2PS", reg, rm)
}

// VEXPANDPD
// Expand packed double-precision floating-point values from xmm2/m128 to xmm1 using writemask k1.
// Expand packed double-precision floating-point values from ymm2/m256 to ymm1 using writemask k1.
// Expand packed double-precision floating-point values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1445
func VEXPANDPD(reg, rm interface{}) {
	unsafe.Asm("VEXPANDPD", reg, rm)
}

// VEXTRACTF128
// Extract 128 bits of packed floating-point values from ymm2 and store results in xmm1/m128.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF128(rm, reg, imm interface{}) {
	unsafe.Asm("VEXTRACTF128", rm, reg, imm)
}

// VEXTRACTF32X4
// Extract 128 bits of packed single-precision floating- point values from ymm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF32X4() {
	unsafe.Asm("VEXTRACTF32X4", nil)
}

// VEXTRACTF32X8
// Extract 256 bits of packed single-precision floating- point values from zmm2 and store results in ymm1/m256 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF32X8() {
	unsafe.Asm("VEXTRACTF32X8", nil)
}

// VEXTRACTF32x4
// Extract 128 bits of packed single-precision floating- point values from zmm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF32x4() {
	unsafe.Asm("VEXTRACTF32x4", nil)
}

// VEXTRACTF64X2
// Extract 128 bits of packed double-precision floating-point values from ymm2 and store results in xmm1/m128 subject to writemask k1.
// Extract 128 bits of packed double-precision floating-point values from zmm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF64X2() {
	unsafe.Asm("VEXTRACTF64X2", nil)
}

// VEXTRACTF64x4
// Extract 256 bits of packed double-precision floating-point values from zmm2 and store results in ymm1/m256 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1455
func VEXTRACTF64x4() {
	unsafe.Asm("VEXTRACTF64x4", nil)
}

// VEXTRACTI128
// Extract 128 bits of integer data from ymm2 and store results in xmm1/m128.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI128(rm, reg, imm interface{}) {
	unsafe.Asm("VEXTRACTI128", rm, reg, imm)
}

// VEXTRACTI32X4
// Extract 128 bits of double-word integer values from ymm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI32X4() {
	unsafe.Asm("VEXTRACTI32X4", nil)
}

// VEXTRACTI32X8
// Extract 256 bits of double-word integer values from zmm2 and store results in ymm1/m256 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI32X8() {
	unsafe.Asm("VEXTRACTI32X8", nil)
}

// VEXTRACTI32x4
// Extract 128 bits of double-word integer values from zmm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI32x4() {
	unsafe.Asm("VEXTRACTI32x4", nil)
}

// VEXTRACTI64X2
// Extract 128 bits of quad-word integer values from ymm2 and store results in xmm1/m128 subject to writemask k1.
// Extract 128 bits of quad-word integer values from zmm2 and store results in xmm1/m128 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI64X2() {
	unsafe.Asm("VEXTRACTI64X2", nil)
}

// VEXTRACTI64x4
// Extract 256 bits of quad-word integer values from zmm2 and store results in ymm1/m256 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1462
func VEXTRACTI64x4() {
	unsafe.Asm("VEXTRACTI64x4", nil)
}

// VFIXUPIMMPD
// Fix up special numbers in float64 vector xmm1, float64 vector xmm2 and int64 vector xmm3/m128/m64bcst and store the result in xmm1, under writemask.
// Fix up special numbers in float64 vector ymm1, float64 vector ymm2 and int64 vector ymm3/m256/m64bcst and store the result in ymm1, under writemask.
// Fix up elements of float64 vector in zmm2 using int64 vector table in zmm3/m512/m64bcst, combine with preserved elements from zmm1, and store the result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1468
func VFIXUPIMMPD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VFIXUPIMMPD", reg, evex, rm, imm)
}

// VFIXUPIMMPS
// Fix up special numbers in float32 vector xmm1, float32 vector xmm2 and int32 vector xmm3/m128/m32bcst and store the result in xmm1, under writemask.
// Fix up special numbers in float32 vector ymm1, float32 vector ymm2 and int32 vector ymm3/m256/m32bcst and store the result in ymm1, under writemask.
// Fix up elements of float32 vector in zmm2 using int32 vector table in zmm3/m512/m32bcst, combine with preserved elements from zmm1, and store the result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1472
func VFIXUPIMMPS(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VFIXUPIMMPS", reg, evex, rm, imm)
}

// VFIXUPIMMSD
// Fix up a float64 number in the low quadword element of xmm2 using scalar int32 table in xmm3/m64 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1476
func VFIXUPIMMSD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VFIXUPIMMSD", reg, evex, rm, imm)
}

// VFIXUPIMMSS
// Fix up a float32 number in the low doubleword element in xmm2 using scalar int32 table in xmm3/m32 and store the result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1479
func VFIXUPIMMSS(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VFIXUPIMMSS", reg, evex, rm, imm)
}

// VFMADD132PD_FV
// Multiply packed double-precision floating-point values from ymm1 and ymm3/m256/m64bcst, add to ymm2 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm3/m512/m64bcst, add to zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD132PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD132PD", reg, evex, rm)
}

// VFMADD132PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm3/m128/m64bcst, add to xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from xmm1 and xmm3/mem, add to xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/mem, add to ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD132PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD132PD", reg, vex, rm)
}

// VFMADD132PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm3/m128/m32bcst, add to xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/m256/m32bcst, add to ymm2 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm3/m512/m32bcst, add to zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD132PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD132PS", reg, evex, rm)
}

// VFMADD132PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm3/mem, add to xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/mem, add to ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD132PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD132PS", reg, vex, rm)
}

// VFMADD132SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/m64, add to xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD132SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD132SD", reg, vex, rm)
}

// VFMADD132SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/m64, add to xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD132SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD132SD", reg, evex, rm)
}

// VFMADD213PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm2, add to xmm3/m128/m64bcst and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, add to ymm3/m256/m64bcst and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm2, add to zmm3/m512/m64bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD213PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD213PD", reg, evex, rm)
}

// VFMADD213PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm2, add to xmm3/mem and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, add to ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD213PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD213PD", reg, vex, rm)
}

// VFMADD213PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm2, add to xmm3/m128/m32bcst and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, add to ymm3/m256/m32bcst and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm2, add to zmm3/m512/m32bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD213PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD213PS", reg, evex, rm)
}

// VFMADD213PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm2, add to xmm3/mem and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, add to ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD213PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD213PS", reg, vex, rm)
}

// VFMADD213SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, add to xmm3/m64 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD213SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD213SD", reg, vex, rm)
}

// VFMADD213SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, add to xmm3/m64 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD213SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD213SD", reg, evex, rm)
}

// VFMADD231PD_FV
// Multiply packed double-precision floating-point values from xmm2 and xmm3/m128/m64bcst, add to xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/m256/m64bcst, add to ymm1 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm2 and zmm3/m512/m64bcst, add to zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD231PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD231PD", reg, evex, rm)
}

// VFMADD231PD_RVM
// Multiply packed double-precision floating-point values from xmm2 and xmm3/mem, add to xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/mem, add to ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1482
func VFMADD231PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD231PD", reg, vex, rm)
}

// VFMADD231PS_FV
// Multiply packed single-precision floating-point values from xmm2 and xmm3/m128/m32bcst, add to xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/m256/m32bcst, add to ymm1 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm2 and zmm3/m512/m32bcst, add to zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD231PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD231PS", reg, evex, rm)
}

// VFMADD231PS_RVM
// Multiply packed single-precision floating-point values from xmm2 and xmm3/mem, add to xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/mem, add to ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1489
func VFMADD231PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD231PS", reg, vex, rm)
}

// VFMADD231SD_RVM
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/m64, add to xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD231SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMADD231SD", reg, vex, rm)
}

// VFMADD231SD_T1S
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/m64, add to xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1496
func VFMADD231SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMADD231SD", reg, evex, rm)
}

// VFMSUB132PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm3/m128/m64bcst, subtract xmm2 and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/m256/m64bcst, subtract ymm2 and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm1 and zmm3/m512/m64bcst, subtract zmm2 and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB132PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB132PD", reg, evex, rm)
}

// VFMSUB132PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm3/mem, subtract xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/mem, subtract ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB132PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB132PD", reg, vex, rm)
}

// VFMSUB132SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/m64, subtract xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB132SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB132SD", reg, vex, rm)
}

// VFMSUB132SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/m64, subtract xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB132SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB132SD", reg, evex, rm)
}

// VFMSUB213PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm2, subtract xmm3/m128/m64bcst and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, subtract ymm3/m256/m64bcst and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm1 and zmm2, subtract zmm3/m512/m64bcst and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB213PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB213PD", reg, evex, rm)
}

// VFMSUB213PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm2, subtract xmm3/mem and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, subtract ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB213PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB213PD", reg, vex, rm)
}

// VFMSUB213SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, subtract xmm3/m64 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB213SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB213SD", reg, vex, rm)
}

// VFMSUB213SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, subtract xmm3/m64 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB213SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB213SD", reg, evex, rm)
}

// VFMSUB231PD_FV
// Multiply packed double-precision floating-point values from xmm2 and xmm3/m128/m64bcst, subtract xmm1 and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/m256/m64bcst, subtract ymm1 and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm2 and zmm3/m512/m64bcst, subtract zmm1 and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB231PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB231PD", reg, evex, rm)
}

// VFMSUB231PD_RVM
// Multiply packed double-precision floating-point values from xmm2 and xmm3/mem, subtract xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/mem, subtract ymm1 and put result in ymm1.S
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1541
func VFMSUB231PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB231PD", reg, vex, rm)
}

// VFMSUB231SD_RVM
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/m64, subtract xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB231SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUB231SD", reg, vex, rm)
}

// VFMSUB231SD_T1S
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/m64, subtract xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1555
func VFMSUB231SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUB231SD", reg, evex, rm)
}

// VFMSUBADD132PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm3/m128/m64bcst, subtract/add elements in xmm2 and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/m256/m64bcst, subtract/add elements in ymm2 and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm1 and zmm3/m512/m64bcst, subtract/add elements in zmm2 and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD132PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUBADD132PD", reg, evex, rm)
}

// VFMSUBADD132PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm3/mem, subtract/add elements in xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/mem, subtract/add elements in ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD132PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUBADD132PD", reg, vex, rm)
}

// VFMSUBADD213PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm2, subtract/add elements in xmm3/m128/m64bcst and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, subtract/add elements in ymm3/m256/m64bcst and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm1 and zmm2, subtract/add elements in zmm3/m512/m64bcst and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD213PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUBADD213PD", reg, evex, rm)
}

// VFMSUBADD213PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm2, subtract/add elements in xmm3/mem and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, subtract/add elements in ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD213PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUBADD213PD", reg, vex, rm)
}

// VFMSUBADD231PD_FV
// Multiply packed double-precision floating-point values from xmm2 and xmm3/m128/m64bcst, subtract/add elements in xmm1 and put result in xmm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/m256/m64bcst, subtract/add elements in ymm1 and put result in ymm1 subject to writemask k1.
// Multiply packed double-precision floating-point values from zmm2 and zmm3/m512/m64bcst, subtract/add elements in zmm1 and put result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD231PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFMSUBADD231PD", reg, evex, rm)
}

// VFMSUBADD231PD_RVM
// Multiply packed double-precision floating-point values from xmm2 and xmm3/mem, subtract/add elements in xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/mem, subtract/add elements in ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1521
func VFMSUBADD231PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFMSUBADD231PD", reg, vex, rm)
}

// VFNMADD132PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm3/m128/m64bcst, negate the multiplication result and add to xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/m256/m64bcst, negate the multiplication result and add to ymm2 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm3/m512/m64bcst, negate the multiplication result and add to zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD132PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD132PD", reg, evex, rm)
}

// VFNMADD132PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm3/mem, negate the multiplication result and add to xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/mem, negate the multiplication result and add to ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD132PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD132PD", reg, vex, rm)
}

// VFNMADD132PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm3/m128/m32bcst, negate the multiplication result and add to xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/m256/m32bcst, negate the multiplication result and add to ymm2 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm3/m512/m32bcst, negate the multiplication result and add to zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD132PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD132PS", reg, evex, rm)
}

// VFNMADD132PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm3/mem, negate the multiplication result and add to xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/mem, negate the multiplication result and add to ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD132PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD132PS", reg, vex, rm)
}

// VFNMADD132SS_RVM
// Multiply scalar single-precision floating-point value from xmm1 and xmm3/m32, negate the multiplication result and add to xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD132SS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD132SS", reg, vex, rm)
}

// VFNMADD132SS_T1S
// Multiply scalar single-precision floating-point value from xmm1 and xmm3/m32, negate the multiplication result and add to xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD132SS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD132SS", reg, evex, rm)
}

// VFNMADD213PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm2, negate the multiplication result and add to xmm3/m128/m64bcst and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, negate the multiplication result and add to ymm3/m256/m64bcst and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm2, negate the multiplication result and add to zmm3/m512/m64bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD213PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD213PD", reg, evex, rm)
}

// VFNMADD213PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm2, negate the multiplication result and add to xmm3/mem and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, negate the multiplication result and add to ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD213PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD213PD", reg, vex, rm)
}

// VFNMADD213PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm2, negate the multiplication result and add to xmm3/m128/m32bcst and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, negate the multiplication result and add to ymm3/m256/m32bcst and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm2, negate the multiplication result and add to zmm3/m512/m32bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD213PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD213PS", reg, evex, rm)
}

// VFNMADD213PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm2, negate the multiplication result and add to xmm3/mem and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, negate the multiplication result and add to ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD213PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD213PS", reg, vex, rm)
}

// VFNMADD213SS_RVM
// Multiply scalar single-precision floating-point value from xmm1 and xmm2, negate the multiplication result and add to xmm3/m32 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD213SS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD213SS", reg, vex, rm)
}

// VFNMADD213SS_T1S
// Multiply scalar single-precision floating-point value from xmm1 and xmm2, negate the multiplication result and add to xmm3/m32 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD213SS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD213SS", reg, evex, rm)
}

// VFNMADD231PD_FV
// Multiply packed double-precision floating-point values from xmm2 and xmm3/m128/m64bcst, negate the multiplication result and add to xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/m256/m64bcst, negate the multiplication result and add to ymm1 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm2 and zmm3/m512/m64bcst, negate the multiplication result and add to zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD231PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD231PD", reg, evex, rm)
}

// VFNMADD231PD_RVM
// Multiply packed double-precision floating-point values from xmm2 and xmm3/mem, negate the multiplication result and add to xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/mem, negate the multiplication result and add to ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1561
func VFNMADD231PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD231PD", reg, vex, rm)
}

// VFNMADD231PS_FV
// Multiply packed single-precision floating-point values from xmm2 and xmm3/m128/m32bcst, negate the multiplication result and add to xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/m256/m32bcst, negate the multiplication result and add to ymm1 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm2 and zmm3/m512/m32bcst, negate the multiplication result and add to zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD231PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD231PS", reg, evex, rm)
}

// VFNMADD231PS_RVM
// Multiply packed single-precision floating-point values from xmm2 and xmm3/mem, negate the multiplication result and add to xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/mem, negate the multiplication result and add to ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1568
func VFNMADD231PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD231PS", reg, vex, rm)
}

// VFNMADD231SS_RVM
// Multiply scalar single-precision floating-point value from xmm2 and xmm3/m32, negate the multiplication result and add to xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD231SS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMADD231SS", reg, vex, rm)
}

// VFNMADD231SS_T1S
// Multiply scalar single-precision floating-point value from xmm2 and xmm3/m32, negate the multiplication result and add to xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1577
func VFNMADD231SS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMADD231SS", reg, evex, rm)
}

// VFNMSUB132PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm3/m128/m64bcst, negate the multiplication result and subtract xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/m256/m64bcst, negate the multiplication result and subtract ymm2 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm3/m512/m64bcst, negate the multiplication result and subtract zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB132PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB132PD", reg, evex, rm)
}

// VFNMSUB132PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm3/mem, negate the multiplication result and subtract xmm2 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm3/mem, negate the multiplication result and subtract ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB132PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB132PD", reg, vex, rm)
}

// VFNMSUB132PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm3/m128/m32bcst, negate the multiplication result and subtract xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/m256/m32bcst, negate the multiplication result and subtract ymm2 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm3/m512/m32bcst, negate the multiplication result and subtract zmm2 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB132PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB132PS", reg, evex, rm)
}

// VFNMSUB132PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm3/mem, negate the multiplication result and subtract xmm2 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm3/mem, negate the multiplication result and subtract ymm2 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB132PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB132PS", reg, vex, rm)
}

// VFNMSUB132SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/mem, negate the multiplication result and subtract xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB132SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB132SD", reg, vex, rm)
}

// VFNMSUB132SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm3/m64, negate the multiplication result and subtract xmm2 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB132SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB132SD", reg, evex, rm)
}

// VFNMSUB213PD_FV
// Multiply packed double-precision floating-point values from xmm1 and xmm2, negate the multiplication result and subtract xmm3/m128/m64bcst and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, negate the multiplication result and subtract ymm3/m256/m64bcst and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm1 and zmm2, negate the multiplication result and subtract zmm3/m512/m64bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB213PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB213PD", reg, evex, rm)
}

// VFNMSUB213PD_RVM
// Multiply packed double-precision floating-point values from xmm1 and xmm2, negate the multiplication result and subtract xmm3/mem and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm1 and ymm2, negate the multiplication result and subtract ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB213PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB213PD", reg, vex, rm)
}

// VFNMSUB213PS_FV
// Multiply packed single-precision floating-point values from xmm1 and xmm2, negate the multiplication result and subtract xmm3/m128/m32bcst and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, negate the multiplication result and subtract ymm3/m256/m32bcst and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm1 and zmm2, negate the multiplication result and subtract zmm3/m512/m32bcst and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB213PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB213PS", reg, evex, rm)
}

// VFNMSUB213PS_RVM
// Multiply packed single-precision floating-point values from xmm1 and xmm2, negate the multiplication result and subtract xmm3/mem and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm1 and ymm2, negate the multiplication result and subtract ymm3/mem and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB213PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB213PS", reg, vex, rm)
}

// VFNMSUB213SD_RVM
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, negate the multiplication result and subtract xmm3/mem and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB213SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB213SD", reg, vex, rm)
}

// VFNMSUB213SD_T1S
// Multiply scalar double-precision floating-point value from xmm1 and xmm2, negate the multiplication result and subtract xmm3/m64 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB213SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB213SD", reg, evex, rm)
}

// VFNMSUB231PD_FV
// Multiply packed double-precision floating-point values from xmm2 and xmm3/m128/m64bcst, negate the multiplication result and subtract xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/m256/m64bcst, negate the multiplication result and subtract ymm1 and put result in ymm1.
// Multiply packed double-precision floating-point values from zmm2 and zmm3/m512/m64bcst, negate the multiplication result and subtract zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB231PD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB231PD", reg, evex, rm)
}

// VFNMSUB231PD_RVM
// Multiply packed double-precision floating-point values from xmm2 and xmm3/mem, negate the multiplication result and subtract xmm1 and put result in xmm1.
// Multiply packed double-precision floating-point values from ymm2 and ymm3/mem, negate the multiplication result and subtract ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1580
func VFNMSUB231PD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB231PD", reg, vex, rm)
}

// VFNMSUB231PS_FV
// Multiply packed single-precision floating-point values from xmm2 and xmm3/m128/m32bcst, negate the multiplication result subtract add to xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/m256/m32bcst, negate the multiplication result subtract add to ymm1 and put result in ymm1.
// Multiply packed single-precision floating-point values from zmm2 and zmm3/m512/m32bcst, negate the multiplication result subtract add to zmm1 and put result in zmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB231PS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB231PS", reg, evex, rm)
}

// VFNMSUB231PS_RVM
// Multiply packed single-precision floating-point values from xmm2 and xmm3/mem, negate the multiplication result and subtract xmm1 and put result in xmm1.
// Multiply packed single-precision floating-point values from ymm2 and ymm3/mem, negate the multiplication result and subtract ymm1 and put result in ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1586
func VFNMSUB231PS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB231PS", reg, vex, rm)
}

// VFNMSUB231SD_RVM
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/mem, negate the multiplication result and subtract xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB231SD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VFNMSUB231SD", reg, vex, rm)
}

// VFNMSUB231SD_T1S
// Multiply scalar double-precision floating-point value from xmm2 and xmm3/m64, negate the multiplication result and subtract xmm1 and put result in xmm1.
//
// reg: ModRM:reg (r, w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1592
func VFNMSUB231SD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VFNMSUB231SD", reg, evex, rm)
}

// VFPCLASSPS
// Tests the input for the following categories: NaN, +0, -0, +Infinity, -Infinity, denormal, finite negative. The immediate field provides a mask bit for each of these category tests. The masked test results are OR-ed together to form a mask result.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1601
func VFPCLASSPS(reg, rm, imm interface{}) {
	unsafe.Asm("VFPCLASSPS", reg, rm, imm)
}

// VFPCLASSSD
// Tests the input for the following categories: NaN, +0, -0, +Infinity, -Infinity, denormal, finite negative. The immediate field provides a mask bit for each of these category tests. The masked test results are OR-ed together to form a mask result.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1603
func VFPCLASSSD(reg, rm, imm interface{}) {
	unsafe.Asm("VFPCLASSSD", reg, rm, imm)
}

// VFPCLASSSS
// Tests the input for the following categories: NaN, +0, -0, +Infinity, -Infinity, denormal, finite negative. The immediate field provides a mask bit for each of these category tests. The masked test results are OR-ed together to form a mask result.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1605
func VFPCLASSSS(reg, rm, imm interface{}) {
	unsafe.Asm("VFPCLASSSS", reg, rm, imm)
}

// VGATHERDPD_RMV
// Using dword indices specified in vm32x, gather double-pre- cision FP values from memory conditioned on mask speci- fied by xmm2. Conditionally gathered elements are merged into xmm1.
// Using dword indices specified in vm32x, gather double-pre- cision FP values from memory conditioned on mask speci- fied by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1607
func VGATHERDPD_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VGATHERDPD", reg, vsib, vex)
}

// VGATHERDPD_T1S
// Using signed dword indices, gather float64 vector into float64 vector xmm1 using k1 as completion mask.
// Using signed dword indices, gather float64 vector into float64 vector ymm1 using k1 as completion mask.
// Using signed dword indices, gather float64 vector into float64 vector zmm1 using k1 as completion mask.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1617
func VGATHERDPD_T1S(reg, v interface{}) {
	unsafe.Asm("VGATHERDPD", reg, v)
}

// VGATHERDPS_RMV
// Using dword indices specified in vm32x, gather single-preci- sion FP values from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using dword indices specified in vm32y, gather single-preci- sion FP values from memory conditioned on mask specified by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1612
func VGATHERDPS_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VGATHERDPS", reg, vsib, vex)
}

// VGATHERDPS_T1S
// Using signed dword indices, gather single-precision floating- point values from memory using k1 as completion mask.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1617
func VGATHERDPS_T1S(reg, v interface{}) {
	unsafe.Asm("VGATHERDPS", reg, v)
}

// VGATHERPF0DPD
// Using signed dword indices, prefetch sparse byte memory locations containing double-precision data using opmask k1 and T0 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1620
func VGATHERPF0DPD(vsib interface{}) {
	unsafe.Asm("VGATHERPF0DPD", vsib)
}

// VGATHERPF0DPS
// Using signed dword indices, prefetch sparse byte memory locations containing single-precision data using opmask k1 and T0 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1620
func VGATHERPF0DPS(vsib interface{}) {
	unsafe.Asm("VGATHERPF0DPS", vsib)
}

// VGATHERPF0QPD
// Using signed qword indices, prefetch sparse byte memory locations containing double-precision data using opmask k1 and T0 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1620
func VGATHERPF0QPD(vsib interface{}) {
	unsafe.Asm("VGATHERPF0QPD", vsib)
}

// VGATHERPF0QPS
// Using signed qword indices, prefetch sparse byte memory locations containing single-precision data using opmask k1 and T0 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1620
func VGATHERPF0QPS(vsib interface{}) {
	unsafe.Asm("VGATHERPF0QPS", vsib)
}

// VGATHERPF1DPD
// Using signed dword indices, prefetch sparse byte memory locations containing double-precision data using opmask k1 and T1 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1623
func VGATHERPF1DPD(vsib interface{}) {
	unsafe.Asm("VGATHERPF1DPD", vsib)
}

// VGATHERPF1DPS
// Using signed dword indices, prefetch sparse byte memory locations containing single-precision data using opmask k1 and T1 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1623
func VGATHERPF1DPS(vsib interface{}) {
	unsafe.Asm("VGATHERPF1DPS", vsib)
}

// VGATHERPF1QPD
// Using signed qword indices, prefetch sparse byte memory locations containing double-precision data using opmask k1 and T1 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1623
func VGATHERPF1QPD(vsib interface{}) {
	unsafe.Asm("VGATHERPF1QPD", vsib)
}

// VGATHERPF1QPS
// Using signed qword indices, prefetch sparse byte memory locations containing single-precision data using opmask k1 and T1 hint.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1623
func VGATHERPF1QPS(vsib interface{}) {
	unsafe.Asm("VGATHERPF1QPS", vsib)
}

// VGATHERQPD_RMV
// Using qword indices specified in vm64x, gather double-pre- cision FP values from memory conditioned on mask speci- fied by xmm2. Conditionally gathered elements are merged into xmm1.
// Using qword indices specified in vm64y, gather double-pre- cision FP values from memory conditioned on mask speci- fied by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1607
func VGATHERQPD_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VGATHERQPD", reg, vsib, vex)
}

// VGATHERQPD_T1S
// Using signed qword indices, gather float64 vector into float64 vector xmm1 using k1 as completion mask.
// Using signed qword indices, gather float64 vector into float64 vector ymm1 using k1 as completion mask.
// Using signed qword indices, gather float64 vector into float64 vector zmm1 using k1 as completion mask.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1626
func VGATHERQPD_T1S(reg, v interface{}) {
	unsafe.Asm("VGATHERQPD", reg, v)
}

// VGATHERQPS_RMV
// Using qword indices specified in vm64x, gather single-preci- sion FP values from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using qword indices specified in vm64y, gather single-preci- sion FP values from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1612
func VGATHERQPS_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VGATHERQPS", reg, vsib, vex)
}

// VGATHERQPS_T1S
// Using signed qword indices, gather single-precision floating-point values from memory using k1 as completion mask.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1626
func VGATHERQPS_T1S(reg, v interface{}) {
	unsafe.Asm("VGATHERQPS", reg, v)
}

// VGETEXPPS
// Convert the exponent of packed single-precision floating-point values in the source operand to SP FP results representing unbiased integer exponents and stores the results in the destination register.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1647
func VGETEXPPS(reg, rm interface{}) {
	unsafe.Asm("VGETEXPPS", reg, rm)
}

// VGETEXPSS
// Convert the biased exponent (bits 30:23) of the low single- precision floating-point value in xmm3/m32 to a SP FP value representing unbiased integer exponent. Stores the result to xmm1 under the writemask k1 and merge with the other elements of xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1653
func VGETEXPSS(reg, evex, rm interface{}) {
	unsafe.Asm("VGETEXPSS", reg, evex, rm)
}

// VGETMANTPD
// Get Normalized Mantissa from float64 vector xmm2/m128/m64bcst and store the result in xmm1, using imm8 for sign control and mantissa interval normalization, under writemask.
// Get Normalized Mantissa from float64 vector ymm2/m256/m64bcst and store the result in ymm1, using imm8 for sign control and mantissa interval normalization, under writemask.
// Get Normalized Mantissa from float64 vector zmm2/m512/m64bcst and store the result in zmm1, using imm8 for sign control and mantissa interval normalization, under writemask.
//
// Documentation: https://golang.org/s/x86manual#page=1655
func VGETMANTPD() {
	unsafe.Asm("VGETMANTPD", nil)
}

// VGETMANTPS
// Get normalized mantissa from float32 vector xmm2/m128/m32bcst and store the result in xmm1, using imm8 for sign control and mantissa interval normalization, under writemask.
// Get normalized mantissa from float32 vector ymm2/m256/m32bcst and store the result in ymm1, using imm8 for sign control and mantissa interval normalization, under writemask.
// Get normalized mantissa from float32 vector zmm2/m512/m32bcst and store the result in zmm1, using imm8 for sign control and mantissa interval normalization, under writemask.
//
// Documentation: https://golang.org/s/x86manual#page=1659
func VGETMANTPS() {
	unsafe.Asm("VGETMANTPS", nil)
}

// VGETMANTSD
// Extract the normalized mantissa of the low float64 element in xmm3/m64 using imm8 for sign control and mantissa interval normalization. Store the mantissa to xmm1 under the writemask k1 and merge with the other elements of xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1662
func VGETMANTSD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VGETMANTSD", reg, evex, rm, imm)
}

// VHADDPD
// Horizontal add packed double-precision floating-point values from xmm2 and xmm3/mem.
// Horizontal add packed double-precision floating-point values from ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=529
func VHADDPD(reg, vex, rm interface{}) {
	unsafe.Asm("VHADDPD", reg, vex, rm)
}

// VHADDPS
// Horizontal add packed single-precision floating-point values from xmm2 and xmm3/mem.
// Horizontal add packed single-precision floating-point values from ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=532
func VHADDPS(reg, vex, rm interface{}) {
	unsafe.Asm("VHADDPS", reg, vex, rm)
}

// VHSUBPD
// Horizontal subtract packed double-precision floating-point values from xmm2 and xmm3/mem.
// Horizontal subtract packed double-precision floating-point values from ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=536
func VHSUBPD(reg, vex, rm interface{}) {
	unsafe.Asm("VHSUBPD", reg, vex, rm)
}

// VHSUBPS
// Horizontal subtract packed single-precision floating-point values from xmm2 and xmm3/mem.
// Horizontal subtract packed single-precision floating-point values from ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=539
func VHSUBPS(reg, vex, rm interface{}) {
	unsafe.Asm("VHSUBPS", reg, vex, rm)
}

// VINSERTF128
// Insert 128 bits of packed floating-point values from xmm3/m128 and the remaining values from ymm2 into ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1666
func VINSERTF128(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VINSERTF128", reg, vex, rm, imm)
}

// VINSERTF32X4
// Insert 128 bits of packed single-precision floating- point values from xmm3/m128 and the remaining values from ymm2 into ymm1 under writemask k1.
// Insert 128 bits of packed single-precision floating- point values from xmm3/m128 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1666
func VINSERTF32X4() {
	unsafe.Asm("VINSERTF32X4", nil)
}

// VINSERTF32X8
// Insert 256 bits of packed single-precision floating- point values from ymm3/m256 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1666
func VINSERTF32X8() {
	unsafe.Asm("VINSERTF32X8", nil)
}

// VINSERTF64X2
// Insert 128 bits of packed double-precision floating- point values from xmm3/m128 and the remaining values from ymm2 into ymm1 under writemask k1.
// Insert 128 bits of packed double-precision floating- point values from xmm3/m128 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1666
func VINSERTF64X2() {
	unsafe.Asm("VINSERTF64X2", nil)
}

// VINSERTF64X4
// Insert 256 bits of packed double-precision floating- point values from ymm3/m256 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1666
func VINSERTF64X4() {
	unsafe.Asm("VINSERTF64X4", nil)
}

// VINSERTI128
// Insert 128 bits of integer data from xmm3/m128 and the remaining values from ymm2 into ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1670
func VINSERTI128(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VINSERTI128", reg, vex, rm, imm)
}

// VINSERTI32X4
// Insert 128 bits of packed doubleword integer values from xmm3/m128 and the remaining values from ymm2 into ymm1 under writemask k1.
// Insert 128 bits of packed doubleword integer values from xmm3/m128 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1670
func VINSERTI32X4() {
	unsafe.Asm("VINSERTI32X4", nil)
}

// VINSERTI32X8
// Insert 256 bits of packed doubleword integer values from ymm3/m256 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1670
func VINSERTI32X8() {
	unsafe.Asm("VINSERTI32X8", nil)
}

// VINSERTI64X2
// Insert 128 bits of packed quadword integer values from xmm3/m128 and the remaining values from ymm2 into ymm1 under writemask k1.
// Insert 128 bits of packed quadword integer values from xmm3/m128 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1670
func VINSERTI64X2() {
	unsafe.Asm("VINSERTI64X2", nil)
}

// VINSERTI64X4
// Insert 256 bits of packed quadword integer values from ymm3/m256 and the remaining values from zmm2 into zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1670
func VINSERTI64X4() {
	unsafe.Asm("VINSERTI64X4", nil)
}

// VINSERTPS_RVMI
// Insert a single-precision floating-point value selected by imm8 from xmm3/m32 and merge with values in xmm2 at the specified destination element specified by imm8 and write out the result and zero out destination elements in xmm1 as indicated in imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=556
func VINSERTPS_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VINSERTPS", reg, vex, rm, imm)
}

// VINSERTPS_T1S
// Insert a single-precision floating-point value selected by imm8 from xmm3/m32 and merge with values in xmm2 at the specified destination element specified by imm8 and write out the result and zero out destination elements in xmm1 as indicated in imm8.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=556
func VINSERTPS_T1S(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VINSERTPS", reg, evex, rm, imm)
}

// VLDDQU
// Load unaligned packed integer values from mem to xmm1.
// Load unaligned packed integer values from mem to ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=620
func VLDDQU(reg, rm interface{}) {
	unsafe.Asm("VLDDQU", reg, rm)
}

// VLDMXCSR
// Load MXCSR register from m32.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=622
func VLDMXCSR(rm interface{}) {
	unsafe.Asm("VLDMXCSR", rm)
}

// VMASKMOVDQU
// Selectively write bytes from xmm1 to memory location using the byte mask in xmm2. The default memory location is specified by DS:DI/EDI/RDI.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=660
func VMASKMOVDQU(reg, rm interface{}) {
	unsafe.Asm("VMASKMOVDQU", reg, rm)
}

// VMASKMOVPD_MVR
// Conditionally store packed double-precision values from xmm2 using mask in xmm1.
// Conditionally store packed double-precision values from ymm2 using mask in ymm1.
//
// rm: ModRM:r/m (w)
// vex: VEX.vvvv (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1674
func VMASKMOVPD_MVR(rm, vex, reg interface{}) {
	unsafe.Asm("VMASKMOVPD", rm, vex, reg)
}

// VMASKMOVPD_RVM
// Conditionally load packed double-precision values from m128 using mask in xmm2 and store in xmm1.
// Conditionally load packed double-precision values from m256 using mask in ymm2 and store in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1674
func VMASKMOVPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMASKMOVPD", reg, vex, rm)
}

// VMASKMOVPS_MVR
// Conditionally store packed single-precision values from xmm2 using mask in xmm1.
// Conditionally store packed single-precision values from ymm2 using mask in ymm1.
//
// rm: ModRM:r/m (w)
// vex: VEX.vvvv (r)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1674
func VMASKMOVPS_MVR(rm, vex, reg interface{}) {
	unsafe.Asm("VMASKMOVPS", rm, vex, reg)
}

// VMASKMOVPS_RVM
// Conditionally load packed single-precision values from m128 using mask in xmm2 and store in xmm1.
// Conditionally load packed single-precision values from m256 using mask in ymm2 and store in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1674
func VMASKMOVPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMASKMOVPS", reg, vex, rm)
}

// VMAXPD_FV
// Return the maximum packed double-precision floating-point values between xmm2 and xmm3/m128/m64bcst and store result in xmm1 subject to writemask k1.
// Return the maximum packed double-precision floating-point values between ymm2 and ymm3/m256/m64bcst and store result in ymm1 subject to writemask k1.
// Return the maximum packed double-precision floating-point values between zmm2 and zmm3/m512/m64bcst and store result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=664
func VMAXPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMAXPD", reg, evex, rm)
}

// VMAXPD_RVM
// Return the maximum double-precision floating-point values between xmm2 and xmm3/m128.
// Return the maximum packed double-precision floating-point values between ymm2 and ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=664
func VMAXPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMAXPD", reg, vex, rm)
}

// VMAXPS_FV
// Return the maximum packed single-precision floating-point values between xmm2 and xmm3/m128/m32bcst and store result in xmm1 subject to writemask k1.
// Return the maximum packed single-precision floating-point values between ymm2 and ymm3/m256/m32bcst and store result in ymm1 subject to writemask k1.
// Return the maximum packed single-precision floating-point values between zmm2 and zmm3/m512/m32bcst and store result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=667
func VMAXPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMAXPS", reg, evex, rm)
}

// VMAXPS_RVM
// Return the maximum single-precision floating-point values between xmm2 and xmm3/mem.
// Return the maximum single-precision floating-point values between ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=667
func VMAXPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMAXPS", reg, vex, rm)
}

// VMAXSD_RVM
// Return the maximum scalar double-precision floating-point value between xmm3/m64 and xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=670
func VMAXSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMAXSD", reg, vex, rm)
}

// VMAXSD_T1S
// Return the maximum scalar double-precision floating-point value between xmm3/m64 and xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=670
func VMAXSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMAXSD", reg, evex, rm)
}

// VMAXSS_RVM
// Return the maximum scalar single-precision floating-point value between xmm3/m32 and xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=672
func VMAXSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMAXSS", reg, vex, rm)
}

// VMAXSS_T1S
// Return the maximum scalar single-precision floating-point value between xmm3/m32 and xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=672
func VMAXSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMAXSS", reg, evex, rm)
}

// VMINPD_FV
// Return the minimum packed double-precision floating-point values between xmm2 and xmm3/m128/m64bcst and store result in xmm1 subject to writemask k1.
// Return the minimum packed double-precision floating-point values between ymm2 and ymm3/m256/m64bcst and store result in ymm1 subject to writemask k1.
// Return the minimum packed double-precision floating-point values between zmm2 and zmm3/m512/m64bcst and store result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=675
func VMINPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMINPD", reg, evex, rm)
}

// VMINPD_RVM
// Return the minimum double-precision floating-point values between xmm2 and xmm3/mem.
// Return the minimum packed double-precision floating-point values between ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=675
func VMINPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMINPD", reg, vex, rm)
}

// VMINPS_FV
// Return the minimum packed single-precision floating-point values between xmm2 and xmm3/m128/m32bcst and store result in xmm1 subject to writemask k1.
// Return the minimum packed single-precision floating-point values between ymm2 and ymm3/m256/m32bcst and store result in ymm1 subject to writemask k1.
// Return the minimum packed single-precision floating-point values between zmm2 and zmm3/m512/m32bcst and store result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=678
func VMINPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMINPS", reg, evex, rm)
}

// VMINPS_RVM
// Return the minimum single-precision floating-point values between xmm2 and xmm3/mem.
// Return the minimum single double-precision floating-point values between ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=678
func VMINPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMINPS", reg, vex, rm)
}

// VMINSD_RVM
// Return the minimum scalar double-precision floating- point value between xmm3/m64 and xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=681
func VMINSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMINSD", reg, vex, rm)
}

// VMINSD_T1S
// Return the minimum scalar double-precision floating- point value between xmm3/m64 and xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=681
func VMINSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMINSD", reg, evex, rm)
}

// VMINSS_RVM
// Return the minimum scalar single-precision floating- point value between xmm3/m32 and xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=683
func VMINSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMINSS", reg, vex, rm)
}

// VMINSS_T1S
// Return the minimum scalar single-precision floating- point value between xmm3/m32 and xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=683
func VMINSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMINSS", reg, evex, rm)
}

// VMOVAPD_FVM_MR
// Move aligned packed double-precision floating- point values from xmm1 to xmm2/m128 using writemask k1.
// Move aligned packed double-precision floating- point values from ymm1 to ymm2/m256 using writemask k1.
// Move aligned packed double-precision floating- point values from zmm1 to zmm2/m512 using writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func VMOVAPD_FVM_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVAPD", rm, reg)
}

// VMOVAPD_FVM_RM
// Move aligned packed double-precision floating- point values from xmm2/m128 to xmm1 using writemask k1.
// Move aligned packed double-precision floating- point values from ymm2/m256 to ymm1 using writemask k1.
// Move aligned packed double-precision floating- point values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func VMOVAPD_FVM_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVAPD", reg, rm)
}

// VMOVAPD_MR
// Move aligned packed double-precision floating- point values from xmm1 to xmm2/mem.
// Move aligned packed double-precision floating- point values from ymm1 to ymm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func VMOVAPD_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVAPD", rm, reg)
}

// VMOVAPD_RM
// Move aligned packed double-precision floating- point values from xmm2/mem to xmm1.
// Move aligned packed double-precision floating- point values from ymm2/mem to ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=697
func VMOVAPD_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVAPD", reg, rm)
}

// VMOVAPS_FVM_MR
// Move aligned packed single-precision floating-point values from xmm1 to xmm2/m128 using writemask k1.
// Move aligned packed single-precision floating-point values from ymm1 to ymm2/m256 using writemask k1.
// Move aligned packed single-precision floating-point values from zmm1 to zmm2/m512 using writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func VMOVAPS_FVM_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVAPS", rm, reg)
}

// VMOVAPS_FVM_RM
// Move aligned packed single-precision floating-point values from xmm2/m128 to xmm1 using writemask k1.
// Move aligned packed single-precision floating-point values from ymm2/m256 to ymm1 using writemask k1.
// Move aligned packed single-precision floating-point values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func VMOVAPS_FVM_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVAPS", reg, rm)
}

// VMOVAPS_MR
// Move aligned packed single-precision floating-point values from xmm1 to xmm2/mem.
// Move aligned packed single-precision floating-point values from ymm1 to ymm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func VMOVAPS_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVAPS", rm, reg)
}

// VMOVAPS_RM
// Move aligned packed single-precision floating-point values from xmm2/mem to xmm1.
// Move aligned packed single-precision floating-point values from ymm2/mem to ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=701
func VMOVAPS_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVAPS", reg, rm)
}

// VMOVD
// Move doubleword from r/m32 to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=707
func VMOVD(reg, rm interface{}) {
	unsafe.Asm("VMOVD", reg, rm)
}

// VMOVDQA_MR
// Move aligned packed integer values from xmm1 to xmm2/mem.
// Move aligned packed integer values from ymm1 to ymm2/mem.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVDQA", rm, reg)
}

// VMOVDQA_RM
// Move aligned packed integer values from xmm2/mem to xmm1.
// Move aligned packed integer values from ymm2/mem to ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVDQA", reg, rm)
}

// VMOVDQA32_FVM_MR
// Move aligned packed doubleword integer values from xmm1 to xmm2/m128 using writemask k1.
// Move aligned packed doubleword integer values from ymm1 to ymm2/m256 using writemask k1.
// Move aligned packed doubleword integer values from zmm1 to zmm2/m512 using writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA32_FVM_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVDQA32", rm, reg)
}

// VMOVDQA32_FVM_RM
// Move aligned packed doubleword integer values from xmm2/m128 to xmm1 using writemask k1.
// Move aligned packed doubleword integer values from ymm2/m256 to ymm1 using writemask k1.
// Move aligned packed doubleword integer values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA32_FVM_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVDQA32", reg, rm)
}

// VMOVDQA64_FVM_MR
// Move aligned packed quadword integer values from xmm1 to xmm2/m128 using writemask k1.
// Move aligned packed quadword integer values from ymm1 to ymm2/m256 using writemask k1.
// Move aligned packed quadword integer values from zmm1 to zmm2/m512 using writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA64_FVM_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVDQA64", rm, reg)
}

// VMOVDQA64_FVM_RM
// Move aligned quadword integer values from xmm2/m128 to xmm1 using writemask k1.
// Move aligned quadword integer values from ymm2/m256 to ymm1 using writemask k1.
// Move aligned packed quadword integer values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=714
func VMOVDQA64_FVM_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVDQA64", reg, rm)
}

// VMOVHLPS
// Merge two packed single-precision floating-point values from high quadword of xmm3 and low quadword of xmm2.
//
// reg: ModRM:reg (w)
// vvvv: vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=728
func VMOVHLPS(reg, vvvv, rm interface{}) {
	unsafe.Asm("VMOVHLPS", reg, vvvv, rm)
}

// VMOVLHPS
// Merge two packed single-precision floating-point values from low quadword of xmm3 and low quadword of xmm2.
//
// reg: ModRM:reg (w)
// vvvv: vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=734
func VMOVLHPS(reg, vvvv, rm interface{}) {
	unsafe.Asm("VMOVLHPS", reg, vvvv, rm)
}

// VMOVMSKPD
// Extract 2-bit sign mask from xmm2 and store in reg. The upper bits of r32 or r64 are zeroed.
// Extract 4-bit sign mask from ymm2 and store in reg. The upper bits of r32 or r64 are zeroed.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=740
func VMOVMSKPD(reg, rm interface{}) {
	unsafe.Asm("VMOVMSKPD", reg, rm)
}

// VMOVMSKPS
// Extract 4-bit sign mask from xmm2 and store in reg. The upper bits of r32 or r64 are zeroed.
// Extract 8-bit sign mask from ymm2 and store in reg. The upper bits of r32 or r64 are zeroed.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=742
func VMOVMSKPS(reg, rm interface{}) {
	unsafe.Asm("VMOVMSKPS", reg, rm)
}

// VMOVNTDQ_FVM
// Move packed integer values in xmm1 to m128 using non- temporal hint.
// Move packed integer values in zmm1 to m256 using non- temporal hint.
// Move packed integer values in zmm1 to m512 using non- temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=746
func VMOVNTDQ_FVM(rm, reg interface{}) {
	unsafe.Asm("VMOVNTDQ", rm, reg)
}

// VMOVNTDQ_MR
// Move packed integer values in xmm1 to m128 using non- temporal hint.
// Move packed integer values in ymm1 to m256 using non- temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=746
func VMOVNTDQ_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVNTDQ", rm, reg)
}

// VMOVNTDQA_FVM
// Move 128-bit data from m128 to xmm using non-temporal hint if WC memory type.
// Move 256-bit data from m256 to ymm using non-temporal hint if WC memory type.
// Move 512-bit data from m512 to zmm using non-temporal hint if WC memory type.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=744
func VMOVNTDQA_FVM(reg, rm interface{}) {
	unsafe.Asm("VMOVNTDQA", reg, rm)
}

// VMOVNTDQA_RM
// Move double quadword from m128 to xmm using non- temporal hint if WC memory type.
// Move 256-bit data from m256 to ymm using non-temporal hint if WC memory type.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=744
func VMOVNTDQA_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVNTDQA", reg, rm)
}

// VMOVNTPD_FVM
// Move packed double-precision values in xmm1 to m128 using non-temporal hint.
// Move packed double-precision values in ymm1 to m256 using non-temporal hint.
// Move packed double-precision values in zmm1 to m512 using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=750
func VMOVNTPD_FVM(rm, reg interface{}) {
	unsafe.Asm("VMOVNTPD", rm, reg)
}

// VMOVNTPD_MR
// Move packed double-precision values in xmm1 to m128 using non-temporal hint.
// Move packed double-precision values in ymm1 to m256 using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=750
func VMOVNTPD_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVNTPD", rm, reg)
}

// VMOVNTPS_FVM
// Move packed single-precision values in xmm1 to m128 using non-temporal hint.
// Move packed single-precision values in ymm1 to m256 using non-temporal hint.
// Move packed single-precision values in zmm1 to m512 using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=752
func VMOVNTPS_FVM(rm, reg interface{}) {
	unsafe.Asm("VMOVNTPS", rm, reg)
}

// VMOVNTPS_MR
// Move packed single-precision values xmm1 to mem using non-temporal hint.
// Move packed single-precision values ymm1 to mem using non-temporal hint.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=752
func VMOVNTPS_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVNTPS", rm, reg)
}

// VMOVQ_MR
// Move quadword from xmm2 register to xmm1/m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func VMOVQ_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVQ", rm, reg)
}

// VMOVQ_RM
// Move quadword from xmm2 to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func VMOVQ_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVQ", reg, rm)
}

// VMOVQ_T1S_MR
// Move quadword from xmm2 register to xmm1/m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func VMOVQ_T1S_MR(rm, reg interface{}) {
	unsafe.Asm("VMOVQ", rm, reg)
}

// VMOVQ_T1S_RM
// Move quadword from xmm2/m64 to xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=755
func VMOVQ_T1S_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVQ", reg, rm)
}

// VMOVSHDUP_FVM
// Move odd index single-precision floating-point values from xmm2/m128 and duplicate each element into xmm1 under writemask.
// Move odd index single-precision floating-point values from ymm2/m256 and duplicate each element into ymm1 under writemask.
// Move odd index single-precision floating-point values from zmm2/m512 and duplicate each element into zmm1 under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=766
func VMOVSHDUP_FVM(reg, rm interface{}) {
	unsafe.Asm("VMOVSHDUP", reg, rm)
}

// VMOVSHDUP_RM
// Move odd index single-precision floating-point values from xmm2/mem and duplicate each element into xmm1.
// Move odd index single-precision floating-point values from ymm2/mem and duplicate each element into ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=766
func VMOVSHDUP_RM(reg, rm interface{}) {
	unsafe.Asm("VMOVSHDUP", reg, rm)
}

// VMPSADBW
// Sums absolute 8-bit integer difference of adjacent groups of 4 byte integers in xmm2 and xmm3/m128 and writes the results in xmm1. Starting offsets within xmm2 and xmm3/m128 are determined by imm8.
// Sums absolute 8-bit integer difference of adjacent groups of 4 byte integers in xmm2 and ymm3/m128 and writes the results in ymm1. Starting offsets within ymm2 and xmm3/m128 are determined by imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=788
func VMPSADBW(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VMPSADBW", reg, vex, rm, imm)
}

// VMULPD_FV
// Multiply packed double-precision floating-point values from xmm3/m128/m64bcst to xmm2 and store result in xmm1.
// Multiply packed double-precision floating-point values from ymm3/m256/m64bcst to ymm2 and store result in ymm1.
// Multiply packed double-precision floating-point values in zmm3/m512/m64bcst with zmm2 and store result in zmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=798
func VMULPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMULPD", reg, evex, rm)
}

// VMULPD_RVM
// Multiply packed double-precision floating-point values in xmm3/m128 with xmm2 and store result in xmm1.
// Multiply packed double-precision floating-point values in ymm3/m256 with ymm2 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=798
func VMULPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMULPD", reg, vex, rm)
}

// VMULPS_FV
// Multiply packed single-precision floating-point values from xmm3/m128/m32bcst to xmm2 and store result in xmm1.
// Multiply packed single-precision floating-point values from ymm3/m256/m32bcst to ymm2 and store result in ymm1.
// Multiply packed single-precision floating-point values in zmm3/m512/m32bcst with zmm2 and store result in zmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=801
func VMULPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VMULPS", reg, evex, rm)
}

// VMULPS_RVM
// Multiply packed single-precision floating-point values in xmm3/m128 with xmm2 and store result in xmm1.
// Multiply packed single-precision floating-point values in ymm3/m256 with ymm2 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=801
func VMULPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMULPS", reg, vex, rm)
}

// VMULSD_RVM
// Multiply the low double-precision floating-point value in xmm3/m64 by low double-precision floating-point value in xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=804
func VMULSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMULSD", reg, vex, rm)
}

// VMULSD_T1S
// Multiply the low double-precision floating-point value in xmm3/m64 by low double-precision floating-point value in xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=804
func VMULSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMULSD", reg, evex, rm)
}

// VMULSS_RVM
// Multiply the low single-precision floating-point value in xmm3/m32 by the low single-precision floating-point value in xmm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=806
func VMULSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VMULSS", reg, vex, rm)
}

// VMULSS_T1S
// Multiply the low single-precision floating-point value in xmm3/m32 by the low single-precision floating-point value in xmm2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=806
func VMULSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VMULSS", reg, evex, rm)
}

// VORPD_FV
// Return the bitwise logical OR of packed double-precision floating-point values in xmm2 and xmm3/m128/m64bcst subject to writemask k1.
// Return the bitwise logical OR of packed double-precision floating-point values in ymm2 and ymm3/m256/m64bcst subject to writemask k1.
// Return the bitwise logical OR of packed double-precision floating-point values in zmm2 and zmm3/m512/m64bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=820
func VORPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VORPD", reg, evex, rm)
}

// VORPD_RVM
// Return the bitwise logical OR of packed double-precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical OR of packed double-precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=820
func VORPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VORPD", reg, vex, rm)
}

// VORPS_FV
// Return the bitwise logical OR of packed single-precision floating-point values in xmm2 and xmm3/m128/m32bcst subject to writemask k1.
// Return the bitwise logical OR of packed single-precision floating-point values in ymm2 and ymm3/m256/m32bcst subject to writemask k1.
// Return the bitwise logical OR of packed single-precision floating-point values in zmm2 and zmm3/m512/m32bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=823
func VORPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VORPS", reg, evex, rm)
}

// VORPS_RVM
// Return the bitwise logical OR of packed single-precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical OR of packed single-precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=823
func VORPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VORPS", reg, vex, rm)
}

// VPABSB_FVM
// Compute the absolute value of bytes in xmm2/m128 and store UNSIGNED result in xmm1 using writemask k1.
// Compute the absolute value of bytes in ymm2/m256 and store UNSIGNED result in ymm1 using writemask k1.
// Compute the absolute value of bytes in zmm2/m512 and store UNSIGNED result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func VPABSB_FVM(reg, rm interface{}) {
	unsafe.Asm("VPABSB", reg, rm)
}

// VPABSB_RM
// Compute the absolute value of bytes in xmm2/m128 and store UNSIGNED result in xmm1.
// Compute the absolute value of bytes in ymm2/m256 and store UNSIGNED result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func VPABSB_RM(reg, rm interface{}) {
	unsafe.Asm("VPABSB", reg, rm)
}

// VPABSD
// Compute the absolute value of 32- bit integers in xmm2/m128 and store UNSIGNED result in xmm1.
// Compute the absolute value of 32-bit integers in ymm2/m256 and store UNSIGNED result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func VPABSD(reg, rm interface{}) {
	unsafe.Asm("VPABSD", reg, rm)
}

// VPABSW_FVM
// Compute the absolute value of 16-bit integers in xmm2/m128 and store UNSIGNED result in xmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func VPABSW_FVM(reg, rm interface{}) {
	unsafe.Asm("VPABSW", reg, rm)
}

// VPABSW_RM
// Compute the absolute value of 16- bit integers in xmm2/m128 and store UNSIGNED result in xmm1.
// Compute the absolute value of 16-bit integers in ymm2/m256 and store UNSIGNED result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=832
func VPABSW_RM(reg, rm interface{}) {
	unsafe.Asm("VPABSW", reg, rm)
}

// VPACKSSDW_FV
// Converts packed signed doubleword integers from xmm2 and from xmm3/m128/m32bcst into packed signed word integers in xmm1 using signed saturation under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func VPACKSSDW_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPACKSSDW", reg, evex, rm)
}

// VPACKSSDW_RVM
// Converts 4 packed signed doubleword integers from xmm2 and from xmm3/m128 into 8 packed signed word integers in xmm1 using signed saturation.
// Converts 8 packed signed doubleword integers from ymm2 and from ymm3/m256 into 16 packed signed word integers in ymm1using signed saturation.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func VPACKSSDW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPACKSSDW", reg, vex, rm)
}

// VPACKSSWB_FVM
// Converts packed signed word integers from xmm2 and from xmm3/m128 into packed signed byte integers in xmm1 using signed saturation under writemask k1.
// Converts packed signed word integers from ymm2 and from ymm3/m256 into packed signed byte integers in ymm1 using signed saturation under writemask k1.
// Converts packed signed word integers from zmm2 and from zmm3/m512 into packed signed byte integers in zmm1 using signed saturation under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func VPACKSSWB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPACKSSWB", reg, evex, rm)
}

// VPACKSSWB_RVM
// Converts 8 packed signed word integers from xmm2 and from xmm3/m128 into 16 packed signed byte integers in xmm1 using signed saturation.
// Converts 16 packed signed word integers from ymm2 and from ymm3/m256 into 32 packed signed byte integers in ymm1 using signed saturation.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=838
func VPACKSSWB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPACKSSWB", reg, vex, rm)
}

// VPACKUSDW_FV
// Convert packed signed doubleword integers from xmm2 and packed signed doubleword integers from xmm3/m128/m32bcst into packed unsigned word integers in xmm1 using unsigned saturation under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=846
func VPACKUSDW_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPACKUSDW", reg, evex, rm)
}

// VPACKUSDW_RVM
// Convert 4 packed signed doubleword integers from xmm2 and 4 packed signed doubleword integers from xmm3/m128 into 8 packed unsigned word integers in xmm1 using unsigned saturation.
// Convert 8 packed signed doubleword integers from ymm2 and 8 packed signed doubleword integers from ymm3/m256 into 16 packed unsigned word integers in ymm1 using unsigned saturation.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=846
func VPACKUSDW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPACKUSDW", reg, vex, rm)
}

// VPACKUSWB_FVM
// Converts signed word integers from xmm2 and signed word integers from xmm3/m128 into unsigned byte integers in xmm1 using unsigned saturation under writemask k1.
// Converts signed word integers from ymm2 and signed word integers from ymm3/m256 into unsigned byte integers in ymm1 using unsigned saturation under writemask k1.
// Converts signed word integers from zmm2 and signed word integers from zmm3/m512 into unsigned byte integers in zmm1 using unsigned saturation under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=851
func VPACKUSWB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPACKUSWB", reg, evex, rm)
}

// VPACKUSWB_RVM
// Converts 8 signed word integers from xmm2 and 8 signed word integers from xmm3/m128 into 16 unsigned byte integers in xmm1 using unsigned saturation.
// Converts 16 signed word integers from ymm2 and 16signed word integers from ymm3/m256 into 32 unsigned byte integers in ymm1 using unsigned saturation.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=851
func VPACKUSWB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPACKUSWB", reg, vex, rm)
}

// VPADDSB_FVM
// Add packed signed byte integers from xmm2, and xmm3/m128 and store the saturated results in xmm1 under writemask k1.
// Add packed signed byte integers from ymm2, and ymm3/m256 and store the saturated results in ymm1 under writemask k1.
// Add packed signed byte integers from zmm2, and zmm3/m512 and store the saturated results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func VPADDSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPADDSB", reg, evex, rm)
}

// VPADDSB_RVM
// Add packed signed byte integers from xmm3/m128 and xmm2 saturate the results.
// Add packed signed byte integers from ymm2, and ymm3/m256 and store the saturated results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func VPADDSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPADDSB", reg, vex, rm)
}

// VPADDSW_FVM
// Add packed signed word integers from xmm2, and xmm3/m128 and store the saturated results in xmm1 under writemask k1.
// Add packed signed word integers from ymm2, and ymm3/m256 and store the saturated results in ymm1 under writemask k1.
// Add packed signed word integers from zmm2, and zmm3/m512 and store the saturated results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func VPADDSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPADDSW", reg, evex, rm)
}

// VPADDSW_RVM
// Add packed signed word integers from xmm3/m128 and xmm2 and saturate the results.
// Add packed signed word integers from ymm2, and ymm3/m256 and store the saturated results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=863
func VPADDSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPADDSW", reg, vex, rm)
}

// VPADDUSB_FVM
// Add packed unsigned byte integers from xmm2, and xmm3/m128 and store the saturated results in xmm1 under writemask k1.
// Add packed unsigned byte integers from ymm2, and ymm3/m256 and store the saturated results in ymm1 under writemask k1.
// Add packed unsigned byte integers from zmm2, and zmm3/m512 and store the saturated results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func VPADDUSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPADDUSB", reg, evex, rm)
}

// VPADDUSB_RVM
// Add packed unsigned byte integers from xmm3/m128 to xmm2 and saturate the results.
// Add packed unsigned byte integers from ymm2, and ymm3/m256 and store the saturated results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func VPADDUSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPADDUSB", reg, vex, rm)
}

// VPADDUSW_FVM
// Add packed unsigned word integers from xmm2, and xmm3/m128 and store the saturated results in xmm1 under writemask k1.
// Add packed unsigned word integers from ymm2, and ymm3/m256 and store the saturated results in ymm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func VPADDUSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPADDUSW", reg, evex, rm)
}

// VPADDUSW_RVM
// Add packed unsigned word integers from xmm3/m128 to xmm2 and saturate the results.
// Add packed unsigned word integers from ymm2, and ymm3/m256 and store the saturated results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=867
func VPADDUSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPADDUSW", reg, vex, rm)
}

// VPALIGNR_FVM
// Concatenate xmm2 and xmm3/m128 into a 32- byte intermediate result, extract byte aligned result shifted to the right by constant value in imm8 and result is stored in xmm1.
// Concatenate pairs of 16 bytes in ymm2 and ymm3/m256 into 32-byte intermediate result, extract byte-aligned, 16-byte result shifted to the right by constant values in imm8 from each intermediate result, and two 16-byte results are stored in ymm1.
// Concatenate pairs of 16 bytes in zmm2 and zmm3/m512 into 32-byte intermediate result, extract byte-aligned, 16-byte result shifted to the right by constant values in imm8 from each intermediate result, and four 16-byte results are stored in zmm1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=871
func VPALIGNR_FVM(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VPALIGNR", reg, evex, rm, imm)
}

// VPALIGNR_RVMI
// Concatenate xmm2 and xmm3/m128, extract byte aligned result shifted to the right by constant value in imm8 and result is stored in xmm1.
// Concatenate pairs of 16 bytes in ymm2 and ymm3/m256 into 32-byte intermediate result, extract byte-aligned, 16-byte result shifted to the right by constant values in imm8 from each intermediate result, and two 16-byte results are stored in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=871
func VPALIGNR_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPALIGNR", reg, vex, rm, imm)
}

// VPAND
// Bitwise AND of xmm3/m128 and xmm.
// Bitwise AND of ymm2, and ymm3/m256 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=875
func VPAND(reg, vex, rm interface{}) {
	unsafe.Asm("VPAND", reg, vex, rm)
}

// VPANDD
// Bitwise AND of packed doubleword integers in xmm2 and xmm3/m128/m32bcst and store result in xmm1 using writemask k1.
// Bitwise AND of packed doubleword integers in ymm2 and ymm3/m256/m32bcst and store result in ymm1 using writemask k1.
// Bitwise AND of packed doubleword integers in zmm2 and zmm3/m512/m32bcst and store result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=875
func VPANDD(reg, evex, rm interface{}) {
	unsafe.Asm("VPANDD", reg, evex, rm)
}

// VPANDN
// Bitwise AND NOT of xmm3/m128 and xmm2.
// Bitwise AND NOT of ymm2, and ymm3/m256 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=878
func VPANDN(reg, vex, rm interface{}) {
	unsafe.Asm("VPANDN", reg, vex, rm)
}

// VPANDND
// Bitwise AND NOT of packed doubleword integers in xmm2 and xmm3/m128/m32bcst and store result in xmm1 using writemask k1.
// Bitwise AND NOT of packed doubleword integers in ymm2 and ymm3/m256/m32bcst and store result in ymm1 using writemask k1.
// Bitwise AND NOT of packed doubleword integers in zmm2 and zmm3/m512/m32bcst and store result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=878
func VPANDND(reg, evex, rm interface{}) {
	unsafe.Asm("VPANDND", reg, evex, rm)
}

// VPANDNQ
// Bitwise AND NOT of packed quadword integers in xmm2 and xmm3/m128/m64bcst and store result in xmm1 using writemask k1.
// Bitwise AND NOT of packed quadword integers in ymm2 and ymm3/m256/m64bcst and store result in ymm1 using writemask k1.
// Bitwise AND NOT of packed quadword integers in zmm2 and zmm3/m512/m64bcst and store result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=878
func VPANDNQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPANDNQ", reg, evex, rm)
}

// VPANDQ
// Bitwise AND of packed quadword integers in xmm2 and xmm3/m128/m64bcst and store result in xmm1 using writemask k1.
// Bitwise AND of packed quadword integers in ymm2 and ymm3/m256/m64bcst and store result in ymm1 using writemask k1.
// Bitwise AND of packed quadword integers in zmm2 and zmm3/m512/m64bcst and store result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=875
func VPANDQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPANDQ", reg, evex, rm)
}

// VPAVGB_FVM
// Average packed unsigned byte integers from xmm2, and xmm3/m128 with rounding and store to xmm1 under writemask k1.
// Average packed unsigned byte integers from ymm2, and ymm3/m256 with rounding and store to ymm1 under writemask k1.
// Average packed unsigned byte integers from zmm2, and zmm3/m512 with rounding and store to zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=882
func VPAVGB_FVM() {
	unsafe.Asm("VPAVGB", nil)
}

// VPAVGB_RVM
// Average packed unsigned byte integers from xmm3/m128 and xmm2 with rounding.
// Average packed unsigned byte integers from ymm2, and ymm3/m256 with rounding and store to ymm1.
//
// Documentation: https://golang.org/s/x86manual#page=882
func VPAVGB_RVM() {
	unsafe.Asm("VPAVGB", nil)
}

// VPAVGW_FVM
// Average packed unsigned word integers from xmm2, xmm3/m128 with rounding to xmm1 under writemask k1.
// Average packed unsigned word integers from ymm2, ymm3/m256 with rounding to ymm1 under writemask k1.
// Average packed unsigned word integers from zmm2, zmm3/m512 with rounding to zmm1 under writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=882
func VPAVGW_FVM() {
	unsafe.Asm("VPAVGW", nil)
}

// VPAVGW_RVM
// Average packed unsigned word integers from xmm3/m128 and xmm2 with rounding.
// Average packed unsigned word integers from ymm2, ymm3/m256 with rounding to ymm1.
//
// Documentation: https://golang.org/s/x86manual#page=882
func VPAVGW_RVM() {
	unsafe.Asm("VPAVGW", nil)
}

// VPBLENDD
// Select dwords from xmm2 and xmm3/m128 from mask specified in imm8 and store the values into xmm1.
// Select dwords from ymm2 and ymm3/m256 from mask specified in imm8 and store the values into ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1677
func VPBLENDD(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPBLENDD", reg, vex, rm, imm)
}

// VPBLENDMB
// Blend byte integer vector xmm2 and byte vector xmm3/m128 and store the result in xmm1, under control mask.
// Blend byte integer vector ymm2 and byte vector ymm3/m256 and store the result in ymm1, under control mask.
// Blend byte integer vector zmm2 and byte vector zmm3/m512 and store the result in zmm1, under control mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1679
func VPBLENDMB(reg, evex, rm interface{}) {
	unsafe.Asm("VPBLENDMB", reg, evex, rm)
}

// VPBLENDMD
// Blend doubleword integer vector xmm2 and doubleword vector xmm3/m128/m32bcst and store the result in xmm1, under control mask.
// Blend doubleword integer vector ymm2 and doubleword vector ymm3/m256/m32bcst and store the result in ymm1, under control mask.
// Blend doubleword integer vector zmm2 and doubleword vector zmm3/m512/m32bcst and store the result in zmm1, under control mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1681
func VPBLENDMD(reg, evex, rm interface{}) {
	unsafe.Asm("VPBLENDMD", reg, evex, rm)
}

// VPBLENDMQ
// Blend quadword integer vector xmm2 and quadword vector xmm3/m128/m64bcst and store the result in xmm1, under control mask.
// Blend quadword integer vector ymm2 and quadword vector ymm3/m256/m64bcst and store the result in ymm1, under control mask.
// Blend quadword integer vector zmm2 and quadword vector zmm3/m512/m64bcst and store the result in zmm1, under control mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1681
func VPBLENDMQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPBLENDMQ", reg, evex, rm)
}

// VPBLENDMW
// Blend word integer vector xmm2 and word vector xmm3/m128 and store the result in xmm1, under control mask.
// Blend word integer vector ymm2 and word vector ymm3/m256 and store the result in ymm1, under control mask.
// Blend word integer vector zmm2 and word vector zmm3/m512 and store the result in zmm1, under control mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1679
func VPBLENDMW(reg, evex, rm interface{}) {
	unsafe.Asm("VPBLENDMW", reg, evex, rm)
}

// VPBLENDVB
// Select byte values from xmm2 and xmm3/m128 using mask bits in the specified mask register, xmm4, and store the values into xmm1.
// Select byte values from ymm2 and ymm3/m256 from mask specified in the high bit of each byte in ymm4 and store the values into ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8[7:4]
//
// Documentation: https://golang.org/s/x86manual#page=886
func VPBLENDVB(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPBLENDVB", reg, vex, rm, imm)
}

// VPBLENDW
// Select words from xmm2 and xmm3/m128 from mask specified in imm8 and store the values into xmm1.
// Select words from ymm2 and ymm3/m256 from mask specified in imm8 and store the values into ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=890
func VPBLENDW(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPBLENDW", reg, vex, rm, imm)
}

// VPBROADCASTB_RM
// Broadcast a byte integer in the source operand to sixteen locations in xmm1.
// Broadcast a byte integer in the source operand to thirty-two locations in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTB_RM(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTB", reg, rm)
}

// VPBROADCASTB_T1S
// Broadcast a byte integer in the source operand to locations in xmm1 subject to writemask k1.
// Broadcast a byte integer in the source operand to locations in ymm1 subject to writemask k1.
// Broadcast a byte integer in the source operand to 64 locations in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTB_T1S() {
	unsafe.Asm("VPBROADCASTB", nil)
}

// VPBROADCASTD_RM
// Broadcast a dword integer in the source operand to four locations in xmm1.
// Broadcast a dword integer in the source operand to eight locations in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTD_RM(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTD", reg, rm)
}

// VPBROADCASTD_T1S
// Broadcast a dword integer in the source operand to locations in xmm1 subject to writemask k1.
// Broadcast a dword integer in the source operand to locations in ymm1 subject to writemask k1.
// Broadcast a dword integer in the source operand to locations in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTD_T1S() {
	unsafe.Asm("VPBROADCASTD", nil)
}

// VPBROADCASTMB2Q
// Broadcast low byte value in k1 to two locations in xmm1.
// Broadcast low byte value in k1 to four locations in ymm1.
// Broadcast low byte value in k1 to eight locations in zmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1375
func VPBROADCASTMB2Q(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTMB2Q", reg, rm)
}

// VPBROADCASTMW2D
// Broadcast low word value in k1 to four locations in xmm1.
// Broadcast low word value in k1 to eight locations in ymm1.
// Broadcast low word value in k1 to sixteen locations in zmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1375
func VPBROADCASTMW2D(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTMW2D", reg, rm)
}

// VPBROADCASTQ_RM
// Broadcast a qword element in source operand to two locations in xmm1.
// Broadcast a qword element in source operand to four locations in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTQ", reg, rm)
}

// VPBROADCASTQ_T1S
// Broadcast a qword element in source operand to locations in xmm1 subject to writemask k1.
// Broadcast a qword element in source operand to locations in ymm1 subject to writemask k1.
// Broadcast a qword element in source operand to locations in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTQ_T1S() {
	unsafe.Asm("VPBROADCASTQ", nil)
}

// VPBROADCASTW_RM
// Broadcast a word integer in the source operand to eight locations in xmm1.
// Broadcast a word integer in the source operand to sixteen locations in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTW_RM(reg, rm interface{}) {
	unsafe.Asm("VPBROADCASTW", reg, rm)
}

// VPBROADCASTW_T1S
// Broadcast a word integer in the source operand to locations in xmm1 subject to writemask k1.
// Broadcast a word integer in the source operand to locations in ymm1 subject to writemask k1.
// Broadcast a word integer in the source operand to 32 locations in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1687
func VPBROADCASTW_T1S() {
	unsafe.Asm("VPBROADCASTW", nil)
}

// VPCLMULQDQ
// Carry-less multiplication of one quadword of xmm2 by one quadword of xmm3/m128, stores the 128-bit result in xmm1. The imme- diate is used to determine which quadwords of xmm2 and xmm3/m128 should be used.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=893
func VPCLMULQDQ(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPCLMULQDQ", reg, vex, rm, imm)
}

// VPCMPEQB_FVM
// Compare packed bytes in xmm3/m128 and xmm2 for equality and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func VPCMPEQB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPEQB", reg, evex, rm)
}

// VPCMPEQB_RVM
// Compare packed bytes in xmm3/m128 and xmm2 for equality.
// Compare packed bytes in ymm3/m256 and ymm2 for equality.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func VPCMPEQB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPEQB", reg, vex, rm)
}

// VPCMPEQD_FV
// Compare Equal between int32 vector xmm2 and int32 vector xmm3/m128/m32bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Equal between int32 vector ymm2 and int32 vector ymm3/m256/m32bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Equal between int32 vectors in zmm2 and zmm3/m512/m32bcst, and set destination k1 according to the comparison results under writemask k2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func VPCMPEQD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPEQD", reg, evex, rm)
}

// VPCMPEQD_RVM
// Compare packed doublewords in xmm3/m128 and xmm2 for equality.
// Compare packed doublewords in ymm3/m256 and ymm2 for equality.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func VPCMPEQD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPEQD", reg, vex, rm)
}

// VPCMPEQQ_FV
// Compare Equal between int64 vector xmm2 and int64 vector xmm3/m128/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Equal between int64 vector ymm2 and int64 vector ymm3/m256/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Equal between int64 vector zmm2 and int64 vector zmm3/m512/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=902
func VPCMPEQQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPEQQ", reg, evex, rm)
}

// VPCMPEQQ_RVM
// Compare packed quadwords in xmm3/m128 and xmm2 for equality.
// Compare packed quadwords in ymm3/m256 and ymm2 for equality.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=902
func VPCMPEQQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPEQQ", reg, vex, rm)
}

// VPCMPEQW
// Compare packed words in xmm3/m128 and xmm2 for equality.
// Compare packed words in ymm3/m256 and ymm2 for equality.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=896
func VPCMPEQW(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPEQW", reg, vex, rm)
}

// VPCMPESTRI
// Perform a packed comparison of string data with explicit lengths, generating an index, and storing the result in ECX.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=905
func VPCMPESTRI(reg, rm, imm interface{}) {
	unsafe.Asm("VPCMPESTRI", reg, rm, imm)
}

// VPCMPESTRM
// Perform a packed comparison of string data with explicit lengths, generating a mask, and storing the result in XMM0.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=907
func VPCMPESTRM(reg, rm, imm interface{}) {
	unsafe.Asm("VPCMPESTRM", reg, rm, imm)
}

// VPCMPGTB_FVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 for greater than, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare packed signed byte integers in ymm2 and ymm3/m256 for greater than, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func VPCMPGTB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPGTB", reg, evex, rm)
}

// VPCMPGTB_RVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 for greater than.
// Compare packed signed byte integers in ymm2 and ymm3/m256 for greater than.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func VPCMPGTB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPGTB", reg, vex, rm)
}

// VPCMPGTD_FV
// Compare Greater between int32 vector xmm2 and int32 vector xmm3/m128/m32bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Greater between int32 vector ymm2 and int32 vector ymm3/m256/m32bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Greater between int32 elements in zmm2 and zmm3/m512/m32bcst, and set destination k1 according to the comparison results under writemask. k2.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func VPCMPGTD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPGTD", reg, evex, rm)
}

// VPCMPGTD_RVM
// Compare packed signed doubleword integers in xmm2 and xmm3/m128 for greater than.
// Compare packed signed doubleword integers in ymm2 and ymm3/m256 for greater than.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func VPCMPGTD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPGTD", reg, vex, rm)
}

// VPCMPGTQ_FV
// Compare Greater between int64 vector xmm2 and int64 vector xmm3/m128/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Greater between int64 vector ymm2 and int64 vector ymm3/m256/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
// Compare Greater between int64 vector zmm2 and int64 vector zmm3/m512/m64bcst, and set vector mask k1 to reflect the zero/nonzero status of each element of the result, under writemask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=915
func VPCMPGTQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPCMPGTQ", reg, evex, rm)
}

// VPCMPGTQ_RVM
// Compare packed signed qwords in xmm2 and xmm3/m128 for greater than.
// Compare packed signed qwords in ymm2 and ymm3/m256 for greater than.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=915
func VPCMPGTQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPGTQ", reg, vex, rm)
}

// VPCMPGTW
// Compare packed signed word integers in xmm2 and xmm3/m128 for greater than.
// Compare packed signed word integers in ymm2 and ymm3/m256 for greater than.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=909
func VPCMPGTW(reg, vex, rm interface{}) {
	unsafe.Asm("VPCMPGTW", reg, vex, rm)
}

// VPCMPISTRI
// Perform a packed comparison of string data with implicit lengths, generating an index, and storing the result in ECX.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=918
func VPCMPISTRI(reg, rm, imm interface{}) {
	unsafe.Asm("VPCMPISTRI", reg, rm, imm)
}

// VPCMPISTRM
// Perform a packed comparison of string data with implicit lengths, generating a Mask, and storing the result in XMM0.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=920
func VPCMPISTRM(reg, rm, imm interface{}) {
	unsafe.Asm("VPCMPISTRM", reg, rm, imm)
}

// VPCMPQ
// Compare packed signed quadword integer values in xmm3/m128/m64bcst and xmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed signed quadword integer values in ymm3/m256/m64bcst and ymm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed signed quadword integer values in zmm3/m512/m64bcst and zmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1701
func VPCMPQ(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VPCMPQ", reg, evex, rm, imm)
}

// VPCMPUQ
// Compare packed unsigned quadword integer values in xmm3/m128/m64bcst and xmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed unsigned quadword integer values in ymm3/m256/m64bcst and ymm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed unsigned quadword integer values in zmm3/m512/m64bcst and zmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1701
func VPCMPUQ(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VPCMPUQ", reg, evex, rm, imm)
}

// VPCMPUW
// Compare packed unsigned word integers in xmm3/m128 and xmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed unsigned word integers in ymm3/m256 and ymm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// vvvv: vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1704
func VPCMPUW(reg, vvvv, rm, imm interface{}) {
	unsafe.Asm("VPCMPUW", reg, vvvv, rm, imm)
}

// VPCMPW
// Compare packed signed word integers in xmm3/m128 and xmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed signed word integers in ymm3/m256 and ymm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
// Compare packed signed word integers in zmm3/m512 and zmm2 using bits 2:0 of imm8 as a comparison predicate with writemask k2 and leave the result in mask register k1.
//
// reg: ModRM:reg (w)
// vvvv: vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1704
func VPCMPW(reg, vvvv, rm, imm interface{}) {
	unsafe.Asm("VPCMPW", reg, vvvv, rm, imm)
}

// VPCOMPRESSD
// Compress packed doubleword integer values from xmm2 to xmm1/m128 using controlmask k1.
// Compress packed doubleword integer values from ymm2 to ymm1/m256 using controlmask k1.
// Compress packed doubleword integer values from zmm2 to zmm1/m512 using controlmask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1707
func VPCOMPRESSD(rm, reg interface{}) {
	unsafe.Asm("VPCOMPRESSD", rm, reg)
}

// VPERM2F128
// Permute 128-bit floating-point fields in ymm2 and ymm3/mem using controls from imm8 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1714
func VPERM2F128(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPERM2F128", reg, vex, rm, imm)
}

// VPERM2I128
// Permute 128-bit integer data in ymm2 and ymm3/mem using controls from imm8 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1716
func VPERM2I128(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPERM2I128", reg, vex, rm, imm)
}

// VPERMD_FV
// Permute doublewords in ymm3/m256/m32bcst using indexes in ymm2 and store the result in ymm1 using writemask k1.
// Permute doublewords in zmm3/m512/m32bcst using indices in zmm2 and store the result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1718
func VPERMD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPERMD", reg, evex, rm)
}

// VPERMD_RVM
// Permute doublewords in ymm3/m256 using indices in ymm2 and store the result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1718
func VPERMD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPERMD", reg, vex, rm)
}

// VPERMPS_FV
// Permute single-precision floating-point elements in ymm3/m256/m32bcst using indexes in ymm2 and store the result in ymm1 subject to write mask k1.
// Permute single-precision floating-point values in zmm3/m512/m32bcst using indices in zmm2 and store the result in zmm1 subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1740
func VPERMPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPERMPS", reg, evex, rm)
}

// VPERMPS_RVM
// Permute single-precision floating-point elements in ymm3/m256 using indices in ymm2 and store the result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1740
func VPERMPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPERMPS", reg, vex, rm)
}

// VPERMW
// Permute word integers in xmm3/m128 using indexes in xmm2 and store the result in xmm1 using writemask k1.
// Permute word integers in ymm3/m256 using indexes in ymm2 and store the result in ymm1 using writemask k1.
// Permute word integers in zmm3/m512 using indexes in zmm2 and store the result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1718
func VPERMW(reg, vex, rm interface{}) {
	unsafe.Asm("VPERMW", reg, vex, rm)
}

// VPEXPANDD
// Expand packed double-word integer values from xmm2/m128 to xmm1 using writemask k1.
// Expand packed double-word integer values from ymm2/m256 to ymm1 using writemask k1.
// Expand packed double-word integer values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1746
func VPEXPANDD(reg, rm interface{}) {
	unsafe.Asm("VPEXPANDD", reg, rm)
}

// VPEXPANDQ
// Expand packed quad-word integer values from xmm2/m128 to xmm1 using writemask k1.
// Expand packed quad-word integer values from ymm2/m256 to ymm1 using writemask k1.
// Expand packed quad-word integer values from zmm2/m512 to zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1748
func VPEXPANDQ(reg, rm interface{}) {
	unsafe.Asm("VPEXPANDQ", reg, rm)
}

// VPEXTRB_MRI
// Extract a byte integer value from xmm2 at the source byte offset specified by imm8 into reg or m8. The upper bits of r64/r32 is filled with zeros.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRB_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("VPEXTRB", rm, reg, imm)
}

// VPEXTRB_T1S_MRI
// Extract a byte integer value from xmm2 at the source byte offset specified by imm8 into reg or m8. The upper bits of r64/r32 is filled with zeros.
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRB_T1S_MRI() {
	unsafe.Asm("VPEXTRB", nil)
}

// VPEXTRD_MRI
// Extract a dword integer value from xmm2 at the source dword offset specified by imm8 into r32/m32.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRD_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("VPEXTRD", rm, reg, imm)
}

// VPEXTRD_T1S_MRI
// Extract a dword integer value from xmm2 at the source dword offset specified by imm8 into r32/m32.
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRD_T1S_MRI() {
	unsafe.Asm("VPEXTRD", nil)
}

// VPEXTRQ_MRI
// Extract a qword integer value from xmm2 at the source dword offset specified by imm8 into r64/m64.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRQ_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("VPEXTRQ", rm, reg, imm)
}

// VPEXTRQ_T1S_MRI
// Extract a qword integer value from xmm2 at the source dword offset specified by imm8 into r64/m64.
//
// Documentation: https://golang.org/s/x86manual#page=926
func VPEXTRQ_T1S_MRI() {
	unsafe.Asm("VPEXTRQ", nil)
}

// VPEXTRW_MRI
// Extract a word integer value from xmm2 at the source word offset specified by imm8 into reg or m16. The upper bits of r64/r32 is filled with zeros.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=929
func VPEXTRW_MRI(rm, reg, imm interface{}) {
	unsafe.Asm("VPEXTRW", rm, reg, imm)
}

// VPEXTRW_RMI
// Extract the word specified by imm8 from xmm1 and move it to reg, bits 15:0. Zero- extend the result. The upper bits of r64/r32 is filled with zeros.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=929
func VPEXTRW_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("VPEXTRW", reg, rm, imm)
}

// VPGATHERDD_RMV
// Using dword indices specified in vm32x, gather dword val- ues from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using dword indices specified in vm32y, gather dword from memory conditioned on mask specified by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1629
func VPGATHERDD_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VPGATHERDD", reg, vsib, vex)
}

// VPGATHERDD_T1S
// Using signed dword indices, gather dword values from memory using writemask k1 for merging-masking.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1633
func VPGATHERDD_T1S(reg, v interface{}) {
	unsafe.Asm("VPGATHERDD", reg, v)
}

// VPGATHERDQ_RMV
// Using dword indices specified in vm32x, gather qword val- ues from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using dword indices specified in vm32x, gather qword val- ues from memory conditioned on mask specified by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1636
func VPGATHERDQ_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VPGATHERDQ", reg, vsib, vex)
}

// VPGATHERDQ_T1S
// Using signed dword indices, gather quadword values from memory using writemask k1 for merging-masking.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1633
func VPGATHERDQ_T1S(reg, v interface{}) {
	unsafe.Asm("VPGATHERDQ", reg, v)
}

// VPGATHERQD_RMV
// Using qword indices specified in vm64x, gather dword val- ues from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using qword indices specified in vm64y, gather dword val- ues from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1629
func VPGATHERQD_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VPGATHERQD", reg, vsib, vex)
}

// VPGATHERQD_T1S
// Using signed qword indices, gather dword values from memory using writemask k1 for merging-masking.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1641
func VPGATHERQD_T1S(reg, v interface{}) {
	unsafe.Asm("VPGATHERQD", reg, v)
}

// VPGATHERQQ_RMV
// Using qword indices specified in vm64x, gather qword val- ues from memory conditioned on mask specified by xmm2. Conditionally gathered elements are merged into xmm1.
// Using qword indices specified in vm64y, gather qword val- ues from memory conditioned on mask specified by ymm2. Conditionally gathered elements are merged into ymm1.
//
// reg: ModRM:reg (r, w)
// vsib: vsib (r)
// vex: VEX.vvvv (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1636
func VPGATHERQQ_RMV(reg, vsib, vex interface{}) {
	unsafe.Asm("VPGATHERQQ", reg, vsib, vex)
}

// VPGATHERQQ_T1S
// Using signed qword indices, gather quadword values from memory using writemask k1 for merging-masking.
//
// reg: ModRM:reg (w)
// v: VectorReg(R): VSIB:index
//
// Documentation: https://golang.org/s/x86manual#page=1641
func VPGATHERQQ_T1S(reg, v interface{}) {
	unsafe.Asm("VPGATHERQQ", reg, v)
}

// VPHADDD
// Add 32-bit integers horizontally, pack to xmm1.
// Add 32-bit signed integers horizontally, pack to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=932
func VPHADDD(reg, vex, rm interface{}) {
	unsafe.Asm("VPHADDD", reg, vex, rm)
}

// VPHADDSW
// Add 16-bit signed integers horizontally, pack saturated integers to xmm1.
// Add 16-bit signed integers horizontally, pack saturated integers to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=936
func VPHADDSW(reg, vex, rm interface{}) {
	unsafe.Asm("VPHADDSW", reg, vex, rm)
}

// VPHADDW
// Add 16-bit integers horizontally, pack to xmm1.
// Add 16-bit signed integers horizontally, pack to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=932
func VPHADDW(reg, vex, rm interface{}) {
	unsafe.Asm("VPHADDW", reg, vex, rm)
}

// VPHMINPOSUW
// Find the minimum unsigned word in xmm2/m128 and place its value in the low word of xmm1 and its index in the second- lowest word of xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=938
func VPHMINPOSUW(reg, rm interface{}) {
	unsafe.Asm("VPHMINPOSUW", reg, rm)
}

// VPHSUBD
// Subtract 32-bit signed integers horizontally, pack to xmm1.
// Subtract 32-bit signed integers horizontally, pack to ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=940
func VPHSUBD(reg, vex, rm interface{}) {
	unsafe.Asm("VPHSUBD", reg, vex, rm)
}

// VPHSUBSW
// Subtract 16-bit signed integer horizontally, pack saturated integers to xmm1.
// Subtract 16-bit signed integer horizontally, pack saturated integers to ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=943
func VPHSUBSW(reg, vex, rm interface{}) {
	unsafe.Asm("VPHSUBSW", reg, vex, rm)
}

// VPHSUBW
// Subtract 16-bit signed integers horizontally, pack to xmm1.
// Subtract 16-bit signed integers horizontally, pack to ymm1.
//
// reg: ModRM:reg (r, w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=940
func VPHSUBW(reg, vex, rm interface{}) {
	unsafe.Asm("VPHSUBW", reg, vex, rm)
}

// VPINSRB_RVMI
// Merge a byte integer value from r32/m8 and rest from xmm2 into xmm1 at the byte offset in imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRB_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPINSRB", reg, vex, rm, imm)
}

// VPINSRB_T1S__RVMI
// Merge a byte integer value from r32/m8 and rest from xmm2 into xmm1 at the byte offset in imm8.
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRB_T1S__RVMI() {
	unsafe.Asm("VPINSRB", nil)
}

// VPINSRD_RVMI
// Insert a dword integer value from r32/m32 and rest from xmm2 into xmm1 at the dword offset in imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRD_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPINSRD", reg, vex, rm, imm)
}

// VPINSRD_T1S__RVMI
// Insert a dword integer value from r32/m32 and rest from xmm2 into xmm1 at the dword offset in imm8.
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRD_T1S__RVMI() {
	unsafe.Asm("VPINSRD", nil)
}

// VPINSRQ_RVMI
// Insert a qword integer value from r64/m64 and rest from xmm2 into xmm1 at the qword offset in imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRQ_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPINSRQ", reg, vex, rm, imm)
}

// VPINSRQ_T1S__RVMI
// Insert a qword integer value from r64/m64 and rest from xmm2 into xmm1 at the qword offset in imm8.
//
// Documentation: https://golang.org/s/x86manual#page=945
func VPINSRQ_T1S__RVMI() {
	unsafe.Asm("VPINSRQ", nil)
}

// VPINSRW_RVMI
// Insert a word integer value from r32/m16 and rest from xmm2 into xmm1 at the word offset in imm8.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=948
func VPINSRW_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VPINSRW", reg, vex, rm, imm)
}

// VPINSRW_T1S__RVMI
// Insert a word integer value from r32/m16 and rest from xmm2 into xmm1 at the word offset in imm8.
//
// Documentation: https://golang.org/s/x86manual#page=948
func VPINSRW_T1S__RVMI() {
	unsafe.Asm("VPINSRW", nil)
}

// VPMADDUBSW_FVM
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to xmm1 under writemask k1.
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to ymm1 under writemask k1.
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=950
func VPMADDUBSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMADDUBSW", reg, evex, rm)
}

// VPMADDUBSW_RVM
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to xmm1.
// Multiply signed and unsigned bytes, add horizontal pair of signed words, pack saturated signed-words to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=950
func VPMADDUBSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMADDUBSW", reg, vex, rm)
}

// VPMADDWD_FVM
// Multiply the packed word integers in xmm2 by the packed word integers in xmm3/m128, add adjacent doubleword results, and store in xmm1 under writemask k1.
// Multiply the packed word integers in ymm2 by the packed word integers in ymm3/m256, add adjacent doubleword results, and store in ymm1 under writemask k1.
// Multiply the packed word integers in zmm2 by the packed word integers in zmm3/m512, add adjacent doubleword results, and store in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=953
func VPMADDWD_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMADDWD", reg, evex, rm)
}

// VPMADDWD_RVM
// Multiply the packed word integers in xmm2 by the packed word integers in xmm3/m128, add adjacent doubleword results, and store in xmm1.
// Multiply the packed word integers in ymm2 by the packed word integers in ymm3/m256, add adjacent doubleword results, and store in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=953
func VPMADDWD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMADDWD", reg, vex, rm)
}

// VPMASKMOVD_MVR
// Conditionally store dword values from xmm2 using mask in xmm1.
// Conditionally store dword values from ymm2 using mask in ymm1.
//
// rm: ModRM:r/m (w)
// vex: VEX.vvvv
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1753
func VPMASKMOVD_MVR(rm, vex, reg interface{}) {
	unsafe.Asm("VPMASKMOVD", rm, vex, reg)
}

// VPMASKMOVD_RVM
// Conditionally load dword values from m256 using mask in ymm2 and store in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1753
func VPMASKMOVD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMASKMOVD", reg, vex, rm)
}

// VPMASKMOVQ_MVR
// Conditionally store qword values from xmm2 using mask in xmm1.
// Conditionally store qword values from ymm2 using mask in ymm1.
//
// rm: ModRM:r/m (w)
// vex: VEX.vvvv
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1753
func VPMASKMOVQ_MVR(rm, vex, reg interface{}) {
	unsafe.Asm("VPMASKMOVQ", rm, vex, reg)
}

// VPMASKMOVQ_RVM
// Conditionally load qword values from m128 using mask in xmm2 and store in xmm1.
// Conditionally load qword values from m256 using mask in ymm2 and store in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1753
func VPMASKMOVQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMASKMOVQ", reg, vex, rm)
}

// VPMAXSB_FVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1 under writemask k1.
// Compare packed signed byte integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1 under writemask k1.
// Compare packed signed byte integers in zmm2 and zmm3/m512 and store packed maximum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXSB", reg, evex, rm)
}

// VPMAXSB_RVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1.
// Compare packed signed byte integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMAXSB", reg, vex, rm)
}

// VPMAXSD_FV
// Compare packed signed dword integers in xmm2 and xmm3/m128/m32bcst and store packed maximum values in xmm1 using writemask k1.
// Compare packed signed dword integers in ymm2 and ymm3/m256/m32bcst and store packed maximum values in ymm1 using writemask k1.
// Compare packed signed dword integers in zmm2 and zmm3/m512/m32bcst and store packed maximum values in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXSD", reg, evex, rm)
}

// VPMAXSD_RVM
// Compare packed signed dword integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1.
// Compare packed signed dword integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMAXSD", reg, vex, rm)
}

// VPMAXSQ
// Compare packed signed qword integers in xmm2 and xmm3/m128/m64bcst and store packed maximum values in xmm1 using writemask k1.
// Compare packed signed qword integers in ymm2 and ymm3/m256/m64bcst and store packed maximum values in ymm1 using writemask k1.
// Compare packed signed qword integers in zmm2 and zmm3/m512/m64bcst and store packed maximum values in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXSQ", reg, evex, rm)
}

// VPMAXSW_FVM
// Compare packed signed word integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1 under writemask k1.
// Compare packed signed word integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1 under writemask k1.
// Compare packed signed word integers in zmm2 and zmm3/m512 and store packed maximum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXSW", reg, evex, rm)
}

// VPMAXSW_RVM
// Compare packed signed word integers in xmm3/m128 and xmm2 and store packed maximum values in xmm1.
// Compare packed signed word integers in ymm3/m256 and ymm2 and store packed maximum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=956
func VPMAXSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMAXSW", reg, vex, rm)
}

// VPMAXUB_FVM
// Compare packed unsigned byte integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1 under writemask k1.
// Compare packed unsigned byte integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1 under writemask k1.
// Compare packed unsigned byte integers in zmm2 and zmm3/m512 and store packed maximum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func VPMAXUB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXUB", reg, evex, rm)
}

// VPMAXUB_RVM
// Compare packed unsigned byte integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1.
// Compare packed unsigned byte integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func VPMAXUB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMAXUB", reg, vex, rm)
}

// VPMAXUW_FVM
// Compare packed unsigned word integers in xmm2 and xmm3/m128 and store packed maximum values in xmm1 under writemask k1.
// Compare packed unsigned word integers in ymm2 and ymm3/m256 and store packed maximum values in ymm1 under writemask k1.
// Compare packed unsigned word integers in zmm2 and zmm3/m512 and store packed maximum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func VPMAXUW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMAXUW", reg, evex, rm)
}

// VPMAXUW_RVM
// Compare packed unsigned word integers in xmm3/m128 and xmm2 and store maximum packed values in xmm1.
// Compare packed unsigned word integers in ymm3/m256 and ymm2 and store maximum packed values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=963
func VPMAXUW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMAXUW", reg, vex, rm)
}

// VPMINSB_FVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1 under writemask k1.
// Compare packed signed byte integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1 under writemask k1.
// Compare packed signed byte integers in zmm2 and zmm3/m512 and store packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func VPMINSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINSB", reg, evex, rm)
}

// VPMINSB_RVM
// Compare packed signed byte integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1.
// Compare packed signed byte integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func VPMINSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMINSB", reg, vex, rm)
}

// VPMINSW_FVM
// Compare packed signed word integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1 under writemask k1.
// Compare packed signed word integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1 under writemask k1.
// Compare packed signed word integers in zmm2 and zmm3/m512 and store packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func VPMINSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINSW", reg, evex, rm)
}

// VPMINSW_RVM
// Compare packed signed word integers in xmm3/m128 and xmm2 and return packed minimum values in xmm1.
// Compare packed signed word integers in ymm3/m256 and ymm2 and return packed minimum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=972
func VPMINSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMINSW", reg, vex, rm)
}

// VPMINUB_FVM
// Compare packed unsigned byte integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1 under writemask k1.
// Compare packed unsigned byte integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1 under writemask k1.
// Compare packed unsigned byte integers in zmm2 and zmm3/m512 and store packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func VPMINUB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINUB", reg, evex, rm)
}

// VPMINUB_RVM
// Compare packed unsigned byte integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1.
// Compare packed unsigned byte integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func VPMINUB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMINUB", reg, vex, rm)
}

// VPMINUD_FV
// Compare packed unsigned dword integers in xmm2 and xmm3/m128/m32bcst and store packed minimum values in xmm1 under writemask k1.
// Compare packed unsigned dword integers in ymm2 and ymm3/m256/m32bcst and store packed minimum values in ymm1 under writemask k1.
// Compare packed unsigned dword integers in zmm2 and zmm3/m512/m32bcst and store packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=986
func VPMINUD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINUD", reg, evex, rm)
}

// VPMINUD_RVM
// Compare packed unsigned dword integers in xmm2 and xmm3/m128 and store packed minimum values in xmm1.
// Compare packed unsigned dword integers in ymm2 and ymm3/m256 and store packed minimum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=986
func VPMINUD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMINUD", reg, vex, rm)
}

// VPMINUQ
// Compare packed unsigned qword integers in xmm2 and xmm3/m128/m64bcst and store packed minimum values in xmm1 under writemask k1.
// Compare packed unsigned qword integers in ymm2 and ymm3/m256/m64bcst and store packed minimum values in ymm1 under writemask k1.
// Compare packed unsigned qword integers in zmm2 and zmm3/m512/m64bcst and store packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=986
func VPMINUQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINUQ", reg, evex, rm)
}

// VPMINUW_FVM
// Compare packed unsigned word integers in xmm3/m128 and xmm2 and return packed minimum values in xmm1 under writemask k1.
// Compare packed unsigned word integers in ymm3/m256 and ymm2 and return packed minimum values in ymm1 under writemask k1.
// Compare packed unsigned word integers in zmm3/m512 and zmm2 and return packed minimum values in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func VPMINUW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMINUW", reg, evex, rm)
}

// VPMINUW_RVM
// Compare packed unsigned word integers in xmm3/m128 and xmm2 and return packed minimum values in xmm1.
// Compare packed unsigned word integers in ymm3/m256 and ymm2 and return packed minimum values in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=981
func VPMINUW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMINUW", reg, vex, rm)
}

// VPMOVB2M
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding byte in XMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding byte in YMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding byte in ZMM1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1759
func VPMOVB2M(reg, rm interface{}) {
	unsafe.Asm("VPMOVB2M", reg, rm)
}

// VPMOVD2M
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding doubleword in XMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding doubleword in YMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding doubleword in ZMM1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1759
func VPMOVD2M(reg, rm interface{}) {
	unsafe.Asm("VPMOVD2M", reg, rm)
}

// VPMOVDB
// Converts 16 packed double-word integers from zmm2 into 16 packed byte integers in xmm1/m128 with truncation under writemask k1.
// Converts 4 packed double-word integers from xmm2 into 4 packed byte integers in xmm1/m32 with truncation under writemask k1.
// Converts 8 packed double-word integers from ymm2 into 8 packed byte integers in xmm1/m64 with truncation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1774
func VPMOVDB(rm, reg interface{}) {
	unsafe.Asm("VPMOVDB", rm, reg)
}

// VPMOVDW
// Converts 8 packed double-word integers from ymm2 into 8 packed word integers in xmm1/m128 with truncation under writemask k1.
// Converts 4 packed double-word integers from xmm2 into 4 packed word integers in xmm1/m64 with truncation under writemask k1.
// Converts 16 packed double-word integers from zmm2 into 16 packed word integers in ymm1/m256 with truncation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1778
func VPMOVDW(rm, reg interface{}) {
	unsafe.Asm("VPMOVDW", rm, reg)
}

// VPMOVM2B
// Sets each byte in XMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each byte in YMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each byte in ZMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1756
func VPMOVM2B(reg, rm interface{}) {
	unsafe.Asm("VPMOVM2B", reg, rm)
}

// VPMOVM2D
// Sets each doubleword in XMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each doubleword in YMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each doubleword in ZMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1756
func VPMOVM2D(reg, rm interface{}) {
	unsafe.Asm("VPMOVM2D", reg, rm)
}

// VPMOVM2Q
// Sets each quadword in XMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each quadword in YMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each quadword in ZMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1756
func VPMOVM2Q(reg, rm interface{}) {
	unsafe.Asm("VPMOVM2Q", reg, rm)
}

// VPMOVM2W
// Sets each word in XMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each word in YMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
// Sets each word in ZMM1 to all 1’s or all 0’s based on the value of the corresponding bit in k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1756
func VPMOVM2W(reg, rm interface{}) {
	unsafe.Asm("VPMOVM2W", reg, rm)
}

// VPMOVMSKB
// Move a byte mask of xmm1 to reg. The upper bits of r32 or r64 are filled with zeros.
// Move a 32-bit mask of ymm1 to reg. The upper bits of r64 are filled with zeros.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=990
func VPMOVMSKB(reg, rm interface{}) {
	unsafe.Asm("VPMOVMSKB", reg, rm)
}

// VPMOVQ2M
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding quadword in XMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding quadword in YMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding quadword in ZMM1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1759
func VPMOVQ2M(reg, rm interface{}) {
	unsafe.Asm("VPMOVQ2M", reg, rm)
}

// VPMOVQB
// Converts 2 packed quad-word integers from xmm2 into 2 packed byte integers in xmm1/m16 with truncation under writemask k1.
// Converts 4 packed quad-word integers from ymm2 into 4 packed byte integers in xmm1/m32 with truncation under writemask k1.
// Converts 8 packed quad-word integers from zmm2 into 8 packed byte integers in xmm1/m64 with truncation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1762
func VPMOVQB(rm, reg interface{}) {
	unsafe.Asm("VPMOVQB", rm, reg)
}

// VPMOVSDB
// Converts 16 packed signed double-word integers from zmm2 into 16 packed signed byte integers in xmm1/m128 using signed saturation under writemask k1.
// Converts 4 packed signed double-word integers from xmm2 into 4 packed signed byte integers in xmm1/m32 using signed saturation under writemask k1.
// Converts 8 packed signed double-word integers from ymm2 into 8 packed signed byte integers in xmm1/m64 using signed saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1774
func VPMOVSDB(rm, reg interface{}) {
	unsafe.Asm("VPMOVSDB", rm, reg)
}

// VPMOVSDW
// Converts 8 packed signed double-word integers from ymm2 into 8 packed signed word integers in xmm1/m128 using signed saturation under writemask k1.
// Converts 4 packed signed double-word integers from xmm2 into 4 packed signed word integers in ymm1/m64 using signed saturation under writemask k1.
// Converts 16 packed signed double-word integers from zmm2 into 16 packed signed word integers in ymm1/m256 using signed saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1778
func VPMOVSDW(rm, reg interface{}) {
	unsafe.Asm("VPMOVSDW", rm, reg)
}

// VPMOVSQB
// Converts 2 packed signed quad-word integers from xmm2 into 2 packed signed byte integers in xmm1/m16 using signed saturation under writemask k1.
// Converts 4 packed signed quad-word integers from ymm2 into 4 packed signed byte integers in xmm1/m32 using signed saturation under writemask k1.
// Converts 8 packed signed quad-word integers from zmm2 into 8 packed signed byte integers in xmm1/m64 using signed saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1762
func VPMOVSQB(rm, reg interface{}) {
	unsafe.Asm("VPMOVSQB", rm, reg)
}

// VPMOVSXBD_QVM
// Sign extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1 subject to writemask k1.
// Sign extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 32-bit integers in ymm1 subject to writemask k1.
// Sign extend 16 packed 8-bit integers in the low 16 bytes of xmm2/m128 to 16 packed 32-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBD_QVM() {
	unsafe.Asm("VPMOVSXBD", nil)
}

// VPMOVSXBD_RM
// Sign extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1.
// Sign extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 32-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBD_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXBD", reg, rm)
}

// VPMOVSXBQ_OVM
// Sign extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1 subject to writemask k1.
// Sign extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 64-bit integers in ymm1 subject to writemask k1.
// Sign extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 64-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBQ_OVM() {
	unsafe.Asm("VPMOVSXBQ", nil)
}

// VPMOVSXBQ_RM
// Sign extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1.
// Sign extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 64-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXBQ", reg, rm)
}

// VPMOVSXBW_HVM
// Sign extend 8 packed 8-bit integers in xmm2/m64 to 8 packed 16-bit integers in zmm1.
// Sign extend 16 packed 8-bit integers in xmm2/m128 to 16 packed 16-bit integers in ymm1.
// Sign extend 32 packed 8-bit integers in ymm2/m256 to 32 packed 16-bit integers in zmm1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBW_HVM() {
	unsafe.Asm("VPMOVSXBW", nil)
}

// VPMOVSXBW_RM
// Sign extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 16-bit integers in xmm1.
// Sign extend 16 packed 8-bit integers in xmm2/m128 to 16 packed 16-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXBW_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXBW", reg, rm)
}

// VPMOVSXDQ_HVM
// Sign extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in zmm1 using writemask k1.
// Sign extend 4 packed 32-bit integers in the low 16 bytes of xmm2/m128 to 4 packed 64-bit integers in zmm1 using writemask k1.
// Sign extend 8 packed 32-bit integers in the low 32 bytes of ymm2/m256 to 8 packed 64-bit integers in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXDQ_HVM() {
	unsafe.Asm("VPMOVSXDQ", nil)
}

// VPMOVSXDQ_RM
// Sign extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in xmm1.
// Sign extend 4 packed 32-bit integers in the low 16 bytes of xmm2/m128 to 4 packed 64-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXDQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXDQ", reg, rm)
}

// VPMOVSXWD_HVM
// Sign extend 4 packed 16-bit integers in the low 8 bytes of ymm2/mem to 4 packed 32-bit integers in xmm1 subject to writemask k1.
// Sign extend 8 packed 16-bit integers in the low 16 bytes of ymm2/m128 to 8 packed 32-bit integers in ymm1 subject to writemask k1.
// Sign extend 16 packed 16-bit integers in the low 32 bytes of ymm2/m256 to 16 packed 32-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXWD_HVM() {
	unsafe.Asm("VPMOVSXWD", nil)
}

// VPMOVSXWD_RM
// Sign extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 32-bit integers in xmm1.
// Sign extend 8 packed 16-bit integers in the low 16 bytes of xmm2/m128 to 8 packed 32-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXWD_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXWD", reg, rm)
}

// VPMOVSXWQ_QVM
// Sign extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1 subject to writemask k1.
// Sign extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 64-bit integers in ymm1 subject to writemask k1.
// Sign extend 8 packed 16-bit integers in the low 16 bytes of xmm2/m128 to 8 packed 64-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXWQ_QVM() {
	unsafe.Asm("VPMOVSXWQ", nil)
}

// VPMOVSXWQ_RM
// Sign extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1.
// Sign extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 64-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=992
func VPMOVSXWQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVSXWQ", reg, rm)
}

// VPMOVUSDB
// Converts 16 packed unsigned double-word integers from zmm2 into 16 packed unsigned byte integers in xmm1/m128 using unsigned saturation under writemask k1.
// Converts 4 packed unsigned double-word integers from xmm2 into 4 packed unsigned byte integers in xmm1/m32 using unsigned saturation under writemask k1.
// Converts 8 packed unsigned double-word integers from ymm2 into 8 packed unsigned byte integers in xmm1/m64 using unsigned saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1774
func VPMOVUSDB(rm, reg interface{}) {
	unsafe.Asm("VPMOVUSDB", rm, reg)
}

// VPMOVUSDW
// Converts 8 packed unsigned double-word integers from ymm2 into 8 packed unsigned word integers in xmm1/m128 using unsigned saturation under writemask k1.
// Converts 4 packed unsigned double-word integers from xmm2 into 4 packed unsigned word integers in xmm1/m64 using unsigned saturation under writemask k1.
// Converts 16 packed unsigned double-word integers from zmm2 into 16 packed unsigned word integers in ymm1/m256 using unsigned saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1778
func VPMOVUSDW(rm, reg interface{}) {
	unsafe.Asm("VPMOVUSDW", rm, reg)
}

// VPMOVUSQB
// Converts 2 packed unsigned quad-word integers from xmm2 into 2 packed unsigned byte integers in xmm1/m16 using unsigned saturation under writemask k1.
// Converts 4 packed unsigned quad-word integers from ymm2 into 4 packed unsigned byte integers in xmm1/m32 using unsigned saturation under writemask k1.
// Converts 8 packed unsigned quad-word integers from zmm2 into 8 packed unsigned byte integers in xmm1/m64 using unsigned saturation under writemask k1.
//
// rm: ModRM:r/m (w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1762
func VPMOVUSQB(rm, reg interface{}) {
	unsafe.Asm("VPMOVUSQB", rm, reg)
}

// VPMOVW2M
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding word in XMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding word in YMM1.
// Sets each bit in k1 to 1 or 0 based on the value of the most significant bit of the corresponding word in ZMM1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1759
func VPMOVW2M(reg, rm interface{}) {
	unsafe.Asm("VPMOVW2M", reg, rm)
}

// VPMOVZXBD_QVM
// Zero extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1 subject to writemask k1.
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 32-bit integers in ymm1 subject to writemask k1.
// Zero extend 16 packed 8-bit integers in xmm2/m128 to 16 packed 32-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBD_QVM() {
	unsafe.Asm("VPMOVZXBD", nil)
}

// VPMOVZXBD_RM
// Zero extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 32-bit integers in xmm1.
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 32-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBD_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXBD", reg, rm)
}

// VPMOVZXBQ_OVM
// Zero extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1 subject to writemask k1.
// Zero extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 64-bit integers in ymm1 subject to writemask k1.
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 64-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBQ_OVM() {
	unsafe.Asm("VPMOVZXBQ", nil)
}

// VPMOVZXBQ_RM
// Zero extend 2 packed 8-bit integers in the low 2 bytes of xmm2/m16 to 2 packed 64-bit integers in xmm1.
// Zero extend 4 packed 8-bit integers in the low 4 bytes of xmm2/m32 to 4 packed 64-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXBQ", reg, rm)
}

// VPMOVZXBW_HVM
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 16-bit integers in xmm1.
// Zero extend 16 packed 8-bit integers in xmm2/m128 to 16 packed 16-bit integers in ymm1.
// Zero extend 32 packed 8-bit integers in ymm2/m256 to 32 packed 16-bit integers in zmm1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBW_HVM() {
	unsafe.Asm("VPMOVZXBW", nil)
}

// VPMOVZXBW_RM
// Zero extend 8 packed 8-bit integers in the low 8 bytes of xmm2/m64 to 8 packed 16-bit integers in xmm1.
// Zero extend 16 packed 8-bit integers in xmm2/m128 to 16 packed 16-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXBW_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXBW", reg, rm)
}

// VPMOVZXDQ_HVM
// Zero extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in zmm1 using writemask k1.
// Zero extend 4 packed 32-bit integers in xmm2/m128 to 4 packed 64-bit integers in zmm1 using writemask k1.
// Zero extend 8 packed 32-bit integers in ymm2/m256 to 8 packed 64-bit integers in zmm1 using writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXDQ_HVM() {
	unsafe.Asm("VPMOVZXDQ", nil)
}

// VPMOVZXDQ_RM
// Zero extend 2 packed 32-bit integers in the low 8 bytes of xmm2/m64 to 2 packed 64-bit integers in xmm1.
// Zero extend 4 packed 32-bit integers in xmm2/m128 to 4 packed 64-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXDQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXDQ", reg, rm)
}

// VPMOVZXWD_HVM
// Zero extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 32-bit integers in xmm1 subject to writemask k1.
// Zero extend 8 packed 16-bit integers in xmm2/m128 to 8 packed 32-bit integers in zmm1 subject to writemask k1.
// Zero extend 16 packed 16-bit integers in ymm2/m256 to 16 packed 32-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXWD_HVM() {
	unsafe.Asm("VPMOVZXWD", nil)
}

// VPMOVZXWD_RM
// Zero extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 32-bit integers in xmm1.
// Zero extend 8 packed 16-bit integers xmm2/m128 to 8 packed 32-bit integers in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXWD_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXWD", reg, rm)
}

// VPMOVZXWQ_QVM
// Zero extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1 subject to writemask k1.
// Zero extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 64-bit integers in ymm1 subject to writemask k1.
// Zero extend 8 packed 16-bit integers in xmm2/m128 to 8 packed 64-bit integers in zmm1 subject to writemask k1.
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXWQ_QVM() {
	unsafe.Asm("VPMOVZXWQ", nil)
}

// VPMOVZXWQ_RM
// Zero extend 2 packed 16-bit integers in the low 4 bytes of xmm2/m32 to 2 packed 64-bit integers in xmm1.
// Zero extend 4 packed 16-bit integers in the low 8 bytes of xmm2/m64 to 4 packed 64-bit integers in xmm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1002
func VPMOVZXWQ_RM(reg, rm interface{}) {
	unsafe.Asm("VPMOVZXWQ", reg, rm)
}

// VPMULDQ_FV
// Multiply packed signed doubleword integers in xmm2 by packed signed doubleword integers in xmm3/m128/m64bcst, and store the quadword results in xmm1 using writemask k1.
// Multiply packed signed doubleword integers in ymm2 by packed signed doubleword integers in ymm3/m256/m64bcst, and store the quadword results in ymm1 using writemask k1.
// Multiply packed signed doubleword integers in zmm2 by packed signed doubleword integers in zmm3/m512/m64bcst, and store the quadword results in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1011
func VPMULDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULDQ", reg, evex, rm)
}

// VPMULDQ_RVM
// Multiply packed signed doubleword integers in xmm2 by packed signed doubleword integers in xmm3/m128, and store the quadword results in xmm1.
// Multiply packed signed doubleword integers in ymm2 by packed signed doubleword integers in ymm3/m256, and store the quadword results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1011
func VPMULDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULDQ", reg, vex, rm)
}

// VPMULHRSW_FVM
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to xmm1 under writemask k1.
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to ymm1 under writemask k1.
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1014
func VPMULHRSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULHRSW", reg, evex, rm)
}

// VPMULHRSW_RVM
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to xmm1.
// Multiply 16-bit signed words, scale and round signed doublewords, pack high 16 bits to ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1014
func VPMULHRSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULHRSW", reg, vex, rm)
}

// VPMULHUW_FVM
// Multiply the packed unsigned word integers in xmm2 and xmm3/m128, and store the high 16 bits of the results in xmm1 under writemask k1.
// Multiply the packed unsigned word integers in ymm2 and ymm3/m256, and store the high 16 bits of the results in ymm1 under writemask k1.
// Multiply the packed unsigned word integers in zmm2 and zmm3/m512, and store the high 16 bits of the results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1018
func VPMULHUW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULHUW", reg, evex, rm)
}

// VPMULHUW_RVM
// Multiply the packed unsigned word integers in xmm2 and xmm3/m128, and store the high 16 bits of the results in xmm1.
// Multiply the packed unsigned word integers in ymm2 and ymm3/m256, and store the high 16 bits of the results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1018
func VPMULHUW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULHUW", reg, vex, rm)
}

// VPMULHW_FVM
// Multiply the packed signed word integers in xmm2 and xmm3/m128, and store the high 16 bits of the results in xmm1 under writemask k1.
// Multiply the packed signed word integers in ymm2 and ymm3/m256, and store the high 16 bits of the results in ymm1 under writemask k1.
// Multiply the packed signed word integers in zmm2 and zmm3/m512, and store the high 16 bits of the results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1022
func VPMULHW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULHW", reg, evex, rm)
}

// VPMULHW_RVM
// Multiply the packed signed word integers in xmm2 and xmm3/m128, and store the high 16 bits of the results in xmm1.
// Multiply the packed signed word integers in ymm2 and ymm3/m256, and store the high 16 bits of the results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1022
func VPMULHW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULHW", reg, vex, rm)
}

// VPMULLW_FVM
// Multiply the packed signed word integers in xmm2 and xmm3/m128, and store the low 16 bits of the results in xmm1 under writemask k1.
// Multiply the packed signed word integers in ymm2 and ymm3/m256, and store the low 16 bits of the results in ymm1 under writemask k1.
// Multiply the packed signed word integers in zmm2 and zmm3/m512, and store the low 16 bits of the results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1030
func VPMULLW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULLW", reg, evex, rm)
}

// VPMULLW_RVM
// Multiply the packed dword signed integers in xmm2 and xmm3/m128 and store the low 32 bits of each product in xmm1.
// Multiply the packed signed word integers in ymm2 and ymm3/m256, and store the low 16 bits of the results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1030
func VPMULLW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULLW", reg, vex, rm)
}

// VPMULUDQ_FV
// Multiply packed unsigned doubleword integers in xmm2 by packed unsigned doubleword integers in xmm3/m128/m64bcst, and store the quadword results in xmm1 under writemask k1.
// Multiply packed unsigned doubleword integers in ymm2 by packed unsigned doubleword integers in ymm3/m256/m64bcst, and store the quadword results in ymm1 under writemask k1.
// Multiply packed unsigned doubleword integers in zmm2 by packed unsigned doubleword integers in zmm3/m512/m64bcst, and store the quadword results in zmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1034
func VPMULUDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPMULUDQ", reg, evex, rm)
}

// VPMULUDQ_RVM
// Multiply packed unsigned doubleword integers in xmm2 by packed unsigned doubleword integers in xmm3/m128, and store the quadword results in xmm1.
// Multiply packed unsigned doubleword integers in ymm2 by packed unsigned doubleword integers in ymm3/m256, and store the quadword results in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1034
func VPMULUDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPMULUDQ", reg, vex, rm)
}

// VPOR
// Bitwise OR of xmm2/m128 and xmm3.
// Bitwise OR of ymm2/m256 and ymm3.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1051
func VPOR(reg, vex, rm interface{}) {
	unsafe.Asm("VPOR", reg, vex, rm)
}

// VPORD
// Bitwise OR of packed doubleword integers in xmm2 and xmm3/m128/m32bcst using writemask k1.
// Bitwise OR of packed doubleword integers in ymm2 and ymm3/m256/m32bcst using writemask k1.
// Bitwise OR of packed doubleword integers in zmm2 and zmm3/m512/m32bcst using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1051
func VPORD(reg, evex, rm interface{}) {
	unsafe.Asm("VPORD", reg, evex, rm)
}

// VPORQ
// Bitwise OR of packed quadword integers in xmm2 and xmm3/m128/m64bcst using writemask k1.
// Bitwise OR of packed quadword integers in ymm2 and ymm3/m256/m64bcst using writemask k1.
// Bitwise OR of packed quadword integers in zmm2 and zmm3/m512/m64bcst using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1051
func VPORQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPORQ", reg, evex, rm)
}

// VPSHUFB_FVM
// Shuffle bytes in xmm2 according to contents of xmm3/m128 under write mask k1.
// Shuffle bytes in ymm2 according to contents of ymm3/m256 under write mask k1.
// Shuffle bytes in zmm2 according to contents of zmm3/m512 under write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1064
func VPSHUFB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSHUFB", reg, evex, rm)
}

// VPSHUFB_RVM
// Shuffle bytes in xmm2 according to contents of xmm3/m128.
// Shuffle bytes in ymm2 according to contents of ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1064
func VPSHUFB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSHUFB", reg, vex, rm)
}

// VPSHUFD_FV
// Shuffle the doublewords in xmm2/m128/m32bcst based on the encoding in imm8 and store the result in xmm1 using writemask k1.
// Shuffle the doublewords in ymm2/m256/m32bcst based on the encoding in imm8 and store the result in ymm1 using writemask k1.
// Shuffle the doublewords in zmm2/m512/m32bcst based on the encoding in imm8 and store the result in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1068
func VPSHUFD_FV(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFD", reg, rm, imm)
}

// VPSHUFD_RMI
// Shuffle the doublewords in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
// Shuffle the doublewords in ymm2/m256 based on the encoding in imm8 and store the result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1068
func VPSHUFD_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFD", reg, rm, imm)
}

// VPSHUFHW_FVM
// Shuffle the high words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1 under write mask k1.
// Shuffle the high words in ymm2/m256 based on the encoding in imm8 and store the result in ymm1 under write mask k1.
// Shuffle the high words in zmm2/m512 based on the encoding in imm8 and store the result in zmm1 under write mask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1072
func VPSHUFHW_FVM(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFHW", reg, rm, imm)
}

// VPSHUFHW_RMI
// Shuffle the high words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
// Shuffle the high words in ymm2/m256 based on the encoding in imm8 and store the result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1072
func VPSHUFHW_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFHW", reg, rm, imm)
}

// VPSHUFLW_FVM
// Shuffle the low words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1 under write mask k1.
// Shuffle the low words in ymm2/m256 based on the encoding in imm8 and store the result in ymm1 under write mask k1.
// Shuffle the low words in zmm2/m512 based on the encoding in imm8 and store the result in zmm1 under write mask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1075
func VPSHUFLW_FVM(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFLW", reg, rm, imm)
}

// VPSHUFLW_RMI
// Shuffle the low words in xmm2/m128 based on the encoding in imm8 and store the result in xmm1.
// Shuffle the low words in ymm2/m256 based on the encoding in imm8 and store the result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1075
func VPSHUFLW_RMI(reg, rm, imm interface{}) {
	unsafe.Asm("VPSHUFLW", reg, rm, imm)
}

// VPSIGNB
// Negate/zero/preserve packed byte integers in xmm2 depending on the corresponding sign in xmm3/m128.
// Negate packed byte integers in ymm2 if the corresponding sign in ymm3/m256 is less than zero.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func VPSIGNB(reg, vex, rm interface{}) {
	unsafe.Asm("VPSIGNB", reg, vex, rm)
}

// VPSIGND
// Negate/zero/preserve packed doubleword integers in xmm2 depending on the corresponding sign in xmm3/m128.
// Negate packed doubleword integers in ymm2 if the corresponding sign in ymm3/m256 is less than zero.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func VPSIGND(reg, vex, rm interface{}) {
	unsafe.Asm("VPSIGND", reg, vex, rm)
}

// VPSIGNW
// Negate/zero/preserve packed word integers in xmm2 depending on the corresponding sign in xmm3/m128.
// Negate packed 16-bit integers in ymm2 if the corresponding sign in ymm3/m256 is less than zero.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1079
func VPSIGNW(reg, vex, rm interface{}) {
	unsafe.Asm("VPSIGNW", reg, vex, rm)
}

// VPSLLD_RVM
// Shift doublewords in xmm2 left by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSLLD", reg, vex, rm)
}

// VPSLLD_VMI
// Shift doublewords in xmm2 left by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLD_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSLLD", vex, rm, imm)
}

// VPSLLDQ_FVMI
// Shift xmm2/m128 left by imm8 bytes while shifting in 0s and store result in xmm1.
// Shift ymm2/m256 left by imm8 bytes while shifting in 0s and store result in ymm1.
// Shift zmm2/m512 left by imm8 bytes while shifting in 0s and store result in zmm1.
//
// evex: EVEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1083
func VPSLLDQ_FVMI(evex, rm, imm interface{}) {
	unsafe.Asm("VPSLLDQ", evex, rm, imm)
}

// VPSLLDQ_VMI
// Shift xmm2 left by imm8 bytes while shifting in 0s and store result in xmm1.
// Shift ymm2 left by imm8 bytes while shifting in 0s and store result in ymm1.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1083
func VPSLLDQ_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSLLDQ", vex, rm, imm)
}

// VPSLLQ_RVM
// Shift quadwords in xmm2 left by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSLLQ", reg, vex, rm)
}

// VPSLLQ_VMI
// Shift quadwords in xmm2 left by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLQ_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSLLQ", vex, rm, imm)
}

// VPSLLVD_FV
// Shift doublewords in xmm2 left by amount specified in the corresponding element of xmm3/m128/m32bcst while shifting in 0s using writemask k1.
// Shift doublewords in ymm2 left by amount specified in the corresponding element of ymm3/m256/m32bcst while shifting in 0s using writemask k1.
// Shift doublewords in zmm2 left by amount specified in the corresponding element of zmm3/m512/m32bcst while shifting in 0s using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1801
func VPSLLVD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPSLLVD", reg, evex, rm)
}

// VPSLLVD_RVM
// Shift doublewords in xmm2 left by amount specified in the corresponding element of xmm3/m128 while shifting in 0s.
// Shift doublewords in ymm2 left by amount specified in the corresponding element of ymm3/m256 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1801
func VPSLLVD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSLLVD", reg, vex, rm)
}

// VPSLLVQ_FV
// Shift quadwords in xmm2 left by amount specified in the corresponding element of xmm3/m128/m64bcst while shifting in 0s using writemask k1.
// Shift quadwords in ymm2 left by amount specified in the corresponding element of ymm3/m256/m64bcst while shifting in 0s using writemask k1.
// Shift quadwords in zmm2 left by amount specified in the corresponding element of zmm3/m512/m64bcst while shifting in 0s using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1801
func VPSLLVQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPSLLVQ", reg, evex, rm)
}

// VPSLLVQ_RVM
// Shift quadwords in xmm2 left by amount specified in the corresponding element of xmm3/m128 while shifting in 0s.
// Shift quadwords in ymm2 left by amount specified in the corresponding element of ymm3/m256 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1801
func VPSLLVQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSLLVQ", reg, vex, rm)
}

// VPSLLVW
// Shift words in xmm2 left by amount specified in the corresponding element of xmm3/m128 while shifting in 0s using writemask k1.
// Shift words in ymm2 left by amount specified in the corresponding element of ymm3/m256 while shifting in 0s using writemask k1.
// Shift words in zmm2 left by amount specified in the corresponding element of zmm3/m512 while shifting in 0s using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1801
func VPSLLVW(reg, evex, rm interface{}) {
	unsafe.Asm("VPSLLVW", reg, evex, rm)
}

// VPSLLW_RVM
// Shift words in xmm2 left by amount specified in xmm3/m128 while shifting in 0s.
// Shift words in ymm2 left by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSLLW", reg, vex, rm)
}

// VPSLLW_VMI
// Shift words in xmm2 left by imm8 while shifting in 0s.
// Shift words in ymm2 left by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1085
func VPSLLW_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSLLW", vex, rm, imm)
}

// VPSRAD_RVM
// Shift doublewords in xmm2 right by amount specified in xmm3/m128 while shifting in sign bits.
// Shift doublewords in ymm2 right by amount specified in xmm3/m128 while shifting in sign bits.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1097
func VPSRAD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRAD", reg, vex, rm)
}

// VPSRAD_VMI
// Shift doublewords in xmm2 right by imm8 while shifting in sign bits.
// Shift doublewords in ymm2 right by imm8 while shifting in sign bits.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1097
func VPSRAD_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRAD", vex, rm, imm)
}

// VPSRAVD_FV
// Shift doublewords in xmm2 right by amount specified in the corresponding element of xmm3/m128/m32bcst while shifting in sign bits using writemask k1.
// Shift doublewords in ymm2 right by amount specified in the corresponding element of ymm3/m256/m32bcst while shifting in sign bits using writemask k1.
// Shift doublewords in zmm2 right by amount specified in the corresponding element of zmm3/m512/m32bcst while shifting in sign bits using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1806
func VPSRAVD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPSRAVD", reg, evex, rm)
}

// VPSRAVD_RVM
// Shift doublewords in xmm2 right by amount specified in the corresponding element of xmm3/m128 while shifting in sign bits.
// Shift doublewords in ymm2 right by amount specified in the corresponding element of ymm3/m256 while shifting in sign bits.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1806
func VPSRAVD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRAVD", reg, vex, rm)
}

// VPSRAVQ
// Shift quadwords in xmm2 right by amount specified in the corresponding element of xmm3/m128/m64bcst while shifting in sign bits using writemask k1.
// Shift quadwords in ymm2 right by amount specified in the corresponding element of ymm3/m256/m64bcst while shifting in sign bits using writemask k1.
// Shift quadwords in zmm2 right by amount specified in the corresponding element of zmm3/m512/m64bcst while shifting in sign bits using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1806
func VPSRAVQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPSRAVQ", reg, evex, rm)
}

// VPSRAVW
// Shift words in xmm2 right by amount specified in the corresponding element of xmm3/m128 while shifting in sign bits using writemask k1.
// Shift words in ymm2 right by amount specified in the corresponding element of ymm3/m256 while shifting in sign bits using writemask k1.
// Shift words in zmm2 right by amount specified in the corresponding element of zmm3/m512 while shifting in sign bits using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1806
func VPSRAVW(reg, evex, rm interface{}) {
	unsafe.Asm("VPSRAVW", reg, evex, rm)
}

// VPSRAW_M128
// Shift words in xmm2 right by amount specified in xmm3/m128 while shifting in sign bits using writemask k1.
// Shift words in ymm2 right by amount specified in xmm3/m128 while shifting in sign bits using writemask k1.
// Shift words in zmm2 right by amount specified in xmm3/m128 while shifting in sign bits using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1097
func VPSRAW_M128(reg, evex, rm interface{}) {
	unsafe.Asm("VPSRAW", reg, evex, rm)
}

// VPSRAW_RVM
// Shift words in xmm2 right by amount specified in xmm3/m128 while shifting in sign bits.
// Shift words in ymm2 right by amount specified in xmm3/m128 while shifting in sign bits.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1097
func VPSRAW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRAW", reg, vex, rm)
}

// VPSRAW_VMI
// Shift words in xmm2 right by imm8 while shifting in sign bits.
// Shift words in ymm2 right by imm8 while shifting in sign bits.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1097
func VPSRAW_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRAW", vex, rm, imm)
}

// VPSRLD_RVM
// Shift doublewords in xmm2 right by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRLD", reg, vex, rm)
}

// VPSRLD_VMI
// Shift doublewords in xmm2 right by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLD_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRLD", vex, rm, imm)
}

// VPSRLDQ_FVM
// Shift xmm2/m128 right by imm8 bytes while shifting in 0s and store result in xmm1.
// Shift ymm2/m256 right by imm8 bytes while shifting in 0s and store result in ymm1.
// Shift zmm2/m512 right by imm8 bytes while shifting in 0s and store result in zmm1.
//
// evex: EVEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1107
func VPSRLDQ_FVM(evex, rm, imm interface{}) {
	unsafe.Asm("VPSRLDQ", evex, rm, imm)
}

// VPSRLDQ_VMI
// Shift xmm2 right by imm8 bytes while shifting in 0s.
// Shift ymm1 right by imm8 bytes while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1107
func VPSRLDQ_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRLDQ", vex, rm, imm)
}

// VPSRLQ_RVM
// Shift quadwords in xmm2 right by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRLQ", reg, vex, rm)
}

// VPSRLQ_VMI
// Shift quadwords in xmm2 right by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLQ_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRLQ", vex, rm, imm)
}

// VPSRLW_RVM
// Shift words in xmm2 right by amount specified in xmm3/m128 while shifting in 0s.
// Shift words in ymm2 right by amount specified in xmm3/m128 while shifting in 0s.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSRLW", reg, vex, rm)
}

// VPSRLW_VMI
// Shift words in xmm2 right by imm8 while shifting in 0s.
// Shift words in ymm2 right by imm8 while shifting in 0s.
//
// vex: VEX.vvvv (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1109
func VPSRLW_VMI(vex, rm, imm interface{}) {
	unsafe.Asm("VPSRLW", vex, rm, imm)
}

// VPSUBB_FVM
// Subtract packed byte integers in xmm3/m128 from xmm2 and store in xmm1 using writemask k1.
// Subtract packed byte integers in ymm3/m256 from ymm2 and store in ymm1 using writemask k1.
// Subtract packed byte integers in zmm3/m512 from zmm2 and store in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func VPSUBB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBB", reg, evex, rm)
}

// VPSUBB_RVM
// Subtract packed byte integers in xmm3/m128 from xmm2.
// Subtract packed byte integers in ymm3/m256 from ymm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func VPSUBB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBB", reg, vex, rm)
}

// VPSUBD
// Subtract packed doubleword integers in xmm3/m128 from xmm2.
// Subtract packed doubleword integers in ymm3/m256 from ymm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func VPSUBD(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBD", reg, vex, rm)
}

// VPSUBQ_FV
// Subtract packed quadword integers in xmm3/m128/m64bcst from xmm2 and store in xmm1 using writemask k1.
// Subtract packed quadword integers in ymm3/m256/m64bcst from ymm2 and store in ymm1 using writemask k1.
// Subtract packed quadword integers in zmm3/m512/m64bcst from zmm2 and store in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1128
func VPSUBQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBQ", reg, evex, rm)
}

// VPSUBQ_RVM
// Subtract packed quadword integers in xmm3/m128 from xmm2.
// Subtract packed quadword integers in ymm3/m256 from ymm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1128
func VPSUBQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBQ", reg, vex, rm)
}

// VPSUBSB_FVM
// Subtract packed signed byte integers in xmm3/m128 from packed signed byte integers in xmm2 and saturate results and store in xmm1 using writemask k1.
// Subtract packed signed byte integers in ymm3/m256 from packed signed byte integers in ymm2 and saturate results and store in ymm1 using writemask k1.
// Subtract packed signed byte integers in zmm3/m512 from packed signed byte integers in zmm2 and saturate results and store in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func VPSUBSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBSB", reg, evex, rm)
}

// VPSUBSB_RVM
// Subtract packed signed byte integers in xmm3/m128 from packed signed byte integers in xmm2 and saturate results.
// Subtract packed signed byte integers in ymm3/m256 from packed signed byte integers in ymm2 and saturate results.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func VPSUBSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBSB", reg, vex, rm)
}

// VPSUBSW_FVM
// Subtract packed signed word integers in xmm3/m128 from packed signed word integers in xmm2 and saturate results and store in xmm1 using writemask k1.
// Subtract packed signed word integers in ymm3/m256 from packed signed word integers in ymm2 and saturate results and store in ymm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func VPSUBSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBSW", reg, evex, rm)
}

// VPSUBSW_RVM
// Subtract packed signed word integers in xmm3/m128 from packed signed word integers in xmm2 and saturate results.
// Subtract packed signed word integers in ymm3/m256 from packed signed word integers in ymm2 and saturate results.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1131
func VPSUBSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBSW", reg, vex, rm)
}

// VPSUBUSB_FVM
// Subtract packed unsigned byte integers in xmm3/m128 from packed unsigned byte integers in xmm2, saturate results and store in xmm1 using writemask k1.
// Subtract packed unsigned byte integers in ymm3/m256 from packed unsigned byte integers in ymm2, saturate results and store in ymm1 using writemask k1.
// Subtract packed unsigned byte integers in zmm3/m512 from packed unsigned byte integers in zmm2, saturate results and store in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func VPSUBUSB_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBUSB", reg, evex, rm)
}

// VPSUBUSB_RVM
// Subtract packed unsigned byte integers in xmm3/m128 from packed unsigned byte integers in xmm2 and saturate result.
// Subtract packed unsigned byte integers in ymm3/m256 from packed unsigned byte integers in ymm2 and saturate result.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func VPSUBUSB_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBUSB", reg, vex, rm)
}

// VPSUBUSW_FVM
// Subtract packed unsigned word integers in xmm3/m128 from packed unsigned word integers in xmm2 and saturate results and store in xmm1 using writemask k1.
// Subtract packed unsigned word integers in ymm3/m256 from packed unsigned word integers in ymm2, saturate results and store in ymm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func VPSUBUSW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBUSW", reg, evex, rm)
}

// VPSUBUSW_RVM
// Subtract packed unsigned word integers in xmm3/m128 from packed unsigned word integers in xmm2 and saturate result.
// Subtract packed unsigned word integers in ymm3/m256 from packed unsigned word integers in ymm2 and saturate result.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1135
func VPSUBUSW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBUSW", reg, vex, rm)
}

// VPSUBW_FVM
// Subtract packed word integers in xmm3/m128 from xmm2 and store in xmm1 using writemask k1.
// Subtract packed word integers in ymm3/m256 from ymm2 and store in ymm1 using writemask k1.
// Subtract packed word integers in zmm3/m512 from zmm2 and store in zmm1 using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func VPSUBW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPSUBW", reg, evex, rm)
}

// VPSUBW_RVM
// Subtract packed word integers in xmm3/m128 from xmm2.
// Subtract packed word integers in ymm3/m256 from ymm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1121
func VPSUBW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPSUBW", reg, vex, rm)
}

// VPTEST
// Set ZF and CF depending on bitwise AND and ANDN of sources.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1139
func VPTEST(reg, rm interface{}) {
	unsafe.Asm("VPTEST", reg, rm)
}

// VPUNPCKHBW_FVM
// Interleave high-order bytes from xmm2 and xmm3/m128 into xmm1 register using k1 write mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHBW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKHBW", reg, evex, rm)
}

// VPUNPCKHBW_RVM
// Interleave high-order bytes from xmm2 and xmm3/m128 into xmm1.
// Interleave high-order bytes from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHBW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKHBW", reg, vex, rm)
}

// VPUNPCKHDQ_FV
// Interleave high-order doublewords from xmm2 and xmm3/m128/m32bcst into xmm1 register using k1 write mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKHDQ", reg, evex, rm)
}

// VPUNPCKHDQ_RVM
// Interleave high-order doublewords from xmm2 and xmm3/m128 into xmm1.
// Interleave high-order doublewords from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKHDQ", reg, vex, rm)
}

// VPUNPCKHQDQ_FV
// Interleave high-order quadword from xmm2 and xmm3/m128/m64bcst into xmm1 register using k1 write mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHQDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKHQDQ", reg, evex, rm)
}

// VPUNPCKHQDQ_RVM
// Interleave high-order quadword from xmm2 and xmm3/m128 into xmm1 register.
// Interleave high-order quadword from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHQDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKHQDQ", reg, vex, rm)
}

// VPUNPCKHWD_FVM
// Interleave high-order words from xmm2 and xmm3/m128 into xmm1 register using k1 write mask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHWD_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKHWD", reg, evex, rm)
}

// VPUNPCKHWD_RVM
// Interleave high-order words from xmm2 and xmm3/m128 into xmm1.
// Interleave high-order words from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1143
func VPUNPCKHWD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKHWD", reg, vex, rm)
}

// VPUNPCKLBW_FVM
// Interleave low-order bytes from xmm2 and xmm3/m128 into xmm1 register subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLBW_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKLBW", reg, evex, rm)
}

// VPUNPCKLBW_RVM
// Interleave low-order bytes from xmm2 and xmm3/m128 into xmm1.
// Interleave low-order bytes from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLBW_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKLBW", reg, vex, rm)
}

// VPUNPCKLDQ_FV
// Interleave low-order doublewords from xmm2 and xmm3/m128/m32bcst into xmm1 register subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKLDQ", reg, evex, rm)
}

// VPUNPCKLDQ_RVM
// Interleave low-order doublewords from xmm2 and xmm3/m128 into xmm1.
// Interleave low-order doublewords from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKLDQ", reg, vex, rm)
}

// VPUNPCKLQDQ_FV
// Interleave low-order quadword from zmm2 and zmm3/m512/m64bcst into zmm1 register subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLQDQ_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKLQDQ", reg, evex, rm)
}

// VPUNPCKLQDQ_RVM
// Interleave low-order quadword from xmm2 and xmm3/m128 into xmm1 register.
// Interleave low-order quadword from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLQDQ_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKLQDQ", reg, vex, rm)
}

// VPUNPCKLWD_FVM
// Interleave low-order words from xmm2 and xmm3/m128 into xmm1 register subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLWD_FVM(reg, evex, rm interface{}) {
	unsafe.Asm("VPUNPCKLWD", reg, evex, rm)
}

// VPUNPCKLWD_RVM
// Interleave low-order words from xmm2 and xmm3/m128 into xmm1.
// Interleave low-order words from ymm2 and ymm3/m256 into ymm1 register.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1153
func VPUNPCKLWD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VPUNPCKLWD", reg, vex, rm)
}

// VPXOR
// Bitwise XOR of xmm3/m128 and xmm2.
// Bitwise XOR of ymm3/m256 and ymm2.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1170
func VPXOR(reg, vex, rm interface{}) {
	unsafe.Asm("VPXOR", reg, vex, rm)
}

// VPXORD
// Bitwise XOR of packed doubleword integers in xmm2 and xmm3/m128 using writemask k1.
// Bitwise XOR of packed doubleword integers in ymm2 and ymm3/m256 using writemask k1.
// Bitwise XOR of packed doubleword integers in zmm2 and zmm3/m512/m32bcst using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1170
func VPXORD(reg, evex, rm interface{}) {
	unsafe.Asm("VPXORD", reg, evex, rm)
}

// VPXORQ
// Bitwise XOR of packed quadword integers in xmm2 and xmm3/m128 using writemask k1.
// Bitwise XOR of packed quadword integers in ymm2 and ymm3/m256 using writemask k1.
// Bitwise XOR of packed quadword integers in zmm2 and zmm3/m512/m64bcst using writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1170
func VPXORQ(reg, evex, rm interface{}) {
	unsafe.Asm("VPXORQ", reg, evex, rm)
}

// VRANGEPD
// Calculate two RANGE operation output value from 2 pairs of double-precision floating-point values in xmm2 and xmm3/m128/m32bcst, store the results to xmm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
// Calculate four RANGE operation output value from 4pairs of double-precision floating-point values in ymm2 and ymm3/m256/m32bcst, store the results to ymm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
// Calculate eight RANGE operation output value from 8 pairs of double-precision floating-point values in zmm2 and zmm3/m512/m32bcst, store the results to zmm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1826
func VRANGEPD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRANGEPD", reg, evex, rm, imm)
}

// VRANGEPS
// Calculate four RANGE operation output value from 4 pairs of single-precision floating-point values in xmm2 and xmm3/m128/m32bcst, store the results to xmm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
// Calculate eight RANGE operation output value from 8 pairs of single-precision floating-point values in ymm2 and ymm3/m256/m32bcst, store the results to ymm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
// Calculate 16 RANGE operation output value from 16 pairs of single-precision floating-point values in zmm2 and zmm3/m512/m32bcst, store the results to zmm1 under the writemask k1. Imm8 specifies the comparison and sign of the range operation.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1831
func VRANGEPS(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRANGEPS", reg, evex, rm, imm)
}

// VRANGESD
// Calculate a RANGE operation output value from 2 double- precision floating-point values in xmm2 and xmm3/m64, store the output to xmm1 under writemask. Imm8 specifies the comparison and sign of the range operation.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1835
func VRANGESD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRANGESD", reg, evex, rm, imm)
}

// VRANGESS
// Calculate a RANGE operation output value from 2 single- precision floating-point values in xmm2 and xmm3/m32, store the output to xmm1 under writemask. Imm8 specifies the comparison and sign of the range operation.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1838
func VRANGESS(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRANGESS", reg, evex, rm, imm)
}

// VRCP14PD
// Computes the approximate reciprocals of the packed double- precision floating-point values in xmm2/m128/m64bcst and stores the results in xmm1. Under writemask.
// Computes the approximate reciprocals of the packed double- precision floating-point values in ymm2/m256/m64bcst and stores the results in ymm1. Under writemask.
// Computes the approximate reciprocals of the packed double- precision floating-point values in zmm2/m512/m64bcst and stores the results in zmm1. Under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1841
func VRCP14PD(reg, rm interface{}) {
	unsafe.Asm("VRCP14PD", reg, rm)
}

// VRCP14PS
// Computes the approximate reciprocals of the packed single- precision floating-point values in xmm2/m128/m32bcst and stores the results in xmm1. Under writemask.
// Computes the approximate reciprocals of the packed single- precision floating-point values in ymm2/m256/m32bcst and stores the results in ymm1. Under writemask.
// Computes the approximate reciprocals of the packed single- precision floating-point values in zmm2/m512/m32bcst and stores the results in zmm1. Under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1845
func VRCP14PS(reg, rm interface{}) {
	unsafe.Asm("VRCP14PS", reg, rm)
}

// VRCP14SS
// Computes the approximate reciprocal of the scalar single- precision floating-point value in xmm3/m32 and stores the results in xmm1 using writemask k1. Also, upper double- precision floating-point value (bits[127:32]) from xmm2 is copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1847
func VRCP14SS(reg, evex, rm interface{}) {
	unsafe.Asm("VRCP14SS", reg, evex, rm)
}

// VRCP28PD
// Computes the approximate reciprocals ( < 2^-28 relative error) of the packed double-precision floating-point values in zmm2/m512/m64bcst and stores the results in zmm1. Under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1849
func VRCP28PD(reg, rm interface{}) {
	unsafe.Asm("VRCP28PD", reg, rm)
}

// VRCP28PS
// Computes the approximate reciprocals ( < 2^-28 relative error) of the packed single-precision floating-point values in zmm2/m512/m32bcst and stores the results in zmm1. Under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1853
func VRCP28PS(reg, rm interface{}) {
	unsafe.Asm("VRCP28PS", reg, rm)
}

// VRCP28SD
// Computes the approximate reciprocal ( < 2^-28 relative error) of the scalar double-precision floating-point value in xmm3/m64 and stores the results in xmm1. Under writemask. Also, upper double-precision floating-point value (bits[127:64]) from xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1851
func VRCP28SD(reg, evex, rm interface{}) {
	unsafe.Asm("VRCP28SD", reg, evex, rm)
}

// VRCP28SS
// Computes the approximate reciprocal ( < 2^-28 relative error) of the scalar single-precision floating-point value in xmm3/m32 and stores the results in xmm1. Under writemask. Also, upper 3 single-precision floating-point values (bits[127:32]) from xmm2 is copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1855
func VRCP28SS(reg, evex, rm interface{}) {
	unsafe.Asm("VRCP28SS", reg, evex, rm)
}

// VRCPPS
// Computes the approximate reciprocals of packed single-precision values in xmm2/mem and stores the results in xmm1.
// Computes the approximate reciprocals of packed single-precision values in ymm2/mem and stores the results in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1178
func VRCPPS(reg, rm interface{}) {
	unsafe.Asm("VRCPPS", reg, rm)
}

// VRCPSS
// Computes the approximate reciprocal of the scalar single-precision floating-point value in xmm3/m32 and stores the result in xmm1. Also, upper single precision floating-point values (bits[127:32]) from xmm2 are copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1180
func VRCPSS(reg, vex, rm interface{}) {
	unsafe.Asm("VRCPSS", reg, vex, rm)
}

// VREDUCEPS
// Perform reduction transformation on packed single-precision floating point values in xmm2/m128/m32bcst by subtracting a number of fraction bits specified by the imm8 field. Stores the result in xmm1 register under writemask k1.
// Perform reduction transformation on packed single-precision floating point values in ymm2/m256/m32bcst by subtracting a number of fraction bits specified by the imm8 field. Stores the result in ymm1 register under writemask k1.
// Perform reduction transformation on packed single-precision floating point values in zmm2/m512/m32bcst by subtracting a number of fraction bits specified by the imm8 field. Stores the result in zmm1 register under writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1862
func VREDUCEPS(reg, rm, imm interface{}) {
	unsafe.Asm("VREDUCEPS", reg, rm, imm)
}

// VREDUCESD
// Perform a reduction transformation on a scalar double-precision floating point value in xmm3/m64 by subtracting a number of fraction bits specified by the imm8 field. Also, upper double precision floating-point value (bits[127:64]) from xmm2 are copied to xmm1[127:64]. Stores the result in xmm1 register.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1860
func VREDUCESD(reg, evex, rm interface{}) {
	unsafe.Asm("VREDUCESD", reg, evex, rm)
}

// VRNDSCALEPD
// Rounds packed double-precision floating point values in xmm2/m128/m64bcst to a number of fraction bits specified by the imm8 field. Stores the result in xmm1 register. Under writemask.
// Rounds packed double-precision floating point values in ymm2/m256/m64bcst to a number of fraction bits specified by the imm8 field. Stores the result in ymm1 register. Under writemask.
// Rounds packed double-precision floating-point values in zmm2/m512/m64bcst to a number of fraction bits specified by the imm8 field. Stores the result in zmm1 register using writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1866
func VRNDSCALEPD(reg, rm, imm interface{}) {
	unsafe.Asm("VRNDSCALEPD", reg, rm, imm)
}

// VRNDSCALEPS
// Rounds packed single-precision floating point values in xmm2/m128/m32bcst to a number of fraction bits specified by the imm8 field. Stores the result in xmm1 register. Under writemask.
// Rounds packed single-precision floating point values in ymm2/m256/m32bcst to a number of fraction bits specified by the imm8 field. Stores the result in ymm1 register. Under writemask.
// Rounds packed single-precision floating-point values in zmm2/m512/m32bcst to a number of fraction bits specified by the imm8 field. Stores the result in zmm1 register using writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1872
func VRNDSCALEPS(reg, rm, imm interface{}) {
	unsafe.Asm("VRNDSCALEPS", reg, rm, imm)
}

// VRNDSCALESD
// Rounds scalar double-precision floating-point value in xmm3/m64 to a number of fraction bits specified by the imm8 field. Stores the result in xmm1 register.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1870
func VRNDSCALESD(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRNDSCALESD", reg, evex, rm, imm)
}

// VRNDSCALESS
// Rounds scalar single-precision floating-point value in xmm3/m32 to a number of fraction bits specified by the imm8 field. Stores the result in xmm1 register under writemask.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1875
func VRNDSCALESS(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VRNDSCALESS", reg, evex, rm, imm)
}

// VROUNDPD
// Round packed double-precision floating-point values in xmm2/m128 and place the result in xmm1. The rounding mode is determined by imm8.
// Round packed double-precision floating-point values in ymm2/m256 and place the result in ymm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1216
func VROUNDPD(reg, rm, imm interface{}) {
	unsafe.Asm("VROUNDPD", reg, rm, imm)
}

// VROUNDPS
// Round packed single-precision floating-point values in xmm2/m128 and place the result in xmm1. The rounding mode is determined by imm8.
// Round packed single-precision floating-point values in ymm2/m256 and place the result in ymm1. The rounding mode is determined by imm8.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1219
func VROUNDPS(reg, rm, imm interface{}) {
	unsafe.Asm("VROUNDPS", reg, rm, imm)
}

// VROUNDSD
// Round the low packed double precision floating-point value in xmm3/m64 and place the result in xmm1. The rounding mode is determined by imm8. Upper packed double precision floating-point value (bits[127:64]) from xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1222
func VROUNDSD(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VROUNDSD", reg, vex, rm, imm)
}

// VROUNDSS
// Round the low packed single precision floating-point value in xmm3/m32 and place the result in xmm1. The rounding mode is determined by imm8. Also, upper packed single precision floating-point values (bits[127:32]) from xmm2 are copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1224
func VROUNDSS(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VROUNDSS", reg, vex, rm, imm)
}

// VRSQRT14PD
// Computes the approximate reciprocal square roots of the packed double-precision floating-point values in xmm2/m128/m64bcst and stores the results in xmm1. Under writemask.
// Computes the approximate reciprocal square roots of the packed double-precision floating-point values in ymm2/m256/m64bcst and stores the results in ymm1. Under writemask.
// Computes the approximate reciprocal square roots of the packed double-precision floating-point values in zmm2/m512/m64bcst and stores the results in zmm1 under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1877
func VRSQRT14PD(reg, rm interface{}) {
	unsafe.Asm("VRSQRT14PD", reg, rm)
}

// VRSQRT14PS
// Computes the approximate reciprocal square roots of the packed single-precision floating-point values in xmm2/m128/m32bcst and stores the results in xmm1. Under writemask.
// Computes the approximate reciprocal square roots of the packed single-precision floating-point values in ymm2/m256/m32bcst and stores the results in ymm1. Under writemask.
// Computes the approximate reciprocal square roots of the packed single-precision floating-point values in zmm2/m512/m32bcst and stores the results in zmm1. Under writemask.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1881
func VRSQRT14PS(reg, rm interface{}) {
	unsafe.Asm("VRSQRT14PS", reg, rm)
}

// VRSQRT14SD
// Computes the approximate reciprocal square root of the scalar double-precision floating-point value in xmm3/m64 and stores the result in the low quadword element of xmm1 using writemask k1. Bits[127:64] of xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1879
func VRSQRT14SD(reg, evex, rm interface{}) {
	unsafe.Asm("VRSQRT14SD", reg, evex, rm)
}

// VRSQRT14SS
// Computes the approximate reciprocal square root of the scalar single-precision floating-point value in xmm3/m32 and stores the result in the low doubleword element of xmm1 using writemask k1. Bits[127:32] of xmm2 is copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1883
func VRSQRT14SS(reg, vex, rm interface{}) {
	unsafe.Asm("VRSQRT14SS", reg, vex, rm)
}

// VRSQRT28PD
// Computes approximations to the Reciprocal square root (<2^- 28 relative error) of the packed double-precision floating-point values from zmm2/m512/m64bcst and stores result in zmm1with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1885
func VRSQRT28PD(reg, rm interface{}) {
	unsafe.Asm("VRSQRT28PD", reg, rm)
}

// VRSQRT28PS
// Computes approximations to the Reciprocal square root (<2^-28 relative error) of the packed single-precision floating-point values from zmm2/m512/m32bcst and stores result in zmm1with writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1889
func VRSQRT28PS(reg, rm interface{}) {
	unsafe.Asm("VRSQRT28PS", reg, rm)
}

// VRSQRT28SD
// Computes approximate reciprocal square root (<2^-28 relative error) of the scalar double-precision floating-point value from xmm3/m64 and stores result in xmm1with writemask k1. Also, upper double-precision floating-point value (bits[127:64]) from xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1887
func VRSQRT28SD(reg, evex, rm interface{}) {
	unsafe.Asm("VRSQRT28SD", reg, evex, rm)
}

// VRSQRT28SS
// Computes approximate reciprocal square root (<2^-28 relative error) of the scalar single-precision floating-point value from xmm3/m32 and stores result in xmm1with writemask k1. Also, upper 3 single-precision floating-point value (bits[127:32]) from xmm2 is copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1891
func VRSQRT28SS(reg, evex, rm interface{}) {
	unsafe.Asm("VRSQRT28SS", reg, evex, rm)
}

// VRSQRTPS
// Computes the approximate reciprocals of the square roots of packed single-precision values in xmm2/mem and stores the results in xmm1.
// Computes the approximate reciprocals of the square roots of packed single-precision values in ymm2/mem and stores the results in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1228
func VRSQRTPS(reg, rm interface{}) {
	unsafe.Asm("VRSQRTPS", reg, rm)
}

// VRSQRTSS
// Computes the approximate reciprocal of the square root of the low single precision floating-point value in xmm3/m32 and stores the results in xmm1. Also, upper single precision floating-point values (bits[127:32]) from xmm2 are copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1230
func VRSQRTSS(reg, vex, rm interface{}) {
	unsafe.Asm("VRSQRTSS", reg, vex, rm)
}

// VSCALEFSD
// Scale the scalar double-precision floating-point values in xmm2 using the value from xmm3/m64. Under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1896
func VSCALEFSD(reg, evex, rm interface{}) {
	unsafe.Asm("VSCALEFSD", reg, evex, rm)
}

// VSCATTERPF0DPD
// Using signed dword indices, prefetch sparse byte memory locations containing double-precision data using writemask k1 and T0 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1907
func VSCATTERPF0DPD(vsib interface{}) {
	unsafe.Asm("VSCATTERPF0DPD", vsib)
}

// VSCATTERPF0DPS
// Using signed dword indices, prefetch sparse byte memory locations containing single-precision data using writemask k1 and T0 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1907
func VSCATTERPF0DPS(vsib interface{}) {
	unsafe.Asm("VSCATTERPF0DPS", vsib)
}

// VSCATTERPF0QPD
// Using signed qword indices, prefetch sparse byte memory locations containing double-precision data using writemask k1 and T0 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1907
func VSCATTERPF0QPD(vsib interface{}) {
	unsafe.Asm("VSCATTERPF0QPD", vsib)
}

// VSCATTERPF0QPS
// Using signed qword indices, prefetch sparse byte memory locations containing single-precision data using writemask k1 and T0 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1907
func VSCATTERPF0QPS(vsib interface{}) {
	unsafe.Asm("VSCATTERPF0QPS", vsib)
}

// VSCATTERPF1DPD
// Using signed dword indices, prefetch sparse byte memory locations containing double-precision data using writemask k1 and T1 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1909
func VSCATTERPF1DPD(vsib interface{}) {
	unsafe.Asm("VSCATTERPF1DPD", vsib)
}

// VSCATTERPF1DPS
// Using signed dword indices, prefetch sparse byte memory locations containing single-precision data using writemask k1 and T1 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1909
func VSCATTERPF1DPS(vsib interface{}) {
	unsafe.Asm("VSCATTERPF1DPS", vsib)
}

// VSCATTERPF1QPD
// Using signed qword indices, prefetch sparse byte memory locations containing double-precision data using writemask k1 and T1 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1909
func VSCATTERPF1QPD(vsib interface{}) {
	unsafe.Asm("VSCATTERPF1QPD", vsib)
}

// VSCATTERPF1QPS
// Using signed qword indices, prefetch sparse byte memory locations containing single-precision data using writemask k1 and T1 hint with intent to write.
//
// vsib: vsib (r)
//
// Documentation: https://golang.org/s/x86manual#page=1909
func VSCATTERPF1QPS(vsib interface{}) {
	unsafe.Asm("VSCATTERPF1QPS", vsib)
}

// VSHUFF32X4
// Shuffle 128-bit packed single-precision floating-point values selected by imm8 from ymm2 and ymm3/m256/m32bcst and place results in ymm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFF32X4(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFF32X4", reg, evex, rm, imm)
}

// VSHUFF32x4
// Shuffle 128-bit packed single-precision floating-point values selected by imm8 from zmm2 and zmm3/m512/m32bcst and place results in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFF32x4(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFF32x4", reg, evex, rm, imm)
}

// VSHUFF64X2
// Shuffle 128-bit packed double-precision floating-point values selected by imm8 from ymm2 and ymm3/m256/m64bcst and place results in ymm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFF64X2(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFF64X2", reg, evex, rm, imm)
}

// VSHUFF64x2
// Shuffle 128-bit packed double-precision floating-point values selected by imm8 from zmm2 and zmm3/m512/m64bcst and place results in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFF64x2(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFF64x2", reg, evex, rm, imm)
}

// VSHUFI32X4
// Shuffle 128-bit packed double-word values selected by imm8 from ymm2 and ymm3/m256/m32bcst and place results in ymm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFI32X4(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFI32X4", reg, evex, rm, imm)
}

// VSHUFI32x4
// Shuffle 128-bit packed double-word values selected by imm8 from zmm2 and zmm3/m512/m32bcst and place results in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFI32x4(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFI32x4", reg, evex, rm, imm)
}

// VSHUFI64X2
// Shuffle 128-bit packed quad-word values selected by imm8 from ymm2 and ymm3/m256/m64bcst and place results in ymm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFI64X2(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFI64X2", reg, evex, rm, imm)
}

// VSHUFI64x2
// Shuffle 128-bit packed quad-word values selected by imm8 from zmm2 and zmm3/m512/m64bcst and place results in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1911
func VSHUFI64x2(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFI64x2", reg, evex, rm, imm)
}

// VSHUFPD_FV
// Shuffle two paris of double-precision floating-point values from xmm2 and xmm3/m128/m64bcst using imm8 to select from each pair. store interleaved results in xmm1 subject to writemask k1.
// Shuffle four paris of double-precision floating-point values from ymm2 and ymm3/m256/m64bcst using imm8 to select from each pair. store interleaved results in ymm1 subject to writemask k1.
// Shuffle eight paris of double-precision floating-point values from zmm2 and zmm3/m512/m64bcst using imm8 to select from each pair. store interleaved results in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1269
func VSHUFPD_FV(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFPD", reg, evex, rm, imm)
}

// VSHUFPD_RVMI
// Shuffle two pairs of double-precision floating-point values from xmm2 and xmm3/m128 using imm8 to select from each pair, interleaved result is stored in xmm1.
// Shuffle four pairs of double-precision floating-point values from ymm2 and ymm3/m256 using imm8 to select from each pair, interleaved result is stored in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1269
func VSHUFPD_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VSHUFPD", reg, vex, rm, imm)
}

// VSHUFPS_FV
// Select from quadruplet of single-precision floating- point values in xmm1 and xmm2/m128 using imm8, interleaved result pairs are stored in xmm1, subject to writemask k1.
// Select from quadruplet of single-precision floating- point values in ymm2 and ymm3/m256 using imm8, interleaved result pairs are stored in ymm1, subject to writemask k1.
// Select from quadruplet of single-precision floating- point values in zmm2 and zmm3/m512 using imm8, interleaved result pairs are stored in zmm1, subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1274
func VSHUFPS_FV(reg, evex, rm, imm interface{}) {
	unsafe.Asm("VSHUFPS", reg, evex, rm, imm)
}

// VSHUFPS_RVMI
// Select from quadruplet of single-precision floating- point values in xmm1 and xmm2/m128 using imm8, interleaved result pairs are stored in xmm1.
// Select from quadruplet of single-precision floating- point values in ymm2 and ymm3/m256 using imm8, interleaved result pairs are stored in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1274
func VSHUFPS_RVMI(reg, vex, rm, imm interface{}) {
	unsafe.Asm("VSHUFPS", reg, vex, rm, imm)
}

// VSQRTPD_FV
// Computes Square Roots of the packed double-precision floating-point values in xmm2/m128/m64bcst and stores the result in xmm1 subject to writemask k1.
// Computes Square Roots of the packed double-precision floating-point values in ymm2/m256/m64bcst and stores the result in ymm1 subject to writemask k1.
// Computes Square Roots of the packed double-precision floating-point values in zmm2/m512/m64bcst and stores the result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1284
func VSQRTPD_FV(reg, rm interface{}) {
	unsafe.Asm("VSQRTPD", reg, rm)
}

// VSQRTPD_RM
// Computes Square Roots of the packed double-precision floating-point values in xmm2/m128 and stores the result in xmm1.
// Computes Square Roots of the packed double-precision floating-point values in ymm2/m256 and stores the result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1284
func VSQRTPD_RM(reg, rm interface{}) {
	unsafe.Asm("VSQRTPD", reg, rm)
}

// VSQRTPS_FV
// Computes Square Roots of the packed single-precision floating-point values in xmm2/m128/m32bcst and stores the result in xmm1 subject to writemask k1.
// Computes Square Roots of the packed single-precision floating-point values in ymm2/m256/m32bcst and stores the result in ymm1 subject to writemask k1.
// Computes Square Roots of the packed single-precision floating-point values in zmm2/m512/m32bcst and stores the result in zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1287
func VSQRTPS_FV(reg, rm interface{}) {
	unsafe.Asm("VSQRTPS", reg, rm)
}

// VSQRTPS_RM
// Computes Square Roots of the packed single-precision floating-point values in xmm2/m128 and stores the result in xmm1.
// Computes Square Roots of the packed single-precision floating-point values in ymm2/m256 and stores the result in ymm1.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1287
func VSQRTPS_RM(reg, rm interface{}) {
	unsafe.Asm("VSQRTPS", reg, rm)
}

// VSQRTSD_RVM
// Computes square root of the low double-precision floating- point value in xmm3/m64 and stores the results in xmm1. Also, upper double-precision floating-point value (bits[127:64]) from xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1290
func VSQRTSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSQRTSD", reg, vex, rm)
}

// VSQRTSD_T1S
// Computes square root of the low double-precision floating- point value in xmm3/m64 and stores the results in xmm1 under writemask k1. Also, upper double-precision floating- point value (bits[127:64]) from xmm2 is copied to xmm1[127:64].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1290
func VSQRTSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VSQRTSD", reg, evex, rm)
}

// VSQRTSS_RVM
// Computes square root of the low single-precision floating-point value in xmm3/m32 and stores the results in xmm1. Also, upper single-precision floating-point values (bits[127:32]) from xmm2 are copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1292
func VSQRTSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSQRTSS", reg, vex, rm)
}

// VSQRTSS_T1S
// Computes square root of the low single-precision floating-point value in xmm3/m32 and stores the results in xmm1 under writemask k1. Also, upper single-precision floating-point values (bits[127:32]) from xmm2 are copied to xmm1[127:32].
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1292
func VSQRTSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VSQRTSS", reg, evex, rm)
}

// VSTMXCSR
// Store contents of MXCSR register to m32.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1299
func VSTMXCSR(rm interface{}) {
	unsafe.Asm("VSTMXCSR", rm)
}

// VSUBPD_FV
// Subtract packed double-precision floating-point values from xmm3/m128/m64bcst to xmm2 and store result in xmm1 with writemask k1.
// Subtract packed double-precision floating-point values from ymm3/m256/m64bcst to ymm2 and store result in ymm1 with writemask k1.
// Subtract packed double-precision floating-point values from zmm3/m512/m64bcst to zmm2 and store result in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1308
func VSUBPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VSUBPD", reg, evex, rm)
}

// VSUBPD_RVM
// Subtract packed double-precision floating-point values in xmm3/mem from xmm2 and store result in xmm1.
// Subtract packed double-precision floating-point values in ymm3/mem from ymm2 and store result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1308
func VSUBPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSUBPD", reg, vex, rm)
}

// VSUBPS_FV
// Subtract packed single-precision floating-point values from xmm3/m128/m32bcst to xmm2 and stores result in xmm1 with writemask k1.
// Subtract packed single-precision floating-point values from ymm3/m256/m32bcst to ymm2 and stores result in ymm1 with writemask k1.
// Subtract packed single-precision floating-point values in zmm3/m512/m32bcst from zmm2 and stores result in zmm1 with writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1311
func VSUBPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VSUBPS", reg, evex, rm)
}

// VSUBPS_RVM
// Subtract packed single-precision floating-point values in xmm3/mem from xmm2 and stores result in xmm1.
// Subtract packed single-precision floating-point values in ymm3/mem from ymm2 and stores result in ymm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1311
func VSUBPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSUBPS", reg, vex, rm)
}

// VSUBSD_RVM
// Subtract the low double-precision floating-point value in xmm3/m64 from xmm2 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1314
func VSUBSD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSUBSD", reg, vex, rm)
}

// VSUBSD_T1S
// Subtract the low double-precision floating-point value in xmm3/m64 from xmm2 and store the result in xmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1314
func VSUBSD_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VSUBSD", reg, evex, rm)
}

// VSUBSS_RVM
// Subtract the low single-precision floating-point value in xmm3/m32 from xmm2 and store the result in xmm1.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1316
func VSUBSS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VSUBSS", reg, vex, rm)
}

// VSUBSS_T1S
// Subtract the low single-precision floating-point value in xmm3/m32 from xmm2 and store the result in xmm1 under writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1316
func VSUBSS_T1S(reg, evex, rm interface{}) {
	unsafe.Asm("VSUBSS", reg, evex, rm)
}

// VTESTPD
// Set ZF and CF depending on sign bit AND and ANDN of packed double-precision floating-point sources.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1916
func VTESTPD(reg, rm interface{}) {
	unsafe.Asm("VTESTPD", reg, rm)
}

// VTESTPS
// Set ZF and CF depending on sign bit AND and ANDN of packed single-precision floating-point sources.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1916
func VTESTPS(reg, rm interface{}) {
	unsafe.Asm("VTESTPS", reg, rm)
}

// VUCOMISD_RM
// Compare low double-precision floating-point values in xmm1 and xmm2/mem64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1335
func VUCOMISD_RM(reg, rm interface{}) {
	unsafe.Asm("VUCOMISD", reg, rm)
}

// VUCOMISD_T1S
// Compare low double-precision floating-point values in xmm1 and xmm2/m64 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1335
func VUCOMISD_T1S(reg, rm interface{}) {
	unsafe.Asm("VUCOMISD", reg, rm)
}

// VUCOMISS_RM
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1337
func VUCOMISS_RM(reg, rm interface{}) {
	unsafe.Asm("VUCOMISS", reg, rm)
}

// VUCOMISS_T1S
// Compare low single-precision floating-point values in xmm1 and xmm2/mem32 and set the EFLAGS flags accordingly.
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1337
func VUCOMISS_T1S(reg, rm interface{}) {
	unsafe.Asm("VUCOMISS", reg, rm)
}

// VUNPCKHPD_FV
// Unpacks and Interleaves double precision floating-point values from high quadwords of xmm2 and xmm3/m128/m64bcst subject to writemask k1.
// Unpacks and Interleaves double precision floating-point values from high quadwords of ymm2 and ymm3/m256/m64bcst subject to writemask k1.
// Unpacks and Interleaves double-precision floating-point values from high quadwords of zmm2 and zmm3/m512/m64bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1340
func VUNPCKHPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VUNPCKHPD", reg, evex, rm)
}

// VUNPCKHPD_RVM
// Unpacks and Interleaves double-precision floating-point values from high quadwords of ymm2 and ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1340
func VUNPCKHPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VUNPCKHPD", reg, vex, rm)
}

// VUNPCKHPD__RVM
// Unpacks and Interleaves double-precision floating-point values from high quadwords of xmm2 and xmm3/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1340
func VUNPCKHPD__RVM() {
	unsafe.Asm("VUNPCKHPD", nil)
}

// VUNPCKHPS_FV
// Unpacks and Interleaves single-precision floating-point values from high quadwords of xmm2 and xmm3/m128/m32bcst and write result to xmm1 subject to writemask k1.
// Unpacks and Interleaves single-precision floating-point values from high quadwords of ymm2 and ymm3/m256/m32bcst and write result to ymm1 subject to writemask k1.
// Unpacks and Interleaves single-precision floating-point values from high quadwords of zmm2 and zmm3/m512/m32bcst and write result to zmm1 subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1344
func VUNPCKHPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VUNPCKHPS", reg, evex, rm)
}

// VUNPCKHPS_RVM
// Unpacks and Interleaves single-precision floating-point values from high quadwords of ymm2 and ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1344
func VUNPCKHPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VUNPCKHPS", reg, vex, rm)
}

// VUNPCKHPS__RVM
// Unpacks and Interleaves single-precision floating-point values from high quadwords of xmm2 and xmm3/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1344
func VUNPCKHPS__RVM() {
	unsafe.Asm("VUNPCKHPS", nil)
}

// VUNPCKLPD_FV
// Unpacks and Interleaves double precision floating-point values from low quadwords of xmm2 and xmm3/m128/m64bcst subject to write mask k1.
// Unpacks and Interleaves double precision floating-point values from low quadwords of ymm2 and ymm3/m256/m64bcst subject to write mask k1.
// Unpacks and Interleaves double-precision floating-point values from low quadwords of zmm2 and zmm3/m512/m64bcst subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1348
func VUNPCKLPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VUNPCKLPD", reg, evex, rm)
}

// VUNPCKLPD_RVM
// Unpacks and Interleaves double-precision floating-point values from low quadwords of ymm2 and ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1348
func VUNPCKLPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VUNPCKLPD", reg, vex, rm)
}

// VUNPCKLPD__RVM
// Unpacks and Interleaves double-precision floating-point values from low quadwords of xmm2 and xmm3/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1348
func VUNPCKLPD__RVM() {
	unsafe.Asm("VUNPCKLPD", nil)
}

// VUNPCKLPS_FV
// Unpacks and Interleaves single-precision floating-point values from low quadwords of xmm2 and xmm3/mem and write result to xmm1 subject to write mask k1.
// Unpacks and Interleaves single-precision floating-point values from low quadwords of ymm2 and ymm3/mem and write result to ymm1 subject to write mask k1.
// Unpacks and Interleaves single-precision floating-point values from low quadwords of zmm2 and zmm3/m512/m32bcst and write result to zmm1 subject to write mask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1352
func VUNPCKLPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VUNPCKLPS", reg, evex, rm)
}

// VUNPCKLPS_RVM
// Unpacks and Interleaves single-precision floating-point values from low quadwords of ymm2 and ymm3/m256.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1352
func VUNPCKLPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VUNPCKLPS", reg, vex, rm)
}

// VUNPCKLPS__RVM
// Unpacks and Interleaves single-precision floating-point values from low quadwords of xmm2 and xmm3/m128.
//
// Documentation: https://golang.org/s/x86manual#page=1352
func VUNPCKLPS__RVM() {
	unsafe.Asm("VUNPCKLPS", nil)
}

// VXORPD_FV
// Return the bitwise logical XOR of packed double- precision floating-point values in xmm2 and xmm3/m128/m64bcst subject to writemask k1.
// Return the bitwise logical XOR of packed double- precision floating-point values in ymm2 and ymm3/m256/m64bcst subject to writemask k1.
// Return the bitwise logical XOR of packed double- precision floating-point values in zmm2 and zmm3/m512/m64bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1952
func VXORPD_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VXORPD", reg, evex, rm)
}

// VXORPD_RVM
// Return the bitwise logical XOR of packed double- precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv (r)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1952
func VXORPD_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VXORPD", reg, vex, rm)
}

// VXORPD__RVM
// Return the bitwise logical XOR of packed double- precision floating-point values in xmm2 and xmm3/mem.
//
// Documentation: https://golang.org/s/x86manual#page=1952
func VXORPD__RVM() {
	unsafe.Asm("VXORPD", nil)
}

// VXORPS_FV
// Return the bitwise logical XOR of packed single- precision floating-point values in xmm2 and xmm3/m128/m32bcst subject to writemask k1.
// Return the bitwise logical XOR of packed single- precision floating-point values in ymm2 and ymm3/m256/m32bcst subject to writemask k1.
// Return the bitwise logical XOR of packed single- precision floating-point values in zmm2 and zmm3/m512/m32bcst subject to writemask k1.
//
// reg: ModRM:reg (w)
// evex: EVEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1955
func VXORPS_FV(reg, evex, rm interface{}) {
	unsafe.Asm("VXORPS", reg, evex, rm)
}

// VXORPS_RVM
// Return the bitwise logical XOR of packed single- precision floating-point values in xmm2 and xmm3/mem.
// Return the bitwise logical XOR of packed single- precision floating-point values in ymm2 and ymm3/mem.
//
// reg: ModRM:reg (w)
// vex: VEX.vvvv
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1955
func VXORPS_RVM(reg, vex, rm interface{}) {
	unsafe.Asm("VXORPS", reg, vex, rm)
}

// VZEROALL
// Zero all YMM registers.
//
// Documentation: https://golang.org/s/x86manual#page=1919
func VZEROALL() {
	unsafe.Asm("VZEROALL", nil)
}

// VZEROUPPER
// Zero upper 128 bits of all YMM registers.
//
// Documentation: https://golang.org/s/x86manual#page=1921
func VZEROUPPER() {
	unsafe.Asm("VZEROUPPER", nil)
}

// WAIT
// Check pending unmasked floating-point exceptions.
//
// Documentation: https://golang.org/s/x86manual#page=1923
func WAIT() {
	unsafe.Asm("WAIT", nil)
}

// WBINVD
// Write back and flush Internal caches; initiate writing-back and flushing of external caches.
//
// Documentation: https://golang.org/s/x86manual#page=1924
func WBINVD() {
	unsafe.Asm("WBINVD", nil)
}

// WRFSBASE
// Load the FS base address with the 32-bit value in the source register.
// Load the FS base address with the 64-bit value in the source register.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1926
func WRFSBASE(rm interface{}) {
	unsafe.Asm("WRFSBASE", rm)
}

// WRGSBASE
// Load the GS base address with the 32-bit value in the source register.
// Load the GS base address with the 64-bit value in the source register.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1926
func WRGSBASE(rm interface{}) {
	unsafe.Asm("WRGSBASE", rm)
}

// WRMSR
// Write the value in EDX:EAX to MSR specified by ECX.
//
// Documentation: https://golang.org/s/x86manual#page=1928
func WRMSR() {
	unsafe.Asm("WRMSR", nil)
}

// WRPKRU
// Writes EAX into PKRU.
//
// Documentation: https://golang.org/s/x86manual#page=1930
func WRPKRU() {
	unsafe.Asm("WRPKRU", nil)
}

// XABORT
// Causes an RTM abort if in RTM execution
//
// imm: imm8
//
// Documentation: https://golang.org/s/x86manual#page=1935
func XABORT(imm interface{}) {
	unsafe.Asm("XABORT", imm)
}

// XACQUIRE
// A hint used with an “XACQUIRE-enabled“ instruction to start lock elision on the instruction memory operand address.
//
// Documentation: https://golang.org/s/x86manual#page=1931
func XACQUIRE() {
	unsafe.Asm("XACQUIRE", nil)
}

// XADD
// Exchange r16 and r/m16; load sum into r/m16.
// Exchange r32 and r/m32; load sum into r/m32.
// Exchange r64 and r/m64; load sum into r/m64.
// Exchange r8 and r/m8; load sum into r/m8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1937
func XADD(rm, reg interface{}) {
	unsafe.Asm("XADD", rm, reg)
}

// XBEGIN
// Specifies the start of an RTM region. Provides a 16-bit relative offset to compute the address of the fallback instruction address at which execution resumes following an RTM abort.
// Specifies the start of an RTM region. Provides a 32-bit relative offset to compute the address of the fallback instruction address at which execution resumes following an RTM abort.
//
// offset: Offset
//
// Documentation: https://golang.org/s/x86manual#page=1939
func XBEGIN(offset interface{}) {
	unsafe.Asm("XBEGIN", offset)
}

// XCHG_MR
// Exchange r16 with word from r/m16.
// Exchange r32 with doubleword from r/m32.
// Exchange r64 with quadword from r/m64.
// Exchange r8 (byte register) with byte from r/m8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1942
func XCHG_MR(rm, reg interface{}) {
	unsafe.Asm("XCHG", rm, reg)
}

// XCHG_O
// Exchange r16 with AX.
// Exchange r32 with EAX.
// Exchange r64 with RAX.
// Exchange AX with r16.
// Exchange EAX with r32.
// Exchange RAX with r64.
//
// ax: AX (r, w)
// opcode: opcode + rd (r, w)
//
// Documentation: https://golang.org/s/x86manual#page=1942
func XCHG_O(ax, opcode interface{}) {
	unsafe.Asm("XCHG", ax, opcode)
}

// XCHG_RM
// Exchange word from r/m16 with r16.
// Exchange doubleword from r/m32 with r32.
// Exchange quadword from r/m64 with r64.
// Exchange byte from r/m8 with r8 (byte register).
//
// reg: ModRM:reg (w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1942
func XCHG_RM(reg, rm interface{}) {
	unsafe.Asm("XCHG", reg, rm)
}

// XEND
// Specifies the end of an RTM code region.
//
// Documentation: https://golang.org/s/x86manual#page=1944
func XEND() {
	unsafe.Asm("XEND", nil)
}

// XGETBV
// Reads an XCR specified by ECX into EDX:EAX.
//
// Documentation: https://golang.org/s/x86manual#page=1946
func XGETBV() {
	unsafe.Asm("XGETBV", nil)
}

// XLATB
// Set AL to memory byte DS:[(E)BX + unsigned AL].
// Set AL to memory byte [RBX + unsigned AL].
//
// Documentation: https://golang.org/s/x86manual#page=1948
func XLATB() {
	unsafe.Asm("XLATB", nil)
}

// XOR_I
// AL XOR imm8.
// AX XOR imm16.
// EAX XOR imm32.
// RAX XOR imm32 (sign-extended).
//
// al: AL/AX/EAX/RAX
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1950
func XOR_I(al, imm interface{}) {
	unsafe.Asm("XOR", al, imm)
}

// XOR_MI
// r/m16 XOR imm16.
// r/m16 XOR imm8 (sign-extended).
// r/m32 XOR imm32.
// r/m32 XOR imm8 (sign-extended).
// r/m64 XOR imm32 (sign-extended).
// r/m64 XOR imm8 (sign-extended).
// r/m8 XOR imm8.
//
// rm: ModRM:r/m (r, w)
// imm: imm8/16/32
//
// Documentation: https://golang.org/s/x86manual#page=1950
func XOR_MI(rm, imm interface{}) {
	unsafe.Asm("XOR", rm, imm)
}

// XOR_MR
// r/m16 XOR r16.
// r/m32 XOR r32.
// r/m64 XOR r64.
// r/m8 XOR r8.
//
// rm: ModRM:r/m (r, w)
// reg: ModRM:reg (r)
//
// Documentation: https://golang.org/s/x86manual#page=1950
func XOR_MR(rm, reg interface{}) {
	unsafe.Asm("XOR", rm, reg)
}

// XOR_RM
// r16 XOR r/m16.
// r32 XOR r/m32.
// r64 XOR r/m64.
// r8 XOR r/m8.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1950
func XOR_RM(reg, rm interface{}) {
	unsafe.Asm("XOR", reg, rm)
}

// XORPD
// Return the bitwise logical XOR of packed double- precision floating-point values in xmm1 and xmm2/mem.
//
// Documentation: https://golang.org/s/x86manual#page=1952
func XORPD() {
	unsafe.Asm("XORPD", nil)
}

// XORPS
// Return the bitwise logical XOR of packed single- precision floating-point values in xmm1 and xmm2/mem.
//
// reg: ModRM:reg (r, w)
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1955
func XORPS(reg, rm interface{}) {
	unsafe.Asm("XORPS", reg, rm)
}

// XRELEASE
// A hint used with an “XRELEASE-enabled“ instruction to end lock elision on the instruction memory operand address.
//
// Documentation: https://golang.org/s/x86manual#page=1931
func XRELEASE() {
	unsafe.Asm("XRELEASE", nil)
}

// XRSTOR
// Restore state components specified by EDX:EAX from mem.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1958
func XRSTOR(rm interface{}) {
	unsafe.Asm("XRSTOR", rm)
}

// XRSTOR64
// Restore state components specified by EDX:EAX from mem.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1958
func XRSTOR64(rm interface{}) {
	unsafe.Asm("XRSTOR64", rm)
}

// XRSTORS
// Restore state components specified by EDX:EAX from mem.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1962
func XRSTORS(rm interface{}) {
	unsafe.Asm("XRSTORS", rm)
}

// XRSTORS64
// Restore state components specified by EDX:EAX from mem.
//
// rm: ModRM:r/m (r)
//
// Documentation: https://golang.org/s/x86manual#page=1962
func XRSTORS64(rm interface{}) {
	unsafe.Asm("XRSTORS64", rm)
}

// XSAVE
// Save state components specified by EDX:EAX to mem.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1966
func XSAVE(rm interface{}) {
	unsafe.Asm("XSAVE", rm)
}

// XSAVE64
// Save state components specified by EDX:EAX to mem.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1966
func XSAVE64(rm interface{}) {
	unsafe.Asm("XSAVE64", rm)
}

// XSAVEC
// Save state components specified by EDX:EAX to mem with compaction.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1969
func XSAVEC(rm interface{}) {
	unsafe.Asm("XSAVEC", rm)
}

// XSAVEC64
// Save state components specified by EDX:EAX to mem with compaction.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1969
func XSAVEC64(rm interface{}) {
	unsafe.Asm("XSAVEC64", rm)
}

// XSAVEOPT
// Save state components specified by EDX:EAX to mem, optimizing if possible.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1972
func XSAVEOPT(rm interface{}) {
	unsafe.Asm("XSAVEOPT", rm)
}

// XSAVEOPT64
// Save state components specified by EDX:EAX to mem, optimizing if possible.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1972
func XSAVEOPT64(rm interface{}) {
	unsafe.Asm("XSAVEOPT64", rm)
}

// XSAVES
// Save state components specified by EDX:EAX to mem with compaction, optimizing if possible.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1975
func XSAVES(rm interface{}) {
	unsafe.Asm("XSAVES", rm)
}

// XSAVES64
// Save state components specified by EDX:EAX to mem with compaction, optimizing if possible.
//
// rm: ModRM:r/m (w)
//
// Documentation: https://golang.org/s/x86manual#page=1975
func XSAVES64(rm interface{}) {
	unsafe.Asm("XSAVES64", rm)
}

// XSETBV
// Write the value in EDX:EAX to the XCR specified by ECX.
//
// Documentation: https://golang.org/s/x86manual#page=1978
func XSETBV() {
	unsafe.Asm("XSETBV", nil)
}

// XTEST
// Test if executing in a transactional region
//
// Documentation: https://golang.org/s/x86manual#page=1980
func XTEST() {
	unsafe.Asm("XTEST", nil)
}
