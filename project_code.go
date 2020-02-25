package main

import (
	"encoding/json"
	"fmt"
	"strconv" //"github.com/hyperledger/fabric/core/chaincode/shim"
	//pb "github.com/hyperledger/fabric/protos/peer"
	"time"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type banka struct {
	ID                   string
	Naziv                string
	GodinaPorekla        int
	DrzavaPorekla        string
	DrzaveUKojimaPosluje []string
	Korisnici            []korisnik
}

type korisnik struct {
	ID            string
	BrojRacuna    string
	Ime           string
	Prezime       string
	Email         string
	KolicinaNovca float64
	Krediti       []kredit
	Transakcije   []transakcija
}

type transakcija struct {
	ID            string
	Datum         string
	IDUplatioc    string
	IDPrimalac    string
	KolicinaNovca float64
}

type kredit struct {
	ID                 string
	DatumOdobrenja     string
	DatumZavrsetka     string
	VelicinaRate       float64
	Kamata             float64
	BrojRata           int
	BrojOtplacenihRate int
	DobijeniNovac      float64
}

//Global variables for IDs
var korisnikId int
var kreditId int
var bankaId int
var transakcijaId int

type SimpleChaincode struct {
}

//Init function
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	var kredit1 = kredit{"kr1", "1.1.2014.", "1.1.2016.", 10.0, 2.5, 24, 24, 2400.0}
	var kredit2 = kredit{"kr2", "2.1.2014.", "2.1.2016.", 10.0, 2.5, 24, 20, 2400.0}
	var kredit3 = kredit{"kr3", "3.1.2014.", "3.1.2016.", 10.0, 2.5, 24, 24, 2400.0}
	var kredit4 = kredit{"kr4", "4.1.2014.", "4.1.2016.", 10.0, 2.5, 24, 20, 2400.0}
	kreditId = 5
	var zaK1 = make([]kredit, 0, 20)
	zaK1 = append(zaK1, kredit1)
	zaK1 = append(zaK1, kredit2)

	var zaK2 = make([]kredit, 0, 20)
	zaK2 = append(zaK2, kredit3)

	var zaK3 = make([]kredit, 0, 20)
	zaK3 = append(zaK3, kredit4)

	var transakcija1 = transakcija{"t1", "1.1.2020", "ko1", "ko2", 100.0}
	var transakcija2 = transakcija{"t2", "2.1.2020", "ko2", "ko1", 200.0}
	var transakcija3 = transakcija{"t3", "3.1.2020", "ko1", "ko3", 300.0}
	var transakcija4 = transakcija{"t4", "4.1.2020", "ko3", "ko2", 400.0}
	transakcijaId = 5

	var tzaK1 = make([]transakcija, 0, 100)
	tzaK1 = append(tzaK1, transakcija1)
	tzaK1 = append(tzaK1, transakcija2)
	tzaK1 = append(tzaK1, transakcija3)

	var tzaK2 = make([]transakcija, 0, 100)
	tzaK2 = append(tzaK2, transakcija1)
	tzaK2 = append(tzaK2, transakcija2)
	tzaK2 = append(tzaK2, transakcija4)

	var tzaK3 = make([]transakcija, 0, 100)
	tzaK3 = append(tzaK3, transakcija3)
	tzaK3 = append(tzaK3, transakcija4)

	var korisnik1 = korisnik{"ko1", "1", "Pera", "Peric", "pera@gmail.com", 1000.0, zaK1, tzaK1}
	var korisnik2 = korisnik{"ko2", "2", "Marko", "Markovic", "mare@gmail.com", 2000.0, zaK2, tzaK2}
	var korisnik3 = korisnik{"ko3", "3", "Zika", "Zikic", "zika@gmail.com", 3000.0, zaK3, tzaK3}
	korisnikId = 4

	var drzave1 = make([]string, 0, 20)
	drzave1 = append(drzave1, "Srbija")
	drzave1 = append(drzave1, "Bosna i Hercegovina")
	drzave1 = append(drzave1, "Hrvatska")
	var korisnici1 = make([]korisnik, 0, 20)
	korisnici1 = append(korisnici1, korisnik1)
	var banka1 = banka{"b1", "Prva banka", 1990, "Srbija", drzave1, korisnici1}

	var drzave2 = make([]string, 0, 20)
	drzave2 = append(drzave1, "Srbija")
	drzave2 = append(drzave1, "Makedonija")
	drzave2 = append(drzave1, "Madjarska")
	var korisnici2 = make([]korisnik, 0, 20)
	korisnici2 = append(korisnici2, korisnik2)
	korisnici2 = append(korisnici2, korisnik3)
	var banka2 = banka{"b2", "Druga banka", 1987, "Srbija", drzave2, korisnici2}
	bankaId = 3

	// Write the state to the ledger

	//Krediti
	ajson, _ := json.Marshal(kredit1)
	err := stub.PutState(kredit1.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(kredit2)
	err = stub.PutState(kredit2.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(kredit3)
	err = stub.PutState(kredit3.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(kredit4)
	err = stub.PutState(kredit4.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Korisnici
	ajson, _ = json.Marshal(korisnik1)
	err = stub.PutState(korisnik1.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(korisnik2)
	err = stub.PutState(korisnik2.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(korisnik3)
	err = stub.PutState(korisnik3.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Banke
	ajson, _ = json.Marshal(banka1)
	err = stub.PutState(banka1.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(banka2)
	err = stub.PutState(banka2.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Transakcije
	ajson, _ = json.Marshal(transakcija1)
	err = stub.PutState(transakcija1.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transakcija2)
	err = stub.PutState(transakcija2.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transakcija3)
	err = stub.PutState(transakcija3.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(transakcija4)
	err = stub.PutState(transakcija4.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	/*if function == "delete" {
		return t.delete(stub, args)
	}*/
	/*if function == "query" {
		return t.query(stub, args)
	}*/
	if function == "dodajKorisnika" {
		return t.dodajKorisnika(stub, args)
	}
	/*if function == "dodajTransakciju" {
		return t.dodajTransakciju(stub, args)
	}*/
	if function == "dodajKorisnikaBanci" {
		return t.dodajKorisnikaBanci(stub, args)
	}
	/*if function == "podizanjeKredita" {
		return t.podizanjeKredita(stub, args)
	}*/

	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) dodajKorisnika(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var brRacuna, ime, prezime, email, kolicinaNovca string // Entities

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 4, function followed by 2 names and 1 value")
	}

	brRacuna = args[0]
	ime = args[1]
	prezime = args[2]
	email = args[3]
	//TODO ovde moras odraditi proveru je l dobar parse
	var kolicinaNovcaFloat, error = strconv.ParseFloat(kolicinaNovca, 64)

	if error != nil {
		return shim.Error("Netacna vrednost za kolicinu novca, mora biti realan broj!")
	}

	korId := "ko" + strconv.Itoa(korisnikId)
	korisnikId = korisnikId + 1

	var newKrediti = make([]kredit, 0, 20)
	var newTransakcije = make([]transakcija, 0, 100)
	//TODO mora i provera da li vec postoji korisnik sa ovim ID
	var newKorisnik = korisnik{korId, brRacuna, ime, prezime, email, kolicinaNovcaFloat, newKrediti, newTransakcije}

	ajson, _ := json.Marshal(newKorisnik)
	err := stub.PutState(newKorisnik.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) dodajKorisnikaBanci(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var korisnikKey, bankaKey string // Entities

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	korisnikKey = args[0]
	bankaKey = args[1]

	//load korisnik
	loadedKorisnik, err := stub.GetState(korisnikKey)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + korisnikKey + "\"}"
		return shim.Error(jsonResp)
	}
	if loadedKorisnik == nil || len(loadedKorisnik) == 0 {
		jsonResp := "{\"Error\":\" " + korisnikKey + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}

	korisnikFromJson := korisnik{}
	err = json.Unmarshal(loadedKorisnik, &korisnikFromJson)
	if err != nil {
		return shim.Error("Failed to unmarshall Korisnik")
	}

	//load banka
	loadedBanka, err := stub.GetState(bankaKey)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + bankaKey + "\"}"
		return shim.Error(jsonResp)
	}
	if loadedBanka == nil || len(loadedBanka) == 0 {
		jsonResp := "{\"Error\":\" " + bankaKey + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}

	bankaFromJson := banka{}
	err = json.Unmarshal(loadedBanka, &bankaFromJson)
	if err != nil {
		return shim.Error("Failed to unmarshall Banka")
	}

	for i := 0; i < len(bankaFromJson.Korisnici); i++ {
		if bankaFromJson.Korisnici[i].ID == korisnikKey {
			jsonResp := "{\"Error\":\" Korisnik sa id-em: " + korisnikKey + " vec postoji u ovoj banci!" + "\"}"
			return shim.Error(jsonResp)
		}
	}

	bankaFromJson.Korisnici = append(bankaFromJson.Korisnici, korisnikFromJson)

	ajson, _ := json.Marshal(bankaFromJson)
	err = stub.PutState(bankaFromJson.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}
	//mrs
	return shim.Success(nil)
}

func (t *SimpleChaincode) podizanjeKredita(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var idKorisnika, iznosKredita, brojRata, kamata string // Entities

	if len(args) != 5 {
		return shim.Error("Netacan broj argumenata! Ocekivani argumenti: ID korisnika, iznos kredita, broj rata, kamata")
	}

	idKorisnika = args[0]
	iznosKredita = args[1]
	brojRata = args[2]
	kamata = args[3]

	//load korisnik
	loadedKorisnik, err := stub.GetState(idKorisnika)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + idKorisnika + "\"}"
		return shim.Error(jsonResp)
	}
	if loadedKorisnik == nil || len(loadedKorisnik) == 0 {
		jsonResp := "{\"Error\":\" " + idKorisnika + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}

	korisnikFromJson := korisnik{}
	err = json.Unmarshal(loadedKorisnik, &korisnikFromJson)
	if err != nil {
		return shim.Error("Failed to unmarshall Korisnik")
	}

	for i := 0; i < len(korisnikFromJson.Krediti); i++ {
		if korisnikFromJson.Krediti[i].BrojOtplacenihRate != korisnikFromJson.Krediti[i].BrojRata {
			return shim.Error("Zeljeni korisnik ima neotplacen kredit!")
		}
	}

	sumaUplata := 0.0
	brojUplata := 0
	for i := 0; i < len(korisnikFromJson.Transakcije); i++ {
		if korisnikFromJson.Transakcije[i].IDPrimalac == korisnikFromJson.ID {
			sumaUplata += korisnikFromJson.Transakcije[i].KolicinaNovca
			brojUplata++
		}
	}
	maxIznosKredita := (sumaUplata / float64(brojUplata)) * 5.0

	iznosKreditaFloat, floatError := strconv.ParseFloat(iznosKredita, 64)
	if floatError != nil {
		return shim.Error("Greska pri parsiranju broja za zeljeni iznos kredita!")
	}

	if iznosKreditaFloat > maxIznosKredita {
		return shim.Error("Zeljeni iznos kredita je veci od dozvoljenog!")
	}

	kamataFloat, kamataErr := strconv.ParseFloat(kamata, 64)
	if kamataErr != nil {
		return shim.Error("Zeljeni iznos kamate mora biti realan broj u procentima!")
	}

	brojRataInt, rateErr := strconv.Atoi(brojRata)
	if rateErr != nil {
		return shim.Error("Broj rata mora biti ceo broj!")
	}

	mesecnaRata := (iznosKreditaFloat * kamataFloat * 0.01) / float64(brojRataInt)
	datumOdobravanja := time.Now().Format("01-02-2006")
	var datumZavrsetka = ""

	idKredita := "kr" + strconv.Itoa(kreditId)
	kreditId = kreditId + 1

	var newKredit = kredit{idKredita, datumOdobravanja, datumZavrsetka, mesecnaRata, kamataFloat, brojRataInt, 0, iznosKreditaFloat}
	korisnikFromJson.Krediti = append(korisnikFromJson.Krediti, newKredit)
	//TODO da li da stavim i kredit na state illi samo da update u korisnika?
	ajson, _ := json.Marshal(newKredit)
	err = stub.PutState(newKredit.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	ajson, _ = json.Marshal(korisnikFromJson)
	err = stub.PutState(korisnikFromJson.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) uplataRateKredita(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var idKorisnika string

	idKorisnika = args[0]

	//load korisnik
	loadedKorisnik, err := stub.GetState(idKorisnika)

	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + idKorisnika + "\"}"
		return shim.Error(jsonResp)
	}
	if loadedKorisnik == nil || len(loadedKorisnik) == 0 {
		jsonResp := "{\"Error\":\" " + idKorisnika + " does not exit " + "\"}"
		return shim.Error(jsonResp)
	}

	korisnikFromJson := korisnik{}
	err = json.Unmarshal(loadedKorisnik, &korisnikFromJson)
	if err != nil {
		return shim.Error("Failed to unmarshall Korisnik")
	}

	var rbrKredita = -1
	for i := 0; i < len(korisnikFromJson.Krediti); i++ {
		if korisnikFromJson.Krediti[i].BrojOtplacenihRate != korisnikFromJson.Krediti[i].BrojRata {
			rbrKredita = i
		}
	}

	if rbrKredita == -1 {
		return shim.Error("Nemate zaduzen kredit!")
	}

	if korisnikFromJson.Krediti[rbrKredita].VelicinaRate > korisnikFromJson.KolicinaNovca {
		return shim.Error("Nemate dovoljno novca na racunu da platite ratu za kredit!")
	}

	korisnikFromJson.Krediti[rbrKredita].BrojOtplacenihRate++
	korisnikFromJson.KolicinaNovca -= korisnikFromJson.Krediti[rbrKredita].VelicinaRate

	ajson, _ := json.Marshal(korisnikFromJson)
	err = stub.PutState(korisnikFromJson.ID, ajson)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}
func main() {

}
