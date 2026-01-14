package pkg

import (
	"crypto/rand"
	"math/big"
)

func GenerateOTP(length int) (string, error) {
    const digits = "0123456789"
    
    otp := make([]byte, length)
    
    maxIndex := big.NewInt(int64(len(digits)))

    for i := 0; i < length; i++ {
        num, err := rand.Int(rand.Reader, maxIndex)
        if err != nil {
            return "", err
        }
        
        otp[i] = digits[num.Int64()]
    }

    return string(otp), nil
}