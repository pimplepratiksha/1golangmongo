package main
import (
	"fmt"
	"os"
	"strings"
	dbrepo "assignment1/dbrepository"
	mongoutils "assignment1/utils"
	domain "assignment1/domain"
)
func main() {
	mongoSession, err1 := mongoutils.RegisterMongoSession(os.Getenv("MONGO_HOST"))
	fmt.Println(mongoSession, err1)
	dbname := "restaurant"
	repoaccess := dbrepo.NewMongoRepository(mongoSession, dbname)

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
			case "--postcode":
				final,_=repoaccess.FindByTypeOfPostCode(arr[1])
			default:
				fmt.Println("invalid argument")
				return 
			}
			for _,z:=range final {
				fmt.Println(z)
			}
		}else if(os.Args[1]=="count"){
			ip=os.Args[2]
			arr:=strings.Split(ip,"=")
			switch(arr[0]){
			case "--type_of_food":
				fcnt,_:=repoaccess.CountByTypeOfFood(arr[1])
				fmt.Println(fcnt)
			case "--postcode":
				fcnt,_:=repoaccess.CountByTypeOfPostCode(arr[1])
				fmt.Println(fcnt)
			default:
				fmt.Println("invalid argument")
				return 
			}
		}else if(os.Args[1]=="search"){	
			query:=os.Args[2]
			final,_:=repoaccess.Search(query)
			for _,z:=range final {
				fmt.Println(z)
			}			
			}
	}
}
