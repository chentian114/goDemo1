# go 学习练习案例

go语言特点
    1.静态类型、编译型的开源语言
    2.脚本化的语法，支持多种编程范式（函数式&面向对象）
    3.原生、给力的并发编程支持

go语言的优劣势
    优势
    1.脚本化的语法
    2.静态类型+编译型，程序运行速度有保障
    3.原生的扶持并发编程（降低开发、维护成本；程序可以更好的执行）
    劣势
    1.语法糖并没有Python和Ruby那么多
    2.目前的程序运行速度还不及C（已超过C++和Java）
    3.第三方函数库暂时不像绝对主流的编程语言那么多

go源码文件
 .go为后缀，内容以go语言代码组织的文件

go源码文件分类
    1.命令源码文件（声明自己属于main代码包、包含无参数声明和结果声明的main函数）
        同一个代码包中强烈不建议直接包含多个命令源码文件
    2.库源码文件（不具备命令源码文件两个特征的源码文件）
    3.测试源码文件（不具备命令源码文件两个特征的源码文件，名称以_test.go为后缀）
        测试函数以Test（功能测试函数）或Benchmark为前缀（性能测试函数）

代码包的作用
    编译和归档go程序的最基本单位
    代码划分、集结和依赖的有效组织形式，也是权限控制的辅助手段

程序实体与关键字
    在go语言中，变量、常量、函数、结构体和接口被统称为“程序实体”，而它们的名字被统称为“标识符”。
    标识符可以是任何Unicode编码可以表示的字母、数字及下划线。首字母不能是数字或下划线。
    对程序实体的访问权限控制只能通过它们的名字来实现。名字首字母大写的程序实体可以被任何代码包访问到。而名字
首字母小写的程序实体则只能被同一个代码包中的代码所访问。
    关键字不能作为标识符。

变量和常量
    声明变量的关键字var，声明常量的关键字const。
    常量只能被赋予基本数据类型的值本身。
    var num1 int = 1  // 普通赋值： var、变量名称、变量类型、=、变量值
    var num2,num3 int = 2,3 //平行赋值
    var (
        num4 int = 4
        num5 int = 5
    ) // 多行赋值
    上述三种赋值也适用于常量，常量不能只声明不赋值。

