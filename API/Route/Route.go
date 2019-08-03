package main

import(
	"encoding/json"
	"net/http"
    "fmt"
    "API/Controller/Sort"
    "API/Controller/Sum"
)

type data struct{
    textAwal string 
    hasilSort string 
}


type datas struct{
    hidup int
    mati int
}

func Sums(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "POST" {
        var text = r.FormValue("text")
        if text == "" {
            fmt.Println("text is Empty")
            w.WriteHeader(500)
            return
        }
        sumHidup, sumMati := Sum.Sum(text)

        datas := []int{sumHidup, sumMati}

        result, err := json.Marshal(datas)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
        w.Write(result)
		return

    }
}

func Shorter(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "POST" {
        var text = r.FormValue("text")

        if text == "" {
            fmt.Println("text is Empty")
            w.WriteHeader(500)
            return
        }

        data := []string{text,Sort.Sort(text)}

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
    http.HandleFunc("/sum", Sums)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}