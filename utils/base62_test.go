package utils

import (
	"math"
	"testing"
)

func TestDecode(t *testing.T) {
	testSet := []string{
		"a",
		"b",
		"bM",
		"k9viXaIfiWh",
	}

	resSet := []int64{
		0,
		1,
		100,
		math.MaxInt64,
	}

	caseNum := 4
	for i := 0; i < caseNum; i++{
		if res, err := Decode(testSet[i]); res != resSet[i] || err != nil{
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("Convert %s fail, Except %d, get %d", testSet[i], resSet[i], res)
		}
	}
}

func TestEncode(t *testing.T) {
	testSet := []int64{
		0,
		1,
		100,
		math.MaxInt64,
	}

	resSet := []string{
		"a",
		"b",
		"bM",
		"k9viXaIfiWh",
	}

	caseNum := 4
	for i := 0; i < caseNum; i++{
		if res := Encode(testSet[i]); res != resSet[i] {
			t.Fatalf("Convert %d fail, Except %s, get %s", testSet[i], resSet[i], res)
		}
	}

}