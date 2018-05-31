package main

import (

	"math"
	"net/http"
	"sort"
	"github.com/gorilla/mux"
	"strconv"
	"io"
	"encoding/json"

	"fmt"
	"log"
)

type Point struct {
	Name string	`json:"name"`
	X    float64	`json:"x"`
	Y    float64	`json:y`
}

type UserPref struct {
	Name     string  `json:"name"`
	Distance float64 `json:"dist"`
}
type UserPrefArr []UserPref

func (v UserPrefArr) Len() int {
	return len(v)
}
func (v UserPrefArr) Less(i, j int) bool {
	return v[i].Distance < v[j].Distance
}
func (v UserPrefArr) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}
func calculatePointToPoint(source, destination Point) float64 {
	x := math.Pow(float64(source.X-destination.X), 2)
	y := math.Pow(float64(source.Y-destination.Y), 2)
	dis := math.Sqrt(x + y)
	return dis

}

func calculateDistanceTopThree(x, y float64) []UserPref {
	p := Point{"user_location", x, y}
	var userList UserPrefArr

	for _, v := range point_arr {
		tempUser := UserPref{v.Name, calculatePointToPoint(p, v)}
		userList = append(userList, tempUser)
	}


	sort.Sort(userList)

	return userList[0:3]

}

func NeedCoffee(w http.ResponseWriter, r *http.Request) {
	map_var := r.URL.Query()
	x, ok := map_var["x"]
	y, ok_y := map_var["y"]

	if ok && ok_y && len(x)==1 && len(y)==1{
		_x,e:=strconv.ParseFloat(x[0],64)
		_y,e:=strconv.ParseFloat(y[0],64)
		if e!=nil{
			log.Println("Handle error of variable not given in api",e)

		}
		str,_:=json.Marshal(calculateDistanceTopThree(_x,_y))
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")


		io.WriteString(w,string(str))

	} else {
		io.WriteString(w,string("Bad Request"))
	}

}

func AddCoffee(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(decoder)
	var t []Point
	err := decoder.Decode(&t)
	if err != nil {
		log.Print(err)
		io.WriteString(w,string("Bad Request"))
	}
	defer r.Body.Close()
	for _,v:=range t{
		ok,_:=in_array(v,point_arr)
		if!ok{
			point_arr = append(point_arr,v)
			io.WriteString(w,string(v.Name+"\tAdded\n"))
		}else{
			io.WriteString(w,string(v.Name+"\tAlready exist\n"))
		}
	}


}

func in_array(val Point, array []Point) (exists bool, index int) {
	exists = false
	index = -1;

	for i, v := range array {
		if val.Name == v.Name &&  val.X == v.X && val.Y == v.Y {
			index = i
			exists = true
			return
		}
	}

	return
}
var point_arr []Point
func intiation(){
	point_arr = append(point_arr, Point{"Starbucks Seattle", 47.5809, -122.3160})
	point_arr = append(point_arr, Point{"Starbucks SF", 37.5209, -122.3340})
	point_arr = append(point_arr, Point{"Starbucks Moscow", 55.752047, 37.595242})
	point_arr = append(point_arr, Point{"Starbucks Seattle 2", 47.5869, -122.3368})
	point_arr = append(point_arr, Point{"Starbucks Rio De Janeiro", -22.923489, -43.234418})
	point_arr = append(point_arr, Point{"Starbucks Sydney", -33.871843, 151.206767})
}
func main() {
	intiation()
	r:=mux.NewRouter()
	r.HandleFunc("/needcoffee",NeedCoffee).Methods("GET")
	r.HandleFunc("/addcoffees",AddCoffee).Methods("POST")
	http.ListenAndServe("localhost:3000", r)

}
