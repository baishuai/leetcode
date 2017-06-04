package p591

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test0(t *testing.T) {
	assert.Equal(t, true, isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
}

func Test1(t *testing.T) {
	assert.Equal(t, true, isValid("<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"))
}

func Test2(t *testing.T) {
	assert.Equal(t, false, isValid("<A>  <B> </A>   </B>"))
}

func Test3(t *testing.T) {
	assert.Equal(t, false, isValid("tes <A>  <B> </A>   </B>"))
}

func Test4(t *testing.T) {
	assert.Equal(t, false, isValid("<A2>  <B> </A>   </B>"))
}

func Test5(t *testing.T) {
	assert.Equal(t, false, isValid("<ASDDSDFSDEWRW>  <B> </A>   </B>"))
}

func Test6(t *testing.T) {
	assert.Equal(t, false, isValid("<DIV>>>  ![cdata[]] <![CDATA[<div>]>]>]>>]</DIV>"))
}

func Test7(t *testing.T) {
	assert.Equal(t, false, isValid("<A>  <B> </2ADDDDDDDDDD>   </B>"))
	assert.Equal(t, false, isValid("<A>  <B> </223>"))
	assert.Equal(t, false, isValid("<A>  <B> </>   </B>"))
	assert.Equal(t, false, isValid("<A"))
}
