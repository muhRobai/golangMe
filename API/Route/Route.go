package main

import(
	"encoding/json"
	"net/http"
    "fmt"
    "strings"
    "sort"
)

type data struct{
    textAwal string
    hasilSort string
}

func Sorts(text string)string{
    var role =[]string{"a","i","u","e","o","A","I","U","E","O"}
    var exceptRole = "bcdfghjklmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ"
    var hasil = []string{}
    var hasil2 = []string{}
    var txtArray = strings.Split(exceptRole,"")
    var exceptRest = 0
   
    for j := 0; j < len(role); j++ {
        if strings.ContainsAny(text, role[j]) {
            exceptRest = strings.Count(text,role[j])
            for k := 0; k < exceptRest; k++ {
                hasil = append(hasil,role[j] )
            }
            sort.Strings(hasil)
        }
    }

    for i := 0; i < len(txtArray); i++ {
        if strings.ContainsAny(text, txtArray[i]) {
            exceptRest = strings.Count(text,txtArray[i])
            for k := 0; k < exceptRest; k++ {
                hasil2 = append(hasil2,txtArray[i])
            }
            sort.Strings(hasil2)
        }
    }

    for l := 0; l < len(hasil2); l++ {
        hasil = append(hasil, hasil2[l])
    }
    return strings.Join(hasil,"")
}


func Shorter(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "POST" {
        var text = r.FormValue("text")

        data := []string{text,Sorts(text)}

        result, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
    }
}

func main() {

    http.HandleFunc("/sort", Shorter)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}