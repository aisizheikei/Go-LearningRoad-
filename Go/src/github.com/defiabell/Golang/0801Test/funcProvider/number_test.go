package funcProvider

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	got := Add(1, 2)
	want := 3
	if !reflect.DeepEqual(got, want) {
		//测试用例失败
		t.Errorf("want: %v, but got: %v\n", want, got)
	}
}

func Test2(t *testing.T) {
	got := Add(1, 4)
	want := 4
	if !reflect.DeepEqual(got, want) {
		//测试用例失败
		t.Errorf("want: %v, but got: %v\n", want, got)
	}
}
func Test3(t *testing.T) {
	args1 := make([]int, 1, 2)
	args2 := make([]int, 2, 2)
	t.Errorf("test3 result:%v\n", reflect.DeepEqual(args1, args2))

}
