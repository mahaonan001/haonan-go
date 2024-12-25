package pwdcode

import (
	"fmt"
	"testing"
)

func TestSecret(t *testing.T) {
	s := []byte("枯藤老树昏鸦")
	es, err := EnPwdCode(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(es)
	ds, err := DePwdCode(es)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(ds))
}
