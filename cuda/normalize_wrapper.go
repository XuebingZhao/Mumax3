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

// CUDA handle for normalize kernel
var normalize_code cu.Function

// Stores the arguments for normalize kernel invocation
type normalize_args_t struct {
	arg_vx  unsafe.Pointer
	arg_vy  unsafe.Pointer
	arg_vz  unsafe.Pointer
	arg_vol unsafe.Pointer
	arg_N   int
	argptr  [5]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for normalize kernel invocation
var normalize_args normalize_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	normalize_args.argptr[0] = unsafe.Pointer(&normalize_args.arg_vx)
	normalize_args.argptr[1] = unsafe.Pointer(&normalize_args.arg_vy)
	normalize_args.argptr[2] = unsafe.Pointer(&normalize_args.arg_vz)
	normalize_args.argptr[3] = unsafe.Pointer(&normalize_args.arg_vol)
	normalize_args.argptr[4] = unsafe.Pointer(&normalize_args.arg_N)
}

// Wrapper for normalize CUDA kernel, asynchronous.
func k_normalize_async(vx unsafe.Pointer, vy unsafe.Pointer, vz unsafe.Pointer, vol unsafe.Pointer, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	normalize_args.Lock()
	defer normalize_args.Unlock()

	if normalize_code == 0 {
		normalize_code = fatbinLoad(normalize_map, "normalize")
	}

	normalize_args.arg_vx = vx
	normalize_args.arg_vy = vy
	normalize_args.arg_vz = vz
	normalize_args.arg_vol = vol
	normalize_args.arg_N = N

	args := normalize_args.argptr[:]
	cu.LaunchKernel(normalize_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for normalize kernel.
var normalize_map = map[int]string{0: "",
	20: normalize_ptx_20,
	30: normalize_ptx_30,
	35: normalize_ptx_35}

// normalize PTX code for various compute capabilities.
const (
	normalize_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u64 normalize_param_3,
	.param .u32 normalize_param_4
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<22>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd9, [normalize_param_0];
	ld.param.u64 	%rd10, [normalize_param_1];
	ld.param.u64 	%rd11, [normalize_param_2];
	ld.param.u64 	%rd8, [normalize_param_3];
	ld.param.u32 	%r2, [normalize_param_4];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd8;
	.loc 1 7 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 8 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_8;

	.loc 1 10 1
	setp.ne.s64	%p2, %rd8, 0;
	@%p2 bra 	BB0_3;

	mov.f32 	%f20, 0f3F800000;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd4, %rd12;
	.loc 1 10 1
	ld.global.f32 	%f20, [%rd13];

BB0_4:
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd5, %rd3, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f10, [%rd5];
	mul.f32 	%f3, %f20, %f10;
	add.s64 	%rd6, %rd2, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f11, [%rd6];
	mul.f32 	%f4, %f20, %f11;
	add.s64 	%rd7, %rd1, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f12, [%rd7];
	mul.f32 	%f5, %f20, %f12;
	.loc 1 12 1
	mul.f32 	%f13, %f4, %f4;
	fma.rn.f32 	%f14, %f3, %f3, %f13;
	fma.rn.f32 	%f15, %f5, %f5, %f14;
	.loc 2 3055 10
	sqrt.rn.f32 	%f6, %f15;
	.loc 1 12 125
	setp.neu.f32	%p3, %f6, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f21, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f21, %f6;

BB0_7:
	mul.f32 	%f17, %f21, %f3;
	.loc 1 13 1
	st.global.f32 	[%rd5], %f17;
	mul.f32 	%f18, %f21, %f4;
	.loc 1 14 1
	st.global.f32 	[%rd6], %f18;
	mul.f32 	%f19, %f21, %f5;
	.loc 1 15 1
	st.global.f32 	[%rd7], %f19;

BB0_8:
	.loc 1 17 2
	ret;
}


`
	normalize_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u64 normalize_param_3,
	.param .u32 normalize_param_4
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<22>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd9, [normalize_param_0];
	ld.param.u64 	%rd10, [normalize_param_1];
	ld.param.u64 	%rd11, [normalize_param_2];
	ld.param.u64 	%rd8, [normalize_param_3];
	ld.param.u32 	%r2, [normalize_param_4];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd8;
	.loc 1 7 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 8 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_8;

	.loc 1 10 1
	setp.ne.s64	%p2, %rd8, 0;
	@%p2 bra 	BB0_3;

	mov.f32 	%f20, 0f3F800000;
	bra.uni 	BB0_4;

BB0_3:
	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd4, %rd12;
	.loc 1 10 1
	ld.global.f32 	%f20, [%rd13];

BB0_4:
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd5, %rd3, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f10, [%rd5];
	mul.f32 	%f3, %f20, %f10;
	add.s64 	%rd6, %rd2, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f11, [%rd6];
	mul.f32 	%f4, %f20, %f11;
	add.s64 	%rd7, %rd1, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f12, [%rd7];
	mul.f32 	%f5, %f20, %f12;
	.loc 1 12 1
	mul.f32 	%f13, %f4, %f4;
	fma.rn.f32 	%f14, %f3, %f3, %f13;
	fma.rn.f32 	%f15, %f5, %f5, %f14;
	.loc 2 3055 10
	sqrt.rn.f32 	%f6, %f15;
	.loc 1 12 125
	setp.neu.f32	%p3, %f6, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f21, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f21, %f6;

BB0_7:
	mul.f32 	%f17, %f21, %f3;
	.loc 1 13 1
	st.global.f32 	[%rd5], %f17;
	mul.f32 	%f18, %f21, %f4;
	.loc 1 14 1
	st.global.f32 	[%rd6], %f18;
	mul.f32 	%f19, %f21, %f5;
	.loc 1 15 1
	st.global.f32 	[%rd7], %f19;

BB0_8:
	.loc 1 17 2
	ret;
}


`
	normalize_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


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

.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u64 normalize_param_3,
	.param .u32 normalize_param_4
)
{
	.reg .pred 	%p<4>;
	.reg .s32 	%r<9>;
	.reg .f32 	%f<22>;
	.reg .s64 	%rd<15>;


	ld.param.u64 	%rd9, [normalize_param_0];
	ld.param.u64 	%rd10, [normalize_param_1];
	ld.param.u64 	%rd11, [normalize_param_2];
	ld.param.u64 	%rd8, [normalize_param_3];
	ld.param.u32 	%r2, [normalize_param_4];
	cvta.to.global.u64 	%rd1, %rd11;
	cvta.to.global.u64 	%rd2, %rd10;
	cvta.to.global.u64 	%rd3, %rd9;
	cvta.to.global.u64 	%rd4, %rd8;
	.loc 1 7 1
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	.loc 1 8 1
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB2_8;

	.loc 1 10 1
	setp.ne.s64	%p2, %rd8, 0;
	@%p2 bra 	BB2_3;

	mov.f32 	%f20, 0f3F800000;
	bra.uni 	BB2_4;

BB2_3:
	mul.wide.s32 	%rd12, %r1, 4;
	add.s64 	%rd13, %rd4, %rd12;
	.loc 1 10 1
	ld.global.nc.f32 	%f20, [%rd13];

BB2_4:
	mul.wide.s32 	%rd14, %r1, 4;
	add.s64 	%rd5, %rd3, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f10, [%rd5];
	mul.f32 	%f3, %f20, %f10;
	add.s64 	%rd6, %rd2, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f11, [%rd6];
	mul.f32 	%f4, %f20, %f11;
	add.s64 	%rd7, %rd1, %rd14;
	.loc 1 11 1
	ld.global.f32 	%f12, [%rd7];
	mul.f32 	%f5, %f20, %f12;
	.loc 1 12 1
	mul.f32 	%f13, %f4, %f4;
	fma.rn.f32 	%f14, %f3, %f3, %f13;
	fma.rn.f32 	%f15, %f5, %f5, %f14;
	.loc 3 3055 10
	sqrt.rn.f32 	%f6, %f15;
	.loc 1 12 125
	setp.neu.f32	%p3, %f6, 0f00000000;
	@%p3 bra 	BB2_6;

	mov.f32 	%f21, 0f00000000;
	bra.uni 	BB2_7;

BB2_6:
	rcp.rn.f32 	%f21, %f6;

BB2_7:
	mul.f32 	%f17, %f21, %f3;
	.loc 1 13 1
	st.global.f32 	[%rd5], %f17;
	mul.f32 	%f18, %f21, %f4;
	.loc 1 14 1
	st.global.f32 	[%rd6], %f18;
	mul.f32 	%f19, %f21, %f5;
	.loc 1 15 1
	st.global.f32 	[%rd7], %f19;

BB2_8:
	.loc 1 17 2
	ret;
}


`
)
