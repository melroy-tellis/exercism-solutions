package variablelengthquantity

import "fmt"

func EncodeVarint(input []uint32) []byte {
    output := []byte{}
    for _, value := range(input) {
        encodedValue := []byte{byte(value&0x7f)}
        value >>= 7
        for value > 0 {
            encodedValue = append([]byte{byte((value&0x7f)|0x80)}, encodedValue...)
            value >>= 7
        }
        output = append(output, encodedValue...)
    }
    return output
    
}

func DecodeVarint(input []byte) ([]uint32, error) {
    output := []uint32{}
    decodedValue := uint32(0)
    isMsbClear := false
    for _, value := range(input) {
        decodedValue = (decodedValue<<7) | uint32(value&0x7f)
        isMsbClear = (value&0x80) == 0
        
		if isMsbClear {
            output = append(output, decodedValue)
            decodedValue = uint32(0)
        }
    }
    if !isMsbClear {
        return nil, fmt.Errorf("incomplete sequence, last byte does not have a cleared MSB: %v", input)
    }
    return output, nil
    
}
