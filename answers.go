package main

func YouCanDoIt() {
	var answers []string

	// ------------------------------------------------
	// BEGINNER ANSWERS
	// ------------------------------------------------

	answers = append(answers,  `The blank identifier is usually used when you want to declare
								a var but not use it. like for example when a functions retuns 
								multiple values but you want to use only one. Or when you import a package 
								but you want to use its side effects only and not its exported vars and funcs, 
								How does it achieve this? Thanks to the blank identifier, only the init funcs are executed`)

	answers = append(answers,  `When you declare a var it means that you are declare the name and type of a var. 
									You are telling the compiler to give you a certain amount of space according to the type
								When you assign a var it means that you are assigning a value to the var
								When you are initializing a var it means you are both declaring and assign to the var`)

	answers = append(answers,  `Allocationg is when you are asking the complier to assign space in memory for a var.
								Initialization is when you allocate a var and assign a value to it as well`)

	answers = append(answers,  `The new keyword allocates memory for you and initialize the vars to zero-value
									and returns the pointer of the var
								the make keyword allocates and initialize to zero-value, but you can also
									specify the length and capacity, and it returns the value of the data structure`)

	answers = append(answers,  `The first one initialize the slice to zero-values, and its length and capacity to second parameter
								The second one initialize the slice to zero-values, its length to 0, and capacity to 10
								The third one is the same one as the first`)

	answers = append(answers,  `You can use as a key all types that are comparable.
								Comparable Types = Every types apart funcs, and slices and arrays 
								that have non-comparable types themself `)

	answers = append(answers,  `To write idiomatic code it means write Go in a certain way / in a standard way.
								Like the comma ok idiom, like short declaration operator`)

	answers = append(answers,  `It depends, but most of time I'll say it better to go for readability,
								because Go its already performant, since it s a compiled language.
								Also simplicity / maintability / readablity is the "slogan" of Golang`)

	answers = append(answers,  `Rob Pike invented, along-side with other google engineers.
								Go was invented to solve scaling problems, and doing it in the simplest way`)

	answers = append(answers,  `Strongly typed means that you can't cast / convert types implicitly. 
								It means you cannot add an integer to a string without expliicitly converting one of them`)

	answers = append(answers,  `The var keyword is used to declare vars. when you want to declare a var and 
								you don't want to assign a value, you would use var`)

	answers = append(answers,  `word size refers the number of bits processed, stored by a computer's processor or memory
								Common word sizes are 16 / 32 /64 bits`)

	answers = append(answers,  `A complier is software program that transalates high-level code into machine code 
								that the computer can understand`)

	answers = append(answers,  `A garbage collrect is a built-in mechanism for automatic memory management.
								It tracks memory usage and reclaims memory that is no longer used by the program .

								Tri-color: Memory is categorized into three colors: white, gray, and black. Initially, 
									all objects are white, indicating they have not been visited yet. 
									When an object is discovered but not yet processed, it's marked as gray. 
									Once it's processed, it's marked as black.

								Mark Phase: The garbage collector starts traversing the object graph, 
									beginning from the roots (e.g., global variables, stack, etc.). 
									It recursively visits objects reachable from the roots, marking them as gray. 
									This process continues until no more gray objects are left,
									meaning all reachable objects are marked.

								Sweep Phase: The garbage collector iterates through all allocated memory, 
									reclaiming memory from objects that are not marked black (i.e., unreachable). 
									It also resets the color of all remaining live objects back to white
									for the next garbage collection cycle.

								The garbage collection in Go triggers automatically when certain conditions are met, 
								such as when the heap size exceeds a predefined threshold, or when the application is idle. 
								The Go runtime system manages the scheduling of garbage collection, 
								aiming to minimize its impact on the performance of the running program.`)

	answers = append(answers,  `Go handle package management though Go Mudules. How does it work?
								In go.mod file you specify the details of your dependecy/module (path, versions and indirect dependecy).
								It automatically downloads the dependecy for you. You also have go.sum file where it stores cryptographic 
								checksums of the dependency, preventing unwanted modifications of the packages.`)

	answers = append(answers,  `Byte is a size unit of data. it's the equivalent of 8-bits. 
								Code points represents a character in the Unicode standard. 
									In Go it s called rune, which is an alias for int32
								UTF-8 is a variable-length encoding for Unicode characters.
									It encodes each Unicode character as one to four bytes
									In Go, string literals are UTF-8 encoded by default.

								When iterating over a string using a for range loop, 
									Go treats each iteration as decoding a UTF-8-encoded rune. 
									This allows for easy iteration over characters in a string, 
									regardless of whether they are single-byte or multi-byte characters.`)

	answers = append(answers,  `Embedding a struct means including a struct inside another one as a field.
									So the outer struct can access the methods/fields of the embedded struct
								Inner-type Promotion is the feature of directly accessing the methods of an embedded struct,
								without needed to give a name to the embedded struct`)

	answers = append(answers, ` remove the person var name

	type human struct {
		name  string
		email string
	}
	
	type secretAgent struct {
		human
		id  string
	}

	`)

	answers = append(answers, ` initialize the map
	
	package main

	import "fmt"
	
	func main() {
		var m make(map[string]int)
		m["b"] = 42
		fmt.Println(m)
	}
	
	`)

	answers = append(answers, `Either declare and initialize xi1 outside the for loop, or xi1 is a syntax error

	func slice() {
		xi := make([]int, 10)
		for i := 0; i < 10; i++ {
			xi = append(xi, i)
		}
		fmt.Println(xi)
	}		
	`)

	answers = append(answers, `
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		// the first two means what indexes, the last (in conjunction with the first) is the length and capacity
		// the last number is exclusive
		xs := sports[1:3:3]
		// this changes the previoues "surf"
		xs[0] = "CHANGED"

		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)

	answers = append(answers, `Fix this code so that when xs[0] is changed this doesn't change 'sports':
	
	func main() {
		sports := make([]string, 5)
		sports[0] = "ski"
		sports[1] = "surf"
		sports[2] = "swim"
		sports[3] = "sail"
		sports[4] = "sumo wrestling"
	
		xs := make(string[], len(sports[1:3]), len(sports[1:3]))
		copy(xs, sports[1:3])
		xs[0] = "CHANGED"
		inspectSlice(sports)
		inspectSlice(xs)
	}
	
	func inspectSlice(xs []string) {
		fmt.Printf("len: %v \ncap: %v \n", len(xs), cap(xs))
		for i := range xs {
			fmt.Printf("%p \t %v \n", &xs[i], xs[i])
		}
	}
	`)

	// ------------------------------------------------
	// INTERMEDIATE ANSWERS
	// ------------------------------------------------

	answers = append(answers,  `The comma ok idiom is a pattern used for type assertions and 
								checking if a value is present in a collection, such as a map or a channel.`)

	answers = append(answers,  `Technically it is not. Go does not inherit the classic OOP patterns.
								it incorporates elements of OOP but also introduces its own unique features and paradigms.
								It does include Encapsulation, Abstraction, and Polymorphism, but not Inheritence`)

	answers = append(answers,  `It refers to how string are represented in memory. 
								A string is composed of two parts: a pointer to the the array of bytes stroed in memory
								and a length of the string`)

	answers = append(answers,  `It refers to how a slice is represented in memory.
								A slice is composed of three parts: a pointer to the array, the length (the number of elements in the slice)
								and the capacity (the maximum number of elements that the slice can hold without requiring reallocation of the underlying array)`)

	answers = append(answers,  `The difference is that the nil slice is not initialized (it's not referring to any underlying array).
								Meanwhile the empty slice is initalized with a length of 0, but it has an underlying array. It is not  nil `)

	answers = append(answers,  `Typed constants have explicit types, while untyped constants don't, 
								but at runtime Go assign a type to them based on thei context`)

	answers = append(answers,  `In Go, errors are represented as values. This means errors can be assigned to variables,
								passed as arguments, and returned from functions just like any other value.
								Go has explicit error handling, forcing developers to handle errors, 
								something that can be ignored in other languages`)

	answers = append(answers,  `Conversion is the act of converting a type to another compatible type. (like int32 to int64)
								Cùasting is the act of converting a type to another, even if not compatible.
								There is no casting in Go by design, since it can lead to unexpected behaviours `)

	answers = append(answers,  `Nil is an identifier, used to represent the absence of a value or reference for certain types.
								The only similarity with "null" is for pointers, meaning that
								nil represnts a zero memory address (a memory space reserved by the OS)`)

	answers = append(answers,  `You need create a go file with suffix "_test.go" and then create a the function 
								with the parameter "* testing.B" and when this command to run the benchmark go test -bench=.`)

	answers = append(answers,  `int can represent negative values, while uint cannot.
								Mechanical sympathy means design software while having ni mind how hardware
								interacts with these types`)

	answers = append(answers,  `Pointers are variables that stores memory address of other variables.
								This allowa indirect referencing and manipulation of data stored in memory.`)

	answers = append(answers,  ``)
}