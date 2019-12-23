// AUTOGENERATED FILE
// +build !codeanalysis
// +build !syz_target syz_target,syz_os_test,syz_arch_32_shmem

package gen

import . "github.com/google/syzkaller/prog"
import . "github.com/google/syzkaller/sys/test"

func init() {
	RegisterTarget(&Target{OS: "test", Arch: "32_shmem", Revision: revision_32_shmem, PtrSize: 4, PageSize: 8192, NumPages: 2048, DataOffset: 536870912, Syscalls: syscalls_32_shmem, Resources: resources_32_shmem, Structs: structDescs_32_shmem, Consts: consts_32_shmem}, InitTarget)
}

var resources_32_shmem = []*ResourceDesc(nil)

var structDescs_32_shmem = []*KeyedStruct{
	{Key: StructKey{Name: "align0"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "align0", TypeSize: 20}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f0", TypeSize: 2}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f3", TypeSize: 2}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f4", TypeSize: 8}}},
	}}},
	{Key: StructKey{Name: "compare_data"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "compare_data", IsVarlen: true}, Fields: []Type{
		&StructType{Key: StructKey{Name: "align0"}, FldName: "align0"},
		&StructType{Key: StructKey{Name: "syz_bf_struct0"}, FldName: "bf0"},
		&StructType{Key: StructKey{Name: "syz_bf_struct1"}, FldName: "bf1"},
		&StructType{Key: StructKey{Name: "syz_bf_struct2"}, FldName: "bf2"},
		&StructType{Key: StructKey{Name: "syz_bf_struct3"}, FldName: "bf3"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct4]"}, FldName: "bf4"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct5]"}, FldName: "bf5"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct6]"}, FldName: "bf6"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct7]"}, FldName: "bf7"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct8]"}, FldName: "bf8"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct9]"}, FldName: "bf9"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct10]"}, FldName: "bf10"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct11]"}, FldName: "bf11"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct12]"}, FldName: "bf12"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct13]"}, FldName: "bf13"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct14]"}, FldName: "bf14"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct15]"}, FldName: "bf15"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct16]"}, FldName: "bf16"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct17]"}, FldName: "bf17"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct18]"}, FldName: "bf18"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct19]"}, FldName: "bf19"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct20]"}, FldName: "bf20"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct21]"}, FldName: "bf21"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct22]"}, FldName: "bf22"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct23]"}, FldName: "bf23"},
		&StructType{Key: StructKey{Name: "syz_bf_align[syz_bf_struct24]"}, FldName: "bf24"},
		&BufferType{TypeCommon: TypeCommon{TypeName: "string", FldName: "str", IsVarlen: true}, Kind: 2},
		&BufferType{TypeCommon: TypeCommon{TypeName: "array", FldName: "blob", IsVarlen: true}},
		&ArrayType{TypeCommon: TypeCommon{TypeName: "array", FldName: "arr16be", IsVarlen: true}, Type: &IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16be", TypeSize: 2}, ArgFormat: 1}}},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct10]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct10]", TypeSize: 16}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct10"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct11]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct11]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct11"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct12]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct12]", TypeSize: 12}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct12"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct13]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct13]", TypeSize: 12}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct13"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct14]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct14]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct14"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct15]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct15]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct15"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct16]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct16]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct16"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct17]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct17]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct17"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct18]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct18]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct18"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct19]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct19]", TypeSize: 6}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct19"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct20]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct20]", TypeSize: 12}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct20"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct21]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct21]", TypeSize: 12}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct21"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct22]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct22]", TypeSize: 3}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&StructType{Key: StructKey{Name: "syz_bf_struct22"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct23]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct23]", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&StructType{Key: StructKey{Name: "syz_bf_struct23"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct24]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct24]", TypeSize: 3}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&StructType{Key: StructKey{Name: "syz_bf_struct24"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct4]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct4]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct4"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct5]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct5]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct5"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct6]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct6]", TypeSize: 6}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct6"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct7]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct7]", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct7"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct8]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct8]", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct8"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_align[syz_bf_struct9]"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_align[syz_bf_struct9]", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&StructType{Key: StructKey{Name: "syz_bf_struct9"}, FldName: "f1"},
	}}},
	{Key: StructKey{Name: "syz_bf_struct0"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct0", TypeSize: 24}, Fields: []Type{
		&FlagsType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "syz_bf_flags", FldName: "f0", TypeSize: 4}, BitfieldLen: 10, BitfieldUnit: 2}, Vals: []uint64{0, 1, 2}, BitMask: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1", TypeSize: 8}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "f2"}, BitfieldLen: 5, BitfieldUnit: 2}, Val: 66},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f3"}, BitfieldOff: 5, BitfieldLen: 6, BitfieldUnit: 2}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "f4", TypeSize: 4}, BitfieldOff: 11, BitfieldLen: 15, BitfieldUnit: 4}, Val: 66},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "f5", TypeSize: 2}, BitfieldLen: 11, BitfieldUnit: 2}, Path: []string{"parent"}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "f6", TypeSize: 2}, ArgFormat: 1, BitfieldLen: 11, BitfieldUnit: 2}, Path: []string{"parent"}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f7", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct1"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct1", TypeSize: 8}, Fields: []Type{
		&StructType{Key: StructKey{Name: "syz_bf_struct1_internal"}, FldName: "f0"},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct10"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct10", TypeSize: 12}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f0"}, BitfieldLen: 4, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1"}, BitfieldOff: 4, BitfieldLen: 4, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f2"}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f3"}, BitfieldOff: 12, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f4"}, BitfieldOff: 24, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f5"}, BitfieldOff: 36, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f6", TypeSize: 8}, BitfieldOff: 48, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f7", TypeSize: 2}, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f8", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct11"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct11", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f3"}, BitfieldOff: 24, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 3}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f4", TypeSize: 1}, BitfieldOff: 28, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 3}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct12"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct12", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f3"}, BitfieldOff: 24, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 3}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f4", TypeSize: 1}, BitfieldOff: 28, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 3}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f5", TypeSize: 1}, BitfieldLen: 4, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f6", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct13"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct13", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f0", TypeSize: 2}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 2}, BitfieldOff: 16, BitfieldLen: 12, BitfieldUnit: 4, BitfieldUnitOff: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f2", TypeSize: 2}, BitfieldLen: 12, BitfieldUnit: 4}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct14"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct14", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f2", TypeSize: 2}, BitfieldOff: 16, BitfieldLen: 12, BitfieldUnit: 4, BitfieldUnitOff: 2}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct15"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct15", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 12, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 2}, BitfieldOff: 4, BitfieldLen: 4, BitfieldUnit: 2}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct16"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct16", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1"}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldOff: 12, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct17"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct17", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 6, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldLen: 4, BitfieldUnit: 2}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct18"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct18", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldLen: 6, BitfieldUnit: 2}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct19"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct19", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f1", TypeSize: 2}, BitfieldLen: 12, BitfieldUnit: 2}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct1_internal"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct1_internal", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f0"}, BitfieldLen: 10, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1"}, BitfieldOff: 10, BitfieldLen: 10, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f2", TypeSize: 4}, BitfieldOff: 20, BitfieldLen: 10, BitfieldUnit: 4}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct2"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct2", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f0"}, BitfieldLen: 4, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1"}, BitfieldOff: 4, BitfieldLen: 8, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f2"}, BitfieldOff: 12, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f3"}, BitfieldOff: 24, BitfieldLen: 20, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f4", TypeSize: 8}, BitfieldOff: 44, BitfieldLen: 16, BitfieldUnit: 8}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct20"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct20", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f2"}, BitfieldOff: 16, BitfieldLen: 4, BitfieldUnit: 8, BitfieldUnitOff: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f3"}, BitfieldOff: 20, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f4", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 2}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f5", TypeSize: 2}, BitfieldOff: 4, BitfieldLen: 4, BitfieldUnit: 1}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 3}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct21"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct21", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f0", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f1", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 8, BitfieldUnit: 2, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f3", TypeSize: 2}}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct22"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct22", TypeSize: 2}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1"}, BitfieldLen: 4, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldOff: 4, BitfieldLen: 4, BitfieldUnit: 2}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct23"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct23", TypeSize: 3}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1"}, BitfieldLen: 4, BitfieldUnit: 4}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 2}, BitfieldOff: 4, BitfieldLen: 6, BitfieldUnit: 2}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct24"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct24", TypeSize: 2}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64", FldName: "f1", TypeSize: 1}, BitfieldLen: 4, BitfieldUnit: 8}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct3"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct3", TypeSize: 8}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f0"}, ArgFormat: 1, BitfieldLen: 4, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f1"}, ArgFormat: 1, BitfieldOff: 4, BitfieldLen: 8, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f2"}, ArgFormat: 1, BitfieldOff: 12, BitfieldLen: 12, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f3"}, ArgFormat: 1, BitfieldOff: 24, BitfieldLen: 20, BitfieldUnit: 8}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int64be", FldName: "f4", TypeSize: 8}, ArgFormat: 1, BitfieldOff: 44, BitfieldLen: 16, BitfieldUnit: 8}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct4"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct4", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct5"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct5", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct6"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct6", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f2", TypeSize: 1}}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 1}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct7"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct7", TypeSize: 2}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f1", TypeSize: 1}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
	}}},
	{Key: StructKey{Name: "syz_bf_struct8"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct8", TypeSize: 4}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "f1"}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 4, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldOff: 12, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "pad", TypeSize: 2}}, IsPad: true},
	}}},
	{Key: StructKey{Name: "syz_bf_struct9"}, Desc: &StructDesc{TypeCommon: TypeCommon{TypeName: "syz_bf_struct9", TypeSize: 2}, Fields: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int8", FldName: "f0", TypeSize: 1}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f1"}, BitfieldOff: 8, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int16", FldName: "f2", TypeSize: 1}, BitfieldOff: 12, BitfieldLen: 4, BitfieldUnit: 2, BitfieldUnitOff: 1}},
	}}},
}

var syscalls_32_shmem = []*Syscall{
	{Name: "syz_compare", CallName: "syz_compare", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "want", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "string", IsVarlen: true}, Kind: 2}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "bytesize", FldName: "want_len", TypeSize: 4}}, BitSize: 8, Path: []string{"want"}},
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "got", TypeSize: 4}, Type: &UnionType{Key: StructKey{Name: "compare_data"}}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "bytesize", FldName: "got_len", TypeSize: 4}}, BitSize: 8, Path: []string{"got"}},
	}},
	{Name: "syz_compare_int$2", CallName: "syz_compare_int", MissingArgs: 2, Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 4}}, Val: 2},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 4}}},
	}},
	{Name: "syz_compare_int$3", CallName: "syz_compare_int", MissingArgs: 1, Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 4}}, Val: 3},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v2", TypeSize: 4}}},
	}},
	{Name: "syz_compare_int$4", CallName: "syz_compare_int", Args: []Type{
		&ConstType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "const", FldName: "n", TypeSize: 4}}, Val: 4},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v0", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v1", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v2", TypeSize: 4}}},
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "intptr", FldName: "v3", TypeSize: 4}}},
	}},
	{Name: "syz_errno", CallName: "syz_errno", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "v", TypeSize: 4}}},
	}},
	{Name: "syz_execute_func", CallName: "syz_execute_func", Args: []Type{
		&PtrType{TypeCommon: TypeCommon{TypeName: "ptr", FldName: "text", TypeSize: 4}, Type: &BufferType{TypeCommon: TypeCommon{TypeName: "text", IsVarlen: true}, Kind: 4}},
	}},
	{Name: "syz_exit", CallName: "syz_exit", Args: []Type{
		&IntType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "int32", FldName: "status", TypeSize: 4}}},
	}},
	{Name: "syz_mmap", CallName: "syz_mmap", Args: []Type{
		&VmaType{TypeCommon: TypeCommon{TypeName: "vma", FldName: "addr", TypeSize: 4}},
		&LenType{IntTypeCommon: IntTypeCommon{TypeCommon: TypeCommon{TypeName: "len", FldName: "len", TypeSize: 4}}, Path: []string{"addr"}},
	}},
}

var consts_32_shmem = []ConstValue{
	{Name: "IPPROTO_ICMPV6", Value: 58},
	{Name: "IPPROTO_TCP", Value: 6},
	{Name: "IPPROTO_UDP", Value: 17},
	{Name: "ONLY_32BITS_CONST", Value: 1},
}

const revision_32_shmem = "bfebe089857c90dd10dd5e67d52fe6357070d402"
