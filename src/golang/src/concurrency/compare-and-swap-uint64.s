TEXT ·CompareAndSwapUint64(SB),NOSPLIT,$0-21  // HL
        MOVL	addr+0(FP), BP
        TESTL	$7, BP
        JZ	2(PC)
        MOVL	0, AX // crash with nil ptr deref
        MOVL	old_lo+4(FP), AX
        MOVL	old_hi+8(FP), DX
        MOVL	new_lo+12(FP), BX
        MOVL	new_hi+16(FP), CX
        // CMPXCHG8B was introduced on the Pentium.
        LOCK
        CMPXCHG8B	0(BP)    // HL
        SETEQ	swapped+20(FP)
        RET
