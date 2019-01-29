package util

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestRsaVerifySha1(t *testing.T) {
	var (
		byte_myPrikey []byte
		err           error
		myPrikey      = `MIIEowIBAAKCAQEAlV1h5xCHGLpeHpl849yuZkDOf9fhvazgbpv5+1MxvEiYEur5iP0YW0peJRgqKwUMAlDVC0Eo3Nv3LAA8/qYrAbtv7TfDponFI2QILxgoACvzBy+YzENWVMYB8GBjc+mslslEjtCLZB5ZOOEWrw2jfE7pXqSNXfvw0IoMSxo79W3umevcQhKENbchHhS/q8gPH2Qhs/6NUPRjtMY6DuI+6LLpwAXWayT+YO1uuAXrz2t5PjSmSlAYBwEPuhYOG0RseQb7nrRJMzzi+QmWpY2F80zYlH7umqV/g5Vd7BvMwr8jXYhsW0Kpe0K7saxVXUEA5Wo8i4xRmWQY1ykO7na6DwIDAQABAoIBAEN7SiUj0XV5ldZLUCJ+bIikFOereEDCny855woGPz0qqxI4+P5MIrnz8m7d6QdAo3lUXbxU9wo5kwVdIjGGj32b2miZPhq6ucdQhWYJiuDw+j6v4V+/uHxabTvVwHj5BNGnIwNG8wMLtxG8mAwEpG0gFc19mKTyk40UIZO4yYgLLwUtTp8l9E8Xg3xN37idecNnQ8GNEg+Qy0ue04bZM1BkwdQhEZtMlMFKfc6r62APBVM3Z4dxsTCZe2lbTMv3UhYhFP6/bhEsUYrDqRpPvQVJS7cSb+PdpGcMqoMNwcvBEpMtm1l0U2vJ9DPvWYVQEv5V3xgYK6eOFEToNYdpwokCgYEAw8+N94UwuZKwc1XQNb1NdLEuI0egHS9nnVYaBOM/I6s1YeJX39w5uDeIVl4ZnIhChUOOFB8EbxFvKXLTVzQ+qLrPXvDuZmJvq2b+H0nFTx5y4PN2c7U7Ue+rUoakaTumpYuRtzjIvTVUGkUa471+OTxBAUiHodNaR9yRpVx/3eUCgYEAw0b4arraleDfYuaT227wlWQT2W24G/cXLntKOs97SAwhLkA8hV67r5ellb46WSbjvQOpg0gwSjR7L7OmTEk7TOHfOQJK9V+YkLXR/RhsHSlEuD3sJ/GRA33bjr2O7uWqvqVXC6wSY0+COhIOjD1LesA61ME/duTP1DijyiqxmOMCgYA4LGXs2U/WHOfz3m1hzVHYJTA4PMcJOF1APMAwIMUvRWGGGDnfZb1FROEe7dXpGwoCUxQCX7eU6Wp7eI56mOlU3Gq7MOEjjB+/C0fhz/cDsJeCQzX7EcXXxqrefPPToKI5IaYG6wpjhVYAR0zkgqsgXlHDvXvzh7+BmBxMVRH1pQKBgQCOF+5be2vi86aZCL8+RRO7IP9wj24Qq1Oq9vukn6VNX8YlKYjgY3ae5vIObEV46dt9hlSqurSoyld49nhjukX0Q8dybECvG8igRC1wxXymG4ltp2FYD2c96y8ARt7i5yu/XgBg9ezLZueT6d+8HD34LDii55uOoCC6hBxq13YzkQKBgFFcz9aWIHTDg8ea6ze8v85HvJafRnXt/1I85LbS/k/uGDmPbFn73JrswfNtZN0EeVs1epWnFnzEVtf1eq6g41LqkLshic653UckogObZAfkN4C3M9a8CUyz1+FZmv4nfBuK7oEIV0QAkDcAV5T/WWFRt3SYSd9ybJUkDzzpGO+t`
		//myPubkey      = `MIIBCgKCAQEAlV1h5xCHGLpeHpl849yuZkDOf9fhvazgbpv5+1MxvEiYEur5iP0YW0peJRgqKwUMAlDVC0Eo3Nv3LAA8/qYrAbtv7TfDponFI2QILxgoACvzBy+YzENWVMYB8GBjc+mslslEjtCLZB5ZOOEWrw2jfE7pXqSNXfvw0IoMSxo79W3umevcQhKENbchHhS/q8gPH2Qhs/6NUPRjtMY6DuI+6LLpwAXWayT+YO1uuAXrz2t5PjSmSlAYBwEPuhYOG0RseQb7nrRJMzzi+QmWpY2F80zYlH7umqV/g5Vd7BvMwr8jXYhsW0Kpe0K7saxVXUEA5Wo8i4xRmWQY1ykO7na6DwIDAQAB`
		message = []byte(`{"autoRenewing":"","orderId":"111111","packageName":"com.test.android.name1","productId":"","purchaseTime":"20200230","purchaseState":"","developerPayload":"","purchaseToken":"token1"}`)
	)

	hashed := sha1.Sum(message)

	//rsaPrivateKey, err := rsa.GenerateKey(rng, 2048)

	if byte_myPrikey, err = base64.StdEncoding.DecodeString(myPrikey); err != nil {
		t.Errorf("Error from signing: %s\n", err)
		return
	}
	rsaPrivateKey, err := x509.ParsePKCS1PrivateKey(byte_myPrikey)

	signature, err := rsa.SignPKCS1v15(nil, rsaPrivateKey, crypto.SHA1, hashed[:])
	if err != nil {
		t.Errorf("Error from signing: %s\n", err)
		return
	}

	b64pk := base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&rsaPrivateKey.PublicKey))

	b64sign := base64.StdEncoding.EncodeToString(signature)

	//加上 jdk 生成秘钥的oid , 在RsaVerifySha1 会去掉
	zeroSlice := make([]byte, 24)
	b64pk_byte, _ := base64.StdEncoding.DecodeString(b64pk)
	b64pk_byte = append(zeroSlice, b64pk_byte...)
	b64pk = base64.StdEncoding.EncodeToString(b64pk_byte)
	//end

	if err := RsaVerifySha1(b64pk, message[:], b64sign); err != nil {
		t.Errorf("RsaVerifySha1 %s\n", err)
		return
	}

	fmt.Printf("pk= %s  \n", b64pk)
	fmt.Printf("sign= %s \n", b64sign)

}
