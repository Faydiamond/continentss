package continentss

import "fmt"

var logMessage = "[LOG]"

// Version of the calculator
var Version = "1.0"

func WhatCountry(country string) string {
	fmt.Println("Digite una pais ")
	var continent string
	switch country {
	case "armenia", "afganistán", "arabia saudita", "azerbaiyán ", "banglades", "barein", "birmania ", "brunei", "butan", "camboya",
		"catar", "china", "chipre", "corea del norte", "corea del sur", "emiratos arabes unidos", "filipinas", "georgia ", "india", "indonesia", "irak",
		"iran", "israel", "japon", "jordania", "kazajistan", "kirguistan", "kuwait", "laos", "libano", "malasia", "maldivas", "mongolia", "nepal",
		"Oman", "Pakistan", "singapur", "siria", "sri lanka", "tayikistan", "tailandia", "timor oriental", "Turkmenistan", "turquía", "uzbekistan",
		"vietnam", "yemen":
		continent = "Asia"
	case "angola", "argelia", "benin", "botsuana", "burkina faso", "burundi", "cabo verde", "camerun", "centroafricana", "comores", "costa de marfil", "chad",
		"egipto", "eritrea", "etiopia", "gabon", "gambia", "ghana", "guinea", "guinea-bissau", "guinea ecuatorial", "kenia", "lesoto", "liberia", "libia",
		"madagascar", "malaui", "mali", "marruecos", "mauricio", "mauritania", "mozambique", "namibia", "nigeria", "ruanda", "republica del congo",
		"república democrática del congo", "santo tome y principe", "senegal", "seychelles", "sierra leona", "somalia", "sudafrica", "Sudan", "sudán del sur",
		"suazilandia", "tanzania", "tunez", "togo", "uganda", "yibuti", "zambia", "zimbabue":
		continent = "Africa"
	case "albania", "andorra", "belgica", "bosnia y herzegovina", "checo", "croacia", "eslovaquia", "españa", "finlandia", "gran bretaña", "holanda", "italia",
		"islandia", "liechtenstein", "luxemburgo", "moldavia", "monaco", "polonia", "rumania", "san marino", "suecia", "ucrania", "suiza", "serbia y montenegro",
		"rusia", "portugal", "noruega", "malta", "macedonia", "lituania", "letonia", "irlanda", "hungria", "grecia", "francia", "estonia", "eslovenia",
		"dinamarca", "bulgaria", "bielorusia", "austria", "alemania":
		continent = "Europa"
	case "antigua y barbuda", "barbados", "costa rica", "dominica", "jamaica", "mexico", "republica dominicana", "trinidad y tobago", "panama", "estados unidos",
		"granda", "cuba", "canada", "bahamas":
		continent = "América del Norte"
	case "argentina", "brasil", "colombia", "guyana", "surinam", "venezuela", "bolivia", "chile", "ecuador", "peru", "uruguay":
		continent = "América del Sur"
	case "australia", "islas de cook", "micronesie", "papua nueva guinea", "tonga", "fiji", "kiribati", "nueva zelanda", "Samoa", "Vanuatu":
		continent = "Oceania"
	}
	return continent
}
