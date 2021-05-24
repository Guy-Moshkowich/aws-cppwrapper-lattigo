// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

#include "marshaler.h"
#include <iterator>
#include <sstream>
#include <vector>

using namespace std;

namespace latticpp {

    void writeToStream(void* ostreamPtr, void* data, uint64_t len) {
        (*((ostream*)ostreamPtr)).write((const char*)data, len);
    }

    void marshalBinaryCiphertext(Ciphertext ct, std::ostream &stream) {
        lattigo_marshalBinaryCiphertext(ct.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryParameters(Parameters params, std::ostream &stream) {
        lattigo_marshalBinaryParameters(params.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinarySecretKey(SecretKey sk, std::ostream &stream) {
        lattigo_marshalBinarySecretKey(sk.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryPublicKey(PublicKey pk, std::ostream &stream) {
        lattigo_marshalBinaryPublicKey(pk.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryEvaluationKey(EvaluationKey evaKey, std::ostream &stream) {
        lattigo_marshalBinaryEvaluationKey(evaKey.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    void marshalBinaryRotationKeys(RotationKeys rotKeys, std::ostream &stream) {
        lattigo_marshalBinaryRotationKeys(rotKeys.getRawHandle(), &writeToStream, (void*)(&stream));
    }

    Ciphertext unmarshalBinaryCiphertext(istream &stream) {
        // Note: the next line is a well-known hard parsing problem for C++.
        // See https://stackoverflow.com/questions/4423361/constructing-a-vector-with-istream-iterators
        // and https://arstechnica.com/civis/viewtopic.php?f=20&t=767929
        // In addition to the difficult parsing problem, you also must import the <vector> and <iterator>
        // headers. Without them, you get obscure errors.
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return Ciphertext(lattigo_unmarshalBinaryCiphertext(buffer.data(), buffer.size()));
    }

    Parameters unmarshalBinaryParameters(istream &stream) {
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return Parameters(lattigo_unmarshalBinaryParameters(buffer.data(), buffer.size()));
    }

    SecretKey unmarshalBinarySecretKey(istream &stream) {
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return SecretKey(lattigo_unmarshalBinarySecretKey(buffer.data(), buffer.size()));
    }

    PublicKey unmarshalBinaryPublicKey(istream &stream) {
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return PublicKey(lattigo_unmarshalBinaryPublicKey(buffer.data(), buffer.size()));
    }

    EvaluationKey unmarshalBinaryEvaluationKey(istream &stream) {
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return EvaluationKey(lattigo_unmarshalBinaryEvaluationKey(buffer.data(), buffer.size()));
    }

    RotationKeys unmarshalBinaryRotationKeys(istream &stream) {
        vector<char> buffer(istreambuf_iterator<char>{stream}, {});
        return RotationKeys(lattigo_unmarshalBinaryRotationKeys(buffer.data(), buffer.size()));
    }
}  // namespace latticpp