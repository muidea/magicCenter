package util

import "testing"

func TestIntArray2Str(t *testing.T) {
	tempArray := []int{1, 2}
	str := IntArray2Str(tempArray)
	if str != "1,2" {
		t.Errorf("IntArray2Str failed, %s", str)
	}

	tempArray = []int{}
	str = IntArray2Str(tempArray)
	if str != "" {
		t.Errorf("IntArray2Str failed, %s", str)
	}
}

func TestStr2IntArray(t *testing.T) {
	str := ""
	tempArray, ok := Str2IntArray(str)
	if !ok || len(tempArray) > 0 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}

	str = "1"
	tempArray, ok = Str2IntArray(str)
	if !ok || len(tempArray) != 1 || tempArray[0] != 1 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}

	str = ",1"
	tempArray, ok = Str2IntArray(str)
	if !ok || len(tempArray) != 1 || tempArray[0] != 1 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}
	str = "1,"
	tempArray, ok = Str2IntArray(str)
	if !ok || len(tempArray) != 1 || tempArray[0] != 1 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}
	str = ",1,"
	tempArray, ok = Str2IntArray(str)
	if !ok || len(tempArray) != 1 || tempArray[0] != 1 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}

	str = ",1,2,3,4"
	tempArray, ok = Str2IntArray(str)
	if !ok || len(tempArray) != 4 || tempArray[0] != 1 {
		t.Errorf("Str2IntArray failed, ok=%d, len(tempArray)=%d", ok, len(tempArray))
	}
}
