# Go_Intro-
## What is Golang?
Go or Golang is a language created by Google originally with purpose to resolve internal problems in the company. But eventually this escalated into the development universe. 

Go is created in C++ languaje, but its syntax is a lot more friendly. Was thought to take advantage of the the last advances in hardware, multiprocessors, huge amounts of storage and parallelism.

Is a compiled language and could create .EXE files compatible with all operating systems

Go proved to be an ideal language for large developments with thousands and thousands of concurrent users.

## Conventions
Go Encourages developers to have good practices

- ';' symbol is not necesary
- The compiler show alerts when detect an unused variable or an unused package.
- Go functions could return more than one value!
- To iterate only exist the FOR instruction (does not exist While, Do Until..., or  similar )
- Go is not an OOP language but could resolve it using Structures, Methods and Interfaces
- Scope of variables is determined by names in Capital letters or lower letters

<hr />

### Exercises Notes
#### Exercise 01 - Hello World
Basic Hello World Structure

     package main
      
    import  "fmt"

    func  main() {
    
    fmt.Println("Hi Golang")
    }

> Its important to have the bracket inline with the declaration of te function. Put it below is incorrect.

#### Exercise 02 - Variables
Variables could have }**Global Scope** *( Outsite Functions)* or **Local Scope** *(in functions)*

##### Ways to declare variables
    //var name type
    var number1 int
/
    number2 := 4
    number2 = 5
> Is not allowed to declare two times the same variable by using the two dots:

    var number1, number2, number3 int

    number1, number2, number3, text := 2, 24, 43, "Hello, my name is David"

>You cant reassing or change the type of a variable. For example, you cant convert an int into float

##### Casting variable
    // type(variable)
    int(floatVariable)

For strings you need other functions to it works
    fmt.Sprintf("%d", intVariable)
>The  first parametter is called Verb. Further we are going to see the complete list of verbs

You can also import another package strconv
    strcon.Itoa(int(number)

#### Exercise 3 - Control Structures
##### if statement
    if condition {
		//instructions
	}else if condition2{
		//instructions
	}else{
        //instructions
    }

>In go is possible to reassign a value of a variable in the if instruction.
**An illustrative example:**
    estado:=true
    if estado=false; estado=true {
		//instructions
	}else{
		//instructions
	}

##### switch case 
    switch expression {
	case condition:
		//instruction
    default:
        //instruction
	} 

####  Show and accept data

    fmt.Println("Enter the number: ")
	fmt.Scanf("%d", &variable)

Scanf functions well in linux but it have some issues compiling in windows os. So you can use instead the function Scanln

    fmt.Println("Enter the number: ")
	fmt.Scanln(&variable)

Also you can use the bufio and os packages to resolve this action

    scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		leyenda = scanner.Text()
	}
    fmt.Println(leyenda, resultado)


#### Structures
Una estructura, permite generar una definicion de propiedades que van a tener los objetos

#### Interfaces
Interfaces define conducts, operations, behavior

#### DEFER PANIC RECOVER
##### DEFER
>Es una instruccion que se ejecuta si o si cuando se detecta que una funcion se va por un return o por un error o por un fin de funcion

#### Routines
##### Canal
>Un canal es un espacio de memoria de dialogo en el que van a dialogar distintas rutinas y cuando se aloje un espacio de memoria la rutina que esta pidiendo un valor a cambio va a actuar en consecuencia.  

#### Middlewares
> Son interceptores que permiten ejecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables
