package main


import(
	"fmt"
	"net/http"
	"html/template"
	"strings"
	"log"
	"os"
	"net"
)



func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello World")
	
}



func login(w http.ResponseWriter , r *http.Request){
	fmt.Println("Method:",r.Method)
	if(r.Method == "GET"){
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}else
	{
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
	
	
}



func main(){
	if(len(os.Args)==1){
		http.HandleFunc("/",sayHello)
		http.HandleFunc("/login",login)

		err:=http.ListenAndServe(":9090",nil)
		if err != nil {
			log.Fatal("ListenAndServe",err)
		}
		
	}else
	{
		addr := net.ParseIP(os.Args[1])
		if(addr == nil){
			fmt.Println("wrong IP addr");
		}else{
			fmt.Println("IP addr :",addr.String())
		}
			
	}
	
		
}