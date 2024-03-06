// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package ckks

// cgo will automatically generate a struct for functions which return multiple values,
// but the auto-generated struct with generated names loses its semantic value. We opt
// to define our own struct here.

/*
#include "stdint.h"
struct Lattigo_KeyPairHandle {
  uint64_t sk;
  uint64_t pk;
};
*/
import "C"

import (
	"fmt"
	"github.com/tuneinsight/lattigo/v4/ckks"
	"github.com/tuneinsight/lattigo/v4/ckks/bootstrapping"
	"github.com/tuneinsight/lattigo/v4/rlwe"
	"lattigo-cpp/marshal"
	"unsafe"
)

// https://github.com/golang/go/issues/35715#issuecomment-791039692
type Handle5 = uint64

func getStoredKeyGenerator(keygenHandle Handle5) *rlwe.KeyGenerator {
	ref := marshal.CrossLangObjMap.Get(keygenHandle)
	return (*rlwe.KeyGenerator)(ref.Ptr)
}

func getStoredSecretKey(skHandle Handle5) *rlwe.SecretKey {
	ref := marshal.CrossLangObjMap.Get(skHandle)
	return (*rlwe.SecretKey)(ref.Ptr)
}

func getStoredPublicKey(pkHandle Handle5) *rlwe.PublicKey {
	ref := marshal.CrossLangObjMap.Get(pkHandle)
	return (*rlwe.PublicKey)(ref.Ptr)
}

func getStoredEvaluationKey(evalKeyHandle Handle5) *rlwe.EvaluationKey {
	ref := marshal.CrossLangObjMap.Get(evalKeyHandle)
	return (*rlwe.EvaluationKey)(ref.Ptr)
}

func getStoredRelinKey(relinKeyHandle Handle5) *rlwe.RelinearizationKey {
	ref := marshal.CrossLangObjMap.Get(relinKeyHandle)
	return (*rlwe.RelinearizationKey)(ref.Ptr)
}

func getStoredRotationKeys(rotKeysHandle Handle5) *rlwe.RotationKeySet {
	ref := marshal.CrossLangObjMap.Get(rotKeysHandle)
	return (*rlwe.RotationKeySet)(ref.Ptr)
}

func getStoredSwitchingKey(switchingKeyHandle Handle5) *rlwe.SwitchingKey {
	ref := marshal.CrossLangObjMap.Get(switchingKeyHandle)
	return (*rlwe.SwitchingKey)(ref.Ptr)
}

func getStoredBootstrappingKey(bootKeyHandle Handle5) *bootstrapping.EvaluationKeys {
	ref := marshal.CrossLangObjMap.Get(bootKeyHandle)
	return (*bootstrapping.EvaluationKeys)(ref.Ptr)
}

//export lattigo_newKeyGenerator
func lattigo_newKeyGenerator(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	var keyGenerator rlwe.KeyGenerator
	keyGenerator = ckks.NewKeyGenerator(*paramPtr)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&keyGenerator))
}

//export lattigo_newSecretKey
func lattigo_newSecretKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rlwe.NewSecretKey((*paramPtr).Parameters)))
}

//export lattigo_newPublicKey
func lattigo_newPublicKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rlwe.NewPublicKey((*paramPtr).Parameters)))
}

//export lattigo_newRelinearizationKey
func lattigo_newRelinearizationKey(paramHandle Handle5) Handle5 {
	paramPtr := getStoredParameters(paramHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rlwe.NewRelinearizationKey((*paramPtr).Parameters, 1)))
}

//export lattigo_newRotationKeys
func lattigo_newRotationKeys(paramHandle Handle5, galoisElements *C.uint64_t, galoisElementsLen uint64) Handle5 {
	paramPtr := getStoredParameters(paramHandle)

	galoisElementsTmp := make([]uint64, galoisElementsLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(galoisElements))
	for i := range galoisElementsTmp {
		galoisElementsTmp[i] = *(*uint64)(unsafe.Pointer(basePtrIn + size*uintptr(i)))
	}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rlwe.NewRotationKeySet((*paramPtr).Parameters, galoisElementsTmp)))
}

//export lattigo_genSecretKey
func lattigo_genSecretKey(keygenHandle Handle5) Handle5 {
	keygen := getStoredKeyGenerator(keygenHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenSecretKey()))
}

//export lattigo_copyNewSecretKey
func lattigo_copyNewSecretKey(skHandle Handle5) Handle5 {
	sk := getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(sk.CopyNew()))
}

//export lattigo_polyQPSecretKey
func lattigo_polyQPSecretKey(skHandle Handle5) Handle5 {
	sk := getStoredSecretKey(skHandle)
	polyQP := sk.Value
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&polyQP))
}

//export lattigo_genPublicKey
func lattigo_genPublicKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	keygen := getStoredKeyGenerator(keygenHandle)
	sk := getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenPublicKey(sk)))
}

//export lattigo_genKeyPair
func lattigo_genKeyPair(keygenHandle Handle5) C.struct_Lattigo_KeyPairHandle {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	var pk *rlwe.PublicKey
	sk, pk = (*keygen).GenKeyPair()
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

//export lattigo_genKeyPairSparse
func lattigo_genKeyPairSparse(keygenHandle Handle5, hw uint64) C.struct_Lattigo_KeyPairHandle {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	var pk *rlwe.PublicKey
	sk = (*keygen).GenSecretKeyWithHammingWeight(int(hw))
	pk = (*keygen).GenPublicKey(sk)
	var kpHandle C.struct_Lattigo_KeyPairHandle
	kpHandle.sk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(sk)))
	kpHandle.pk = C.uint64_t(marshal.CrossLangObjMap.Add(unsafe.Pointer(pk)))
	return kpHandle
}

// only generates relinearization keys for ciphertexts up to degree 2
// (i.e., you must relinearize after each ct/ct multiplication)
//
//export lattigo_genRelinearizationKey
func lattigo_genRelinearizationKey(keygenHandle Handle5, skHandle Handle5) Handle5 {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)
	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer((*keygen).GenRelinearizationKey(sk, 2)))
}

// Positive k is for left rotation by k positions
// Negative k is equivalent to a right rotation by k positions
//
//export lattigo_genRotationKeysForRotations
func lattigo_genRotationKeysForRotations(keygenHandle Handle5, skHandle Handle5, ks *C.int64_t, ksLen uint64, includeConjugate uint8) Handle5 {
	var keygen *rlwe.KeyGenerator
	keygen = getStoredKeyGenerator(keygenHandle)

	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)

	rotations := make([]int, ksLen)
	size := unsafe.Sizeof(uint64(0))
	basePtrIn := uintptr(unsafe.Pointer(ks))
	for i := range rotations {
		rotations[i] = int(*(*int64)(unsafe.Pointer(basePtrIn + size*uintptr(i))))
	}
	includeConjugateVal := (includeConjugate == 1)
	var rotKeys *rlwe.RotationKeySet
	// The second argument determines if conjugation keys are generated or not. This wrapper API does
	// not support generating a conjugation key.
	rotKeys = (*keygen).GenRotationKeysForRotations(rotations, includeConjugateVal, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotKeys))
}

//export lattigo_getSwitchingKey
func lattigo_getSwitchingKey(switchingKeyHandle Handle5, galEl uint64) Handle5 {
	rotKeys := getStoredRotationKeys(switchingKeyHandle)
	rotationKey := rotKeys.Keys[galEl]
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotationKey))
}

//export lattigo_setSwitchingKey
func lattigo_setSwitchingKey(rotKeysHandle, switchingKeyHandle Handle5, galEl uint64) {
	rotKeys := getStoredRotationKeys(rotKeysHandle)
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	rotKeys.Keys[galEl] = rotKey
}

//export lattigo_switchingKeyExist
func lattigo_switchingKeyExist(switchingKeyHandle Handle5, galEl uint64) uint64 {
	rotKeys := getStoredRotationKeys(switchingKeyHandle)
	_, exist := rotKeys.Keys[galEl]
	if exist {
		return uint64(1)
	} else {
		return uint64(0)
	}
}

//export lattigo_getNumSwitchingKeys
func lattigo_getNumSwitchingKeys(rotKeysHandle Handle5) uint64 {
	rotKeys := getStoredRotationKeys(rotKeysHandle)
	return uint64(len(rotKeys.Keys))
}

//export lattigo_getGaloisElementsOfRotationKeys
func lattigo_getGaloisElementsOfRotationKeys(rotKeysHandle Handle5, outValues *C.uint64_t) {
	rotKeys := getStoredRotationKeys(rotKeysHandle)
	galoisElements := make([]uint64, len(rotKeys.Keys))

	i := 0
	for k := range rotKeys.Keys {
		galoisElements[i] = k
		i++
	}

	size := unsafe.Sizeof(uint64(0))
	basePtr := uintptr(unsafe.Pointer(outValues))
	for i := range galoisElements {
		*(*uint64)(unsafe.Pointer(basePtr + size*uintptr(i))) = galoisElements[i]
	}
}

//export lattigo_copyNewSwitchingKey
func lattigo_copyNewSwitchingKey(switchingKeyHandle Handle5) Handle5 {
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rotKey.CopyNew()))
}

//export lattigo_numOfDecomp
func lattigo_numOfDecomp(switchingKeyHandle Handle5) uint64 {
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	return uint64(len(rotKey.Value))
}

//export lattigo_galoisElementForColumnRotationBy
func lattigo_galoisElementForColumnRotationBy(paramHandle Handle5, rotationStep uint64) uint64 {
	param := getStoredParameters(paramHandle)
	return uint64(param.GaloisElementForColumnRotationBy(int(rotationStep)))
}

//export lattigo_switchingKeyIsCorrect
func lattigo_switchingKeyIsCorrect(switchingKeyHandle Handle5, galEl uint64, skHandle Handle5, paramHandle Handle5, log2Bound uint64) uint64 {
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	sk := getStoredSecretKey(skHandle)
	param := getStoredParameters(paramHandle)
	rotKey.GadgetCiphertext.CopyNew()
	rotKey.CopyNew()
	isCorrect := rlwe.RotationKeyIsCorrect(rotKey.CopyNew(), galEl, sk.CopyNew(), param.Parameters, int(log2Bound))
	if isCorrect {
		return uint64(1)
	} else {
		return uint64(0)
	}
}

//export lattigo_getCiphertextQP
func lattigo_getCiphertextQP(switchingKeyHandle Handle5, i, j uint64) Handle5 {
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&(rotKey.Value[i][j])))
}

//export lattigo_setCiphertextQP
func lattigo_setCiphertextQP(switchingKeyHandle, ctQPHandle Handle5, i, j uint64) {
	rotKey := getStoredSwitchingKey(switchingKeyHandle)
	ctQP := getStoredCiphertextQP(ctQPHandle)
	rotKey.Value[i][j] = *ctQP
}

//export lattigo_makeEvaluationKeyOnlyRelin
func lattigo_makeEvaluationKeyOnlyRelin(relinKeyHandle Handle5) Handle5 {
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	var evalKey rlwe.EvaluationKey
	evalKey = rlwe.EvaluationKey{Rlk: relinKey}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evalKey))
}

//export lattigo_makeEvaluationKey
func lattigo_makeEvaluationKey(relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	var relinKey *rlwe.RelinearizationKey
	relinKey = getStoredRelinKey(relinKeyHandle)

	var rotKeys *rlwe.RotationKeySet
	rotKeys = getStoredRotationKeys(rotKeyHandle)

	var evalKey rlwe.EvaluationKey
	evalKey = rlwe.EvaluationKey{Rlk: relinKey, Rtks: rotKeys}

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&evalKey))
}

//export lattigo_makeEmptyEvaluationKey
func lattigo_makeEmptyEvaluationKey() Handle5 {
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&rlwe.EvaluationKey{}))
}

//export lattigo_setRelinKeyForEvaluationKey
func lattigo_setRelinKeyForEvaluationKey(evalKeyHandle Handle5, relinKeyHandle Handle5) {
	evalKey := getStoredEvaluationKey(evalKeyHandle)
	evalKey.Rlk = getStoredRelinKey(relinKeyHandle)
}

//export lattigo_setRotKeysForEvaluationKey
func lattigo_setRotKeysForEvaluationKey(evalKeyHandle Handle5, rotKeysHandle Handle5) {
	evalKey := getStoredEvaluationKey(evalKeyHandle)
	evalKey.Rtks = getStoredRotationKeys(rotKeysHandle)
}

//export lattigo_genBootstrappingKey
func lattigo_genBootstrappingKey(keygenHandle Handle5, paramHandle Handle5, btpParamsHandle Handle5, skHandle Handle5, relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var btpParams *bootstrapping.Parameters
	btpParams = getStoredBootstrappingParameters(btpParamsHandle)

	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)

	var btpKey bootstrapping.EvaluationKeys
	btpKey = bootstrapping.GenEvaluationKeys(*btpParams, *params, sk)

	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&btpKey))
}

//export lattigo_genSwkDenseToSparse
func lattigo_genSwkDenseToSparse(paramHandle Handle5, btpParamsHandle Handle5, skHandle Handle5) Handle5 {
	btpParams := getStoredBootstrappingParameters(btpParamsHandle)
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	swkDtS, _ := btpParams.GenEncapsulationSwitchingKeys(*params, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&swkDtS))
}

//export lattigo_genSwkSparseToDense
func lattigo_genSwkSparseToDense(paramHandle Handle5, btpParamsHandle Handle5, skHandle Handle5) Handle5 {
	btpParams := getStoredBootstrappingParameters(btpParamsHandle)
	params := getStoredParameters(paramHandle)
	sk := getStoredSecretKey(skHandle)
	_, swkStD := btpParams.GenEncapsulationSwitchingKeys(*params, sk)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&swkStD))
}

//export lattigo_genBootstrappingKey2
func lattigo_genBootstrappingKey2(keygenHandle Handle5, paramHandle Handle5, btpParamsHandle Handle5, skHandle Handle5, relinKeyHandle Handle5, rotKeyHandle Handle5) Handle5 {
	fmt.Println("lattigo_genBootstrappingKey2")
	var params *ckks.Parameters
	params = getStoredParameters(paramHandle)

	var btpParams *bootstrapping.Parameters
	btpParams = getStoredBootstrappingParameters(btpParamsHandle)

	var sk *rlwe.SecretKey
	sk = getStoredSecretKey(skHandle)
	rotkeys := getStoredRotationKeys(rotKeyHandle)
	rlk := getStoredRelinKey(relinKeyHandle)
	swkDtS, swkStD := btpParams.GenEncapsulationSwitchingKeys(*params, sk)

	var btpKey bootstrapping.EvaluationKeys
	btpKey = bootstrapping.EvaluationKeys{
		EvaluationKey: rlwe.EvaluationKey{
			Rlk:  rlk,
			Rtks: rotkeys},
		SwkDtS: swkDtS,
		SwkStD: swkStD,
	}
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(&btpKey))

}

//export lattigo_newSwitchingKey
func lattigo_newSwitchingKey(paramsHandle Handle5, levelQ, levelP uint64) Handle5 {
	params := getStoredParameters(paramsHandle)
	return marshal.CrossLangObjMap.Add(unsafe.Pointer(rlwe.NewSwitchingKey(params.Parameters, int(levelQ), int(levelP))))
}
