package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Card struct {
	Front string `json:"front"` //front tag for json
	Back  string `json:"back"`  //back tag for json
}

type Deck struct {
	Cards []Card //deck to hold the cards
}

var deck = []Card{} //global declaration of deck (of type Deck) allows its value to be accessed by functions below

//set the cards in deck
func setInitialDeck() {
	card1 := Card{"Viruses", "These infect applications attaching themselves to the initialization sequence. The virus replicates itself, infecting other code in the computer system.Viruses can also attach themselves to executable code or associate themselves with a file by creating a virus file with the same name  but with an .exe extension, thus creating a decoy which carries the virus"}
	card2 := Card{"Trojans", "a program hiding inside a useful program with malicious purposes. Unlike viruses, a trojan doesn’t replicate itself and it is commonly used to establish a backdoor to be exploited by attackers."}
	card3 := Card{"Worms—unlike viruses", "They don’t attack the host, being self-contained programs that propagate across networks and computers. Worms are often installed through email attachments, sending a copy of themselves to every contact in the infected computer email list. They are commonly used to overload an email server and achieve a denial-of-service attack."}
	card4 := Card{"Ransomware", "A type of malware that denies access to the victim data, threatening to publish or delete it unless a ransom is paid. Advanced ransomware uses cryptoviral extortion, encrypting the victim’s data so that it is impossible to decrypt without the decryption key."}
	card5 := Card{"Spyware", "A type of program installed to collect information about users, their systems or browsing habits, sending the data to a remote user. The attacker can then use the information for blackmailing purposes or download and install other malicious programs from the web."}
	card6 := Card{"Active attacks", "Involve some modification of the data stream or the creation of a false stream"}
	card7 := Card{"A masquerade", "It takes place when one entity pretends to be a different entity. A masquerade attack usually includes one of the other forms of active attack"}
	card8 := Card{"The denial of service attack", "It Prevents or inhibits the normal use or management of communication facilities."}
	deck = append(deck, card1, card2, card3, card4, card5, card6, card7, card8)

}

//main file handler
func serveMainFile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	b, err := json.MarshalIndent(deck, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, template.JS(b)) //escapes stored JSON data within a script element for proper front-end consumption
}

func cardform(w http.ResponseWriter, r *http.Request) { //handler to serve form for adding custom cards when /form handler is invoked
	t, _ := template.ParseFiles("add.html")
	t.Execute(w, nil)
}

func cardadd(w http.ResponseWriter, r *http.Request) {
	var newcard Card                               //defining newcard (of type Card) to hold form data
	newcard.Front = r.FormValue("term")            //HTML form input hooks for storage
	newcard.Back = r.FormValue("definition")       //
	deck = append(deck, newcard)                   //append form inputs to our globally stored data for template injection
	b, err := json.MarshalIndent(deck, "", "    ") //format deck for legibility
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	os.Stdout.Write(b)                      //each new card is written to the console on submission to confirm proper form function and to test for Marshalling/JSON errors.
	t, _ := template.ParseFiles("add.html") //re-renders blank form template
	t.Execute(w, nil)
}

func main() { //Main function for the program
	setInitialDeck()
	http.HandleFunc("/", serveMainFile) //renders main.html template with embedded JS for flipping/cycling through cards injected as formatted JSON
	http.HandleFunc("/form", cardform)  //renders add.html
	http.HandleFunc("/add", cardadd)    //writes POST containing newcard data, adds it to globally stored deck, and re-renders add.html
	http.ListenAndServe(":8080", nil)   //port info
}
