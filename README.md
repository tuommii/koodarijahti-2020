# Koodarijahti 2020

Mielestäni [tehtävänanto](https://github.com/tuommii/vincit/blob/master/Tehtava.pdf) oli **erittäin hauska** ja **hyvin suunniteltu**, jatkuvasti teki mieli lisätä jokin uusi ominaisuus.

<img src="https://github.com/tuommii/vincit/blob/master/screenshot.png" width="300">

Projekti livenä: [Heroku app](https://multiplayer-button.herokuapp.com/) (27.11.2020)

## Ratkaisuni

### Palvelinsovellus
Tein palvelinsovelluksen **go**-kielellä, **ilman** riippuvuuksia. Kaikki pelin tilaan vaikuttavat muutokset
tapahtuvat palvelimella.

### Selainsovellus
Selain-puolella tehtävänannossa toivottiin käytettävän jotain nykyaikaista työkalua, niinpä hylkäsin ajatukseni
tehdä kaikki itse ja käytin **Vue.js**-kirjastoa, jota olin kokeillut jo joskus ennen 1.0-versiota
(omasi jo silloin hyvän dokumentaation).

## Muutama huomio
Siltä varalta, ettei koodini olekaan niin helposti luettavaa kuin toivon 😉

API:
```markdown
- /state
	- Ainoastaan palauttaa pelin nykyisen tilan
- /click
	- Muuttaa pelin tilaa sekä palauttaa sen
- /reset
	- Palauttaa pelaajan arvot lähtötilanteeseen.
```

* **Ääniefektit** klikkaukselle, voitolle ja pelin loppumiselle.
* Laskuri nollaantuu aina 500:n painalluksen jälkeen, jottei ikinä mennä integerin maksimiarvon yli.
Tämä ei vaikuta palkintojenjakoon.
* Pelaajat yksilöidään IP-osoitteella, joka toimii myös map-tyypin avaimena. Näin pelaajan tila säilyy, vaikka selain
välissä suljettaisiinkin. En siis käyttänyt tietokantaa, vaikka MySQL ja MongoDB ovatkin tuttuja.
* Herokun Hobby Dyno uudelleenkäynnistyy tietyin väliajoin, joten samalla se tulee hoitaneeksi muistin vapauttamisen ;)
* Ainoastaan joka kymmenes painallus voi sisältää palkinnon, turha tarkistaa joka kerta.
* Palvelin loggaa pyynnöt middleware-funktiolla
* Painikkeelle on asetettu pieni viive klikkausten välille.
* Jos ympäristömuuttujaa __PORT__ ei löydy, kuunnellaan porttia 3000.
* Palvelinsovelluksen koodi on runsaasti kommentoitua ainoastaan gofmt:n vuoksi.
* Ääniasetusta ei tallenneta, joten se palautuu sivulatauksen jälkeen.

## Kokeile itse
Go tulee olla asennettuna. Kloonaa repo ja suorita **hakemiston juuressa**:

```go build -o server```

```./server```

Avaa selaimessa [http://localhost:3000/](http://localhost:3000/)
