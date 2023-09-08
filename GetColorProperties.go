package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)


const(
	word 	string ="The random color is:"
	apiUrl 	string = "https://www.thecolorapi.com/"

)
var (
	color string
	Uri string
	
)

 type RGBColor struct{
	//rgb color
 	r int 
 	g int 
 	b int 
}

type ColorInfo struct {
	Hex struct {
		Value string `json:"value"`
		Clean string `json:"clean"`
	} `json:"hex"`
	RGB struct {
		Fraction struct {
			R float64 `json:"r"`
			G float64 `json:"g"`
			B float64 `json:"b"`
		} `json:"fraction"`
		R     int    `json:"r"`
		G     int    `json:"g"`
		B     int    `json:"b"`
		Value string `json:"value"`
	} `json:"rgb"`
	HSL struct {
		Fraction struct {
			H float64 `json:"h"`
			S float64 `json:"s"`
			L float64 `json:"l"`
		} `json:"fraction"`
		H     int    `json:"h"`
		S     int    `json:"s"`
		L     int    `json:"l"`
		Value string `json:"value"`
	} `json:"hsl"`
	HSV struct {
		Fraction struct {
			H float64 `json:"h"`
			S float64 `json:"s"`
			V float64 `json:"v"`
		} `json:"fraction"`
		H     int    `json:"h"`
		S     int    `json:"s"`
		V     int    `json:"v"`
		Value string `json:"value"`
	} `json:"hsv"`
	Name struct {
		Value             string `json:"value"`
		ClosestNamedHex   string `json:"closest_named_hex"`
		ExactMatchName    bool   `json:"exact_match_name"`
		Distance          int    `json:"distance"`
	} `json:"name"`
	CMYK struct {
		Fraction struct {
			C float64 `json:"c"`
			M float64 `json:"m"`
			Y float64 `json:"y"`
			K float64 `json:"k"`
		} `json:"fraction"`
		Value string `json:"value"`
		C     int    `json:"c"`
		M     int    `json:"m"`
		Y     int    `json:"y"`
		K     int    `json:"k"`
	} `json:"cmyk"`
	XYZ struct {
		Fraction struct {
			X float64 `json:"X"`
			Y float64 `json:"Y"`
			Z float64 `json:"Z"`
		} `json:"fraction"`
		Value string `json:"value"`
		X     int    `json:"X"`
		Y     int    `json:"Y"`
		Z     int    `json:"Z"`
	} `json:"XYZ"`
	Image struct {
		Bare  string `json:"bare"`
		Named string `json:"named"`
	} `json:"image"`
	Contrast struct {
		Value string `json:"value"`
	} `json:"contrast"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Embedded struct{} `json:"_embedded"`
}





func assignColor() RGBColor{
	return RGBColor{
		r :rand.Intn(255),
		g :rand.Intn(255),
		b :rand.Intn(255),
	}
		
}


func main(){
	GenerateValues:=assignColor()
	color:= strconv.Itoa(GenerateValues.r)+","+strconv.Itoa(GenerateValues.g)+","+strconv.Itoa(GenerateValues.b)

	Uri:= apiUrl+"id?rgb="+color+"&format=json"
	

	resp, err:= http.Get(Uri)
	if err!= nil{
		fmt.Println("Error get")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK {
		fmt.Println("Failed with code:",resp.StatusCode)	
		return
	}

	body,err:= io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error",err)
		return
		
	}

	var thiscolor ColorInfo
	json.Unmarshal(body,&thiscolor)
	fmt.Println(thiscolor.Name)
	//add case for exact name eq false :3
			

}