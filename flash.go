package main

import (
	"net/http"
)

type FlashCard struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

type Holder struct {
	Cards []FlashCard // to hold the cards
}

var holder = []FlashCard{}

//set the cards in holder
func setInitial() {
	card1 := FlashCard{"Viruses", "These infect applications attaching themselves to the initialization sequence. The virus replicates itself, infecting other code in the computer system.Viruses can also attach themselves to executable code or associate themselves with a file by creating a virus file with the same name  but with an .exe extension, thus creating a decoy which carries the virus"}
	card2 := FlashCard{"Trojans", "a program hiding inside a useful program with malicious purposes. Unlike viruses, a trojan doesn’t replicate itself and it is commonly used to establish a backdoor to be exploited by attackers."}
	card3 := FlashCard{"Worms—unlike viruses", "They don’t attack the host, being self-contained programs that propagate across networks and computers. Worms are often installed through email attachments, sending a copy of themselves to every contact in the infected computer email list. They are commonly used to overload an email server and achieve a denial-of-service attack."}
	card4 := FlashCard{"Ransomware", "A type of malware that denies access to the victim data, threatening to publish or delete it unless a ransom is paid. Advanced ransomware uses cryptoviral extortion, encrypting the victim’s data so that it is impossible to decrypt without the decryption key."}
	card5 := FlashCard{"Spyware", "A type of program installed to collect information about users, their systems or browsing habits, sending the data to a remote user. The attacker can then use the information for blackmailing purposes or download and install other malicious programs from the web."}
	card6 := FlashCard{"Active attacks", "Involve some modification of the data stream or the creation of a false stream"}
	card7 := FlashCard{"A masquerade", "It takes place when one entity pretends to be a different entity. A masquerade attack usually includes one of the other forms of active attack"}
	card8 := FlashCard{"The denial of service attack", "It Prevents or inhibits the normal use or management of communication facilities."}
	holder = append(holder, card1, card2, card3, card4, card5, card6, card7, card8)

}

func main() {
	setInitial()
	http.ListenAndServe(":3636", nil)
}
