package patrol

import (
	"strconv"
	"encoding/json"
	"net/http"
	"webcenter/model"
)

type patrolLineController struct {
}

type Result struct {
	Result string
	Route model.Routeline
}

func (this *patrolLineController) QueryAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Header().Set("charset", "utf-8")
	
	id, err := strconv.Atoi(r.URL.RawQuery)
	if err != nil {
		id = 1
	}
	
	var result string
	routeline,found := model.GetRouline(id)
	if found {
		result = "OK"
	} else {
		result = "NOK"
	}
		
	resultInfo := Result{Result:result,Route:routeline}	
	b, err := json.Marshal(&resultInfo)
    if err != nil {
        return
    }
    w.Write(b)
}
