package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"

	solmerkle "github.com/0xKiwi/sol-merkle-tree-go"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


const (
	jsonUrl  = "https://url_to_json.com/data.json"
	jsonPath = "merkle_tree_name.json"
	apiKeyServer = "API_KEY_SERVER"
	apiKeyFront = "API_KEY_FRONT"
)


func WriteDataToFileAsJSON(data interface{}, filedir string) error {
	f, err := os.Create(filedir)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	
	return encoder.Encode(data)

}

type objectJson struct {
    A string `json:"a"`
    V string `json:"v"`
}

func ReadJson(url string) ([]objectJson, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var records []objectJson
	err = json.Unmarshal(body, &records)
	if err != nil {
		return nil, err
	}


	for i := range records {
		v, success := new(big.Float).SetString(records[i].V)
		if !success {
			fmt.Printf("Errore nella conversione del valore 'v' all'indice %d\n", i)
			continue
		}

		ethers := new(big.Float).Mul(v, big.NewFloat(1e18))
		ethersInt, _ := ethers.Int(nil)

		records[i].V = ethersInt.String()
	}

	return records, nil
}


func (d objectJson) Hash() []byte {
	address := common.HexToAddress(d.A)
	value, _ := new(big.Int).SetString(d.V, 10)
	// hexStr := d.A[2:]
	// bytes, _ := hex.DecodeString(hexStr)
	// num := new(big.Int)
	// num.SetBytes(bytes)
	// fmt.Println("Valore convertito:", num)
	
	packed := append(
		common.LeftPadBytes(address.Bytes(), 32),
		common.LeftPadBytes(value.Bytes(), 32)...,
	)

	return crypto.Keccak256(packed)
}


// func ParseToDataMerkleTree(tmp [][]string) []Data {
// 	var data []Data
// 	for _, t := range tmp {
// 		dt := Data{
// 			id: t[0],
// 			priceId:  t[1][1:],
// 		}
// 		data = append(data, dt)
// 	}

// 	return data
// }

type MerkleTree struct {
	//mk merkle.Tree
	mk    *solmerkle.MerkleTree
	leave [][]byte
	clearData [][]string
}

type Proof struct {
	Leaf  string   `json:"leaf"`
	Proof []string `json:"proof"`
	ClearData[]string `json:"clearData"`
}

func NewTree(data []objectJson) (MerkleTree, error) {
	var leaves [][]byte
	var clearData [][]string
	for _, d := range data {
		leaves = append(leaves, d.Hash())
		clearData = append(clearData, []string{d.A, d.V})
		fmt.Printf("leaf created for LAND %s: %x\n", d.A, d.Hash())
	}

	tree, err := solmerkle.GenerateTreeFromHashedItems(leaves)
	if err != nil {
		return MerkleTree{}, fmt.Errorf("could not generate trie: %v", err)
	}

	return MerkleTree{
		mk:    tree,
		leave: leaves,
		clearData: clearData,
	}, nil
}

func (tree MerkleTree) Root() string {
	return hexutil.Encode(tree.mk.Root())
}

func (tree MerkleTree) ProofByIndex(index int) ([]string, error) {
	proof, err := tree.Proof(tree.leave[index])
	if err != nil {
		return nil, err
	}

	return proof, nil
}

func (tree MerkleTree) Proof(data []byte) ([]string, error) {
	var proofStringArr []string
	proof, err := tree.mk.MerkleProof(data)
	if err != nil {
		return nil, fmt.Errorf("could not generate proof: %v", err)
	}

	for _, p := range proof {
		proofStringArr = append(proofStringArr, hexutil.Encode(p))
	}

	return proofStringArr, nil
}

func (tree MerkleTree) AllProof() ([]Proof, error) {
	var proof []Proof
	cont := 0

	for _, leaf := range tree.leave {

		leafString := hexutil.Encode(leaf)
		currentClearData := tree.clearData[cont][:]
		pr, err := tree.Proof(leaf)
		if err != nil {
			return nil, err
		}

		proof = append(proof, Proof{
			Leaf:  leafString,
			Proof: pr,
			ClearData: currentClearData,
		})
		
		cont++
	}
	return proof, nil
}

func (tree MerkleTree) Verify(index int) (bool, error) {
	root := tree.mk.Root()
	proof, err := tree.mk.MerkleProof(tree.leave[index])
	if err != nil {
		return false, fmt.Errorf("could not generate proof: %v", err)
	}
	leaf := tree.leave[index]
	return solmerkle.VerifyMerkleBranch(root, leaf, proof), nil
}


func createMerkle() {
	
	jsonData, err := ReadJson(jsonUrl)
	 fmt.Println(jsonData)

	if err != nil {
		fmt.Println(err)
		return
	}


	mkTree, err := NewTree(jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	proof, err := mkTree.AllProof()
	if err != nil {
		fmt.Println(err)
		return
	}

	type TreeData struct {
		Root  string   `json:"root"`
		Proof []Proof `json:"proof"`
	}

	treeData := TreeData{
		Root:  mkTree.Root(),
		Proof: proof,
	}

	err = WriteDataToFileAsJSON(treeData, jsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "what are you looking for here, dude?")
}

func merkle(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("x-api-key")

	if apiKey != "" && apiKey == apiKeyServer {
		createMerkle()
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}


type ProofFront struct {
	Leaf      string   `json:"leaf"`
	Proof     []string `json:"proof"`
	ClearData []string `json:"clearData"`
}

type RootObject struct {
	Root  string      `json:"root"`
	Proof []ProofFront `json:"proof"`
}

func searchAddress(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	apiKey := r.Header.Get("x-api-key")
	if apiKey != "" && apiKey == apiKeyFront {

		data, err := ioutil.ReadFile(jsonPath)
		if err != nil {
			fmt.Println("Errore durante la lettura del file JSON:", err)
			return
		}
		
		var rootObj RootObject
		err = json.Unmarshal(data, &rootObj)
		if err != nil {
			fmt.Println("Errore durante il parsing del JSON:", err)
			return
		}

		for _, proof := range rootObj.Proof {
			if proof.ClearData[0] == address {
				newJSON := struct {
					Proof     []string `json:"proof"`
					ClearData []string `json:"clearData"`
				}{
					Proof:     proof.Proof,
					ClearData: proof.ClearData,
				}

				
				newJSONString, err := json.MarshalIndent(newJSON, "", "  ")
				if err != nil {
					fmt.Println("Errore durante la conversione del JSON in stringa:", err)
					return
				}

		
				fmt.Fprint(w, string(newJSONString))
				return
			}
		}

		fmt.Fprint(w, "non presente")
	} else {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	http.HandleFunc("/merkle", merkle)
	http.HandleFunc("/search", searchAddress)
	http.ListenAndServe(":8080", nil)
}


