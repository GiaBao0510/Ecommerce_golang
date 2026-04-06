package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {

	// --- Nguyên mẫu
	// var (
	// 	input = 2
	// 	output = 25
	// )

	// actual := AddOne(2)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input: %d, output: %d, actual: %d", input, input, output, actual)
	// }


	// --- Sử dụng thư viện assert
	assert.Equal(t, AddOne(5), 6, "AddOne(5) should be 6")
	assert.NotEqual(t, AddOne(2), 3, "AddOne(2) should not be 2")
	assert.Nil(t, nil, nil)
}

// Hàm require khi đang test mà failed thì những hàm đằng sau sẽ không được thực thi, dừng ngay lâp tức
func TestRequire(t *testing.T) {
	require.Equal(t,  2, 3)
	fmt.Println("---> Not Excuting")
}

// Hàm assert khi đang test mà failed thì những hàm đằng sau vẫn được thực thi
func TestAssert(t *testing.T) {
	assert.Equal(t,  2, 3)
	fmt.Println("---> Excuting")
}