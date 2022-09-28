package accesslog

import "fmt"

func ExampleTokenize() {
	tokens := tokenize("")
	fmt.Printf("Tokens [] : %v\n", tokens)

	tokens = tokenize(",")
	fmt.Printf("Tokens [,] : %v\n", tokens)

	tokens = tokenize(",,,")
	fmt.Printf("Tokens [,,,] : %v\n", tokens)

	tokens = tokenize(",start_time")
	fmt.Printf("Tokens [,start_time] : %v\n", tokens)

	tokens = tokenize(",start_time,,duration")
	fmt.Printf("Tokens [,start_time,,duration] : %v\n", tokens)

	//Output:
	// Tokens [] : []
	// Tokens [,] : []
	// Tokens [,,,] : []
	// Tokens [,start_time] : [start_time]
	// Tokens [,start_time,,duration] : [start_time duration]
}

func ExampleCreateEntity() {
	ingress := CSVAttributes{App: "", Custom: "", RequestHeaders: "", ResponseHeaders: "", ResponseTrailers: "", Cookies: ""}
	egress := CSVAttributes{App: "", Custom: "", RequestHeaders: "", ResponseHeaders: "", ResponseTrailers: "", Cookies: ""}
	view := CreateEntity(&ingress, &egress)
	fmt.Printf("Ingress View : %v\n", view.Ingress)
	fmt.Printf("Egress View : %v\n", view.Egress)

	ingress2 := CSVAttributes{App: "ingress_log_attr", Custom: "ingress_custom", RequestHeaders: "ingress_req_header1,ingress_req_header2", ResponseHeaders: "ingress_resp_header", ResponseTrailers: "ingress_resp_trailer", Cookies: "ingress_cookie"}
	egress2 := CSVAttributes{App: "egress_log_attr", Custom: "egress_custom", RequestHeaders: "egress_req_header1,egress_req_header2", ResponseHeaders: "egress_resp_header", ResponseTrailers: "egress_resp_trailer", Cookies: "egress_cookie1,egress_cookie2"}
	view2 := CreateEntity(&ingress2, &egress2)
	fmt.Printf("Ingress View : %v\n", view2.Ingress)
	fmt.Printf("Egress View : %v\n", view2.Egress)

	//Output:
	// Ingress View : {[] [] [] [] [] []}
	// Egress View : {[] [] [] [] [] []}
	// Ingress View : {[ingress_log_attr] [ingress_custom] [ingress_req_header1 ingress_req_header2] [ingress_resp_header] [ingress_resp_trailer] [ingress_cookie]}
	// Egress View : {[egress_log_attr] [egress_custom] [egress_req_header1 egress_req_header2] [egress_resp_header] [egress_resp_trailer] [egress_cookie1 egress_cookie2]}
}

func ExampleVersionedEntityEmpty() {
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

func ExampleVersionedEntitySetEntity() {
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
	// Entity : {1.2.3 {[] [] [] [] [] []} {[] [] [] [] [] []}}
	//
	// New version [1.2.3] : true
	// Index valid : true
	// Hash : 364654070
	// IsEmptyEntity : false
	// Entity : {1.2.4 {[] [] [] [] [] []} {[] [] [] [] [] []}}
}
