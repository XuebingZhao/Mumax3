package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for reducemaxvecdiff2 kernel
var reducemaxvecdiff2_code cu.Function

// Stores the arguments for reducemaxvecdiff2 kernel invocation
type reducemaxvecdiff2_args_t struct {
	arg_x1      unsafe.Pointer
	arg_y1      unsafe.Pointer
	arg_z1      unsafe.Pointer
	arg_x2      unsafe.Pointer
	arg_y2      unsafe.Pointer
	arg_z2      unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [9]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducemaxvecdiff2 kernel invocation
var reducemaxvecdiff2_args reducemaxvecdiff2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducemaxvecdiff2_args.argptr[0] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_x1)
	reducemaxvecdiff2_args.argptr[1] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_y1)
	reducemaxvecdiff2_args.argptr[2] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_z1)
	reducemaxvecdiff2_args.argptr[3] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_x2)
	reducemaxvecdiff2_args.argptr[4] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_y2)
	reducemaxvecdiff2_args.argptr[5] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_z2)
	reducemaxvecdiff2_args.argptr[6] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_dst)
	reducemaxvecdiff2_args.argptr[7] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_initVal)
	reducemaxvecdiff2_args.argptr[8] = unsafe.Pointer(&reducemaxvecdiff2_args.arg_n)
}

// Wrapper for reducemaxvecdiff2 CUDA kernel, asynchronous.
func k_reducemaxvecdiff2_async(x1 unsafe.Pointer, y1 unsafe.Pointer, z1 unsafe.Pointer, x2 unsafe.Pointer, y2 unsafe.Pointer, z2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	reducemaxvecdiff2_args.Lock()
	defer reducemaxvecdiff2_args.Unlock()

	if reducemaxvecdiff2_code == 0 {
		reducemaxvecdiff2_code = fatbinLoad(reducemaxvecdiff2_map, "reducemaxvecdiff2")
	}

	reducemaxvecdiff2_args.arg_x1 = x1
	reducemaxvecdiff2_args.arg_y1 = y1
	reducemaxvecdiff2_args.arg_z1 = z1
	reducemaxvecdiff2_args.arg_x2 = x2
	reducemaxvecdiff2_args.arg_y2 = y2
	reducemaxvecdiff2_args.arg_z2 = z2
	reducemaxvecdiff2_args.arg_dst = dst
	reducemaxvecdiff2_args.arg_initVal = initVal
	reducemaxvecdiff2_args.arg_n = n

	args := reducemaxvecdiff2_args.argptr[:]
	cu.LaunchKernel(reducemaxvecdiff2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for reducemaxvecdiff2 kernel.
var reducemaxvecdiff2_map = map[int]string{0: "",
	20: reducemaxvecdiff2_ptx_20,
	30: reducemaxvecdiff2_ptx_30,
	35: reducemaxvecdiff2_ptx_35}

// reducemaxvecdiff2 PTX code for various compute capabilities.
const (
	reducemaxvecdiff2_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.visible .entry reducemaxvecdiff2(
	.param .u64 reducemaxvecdiff2_param_0,
	.param .u64 reducemaxvecdiff2_param_1,
	.param .u64 reducemaxvecdiff2_param_2,
	.param .u64 reducemaxvecdiff2_param_3,
	.param .u64 reducemaxvecdiff2_param_4,
	.param .u64 reducemaxvecdiff2_param_5,
	.param .u64 reducemaxvecdiff2_param_6,
	.param .f32 reducemaxvecdiff2_param_7,
	.param .u32 reducemaxvecdiff2_param_8
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<28>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecdiff2$__cuda_local_var_33837_35_non_const_sdata[2048];

	ld.param.u64 	%rd9, [reducemaxvecdiff2_param_0];
	ld.param.u64 	%rd10, [reducemaxvecdiff2_param_1];
	ld.param.u64 	%rd11, [reducemaxvecdiff2_param_2];
	ld.param.u64 	%rd12, [reducemaxvecdiff2_param_3];
	ld.param.u64 	%rd13, [reducemaxvecdiff2_param_4];
	ld.param.u64 	%rd14, [reducemaxvecdiff2_param_5];
	ld.param.u64 	%rd15, [reducemaxvecdiff2_param_6];
	ld.param.f32 	%f40, [reducemaxvecdiff2_param_7];
	ld.param.u32 	%r9, [reducemaxvecdiff2_param_8];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd11;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd10;
	cvta.to.global.u64 	%rd6, %rd12;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 14 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 14 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 14 1
	mul.wide.s32 	%rd16, %r15, 4;
	add.s64 	%rd17, %rd7, %rd16;
	add.s64 	%rd18, %rd6, %rd16;
	ld.global.f32 	%f5, [%rd18];
	ld.global.f32 	%f6, [%rd17];
	sub.f32 	%f7, %f6, %f5;
	add.s64 	%rd19, %rd5, %rd16;
	add.s64 	%rd20, %rd4, %rd16;
	ld.global.f32 	%f8, [%rd20];
	ld.global.f32 	%f9, [%rd19];
	sub.f32 	%f10, %f9, %f8;
	mul.f32 	%f11, %f10, %f10;
	fma.rn.f32 	%f12, %f7, %f7, %f11;
	add.s64 	%rd21, %rd3, %rd16;
	add.s64 	%rd22, %rd2, %rd16;
	ld.global.f32 	%f13, [%rd22];
	ld.global.f32 	%f14, [%rd21];
	sub.f32 	%f15, %f14, %f13;
	fma.rn.f32 	%f16, %f15, %f15, %f12;
	.loc 2 2770 10
	max.f32 	%f40, %f40, %f16;
	.loc 1 14 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 14 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 14 1
	mul.wide.s32 	%rd23, %r2, 4;
	mov.u64 	%rd24, reducemaxvecdiff2$__cuda_local_var_33837_35_non_const_sdata;
	add.s64 	%rd8, %rd24, %rd23;
	st.shared.f32 	[%rd8], %f40;
	bar.sync 	0;
	.loc 1 14 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 14 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 14 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	.loc 1 14 1
	ld.shared.f32 	%f17, [%rd8];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd25, %r12, 4;
	add.s64 	%rd27, %rd24, %rd25;
	ld.shared.f32 	%f18, [%rd27];
	.loc 2 2770 10
	max.f32 	%f19, %f17, %f18;
	st.shared.f32 	[%rd8], %f19;

BB0_5:
	.loc 1 14 1
	bar.sync 	0;
	.loc 1 14 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 14 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 14 1
	ld.volatile.shared.f32 	%f20, [%rd8];
	ld.volatile.shared.f32 	%f21, [%rd8+128];
	.loc 2 2770 10
	max.f32 	%f22, %f20, %f21;
	st.volatile.shared.f32 	[%rd8], %f22;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f23, [%rd8+64];
	ld.volatile.shared.f32 	%f24, [%rd8];
	.loc 2 2770 10
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd8], %f25;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f26, [%rd8+32];
	ld.volatile.shared.f32 	%f27, [%rd8];
	.loc 2 2770 10
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd8], %f28;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f29, [%rd8+16];
	ld.volatile.shared.f32 	%f30, [%rd8];
	.loc 2 2770 10
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd8], %f31;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f32, [%rd8+8];
	ld.volatile.shared.f32 	%f33, [%rd8];
	.loc 2 2770 10
	max.f32 	%f34, %f33, %f32;
	st.volatile.shared.f32 	[%rd8], %f34;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f35, [%rd8+4];
	ld.volatile.shared.f32 	%f36, [%rd8];
	.loc 2 2770 10
	max.f32 	%f37, %f36, %f35;
	st.volatile.shared.f32 	[%rd8], %f37;

BB0_8:
	.loc 1 14 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 14 1
	ld.shared.f32 	%f38, [reducemaxvecdiff2$__cuda_local_var_33837_35_non_const_sdata];
	.loc 2 2750 10
	abs.f32 	%f39, %f38;
	.loc 1 14 37
	mov.b32 	 %r13, %f39;
	.loc 2 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB0_10:
	.loc 1 15 2
	ret;
}


`
	reducemaxvecdiff2_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.visible .entry reducemaxvecdiff2(
	.param .u64 reducemaxvecdiff2_param_0,
	.param .u64 reducemaxvecdiff2_param_1,
	.param .u64 reducemaxvecdiff2_param_2,
	.param .u64 reducemaxvecdiff2_param_3,
	.param .u64 reducemaxvecdiff2_param_4,
	.param .u64 reducemaxvecdiff2_param_5,
	.param .u64 reducemaxvecdiff2_param_6,
	.param .f32 reducemaxvecdiff2_param_7,
	.param .u32 reducemaxvecdiff2_param_8
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<28>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecdiff2$__cuda_local_var_33910_35_non_const_sdata[2048];

	ld.param.u64 	%rd9, [reducemaxvecdiff2_param_0];
	ld.param.u64 	%rd10, [reducemaxvecdiff2_param_1];
	ld.param.u64 	%rd11, [reducemaxvecdiff2_param_2];
	ld.param.u64 	%rd12, [reducemaxvecdiff2_param_3];
	ld.param.u64 	%rd13, [reducemaxvecdiff2_param_4];
	ld.param.u64 	%rd14, [reducemaxvecdiff2_param_5];
	ld.param.u64 	%rd15, [reducemaxvecdiff2_param_6];
	ld.param.f32 	%f40, [reducemaxvecdiff2_param_7];
	ld.param.u32 	%r9, [reducemaxvecdiff2_param_8];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd11;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd10;
	cvta.to.global.u64 	%rd6, %rd12;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 14 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 14 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 14 1
	mul.wide.s32 	%rd16, %r15, 4;
	add.s64 	%rd17, %rd7, %rd16;
	add.s64 	%rd18, %rd6, %rd16;
	ld.global.f32 	%f5, [%rd18];
	ld.global.f32 	%f6, [%rd17];
	sub.f32 	%f7, %f6, %f5;
	add.s64 	%rd19, %rd5, %rd16;
	add.s64 	%rd20, %rd4, %rd16;
	ld.global.f32 	%f8, [%rd20];
	ld.global.f32 	%f9, [%rd19];
	sub.f32 	%f10, %f9, %f8;
	mul.f32 	%f11, %f10, %f10;
	fma.rn.f32 	%f12, %f7, %f7, %f11;
	add.s64 	%rd21, %rd3, %rd16;
	add.s64 	%rd22, %rd2, %rd16;
	ld.global.f32 	%f13, [%rd22];
	ld.global.f32 	%f14, [%rd21];
	sub.f32 	%f15, %f14, %f13;
	fma.rn.f32 	%f16, %f15, %f15, %f12;
	.loc 2 2770 10
	max.f32 	%f40, %f40, %f16;
	.loc 1 14 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 14 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 14 1
	mul.wide.s32 	%rd23, %r2, 4;
	mov.u64 	%rd24, reducemaxvecdiff2$__cuda_local_var_33910_35_non_const_sdata;
	add.s64 	%rd8, %rd24, %rd23;
	st.shared.f32 	[%rd8], %f40;
	bar.sync 	0;
	.loc 1 14 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 14 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 14 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB0_5;

	.loc 1 14 1
	ld.shared.f32 	%f17, [%rd8];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd25, %r12, 4;
	add.s64 	%rd27, %rd24, %rd25;
	ld.shared.f32 	%f18, [%rd27];
	.loc 2 2770 10
	max.f32 	%f19, %f17, %f18;
	st.shared.f32 	[%rd8], %f19;

BB0_5:
	.loc 1 14 1
	bar.sync 	0;
	.loc 1 14 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 14 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 14 1
	ld.volatile.shared.f32 	%f20, [%rd8];
	ld.volatile.shared.f32 	%f21, [%rd8+128];
	.loc 2 2770 10
	max.f32 	%f22, %f20, %f21;
	st.volatile.shared.f32 	[%rd8], %f22;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f23, [%rd8+64];
	ld.volatile.shared.f32 	%f24, [%rd8];
	.loc 2 2770 10
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd8], %f25;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f26, [%rd8+32];
	ld.volatile.shared.f32 	%f27, [%rd8];
	.loc 2 2770 10
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd8], %f28;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f29, [%rd8+16];
	ld.volatile.shared.f32 	%f30, [%rd8];
	.loc 2 2770 10
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd8], %f31;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f32, [%rd8+8];
	ld.volatile.shared.f32 	%f33, [%rd8];
	.loc 2 2770 10
	max.f32 	%f34, %f33, %f32;
	st.volatile.shared.f32 	[%rd8], %f34;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f35, [%rd8+4];
	ld.volatile.shared.f32 	%f36, [%rd8];
	.loc 2 2770 10
	max.f32 	%f37, %f36, %f35;
	st.volatile.shared.f32 	[%rd8], %f37;

BB0_8:
	.loc 1 14 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 14 1
	ld.shared.f32 	%f38, [reducemaxvecdiff2$__cuda_local_var_33910_35_non_const_sdata];
	.loc 2 2750 10
	abs.f32 	%f39, %f38;
	.loc 1 14 37
	mov.b32 	 %r13, %f39;
	.loc 2 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB0_10:
	.loc 1 15 2
	ret;
}


`
	reducemaxvecdiff2_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64

.global .align 1 .b8 $str[11] = {95, 95, 67, 85, 68, 65, 95, 70, 84, 90, 0};

.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry reducemaxvecdiff2(
	.param .u64 reducemaxvecdiff2_param_0,
	.param .u64 reducemaxvecdiff2_param_1,
	.param .u64 reducemaxvecdiff2_param_2,
	.param .u64 reducemaxvecdiff2_param_3,
	.param .u64 reducemaxvecdiff2_param_4,
	.param .u64 reducemaxvecdiff2_param_5,
	.param .u64 reducemaxvecdiff2_param_6,
	.param .f32 reducemaxvecdiff2_param_7,
	.param .u32 reducemaxvecdiff2_param_8
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<17>;
	.reg .f32 	%f<41>;
	.reg .s64 	%rd<28>;
	// demoted variable
	.shared .align 4 .b8 reducemaxvecdiff2$__cuda_local_var_34073_35_non_const_sdata[2048];

	ld.param.u64 	%rd9, [reducemaxvecdiff2_param_0];
	ld.param.u64 	%rd10, [reducemaxvecdiff2_param_1];
	ld.param.u64 	%rd11, [reducemaxvecdiff2_param_2];
	ld.param.u64 	%rd12, [reducemaxvecdiff2_param_3];
	ld.param.u64 	%rd13, [reducemaxvecdiff2_param_4];
	ld.param.u64 	%rd14, [reducemaxvecdiff2_param_5];
	ld.param.u64 	%rd15, [reducemaxvecdiff2_param_6];
	ld.param.f32 	%f40, [reducemaxvecdiff2_param_7];
	ld.param.u32 	%r9, [reducemaxvecdiff2_param_8];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd14;
	cvta.to.global.u64 	%rd3, %rd11;
	cvta.to.global.u64 	%rd4, %rd13;
	cvta.to.global.u64 	%rd5, %rd10;
	cvta.to.global.u64 	%rd6, %rd12;
	cvta.to.global.u64 	%rd7, %rd9;
	.loc 1 14 1
	mov.u32 	%r16, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r15, %r16, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r16, %r11;
	.loc 1 14 1
	setp.ge.s32	%p1, %r15, %r9;
	@%p1 bra 	BB2_2;

BB2_1:
	.loc 1 14 1
	mul.wide.s32 	%rd16, %r15, 4;
	add.s64 	%rd17, %rd7, %rd16;
	add.s64 	%rd18, %rd6, %rd16;
	ld.global.f32 	%f5, [%rd18];
	ld.global.f32 	%f6, [%rd17];
	sub.f32 	%f7, %f6, %f5;
	add.s64 	%rd19, %rd5, %rd16;
	add.s64 	%rd20, %rd4, %rd16;
	ld.global.f32 	%f8, [%rd20];
	ld.global.f32 	%f9, [%rd19];
	sub.f32 	%f10, %f9, %f8;
	mul.f32 	%f11, %f10, %f10;
	fma.rn.f32 	%f12, %f7, %f7, %f11;
	add.s64 	%rd21, %rd3, %rd16;
	add.s64 	%rd22, %rd2, %rd16;
	ld.global.f32 	%f13, [%rd22];
	ld.global.f32 	%f14, [%rd21];
	sub.f32 	%f15, %f14, %f13;
	fma.rn.f32 	%f16, %f15, %f15, %f12;
	.loc 3 2770 10
	max.f32 	%f40, %f40, %f16;
	.loc 1 14 1
	add.s32 	%r15, %r15, %r4;
	.loc 1 14 1
	setp.lt.s32	%p2, %r15, %r9;
	@%p2 bra 	BB2_1;

BB2_2:
	.loc 1 14 1
	mul.wide.s32 	%rd23, %r2, 4;
	mov.u64 	%rd24, reducemaxvecdiff2$__cuda_local_var_34073_35_non_const_sdata;
	add.s64 	%rd8, %rd24, %rd23;
	st.shared.f32 	[%rd8], %f40;
	bar.sync 	0;
	.loc 1 14 1
	setp.lt.u32	%p3, %r16, 66;
	@%p3 bra 	BB2_6;

BB2_3:
	.loc 1 14 1
	mov.u32 	%r7, %r16;
	shr.u32 	%r16, %r7, 1;
	.loc 1 14 1
	setp.ge.u32	%p4, %r2, %r16;
	@%p4 bra 	BB2_5;

	.loc 1 14 1
	ld.shared.f32 	%f17, [%rd8];
	add.s32 	%r12, %r16, %r2;
	mul.wide.u32 	%rd25, %r12, 4;
	add.s64 	%rd27, %rd24, %rd25;
	ld.shared.f32 	%f18, [%rd27];
	.loc 3 2770 10
	max.f32 	%f19, %f17, %f18;
	st.shared.f32 	[%rd8], %f19;

BB2_5:
	.loc 1 14 1
	bar.sync 	0;
	.loc 1 14 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB2_3;

BB2_6:
	.loc 1 14 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB2_8;

	.loc 1 14 1
	ld.volatile.shared.f32 	%f20, [%rd8];
	ld.volatile.shared.f32 	%f21, [%rd8+128];
	.loc 3 2770 10
	max.f32 	%f22, %f20, %f21;
	st.volatile.shared.f32 	[%rd8], %f22;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f23, [%rd8+64];
	ld.volatile.shared.f32 	%f24, [%rd8];
	.loc 3 2770 10
	max.f32 	%f25, %f24, %f23;
	st.volatile.shared.f32 	[%rd8], %f25;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f26, [%rd8+32];
	ld.volatile.shared.f32 	%f27, [%rd8];
	.loc 3 2770 10
	max.f32 	%f28, %f27, %f26;
	st.volatile.shared.f32 	[%rd8], %f28;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f29, [%rd8+16];
	ld.volatile.shared.f32 	%f30, [%rd8];
	.loc 3 2770 10
	max.f32 	%f31, %f30, %f29;
	st.volatile.shared.f32 	[%rd8], %f31;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f32, [%rd8+8];
	ld.volatile.shared.f32 	%f33, [%rd8];
	.loc 3 2770 10
	max.f32 	%f34, %f33, %f32;
	st.volatile.shared.f32 	[%rd8], %f34;
	.loc 1 14 1
	ld.volatile.shared.f32 	%f35, [%rd8+4];
	ld.volatile.shared.f32 	%f36, [%rd8];
	.loc 3 2770 10
	max.f32 	%f37, %f36, %f35;
	st.volatile.shared.f32 	[%rd8], %f37;

BB2_8:
	.loc 1 14 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB2_10;

	.loc 1 14 1
	ld.shared.f32 	%f38, [reducemaxvecdiff2$__cuda_local_var_34073_35_non_const_sdata];
	.loc 3 2750 10
	abs.f32 	%f39, %f38;
	.loc 1 14 37
	mov.b32 	 %r13, %f39;
	.loc 3 3781 3
	atom.global.max.s32 	%r14, [%rd1], %r13;

BB2_10:
	.loc 1 15 2
	ret;
}


`
)
