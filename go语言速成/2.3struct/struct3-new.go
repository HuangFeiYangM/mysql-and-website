//尽管缺少构造器，Go 语言却有一个内置的 new 函数，使用它来分配类型所需要的内存。 new(X) 的结果与 &X{} 相同。

package main

import "fmt"

type Saiyan struct {
	Name  string
	Power int
}

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan2 struct {
	*Person
	Power int
}

func main() {
	//new(X) 的结果与 &X{} 相同。
	goku := new(Saiyan)
	// same as
	//goku := &Saiyan{}

	goku.Name = "hfy"
	goku.Power = 114514

	//或者:
	// goku := &Saiyan {
	//     Name: "goku",
	//     Power: 9000,
	// }

	//1
	goku2 := &Saiyan2{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku2.Introduce()

	fmt.Println(goku2)

	//2
	//下面这样也完全有效：
	goku3 := &Saiyan2{
		Person: &Person{"Goku"},
	}
	fmt.Println(goku3.Name)
	fmt.Println(goku3.Person.Name)

}

// ## 组合

// Go 支持组合， 这是将一个结构包含进另一个结构的行为。在某些语言中，这种行为叫做 特质 或者 混合。 没有明确的组合机制的语言总是可以做到这一点。在 Java 中， 可以使用 *继承* 来扩展结构。但是在脚本中并没有这种选项， 混合将会被写成如下形式：

// ```java
// public class Person {
//   private String name;

//   public String getName() {
//     return this.name;
//   }
// }

// public class Saiyan {
//   // Saiyan 中包含着 person 对象
//   private Person person;

//   // 将请求转发到 person 中
//   public String getName() {
//     return this.person.getName();
//   }
//   ...
// }
// ```

// 这可能会非常繁琐。`Person` 的每个方法都需要在 `Saiyan` 中重复。Go 避免了这种复杂性：

// ```go
// type Person struct {
//   Name string
// }

// func (p *Person) Introduce() {
//   fmt.Printf("Hi, I'm %s\n", p.Name)
// }

// type Saiyan struct {
//   *Person
//   Power int
// }

// // 使用它
// goku := &Saiyan{
//   Person: &Person{"Goku"},
//   Power: 9001,
// }
// goku.Introduce()
// ```

// `Saiyan`  结构体有一个 `Person` 类型的字段。由于我们没有显式地给它一个字段名，所以我们可以隐式地访问组合类型的字段和函数。然而，Go 编译器确实给了它一个字段名Person，下面这样完全有效：

// ```go
// goku := &Saiyan{
//   Person: &Person{"Goku"},
// }
// fmt.Println(goku.Name)
// fmt.Println(goku.Person.Name)
// ```

// 上面两个都打印 「Goku」。

// 组合比继承更好吗？许多人认为它是一种更好的组织代码的方式。当使用继承的时候，你的类和超类紧密耦合在一起，你最终专注于结构而不是行为。
