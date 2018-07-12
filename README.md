# go 学习练习案例

### go语言特点
    * 1.静态类型、编译型的开源语言
    * 2.脚本化的语法，支持多种编程范式（函数式&面向对象）
    * 3.原生、给力的并发编程支持

### go语言的优劣势
    * 优势
    * 1.脚本化的语法
    * 2.静态类型+编译型，程序运行速度有保障
    * 3.原生的扶持并发编程（降低开发、维护成本；程序可以更好的执行）
    * 劣势
    * 1.语法糖并没有Python和Ruby那么多
    * 2.目前的程序运行速度还不及C（已超过C++和Java）
    * 3.第三方函数库暂时不像绝对主流的编程语言那么多

### go源码文件
    * .go为后缀，内容以go语言代码组织的文件

### go源码文件分类
    * 1.命令源码文件（声明自己属于main代码包、包含无参数声明和结果声明的main函数）
        同一个代码包中强烈不建议直接包含多个命令源码文件
    * 2.库源码文件（不具备命令源码文件两个特征的源码文件）
    * 3.测试源码文件（不具备命令源码文件两个特征的源码文件，名称以_test.go为后缀）
        测试函数以Test（功能测试函数）或Benchmark为前缀（性能测试函数）

### 代码包的作用
    * 编译和归档go程序的最基本单位
    * 代码划分、集结和依赖的有效组织形式，也是权限控制的辅助手段

### 程序实体与关键字
    * 在go语言中，变量、常量、函数、结构体和接口被统称为“程序实体”，而它们的名字被统称为“标识符”。
    * 标识符可以是任何Unicode编码可以表示的字母、数字及下划线。首字母不能是数字或下划线。
    * 对程序实体的访问权限控制只能通过它们的名字来实现。名字首字母大写的程序实体可以被任何代码包访问到。而名字首字母小写的程序实体则只能被同一个代码包中的代码所访问。
    * 关键字不能作为标识符。

### 变量和常量
    * 声明变量的关键字var，声明常量的关键字const。
    * 常量只能被赋予基本数据类型的值本身。
    * var num1 int = 1  // 普通赋值： var、变量名称、变量类型、=、变量值
    * var num2,num3 int = 2,3 //平行赋值
    * var (
    *     num4 int = 4
    *     num5 int = 5
    * ) // 多行赋值
    * 上述三种赋值也适用于常量，常量不能只声明不赋值。

### 浮点数类型
    * 浮点数类型有两个，即float32和float64。存储这两个类型的值的空间分别需要4个字节和8个字节。
    * 浮点数类型的值一般由整数部分、小数点“.”和小数部分组成。
    * 指数部分由“E”或“e”以及一个带正负号的10进制数组成。比如，3.7E-2表示浮点数0.037。又比如，3.7E+1表示浮点数37
    * 有一点需要注意，在Go语言里，浮点数的相关部分只能由10进制表示法表示，而不能由8进制表示法或16进制表示法表示。

### 数组
    * 一个数组（Array）就是一个可以容纳若干类型相同的元素的容器。这个容器的大小（即数组的长度）是固定的，且是体现在数组的类型字面量之中的。
    * type MyNumbers [3]int  注：类型声明语句由关键字type、类型名称和类型字面量组成。用于表示某个类型的值的字面表示可被称为值字面量
    * var numbers = [3]int{1, 2, 3} 这是一条变量声明语句。它在声明变量的同时为该变量赋值。
    * var numbers = [...]int{1, 2, 3} 在其中的类型字面量中省略代表其长度的数字
    * numbers[0] // 会得到第一个元素
    * var length = len(numbers)  len是Go语言的内建函数的名称。该函数用于获取字符串、数组、切片、字典或通道类型的值的长度。
    * var numbers2 [5]int 只声明一个数组类型的变量而不为它赋值，那么该变量的值将会是指定长度的、其中各元素均为元素类型的零值。[5]int{0, 0, 0, 0, 0}

### 切片类型
    * 切片（Slice）与数组一样，也是可以容纳若干类型相同的元素的容器。与数组不同的是，无法通过切片类型来确定其值的长度。
    * 表示切片类型的字面量如 []int；
    * 对一个切片类型的声明可以这样：type MySlice []int，MySlice即为切片类型[]int的一个别名类型。 对切片值的表示也与数组值也极其相似，如：[]int{1, 2, 3}
    * 实施切片操作的方式就是切片表达式。举例如下：
    * var numbers3 = [5]int{1, 2, 3, 4, 5}
    * var slice1 = numbers3[1:4]  切片表达式numbers3[1:4]的求值结果为[]int{2, 3, 4}。实际上，slice1这个切片值的底层数组正是numbers3的值。
    * 在一个切片值上实施切片操作 var slice2 = slice1[1:3]  slice2的值为[]int{3, 4}
    * 作为切片表达式求值结果的切片值的长度总是为元素上界索引与元素下界索引的差值。
    * 一个切片值的容量即为它的第一个元素值在其底层数组中的索引值与该数组长度的差值的绝对值。
    *    为了获取数组、切片或通道类型的值的容量，我们可以使用内建函数cap，如：var capacity2 int = cap(slice2)
    * 切片类型属于引用类型。它的零值即为nil。

### 字典类型
    * Go语言的字典（Map）类型其实是哈希表（Hash Table）的一个实现。字典用于存储键-元素对（更通俗的说法是键-值对）的无序集合。
    *     注意，同一个字典中的每个键都是唯一的。
    * 字典类型的字面量如下：map[K]T， K”意为键的类型，而“T”则代表元素（或称值）的类型。如： map[int]string。字典的键类型必须是可比较的，否则会引起错误。
    * 以字典类型map[int]string为例，它的值的字面量可以是这样的： map[int]string{1: "a", 2: "b", 3: "c"}
    * mm := map[int]string{1: "a", 2: "b", 3: "c"}  运用索引表达式取出字典中的值，
    *     就像这样：b := mm[2] 放入方括号中的不再是索引值，而是与我们要取出的值对应的那个键。
    * e, ok := mm[5] 针对字典的索引表达式可以有两个求值结果。第二个求值结果是bool类型的。它用于表明字典值中是否存在指定的键值对。不存在键5，变量ok必为false
    * 从字典中删除键值对的方法非常简单，仅仅是调用内建函数delete而已，就像这样：delete(mm, 4)
    * 与切片类型相同，字典类型属于引用类型。它的零值即为nil

### 通道类型
    * 通道（Channel）是Go语言中一种非常独特的数据结构。它可用于在不同Goroutine之间传递类型化的数据，并且是并发安全的。
    *     Goroutine（也称为Go程序）可以被看做是承载可被并发执行的代码块的载体。
    *     它们由Go语言的运行时系统调度，并依托操作系统线程（又称内核线程）来并发地执行其中的代码块。
    * 通道类型的表示方法很简单，仅由两部分组成，如下：chan T  左边是代表通道类型的关键字chan，而右边代表该通道类型允许传递的数据的类型。
    * 无法用字面量来为通道类型的变量赋值。只能通过调用内建函数make来达到目的。make函数可接受两个参数。
    *     第一个参数是代表了将被初始化的值的类型的字面量，而第二个参数则是值的长度。make(chan int, 5)
    * 声明一个通道类型的变量，并为其赋值： ch1 := make(chan string, 5)
    * 可以使用接收操作符<-向通道值发送数据。 ch1 <- "value1"
    * 从ch1那里接收字符串。 value := <- ch1
    * value, ok := <- ch1。这样做的目的同样是为了消除与零值有关的歧义。这里的变量ok的值同样是bool类型的。
    *     它代表了通道值的状态，true代表通道值有效，而false则代表通道值已无效（或称已关闭）。
    * 关闭通道值，可以通过调用内建函数close来达到目的，就像这样：close(ch1) 通道值的重复关闭会引发运行时恐慌。这会使程序崩溃。
    * 通道类型属于引用类型。它的零值即为nil
    * 通道有带缓冲和非缓冲之分。缓冲通道中可以缓存N个数据。我们在初始化一个通道值的时候必须指定这个N。相对的，非缓冲通道不会缓存任何数据。
    *     发送方在向通道值发送数据的时候会立即被阻塞，直到有某一个接收方已从该通道值中接收了这条数据。
    * 非缓冲的通道值的初始化方法如下：make(chan int, 0)
    * 我们还可以以数据在通道中的传输方向为依据来划分通道。默认情况下，通道都是双向的，即双向通道。
    *     如果数据只能在通道中单向传输，那么该通道就被称作单向通道。
    * type Receiver <-chan int 类型Receiver代表了一个只可从中接收数据的单向通道类型。这样的通道也被称为接收通道。
    * type Sender chan<- int  声明一个发送通道类型
    * 单向通道的主要作用是约束程序对通道值的使用方式。比如，我们调用一个函数时给予它一个发送通道作为参数，以此来约束它只能向该通道发送数据。

### 函数
    * 在Go语言中，函数是一等（first-class）类型。这意味着，我们可以把函数作为值来传递和使用。
    * 函数代表着这样一个过程：它接受若干输入（参数），并经过一些步骤（语句）的执行之后再返回输出（结果）。特别的是，Go语言中的函数可以返回多个结果。
    * 函数类型的字面量由关键字func、由圆括号包裹参数声明列表、空格以及可以由圆括号包裹的结果声明列表组成。 func(input1 string ,input2 string) string
    * 函数类型声明 type MyFunc func(input1 string ,input2 string) string
    * 编写函数的时候需要先写关键字func和函数名称，后跟参数声明列表和结果声明列表，最后是由花括号包裹的语句列表。
    * func myFunc(part1 string, part2 string) (result string) {
    *     result = part1 + part2
    *     return
    * }
    * 如果结果声明是带名称的，那么它就相当于一个已被声明但未被显式赋值的变量。我们可以为它赋值且在return语句中省略掉需要返回的结果值。
    * 函数myFunc是函数类型MyFunc的一个实现。
    * 实际上，只要一个函数的参数声明列表和结果声明列表中的数据类型的顺序和名称与某一个函数类型完全一致，前者就是后者的一个实现。
    * var splice func(string, string) string // 等价于 var splice MyFunc
    * 然后把函数myFunc赋给它：splice = myFunc
    * 实施调用动作 splice("1", "2")
    * var splice = func(part1 string, part2 string) string {
    *     return part1 + part2
    * }
    * 在这个示例中，我们直接使用了一个匿名函数来初始化splice变量。
    * var result = func(part1 string, part2 string) string {
    *     return part1 + part2
    * }("1", "2")
    * 这里的result变量的类型不是函数类型，而与后面的匿名函数的结果类型是相同的。
    * 函数类型的零值是nil。

### 结构体和方法
    * Go语言的结构体类型（Struct）比函数类型更加灵活。它可以封装属性和操作。
    * 结构体类型的字面量由关键字type、类型名称、关键字struct，以及由花括号包裹的若干字段声明组成。
    *     其中，每个字段声明独占一行并由字段名称（可选）和字段类型组成。
    * type Person struct {
    *  Name   string
    *  Gender string
    *  Age    uint8
    * }
    * 用字面量创建出一个该类型的值 Person{Name: "Robert", Gender: "Male", Age: 33}
    * 如果这里的键值对的顺序与其类型中的字段声明完全相同的话 Person{"Robert", "Male", 33}
    * 匿名结构体
    * p := struct {
    *     Name   string
    *     Gender string
    *     Age    uint8
    * }{"Robert", "Male", 33}
    * 匿名结构体最大的用处就是在内部临时创建一个结构以封装数据，而不必正式为其声明相关规则。而在涉及到对外的场景中，我强烈建议使用正式的结构体类型。
    * 结构体类型可以拥有若干方法。所谓方法，其实就是一种特殊的函数。它可以依附于某个自定义类型。
    * 方法的特殊在于它的声明包含了一个接收者声明。这里的接收者指代它所依附的那个类型。
    * func (person *Person) Grow() {
    *     person.Age++
    * }
    * 在关键字func和名称Grow之间的那个圆括号及其包含的内容就是接收者声明。
    * 其中的内容由两部分组成。第一部分是代表它依附的那个类型的值的标识符。第二部分是它依附的那个类型的名称。
    * 后者表明了依附关系，而前者则使得在该方法中的代码可以使用到该类型的值
    * p := Person{"Robert", "Male", 33}
    * p.Grow()
    * 需要注意的是，在Grow方法的接收者声明中的那个类型是*Person，而不是Person。实际上，前者是后者的指针类型。
    * 结构体类型属于值类型。它的零值并不是nil，而是其中字段的值均为相应类型的零值的值。

### 接口
    * 在Go语言中，一个接口类型总是代表着某一种类型（即所有实现它的类型）的行为。
    * 一个接口类型的声明通常会包含关键字type、类型名称、关键字interface以及由花括号包裹的若干方法声明。示例如下：
    * type Animal interface {
    *     Grow()
    *     Move(string) string
    * }
    * 接口类型中的方法声明是普通的方法声明的简化形式。
    * 它们只包括方法名称、参数声明列表和结果声明列表。其中的参数的名称和结果的名称都可以被省略。不过，出于文档化的目的，我还是建议大家在这里写上它们。
    * Move(new string) (old string)
    * 如果一个数据类型所拥有的方法集合中包含了某一个接口类型中的所有方法声明的实现，那么就可以说这个数据类型实现了那个接口类型。
    * 空接口类型即是不包含任何方法声明的接口类型，用interface{}表示，常简称为空接口。
    * 正因为空接口的定义，Go语言中的包含预定义的任何数据类型都可以被看做是空接口的实现。
    * 判定*Person类型实现了Animal接口，在Go语言中，可以用类型断言来实现。
    * 不能在一个非接口类型的值上应用类型断言来判定它是否属于某一个接口类型的。我们必须先把前者转换成空接口类型的值。
    * *Person类型转换成空接口类型的值，就像这样：
    * p := Person{"Robert", "Male", 33, "Beijing"}
    * v := interface{}(&p)
    * 在类型字面量后跟由圆括号包裹的值就构成了一个类型转换表达式，意为将后者转换为前者类型的值。
    * 我们就可以在v上应用类型断言了，即：
    * h, ok := v.(Animal)
    * 类型断言表达式v.(Animal)的求值结果可以有两个。第一个结果是被转换后的那个目标类型（这里是Animal）的值，而第二个结果则是转换操作成功与否的标志。

### if语句
    * Go语言的流程控制主要包括条件分支、循环和并发。条件表达式是指其结果类型是bool的表达式。
    * if语句一般会由关键字if、条件表达式和由花括号包裹的代码块组成。
    * var number int
    * if number := 4; 100 > number {
    *     number += 3
    * }
    * 变量赋值直接加入到if子句中。被叫做if语句的初始化子句。它应被放置在if关键字和条件表达式之间，并与前者由空格分隔
    * 标识符的重声明。if语句内部对number的访问和赋值都只会涉及到第二次声明的那个number变量。这种现象也被叫做标识符的遮蔽。

### switch语句
    * switch语句提供了一个多分支条件执行的方法。case代表分支。每一个case可以携带一个表达式或一个类型说明符。
    * ase表达式的结果类型需要与switch表达式的结果类型一致。
    * var name string
    * // 省略若干条语句
    * switch name {
    * case "Golang":
    *     fmt.Println("A programming language from Google.")
    * case "Rust":
    *     fmt.Println("A programming language from Mozilla.")
    * default:
    *     fmt.Println("Unknown!")
    * }
    * 与if语句一样，switch语句还可以包含初始化子句，且其出现位置和写法也如出一辙。
    * 类型switch语句。紧随case关键字的不是表达式，而是类型说明符。
    * 类型说明符由若干个类型字面量组成，且多个类型字面量之间由英文逗号分隔。
    * switch表达式是非常特殊的。这种特殊的表达式也起到了类型断言的作用，但其表现形式很特殊，如：v.(type)
    * v必须代表一个接口类型的值。该类表达式只能出现在类型switch语句中，且只能充当switch表达式。
    * v := 11
    * switch i := interface{}(v).(type) {
    * case int, int8, int16, int32, int64:
    *     fmt.Printf("A signed integer: %d. The type is %T. \n", i, i)
    * case uint, uint8, uint16, uint32, uint64:
    *     fmt.Printf("A unsigned integer: %d. The type is %T. \n", i, i)
    * default:
    *     fmt.Println("Unknown!")
    * }
    * fallthrough。它既是一个关键字，又可以代表一条语句。
    * fallthrough语句可被包含在表达式switch语句中的case语句中。它的作用是使控制权流转到下一个case。
    * fallthrough语句仅能作为case语句中的最后一条语句出现。并且，包含它的case语句不能是其所属switch语句的最后一条case语句。

### for语句
    * for语句代表着循环。一条语句通常由关键字for、初始化子句、条件表达式、后置子句和以花括号包裹的代码块组成。
    * for i := 0; i < 10; i++ {
    *     fmt.Print(i, " ")
    * }
    * range子句包含一个或两个迭代变量（用于与迭代出的值绑定）、特殊标记:=或=、关键字range以及range表达式。
    * for i, v := range "Go语言" {
    *     fmt.Printf("%d: %c\n", i, v)
    * }
    * 对于字符串类型的被迭代值来说，for语句每次会迭代出两个值。
    *     第一个值代表第二个值在字符串中的索引，而第二个值则代表该字符串中的某一个字符。迭代是以索引递增的顺序进行的。
    * 对于数组值、数组的指针值和切片之来说，range子句每次也会迭代出两个值。
    *     其中，第一个值会是第二个值在被迭代值中的索引，而第二个值则是被迭代值中的某一个元素。
    * 对于字典值来说，range子句每次仍然会迭代出两个值。
    *     显然，第一个值是字典中的某一个键，而第二个值则是该键对应的那个值。
    * 携带range子句的for语句还可以应用于一个通道值之上。其作用是不断地从该通道值中接收数据，不过每次只会接收一个值。
    *     注意，如果通道值中没有数据，那么for语句的执行会处于阻塞状态。无论怎样，这样的循环会一直进行下去。直至该通道值被关闭，for语句的执行才会结束。
    * break语句和continue语句。它们都可以被放置在for语句的代码块中。
    *     前者被执行时会使其所属的for语句的执行立即结束，
    *     而后者被执行时会使当次迭代被中止（当次迭代的后续语句会被忽略）而直接进入到下一次迭代。

### select语句
    * select语句属于条件分支流程控制方法，不过它只能用于通道。
    * 它可以包含若干条case语句，并根据条件选择其中的一个执行。
    * select语句中的case关键字只能后跟用于通道的发送操作的表达式以及接收操作的表达式或语句。
    * ch1 := make(chan int, 1)
    * ch2 := make(chan int, 1)
    * // 省略若干条语句
    * select {
    * case e1 := <-ch1:
    *     fmt.Printf("1th case is selected. e1=%v.\n", e1)
    * case e2 := <-ch2:
    *     fmt.Printf("2th case is selected. e2=%v.\n", e2)
    * default:
    *     fmt.Println("No data!")
    * }
    * 如果该select语句被执行时通道ch1和ch2中都没有任何数据，那么肯定只有default case会被执行。
    * 对于包含通道接收操作的case来讲，其执行条件就是通道中存在数据
    * 如果在当时有数据的通道多于一个，那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。
    * 如果一条select语句中不存在default case， 并且在被执行时其中的所有case都不满足执行条件，那么它的执行将会被阻塞！
    *     当前流程的进行也会因此而停滞。
    *     直到其中一个case满足了执行条件，执行才会继续。
    * break语句也可以被包含在select语句中的case语句中。
    *     它的作用是立即结束当前的select语句的执行，不论其所属的case语句中是否还有未被执行的语句。

### defer语句
    * defer语句仅能被放置在函数或方法中。它由关键字defer和一个调用表达式组成。
    * 这里的调用表达式所表示的既不能是对Go语言内建函数的调用也不能是对Go语言标准库代码包unsafe中的那些函数的调用。
    * func readFile(path string) ([]byte, error) {
    *     file, err := os.Open(path)
    *     if err != nil {
    *         return nil, err
    *     }
    *     defer file.Close()
    *     return ioutil.ReadAll(file)
    * }
    * defer它的确切的执行时机是在其所属的函数（这里是readFile）的执行即将结束的那个时刻。
    * 无论readFile函数正常地返回了结果还是由于在其执行期间有运行时恐慌发生而被剥夺了流程控制权，其中的file.Close()都会在该函数即将退出那一刻被执行。
    * 当一个函数中存在多个defer语句时，它们携带的表达式语句的执行顺序一定是它们的出现顺序的倒序。

### 异常处理——error
    * error是Go语言内置的一个接口类型
    * type error interface {
    *     Error() string
    * }
    * 只要一个类型的方法集合包含了名为Error、无参数声明且仅声明了一个string类型的结果的方法，就相当于实现了error接口。
    * 我们需要创造出错误（即error类型的值）并把它传递给上层程序。这很简单。只需调用标准库代码包errors的New函数即可。
    * if path == "" {
    *     return nil, errors.New("The parameter is invalid!")
    * }

### 异常处理——panic
    * panic可被意译为运行时恐慌
    * 它只有在程序运行的时候才会被“抛出来”。并且，恐慌是会被扩散的。如果我们不显式地处理它的话，程序崩溃。
    * 内建函数panic可以让我们人为地产生一个运行时恐慌。
    * 在Go语言中，内建函数recover就可以做到panic被恢复。recover函数必须要在defer语句中调用才有效。
    *     因为一旦有运行时恐慌发生，当前函数以及在调用栈上的所有代码都是失去对流程的控制权。
    *     只有defer语句携带的函数中的代码才可能在运行时恐慌迅速向调用栈上层蔓延时“拦截到”它。
    * defer func() {
    *     if p := recover(); p != nil {
    *         fmt.Printf("Fatal error: %s\n", p)
    *     }
    * }()
    * 调用recover函数。该函数会返回一个interface{}类型的值。如果p不为nil，那么就说明当前确有运行时恐慌发生。
    * 运行时恐慌代表程序运行过程中的致命错误。我们只应该在必要的时候引发它。


