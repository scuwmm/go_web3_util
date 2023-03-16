package gf256

import "errors"

const lenghOfBytes = 16

// y_1（密码）：y-1 通过密码生成的hash值 []byte后16位
// y0（私钥）：128bit （16byte）的助记词
// y1（守护码）： (y_1 + y0)/2 = y1 ; 由此可以计算 y0 = 2*y1 - y_1
func ToGuardCode(y_1 []byte, y0 []byte) ([]byte, error) {
	if len(y_1) < lenghOfBytes || len(y0) < lenghOfBytes {
		return nil, errors.New("y_1/y0 length should be greater than 16")
	}

	field256 := NewField(0x11d, 2)

	leftArr := y_1[(len(y_1) - lenghOfBytes):]
	MidArr := y0[(len(y0) - lenghOfBytes):]
	rightArr := make([]byte, lenghOfBytes)
	for i, left := range leftArr {
		mid := MidArr[i]
		right := field256.Mul(field256.Add(left, mid), field256.Inv(0x2)) //Inv逆元， 参数1 除以 参数2
		rightArr[i] = right
	}

	return rightArr, nil
}

func ToPriKey(y_1 []byte, y1 []byte) ([]byte, error) {
	if len(y_1) < lenghOfBytes || len(y1) < lenghOfBytes {
		return nil, errors.New("y_1/y1 length should be greater than 16")
	}

	field256 := NewField(0x11d, 2)
	leftArr := y_1[(len(y_1) - lenghOfBytes):]
	rightArr := y1[(len(y_1) - lenghOfBytes):]
	midArr := make([]byte, lenghOfBytes)
	for i, left := range leftArr {
		right := rightArr[i]
		mid := field256.Add(field256.Mul(right, 0x2), left) //参数1 乘以 参数2
		midArr[i] = mid
	}

	return midArr, nil

}
