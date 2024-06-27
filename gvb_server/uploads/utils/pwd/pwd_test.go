package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$O5OEa6KHcBOUxLXqPh9YLuVQmO.IKzi1gQSoUFbLfOp3ErCPcnWjW", "1234"))
}
