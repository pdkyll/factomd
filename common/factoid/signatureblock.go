// Copyright 2017 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package factoid

import (
	"fmt"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

/**************************************
 * ISign
 *
 * Interface for RCB Signatures
 *
 * The signature block holds the signatures that validate one of the RCBs.
 * Each signature has an index, so if the RCD is a multisig, you can know
 * how to apply the signatures to the addresses in the RCD.
 **************************************/

type SignatureBlock struct {
	Signatures []interfaces.ISignature
}

var _ interfaces.ISignatureBlock = (*SignatureBlock)(nil)

func (b *SignatureBlock) IsSameAs(s interfaces.ISignatureBlock) bool {
	if s == nil {
		return b == nil
	}

	sigs := s.GetSignatures()
	if len(b.Signatures) != len(sigs) {
		return false
	}
	for i := range b.Signatures {
		if b.Signatures[i].IsSameAs(sigs[i]) == false {
			return false
		}
	}

	return true
}

func (b SignatureBlock) UnmarshalBinary(data []byte) error {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err := b.UnmarshalBinaryData(data)
	return err
}

func (e *SignatureBlock) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *SignatureBlock) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (b SignatureBlock) String() string {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockString.Observe(float64(time.Now().UnixNano() - callTime))	
	txt, err := b.CustomMarshalText()
	if err != nil {
		return "<error>"
	}
	return string(txt)
}

func (s *SignatureBlock) AddSignature(sig interfaces.ISignature) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockAddSignature.Observe(float64(time.Now().UnixNano() - callTime))	
	if len(s.Signatures) > 0 {
		s.Signatures[0] = sig
	} else {
		s.Signatures = append(s.Signatures, sig)
	}
}

func (s SignatureBlock) GetSignature(index int) interfaces.ISignature {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockGetSignature.Observe(float64(time.Now().UnixNano() - callTime))	
	if len(s.Signatures) <= index {
		return nil
	}
	return s.Signatures[index]
}

func (s SignatureBlock) GetSignatures() []interfaces.ISignature {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockGetSignatures.Observe(float64(time.Now().UnixNano() - callTime))	
	if s.Signatures == nil {
		s.Signatures = make([]interfaces.ISignature, 1, 1)
		s.Signatures[0] = new(FactoidSignature)
	}
	return s.Signatures
}

func (a SignatureBlock) MarshalBinary() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer

	for _, sig := range a.GetSignatures() {
		data, err := sig.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Signature failed to Marshal in RCD_1")
		}
		out.Write(data)
	}

	return out.DeepCopyBytes(), nil
}

func (s SignatureBlock) CustomMarshalText() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockCustomMarshalText.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer

	out.WriteString("Signature Block: \n")
	for _, sig := range s.Signatures {
		out.WriteString(" signature: ")
		txt, err := sig.CustomMarshalText()
		if err != nil {
			return nil, err
		}
		out.Write(txt)
		out.WriteString("\n ")

	}

	return out.DeepCopyBytes(), nil
}

func (s *SignatureBlock) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	s.Signatures = make([]interfaces.ISignature, 1)
	s.Signatures[0] = new(FactoidSignature)
	data, err = s.Signatures[0].UnmarshalBinaryData(data)
	if err != nil {
		return nil, fmt.Errorf("Failure to unmarshal Signature")
	}

	return data, nil
}

func NewSingleSignatureBlock(priv, data []byte) *SignatureBlock {
	callTime := time.Now().UnixNano()
	defer factoidSignatureBlockNewSingleSignatureBlock.Observe(float64(time.Now().UnixNano() - callTime))	
	s := new(SignatureBlock)
	s.AddSignature(NewED25519Signature(priv, data))
	return s
}
