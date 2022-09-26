package accesslog

import "fmt"

func ExampleEmpty() {
	e := CreateVersionedEntity()
	fmt.Printf("Valid index : %v\n", e.index == 0)
	fmt.Printf("Hash : %v\n", e.getState().hash)
	fmt.Printf("IsNewVersion [] : %v\n", e.IsNewVersion(""))
	fmt.Printf("IsNewVersion [1] : %v\n", e.IsNewVersion("1"))
	fmt.Printf("IsEmptyEntity : %v\n", e.IsEmpty())

	//Output:
	// Valid index : true
	// Hash : 0
	// IsNewVersion [] : false
	// IsNewVersion [1] : true
	// IsEmptyEntity : true
}

func ExampleSetEntity() {
	e := CreateVersionedEntity()

	// Set entity
	e.SetEntity(&View{Version: "1.2.3"})
	fmt.Printf("New version [] : %v\n", e.IsNewVersion(""))
	fmt.Printf("Index valid : %v\n", e.index == 1)
	fmt.Printf("Hash : %v\n", e.getState().hash)
	fmt.Printf("IsEmptyEntity : %v\n", e.IsEmpty())
	fmt.Printf("Entity : %v\n\n", e.GetEntity())

	// Set entity
	e.SetEntity(&View{Version: "1.2.4"})
	fmt.Printf("New version [1.2.3] : %v\n", e.IsNewVersion("1.2.3"))
	fmt.Printf("Index valid : %v\n", e.index == 0)
	fmt.Printf("Hash : %v\n", e.getState().hash)
	fmt.Printf("IsEmptyEntity : %v\n", e.IsEmpty())
	fmt.Printf("Entity : %v\n\n", e.GetEntity())

	//Output:
	// New version [] : true
	// Index valid : true
	// Hash : 414986927
	// IsEmptyEntity : false
	// Entity : {1.2.3 [] [] [] [] [] []}
	//
	// New version [1.2.3] : true
	// Index valid : true
	// Hash : 364654070
	// IsEmptyEntity : false
	// Entity : {1.2.4 [] [] [] [] [] []}
}
