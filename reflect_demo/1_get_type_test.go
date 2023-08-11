package reflect_demo

import (
	"fmt"
	"reflect"
	"testing"
)

type Id interface {
	GetName() string
}

type User struct {
	Name string `json:"name" myTag:"hello"`
	age  int    ` json:"age"  myTag:"-"`
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int {
	return u.age
}

// 通过reflect.TypeOf 获取Type类型
// Kind 获取当前反射对象的类型
// Elem 获取当前反射对象指向的元素类型
func Test_Get_Type(t *testing.T) {
	typeInt := reflect.TypeOf(1)
	typeString := reflect.TypeOf("hello")

	fmt.Println(typeInt)    // int
	fmt.Println(typeString) // string
	u1 := &User{}
	typeUser := reflect.TypeOf(u1)
	fmt.Println(typeUser)               // *reflect_demo.User
	fmt.Println(typeUser.Kind())        // ptr
	fmt.Println(typeUser.Elem().Kind()) // struct

	typeUser2 := reflect.TypeOf(*u1)
	fmt.Println(typeUser2)        // *reflect_demo.User
	fmt.Println(typeUser2.Kind()) // struct
}

// 通过反射获取struct 的成员属性
func Test_Get_Member(t *testing.T) {

	user := &User{
		Name: "admin",
		age:  12,
	}
	typeUser := reflect.TypeOf(user)
	if typeUser.Kind() == reflect.Ptr {
		typeUser = typeUser.Elem()
	}
	fieldNum := typeUser.NumField()
	fmt.Printf("%v have %d fields \r\n", typeUser, fieldNum)

	for i := 0; i < fieldNum; i++ {
		field := typeUser.Field(i)
		fmt.Printf("index: %d  filedName: %s, fieldOffest: %v, filedAnonymous: %v, filedType: %v, filedIsExported: %v, filedTags: %v  \r\n",
			i, field.Name, field.Offset, field.Anonymous, field.Type, field.IsExported(), field.Tag,
		)
	}

	// 通过 FiledByName || FiledByIndex 参数是[]int 支持嵌套查询
	age, _ := typeUser.FieldByName("age")
	tag := age.Tag
	myTag := tag.Get("myTag")
	fmt.Printf("age myTag : %s \r\n", myTag)
}

func Test_Get_Method(t *testing.T) {

	up := &User{
		Name: "admin",
		age:  12,
	}
	userPtrType := reflect.TypeOf(up)
	userType := reflect.TypeOf(*up)

	ptrMethodNum := userPtrType.NumMethod()
	structMethodNum := userType.NumMethod()

	fmt.Printf("method: ptr receiver have  %d, struct receiver have %d \r\n", ptrMethodNum, structMethodNum)

	getNameFunc, _ := userPtrType.MethodByName("GetName")
	value := getNameFunc.Func
	call := value.Call([]reflect.Value{
		reflect.ValueOf(up),
	})
	fmt.Println(call[0].Interface())
	// 获取函数信息
	funcType := getNameFunc.Func.Type()
	argInNum := funcType.NumIn()
	in := funcType.In(0)
	argOutNum := funcType.NumOut()
	out := funcType.Out(0)

	fmt.Printf("method %v has input arg %d, type: %v \r\n", funcType.Name(), argInNum, in)
	fmt.Printf("method %v has out arg %d , type :%v  \r\n", funcType.Name(), argOutNum, out)

}

func Test_Is_Impl(t *testing.T) {
	var id *Id
	idType := reflect.TypeOf(id).Elem()

	fmt.Printf("idType kind is interface %t\n", idType.Kind() == reflect.Interface)
	up := &User{
		Name: "admin",
		age:  12,
	}
	t1 := reflect.TypeOf(up)
	t2 := reflect.TypeOf(*up)
	fmt.Printf("t1 implements People interface %t\n", t1.Implements(idType))
	fmt.Printf("t2 implements People interface %t\n", t2.Implements(idType))
}

func Test_Value(t *testing.T) {
	up := &User{
		Name: "admin",
		age:  12,
	}
	upValue := reflect.ValueOf(up)
	fmt.Println(upValue)

	uValue := reflect.ValueOf(*up)
	fmt.Println(uValue)

	// equal : reflect.ValueOf(*up)
	uValue2 := upValue.Elem()
	fmt.Println(uValue2)
	// equal : reflect.ValueOf(up)
	upValue2 := uValue2.Addr()
	fmt.Println(upValue2)
	// value -> Type

	t2 := upValue.Type()
	fmt.Println(t2)

	// get value
	fmt.Println(upValue2.Interface())
	getValue := (upValue2.Interface()).(*User)
	fmt.Println(getValue.GetName())

	// empty value
	var nilUser *User = nil
	v1 := reflect.ValueOf(nilUser)
	fmt.Printf("v1.IsValid() : %v \r\n", v1.IsValid())
	if v1.IsValid() {
		fmt.Printf("v1.IsNil() : %v \r\n", v1.IsNil())
	}
	var zeroUser User
	v2 := reflect.ValueOf(zeroUser)
	fmt.Printf("v2.IsValid() : %v \r\n", v2.IsValid())
	if v2.IsValid() {
		fmt.Printf("v2.IsZero() : %v \r\n", v2.IsZero())
	}

	// interface
	var i interface{}
	vi := reflect.ValueOf(i)
	fmt.Printf("vi.IsValid() : %v \r\n", vi.IsValid())

}

func Test_Set_Value(t *testing.T) {
	up := &User{
		Name: "admin",
		age:  12,
	}
	upValue := reflect.ValueOf(up)
	// Only pointers can be modified, only public fields can be modified
	// get Elem to set value
	user := upValue.Elem()
	user.FieldByName("Name").SetString("ge ge")
	// Usually we don't pay attention to the field name when reflecting, so we can judge whether it can be assigned by Can Set
	fmt.Println(user.FieldByName("age").CanSet())
	fmt.Println(upValue)

}

func Test_Set_Slice_Value(t *testing.T) {
	up1 := &User{
		Name: "admin1",
		age:  12,
	}
	up2 := &User{
		Name: "admin2",
		age:  12,
	}
	users := make([]*User, 2, 10)
	users[0] = up1
	users[1] = up2
	sv := reflect.ValueOf(&users)
	elem := sv.Elem()

	elem.Index(0).Elem().FieldByName("Name").SetString("new new admin")

	fmt.Println(users[0])
	fmt.Println(elem.Cap())
	fmt.Println(len(users))

	elem.SetCap(5)
	fmt.Println(cap(users))
}

// MakeFunc MakeMap MakeSlice MakeChan
func Test_Make(t *testing.T) {
	user := &User{
		Name: "admin",
	}
	var f func(user2 *User) string
	ut := reflect.TypeOf(f)
	makeFunc := reflect.MakeFunc(ut, func(args []reflect.Value) (results []reflect.Value) {
		name := args[0].Elem().FieldByName("Name").String()
		return []reflect.Value{reflect.ValueOf("hello !" + name)}
	})
	fmt.Println(makeFunc.Call([]reflect.Value{reflect.ValueOf(user)})[0].String())

}
