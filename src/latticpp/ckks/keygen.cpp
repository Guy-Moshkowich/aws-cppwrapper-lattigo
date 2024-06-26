// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "keygen.h"

using namespace std;

namespace latticpp {

    KeyGenerator newKeyGenerator(const Parameters &params) {
        return KeyGenerator(lattigo_newKeyGenerator(params.getRawHandle()));
    }

    SwitchingKey getSwitchingKey(const RotationKeys &rotationKeys, uint64_t galEl) {
      return SwitchingKey(lattigo_getSwitchingKey(rotationKeys.getRawHandle(), galEl));
    }

    void setSwitchingKey(const RotationKeys &rotKeys, const SwitchingKey &swk, uint64_t galEl) {
      lattigo_setSwitchingKey(rotKeys.getRawHandle(), swk.getRawHandle(), galEl);
    }

    uint64_t switchingKeyExist(const RotationKeys &rotationKeys, uint64_t galEl) {
      return lattigo_switchingKeyExist(rotationKeys.getRawHandle(), galEl);
    }

    vector<uint64_t> getGaloisElementsOfRotationKeys(const RotationKeys &rotationKeys) {
      vector<uint64_t> res(lattigo_getNumSwitchingKeys(rotationKeys.getRawHandle()));
      lattigo_getGaloisElementsOfRotationKeys(rotationKeys.getRawHandle(), res.data());
      return res;
    }

    SwitchingKey copyNew(const SwitchingKey &swk) {
      return SwitchingKey(lattigo_copyNewSwitchingKey(swk.getRawHandle()));
    }

    uint64_t numOfDecomp(const SwitchingKey &swk) {
      return lattigo_numOfDecomp(swk.getRawHandle());
    }

    uint64_t galoisElementForColumnRotationBy(const Parameters &params, uint64_t rotationStep) {
      return lattigo_galoisElementForColumnRotationBy(params.getRawHandle(),
                                                      rotationStep);
    }

    uint64_t switchingKeyIsCorrect(const SwitchingKey &swk, uint64_t galEl, const SecretKey &sk, const Parameters &params, uint64_t log2Bound) {
      return lattigo_switchingKeyIsCorrect(swk.getRawHandle(), galEl,
                                          sk.getRawHandle(), params.getRawHandle(),
                                          log2Bound);
    }

    SecretKey newSecretKey(const Parameters &params) {
      return SecretKey(lattigo_newSecretKey(params.getRawHandle()));
    }

    SecretKey copyNewSecretKey(const SecretKey &sk) {
      return SecretKey(lattigo_copyNewSecretKey(sk.getRawHandle()));
    }

    PolyQP polyQP(const SecretKey &sk) {
      return PolyQP(lattigo_polyQPSecretKey(sk.getRawHandle()));
    }

    PublicKey newPublicKey(const Parameters &params) {
      return PublicKey(lattigo_newPublicKey(params.getRawHandle()));
    }

    RelinearizationKey newRelinearizationKey(const Parameters &params) {
      return RelinearizationKey(
          lattigo_newRelinearizationKey(params.getRawHandle()));
    }

    RotationKeys newRotationKeys(const Parameters &params,
                                 vector<uint64_t> galoisElements) {
      return RotationKeys(lattigo_newRotationKeys(
          params.getRawHandle(), galoisElements.data(), galoisElements.size()));
    }

    SecretKey genSecretKey(const KeyGenerator &keygen) {
      return SecretKey(lattigo_genSecretKey(keygen.getRawHandle()));
    }

    PublicKey genPublicKey(const KeyGenerator &keygen, const SecretKey &sk) {
      return PublicKey(
          lattigo_genPublicKey(keygen.getRawHandle(), sk.getRawHandle()));
    }

    KeyPairHandle genKeyPair(const KeyGenerator &keygen) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPair(keygen.getRawHandle());
        return KeyPairHandle { SecretKey(kp.sk), PublicKey(kp.pk) };
    }

    KeyPairHandle genKeyPairSparse(const KeyGenerator &keygen, uint64_t hw) {
        Lattigo_KeyPairHandle kp = lattigo_genKeyPairSparse(keygen.getRawHandle(), hw);
        return KeyPairHandle { SecretKey(kp.sk), PublicKey(kp.pk) };
    }

    RelinearizationKey genRelinKey(const KeyGenerator &keygen, const SecretKey &sk) {
        return RelinearizationKey(lattigo_genRelinearizationKey(keygen.getRawHandle(), sk.getRawHandle()));
    }

    RotationKeys genRotationKeysForRotations(const KeyGenerator &keygen, const SecretKey &sk, vector<int> shifts, int includeConjugate) {
        // convert from variable-sized int to fixed-size SIGNED int64_t
        vector<int64_t> fixed_width_shifts(shifts.size());
        for (int i = 0; i < shifts.size(); i++) {
            fixed_width_shifts[i] = static_cast<int64_t>(shifts[i]);
        }
        return RotationKeys(lattigo_genRotationKeysForRotations(keygen.getRawHandle(), sk.getRawHandle(), fixed_width_shifts.data(), shifts.size(), includeConjugate));
    }

    CiphertextQP getCiphertextQP(const SwitchingKey &swk, uint64_t i, uint64_t j) {
      return CiphertextQP(lattigo_getCiphertextQP(swk.getRawHandle(), i, j));
    }

    void setCiphertextQP(SwitchingKey &swk, const CiphertextQP &ctQP, uint64_t i, uint64_t j) {
      lattigo_setCiphertextQP(swk.getRawHandle(), ctQP.getRawHandle(), i, j);
    }

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey) {
        return EvaluationKey(lattigo_makeEvaluationKeyOnlyRelin(relinKey.getRawHandle()));
    }    

    EvaluationKey makeEvaluationKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys) {
        return EvaluationKey(lattigo_makeEvaluationKey(relinKey.getRawHandle(), rotKeys.getRawHandle()));
    }

    EvaluationKey makeEmptyEvaluationKey() {
      return EvaluationKey(lattigo_makeEmptyEvaluationKey());
    }

    void setRelinKeyForEvaluationKey(const EvaluationKey &evalKey,
                                     const RelinearizationKey &relinKey) {
      lattigo_setRelinKeyForEvaluationKey(evalKey.getRawHandle(),
                                          relinKey.getRawHandle());
    }

    void setRotKeysForEvaluationKey(const EvaluationKey &evalKey,
                                    const RotationKeys &rotKeys) {
      lattigo_setRotKeysForEvaluationKey(evalKey.getRawHandle(),
                                         rotKeys.getRawHandle());
    }

  BootstrappingKey genBootstrappingKey(const KeyGenerator &keygen, const Parameters &params, const BootstrappingParameters &bootParams, const SecretKey &sk) {
        return BootstrappingKey(lattigo_genBootstrappingKey(keygen.getRawHandle(), params.getRawHandle(), bootParams.getRawHandle(), sk.getRawHandle()));
    }

    BootstrappingKey genBootstrappingKey(const RelinearizationKey &relinKey, const RotationKeys &rotKeys, const SwitchingKey swkDtS, const SwitchingKey swkStD) {
        return BootstrappingKey(lattigo_genBootstrappingKeyByParams(relinKey.getRawHandle(), rotKeys.getRawHandle(),swkDtS.getRawHandle() , swkStD.getRawHandle()));
    }

    BootstrapSwkPairHandle genBootstrapSwkPair(const Parameters &params, const BootstrappingParameters &bootParams, const SecretKey &sk){
      Lattigo_BootstrapSwkPairHandle btsSwkPairHandle = lattigo_genBootstrapSwkPair(params.getRawHandle(), bootParams.getRawHandle(), sk.getRawHandle());
      return BootstrapSwkPairHandle{ SwitchingKey(btsSwkPairHandle.swkStD), SwitchingKey(btsSwkPairHandle.swkDtS) };
    }

    SwitchingKey newSwitchingKey(const Parameters &params, uint64_t levelQ, uint64_t levelP) {
      return SwitchingKey(lattigo_newSwitchingKey(params.getRawHandle(), levelQ, levelP));
    }
}  // namespace latticpp