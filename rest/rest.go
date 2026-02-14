package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hellicopthecat/learngo/blockchain"
	"github.com/hellicopthecat/learngo/utils"
)

var port string

type URL string

func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"desc"`
	Payload     string `json:"payload"`
	// Payload     string `json:"payload,omitempty"`
	HelloTxt string `json:"helloTxt"`
}
type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMsg string `json:"errorMsg"`
}

// 자바 toString 같은
func (u urlDescription) String() string {
	return u.HelloTxt
}

func documentaion(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
			HelloTxt:    "what the fuc",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Post blocks",
		},
		{
			URL:         URL("/blocks/{height}"),
			Method:      "POST",
			Description: "Post blocks",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	h, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().FindBlock(h)
	if err == blockchain.ErrorNotFound {
		json.NewEncoder(rw).Encode(errorResponse{fmt.Sprint(err)})
	} else {
		json.NewEncoder(rw).Encode(block)
	}
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	handler := mux.NewRouter()
	handler.Use(jsonContentTypeMiddleware)
	handler.HandleFunc("/", documentaion).Methods("GET")
	handler.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	handler.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	port = fmt.Sprintf(":%d", aPort)

	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

/*
mux 는 Multiplexer이다.
Multiplexer는 url로 request를 다루는 건데
url을 지켜보고 내가 원하는 함수를 실행한다.
ListenAndServe를 동시에 사용하면 서버가 실행이 되지 않는다.
이때 핸들러에 serverMux를 지정해주면 해결됨
*/
