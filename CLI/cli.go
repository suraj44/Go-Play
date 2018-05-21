package main
import(
	"fmt"
	"os"
	"strings"
	flag "github.com/ogier/pflag"
	
)

var (
	user string
)

func main() {

	flag.Parse()


	//if user doesn't specify any flags, run this logic
	if flag.NFlag() ==0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	users := strings.Split(user,  ",")
	fmt.Printf("Searching user(s): %s\n", users)

	for _, u := range users {
		result :=  getUsers(u)
		fmt.Println(`Username:	`, result.Login)
		fmt.Println(`Name:		`, result.Name)
		fmt.Println(`Email:		`, result.Email)
		fmt.Println(`Bio:		`, result.Bio)
		fmt.Println("")

	}
}


//init is a special function in Go
//It executes it after importing but before main()
//We're expecting a Posix compliant flag(single or double dash)
//StringVarP(The variable we're binding this flag to, double dash flag, single dash flag, default value if no flag, description of flag)
func init() {
	flag.StringVarP(&user, "user", "u" ,"", "Search Users")
}

