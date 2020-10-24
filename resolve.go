package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chi/chi"
)

func resolve(w http.ResponseWriter, r *http.Request) {
	type RawDid struct {
		DID interface{} `json:"did"`
	}
	DIDstr := chi.URLParam(r, "DID")
	_, ok := getValidID(DIDstr)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"requested identifier is not a JLINC DID"}`))
		return
	}

	url := fmt.Sprintf("https://testnet.did.jlinc.org/%s", DIDstr)
	resp, err := http.Get(url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "error: upstream responded with %s", err)
		return
	}
	if resp.StatusCode == 410 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		fmt.Fprintf(w, `{"response":"revoked"}`)
		return
	}
	if resp.StatusCode != 200 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		fmt.Fprintf(w, `{"response":%q}`, http.StatusText(resp.StatusCode))
		return
	}

	fmt.Printf("Status: %v\n", resp.Status)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	var raw RawDid
	json.Unmarshal(body, &raw)
	did, err := json.MarshalIndent(raw.DID, "", "    ")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/did+ld+json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", did)
}

func getValidID(id string) (string, bool) {
	idParts := strings.Split(id, ":")
	idRxp := regexp.MustCompile(`^[\w\-]+$`) //base64 or base58 string
	if len(idParts) == 3 && idParts[0] == "did" && idParts[1] == "jlinc" && idRxp.MatchString(idParts[2]) {
		return idParts[2], true
	}
	return "", false
}
