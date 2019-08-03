package Sort

import(
    "strings"
    "sort"
)

func Sort(text string) string{
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