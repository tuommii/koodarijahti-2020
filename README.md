# Koodarijahti 2020

Mielestäni tehtävänanto oli erittäin hauska ja hyvin suunniteltu, jatkuvasti halusin lisätä jonkin ominaisuuden.

Projektini livenä: [Heroku app](https://multiplayer-button.herokuapp.com/) (27.11.2020)

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

API
```markdown
- /state
	- Palauttaa ainoastaan pelin nykyisen tilan. Kutsutaan esimerkiksi sivua uudelleenladatessa
- /click
	- Muuttaa pelin tilaa. Kutsutaan aina painalluksen jälkeen
- /reset
	- Nollaa pelaajan arvot.
```
