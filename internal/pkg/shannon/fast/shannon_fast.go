/*
	Package main - transpiled by c2go version: v0.23.3 Berkelium 2018-05-04

	If you have found any issues, please raise an issue at:
	https://github.com/elliotchance/c2go/
*/

// Warning (FieldDecl):  /usr/include/mach/i386/_structs.h:94 : Error : name of FieldDecl is empty

package fast

import "unsafe"
import "github.com/elliotchance/c2go/darwin"

type int8_t int8
type int16_t int16
type int32_t int32
type int64_t int64
type uint8_t uint8
type uint16_t uint16
type uint32_t uint32
type uint64_t uint64
type int_least8_t int8_t
type int_least16_t int16_t
type int_least32_t int32_t
type int_least64_t int64_t
type uint_least8_t uint8_t
type uint_least16_t uint16_t
type uint_least32_t uint32_t
type uint_least64_t uint64_t
type int_fast8_t int8_t
type int_fast16_t int16_t
type int_fast32_t int32_t
type int_fast64_t int64_t
type uint_fast8_t uint8_t
type uint_fast16_t uint16_t
type uint_fast32_t uint32_t
type uint_fast64_t uint64_t
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int32
type __uint32_t uint32
type __int64_t int64
type __uint64_t uint64
type __darwin_intptr_t int32
type __darwin_natural_t uint32
type __darwin_ct_rune_t darwin.CtRuneT
type __mbstate_t struct{ memory unsafe.Pointer }

func (unionVar *__mbstate_t) copy() __mbstate_t {
	var buffer [128]byte
	for i := range buffer {
		buffer[i] = (*((*[128]byte)(unionVar.memory)))[i]
	}
	var newUnion __mbstate_t
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__mbstate_t) __mbstate8() *[128]byte {
	if unionVar.memory == nil {
		var buffer [128]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*[128]byte)(unionVar.memory)
}
func (unionVar *__mbstate_t) _mbstateL() *int64 {
	if unionVar.memory == nil {
		var buffer [128]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int64)(unionVar.memory)
}

type __darwin_mbstate_t int64
type __darwin_ptrdiff_t int32
type __darwin_size_t uint32
type __darwin_va_list int64
type __darwin_wchar_t int32
type __darwin_rune_t __darwin_wchar_t
type __darwin_wint_t int32
type __darwin_clock_t uint32
type __darwin_socklen_t __uint32_t
type __darwin_ssize_t int32
type __darwin_time_t int32
type __darwin_blkcnt_t __int64_t
type __darwin_blksize_t __int32_t
type __darwin_dev_t __int32_t
type __darwin_fsblkcnt_t uint32
type __darwin_fsfilcnt_t uint32
type __darwin_gid_t __uint32_t
type __darwin_id_t __uint32_t
type __darwin_ino64_t __uint64_t
type __darwin_ino_t __darwin_ino64_t
type __darwin_mach_port_name_t __darwin_natural_t
type __darwin_mach_port_t __darwin_mach_port_name_t
type __darwin_mode_t __uint16_t
type __darwin_off_t __int64_t
type __darwin_pid_t __int32_t
type __darwin_sigset_t __uint32_t
type __darwin_suseconds_t __int32_t
type __darwin_uid_t __uint32_t
type __darwin_useconds_t __uint32_t
type __darwin_uuid_t []uint8
type __darwin_uuid_string_t []byte
type __darwin_pthread_handler_rec struct {
	__routine func(interface{})
	__arg     interface{}
	__next    []int64
}
type _opaque_pthread_attr_t struct {
	__sig    int32
	__opaque [56]byte
}
type _opaque_pthread_cond_t struct {
	__sig    int32
	__opaque [40]byte
}
type _opaque_pthread_condattr_t struct {
	__sig    int32
	__opaque [8]byte
}
type _opaque_pthread_mutex_t struct {
	__sig    int32
	__opaque [56]byte
}
type _opaque_pthread_mutexattr_t struct {
	__sig    int32
	__opaque [8]byte
}
type _opaque_pthread_once_t struct {
	__sig    int32
	__opaque [8]byte
}
type _opaque_pthread_rwlock_t struct {
	__sig    int32
	__opaque [192]byte
}
type _opaque_pthread_rwlockattr_t struct {
	__sig    int32
	__opaque [16]byte
}
type _opaque_pthread_t struct {
	__sig           int32
	__cleanup_stack []int64
	__opaque        [8176]byte
}
type __darwin_pthread_attr_t _opaque_pthread_attr_t
type __darwin_pthread_cond_t _opaque_pthread_cond_t
type __darwin_pthread_condattr_t _opaque_pthread_condattr_t
type __darwin_pthread_key_t uint32
type __darwin_pthread_mutex_t _opaque_pthread_mutex_t
type __darwin_pthread_mutexattr_t _opaque_pthread_mutexattr_t
type __darwin_pthread_once_t _opaque_pthread_once_t
type __darwin_pthread_rwlock_t _opaque_pthread_rwlock_t
type __darwin_pthread_rwlockattr_t _opaque_pthread_rwlockattr_t
type __darwin_pthread_t []_opaque_pthread_t
type u_int8_t uint8
type u_int16_t uint16
type u_int32_t uint32
type u_int64_t uint64
type register_t int64_t
type uintptr_t uint32
type user_addr_t u_int64_t
type user_size_t u_int64_t
type user_ssize_t int64_t
type user_long_t int64_t
type user_ulong_t u_int64_t
type user_time_t int64_t
type user_off_t int64_t
type syscall_arg_t u_int64_t
type intptr_t __darwin_intptr_t
type intmax_t int32
type uintmax_t uint32
type shn_ctx struct {
	R     [16]uint32_t
	CRC   [16]uint32_t
	initR [16]uint32_t
	konst uint32_t
	sbuf  uint32_t
	mbuf  uint32_t
	nbuf  int32
}
type __darwin_nl_item int32
type __darwin_wctrans_t int32
type __darwin_wctype_t __uint32_t

const P_ALL int32 = 0
const P_PID int32 = 1
const P_PGID int32 = 2

type idtype_t int32
type pid_t __darwin_pid_t
type id_t __darwin_id_t
type sig_atomic_t int32
type __darwin_i386_thread_state struct {
	__eax    uint32
	__ebx    uint32
	__ecx    uint32
	__edx    uint32
	__edi    uint32
	__esi    uint32
	__ebp    uint32
	__esp    uint32
	__ss     uint32
	__eflags uint32
	__eip    uint32
	__cs     uint32
	__ds     uint32
	__es     uint32
	__fs     uint32
	__gs     uint32
}
type __darwin_fp_control struct {
	__invalid uint16
	__denorm  uint16
	__zdiv    uint16
	__ovrfl   uint16
	__undfl   uint16
	__precis  uint16
	__pc      uint16
	__rc      uint16
}
type __darwin_fp_control_t __darwin_fp_control
type __darwin_fp_status struct {
	__invalid uint16
	__denorm  uint16
	__zdiv    uint16
	__ovrfl   uint16
	__undfl   uint16
	__precis  uint16
	__stkflt  uint16
	__errsumm uint16
	__c0      uint16
	__c1      uint16
	__c2      uint16
	__tos     uint16
	__c3      uint16
	__busy    uint16
}
type __darwin_fp_status_t __darwin_fp_status
type __darwin_mmst_reg struct {
	__mmst_reg  [10]byte
	__mmst_rsrv [6]byte
}
type __darwin_xmm_reg struct {
	__xmm_reg [16]byte
}
type __darwin_ymm_reg struct {
	__ymm_reg [32]byte
}
type __darwin_zmm_reg struct {
	__zmm_reg [64]byte
}
type __darwin_opmask_reg struct {
	__opmask_reg [8]byte
}
type __darwin_i386_float_state struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_rsrv4     [224]byte
	__fpu_reserved1 int32
}
type __darwin_i386_avx_state struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_rsrv4     [224]byte
	__fpu_reserved1 int32
	__avx_reserved1 [64]byte
	__fpu_ymmh0     __darwin_xmm_reg
	__fpu_ymmh1     __darwin_xmm_reg
	__fpu_ymmh2     __darwin_xmm_reg
	__fpu_ymmh3     __darwin_xmm_reg
	__fpu_ymmh4     __darwin_xmm_reg
	__fpu_ymmh5     __darwin_xmm_reg
	__fpu_ymmh6     __darwin_xmm_reg
	__fpu_ymmh7     __darwin_xmm_reg
}
type __darwin_i386_avx512_state struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_rsrv4     [224]byte
	__fpu_reserved1 int32
	__avx_reserved1 [64]byte
	__fpu_ymmh0     __darwin_xmm_reg
	__fpu_ymmh1     __darwin_xmm_reg
	__fpu_ymmh2     __darwin_xmm_reg
	__fpu_ymmh3     __darwin_xmm_reg
	__fpu_ymmh4     __darwin_xmm_reg
	__fpu_ymmh5     __darwin_xmm_reg
	__fpu_ymmh6     __darwin_xmm_reg
	__fpu_ymmh7     __darwin_xmm_reg
	__fpu_k0        __darwin_opmask_reg
	__fpu_k1        __darwin_opmask_reg
	__fpu_k2        __darwin_opmask_reg
	__fpu_k3        __darwin_opmask_reg
	__fpu_k4        __darwin_opmask_reg
	__fpu_k5        __darwin_opmask_reg
	__fpu_k6        __darwin_opmask_reg
	__fpu_k7        __darwin_opmask_reg
	__fpu_zmmh0     __darwin_ymm_reg
	__fpu_zmmh1     __darwin_ymm_reg
	__fpu_zmmh2     __darwin_ymm_reg
	__fpu_zmmh3     __darwin_ymm_reg
	__fpu_zmmh4     __darwin_ymm_reg
	__fpu_zmmh5     __darwin_ymm_reg
	__fpu_zmmh6     __darwin_ymm_reg
	__fpu_zmmh7     __darwin_ymm_reg
}
type __darwin_i386_exception_state struct {
	__trapno     __uint16_t
	__cpu        __uint16_t
	__err        __uint32_t
	__faultvaddr __uint32_t
}
type __darwin_x86_debug_state32 struct {
	__dr0 uint32
	__dr1 uint32
	__dr2 uint32
	__dr3 uint32
	__dr4 uint32
	__dr5 uint32
	__dr6 uint32
	__dr7 uint32
}
type __darwin_x86_thread_state64 struct {
	__rax    __uint64_t
	__rbx    __uint64_t
	__rcx    __uint64_t
	__rdx    __uint64_t
	__rdi    __uint64_t
	__rsi    __uint64_t
	__rbp    __uint64_t
	__rsp    __uint64_t
	__r8     __uint64_t
	__r9     __uint64_t
	__r10    __uint64_t
	__r11    __uint64_t
	__r12    __uint64_t
	__r13    __uint64_t
	__r14    __uint64_t
	__r15    __uint64_t
	__rip    __uint64_t
	__rflags __uint64_t
	__cs     __uint64_t
	__fs     __uint64_t
	__gs     __uint64_t
}
type __darwin_x86_float_state64 struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_xmm8      __darwin_xmm_reg
	__fpu_xmm9      __darwin_xmm_reg
	__fpu_xmm10     __darwin_xmm_reg
	__fpu_xmm11     __darwin_xmm_reg
	__fpu_xmm12     __darwin_xmm_reg
	__fpu_xmm13     __darwin_xmm_reg
	__fpu_xmm14     __darwin_xmm_reg
	__fpu_xmm15     __darwin_xmm_reg
	__fpu_rsrv4     [96]byte
	__fpu_reserved1 int32
}
type __darwin_x86_avx_state64 struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_xmm8      __darwin_xmm_reg
	__fpu_xmm9      __darwin_xmm_reg
	__fpu_xmm10     __darwin_xmm_reg
	__fpu_xmm11     __darwin_xmm_reg
	__fpu_xmm12     __darwin_xmm_reg
	__fpu_xmm13     __darwin_xmm_reg
	__fpu_xmm14     __darwin_xmm_reg
	__fpu_xmm15     __darwin_xmm_reg
	__fpu_rsrv4     [96]byte
	__fpu_reserved1 int32
	__avx_reserved1 [64]byte
	__fpu_ymmh0     __darwin_xmm_reg
	__fpu_ymmh1     __darwin_xmm_reg
	__fpu_ymmh2     __darwin_xmm_reg
	__fpu_ymmh3     __darwin_xmm_reg
	__fpu_ymmh4     __darwin_xmm_reg
	__fpu_ymmh5     __darwin_xmm_reg
	__fpu_ymmh6     __darwin_xmm_reg
	__fpu_ymmh7     __darwin_xmm_reg
	__fpu_ymmh8     __darwin_xmm_reg
	__fpu_ymmh9     __darwin_xmm_reg
	__fpu_ymmh10    __darwin_xmm_reg
	__fpu_ymmh11    __darwin_xmm_reg
	__fpu_ymmh12    __darwin_xmm_reg
	__fpu_ymmh13    __darwin_xmm_reg
	__fpu_ymmh14    __darwin_xmm_reg
	__fpu_ymmh15    __darwin_xmm_reg
}
type __darwin_x86_avx512_state64 struct {
	__fpu_reserved  [2]int32
	__fpu_fcw       __darwin_fp_control
	__fpu_fsw       __darwin_fp_status
	__fpu_ftw       __uint8_t
	__fpu_rsrv1     __uint8_t
	__fpu_fop       __uint16_t
	__fpu_ip        __uint32_t
	__fpu_cs        __uint16_t
	__fpu_rsrv2     __uint16_t
	__fpu_dp        __uint32_t
	__fpu_ds        __uint16_t
	__fpu_rsrv3     __uint16_t
	__fpu_mxcsr     __uint32_t
	__fpu_mxcsrmask __uint32_t
	__fpu_stmm0     __darwin_mmst_reg
	__fpu_stmm1     __darwin_mmst_reg
	__fpu_stmm2     __darwin_mmst_reg
	__fpu_stmm3     __darwin_mmst_reg
	__fpu_stmm4     __darwin_mmst_reg
	__fpu_stmm5     __darwin_mmst_reg
	__fpu_stmm6     __darwin_mmst_reg
	__fpu_stmm7     __darwin_mmst_reg
	__fpu_xmm0      __darwin_xmm_reg
	__fpu_xmm1      __darwin_xmm_reg
	__fpu_xmm2      __darwin_xmm_reg
	__fpu_xmm3      __darwin_xmm_reg
	__fpu_xmm4      __darwin_xmm_reg
	__fpu_xmm5      __darwin_xmm_reg
	__fpu_xmm6      __darwin_xmm_reg
	__fpu_xmm7      __darwin_xmm_reg
	__fpu_xmm8      __darwin_xmm_reg
	__fpu_xmm9      __darwin_xmm_reg
	__fpu_xmm10     __darwin_xmm_reg
	__fpu_xmm11     __darwin_xmm_reg
	__fpu_xmm12     __darwin_xmm_reg
	__fpu_xmm13     __darwin_xmm_reg
	__fpu_xmm14     __darwin_xmm_reg
	__fpu_xmm15     __darwin_xmm_reg
	__fpu_rsrv4     [96]byte
	__fpu_reserved1 int32
	__avx_reserved1 [64]byte
	__fpu_ymmh0     __darwin_xmm_reg
	__fpu_ymmh1     __darwin_xmm_reg
	__fpu_ymmh2     __darwin_xmm_reg
	__fpu_ymmh3     __darwin_xmm_reg
	__fpu_ymmh4     __darwin_xmm_reg
	__fpu_ymmh5     __darwin_xmm_reg
	__fpu_ymmh6     __darwin_xmm_reg
	__fpu_ymmh7     __darwin_xmm_reg
	__fpu_ymmh8     __darwin_xmm_reg
	__fpu_ymmh9     __darwin_xmm_reg
	__fpu_ymmh10    __darwin_xmm_reg
	__fpu_ymmh11    __darwin_xmm_reg
	__fpu_ymmh12    __darwin_xmm_reg
	__fpu_ymmh13    __darwin_xmm_reg
	__fpu_ymmh14    __darwin_xmm_reg
	__fpu_ymmh15    __darwin_xmm_reg
	__fpu_k0        __darwin_opmask_reg
	__fpu_k1        __darwin_opmask_reg
	__fpu_k2        __darwin_opmask_reg
	__fpu_k3        __darwin_opmask_reg
	__fpu_k4        __darwin_opmask_reg
	__fpu_k5        __darwin_opmask_reg
	__fpu_k6        __darwin_opmask_reg
	__fpu_k7        __darwin_opmask_reg
	__fpu_zmmh0     __darwin_ymm_reg
	__fpu_zmmh1     __darwin_ymm_reg
	__fpu_zmmh2     __darwin_ymm_reg
	__fpu_zmmh3     __darwin_ymm_reg
	__fpu_zmmh4     __darwin_ymm_reg
	__fpu_zmmh5     __darwin_ymm_reg
	__fpu_zmmh6     __darwin_ymm_reg
	__fpu_zmmh7     __darwin_ymm_reg
	__fpu_zmmh8     __darwin_ymm_reg
	__fpu_zmmh9     __darwin_ymm_reg
	__fpu_zmmh10    __darwin_ymm_reg
	__fpu_zmmh11    __darwin_ymm_reg
	__fpu_zmmh12    __darwin_ymm_reg
	__fpu_zmmh13    __darwin_ymm_reg
	__fpu_zmmh14    __darwin_ymm_reg
	__fpu_zmmh15    __darwin_ymm_reg
	__fpu_zmm16     __darwin_zmm_reg
	__fpu_zmm17     __darwin_zmm_reg
	__fpu_zmm18     __darwin_zmm_reg
	__fpu_zmm19     __darwin_zmm_reg
	__fpu_zmm20     __darwin_zmm_reg
	__fpu_zmm21     __darwin_zmm_reg
	__fpu_zmm22     __darwin_zmm_reg
	__fpu_zmm23     __darwin_zmm_reg
	__fpu_zmm24     __darwin_zmm_reg
	__fpu_zmm25     __darwin_zmm_reg
	__fpu_zmm26     __darwin_zmm_reg
	__fpu_zmm27     __darwin_zmm_reg
	__fpu_zmm28     __darwin_zmm_reg
	__fpu_zmm29     __darwin_zmm_reg
	__fpu_zmm30     __darwin_zmm_reg
	__fpu_zmm31     __darwin_zmm_reg
}
type __darwin_x86_exception_state64 struct {
	__trapno     __uint16_t
	__cpu        __uint16_t
	__err        __uint32_t
	__faultvaddr __uint64_t
}
type __darwin_x86_debug_state64 struct {
	__dr0 __uint64_t
	__dr1 __uint64_t
	__dr2 __uint64_t
	__dr3 __uint64_t
	__dr4 __uint64_t
	__dr5 __uint64_t
	__dr6 __uint64_t
	__dr7 __uint64_t
}
type __darwin_x86_cpmu_state64 struct {
	__ctrs [16]__uint64_t
}
type __darwin_mcontext32 struct {
	__es __darwin_i386_exception_state
	__ss __darwin_i386_thread_state
	__fs __darwin_i386_float_state
}
type __darwin_mcontext_avx32 struct {
	__es __darwin_i386_exception_state
	__ss __darwin_i386_thread_state
	__fs __darwin_i386_avx_state
}
type __darwin_mcontext_avx512_32 struct {
	__es __darwin_i386_exception_state
	__ss __darwin_i386_thread_state
	__fs __darwin_i386_avx512_state
}
type __darwin_mcontext64 struct {
	__es __darwin_x86_exception_state64
	__ss __darwin_x86_thread_state64
	__fs __darwin_x86_float_state64
}
type __darwin_mcontext_avx64 struct {
	__es __darwin_x86_exception_state64
	__ss __darwin_x86_thread_state64
	__fs __darwin_x86_avx_state64
}
type __darwin_mcontext_avx512_64 struct {
	__es __darwin_x86_exception_state64
	__ss __darwin_x86_thread_state64
	__fs __darwin_x86_avx512_state64
}
type mcontext_t []__darwin_mcontext64
type pthread_attr_t __darwin_pthread_attr_t
type __darwin_sigaltstack struct {
	ss_sp    interface{}
	ss_size  __darwin_size_t
	ss_flags int32
}
type stack_t __darwin_sigaltstack
type __darwin_ucontext struct {
	uc_onstack      int32
	uc_sigmask      __darwin_sigset_t
	uc_stack        __darwin_sigaltstack
	uc_link         []__darwin_ucontext
	uc_mcsize       __darwin_size_t
	uc_mcontext     []__darwin_mcontext64
	__mcontext_data __darwin_mcontext64
}
type ucontext_t __darwin_ucontext
type sigset_t __darwin_sigset_t
type size_t __darwin_size_t
type uid_t __darwin_uid_t
type sigval struct{ memory unsafe.Pointer }

func (unionVar *sigval) copy() sigval {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion sigval
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *sigval) sival_int() *int32 {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *sigval) sival_ptr() *interface{} {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*interface{})(unionVar.memory)
}

type sigevent struct {
	sigev_notify            int32
	sigev_signo             int32
	sigev_value             sigval
	sigev_notify_function   func(sigval)
	sigev_notify_attributes []pthread_attr_t
}
type __siginfo struct {
	si_signo  int32
	si_errno  int32
	si_code   int32
	si_pid    pid_t
	si_uid    uid_t
	si_status int32
	si_addr   interface{}
	si_value  sigval
	si_band   int32
	__pad     [7]uint32
}
type siginfo_t __siginfo
type __sigaction_u struct{ memory unsafe.Pointer }

func (unionVar *__sigaction_u) copy() __sigaction_u {
	var buffer [8]byte
	for i := range buffer {
		buffer[i] = (*((*[8]byte)(unionVar.memory)))[i]
	}
	var newUnion __sigaction_u
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *__sigaction_u) __sa_handler() *func(int32) {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*func(int32))(unionVar.memory)
}
func (unionVar *__sigaction_u) __sa_sigaction() *func(int32, []__siginfo, interface{}) {
	if unionVar.memory == nil {
		var buffer [8]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*func(int32, []__siginfo, interface{}))(unionVar.memory)
}

type sig_t func(int32)
type sigvec struct {
	sv_handler func(int32)
	sv_mask    int32
	sv_flags   int32
}
type sigstack struct {
	ss_sp      []byte
	ss_onstack int32
}
type timeval struct {
	tv_sec  __darwin_time_t
	tv_usec __darwin_suseconds_t
}
type rlim_t __uint64_t
type rusage struct {
	ru_utime    timeval
	ru_stime    timeval
	ru_maxrss   int32
	ru_ixrss    int32
	ru_idrss    int32
	ru_isrss    int32
	ru_minflt   int32
	ru_majflt   int32
	ru_nswap    int32
	ru_inblock  int32
	ru_oublock  int32
	ru_msgsnd   int32
	ru_msgrcv   int32
	ru_nsignals int32
	ru_nvcsw    int32
	ru_nivcsw   int32
}
type rusage_info_t interface{}
type rusage_info_v0 struct {
	ri_uuid               [16]uint8_t
	ri_user_time          uint64_t
	ri_system_time        uint64_t
	ri_pkg_idle_wkups     uint64_t
	ri_interrupt_wkups    uint64_t
	ri_pageins            uint64_t
	ri_wired_size         uint64_t
	ri_resident_size      uint64_t
	ri_phys_footprint     uint64_t
	ri_proc_start_abstime uint64_t
	ri_proc_exit_abstime  uint64_t
}
type rusage_info_v1 struct {
	ri_uuid                  [16]uint8_t
	ri_user_time             uint64_t
	ri_system_time           uint64_t
	ri_pkg_idle_wkups        uint64_t
	ri_interrupt_wkups       uint64_t
	ri_pageins               uint64_t
	ri_wired_size            uint64_t
	ri_resident_size         uint64_t
	ri_phys_footprint        uint64_t
	ri_proc_start_abstime    uint64_t
	ri_proc_exit_abstime     uint64_t
	ri_child_user_time       uint64_t
	ri_child_system_time     uint64_t
	ri_child_pkg_idle_wkups  uint64_t
	ri_child_interrupt_wkups uint64_t
	ri_child_pageins         uint64_t
	ri_child_elapsed_abstime uint64_t
}
type rusage_info_v2 struct {
	ri_uuid                  [16]uint8_t
	ri_user_time             uint64_t
	ri_system_time           uint64_t
	ri_pkg_idle_wkups        uint64_t
	ri_interrupt_wkups       uint64_t
	ri_pageins               uint64_t
	ri_wired_size            uint64_t
	ri_resident_size         uint64_t
	ri_phys_footprint        uint64_t
	ri_proc_start_abstime    uint64_t
	ri_proc_exit_abstime     uint64_t
	ri_child_user_time       uint64_t
	ri_child_system_time     uint64_t
	ri_child_pkg_idle_wkups  uint64_t
	ri_child_interrupt_wkups uint64_t
	ri_child_pageins         uint64_t
	ri_child_elapsed_abstime uint64_t
	ri_diskio_bytesread      uint64_t
	ri_diskio_byteswritten   uint64_t
}
type rusage_info_v3 struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
}
type rusage_info_v4 struct {
	ri_uuid                          [16]uint8_t
	ri_user_time                     uint64_t
	ri_system_time                   uint64_t
	ri_pkg_idle_wkups                uint64_t
	ri_interrupt_wkups               uint64_t
	ri_pageins                       uint64_t
	ri_wired_size                    uint64_t
	ri_resident_size                 uint64_t
	ri_phys_footprint                uint64_t
	ri_proc_start_abstime            uint64_t
	ri_proc_exit_abstime             uint64_t
	ri_child_user_time               uint64_t
	ri_child_system_time             uint64_t
	ri_child_pkg_idle_wkups          uint64_t
	ri_child_interrupt_wkups         uint64_t
	ri_child_pageins                 uint64_t
	ri_child_elapsed_abstime         uint64_t
	ri_diskio_bytesread              uint64_t
	ri_diskio_byteswritten           uint64_t
	ri_cpu_time_qos_default          uint64_t
	ri_cpu_time_qos_maintenance      uint64_t
	ri_cpu_time_qos_background       uint64_t
	ri_cpu_time_qos_utility          uint64_t
	ri_cpu_time_qos_legacy           uint64_t
	ri_cpu_time_qos_user_initiated   uint64_t
	ri_cpu_time_qos_user_interactive uint64_t
	ri_billed_system_time            uint64_t
	ri_serviced_system_time          uint64_t
	ri_logical_writes                uint64_t
	ri_lifetime_max_phys_footprint   uint64_t
	ri_instructions                  uint64_t
	ri_cycles                        uint64_t
	ri_billed_energy                 uint64_t
	ri_serviced_energy               uint64_t
	ri_unused                        [2]uint64_t
}
type rusage_info_current rusage_info_v4
type rlimit struct {
	rlim_cur rlim_t
	rlim_max rlim_t
}
type proc_rlimit_control_wakeupmon struct {
	wm_flags uint32_t
	wm_rate  int32_t
}

// _OSSwapInt16 - transpiled function from  /usr/include/libkern/i386/_OSByteOrder.h:44
// Warning (FieldDecl):  /usr/include/mach/i386/_structs.h:94 : Error : name of FieldDecl is empty
func _OSSwapInt16(_data __uint16_t) __uint16_t {
	return (__uint16_t((int32(uint16((_data)))<<uint64(int32(8)) | int32(uint16((_data)))>>uint64(int32(8)))))
}

// _OSSwapInt32 - transpiled function from  /usr/include/libkern/i386/_OSByteOrder.h:53
func _OSSwapInt32(_data __uint32_t) __uint32_t {
	return __uint32_t((darwin.BSwap32(uint32((__uint32_t(_data))))))
}

// _OSSwapInt64 - transpiled function from  /usr/include/libkern/i386/_OSByteOrder.h:68
func _OSSwapInt64(_data __uint64_t) __uint64_t {
	return __uint64_t((darwin.BSwap64(uint64((__uint64_t(_data))))))
}

type BSstructSatSSusrSincludeSsysSwaitPhD199D2E struct {
	w_Termsig  uint32
	w_Coredump uint32
	w_Retcode  uint32
	w_Filler   uint32
}
type BSstructSatSSusrSincludeSsysSwaitPhD218D2E struct {
	w_Stopval uint32
	w_Stopsig uint32
	w_Filler  uint32
}
type wait struct{ memory unsafe.Pointer }

func (unionVar *wait) copy() wait {
	var buffer [16]byte
	for i := range buffer {
		buffer[i] = (*((*[16]byte)(unionVar.memory)))[i]
	}
	var newUnion wait
	newUnion.memory = unsafe.Pointer(&buffer)
	return newUnion
}
func (unionVar *wait) w_status() *int32 {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*int32)(unionVar.memory)
}
func (unionVar *wait) w_T() *BSstructSatSSusrSincludeSsysSwaitPhD199D2E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSsysSwaitPhD199D2E)(unionVar.memory)
}
func (unionVar *wait) w_S() *BSstructSatSSusrSincludeSsysSwaitPhD218D2E {
	if unionVar.memory == nil {
		var buffer [16]byte
		unionVar.memory = unsafe.Pointer(&buffer)
	}
	return (*BSstructSatSSusrSincludeSsysSwaitPhD218D2E)(unionVar.memory)
}

type ct_rune_t darwin.CtRuneT
type rune_t __darwin_rune_t
type wchar_t __darwin_wchar_t
type div_t struct {
	quot int32
	rem  int32
}
type ldiv_t struct {
	quot int32
	rem  int32
}
type lldiv_t struct {
	quot int64
	rem  int64
}
type dev_t __darwin_dev_t
type mode_t __darwin_mode_t
type rsize_t __darwin_size_t
type errno_t int32
type ssize_t __darwin_ssize_t

// cycle - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:87
/* $Id: shnfast.c 442 2006-05-12 23:22:21Z ggr $ */ //
/* ShannonFast: Shannon stream cipher and MAC -- fast implementation */ //
/*
THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED
WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE AND AGAINST
INFRINGEMENT ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/ //
/* interface, multiplication table and SBox */ //
/*
 * FOLD is how many register cycles need to be performed after combining the
 * last byte of key and non-linear feedback, before every byte depends on every
 * byte of the key. This depends on the feedback and nonlinear functions, and
 * on where they are combined into the register. Making it same as the
 * register length is a safe and conservative choice.
 */ //
/* define IS_LITTLE_ENDIAN for faster operation when appropriate */ //
/* Useful macros -- machine independent little-endian version */ //
/* give correct offset for the current position of the register,
 * where logically R[0] is at position "zero". Note that this works for
 * both the stream register and the CRC register.
 */ //
/* step the shift register */ //
/* After stepping, "zero" moves right one place */ //
/* nonlinear feedback function */ //
/* shift register */ //
//
func cycle(c []shn_ctx) {
	var t uint32_t
	var i int32
	{
		t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
		t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
		t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
		c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
		t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
		t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
		t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
		c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
		c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
	}
	t = c[0].R[int32(0)]
	for i = int32(1); i < int32(16); i++ {
		c[0].R[i-int32(1)] = c[0].R[i]
	}
	c[0].R[int32(16)-int32(1)] = t
}

// crcfunc - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:121
/* The Shannon MAC function is modelled after the concepts of Phelix and SHA.
 * Basically, words to be accumulated in the MAC are incorporated in two
 * different ways:
 * 1. They are incorporated into the stream cipher register at a place
 *    where they will immediately have a nonlinear effect on the state
 * 2. They are incorporated into bit-parallel CRC-16 registers; the
 *    contents of these registers will be used in MAC finalization.
 */ //
/* Accumulate a CRC of input words, later to be fed into MAC.
 * This is actually 32 parallel CRC-16s, using the IBM CRC-16
 * polynomial x^16 + x^15 + x^2 + 1.
 */ //
/* now correct alignment of CRC accumulator */ //
//
func crcfunc(c []shn_ctx, i uint32_t) {
	var t uint32_t
	{
		c[0].CRC[(int32(0)+int32(0))%int32(16)] ^= c[0].CRC[(int32(0)+int32(2))%int32(16)] ^ c[0].CRC[(int32(0)+int32(15))%int32(16)] ^ i
	}
	t = c[0].CRC[int32(0)]
	for i = uint32_t(int32(1)); i < uint32_t((uint32(int32(16)))); i++ {
		c[0].CRC[i-uint32_t((uint32(int32(1))))] = c[0].CRC[i]
	}
	c[0].CRC[int32(16)-int32(1)] = t
}

// macfunc - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:136
/* Normal MAC word processing: do both SHA and CRC.
 */ //
//
func macfunc(c []shn_ctx, i uint32_t) {
	crcfunc(c, uint32_t(i))
	c[0].R[int32(13)] ^= i
}

// shn_initstate - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:145
/* initialise to known state
 */ //
/* Register initialised to Fibonacci numbers; Counter zeroed. */ //
//
func shn_initstate(c []shn_ctx) {
	var i int32
	c[0].R[int32(0)] = uint32_t(int32(1))
	c[0].R[int32(1)] = uint32_t(int32(1))
	for i = int32(2); i < int32(16); i++ {
		c[0].R[i] = c[0].R[i-int32(1)] + c[0].R[i-int32(2)]
	}
	c[0].konst = uint32_t(int32(1771488570))
}

// shn_savestate - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:160
/* Save the current register state
 */ //
//
func shn_savestate(c []shn_ctx) {
	var i int32
	for i = int32(0); i < int32(16); i++ {
		c[0].initR[i] = c[0].R[i]
	}
}

// shn_reloadstate - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:171
/* initialise to previously saved register state
 */ //
//
func shn_reloadstate(c []shn_ctx) {
	var i int32
	for i = int32(0); i < int32(16); i++ {
		c[0].R[i] = c[0].initR[i]
	}
}

// shn_genkonst - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:182
/* Initialise "konst"
 */ //
//
func shn_genkonst(c []shn_ctx) {
	c[0].konst = c[0].R[int32(0)]
}

// shn_diffuse - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:195
/* Load key material into the register
 */ //
/* nonlinear diffusion of register for key and MAC */ //
/* relies on FOLD == N! */ //
//
func shn_diffuse(c []shn_ctx) {
	{
		var t uint32_t
		{
			t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(1)+int32(12))%int32(16)] ^ c[0].R[(int32(1)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(1)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(1)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(1)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(1)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(1)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(1)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(2)+int32(12))%int32(16)] ^ c[0].R[(int32(2)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(2)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(2)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(2)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(2)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(2)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(2)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(3)+int32(12))%int32(16)] ^ c[0].R[(int32(3)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(3)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(3)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(3)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(3)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(3)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(3)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(4)+int32(12))%int32(16)] ^ c[0].R[(int32(4)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(4)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(4)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(4)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(4)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(4)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(4)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(5)+int32(12))%int32(16)] ^ c[0].R[(int32(5)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(5)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(5)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(5)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(5)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(5)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(5)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(6)+int32(12))%int32(16)] ^ c[0].R[(int32(6)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(6)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(6)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(6)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(6)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(6)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(6)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(7)+int32(12))%int32(16)] ^ c[0].R[(int32(7)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(7)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(7)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(7)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(7)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(7)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(7)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(8)+int32(12))%int32(16)] ^ c[0].R[(int32(8)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(8)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(8)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(8)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(8)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(8)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(8)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(9)+int32(12))%int32(16)] ^ c[0].R[(int32(9)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(9)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(9)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(9)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(9)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(9)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(9)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(10)+int32(12))%int32(16)] ^ c[0].R[(int32(10)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(10)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(10)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(10)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(10)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(10)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(10)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(11)+int32(12))%int32(16)] ^ c[0].R[(int32(11)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(11)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(11)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(11)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(11)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(11)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(11)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(12)+int32(12))%int32(16)] ^ c[0].R[(int32(12)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(12)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(12)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(12)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(12)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(12)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(12)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(13)+int32(12))%int32(16)] ^ c[0].R[(int32(13)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(13)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(13)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(13)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(13)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(13)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(13)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(14)+int32(12))%int32(16)] ^ c[0].R[(int32(14)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(14)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(14)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(14)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(14)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(14)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(14)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(12))%int32(16)]
		}
	}
	{
		var t uint32_t
		{
			t = c[0].R[(int32(15)+int32(12))%int32(16)] ^ c[0].R[(int32(15)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
			t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			c[0].R[(int32(15)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(15)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(15)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
			t = c[0].R[(int32(15)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(15))%int32(16)]
			t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
			t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
			c[0].R[(int32(15)+int32(1)+int32(0))%int32(16)] ^= t
			c[0].sbuf = t ^ c[0].R[(int32(15)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(12))%int32(16)]
		}
	}
}

// shn_loadkey - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:221
/* common actions for loading key material
 * Allow non-word-multiple key and nonce materianl
 * Note also initializes the CRC register as a side effect.
 */ //
/* start folding in key */ //
/* if there were any extra key bytes, zero pad to a word */ //
/* i unchanged */ //
/* j unchanged */ //
/* also fold in the length of the key */ //
/* save a copy of the register */ //
/* now diffuse */ //
/* now xor the copy back -- makes key loading irreversible */ //
//
func shn_loadkey(c []shn_ctx, key []uint8_t, keylen int32) {
	var i int32
	var j int32
	var k uint32_t
	var xtra []uint8_t = make([]uint8_t, 4, 4)
	for i = int32(0); i < keylen & ^int32(3); i += int32(4) {
		k = uint32_t(uint8_t(((*[1000000000]uint8_t)(unsafe.Pointer(&key[i]))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*[1000000000]uint8_t)(unsafe.Pointer(&key[i]))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*[1000000000]uint8_t)(unsafe.Pointer(&key[i]))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*[1000000000]uint8_t)(unsafe.Pointer(&key[i]))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
		c[0].R[int32(13)] ^= k
		cycle(c)
	}
	if i < keylen {
		for j = int32(0); i < keylen; i++ {
			xtra[func() int32 {
				defer func() {
					j += 1
				}()
				return j
			}()] = key[i]
		}
		for ; j < int32(4); j++ {
			xtra[j] = uint8_t(int32(0))
		}
		k = uint32_t(uint8_t((xtra)[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t((xtra)[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t((xtra)[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t((xtra)[int32(0)]))&uint32_t((uint32(int32(255))))
		c[0].R[int32(13)] ^= k
		cycle(c)
	}
	c[0].R[int32(13)] ^= uint32_t((uint32_t((uint32(keylen)))))
	cycle(c)
	for i = int32(0); i < int32(16); i++ {
		c[0].CRC[i] = c[0].R[i]
	}
	shn_diffuse(c)
	for i = int32(0); i < int32(16); i++ {
		c[0].R[i] ^= c[0].CRC[i]
	}
}

// shn_key - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:265
/* Published "key" interface
 */ //
//
func shn_key(c []shn_ctx, key []uint8_t, keylen int32) {
	shn_initstate(c)
	shn_loadkey(c, key, keylen)
	shn_genkonst(c)
	shn_savestate(c)
	c[0].nbuf = int32(0)
}

// shn_nonce - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:277
/* Published "nonce" interface
 */ //
//
func shn_nonce(c []shn_ctx, nonce []uint8_t, noncelen int32) {
	shn_reloadstate(c)
	c[0].konst = uint32_t(int32(1771488570))
	shn_loadkey(c, nonce, noncelen)
	shn_genkonst(c)
	c[0].nbuf = int32(0)
}

// shn_stream - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:295
/* XOR pseudo-random bytes into buffer
 * Note: doesn't play well with MAC functions.
 */ //
/* handle any previously buffered bytes */ //
/* do lots at a time, if there's enough to do */ //
/* do small or odd size buffers the slow way */ //
/* handle any trailing bytes */ //
//
func shn_stream(c []shn_ctx, buf []uint8_t, nbytes int32) {
	for c[0].nbuf != int32(0) && nbytes != int32(0) {
		(func() []uint8_t {
			defer func() {
				func() []uint8_t {
					buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
					return buf
				}()
			}()
			return buf
		}())[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) & uint32_t((uint32(int32(255))))))))))))))
		c[0].sbuf >>= uint32_t((uint32(uint64(int32(8)))))
		c[0].nbuf -= int32(8)
		nbytes -= 1
	}
	for nbytes >= int32(16)*int32(4) {
		{
			var t uint32_t
			{
				t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(1)+int32(12))%int32(16)] ^ c[0].R[(int32(1)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(1)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(1)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(1)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(1)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(1)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(1)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(2)+int32(12))%int32(16)] ^ c[0].R[(int32(2)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(2)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(2)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(2)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(2)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(2)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(2)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(3)+int32(12))%int32(16)] ^ c[0].R[(int32(3)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(3)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(3)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(3)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(3)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(3)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(3)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(4)+int32(12))%int32(16)] ^ c[0].R[(int32(4)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(4)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(4)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(4)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(4)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(4)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(4)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(5)+int32(12))%int32(16)] ^ c[0].R[(int32(5)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(5)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(5)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(5)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(5)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(5)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(5)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(6)+int32(12))%int32(16)] ^ c[0].R[(int32(6)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(6)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(6)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(6)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(6)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(6)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(6)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(7)+int32(12))%int32(16)] ^ c[0].R[(int32(7)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(7)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(7)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(7)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(7)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(7)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(7)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(8)+int32(12))%int32(16)] ^ c[0].R[(int32(8)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(8)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(8)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(8)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(8)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(8)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(8)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(9)+int32(12))%int32(16)] ^ c[0].R[(int32(9)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(9)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(9)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(9)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(9)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(9)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(9)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(10)+int32(12))%int32(16)] ^ c[0].R[(int32(10)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(10)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(10)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(10)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(10)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(10)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(10)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(11)+int32(12))%int32(16)] ^ c[0].R[(int32(11)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(11)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(11)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(11)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(11)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(11)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(11)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(12)+int32(12))%int32(16)] ^ c[0].R[(int32(12)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(12)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(12)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(12)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(12)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(12)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(12)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(13)+int32(12))%int32(16)] ^ c[0].R[(int32(13)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(13)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(13)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(13)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(13)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(13)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(13)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(14)+int32(12))%int32(16)] ^ c[0].R[(int32(14)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(14)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(14)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(14)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(14)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(14)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(14)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		{
			var t uint32_t
			{
				t = c[0].R[(int32(15)+int32(12))%int32(16)] ^ c[0].R[(int32(15)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(15)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(15)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(15)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(15)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(15)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(15)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(12))%int32(16)]
			}
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
			}
		}
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(16))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(16) * int32(4)
	}
	for int32(4) <= nbytes {
		cycle(c)
		{
			(buf)[int32(3)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255)))))))))))))))
			(buf)[int32(2)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255)))))))))))))))
			(buf)[int32(1)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255)))))))))))))))
			(buf)[int32(0)] ^= uint8_t((uint8(uint8_t((uint8(int32(uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255)))))))))))))))
		}
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(4)
	}
	if nbytes != int32(0) {
		cycle(c)
		c[0].nbuf = int32(32)
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			(func() []uint8_t {
				defer func() {
					func() []uint8_t {
						buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
						return buf
					}()
				}()
				return buf
			}())[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) & uint32_t((uint32(int32(255))))))))))))))
			c[0].sbuf >>= uint32_t((uint32(uint64(int32(8)))))
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
	}
}

// shn_maconly - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:360
/* accumulate words into MAC without encryption
 * Note that plaintext is accumulated for MAC.
 */ //
/* handle any previously buffered bytes */ //
/* not a whole word yet */ //
/* LFSR already cycled */ //
/* do lots at a time, if there's enough to do */ //
/* do small or odd size buffers the slow way */ //
/* handle any trailing bytes */ //
//
func shn_maconly(c []shn_ctx, buf []uint8_t, nbytes int32) {
	if c[0].nbuf != int32(0) {
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8(((func() []uint8_t {
				defer func() {
					func() []uint8_t {
						buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
						return buf
					}()
				}()
				return buf
			}())[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
		if c[0].nbuf != int32(0) {
			return
		}
		macfunc(c, uint32_t(c[0].mbuf))
	}
	for int32(4)*int32(16) <= nbytes {
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(0)+int32(0))%int32(16)] ^= c[0].CRC[(int32(0)+int32(2))%int32(16)] ^ c[0].CRC[(int32(0)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(0)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(1)+int32(12))%int32(16)] ^ c[0].R[(int32(1)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(1)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(1)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(1)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(1)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(1)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(1)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(1)+int32(0))%int32(16)] ^= c[0].CRC[(int32(1)+int32(2))%int32(16)] ^ c[0].CRC[(int32(1)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(1)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(2)+int32(12))%int32(16)] ^ c[0].R[(int32(2)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(2)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(2)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(2)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(2)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(2)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(2)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(2)+int32(0))%int32(16)] ^= c[0].CRC[(int32(2)+int32(2))%int32(16)] ^ c[0].CRC[(int32(2)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(2)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(3)+int32(12))%int32(16)] ^ c[0].R[(int32(3)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(3)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(3)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(3)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(3)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(3)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(3)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(3)+int32(0))%int32(16)] ^= c[0].CRC[(int32(3)+int32(2))%int32(16)] ^ c[0].CRC[(int32(3)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(3)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(4)+int32(12))%int32(16)] ^ c[0].R[(int32(4)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(4)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(4)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(4)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(4)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(4)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(4)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(4)+int32(0))%int32(16)] ^= c[0].CRC[(int32(4)+int32(2))%int32(16)] ^ c[0].CRC[(int32(4)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(4)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(5)+int32(12))%int32(16)] ^ c[0].R[(int32(5)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(5)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(5)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(5)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(5)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(5)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(5)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(5)+int32(0))%int32(16)] ^= c[0].CRC[(int32(5)+int32(2))%int32(16)] ^ c[0].CRC[(int32(5)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(5)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(6)+int32(12))%int32(16)] ^ c[0].R[(int32(6)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(6)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(6)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(6)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(6)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(6)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(6)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(6)+int32(0))%int32(16)] ^= c[0].CRC[(int32(6)+int32(2))%int32(16)] ^ c[0].CRC[(int32(6)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(6)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(7)+int32(12))%int32(16)] ^ c[0].R[(int32(7)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(7)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(7)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(7)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(7)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(7)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(7)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(7)+int32(0))%int32(16)] ^= c[0].CRC[(int32(7)+int32(2))%int32(16)] ^ c[0].CRC[(int32(7)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(7)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(8)+int32(12))%int32(16)] ^ c[0].R[(int32(8)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(8)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(8)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(8)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(8)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(8)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(8)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(8)+int32(0))%int32(16)] ^= c[0].CRC[(int32(8)+int32(2))%int32(16)] ^ c[0].CRC[(int32(8)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(8)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(9)+int32(12))%int32(16)] ^ c[0].R[(int32(9)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(9)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(9)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(9)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(9)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(9)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(9)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(9)+int32(0))%int32(16)] ^= c[0].CRC[(int32(9)+int32(2))%int32(16)] ^ c[0].CRC[(int32(9)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(9)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(10)+int32(12))%int32(16)] ^ c[0].R[(int32(10)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(10)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(10)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(10)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(10)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(10)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(10)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(10)+int32(0))%int32(16)] ^= c[0].CRC[(int32(10)+int32(2))%int32(16)] ^ c[0].CRC[(int32(10)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(10)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(11)+int32(12))%int32(16)] ^ c[0].R[(int32(11)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(11)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(11)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(11)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(11)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(11)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(11)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(11)+int32(0))%int32(16)] ^= c[0].CRC[(int32(11)+int32(2))%int32(16)] ^ c[0].CRC[(int32(11)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(11)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(12)+int32(12))%int32(16)] ^ c[0].R[(int32(12)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(12)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(12)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(12)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(12)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(12)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(12)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(12)+int32(0))%int32(16)] ^= c[0].CRC[(int32(12)+int32(2))%int32(16)] ^ c[0].CRC[(int32(12)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(12)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(13)+int32(12))%int32(16)] ^ c[0].R[(int32(13)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(13)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(13)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(13)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(13)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(13)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(13)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(13)+int32(0))%int32(16)] ^= c[0].CRC[(int32(13)+int32(2))%int32(16)] ^ c[0].CRC[(int32(13)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(13)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(14)+int32(12))%int32(16)] ^ c[0].R[(int32(14)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(14)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(14)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(14)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(14)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(14)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(14)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(14)+int32(0))%int32(16)] ^= c[0].CRC[(int32(14)+int32(2))%int32(16)] ^ c[0].CRC[(int32(14)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(14)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		{
			var t uint32_t
			var t1 uint32_t
			t1 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				t = c[0].R[(int32(15)+int32(12))%int32(16)] ^ c[0].R[(int32(15)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(15)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(15)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(15)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(15)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(15)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(15)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(12))%int32(16)]
			}
			{
				c[0].CRC[(int32(15)+int32(0))%int32(16)] ^= c[0].CRC[(int32(15)+int32(2))%int32(16)] ^ c[0].CRC[(int32(15)+int32(15))%int32(16)] ^ t1
			}
			c[0].R[(int32(15)+int32(1)+int32(13))%int32(16)] ^= t1
		}
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(16))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(4) * int32(16)
	}
	for int32(4) <= nbytes {
		cycle(c)
		macfunc(c, uint32_t((uint32((uint32_t(uint8_t((buf)[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t((buf)[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t((buf)[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t((buf)[int32(0)]))&uint32_t((uint32(int32(255)))))))))
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(4)
	}
	if nbytes != int32(0) {
		cycle(c)
		c[0].mbuf = uint32_t(int32(0))
		c[0].nbuf = int32(32)
		for nbytes != int32(0) {
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8(((func() []uint8_t {
				defer func() {
					func() []uint8_t {
						buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
						return buf
					}()
				}()
				return buf
			}())[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
	}
}

// shn_encrypt - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:432
/* Combined MAC and encryption.
 * Note that plaintext is accumulated for MAC.
 */ //
/* handle any previously buffered bytes */ //
/* not a whole word yet */ //
/* LFSR already cycled */ //
/* do lots at a time, if there's enough to do */ //
/* do small or odd size buffers the slow way */ //
/* handle any trailing bytes */ //
//
func shn_encrypt(c []shn_ctx, buf []uint8_t, nbytes int32) {
	var t3 uint32_t = uint32_t(int32(0))
	if c[0].nbuf != int32(0) {
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8((buf[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			buf[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) >> uint64(int32(32)-c[0].nbuf) & uint32_t((uint32(int32(255))))))))))))))
			buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
		if c[0].nbuf != int32(0) {
			return
		}
		macfunc(c, uint32_t(c[0].mbuf))
	}
	for int32(4)*int32(16) <= nbytes {
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(0)+int32(0))%int32(16)] ^= c[0].CRC[(int32(0)+int32(2))%int32(16)] ^ c[0].CRC[(int32(0)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(0)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(1)+int32(12))%int32(16)] ^ c[0].R[(int32(1)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(1)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(1)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(1)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(1)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(1)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(1)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(1)+int32(0))%int32(16)] ^= c[0].CRC[(int32(1)+int32(2))%int32(16)] ^ c[0].CRC[(int32(1)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(1)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(2)+int32(12))%int32(16)] ^ c[0].R[(int32(2)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(2)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(2)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(2)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(2)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(2)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(2)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(2)+int32(0))%int32(16)] ^= c[0].CRC[(int32(2)+int32(2))%int32(16)] ^ c[0].CRC[(int32(2)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(2)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(3)+int32(12))%int32(16)] ^ c[0].R[(int32(3)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(3)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(3)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(3)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(3)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(3)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(3)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(3)+int32(0))%int32(16)] ^= c[0].CRC[(int32(3)+int32(2))%int32(16)] ^ c[0].CRC[(int32(3)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(3)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(4)+int32(12))%int32(16)] ^ c[0].R[(int32(4)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(4)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(4)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(4)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(4)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(4)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(4)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(4)+int32(0))%int32(16)] ^= c[0].CRC[(int32(4)+int32(2))%int32(16)] ^ c[0].CRC[(int32(4)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(4)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(5)+int32(12))%int32(16)] ^ c[0].R[(int32(5)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(5)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(5)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(5)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(5)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(5)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(5)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(5)+int32(0))%int32(16)] ^= c[0].CRC[(int32(5)+int32(2))%int32(16)] ^ c[0].CRC[(int32(5)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(5)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(6)+int32(12))%int32(16)] ^ c[0].R[(int32(6)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(6)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(6)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(6)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(6)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(6)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(6)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(6)+int32(0))%int32(16)] ^= c[0].CRC[(int32(6)+int32(2))%int32(16)] ^ c[0].CRC[(int32(6)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(6)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(7)+int32(12))%int32(16)] ^ c[0].R[(int32(7)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(7)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(7)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(7)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(7)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(7)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(7)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(7)+int32(0))%int32(16)] ^= c[0].CRC[(int32(7)+int32(2))%int32(16)] ^ c[0].CRC[(int32(7)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(7)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(8)+int32(12))%int32(16)] ^ c[0].R[(int32(8)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(8)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(8)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(8)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(8)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(8)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(8)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(8)+int32(0))%int32(16)] ^= c[0].CRC[(int32(8)+int32(2))%int32(16)] ^ c[0].CRC[(int32(8)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(8)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(9)+int32(12))%int32(16)] ^ c[0].R[(int32(9)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(9)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(9)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(9)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(9)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(9)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(9)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(9)+int32(0))%int32(16)] ^= c[0].CRC[(int32(9)+int32(2))%int32(16)] ^ c[0].CRC[(int32(9)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(9)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(10)+int32(12))%int32(16)] ^ c[0].R[(int32(10)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(10)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(10)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(10)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(10)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(10)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(10)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(10)+int32(0))%int32(16)] ^= c[0].CRC[(int32(10)+int32(2))%int32(16)] ^ c[0].CRC[(int32(10)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(10)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(11)+int32(12))%int32(16)] ^ c[0].R[(int32(11)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(11)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(11)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(11)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(11)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(11)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(11)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(11)+int32(0))%int32(16)] ^= c[0].CRC[(int32(11)+int32(2))%int32(16)] ^ c[0].CRC[(int32(11)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(11)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(12)+int32(12))%int32(16)] ^ c[0].R[(int32(12)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(12)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(12)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(12)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(12)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(12)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(12)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(12)+int32(0))%int32(16)] ^= c[0].CRC[(int32(12)+int32(2))%int32(16)] ^ c[0].CRC[(int32(12)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(12)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(13)+int32(12))%int32(16)] ^ c[0].R[(int32(13)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(13)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(13)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(13)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(13)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(13)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(13)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(13)+int32(0))%int32(16)] ^= c[0].CRC[(int32(13)+int32(2))%int32(16)] ^ c[0].CRC[(int32(13)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(13)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(14)+int32(12))%int32(16)] ^ c[0].R[(int32(14)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(14)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(14)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(14)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(14)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(14)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(14)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(14)+int32(0))%int32(16)] ^= c[0].CRC[(int32(14)+int32(2))%int32(16)] ^ c[0].CRC[(int32(14)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(14)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(15)+int32(12))%int32(16)] ^ c[0].R[(int32(15)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(15)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(15)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(15)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(15)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(15)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(15)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			{
				c[0].CRC[(int32(15)+int32(0))%int32(16)] ^= c[0].CRC[(int32(15)+int32(2))%int32(16)] ^ c[0].CRC[(int32(15)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(15)+int32(1)+int32(13))%int32(16)] ^= t3
			t3 ^= uint32_t(c[0].sbuf)
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(16))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(4) * int32(16)
	}
	for int32(4) <= nbytes {
		cycle(c)
		t3 = uint32_t(uint8_t((buf)[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t((buf)[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t((buf)[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t((buf)[int32(0)]))&uint32_t((uint32(int32(255))))
		macfunc(c, uint32_t(t3))
		t3 ^= uint32_t(c[0].sbuf)
		{
			(buf)[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
		}
		nbytes -= int32(4)
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4))*unsafe.Sizeof(buf[0]))))[:]
	}
	if nbytes != int32(0) {
		cycle(c)
		c[0].mbuf = uint32_t(int32(0))
		c[0].nbuf = int32(32)
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8((buf[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			buf[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) >> uint64(int32(32)-c[0].nbuf) & uint32_t((uint32(int32(255))))))))))))))
			buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
	}
}

// shn_decrypt - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:514
/* Combined MAC and decryption.
 * Note that plaintext is accumulated for MAC.
 */ //
/* handle any previously buffered bytes */ //
/* not a whole word yet */ //
/* LFSR already cycled */ //
/* now do lots at a time, if there's enough */ //
/* do small or odd size buffers the slow way */ //
/* handle any trailing bytes */ //
//
func shn_decrypt(c []shn_ctx, buf []uint8_t, nbytes int32) {
	var t3 uint32_t = uint32_t(int32(0))
	if c[0].nbuf != int32(0) {
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			buf[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) >> uint64(int32(32)-c[0].nbuf) & uint32_t((uint32(int32(255))))))))))))))
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8((buf[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
		if c[0].nbuf != int32(0) {
			return
		}
		macfunc(c, uint32_t(c[0].mbuf))
	}
	for int32(4)*int32(16) <= nbytes {
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(0)+int32(12))%int32(16)] ^ c[0].R[(int32(0)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(0)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(0)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(0)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(0)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(0)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(0)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(0)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(0)+int32(0))%int32(16)] ^= c[0].CRC[(int32(0)+int32(2))%int32(16)] ^ c[0].CRC[(int32(0)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(0)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(0)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(1)+int32(12))%int32(16)] ^ c[0].R[(int32(1)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(1)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(1)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(1)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(1)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(1)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(1)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(1)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(1)+int32(0))%int32(16)] ^= c[0].CRC[(int32(1)+int32(2))%int32(16)] ^ c[0].CRC[(int32(1)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(1)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(1)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(2)+int32(12))%int32(16)] ^ c[0].R[(int32(2)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(2)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(2)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(2)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(2)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(2)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(2)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(2)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(2)+int32(0))%int32(16)] ^= c[0].CRC[(int32(2)+int32(2))%int32(16)] ^ c[0].CRC[(int32(2)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(2)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(2)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(3)+int32(12))%int32(16)] ^ c[0].R[(int32(3)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(3)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(3)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(3)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(3)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(3)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(3)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(3)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(3)+int32(0))%int32(16)] ^= c[0].CRC[(int32(3)+int32(2))%int32(16)] ^ c[0].CRC[(int32(3)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(3)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(3)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(4)+int32(12))%int32(16)] ^ c[0].R[(int32(4)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(4)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(4)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(4)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(4)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(4)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(4)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(4)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(4)+int32(0))%int32(16)] ^= c[0].CRC[(int32(4)+int32(2))%int32(16)] ^ c[0].CRC[(int32(4)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(4)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(5)+int32(12))%int32(16)] ^ c[0].R[(int32(5)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(5)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(5)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(5)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(5)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(5)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(5)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(5)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(5)+int32(0))%int32(16)] ^= c[0].CRC[(int32(5)+int32(2))%int32(16)] ^ c[0].CRC[(int32(5)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(5)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(5)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(6)+int32(12))%int32(16)] ^ c[0].R[(int32(6)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(6)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(6)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(6)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(6)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(6)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(6)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(6)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(6)+int32(0))%int32(16)] ^= c[0].CRC[(int32(6)+int32(2))%int32(16)] ^ c[0].CRC[(int32(6)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(6)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(6)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(7)+int32(12))%int32(16)] ^ c[0].R[(int32(7)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(7)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(7)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(7)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(7)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(7)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(7)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(7)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(7)+int32(0))%int32(16)] ^= c[0].CRC[(int32(7)+int32(2))%int32(16)] ^ c[0].CRC[(int32(7)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(7)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(7)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(8)+int32(12))%int32(16)] ^ c[0].R[(int32(8)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(8)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(8)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(8)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(8)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(8)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(8)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(8)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(8)+int32(0))%int32(16)] ^= c[0].CRC[(int32(8)+int32(2))%int32(16)] ^ c[0].CRC[(int32(8)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(8)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(8)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(9)+int32(12))%int32(16)] ^ c[0].R[(int32(9)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(9)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(9)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(9)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(9)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(9)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(9)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(9)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(9)+int32(0))%int32(16)] ^= c[0].CRC[(int32(9)+int32(2))%int32(16)] ^ c[0].CRC[(int32(9)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(9)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(9)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(10)+int32(12))%int32(16)] ^ c[0].R[(int32(10)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(10)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(10)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(10)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(10)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(10)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(10)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(10)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(10)+int32(0))%int32(16)] ^= c[0].CRC[(int32(10)+int32(2))%int32(16)] ^ c[0].CRC[(int32(10)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(10)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(10)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(11)+int32(12))%int32(16)] ^ c[0].R[(int32(11)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(11)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(11)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(11)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(11)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(11)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(11)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(11)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(11)+int32(0))%int32(16)] ^= c[0].CRC[(int32(11)+int32(2))%int32(16)] ^ c[0].CRC[(int32(11)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(11)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(11)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(12)+int32(12))%int32(16)] ^ c[0].R[(int32(12)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(12)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(12)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(12)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(12)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(12)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(12)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(12)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(12)+int32(0))%int32(16)] ^= c[0].CRC[(int32(12)+int32(2))%int32(16)] ^ c[0].CRC[(int32(12)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(12)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(12)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(13)+int32(12))%int32(16)] ^ c[0].R[(int32(13)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(13)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(13)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(13)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(13)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(13)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(13)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(13)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(13)+int32(0))%int32(16)] ^= c[0].CRC[(int32(13)+int32(2))%int32(16)] ^ c[0].CRC[(int32(13)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(13)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(13)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(14)+int32(12))%int32(16)] ^ c[0].R[(int32(14)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(14)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(14)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(14)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(14)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(14)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(14)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(14)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(14)+int32(0))%int32(16)] ^= c[0].CRC[(int32(14)+int32(2))%int32(16)] ^ c[0].CRC[(int32(14)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(14)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(14)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		{
			var t uint32_t
			var t3 uint32_t
			{
				t = c[0].R[(int32(15)+int32(12))%int32(16)] ^ c[0].R[(int32(15)+int32(13))%int32(16)] ^ uint32_t(c[0].konst)
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)))
				t ^= t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				c[0].R[(int32(15)+int32(0))%int32(16)] = t ^ (c[0].R[(int32(15)+int32(0))%int32(16)]<<uint64(int32(1)) | c[0].R[(int32(15)+int32(0))%int32(16)]>>uint64(int32(32)-int32(1)))
				t = c[0].R[(int32(15)+int32(1)+int32(2))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(15))%int32(16)]
				t ^= t<<uint64(int32(7)) | t>>uint64(int32(32)-int32(7)) | (t<<uint64(int32(22)) | t>>uint64(int32(32)-int32(22)))
				t ^= t<<uint64(int32(5)) | t>>uint64(int32(32)-int32(5)) | (t<<uint64(int32(19)) | t>>uint64(int32(32)-int32(19)))
				c[0].R[(int32(15)+int32(1)+int32(0))%int32(16)] ^= t
				c[0].sbuf = t ^ c[0].R[(int32(15)+int32(1)+int32(8))%int32(16)] ^ c[0].R[(int32(15)+int32(1)+int32(12))%int32(16)]
			}
			t3 = uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t(((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)]))&uint32_t((uint32(int32(255))))
			t3 ^= uint32_t(c[0].sbuf)
			{
				c[0].CRC[(int32(15)+int32(0))%int32(16)] ^= c[0].CRC[(int32(15)+int32(2))%int32(16)] ^ c[0].CRC[(int32(15)+int32(15))%int32(16)] ^ t3
			}
			c[0].R[(int32(15)+int32(1)+int32(13))%int32(16)] ^= t3
			{
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				((*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(15)*int32(4))*unsafe.Sizeof(buf[0]))))[:])[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
		}
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4)*int32(16))*unsafe.Sizeof(buf[0]))))[:]
		nbytes -= int32(4) * int32(16)
	}
	for int32(4) <= nbytes {
		cycle(c)
		t3 = uint32_t(uint8_t((buf)[int32(3)]))&uint32_t((uint32(int32(255))))<<uint64(int32(24)) | uint32_t(uint8_t((buf)[int32(2)]))&uint32_t((uint32(int32(255))))<<uint64(int32(16)) | uint32_t(uint8_t((buf)[int32(1)]))&uint32_t((uint32(int32(255))))<<uint64(int32(8)) | uint32_t(uint8_t((buf)[int32(0)]))&uint32_t((uint32(int32(255))))
		t3 ^= uint32_t(c[0].sbuf)
		macfunc(c, uint32_t(t3))
		{
			(buf)[int32(3)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(2)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(1)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
			(buf)[int32(0)] = uint8_t((uint8((uint32((t3 >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
		}
		nbytes -= int32(4)
		buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4))*unsafe.Sizeof(buf[0]))))[:]
	}
	if nbytes != int32(0) {
		cycle(c)
		c[0].mbuf = uint32_t(int32(0))
		c[0].nbuf = int32(32)
		for c[0].nbuf != int32(0) && nbytes != int32(0) {
			buf[int32(0)] ^= uint8_t((uint8(uint32((uint8_t((uint8(uint32((uint32_t(c[0].sbuf) >> uint64(int32(32)-c[0].nbuf) & uint32_t((uint32(int32(255))))))))))))))
			c[0].mbuf ^= uint32_t((uint32_t((uint32(int32(uint8((buf[int32(0)]))) << uint64(int32(32)-c[0].nbuf))))))
			buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(1)*unsafe.Sizeof(buf[0]))))[:]
			c[0].nbuf -= int32(8)
			nbytes -= 1
		}
	}
}

// shn_finish - transpiled function from  /Users/thomasalbrighton/Projects/spotigo/docs/shannon/ShannonFast.c:587
/* Having accumulated a MAC, finish processing and return it.
 * Note that any unprocessed bytes are treated as if
 * they were encrypted zero bytes, so plaintext (zero) is accumulated.
 */ //
/* handle any previously buffered bytes */ //
/* LFSR already cycled */ //
/* perturb the MAC to mark end of input.
 * Note that only the stream register is updated, not the CRC. This is an
 * action that can't be duplicated by passing in plaintext, hence
 * defeating any kind of extension attack.
 */ //
/* now add the CRC to the stream register and diffuse it */ //
/* produce output from the stream buffer */ //
//
func shn_finish(c []shn_ctx, buf []uint8_t, nbytes int32) {
	var i int32
	if c[0].nbuf != int32(0) {
		macfunc(c, uint32_t(c[0].mbuf))
	}
	cycle(c)
	c[0].R[int32(13)] ^= uint32_t((uint32_t((uint32(int32(1771488570) ^ c[0].nbuf<<uint64(int32(3)))))))
	c[0].nbuf = int32(0)
	for i = int32(0); i < int32(16); i++ {
		c[0].R[i] ^= c[0].CRC[i]
	}
	shn_diffuse(c)
	for nbytes > int32(0) {
		cycle(c)
		if nbytes >= int32(4) {
			{
				(buf)[int32(3)] = uint8_t((uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(3)) & uint32_t((uint32(int32(255))))))))))
				(buf)[int32(2)] = uint8_t((uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(2)) & uint32_t((uint32(int32(255))))))))))
				(buf)[int32(1)] = uint8_t((uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(1)) & uint32_t((uint32(int32(255))))))))))
				(buf)[int32(0)] = uint8_t((uint8((uint32((c[0].sbuf >> uint64(int32(8)*int32(0)) & uint32_t((uint32(int32(255))))))))))
			}
			nbytes -= int32(4)
			buf = (*(*[1000000000]uint8_t)(unsafe.Pointer(uintptr(unsafe.Pointer(&buf[0])) + (uintptr)(int32(4))*unsafe.Sizeof(buf[0]))))[:]
		} else {
			for i = int32(0); i < nbytes; i++ {
				buf[i] = uint8_t((uint8((uint32((c[0].sbuf >> uint64(int32(8)*i) & uint32_t((uint32(int32(255))))))))))
			}
			break
		}
	}
}
func init() {
}
