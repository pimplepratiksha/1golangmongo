package main
import (
	"fmt"
	"os"
	//"io/ioutil"
	"strings"
	//"encoding/json"
	//"bufio"
	dbrepo "../assignment1/dbrepository"
	mongoutils "../assignment1/utils"
	domain "../assignment1/domain"
)
func main() {
	mongoSession, err1 := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	fmt.Println(mongoSession, err1)
	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)
	fmt.Println(repoaccess)

	var ip string
	var final []*domain.Restaurant 
	if len(os.Args)==1{
		cnt,_:=repoaccess.Insert("r.json")
		fmt.Print("Records inserted - ",cnt)
	}else if len(os.Args)>1{
		if(os.Args[1]=="find"){
			ip=os.Args[2]
			arr:=strings.Split(ip,"=")
			switch(arr[0]){
			case "--type_of_food":
				final,_=repoaccess.FindByTypeOfFood(arr[1])
				fmt.Println(final)
			case "--postcode":
				final,_=repoaccess.FindByTypeOfPostCode(arr[1])
				fmt.Println(final)
			default:
				fmt.Println("invalid argument")
				return 
			}
			for _,z:=range final {
				fmt.Println(z)
			}
		}/*else if(os.Args[1]=="count"){
			ip=os.Args[2]
			arr:=strings.Split(ip,"=")
			switch(arr[0]){
			case "--type_of_food":
				fcnt,_:=repoaccess.CountByTypeOfFood(arr[1])
				fmt.Println(fcnt)
			default:
				fmt.Println("invalid argument")
				return 
			}
		}*/

	}
}
