# Koodarijahti 2020

Mielestäni tehtävänanto oli erittäin hauska ja hyvin suunniteltu, jatkuvasti tekisi mieli lisätä jokin uusi ominaisuus.

Projekti livenä: [Heroku app](https://multiplayer-button.herokuapp.com/) (27.11.2020)

## Ratkaisuni

### Palvelinsovellus
Tein palvelinsovelluksen **go**-kielellä, **ilman** riippuvuuksia. Kaikki pelin tilaan vaikuttavat muutokset
tapahtuvat palvelimella.

### Selainsovellus
Selain-puolella tehtävänannossa toivottiin käytettävän jotain nykyaikaista työkalua, niinpä hylkäsin ajatukseni
tehdä kaikki itse ja käytin **Vue.js**-kirjastoa, jota olin kokeillut jo joskus ennen 1.0-versiota
(omasi jo silloin hyvän dokumentaation).

## Muutama huomio koodista
Siltä varalta, ettei koodini olekaan niin helposti luettavaa kuin toivon:

API:
```markdown
- /state
	- Palauttaa ainoastaan pelin nykyisen tilan
- /click
	- Muuttaa pelin tilaa. Kutsutaan aina painalluksen jälkeen
- /reset
	- Nollaa pelaajan arvot.
```

* Laskuri nollaantuu aina 500:n painalluksen jälkeen, jottei ikinä mennä integerin maksimiarvon yli.
Tämä ei vaikuta palkintojenjakoon.
* Pelaajat yksilöidään IP-osoitteella, joka toimii myös map-tyypin avaimena. Näin pelaajan tila säilyy, vaikka selain
välissä suljettaisiinkin.
* Ainoastaan joka kymmenes painallus voi sisältää palkinnon, turha tarkistaa joka kerta.
* Jos ympäristömuuttujaa __PORT__ ei löydy, kuunnellaan porttia 3000.
* Painikkeelle asetettu pieni viive, ennen kuin sitä voi uudelleen klikata.

## Kokeile itse

### Linux / OS X
Go tulee olla asennettuna. Suorita **hakemiston juuressa**:
```make```
```./bin/server```
