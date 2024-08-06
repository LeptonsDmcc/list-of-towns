package config

import (
	"log"

	"listof/pkg/models"

	"github.com/gorilla/mux"
)

var (
	Countries       []models.Country
	States          []models.State
	LocalGovernments []models.LocalGovernment
	Towns          []models.Town
)

func SetupApp() *mux.Router {
	router := mux.NewRouter()
	initData()
	return router
}

func initData() {
	log.Println("Initializing static data...")

	Countries = []models.Country{
		{ID: 1, Name: "Nigeria"},
	}

	States = []models.State{
		{ID: 1, Name: "Abia", CountryID: 1},
		{ID: 2, Name: "Adamawa", CountryID: 1},
		{ID: 3, Name: "Akwa Ibom", CountryID: 1},			
		{ID: 4, Name: "Anambra", CountryID: 1},
		{ID: 5, Name: "Bauchi", CountryID: 1},
		{ID: 6, Name: "Bayelsa", CountryID: 1},				
		{ID: 7, Name: "Benue", CountryID: 1},
		{ID: 8, Name: "Borno", CountryID: 1},
		{ID: 9, Name: "Cross River", CountryID: 1},			
		{ID: 10, Name: "Delta", CountryID: 1},
		{ID: 11, Name: "Ebonyi", CountryID: 1},
		{ID: 12, Name: "Edo", CountryID: 1},				
		{ID: 13, Name: "Ekiti", CountryID: 1},
		{ID: 14, Name: "Enugu", CountryID: 1},
		{ID: 15, Name: "Gombe", CountryID: 1},				
		{ID: 16, Name: "Imo", CountryID: 1},
		{ID: 17, Name: "Jigawa", CountryID: 1},
		{ID: 18, Name: "Kaduna", CountryID: 1},				
		{ID: 19, Name: "Kano", CountryID: 1},
		{ID: 20, Name: "Katsina", CountryID: 1},
		{ID: 21, Name: "Kebbi", CountryID: 1},				
		{ID: 22, Name: "Kogi", CountryID: 1},
		{ID: 23, Name: "Kwara", CountryID: 1},
		{ID: 24, Name: "Lagos", CountryID: 1},				
		{ID: 25, Name: "Nasarawa", CountryID: 1},
		{ID: 26, Name: "Niger", CountryID: 1},
		{ID: 27, Name: "Ogun", CountryID: 1},				
		{ID: 28, Name: "Ondo", CountryID: 1},
		{ID: 29, Name: "Osun", CountryID: 1},
		{ID: 30, Name: "Oyo", CountryID: 1},				
		{ID: 31, Name: "Plateau", CountryID: 1},
		{ID: 32, Name: "Rivers", CountryID: 1},
		{ID: 33, Name: "Sokoto", CountryID: 1},				
		{ID: 34, Name: "Taraba", CountryID: 1},
		{ID: 35, Name: "Yobe", CountryID: 1},
		{ID: 36, Name: "Zamfara", CountryID: 1},			
		{ID: 37, Name: "Abuja", CountryID: 1},
		
	}

	LocalGovernments = []models.LocalGovernment{

		//Abia State
		{ID: 1, Name: "Aba North", StateID: 1},
		{ID: 2, Name: "Aba South", StateID: 1},
		{ID: 3, Name: "Arochukwu", StateID: 1},
		{ID: 4, Name: "Bende", StateID: 1},
		{ID: 5, Name: "Ikwuano", StateID: 1},
		{ID: 6, Name: "Isiala Ngwa North", StateID: 1},
		{ID: 7, Name: "Isiala Ngwa South", StateID: 1},
		{ID: 8, Name: "Isuikwuato", StateID: 1},
		{ID: 9, Name: "Obingwa", StateID: 1},
		{ID: 10, Name: "Ohafia", StateID: 1},
		{ID: 11, Name: "Osisioma", StateID: 1},
		{ID: 12, Name: "Ugwunagbo", StateID: 1},
		{ID: 13, Name: "Umuahia North", StateID: 1},
		{ID: 14, Name: "Umuahia South", StateID: 1},
		{ID: 15, Name: "Umu Nneochi", StateID: 1},

		//Adamawa State
		{ID: 16, Name: "Demsa", StateID: 2},
		{ID: 17, Name: "Fufore", StateID: 2},
		{ID: 18, Name: "Ganye", StateID: 2},
		{ID: 19, Name: "Girei", StateID: 2},
		{ID: 20, Name: "Gombi", StateID: 2},
		{ID: 21, Name: "Guyuk", StateID: 2},
		{ID: 22, Name: "Hong", StateID: 2},
		{ID: 23, Name: "Jada", StateID: 2},
		{ID: 24, Name: "Lamurde", StateID: 2},
		{ID: 25, Name: "Madagali", StateID: 2},
		{ID: 26, Name: "Maiha", StateID: 2},
		{ID: 27, Name: "Mubi North", StateID: 2},
		{ID: 28, Name: "Mubi South", StateID: 2},
		{ID: 29, Name: "Numan", StateID: 2},
		{ID: 30, Name: "Shelleng", StateID: 2},
		{ID: 31, Name: "Song", StateID: 2},
		{ID: 32, Name: "Toungo", StateID: 2},
		{ID: 33, Name: "Yola North", StateID: 2},
		{ID: 34, Name: "Yola South", StateID: 2},

		//Akwa Ibom State
		{ID: 35, Name: "Abak", StateID: 3},
		{ID: 36, Name: "Eastern Obolo", StateID: 3},
		{ID: 37, Name: "Eket", StateID: 3},
		{ID: 38, Name: "Esit Eket", StateID: 3},
		{ID: 39, Name: "Essien Udim", StateID: 3},
		{ID: 40, Name: "Etim Ekpo", StateID: 3},
		{ID: 41, Name: "Etinan", StateID: 3},
		{ID: 42, Name: "Ibeno", StateID: 3},
		{ID: 43, Name: "Ibesikpo Asutan", StateID: 3},
		{ID: 44, Name: "Ibiono-Ibom", StateID: 3},
		{ID: 45, Name: "Ika", StateID: 3},
		{ID: 46, Name: "Ikono", StateID: 3},
		{ID: 47, Name: "Ikot Abasi", StateID: 3},
		{ID: 48, Name: "Ikot Ekpene", StateID: 3},
		{ID: 49, Name: "Ini", StateID: 3},
		{ID: 50, Name: "Itu", StateID: 3},
		{ID: 51, Name: "Mbo", StateID: 3},
		{ID: 52, Name: "Mkpat-Enin", StateID: 3},
		{ID: 53, Name: "Nsit-Atai", StateID: 3},
		{ID: 54, Name: "Nsit-Ibom", StateID: 3},
		{ID: 55, Name: "Nsit-Ubium", StateID: 3},
		{ID: 56, Name: "Obot Akara", StateID: 3},
		{ID: 57, Name: "Okobo", StateID: 3},
		{ID: 58, Name: "Onna", StateID: 3},
		{ID: 59, Name: "Demsa", StateID: 3},
		{ID: 60, Name: "Demsa", StateID: 3},
		{ID: 61, Name: "Oron", StateID: 3},
		{ID: 62, Name: "Oruk Anam", StateID: 3},
		{ID: 63, Name: "Udung-Uko", StateID: 3},
		{ID: 64, Name: "Ukanafun", StateID: 3},
		{ID: 65, Name: "Uruan", StateID: 3},
		{ID: 66, Name: "Urue-Offong/Oruko", StateID: 3},
		{ID: 67, Name: "Uyo", StateID: 3},

		//Anambra
		{ID: 68, Name: "Aguata", StateID: 4},
		{ID: 69, Name: "Anambra East", StateID: 4},
		{ID: 70, Name: "Anambra West", StateID: 4},
		{ID: 71, Name: "Anaocha", StateID: 4},
		{ID: 72, Name: "Awka North", StateID: 4},
		{ID: 73, Name: "Awka South", StateID: 4},
		{ID: 74, Name: "Awka South", StateID: 4},
		{ID: 75, Name: "Dunukofia", StateID: 4},
		{ID: 76, Name: "Ekwusigo", StateID: 4},
		{ID: 77, Name: "Idemili North", StateID: 4},
		{ID: 78, Name: "Idemili South", StateID: 4},
		{ID: 79, Name: "Ihiala", StateID: 4},
		{ID: 80, Name: "Njikoka", StateID: 4},
		{ID: 81, Name: "Nnewi North", StateID: 4},
		{ID: 82, Name: "Nnewi South", StateID: 4},
		{ID: 83, Name: "Ogbaru", StateID: 4},
		{ID: 84, Name: "Onitsha North", StateID: 4},
		{ID: 85, Name: "Onitsha South", StateID: 4},
		{ID: 86, Name: "Orumba North", StateID: 4},
		{ID: 87, Name: "Orumba South", StateID: 4},
		{ID: 88, Name: "Oyi", StateID: 4},

		//
		{ID: 89, Name: "Alkaleri", StateID: 5},
		{ID: 90, Name: "Bauchi", StateID: 5},
		{ID: 91, Name: "Bogoro", StateID: 5},
		{ID: 92, Name: "Damban", StateID: 5},
		{ID: 93, Name: "Darazo", StateID: 5},
		{ID: 94, Name: "Dass", StateID: 5},
		{ID: 95, Name: "Gamawa", StateID: 5},
		{ID: 96, Name: "Ganjuwa", StateID: 5},
		{ID: 97, Name: "Giade", StateID: 5},
		{ID: 98, Name: "Itas/Gadau", StateID: 5},
		{ID: 99, Name: "Jama'are", StateID: 5},
		{ID: 100, Name: "Katagum", StateID: 5},
		{ID: 101, Name: "Kirfi", StateID: 5},
		{ID: 102, Name: "Misau", StateID: 5},
		{ID: 103, Name: "Ningi", StateID: 5},
		{ID: 104, Name: "Shira", StateID: 5},
		{ID: 105, Name: "Tafawa Balewa", StateID: 5},
		{ID: 106, Name: "Toro", StateID: 5},
		{ID: 107, Name: "Warji", StateID: 5},
		{ID: 108, Name: "Zaki", StateID: 5},

		//Bayelsa
		{ID: 109, Name: "Brass", StateID: 6},
		{ID: 110, Name: "Ekeremor", StateID: 6},
		{ID: 111, Name: "Kolokuma/Opokuma", StateID: 6},
		{ID: 112, Name: "Nembe", StateID: 6},
		{ID: 113, Name: "Ogbia", StateID: 6},
		{ID: 114, Name: "Sagbama", StateID: 6},
		{ID: 115, Name: "Southern Ijaw", StateID: 6},
		{ID: 116, Name: "Yenagoa", StateID: 6},

		//
		{ID: 117, Name: "Ado", StateID: 7},
		{ID: 118, Name: "Agatu", StateID: 7},
		{ID: 119, Name: "Apa", StateID: 7},
		{ID: 120, Name: "Buruku", StateID: 7},
		{ID: 121, Name: "Gboko", StateID: 7},
		{ID: 122, Name: "Guma", StateID: 7},
		{ID: 123, Name: "Gwer East", StateID: 7},
		{ID: 124, Name: "Gwer West", StateID: 7},
		{ID: 125, Name: "Katsina-Ala", StateID: 7},
		{ID: 126, Name: "Konshisha", StateID: 7},
		{ID: 127, Name: "Kwande", StateID: 7},
		{ID: 128, Name: "Logo", StateID: 7},
		{ID: 129, Name: "Makurdi", StateID: 7},
		{ID: 130, Name: "Obi", StateID: 7},
		{ID: 131, Name: "Ogbadibo", StateID: 7},
		{ID: 132, Name: "Ohimini", StateID: 7},
		{ID: 133, Name: "Oju", StateID: 7},
		{ID: 134, Name: "Okpokwu", StateID: 7},
		{ID: 135, Name: "Otukpo", StateID: 7},
		{ID: 136, Name: "Tarka", StateID: 7},
		{ID: 137, Name: "Ukum", StateID: 7},
		{ID: 138, Name: "Ushongo", StateID: 7},
		{ID: 139, Name: "Vandeikya", StateID: 7},

		//
		{ID: 140, Name: "Abadam", StateID: 8},
		{ID: 141, Name: "Askira/Uba", StateID: 8},
		{ID: 142, Name: "Bama", StateID: 8},
		{ID: 143, Name: "Bayo", StateID: 8},
		{ID: 144, Name: "Biu", StateID: 8},
		{ID: 145, Name: "Chibok", StateID: 8},
		{ID: 146, Name: "Damboa", StateID: 8},
		{ID: 147, Name: "Dikwa", StateID: 8},
		{ID: 148, Name: "Gubio", StateID: 8},
		{ID: 149, Name: "Guzamala", StateID: 8},
		{ID: 150, Name: "Gwoza", StateID: 8},
		{ID: 151, Name: "Hawul", StateID: 8},
		{ID: 152, Name: "Jere", StateID: 8},
		{ID: 153, Name: "Kaga", StateID: 8},
		{ID: 154, Name: "Kala/Balge", StateID: 8},
		{ID: 155, Name: "Konduga", StateID: 8},
		{ID: 156, Name: "Kukawa", StateID: 8},
		{ID: 157, Name: "Kwaya Kusar", StateID: 8},
		{ID: 158, Name: "Mafa", StateID: 8},
		{ID: 159, Name: "Magumeri", StateID: 8},
		{ID: 160, Name: "Maiduguri", StateID: 8},
		{ID: 161, Name: "Marte", StateID: 8},
		{ID: 162, Name: "Mobbar", StateID: 8},
		{ID: 163, Name: "Monguno", StateID: 8},
		{ID: 164, Name: "Ngala", StateID: 8},
		{ID: 165, Name: "Nganzai", StateID: 8},
		{ID: 166, Name: "Shani", StateID: 8},

		//
		{ID: 167, Name: "Abi", StateID: 9},
		{ID: 168, Name: "Akamkpa", StateID: 9},
		{ID: 169, Name: "Akpabuyo", StateID: 9},
		{ID: 170, Name: "Bakassi", StateID: 9},
		{ID: 171, Name: "Bekwarra", StateID: 9},
		{ID: 172, Name: "Biase", StateID: 9},
		{ID: 173, Name: "Boki", StateID: 9},
		{ID: 174, Name: "Calabar Municipal", StateID: 9},
		{ID: 175, Name: "Calabar South", StateID: 9},
		{ID: 176, Name: "Etung", StateID: 9},
		{ID: 177, Name: "Ikom", StateID: 9},
		{ID: 178, Name: "Obanliku", StateID: 9},
		{ID: 179, Name: "Obubra", StateID: 9},
		{ID: 180, Name: "Obudu", StateID: 9},
		{ID: 181, Name: "Odukpani", StateID: 9},
		{ID: 182, Name: "Ogoja", StateID: 9},
		{ID: 183, Name: "Yakurr", StateID: 9},
		{ID: 184, Name: "Yala", StateID: 9},

		//Delta
		{ID: 185, Name: "Aniocha North", StateID: 10},
		{ID: 186, Name: "Aniocha South", StateID: 10},
		{ID: 187, Name: "Bomadi", StateID: 10},
		{ID: 188, Name: "Burutu", StateID: 10},
		{ID: 189, Name: "Ethiope East", StateID: 10},
		{ID: 190, Name: "Ethiope West", StateID: 10},
		{ID: 191, Name: "Ika North East", StateID: 10},
		{ID: 192, Name: "Ika South", StateID: 10},
		{ID: 193, Name: "Isoko North", StateID: 10},
		{ID: 194, Name: "Isoko South", StateID: 10},
		{ID: 195, Name: "Ndokwa East", StateID: 10},
		{ID: 196, Name: "Ndokwa West", StateID: 10},
		{ID: 197, Name: "Okpe", StateID: 10},
		{ID: 198, Name: "Oshimili North", StateID: 10},
		{ID: 199, Name: "Oshimili South", StateID: 10},
		{ID: 200, Name: "Patani", StateID: 10},
		{ID: 201, Name: "Sapele", StateID: 10},
		{ID: 202, Name: "Udu", StateID: 10},
		{ID: 203, Name: "Ughelli North", StateID: 10},
		{ID: 204, Name: "Ughelli South", StateID: 10},
		{ID: 205, Name: "Ukwuani", StateID: 10},
		{ID: 206, Name: "Uvwie", StateID: 10},
		{ID: 207, Name: "Warri North", StateID: 10},
		{ID: 208, Name: "Warri South", StateID: 10},
		{ID: 209, Name: "Warri South West", StateID: 10},

		//Ebonyi
		{ID: 210, Name: "Abakaliki", StateID: 11},
		{ID: 211, Name: "Afikpo North", StateID: 11},
		{ID: 212, Name: "Afikpo South", StateID: 11},
		{ID: 213, Name: "Ebonyi", StateID: 11},
		{ID: 214, Name: "Ezza North", StateID: 11},
		{ID: 215, Name: "Ezza South", StateID: 11},
		{ID: 216, Name: "Ikwo", StateID: 11},
		{ID: 217, Name: "Ishielu", StateID: 11},
		{ID: 218, Name: "Ivo", StateID: 11},
		{ID: 219, Name: "Izzi", StateID: 11},
		{ID: 220, Name: "Ohaozara", StateID: 11},
		{ID: 221, Name: "Ohaukwu", StateID: 11},
		{ID: 222, Name: "Onicha", StateID: 11},

		//Edo
		{ID: 223, Name: "Akoko Edo", StateID: 12},
		{ID: 224, Name: "Egor", StateID: 12},
		{ID: 225, Name: "Esan Central", StateID: 12},
		{ID: 226, Name: "Esan North-East", StateID: 12},
		{ID: 227, Name: "Esan South-East", StateID: 12},
		{ID: 228, Name: "Esan West", StateID: 12},
		{ID: 229, Name: "Etsako Central", StateID: 12},
		{ID: 230, Name: "Etsako East", StateID: 12},
		{ID: 231, Name: "Etsako West", StateID: 12},
		{ID: 232, Name: "Igueben", StateID: 12},
		{ID: 233, Name: "Ikpoba-Okha", StateID: 12},
		{ID: 234, Name: "Oredo", StateID: 12},
		{ID: 235, Name: "Orhionmwon", StateID: 12},
		{ID: 236, Name: "Ovia North-East", StateID: 12},
		{ID: 237, Name: "Ovia South-West", StateID: 12},
		{ID: 238, Name: "Owan East", StateID: 12},
		{ID: 239, Name: "Owan West", StateID: 12},
		{ID: 240, Name: "Uhunmwonde", StateID: 12},

		//Ekiti
		{ID: 241, Name: "Ado Ekiti", StateID: 13},
		{ID: 242, Name: "Efon", StateID: 13},
		{ID: 243, Name: "Ekiti East", StateID: 13},
		{ID: 244, Name: "Ekiti South-West", StateID: 13},
		{ID: 245, Name: "Ekiti West", StateID: 13},
		{ID: 246, Name: "Emure", StateID: 13},
		{ID: 247, Name: "Gbonyin", StateID: 13},
		{ID: 248, Name: "Ido Osi", StateID: 13},
		{ID: 249, Name: "Ijero", StateID: 13},
		{ID: 250, Name: "Ikere", StateID: 13},
		{ID: 251, Name: "Ikole", StateID: 13},
		{ID: 252, Name: "Ilejemeje", StateID: 13},
		{ID: 253, Name: "Irepodun/Ifelodun", StateID: 13},
		{ID: 254, Name: "Ise/Orun", StateID: 13},
		{ID: 255, Name: "Moba", StateID: 13},
		{ID: 256, Name: "Oye", StateID: 13},

		//Enugu
		{ID: 257, Name: "Aninri", StateID: 14},
		{ID: 258, Name: "Awgu", StateID: 14},
		{ID: 259, Name: "Enugu East", StateID: 14},
		{ID: 260, Name: "Enugu North", StateID: 14},
		{ID: 261, Name: "Enugu South", StateID: 14},
		{ID: 262, Name: "Ezeagu", StateID: 14},
		{ID: 263, Name: "Igbo Etiti", StateID: 14},
		{ID: 264, Name: "Igbo Eze North", StateID: 14},
		{ID: 265, Name: "Igbo Eze South", StateID: 14},
		{ID: 266, Name: "Isi Uzo", StateID: 14},
		{ID: 267, Name: "Nkanu East", StateID: 14},
		{ID: 268, Name: "Nkanu West", StateID: 14},
		{ID: 269, Name: "Nsukka", StateID: 14},
		{ID: 270, Name: "Oji River", StateID: 14},
		{ID: 271, Name: "Udenu", StateID: 14},
		{ID: 272, Name: "Udi", StateID: 14},
		{ID: 273, Name: "Uzo Uwani", StateID: 14},

		//Gombe
		{ID: 274, Name: "Akko", StateID: 15},
		{ID: 275, Name: "Balanga", StateID: 15},
		{ID: 276, Name: "Billiri", StateID: 15},
		{ID: 277, Name: "Dukku", StateID: 15},
		{ID: 278, Name: "Funakaye", StateID: 15},
		{ID: 279, Name: "Gombe", StateID: 15},
		{ID: 280, Name: "Kaltungo", StateID: 15},
		{ID: 281, Name: "Kwami", StateID: 15},
		{ID: 282, Name: "Nafada", StateID: 15},
		{ID: 283, Name: "Shongom", StateID: 15},
		{ID: 284, Name: "Yamaltu/Deba", StateID: 15},

		//Imo
		{ID: 285, Name: "Aboh Mbaise", StateID: 16},
		{ID: 286, Name: "Ahiazu Mbaise", StateID: 16},
		{ID: 287, Name: "Ehime Mbano", StateID: 16},
		{ID: 288, Name: "Ezinihitte", StateID: 16},
		{ID: 289, Name: "Ideato North", StateID: 16},
		{ID: 290, Name: "Ideato South", StateID: 16},
		{ID: 291, Name: "Ihitte/Uboma", StateID: 16},
		{ID: 292, Name: "Ikeduru", StateID: 16},
		{ID: 293, Name: "Isiala Mbano", StateID: 16},
		{ID: 294, Name: "Isu", StateID: 16},
		{ID: 295, Name: "Mbaitoli", StateID: 16},
		{ID: 296, Name: "Ngor Okpala", StateID: 16},
		{ID: 297, Name: "Njaba", StateID: 16},
		{ID: 298, Name: "Nkwerre", StateID: 16},
		{ID: 299, Name: "Nwangele", StateID: 16},
		{ID: 300, Name: "Obowo", StateID: 16},
		{ID: 301, Name: "Oguta", StateID: 16},
		{ID: 302, Name: "Ohaji/Egbema", StateID: 16},
		{ID: 303, Name: "Okigwe", StateID: 16},
		{ID: 304, Name: "Onuimo", StateID: 16},
		{ID: 305, Name: "Orlu", StateID: 16},
		{ID: 306, Name: "Orsu", StateID: 16},
		{ID: 307, Name: "Oru East", StateID: 16},
		{ID: 308, Name: "Oru West", StateID: 16},
		{ID: 309, Name: "Owerri Municipal", StateID: 16},
		{ID: 310, Name: "Owerri North", StateID: 16},
		{ID: 311, Name: "Owerri West", StateID: 16},

		//Jigawa
		{ID: 312, Name: "Auyo", StateID: 17},
		{ID: 313, Name: "Babura", StateID: 17},
		{ID: 314, Name: "Biriniwa", StateID: 17},
		{ID: 315, Name: "Birnin Kudu", StateID: 17},
		{ID: 316, Name: "Buji", StateID: 17},
		{ID: 317, Name: "Dutse", StateID: 17},
		{ID: 318, Name: "Gagarawa", StateID: 17},
		{ID: 319, Name: "Garki", StateID: 17},
		{ID: 320, Name: "Gumel", StateID: 17},
		{ID: 321, Name: "Guri", StateID: 17},
		{ID: 322, Name: "Gwaram", StateID: 17},
		{ID: 323, Name: "Gwiwa", StateID: 17},
		{ID: 324, Name: "Hadejia", StateID: 17},
		{ID: 325, Name: "Jahun", StateID: 17},
		{ID: 326, Name: "Kafin Hausa", StateID: 17},
		{ID: 327, Name: "Kaugama", StateID: 17},
		{ID: 328, Name: "Kazaure", StateID: 17},
		{ID: 329, Name: "Kiri Kasama", StateID: 17},
		{ID: 330, Name: "Kiyawa", StateID: 17},
		{ID: 331, Name: "Maigatari", StateID: 17},
		{ID: 332, Name: "Malam Madori", StateID: 17},
		{ID: 333, Name: "Miga", StateID: 17},
		{ID: 334, Name: "Ringim", StateID: 17},
		{ID: 335, Name: "Roni", StateID: 17},
		{ID: 336, Name: "Sule Tankarkar", StateID: 17},
		{ID: 337, Name: "Taura", StateID: 17},
		{ID: 338, Name: "Yankwashi", StateID: 17},

		//Kaduna
		{ID: 339, Name: "Birnin Gwari", StateID: 18},
		{ID: 340, Name: "Chikun", StateID: 18},
		{ID: 341, Name: "Giwa", StateID: 18},
		{ID: 342, Name: "Igabi", StateID: 18},
		{ID: 343, Name: "Ikara", StateID: 18},
		{ID: 344, Name: "Jaba", StateID: 18},
		{ID: 345, Name: "Jema'a", StateID: 18},
		{ID: 346, Name: "Kachia", StateID: 18},
		{ID: 347, Name: "Kaduna North", StateID: 18},
		{ID: 348, Name: "Kaduna South", StateID: 18},
		{ID: 349, Name: "Kagarko", StateID: 18},
		{ID: 350, Name: "Kajuru", StateID: 18},
		{ID: 351, Name: "Kaura", StateID: 18},
		{ID: 352, Name: "Kauru", StateID: 18},
		{ID: 353, Name: "Kubau", StateID: 18},
		{ID: 354, Name: "Kudan", StateID: 18},
		{ID: 355, Name: "Lere", StateID: 18},
		{ID: 356, Name: "Makarfi", StateID: 18},
		{ID: 357, Name: "Sabon Gari", StateID: 18},
		{ID: 358, Name: "Sanga", StateID: 18},
		{ID: 359, Name: "Soba", StateID: 18},
		{ID: 360, Name: "Zangon Kataf", StateID: 18},
		{ID: 361, Name: "Zaria", StateID: 18},

		//Kano
		{ID: 362, Name: "Ajingi", StateID: 19},
		{ID: 363, Name: "Albasu", StateID: 19},
		{ID: 364, Name: "Bagwai", StateID: 19},
		{ID: 365, Name: "Bebeji", StateID: 19},
		{ID: 366, Name: "Bichi", StateID: 19},
		{ID: 367, Name: "Bunkure", StateID: 19},
		{ID: 368, Name: "Dala", StateID: 19},
		{ID: 369, Name: "Dambatta", StateID: 19},
		{ID: 370, Name: "Dawakin Kudu", StateID: 19},
		{ID: 371, Name: "Dawakin Tofa", StateID: 19},
		{ID: 372, Name: "Doguwa", StateID: 19},
		{ID: 373, Name: "Fagge", StateID: 19},
		{ID: 374, Name: "Gabasawa", StateID: 19},
		{ID: 375, Name: "Garko", StateID: 19},
		{ID: 376, Name: "Garun Mallam", StateID: 19},
		{ID: 377, Name: "Gaya", StateID: 19},
		{ID: 378, Name: "Gezawa", StateID: 19},
		{ID: 379, Name: "Gwale", StateID: 19},
		{ID: 380, Name: "Gwarzo", StateID: 19},
		{ID: 381, Name: "Kabo", StateID: 19},
		{ID: 382, Name: "Kano Municipal", StateID: 19},
		{ID: 383, Name: "Karaye", StateID: 19},
		{ID: 384, Name: "Kibiya", StateID: 19},
		{ID: 385, Name: "Kiru", StateID: 19},
		{ID: 386, Name: "Kumbotso", StateID: 19},
		{ID: 387, Name: "Kunchi", StateID: 19},
		{ID: 388, Name: "Kura", StateID: 19},
		{ID: 389, Name: "Madobi", StateID: 19},
		{ID: 390, Name: "Makoda", StateID: 19},
		{ID: 391, Name: "Minjibir", StateID: 19},
		{ID: 392, Name: "Nasarawa", StateID: 19},
		{ID: 393, Name: "Rano", StateID: 19},
		{ID: 394, Name: "Rimin Gado", StateID: 19},
		{ID: 395, Name: "Rogo", StateID: 19},
		{ID: 396, Name: "Shanono", StateID: 19},
		{ID: 397, Name: "Sumaila", StateID: 19},
		{ID: 398, Name: "Takai", StateID: 19},
		{ID: 399, Name: "Tarauni", StateID: 19},
		{ID: 400, Name: "Tofa", StateID: 19},
		{ID: 401, Name: "Tsanyawa", StateID: 19},
		{ID: 402, Name: "Tudun Wada", StateID: 19},
		{ID: 403, Name: "Ungogo", StateID: 19},
		{ID: 404, Name: "Warawa", StateID: 19},
		{ID: 405, Name: "Wudil", StateID: 19},
		
		
		//Katsina
		{ID: 406, Name: "Bakori", StateID: 20},
		{ID: 407, Name: "Batagarawa", StateID: 20},
		{ID: 408, Name: "Batsari", StateID: 20},
		{ID: 409, Name: "Baure", StateID: 20},
		{ID: 410, Name: "Bindawa", StateID: 20},
		{ID: 411, Name: "Charanchi", StateID: 20},
		{ID: 412, Name: "Dandume", StateID: 20},
		{ID: 413, Name: "Danja", StateID: 20},
		{ID: 414, Name: "Dan Musa", StateID: 20},
		{ID: 415, Name: "Daura", StateID: 20},
		{ID: 416, Name: "Dutsi", StateID: 20},
		{ID: 417, Name: "Dutsin Ma", StateID: 20},
		{ID: 418, Name: "Faskari", StateID: 20},
		{ID: 419, Name: "Funtua", StateID: 20},
		{ID: 420, Name: "Ingawa", StateID: 20},
		{ID: 421, Name: "Jibia", StateID: 20},
		{ID: 422, Name: "Kafur", StateID: 20},
		{ID: 423, Name: "Kaita", StateID: 20},
		{ID: 424, Name: "Kankara", StateID: 20},
		{ID: 425, Name: "Kankia", StateID: 20},
		{ID: 426, Name: "Katsina", StateID: 20},
		{ID: 427, Name: "Kurfi", StateID: 20},
		{ID: 428, Name: "Kusada", StateID: 20},
		{ID: 429, Name: "Mai'Adua", StateID: 20},
		{ID: 430, Name: "Malumfashi", StateID: 20},
		{ID: 431, Name: "Mani", StateID: 20},
		{ID: 432, Name: "Mashi", StateID: 20},
		{ID: 433, Name: "Matazu", StateID: 20},
		{ID: 434, Name: "Musawa", StateID: 20},
		{ID: 435, Name: "Rimi", StateID: 20},
		{ID: 436, Name: "Sabuwa", StateID: 20},
		{ID: 437, Name: "Safana", StateID: 20},
		{ID: 438, Name: "Sandamu", StateID: 20},
		{ID: 439, Name: "Zango", StateID: 20},

		//Kebbi
		{ID: 440, Name: "Aleiro", StateID: 21},
		{ID: 441, Name: "Arewa Dandi", StateID: 21},
		{ID: 442, Name: "Argungu", StateID: 21},
		{ID: 443, Name: "Augie", StateID: 21},
		{ID: 444, Name: "Bagudo", StateID: 21},
		{ID: 445, Name: "Birnin Kebbi", StateID: 21},
		{ID: 446, Name: "Bunza", StateID: 21},
		{ID: 447, Name: "Dandi", StateID: 21},
		{ID: 448, Name: "Fakai", StateID: 21},
		{ID: 449, Name: "Gwandu", StateID: 21},
		{ID: 450, Name: "Jega", StateID: 21},
		{ID: 451, Name: "Kalgo", StateID: 21},
		{ID: 452, Name: "Koko/Besse", StateID: 21},
		{ID: 453, Name: "Maiyama", StateID: 21},
		{ID: 454, Name: "Ngaski", StateID: 21},
		{ID: 455, Name: "Sakaba", StateID: 21},
		{ID: 456, Name: "Shanga", StateID: 21},
		{ID: 457, Name: "Suru", StateID: 21},
		{ID: 458, Name: "Yauri", StateID: 21},
		{ID: 459, Name: "Zuru", StateID: 21},

		//Kogi
		{ID: 460, Name: "Adavi", StateID: 22},
		{ID: 461, Name: "Ajaokuta", StateID: 22},
		{ID: 462, Name: "Ankpa", StateID: 22},
		{ID: 463, Name: "Bassa", StateID: 22},
		{ID: 464, Name: "Dekina", StateID: 22},
		{ID: 465, Name: "Ibaji", StateID: 22},
		{ID: 466, Name: "Igalamela Odolu", StateID: 22},
		{ID: 467, Name: "Idah", StateID: 22},
		{ID: 468, Name: "Ijumu", StateID: 22},
		{ID: 469, Name: "Kabba/Bunu", StateID: 22},
		{ID: 470, Name: "Kogi", StateID: 22},
		{ID: 471, Name: "Mopa-Muro", StateID: 22},
		{ID: 472, Name: "Ofu", StateID: 22},
		{ID: 473, Name: "Okehi", StateID: 22},
		{ID: 474, Name: "Olamaboro", StateID: 22},
		{ID: 475, Name: "Omala", StateID: 22},
		{ID: 476, Name: "Yagba East", StateID: 22},
		{ID: 477, Name: "Yagba West", StateID: 22},

		//Kwara
		{ID: 478, Name: "Asa", StateID: 23},
		{ID: 479, Name: "Baruten", StateID: 23},
		{ID: 480, Name: "Ekiti", StateID: 23},
		{ID: 481, Name: "Ifelodun", StateID: 23},
		{ID: 482, Name: "Ilorin East", StateID: 23},
		{ID: 483, Name: "Ilorin South", StateID: 23},
		{ID: 484, Name: "Ilorin West", StateID: 23},
		{ID: 485, Name: "Irepodun", StateID: 23},
		{ID: 486, Name: "Ishin", StateID: 23},
		{ID: 487, Name: "Offa", StateID: 23},
		{ID: 488, Name: "Oke Ero", StateID: 23},
		{ID: 489, Name: "Oyun", StateID: 23},
		{ID: 490, Name: "Pategi", StateID: 23},
		
		//Lagos
		{ID: 491, Name: "Ajeromi-Ifelodun", StateID: 24},
		{ID: 492, Name: "Alimosho", StateID: 24},
		{ID: 493, Name: "Apapa", StateID: 24},
		{ID: 494, Name: "Badagry", StateID: 24},
		{ID: 495, Name: "Epe", StateID: 24},
		{ID: 496, Name: "Ibeju-Lekki", StateID: 24},
		{ID: 497, Name: "Ifako-Ijaiye", StateID: 24},
		{ID: 498, Name: "Ikeja", StateID: 24},
		{ID: 499, Name: "Ikorodu", StateID: 24},
		{ID: 500, Name: "Lagos Island", StateID: 24},
		{ID: 501, Name: "Lagos Mainland", StateID: 24},
		{ID: 502, Name: "Lagos South-West", StateID: 24},
		{ID: 503, Name: "Lagos West", StateID: 24},
		{ID: 504, Name: "Mushin", StateID: 24},
		{ID: 505, Name: "Ojo", StateID: 24},
		{ID: 506, Name: "Oshodi-Isolo", StateID: 24},
		{ID: 507, Name: "Shomolu", StateID: 24},
		{ID: 508, Name: "Surulere", StateID: 24},
		{ID: 509, Name: "Akwanga", StateID: 25},
		{ID: 510, Name: "Doma", StateID: 25},
		{ID: 511, Name: "Karu", StateID: 25},
		{ID: 512, Name: "Keffi", StateID: 25},
		{ID: 513, Name: "Kokona", StateID: 25},
		{ID: 514, Name: "Nasarawa", StateID: 25},
		{ID: 515, Name: "Nasarawa Egon", StateID: 25},
		{ID: 516, Name: "Obi", StateID: 25},
		{ID: 517, Name: "Toto", StateID: 25},
		{ID: 518, Name: "Wamba", StateID: 25},
		{ID: 519, Name: "Agaie", StateID: 26},
		{ID: 520, Name: "Agwara", StateID: 26},
		{ID: 521, Name: "Bida", StateID: 26},
		{ID: 522, Name: "Borgu", StateID: 26},
		{ID: 523, Name: "Bosso", StateID: 26},
		{ID: 524, Name: "Chanchaga", StateID: 26},
		{ID: 525, Name: "Edati", StateID: 26},
		{ID: 526, Name: "Gbako", StateID: 26},
		{ID: 527, Name: "Katcha", StateID: 26},
		{ID: 528, Name: "Kewo", StateID: 26},
		{ID: 529, Name: "Kontagora", StateID: 26},
		{ID: 530, Name: "Mokwa", StateID: 26},
		{ID: 531, Name: "Mashegu", StateID: 26},
		{ID: 532, Name: "Munya", StateID: 26},
		{ID: 533, Name: "Paikoro", StateID: 26},
		{ID: 534, Name: "Rafi", StateID: 26},
		{ID: 535, Name: "Rijau", StateID: 26},
		{ID: 536, Name: "Shiroro", StateID: 26},
		{ID: 537, Name: "Suleja", StateID: 26},
		{ID: 538, Name: "Wushishi", StateID: 26},
		{ID: 539, Name: "Abeokuta North", StateID: 27},
		{ID: 540, Name: "Abeokuta South", StateID: 27},
		{ID: 541, Name: "Ado-Odo/Ota", StateID: 27},
		{ID: 542, Name: "Ewekoro", StateID: 27},
		{ID: 543, Name: "Ifo", StateID: 27},
		{ID: 544, Name: "Ijebu East", StateID: 27},
		{ID: 545, Name: "Ijebu North", StateID: 27},
		{ID: 546, Name: "Ijebu North East", StateID: 27},
		{ID: 547, Name: "Ijebu Ode", StateID: 27},
		{ID: 548, Name: "Ikenne", StateID: 27},
		{ID: 549, Name: "Imeko Afon", StateID: 27},
		{ID: 550, Name: "Ipokia", StateID: 27},
		{ID: 551, Name: "Obafemi Owode", StateID: 27},
		{ID: 552, Name: "Odeda", StateID: 27},
		{ID: 553, Name: "Odogbolu", StateID: 27},
		{ID: 554, Name: "Ogun Waterside", StateID: 27},
		{ID: 555, Name: "Remo North", StateID: 27},
		{ID: 556, Name: "Remo South", StateID: 27},
		{ID: 557, Name: "Yewa North", StateID: 27},
		{ID: 558, Name: "Yewa South", StateID: 27},
		{ID: 559, Name: "Akoko North-East", StateID: 28},
		{ID: 560, Name: "Akoko North-West", StateID: 28},
		{ID: 561, Name: "Akoko South-East", StateID: 28},
		{ID: 562, Name: "Akoko South-West", StateID: 28},
		{ID: 563, Name: "Akure North", StateID: 28},
		{ID: 564, Name: "Akure South", StateID: 28},
		{ID: 565, Name: "Ese-Odo", StateID: 28},
		{ID: 566, Name: "Idanre", StateID: 28},
		{ID: 567, Name: "Ifedore", StateID: 28},
		{ID: 568, Name: "Ilaje", StateID: 28},
		{ID: 569, Name: "Ileoluji/Okeigbo", StateID: 28},
		{ID: 570, Name: "Odigbo", StateID: 28},
		{ID: 571, Name: "Okitipupa", StateID: 28},
		{ID: 572, Name: "Ondo East", StateID: 28},
		{ID: 573, Name: "Ondo West", StateID: 28},
		{ID: 574, Name: "Ose", StateID: 28},
		{ID: 575, Name: "Owo", StateID: 28},
		{ID: 576, Name: "Atakunmosa East", StateID: 29},
		{ID: 577, Name: "Atakunmosa West", StateID: 29},
		{ID: 578, Name: "Boluwaduro", StateID: 29},
		{ID: 579, Name: "Boluwaduro", StateID: 29},
		{ID: 580, Name: "Boripe", StateID: 29},
		{ID: 581, Name: "Ede North", StateID: 29},
		{ID: 582, Name: "Ede South", StateID: 29},
		{ID: 583, Name: "Ife Central", StateID: 29},
		{ID: 584, Name: "Ife East", StateID: 29},
		{ID: 585, Name: "Ife North", StateID: 29},
		{ID: 586, Name: "Ife South", StateID: 29},
		{ID: 587, Name: "Ilesa East", StateID: 29},
		{ID: 588, Name: "Ilesa West", StateID: 29},
		{ID: 589, Name: "Ilesa West", StateID: 29},
		{ID: 590, Name: "Ilesa East", StateID: 29},
		{ID: 591, Name: "Irepodun", StateID: 29},
		{ID: 592, Name: "Irewole", StateID: 29},
		{ID: 593, Name: "Isokan", StateID: 29},
		{ID: 594, Name: "Orolu", StateID: 29},
		{ID: 595, Name: "Oshogbo", StateID: 29},
		{ID: 596, Name: "Akinyele", StateID: 30},
		{ID: 597, Name: "Atiba", StateID: 30},
		{ID: 598, Name: "Atisbo", StateID: 30},
		{ID: 599, Name: "Egbeda", StateID: 30},
		{ID: 600, Name: "Ibadan North", StateID: 30},
		{ID: 601, Name: "Ibadan North-East", StateID: 30},
		{ID: 602, Name: "Ibadan North-West", StateID: 30},
		{ID: 603, Name: "Ibadan South-East", StateID: 30},
		{ID: 604, Name: "Ibadan South-West", StateID: 30},
		{ID: 605, Name: "Ibarapa Central", StateID: 30},
		{ID: 606, Name: "Ibarapa East", StateID: 30},
		{ID: 607, Name: "Ibarapa North", StateID: 30},
		{ID: 608, Name: "Iseyin", StateID: 30},
		{ID: 609, Name: "Itesiwaju", StateID: 30},
		{ID: 610, Name: "Kajola", StateID: 30},
		{ID: 611, Name: "Lagelu", StateID: 30},
		{ID: 612, Name: "Ogbomosho North", StateID: 30},
		{ID: 613, Name: "Ogbomosho South", StateID: 30},
		{ID: 614, Name: "Olorunsogo", StateID: 30},
		{ID: 615, Name: "Oriire", StateID: 30},
		{ID: 616, Name: "Oyo East", StateID: 30},
		{ID: 617, Name: "Oyo West", StateID: 30},
		{ID: 618, Name: "Saki East", StateID: 30},
		{ID: 619, Name: "Saki West", StateID: 30},
		{ID: 620, Name: "Barkin Ladi", StateID: 31},
		{ID: 621, Name: "Bassa", StateID: 31},
		{ID: 622, Name: "Bokkos", StateID: 31},
		{ID: 623, Name: "Kanam", StateID: 31},
		{ID: 624, Name: "Kanke", StateID: 31},
		{ID: 625, Name: "Langtang North", StateID: 31},
		{ID: 626, Name: "Langtang South", StateID: 31},
		{ID: 627, Name: "Mangu", StateID: 31},
		{ID: 628, Name: "Pankshin", StateID: 31},
		{ID: 629, Name: "Qua'an Pan", StateID: 31},
		{ID: 630, Name: "Riyom", StateID: 31},
		{ID: 631, Name: "Shendam", StateID: 31},
		{ID: 632, Name: "Wase", StateID: 31},
		{ID: 633, Name: "Abua/Odual", StateID: 32},
		{ID: 634, Name: "Ahoada East", StateID: 32},
		{ID: 635, Name: "Ahoada West", StateID: 32},
		{ID: 636, Name: "Akuku-Toru", StateID: 32},
		{ID: 637, Name: "Andoni", StateID: 32},
		{ID: 638, Name: "Asari-Toru", StateID: 32},
		{ID: 639, Name: "Bonny", StateID: 32},
		{ID: 640, Name: "Degema", StateID: 32},
		{ID: 641, Name: "Emohua", StateID: 32},
		{ID: 642, Name: "Etche", StateID: 32},
		{ID: 643, Name: "Gokana", StateID: 32},
		{ID: 644, Name: "Khana", StateID: 32},
		{ID: 645, Name: "Obio/Akpor", StateID: 32},
		{ID: 646, Name: "Ogba/Egbema/Ndoni", StateID: 32},
		{ID: 647, Name: "Ogu/Bolo", StateID: 32},
		{ID: 648, Name: "Okrika", StateID: 32},
		{ID: 649, Name: "Okrika", StateID: 32},
		{ID: 650, Name: "Omuma", StateID: 32},
		{ID: 651, Name: "Opobo/Nkoro", StateID: 32},
		{ID: 652, Name: "Port Harcourt", StateID: 32},
		{ID: 653, Name: "Tai", StateID: 32},
		{ID: 654, Name: "Binji", StateID: 33},
		{ID: 655, Name: "Bodinga", StateID: 33},
		{ID: 656, Name: "Dange/Shuni", StateID: 33},
		{ID: 657, Name: "Gada", StateID: 33},
		{ID: 658, Name: "Goronyo", StateID: 33},
		{ID: 659, Name: "Gudu", StateID: 33},
		{ID: 660, Name: "Gwadabawa", StateID: 33},
		{ID: 661, Name: "Illela", StateID: 33},
		{ID: 662, Name: "Kebbe", StateID: 33},
		{ID: 663, Name: "Kware", StateID: 33},
		{ID: 664, Name: "Rabah", StateID: 33},
		{ID: 665, Name: "Sokoto North", StateID: 33},
		{ID: 666, Name: "Sokoto South", StateID: 33},
		{ID: 667, Name: "Tambuwal", StateID: 33},
		{ID: 668, Name: "Tangaza", StateID: 33},
		{ID: 669, Name: "Tureta", StateID: 33},
		{ID: 670, Name: "Wamako", StateID: 33},
		{ID: 671, Name: "Wurno", StateID: 33},
		{ID: 672, Name: "Yabo", StateID: 33},
		{ID: 673, Name: "Ardo-Kola", StateID: 34},
		{ID: 674, Name: "Bali", StateID: 34},
		{ID: 675, Name: "Donga", StateID: 34},
		{ID: 676, Name: "Gashaka", StateID: 34},
		{ID: 677, Name: "Gassol", StateID: 34},
		{ID: 678, Name: "Ibi", StateID: 34},
		{ID: 679, Name: "Jalingo", StateID: 34},
		{ID: 680, Name: "Kasim", StateID: 34},
		{ID: 681, Name: "Lau", StateID: 34},
		{ID: 682, Name: "Niger", StateID: 34},
		{ID: 683, Name: "Sardauna", StateID: 34},
		{ID: 684, Name: "Takum", StateID: 34},
		{ID: 685, Name: "Ussa", StateID: 34},
		{ID: 686, Name: "Wukari", StateID: 34},
		{ID: 687, Name: "Yorro", StateID: 34},
		{ID: 688, Name: "Bade", StateID: 35},
		{ID: 689, Name: "Bursari", StateID: 35},
		{ID: 690, Name: "Damaturu", StateID: 35},
		{ID: 691, Name: "Fika", StateID: 35},
		{ID: 692, Name: "Fune", StateID: 35},
		{ID: 693, Name: "Geidam", StateID: 35},
		{ID: 694, Name: "Gujba", StateID: 35},
		{ID: 695, Name: "Gulani", StateID: 35},
		{ID: 696, Name: "Jakusko", StateID: 35},
		{ID: 697, Name: "Karasuwa", StateID: 35},
		{ID: 698, Name: "Machina", StateID: 35},
		{ID: 699, Name: "Nangere", StateID: 35},
		{ID: 700, Name: "Nguru", StateID: 35},
		{ID: 701, Name: "Potiskum", StateID: 35},
		{ID: 702, Name: "Tarmuwa", StateID: 35},
		{ID: 703, Name: "Yunusari", StateID: 35},
		{ID: 704, Name: "Yusufari", StateID: 35},
		{ID: 705, Name: "Anka", StateID: 36},
		{ID: 706, Name: "Bakura", StateID: 36},
		{ID: 707, Name: "Birnin Magaji", StateID: 36},
		{ID: 708, Name: "Bukkuyum", StateID: 36},
		{ID: 709, Name: "Bungudu", StateID: 36},
		{ID: 710, Name: "Gummi", StateID: 36},
		{ID: 711, Name: "Gusau", StateID: 36},
		{ID: 712, Name: "Kaura Namoda", StateID: 36},
		{ID: 713, Name: "Maradun", StateID: 36},
		{ID: 714, Name: "Maru", StateID: 36},
		{ID: 715, Name: "Shinkafi", StateID: 36},
		{ID: 716, Name: "Talata Mafara", StateID: 36},
		{ID: 717, Name: "Zurmi", StateID: 36},
		{ID: 718, Name: "Abaji", StateID: 37},
		{ID: 719, Name: "Abuja Municipal", StateID: 37},
		{ID: 720, Name: "Bwari", StateID: 37},
		{ID: 721, Name: "Gwagwalada", StateID: 37},
		{ID: 722, Name: "Kuje", StateID: 37},
		{ID: 723, Name: "Kwali", StateID: 37},
	}

	Towns = []models.Town{
		// Aba North
		{ID: 1, Name: "Eziama", LGID: 1},
		{ID: 2, Name: "Uratta", LGID: 1},
		{ID: 3, Name: "Osusu", LGID: 1},
		{ID: 4, Name: "Umuola", LGID: 1},
		{ID: 5, Name: "Umuogor", LGID: 1},
	
		// Aba South
		{ID: 6, Name: "Ogbor Hill", LGID: 2},
		{ID: 7, Name: "Umungasi", LGID: 2},
		{ID: 8, Name: "Abayi", LGID: 2},
		{ID: 9, Name: "Ohabiam", LGID: 2},
		{ID: 10, Name: "Asa", LGID: 2},
	
		// Arochukwu
		{ID: 11, Name: "Arochukwu", LGID: 3},
		{ID: 12, Name: "Utughugwu", LGID: 3},
		{ID: 13, Name: "Amuvi", LGID: 3},
		{ID: 14, Name: "Atani", LGID: 3},
		{ID: 15, Name: "Ihechiowa", LGID: 3},
	
		// Bende
		{ID: 16, Name: "Bende", LGID: 4},
		{ID: 17, Name: "Uzuakoli", LGID: 4},
		{ID: 18, Name: "Item", LGID: 4},
		{ID: 19, Name: "Alayi", LGID: 4},
		{ID: 20, Name: "Akara", LGID: 4},
		{ID: 21, Name: "Igbere", LGID: 4},
		{ID: 22, Name: "Ozuitem", LGID: 4},
		{ID: 23, Name: "Itumbauzo", LGID: 4},
		{ID: 79, Name: "Umuhu", LGID: 4 },
	
		// Ikwuano
		{ID: 24, Name: "Ikwuano", LGID: 5},
		{ID: 25, Name: "Ndoro", LGID: 5},
		{ID: 26, Name: "Ibere", LGID: 5},
		{ID: 27, Name: "Oloko", LGID: 5},
		{ID: 28, Name: "Oboro", LGID: 5},
		{ID: 80, Name: "Umudike", LGID: 5},
	
		// Isiala Ngwa North
		{ID: 29, Name: "Isiala Ngwa", LGID: 6},
		{ID: 30, Name: "Osokwa", LGID: 6},
		{ID: 31, Name: "Umuoba", LGID: 6},
		{ID: 32, Name: "Amapu Ntigha", LGID: 6},
		{ID: 33, Name: "Okpuala Ngwa", LGID: 6},
	
		// Isiala Ngwa South
		{ID: 34, Name: "Mbutu", LGID: 7},
		{ID: 35, Name: "Ovu", LGID: 7},
		{ID: 36, Name: "Okpuala", LGID: 7},
		{ID: 37, Name: "Umumaru", LGID: 7},
		{ID: 38, Name: "Omoba", LGID: 7},
	
		// Isuikwuato
		{ID: 39, Name: "Isuikwuato", LGID: 8},
		{ID: 40, Name: "Uturu", LGID: 8},
		{ID: 41, Name: "Ovim", LGID: 8},
		{ID: 42, Name: "Amaba", LGID: 8},
		{ID: 43, Name: "Ndiawa", LGID: 8},
	
		// Obingwa
		{ID: 44, Name: "Ngwa", LGID: 9},
		{ID: 45, Name: "Aba", LGID: 9},
		{ID: 46, Name: "Ohazu", LGID: 9},
		{ID: 47, Name: "Umuoha", LGID: 9},
		{ID: 48, Name: "Uratta", LGID: 9},
	
		// Ohafia
		{ID: 49, Name: "Ohafia", LGID: 10},
		{ID: 50, Name: "Elu", LGID: 10},
		{ID: 51, Name: "Eziukwu", LGID: 10},
		{ID: 52, Name: "Amangwu", LGID: 10},
		{ID: 53, Name: "Nkporo", LGID: 10},
	
		// Osisioma
		{ID: 54, Name: "Osisioma", LGID: 11},
		{ID: 55, Name: "Amanator", LGID: 11},
		{ID: 56, Name: "Aro-Ngwa", LGID: 11},
		{ID: 57, Name: "Ekeakpara", LGID: 11},
		{ID: 58, Name: "Nkwuohia", LGID: 11},
	
		// Ugwunagbo
		{ID: 59, Name: "Ugwunagbo", LGID: 12},
		{ID: 60, Name: "Aba", LGID: 12},
		{ID: 61, Name: "Obuzor", LGID: 12},
		{ID: 62, Name: "Amaekpu", LGID: 12},
		{ID: 63, Name: "Ossah", LGID: 12},
	
		// Umuahia North
		{ID: 64, Name: "Umuahia", LGID: 13},
		{ID: 65, Name: "Isingwu", LGID: 13},
		{ID: 66, Name: "Aboh", LGID: 13},
		{ID: 67, Name: "Umuagu", LGID: 13},
		{ID: 68, Name: "Nkwoegwu", LGID: 13},
	
		// Umuahia South
		{ID: 69, Name: "Afara", LGID: 14},
		{ID: 70, Name: "Olokoro", LGID: 14},
		{ID: 71, Name: "Umuopara", LGID: 14},
		{ID: 72, Name: "Ezeleke", LGID: 14},
		{ID: 73, Name: "Omaghu", LGID: 14},
	
		// Umu Nneochi
		{ID: 74, Name: "Nneochi", LGID: 15},
		{ID: 75, Name: "Nkwa", LGID: 15},
		{ID: 76, Name: "Mbala", LGID: 15},
		{ID: 77, Name: "Amuda", LGID: 15},
		{ID: 78, Name: "Ngodo", LGID: 15},

		// Demsa (LGID: 16)
		{ID: 81, Name: "Demsa", LGID: 16},
		{ID: 82, Name: "Borrong", LGID: 16},
		{ID: 83, Name: "Mbula", LGID: 16},
		{ID: 84, Name: "Nassarawo Demsa", LGID: 16},
		{ID: 85, Name: "Bille", LGID: 16},
		{ID: 86, Name: "Dong", LGID: 16},
		{ID: 87, Name: "Kiri", LGID: 16},

		// Fufore (LGID: 17)
		{ID: 88, Name: "Fufore", LGID: 17},
		{ID: 89, Name: "Gurin", LGID: 17},
		{ID: 90, Name: "Verre", LGID: 17},
		{ID: 91, Name: "Malabu", LGID: 17},
		{ID: 92, Name: "Pariya", LGID: 17},
		{ID: 93, Name: "Ribadu", LGID: 17},

		// Ganye (LGID: 18)
		{ID: 94, Name: "Ganye", LGID: 18},
		{ID: 95, Name: "Sugu", LGID: 18},
		{ID: 96, Name: "Jaggu", LGID: 18},
		{ID: 97, Name: "Gamu", LGID: 18},
		{ID: 98, Name: "Gorong", LGID: 18},

		// Girei (LGID: 19)
		{ID: 99, Name: "Girei", LGID: 19},
		{ID: 100, Name: "Damare", LGID: 19},
		{ID: 101, Name: "Jabbi Lamba", LGID: 19},
		{ID: 102, Name: "Viniklang", LGID: 19},
		{ID: 103, Name: "Wuro Dole", LGID: 19},

		// Gombi (LGID: 20)
		{ID: 104, Name: "Gombi", LGID: 20},
		{ID: 105, Name: "Garkida", LGID: 20},
		{ID: 106, Name: "Ga'anda", LGID: 20},
		{ID: 107, Name: "Lala", LGID: 20},
		{ID: 108, Name: "Tawa", LGID: 20},

		// Guyuk (LGID: 21)
		{ID: 109, Name: "Guyuk", LGID: 21},
		{ID: 110, Name: "Chikila", LGID: 21},
		{ID: 111, Name: "Dumna", LGID: 21},
		{ID: 112, Name: "Banjiram", LGID: 21},
		{ID: 113, Name: "Lokoro", LGID: 21},

		// Hong (LGID: 22)
		{ID: 114, Name: "Hong", LGID: 22},
		{ID: 115, Name: "Garaha", LGID: 22},
		{ID: 116, Name: "Hildi", LGID: 22},
		{ID: 117, Name: "Kwajaffa", LGID: 22},
		{ID: 118, Name: "Mijili", LGID: 22},

		// Jada (LGID: 23)
		{ID: 119, Name: "Jada", LGID: 23},
		{ID: 120, Name: "Leko", LGID: 23},
		{ID: 121, Name: "Nyibango", LGID: 23},
		{ID: 122, Name: "Mbulo", LGID: 23},
		{ID: 123, Name: "Ganye", LGID: 23},

		// Lamurde (LGID: 24)
		{ID: 124, Name: "Lamurde", LGID: 24},
		{ID: 125, Name: "Gyamzo", LGID: 24},
		{ID: 126, Name: "Mbamnga", LGID: 24},
		{ID: 127, Name: "Opalo", LGID: 24},
		{ID: 128, Name: "Sabon Gari", LGID: 24},

		// Madagali (LGID: 25)
		{ID: 129, Name: "Madagali", LGID: 25},
		{ID: 130, Name: "Gulak", LGID: 25},
		{ID: 131, Name: "Shuwa", LGID: 25},
		{ID: 132, Name: "Bitta", LGID: 25},
		{ID: 133, Name: "Kirchinga", LGID: 25},

		// Maiha (LGID: 26)
		{ID: 134, Name: "Maiha", LGID: 26},
		{ID: 135, Name: "Belel", LGID: 26},
		{ID: 136, Name: "Mayo Ngulnyaki", LGID: 26},
		{ID: 137, Name: "Pakara", LGID: 26},

		// Mubi North (LGID: 27)
		{ID: 138, Name: "Mubi North", LGID: 27},
		{ID: 139, Name: "Vimtim", LGID: 27},
		{ID: 140, Name: "Bahuli", LGID: 27},
		{ID: 141, Name: "Muchalla", LGID: 27},
		{ID: 142, Name: "Sukundi", LGID: 27},

		// Mubi South (LGID: 28)
		{ID: 143, Name: "Mubi South", LGID: 28},
		{ID: 144, Name: "Gella", LGID: 28},
		{ID: 145, Name: "Lamorde", LGID: 28},
		{ID: 146, Name: "Kwaja", LGID: 28},
		{ID: 147, Name: "Nassarawo Mubi", LGID: 28},

		// Numan (LGID: 29)
		{ID: 148, Name: "Numan", LGID: 29},
		{ID: 149, Name: "Dong", LGID: 29},
		{ID: 150, Name: "Zumo", LGID: 29},
		{ID: 151, Name: "Gaiya", LGID: 29},
		{ID: 152, Name: "Bare", LGID: 29},

		// Shelleng (LGID: 30)
		{ID: 153, Name: "Shelleng", LGID: 30},
		{ID: 154, Name: "Bachure", LGID: 30},
		{ID: 155, Name: "Gundo", LGID: 30},
		{ID: 156, Name: "Jumbul", LGID: 30},
		{ID: 157, Name: "Njuwan", LGID: 30},

		// Song (LGID: 31)
		{ID: 158, Name: "Song", LGID: 31},
		{ID: 159, Name: "Zumo", LGID: 31},
		{ID: 160, Name: "Gudu Mboi", LGID: 31},
		{ID: 161, Name: "Dumne", LGID: 31},
		{ID: 162, Name: "Kilange", LGID: 31},

		// Toungo (LGID: 32)
		{ID: 163, Name: "Toungo", LGID: 32},
		{ID: 164, Name: "Kiri", LGID: 32},
		{ID: 165, Name: "Dembere", LGID: 32},
		{ID: 166, Name: "Tigari", LGID: 32},
		{ID: 167, Name: "Gari Sarki", LGID: 32},

		// Yola North (LGID: 33)
		{ID: 168, Name: "Yola", LGID: 33},
		{ID: 169, Name: "Jimeta", LGID: 33},
		{ID: 170, Name: "Yolde Pate", LGID: 33},
		{ID: 171, Name: "Doubeli", LGID: 33},
		{ID: 172, Name: "Bako", LGID: 33},

		// Yola South (LGID: 34)
		{ID: 173, Name: "Yola Town", LGID: 34},
		{ID: 174, Name: "Namtari", LGID: 34},
		{ID: 175, Name: "Bole", LGID: 34},
		{ID: 176, Name: "Makama", LGID: 34},
		{ID: 177, Name: "Mbamba", LGID: 34},

				// Abak (LGID: 35)
		{ID: 178, Name: "Abak", LGID: 35},
		{ID: 179, Name: "Afaha Obong", LGID: 35},
		{ID: 180, Name: "Ikot Ekang", LGID: 35},
		{ID: 181, Name: "Ikot Oku Mfang", LGID: 35},
		{ID: 182, Name: "Otoro", LGID: 35},

		// Eastern Obolo (LGID: 36)
		{ID: 183, Name: "Okoroete", LGID: 36},
		{ID: 184, Name: "Iko", LGID: 36},
		{ID: 185, Name: "Amadaka", LGID: 36},
		{ID: 186, Name: "Okiuso", LGID: 36},
		{ID: 187, Name: "Elile", LGID: 36},

		// Eket (LGID: 37)
		{ID: 188, Name: "Eket", LGID: 37},
		{ID: 189, Name: "Esit Urua", LGID: 37},
		{ID: 190, Name: "Idua", LGID: 37},
		{ID: 191, Name: "Mkpok", LGID: 37},
		{ID: 192, Name: "Okon", LGID: 37},

		// Esit Eket (LGID: 38)
		{ID: 193, Name: "Uquo", LGID: 38},
		{ID: 194, Name: "Etebi", LGID: 38},
		{ID: 195, Name: "Ikot Abasi", LGID: 38},
		{ID: 196, Name: "Ikpa", LGID: 38},
		{ID: 197, Name: "Odoro Nkit", LGID: 38},

		// Essien Udim (LGID: 39)
		{ID: 198, Name: "Essien Udim", LGID: 39},
		{ID: 199, Name: "Afaha Ikot Ebak", LGID: 39},
		{ID: 200, Name: "Ekpeyong", LGID: 39},
		{ID: 201, Name: "Ikpe Annang", LGID: 39},
		{ID: 202, Name: "Ikot Ibritam", LGID: 39},

		// Etim Ekpo (LGID: 40)
		{ID: 203, Name: "Etim Ekpo", LGID: 40},
		{ID: 204, Name: "Utu", LGID: 40},
		{ID: 205, Name: "Uruk Ata", LGID: 40},
		{ID: 206, Name: "Ikot Akpan Anwa", LGID: 40},
		{ID: 207, Name: "Ikot Udo", LGID: 40},

		// Etinan (LGID: 41)
		{ID: 208, Name: "Etinan", LGID: 41},
		{ID: 209, Name: "Ikot Ebo", LGID: 41},
		{ID: 210, Name: "Ikot Ekang", LGID: 41},
		{ID: 211, Name: "Ikot Ekwere", LGID: 41},
		{ID: 212, Name: "Ikot Udo", LGID: 41},

		// Ibeno (LGID: 42)
		{ID: 213, Name: "Ibeno", LGID: 42},
		{ID: 214, Name: "Mkpanak", LGID: 42},
		{ID: 215, Name: "Okoroitak", LGID: 42},
		{ID: 216, Name: "Atabrikang", LGID: 42},
		{ID: 217, Name: "Eket", LGID: 42},

		// Ibesikpo Asutan (LGID: 43)
		{ID: 218, Name: "Nung Udoe", LGID: 43},
		{ID: 219, Name: "Ibesikpo", LGID: 43},
		{ID: 220, Name: "Ikot Akpa Edet", LGID: 43},
		{ID: 221, Name: "Ikot Ekwere", LGID: 43},
		{ID: 222, Name: "Ikot Ibritam", LGID: 43},

		// Ibiono-Ibom (LGID: 44)
		{ID: 223, Name: "Oko Ita", LGID: 44},
		{ID: 224, Name: "Ikot Akpan Obong", LGID: 44},
		{ID: 225, Name: "Ikot Ekpene", LGID: 44},
		{ID: 226, Name: "Ikot Ibritam", LGID: 44},
		{ID: 227, Name: "Ikot Obong", LGID: 44},

		// Ika (LGID: 45)
		{ID: 228, Name: "Urua Inyang", LGID: 45},
		{ID: 229, Name: "Achan", LGID: 45},
		{ID: 230, Name: "Ikot Abia", LGID: 45},
		{ID: 231, Name: "Ikot Akpan", LGID: 45},
		{ID: 232, Name: "Ikot Obio Inyang", LGID: 45},

		// Ikono (LGID: 46)
		{ID: 233, Name: "Ibiaku Ntok Okpo", LGID: 46},
		{ID: 234, Name: "Ikot Ekpene", LGID: 46},
		{ID: 235, Name: "Ikot Obio", LGID: 46},
		{ID: 236, Name: "Mbiabong", LGID: 46},
		{ID: 237, Name: "Nung Udoe", LGID: 46},

		// Ikot Abasi (LGID: 47)
		{ID: 238, Name: "Ikot Abasi", LGID: 47},
		{ID: 239, Name: "Ikot Akpaden", LGID: 47},
		{ID: 240, Name: "Ikot Okoro", LGID: 47},
		{ID: 241, Name: "Ikot Ubo", LGID: 47},
		{ID: 242, Name: "Ndiya", LGID: 47},

		// Ikot Ekpene (LGID: 48)
		{ID: 243, Name: "Ikot Ekpene", LGID: 48},
		{ID: 244, Name: "Amanyan", LGID: 48},
		{ID: 245, Name: "Ikot Abasi", LGID: 48},
		{ID: 246, Name: "Ikot Akpa", LGID: 48},
		{ID: 247, Name: "Ikot Ekpene", LGID: 48},

		// Ini (LGID: 49)
		{ID: 248, Name: "Ikot Ada Idem", LGID: 49},
		{ID: 249, Name: "Ikot Ekpene", LGID: 49},
		{ID: 250, Name: "Ikot Ibritam", LGID: 49},
		{ID: 251, Name: "Ikot Obio", LGID: 49},
		{ID: 252, Name: "Nnung Udoe", LGID: 49},

		// Itu (LGID: 50)
		{ID: 253, Name: "Itu", LGID: 50},
		{ID: 254, Name: "Ikot Ada Idem", LGID: 50},
		{ID: 255, Name: "Ikot Ekpene", LGID: 50},
		{ID: 256, Name: "Ikot Ibritam", LGID: 50},
		{ID: 257, Name: "Mbiabong", LGID: 50},

		// Mbo (LGID: 51)
		{ID: 258, Name: "Enwang", LGID: 51},
		{ID: 259, Name: "Ebughu", LGID: 51},
		{ID: 260, Name: "Effiat", LGID: 51},
		{ID: 261, Name: "Ewang", LGID: 51},
		{ID: 262, Name: "Uda", LGID: 51},

		// Mkpat-Enin (LGID: 52)
		{ID: 263, Name: "Mkpat Enin", LGID: 52},
		{ID: 264, Name: "Ikot Akpaden", LGID: 52},
		{ID: 265, Name: "Ikot Ekefre", LGID: 52},
		{ID: 266, Name: "Ikot Esien", LGID: 52},
		{ID: 267, Name: "Ikot Ubo", LGID: 52},

		// Nsit-Atai (LGID: 53)
		{ID: 268, Name: "Okoro Atai", LGID: 53},
		{ID: 269, Name: "Ikot Ntuen", LGID: 53},
		{ID: 270, Name: "Ikot Obong", LGID: 53},
		{ID: 271, Name: "Ikot Udo", LGID: 53},
		{ID: 272, Name: "Mbiabong", LGID: 53},

		// Nsit-Ibom (LGID: 54)
		{ID: 273, Name: "Afaha Offiong", LGID: 54},
		{ID: 274, Name: "Asang", LGID: 54},
		{ID: 275, Name: "Ikot Ada Idem", LGID: 54},
		{ID: 276, Name: "Ikot Akpan", LGID: 54},
		{ID: 277, Name: "Ikot Ekpene", LGID: 54},

		// Nsit-Ubium (LGID: 55)
		{ID: 278, Name: "Ikot Edibon", LGID: 55},
		{ID: 279, Name: "Ikot Esien", LGID: 55},
		{ID: 280, Name: "Ikot Nyo", LGID: 55},
		{ID: 281, Name: "Ikot Obio", LGID: 55},
		{ID: 282, Name: "Ikot Ubo", LGID: 55},

		// Obot Akara (LGID: 56)
		{ID: 283, Name: "Nto Edino", LGID: 56},
		{ID: 284, Name: "Ikot Ekpene", LGID: 56},
		{ID: 285, Name: "Ikot Ibritam", LGID: 56},
		{ID: 286, Name: "Ikot Okoro", LGID: 56},
		{ID: 287, Name: "Ndiya", LGID: 56},

		// Okobo (LGID: 57)
		{ID: 288, Name: "Okopedi", LGID: 57},
		{ID: 289, Name: "Ikot Akpa", LGID: 57},
		{ID: 290, Name: "Ikot Abasi", LGID: 57},
		{ID: 291, Name: "Ikot Ada Idem", LGID: 57},
		{ID: 292, Name: "Ikot Akpan", LGID: 57},

		// Onna (LGID: 58)
		{ID: 293, Name: "Abat", LGID: 58},
		{ID: 294, Name: "Ikot Akpa", LGID: 58},
		{ID: 295, Name: "Ikot Ekang", LGID: 58},
		{ID: 296, Name: "Ikot Ibritam", LGID: 58},
		{ID: 297, Name: "Ikot Ubo", LGID: 58},

		// Oron (LGID: 61)
		{ID: 298, Name: "Oron", LGID: 61},
		{ID: 299, Name: "Enwang", LGID: 61},
		{ID: 300, Name: "Ikot Ada Idem", LGID: 61},
		{ID: 301, Name: "Ikot Ekpene", LGID: 61},
		{ID: 302, Name: "Ikot Ibritam", LGID: 61},

		// Oruk Anam (LGID: 62)
		{ID: 303, Name: "Ikot Ibritam", LGID: 62},
		{ID: 304, Name: "Ikot Obio", LGID: 62},
		{ID: 305, Name: "Ndiya", LGID: 62},
		{ID: 306, Name: "Nung Udoe", LGID: 62},
		{ID: 307, Name: "Otoro", LGID: 62},

		// Udung-Uko (LGID: 63)
		{ID: 308, Name: "Udung Uko", LGID: 63},
		{ID: 309, Name: "Ikot Abasi", LGID: 63},
		{ID: 310, Name: "Ikot Ekpene", LGID: 63},
		{ID: 311, Name: "Ikot Ibritam", LGID: 63},
		{ID: 312, Name: "Ikot Obio", LGID: 63},

		// Ukanafun (LGID: 64)
		{ID: 313, Name: "Ikot Akpa Nkuk", LGID: 64},
		{ID: 314, Name: "Ikot Ibritam", LGID: 64},
		{ID: 315, Name: "Ikot Obio", LGID: 64},
		{ID: 316, Name: "Ikot Okoro", LGID: 64},
		{ID: 317, Name: "Otoro", LGID: 64},

		// Uruan (LGID: 65)
		{ID: 318, Name: "Idu", LGID: 65},
		{ID: 319, Name: "Ikot Akpa Nkuk", LGID: 65},
		{ID: 320, Name: "Ikot Ibritam", LGID: 65},
		{ID: 321, Name: "Ikot Obio", LGID: 65},
		{ID: 322, Name: "Ikot Okoro", LGID: 65},

		// Urue-Offong/Oruko (LGID: 66)
		{ID: 323, Name: "Urue Offong", LGID: 66},
		{ID: 324, Name: "Oruko", LGID: 66},
		{ID: 325, Name: "Ikot Akpa Nkuk", LGID: 66},
		{ID: 326, Name: "Ikot Ibritam", LGID: 66},
		{ID: 327, Name: "Ikot Obio", LGID: 66},

		// Uyo (LGID: 67)
		{ID: 328, Name: "Uyo", LGID: 67},
		{ID: 329, Name: "Ewet Housing Estate", LGID: 67},
		{ID: 330, Name: "Ikot Ekpene", LGID: 67},
		{ID: 331, Name: "Ikot Ibritam", LGID: 67},
		{ID: 332, Name: "Ikot Obio", LGID: 67},

				// Aguata (LGID: 68)
		{ID: 333, Name: "Ekwulobia", LGID: 68},
		{ID: 334, Name: "Igbo-Ukwu", LGID: 68},
		{ID: 335, Name: "Isuofia", LGID: 68},
		{ID: 336, Name: "Aguluezechukwu", LGID: 68},
		{ID: 337, Name: "Ezinifite", LGID: 68},

		// Anambra East (LGID: 69)
		{ID: 338, Name: "Otuocha", LGID: 69},
		{ID: 339, Name: "Aguleri", LGID: 69},
		{ID: 340, Name: "Umuleri", LGID: 69},
		{ID: 341, Name: "Nsugbe", LGID: 69},
		{ID: 342, Name: "Nando", LGID: 69},

		// Anambra West (LGID: 70)
		{ID: 343, Name: "Nzam", LGID: 70},
		{ID: 344, Name: "Olumbanasa", LGID: 70},
		{ID: 345, Name: "Igbokenyi", LGID: 70},
		{ID: 346, Name: "Mmiata", LGID: 70},
		{ID: 347, Name: "Umueze Anam", LGID: 70},

		// Anaocha (LGID: 71)
		{ID: 348, Name: "Neni", LGID: 71},
		{ID: 349, Name: "Agulu", LGID: 71},
		{ID: 350, Name: "Adazi Nnukwu", LGID: 71},
		{ID: 351, Name: "Adazi Ani", LGID: 71},
		{ID: 352, Name: "Adazi Enu", LGID: 71},

		// Awka North (LGID: 72)
		{ID: 353, Name: "Achalla", LGID: 72},
		{ID: 354, Name: "Amansea", LGID: 72},
		{ID: 355, Name: "Amanuke", LGID: 72},
		{ID: 356, Name: "Ebonyi", LGID: 72},
		{ID: 357, Name: "Isu Aniocha", LGID: 72},

		// Awka South (LGID: 73 & 74)
		{ID: 358, Name: "Awka", LGID: 73},
		{ID: 359, Name: "Amawbia", LGID: 73},
		{ID: 360, Name: "Nibo", LGID: 73},
		{ID: 361, Name: "Nise", LGID: 73},
		{ID: 362, Name: "Okpuno", LGID: 73},
		{ID: 363, Name: "Umuokpu", LGID: 74},

		// Dunukofia (LGID: 75)
		{ID: 364, Name: "Ukpo", LGID: 75},
		{ID: 365, Name: "Ifitedunu", LGID: 75},
		{ID: 366, Name: "Umunnachi", LGID: 75},
		{ID: 367, Name: "Dunukofia", LGID: 75},
		{ID: 368, Name: "Ukpo", LGID: 75},

		// Ekwusigo (LGID: 76)
		{ID: 369, Name: "Ozubulu", LGID: 76},
		{ID: 370, Name: "Ihembosi", LGID: 76},
		{ID: 371, Name: "Oraifite", LGID: 76},
		{ID: 372, Name: "Uruagu", LGID: 76},
		{ID: 373, Name: "Uruagu", LGID: 76},

		// Idemili North (LGID: 77)
		{ID: 374, Name: "Ogidi", LGID: 77},
		{ID: 375, Name: "Nkpor", LGID: 77},
		{ID: 376, Name: "Obosi", LGID: 77},
		{ID: 377, Name: "Umuoji", LGID: 77},
		{ID: 378, Name: "Eziowelle", LGID: 77},

		// Idemili South (LGID: 78)
		{ID: 379, Name: "Ojoto", LGID: 78},
		{ID: 380, Name: "Nnobi", LGID: 78},
		{ID: 381, Name: "Alor", LGID: 78},
		{ID: 382, Name: "Awka-Etiti", LGID: 78},
		{ID: 383, Name: "Nnewi", LGID: 78},

		// Ihiala (LGID: 79)
		{ID: 384, Name: "Ihiala", LGID: 79},
		{ID: 385, Name: "Uli", LGID: 79},
		{ID: 386, Name: "Okija", LGID: 79},
		{ID: 387, Name: "Amorka", LGID: 79},
		{ID: 388, Name: "Isseke", LGID: 79},

		// Njikoka (LGID: 80)
		{ID: 389, Name: "Abagana", LGID: 80},
		{ID: 390, Name: "Enugwu-Ukwu", LGID: 80},
		{ID: 391, Name: "Nawfia", LGID: 80},
		{ID: 392, Name: "Nimo", LGID: 80},
		{ID: 393, Name: "Enugwu-Agidi", LGID: 80},

		// Nnewi North (LGID: 81)
		{ID: 394, Name: "Nnewi", LGID: 81},
		{ID: 395, Name: "Nnewi-Ichi", LGID: 81},
		{ID: 396, Name: "Otolo", LGID: 81},
		{ID: 397, Name: "Uruagu", LGID: 81},
		{ID: 398, Name: "Umudim", LGID: 81},

		// Nnewi South (LGID: 82)
		{ID: 399, Name: "Ukpor", LGID: 82},
		{ID: 400, Name: "Akwaihedi", LGID: 82},
		{ID: 401, Name: "Ezinifite", LGID: 82},
		{ID: 402, Name: "Osumenyi", LGID: 82},
		{ID: 403, Name: "Amichi", LGID: 82},

		// Ogbaru (LGID: 83)
		{ID: 404, Name: "Atani", LGID: 83},
		{ID: 405, Name: "Okpoko", LGID: 83},
		{ID: 406, Name: "Odekpe", LGID: 83},
		{ID: 407, Name: "Ogbakuba", LGID: 83},
		{ID: 408, Name: "Akili-Ogidi", LGID: 83},

		// Onitsha North (LGID: 84)
		{ID: 409, Name: "Onitsha", LGID: 84},
		{ID: 410, Name: "Odoakpu", LGID: 84},
		{ID: 411, Name: "Woliwo", LGID: 84},
		{ID: 412, Name: "GRA", LGID: 84},
		{ID: 413, Name: "Ogbunike", LGID: 84},

		// Onitsha South (LGID: 85)
		{ID: 414, Name: "Fegge", LGID: 85},
		{ID: 415, Name: "Woliwo", LGID: 85},
		{ID: 416, Name: "Odoakpu", LGID: 85},
		{ID: 417, Name: "Ogbunike", LGID: 85},
		{ID: 418, Name: "Okpoko", LGID: 85},

		// Orumba North (LGID: 86)
		{ID: 419, Name: "Ajalli", LGID: 86},
		{ID: 420, Name: "Awgbu", LGID: 86},
		{ID: 421, Name: "Amaokpala", LGID: 86},
		{ID: 422, Name: "Okoh", LGID: 86},
		{ID: 423, Name: "Ufuma", LGID: 86},

		// Orumba South (LGID: 87)
		{ID: 424, Name: "Umunze", LGID: 87},
		{ID: 425, Name: "Ogbunka", LGID: 87},
		{ID: 426, Name: "Eziagu", LGID: 87},
		{ID: 427, Name: "Isulo", LGID: 87},
		{ID: 428, Name: "Owerre-Ezukala", LGID: 87},

		// Oyi (LGID: 88)
		{ID: 429, Name: "Nteje", LGID: 88},
		{ID: 430, Name: "Umunya", LGID: 88},
		{ID: 431, Name: "Ogbunike", LGID: 88},
		{ID: 432, Name: "Nkwelle-Ezunaka", LGID: 88},
		{ID: 433, Name: "Awkuzu", LGID: 88},

		// Alkaleri (LGID: 89)
		{ID: 434, Name: "Alkaleri", LGID: 89},
		{ID: 435, Name: "Gwaram", LGID: 89},
		{ID: 436, Name: "Maimadi", LGID: 89},
		{ID: 437, Name: "Yuli", LGID: 89},
		{ID: 438, Name: "Futuk", LGID: 89},

		// Bauchi (LGID: 90)
		{ID: 439, Name: "Bauchi", LGID: 90},
		{ID: 440, Name: "Gwallaga", LGID: 90},
		{ID: 441, Name: "Yelwa", LGID: 90},
		{ID: 442, Name: "Miri", LGID: 90},
		{ID: 443, Name: "Dungal", LGID: 90},

		// Bogoro (LGID: 91)
		{ID: 444, Name: "Bogoro", LGID: 91},
		{ID: 445, Name: "Lusa", LGID: 91},
		{ID: 446, Name: "Lere", LGID: 91},
		{ID: 447, Name: "Boi", LGID: 91},
		{ID: 448, Name: "Gurara", LGID: 91},

		// Damban (LGID: 92)
		{ID: 449, Name: "Damban", LGID: 92},
		{ID: 450, Name: "Tafare", LGID: 92},
		{ID: 451, Name: "Zadawa", LGID: 92},
		{ID: 452, Name: "Yelwa", LGID: 92},
		{ID: 453, Name: "Gar", LGID: 92},

		// Darazo (LGID: 93)
		{ID: 454, Name: "Darazo", LGID: 93},
		{ID: 455, Name: "Lame", LGID: 93},
		{ID: 456, Name: "Gabarin", LGID: 93},
		{ID: 457, Name: "Yautare", LGID: 93},
		{ID: 458, Name: "Larzani", LGID: 93},

		// Dass (LGID: 94)
		{ID: 459, Name: "Dass", LGID: 94},
		{ID: 460, Name: "Bajar", LGID: 94},
		{ID: 461, Name: "Durr", LGID: 94},
		{ID: 462, Name: "Zumbur", LGID: 94},
		{ID: 463, Name: "Bununu", LGID: 94},

		// Gamawa (LGID: 95)
		{ID: 464, Name: "Gamawa", LGID: 95},
		{ID: 465, Name: "Gololo", LGID: 95},
		{ID: 466, Name: "Dabino", LGID: 95},
		{ID: 467, Name: "Gadiya", LGID: 95},
		{ID: 468, Name: "Kafin Madaki", LGID: 95},

		// Ganjuwa (LGID: 96)
		{ID: 469, Name: "Ganjuwa", LGID: 96},
		{ID: 470, Name: "Kubi", LGID: 96},
		{ID: 471, Name: "Gungura", LGID: 96},
		{ID: 472, Name: "Kariya", LGID: 96},
		{ID: 473, Name: "Yali", LGID: 96},

		// Giade (LGID: 97)
		{ID: 474, Name: "Giade", LGID: 97},
		{ID: 475, Name: "Zira", LGID: 97},
		{ID: 476, Name: "Zubo", LGID: 97},
		{ID: 477, Name: "Doga", LGID: 97},
		{ID: 478, Name: "Wawa", LGID: 97},

		// Itas/Gadau (LGID: 98)
		{ID: 479, Name: "Itas", LGID: 98},
		{ID: 480, Name: "Gadau", LGID: 98},
		{ID: 481, Name: "Zuburna", LGID: 98},
		{ID: 482, Name: "Bera", LGID: 98},
		{ID: 483, Name: "Akuyam", LGID: 98},

		// Jama'are (LGID: 99)
		{ID: 484, Name: "Jama'are", LGID: 99},
		{ID: 485, Name: "Hanafari", LGID: 99},
		{ID: 486, Name: "Dogon Jeji", LGID: 99},
		{ID: 487, Name: "Jabuli", LGID: 99},
		{ID: 488, Name: "Gusmau", LGID: 99},

		// Katagum (LGID: 100)
		{ID: 489, Name: "Katagum", LGID: 100},
		{ID: 490, Name: "Azare", LGID: 100},
		{ID: 491, Name: "Madara", LGID: 100},
		{ID: 492, Name: "Gabarin", LGID: 100},
		{ID: 493, Name: "Sakawa", LGID: 100},

		// Kirfi (LGID: 101)
		{ID: 494, Name: "Kirfi", LGID: 101},
		{ID: 495, Name: "Tonga", LGID: 101},
		{ID: 496, Name: "Badara", LGID: 101},
		{ID: 497, Name: "Gobirawa", LGID: 101},
		{ID: 498, Name: "Cheledi", LGID: 101},

		// Misau (LGID: 102)
		{ID: 499, Name: "Misau", LGID: 102},
		{ID: 500, Name: "Gwaram", LGID: 102},
		{ID: 501, Name: "Hardawa", LGID: 102},
		{ID: 502, Name: "Gabra", LGID: 102},
		{ID: 503, Name: "Ajalli", LGID: 102},

		// Ningi (LGID: 103)
		{ID: 504, Name: "Ningi", LGID: 103},
		{ID: 505, Name: "Diga", LGID: 103},
		{ID: 506, Name: "Samu", LGID: 103},
		{ID: 507, Name: "Dutsen Kura", LGID: 103},
		{ID: 508, Name: "Balma", LGID: 103},

		// Shira (LGID: 104)
		{ID: 509, Name: "Yana", LGID: 104},
		{ID: 510, Name: "Jama'are", LGID: 104},
		{ID: 511, Name: "Tsakuwa", LGID: 104},
		{ID: 512, Name: "Jagiwa", LGID: 104},
		{ID: 513, Name: "Dumburum", LGID: 104},

		// Tafawa Balewa (LGID: 105)
		{ID: 514, Name: "Tafawa Balewa", LGID: 105},
		{ID: 515, Name: "Lere", LGID: 105},
		{ID: 516, Name: "Bula", LGID: 105},
		{ID: 517, Name: "Dajin", LGID: 105},
		{ID: 518, Name: "Bangwa", LGID: 105},

		// Toro (LGID: 106)
		{ID: 519, Name: "Toro", LGID: 106},
		{ID: 520, Name: "Rishi", LGID: 106},
		{ID: 521, Name: "Lame", LGID: 106},
		{ID: 522, Name: "Tilde", LGID: 106},
		{ID: 523, Name: "Jama'are", LGID: 106},

		// Warji (LGID: 107)
		{ID: 524, Name: "Warji", LGID: 107},
		{ID: 525, Name: "Dagu", LGID: 107},
		{ID: 526, Name: "Katanga", LGID: 107},
		{ID: 527, Name: "Banjiram", LGID: 107},
		{ID: 528, Name: "Gwaram", LGID: 107},

		// Zaki (LGID: 108)
		{ID: 529, Name: "Zaki", LGID: 108},
		{ID: 530, Name: "Gamawa", LGID: 108},
		{ID: 531, Name: "Gumai", LGID: 108},
		{ID: 532, Name: "Mashio", LGID: 108},
		{ID: 533, Name: "Madauchi", LGID: 108},

				// Brass (LGID: 109)
		{ID: 534, Name: "Brass", LGID: 109},
		{ID: 535, Name: "Odioma", LGID: 109},
		{ID: 536, Name: "Ogbokodo", LGID: 109},
		{ID: 537, Name: "Twon-Brass", LGID: 109},
		{ID: 538, Name: "Foropa", LGID: 109},

		// Ekeremor (LGID: 110)
		{ID: 539, Name: "Ekeremor", LGID: 110},
		{ID: 540, Name: "Oporoma", LGID: 110},
		{ID: 541, Name: "Peremabiri", LGID: 110},
		{ID: 542, Name: "Ogbia", LGID: 110},
		{ID: 543, Name: "Oporoma", LGID: 110},

		// Kolokuma/Opokuma (LGID: 111)
		{ID: 544, Name: "Kolokuma/Opokuma", LGID: 111},
		{ID: 545, Name: "Kolo Creek", LGID: 111},
		{ID: 546, Name: "Ebedebiri", LGID: 111},
		{ID: 547, Name: "Otuokpoti", LGID: 111},
		{ID: 548, Name: "Opokuma", LGID: 111},

		// Nembe (LGID: 112)
		{ID: 549, Name: "Nembe", LGID: 112},
		{ID: 550, Name: "Nembe City", LGID: 112},
		{ID: 551, Name: "Ogbia", LGID: 112},
		{ID: 552, Name: "Kolo Creek", LGID: 112},
		{ID: 553, Name: "Apoi", LGID: 112},

		// Ogbia (LGID: 113)
		{ID: 554, Name: "Ogbia", LGID: 113},
		{ID: 555, Name: "Ogbia Town", LGID: 113},
		{ID: 556, Name: "Ogbia Creek", LGID: 113},
		{ID: 557, Name: "Otuogori", LGID: 113},
		{ID: 558, Name: "Okpoama", LGID: 113},

		// Sagbama (LGID: 114)
		{ID: 559, Name: "Sagbama", LGID: 114},
		{ID: 560, Name: "Sagbama Town", LGID: 114},
		{ID: 561, Name: "Ekeremor", LGID: 114},
		{ID: 562, Name: "Ogboinbiri", LGID: 114},
		{ID: 563, Name: "Otuokpoti", LGID: 114},

		// Southern Ijaw (LGID: 115)
		{ID: 564, Name: "Southern Ijaw", LGID: 115},
		{ID: 565, Name: "Oporoma", LGID: 115},
		{ID: 566, Name: "Peremabiri", LGID: 115},
		{ID: 567, Name: "Yenagoa", LGID: 115},
		{ID: 568, Name: "Okpoama", LGID: 115},

		// Yenagoa (LGID: 116)
		{ID: 569, Name: "Yenagoa", LGID: 116},
		{ID: 570, Name: "Azikoro", LGID: 116},
		{ID: 571, Name: "Opolo", LGID: 116},
		{ID: 572, Name: "Edepie", LGID: 116},
		{ID: 573, Name: "Yenagoa Town", LGID: 116},

		// Ado (LGID: 117)
		{ID: 574, Name: "Ado", LGID: 117},
		{ID: 575, Name: "Agbado", LGID: 117},
		{ID: 576, Name: "Ado Town", LGID: 117},
		{ID: 577, Name: "Eke", LGID: 117},
		{ID: 578, Name: "Oju", LGID: 117},

		// Agatu (LGID: 118)
		{ID: 579, Name: "Agatu", LGID: 118},
		{ID: 580, Name: "Akwu", LGID: 118},
		{ID: 581, Name: "Ogbogu", LGID: 118},
		{ID: 582, Name: "Ugbokpo", LGID: 118},
		{ID: 583, Name: "Obagaji", LGID: 118},

		// Apa (LGID: 119)
		{ID: 584, Name: "Apa", LGID: 119},
		{ID: 585, Name: "Ugep", LGID: 119},
		{ID: 586, Name: "Adoka", LGID: 119},
		{ID: 587, Name: "Akaaka", LGID: 119},
		{ID: 588, Name: "Apir", LGID: 119},

		// Buruku (LGID: 120)
		{ID: 589, Name: "Buruku", LGID: 120},
		{ID: 590, Name: "Zaki-Biam", LGID: 120},
		{ID: 591, Name: "Ishere", LGID: 120},
		{ID: 592, Name: "Akperhe", LGID: 120},
		{ID: 593, Name: "Adikpo", LGID: 120},

		// Gboko (LGID: 121)
		{ID: 594, Name: "Gboko", LGID: 121},
		{ID: 595, Name: "Mbaty", LGID: 121},
		{ID: 596, Name: "Tse-Kucha", LGID: 121},
		{ID: 597, Name: "Ityuluv", LGID: 121},
		{ID: 598, Name: "Mbayion", LGID: 121},

		// Guma (LGID: 122)
		{ID: 599, Name: "Guma", LGID: 122},
		{ID: 600, Name: "Mbakur", LGID: 122},
		{ID: 601, Name: "Mbakon", LGID: 122},
		{ID: 602, Name: "Mbakpa", LGID: 122},
		{ID: 603, Name: "Katsina-Ala", LGID: 122},

		// Gwer East (LGID: 123)
		{ID: 604, Name: "Gwer East", LGID: 123},
		{ID: 605, Name: "Tse-Ginde", LGID: 123},
		{ID: 606, Name: "Nengak", LGID: 123},
		{ID: 607, Name: "Gwer", LGID: 123},
		{ID: 608, Name: "Tse-Kura", LGID: 123},

		// Gwer West (LGID: 124)
		{ID: 609, Name: "Gwer West", LGID: 124},
		{ID: 610, Name: "Naka", LGID: 124},
		{ID: 611, Name: "Guma", LGID: 124},
		{ID: 612, Name: "Mbaakpu", LGID: 124},
		{ID: 613, Name: "Tse-Ato", LGID: 124},

		// Katsina-Ala (LGID: 125)
		{ID: 614, Name: "Katsina-Ala", LGID: 125},
		{ID: 615, Name: "Ugbokpo", LGID: 125},
		{ID: 616, Name: "Ogbadibo", LGID: 125},
		{ID: 617, Name: "Ukum", LGID: 125},
		{ID: 618, Name: "Tse-Kura", LGID: 125},

		// Konshisha (LGID: 126)
		{ID: 619, Name: "Konshisha", LGID: 126},
		{ID: 620, Name: "Ogbom", LGID: 126},
		{ID: 621, Name: "Tse-Taru", LGID: 126},
		{ID: 622, Name: "Mbakor", LGID: 126},
		{ID: 623, Name: "Gboko", LGID: 126},

		// Kwande (LGID: 127)
		{ID: 624, Name: "Kwande", LGID: 127},
		{ID: 625, Name: "Ado", LGID: 127},
		{ID: 626, Name: "Adikpo", LGID: 127},
		{ID: 627, Name: "Mbagwa", LGID: 127},
		{ID: 628, Name: "Mbaku", LGID: 127},

		// Logo (LGID: 128)
		{ID: 629, Name: "Logo", LGID: 128},
		{ID: 630, Name: "Tse-Kura", LGID: 128},
		{ID: 631, Name: "Nengak", LGID: 128},
		{ID: 632, Name: "Mbaakpu", LGID: 128},
		{ID: 633, Name: "Guma", LGID: 128},

		// Makurdi (LGID: 129)
		{ID: 634, Name: "Makurdi", LGID: 129},
		{ID: 635, Name: "Wurukum", LGID: 129},
		{ID: 636, Name: "North Bank", LGID: 129},
		{ID: 637, Name: "Gboko", LGID: 129},
		{ID: 638, Name: "Naka", LGID: 129},

		// Obi (LGID: 130)
		{ID: 639, Name: "Obi", LGID: 130},
		{ID: 640, Name: "Oju", LGID: 130},
		{ID: 641, Name: "Ogbadibo", LGID: 130},
		{ID: 642, Name: "Otukpo", LGID: 130},
		{ID: 643, Name: "Ado", LGID: 130},

		// Ogbadibo (LGID: 131)
		{ID: 644, Name: "Ogbadibo", LGID: 131},
		{ID: 645, Name: "Otukpo", LGID: 131},
		{ID: 646, Name: "Ado", LGID: 131},
		{ID: 647, Name: "Okpokwu", LGID: 131},
		{ID: 648, Name: "Obi", LGID: 131},

		// Ohimini (LGID: 132)
		{ID: 649, Name: "Ohini", LGID: 132},
		{ID: 650, Name: "Ogbadibo", LGID: 132},
		{ID: 651, Name: "Ado", LGID: 132},
		{ID: 652, Name: "Otukpo", LGID: 132},
		{ID: 653, Name: "Obi", LGID: 132},

		// Oju (LGID: 133)
		{ID: 654, Name: "Oju", LGID: 133},
		{ID: 655, Name: "Ugbokpo", LGID: 133},
		{ID: 656, Name: "Adoka", LGID: 133},
		{ID: 657, Name: "Orokam", LGID: 133},
		{ID: 658, Name: "Ogbadibo", LGID: 133},

		// Okpokwu (LGID: 134)
		{ID: 659, Name: "Okpokwu", LGID: 134},
		{ID: 660, Name: "Adoka", LGID: 134},
		{ID: 661, Name: "Ugbokpo", LGID: 134},
		{ID: 662, Name: "Orokam", LGID: 134},
		{ID: 663, Name: "Oju", LGID: 134},

		// Otukpo (LGID: 135)
		{ID: 664, Name: "Otukpo", LGID: 135},
		{ID: 665, Name: "Ugbokpo", LGID: 135},
		{ID: 666, Name: "Adoka", LGID: 135},
		{ID: 667, Name: "Orokam", LGID: 135},
		{ID: 668, Name: "Oju", LGID: 135},

		// Tarka (LGID: 136)
		{ID: 669, Name: "Tarka", LGID: 136},
		{ID: 670, Name: "Mbake", LGID: 136},
		{ID: 671, Name: "Guma", LGID: 136},
		{ID: 672, Name: "Mbaakpu", LGID: 136},
		{ID: 673, Name: "Kwande", LGID: 136},

		// Ukum (LGID: 137)
		{ID: 674, Name: "Ukum", LGID: 137},
		{ID: 675, Name: "Tse-Ginde", LGID: 137},
		{ID: 676, Name: "Mbakur", LGID: 137},
		{ID: 677, Name: "Tse-Kura", LGID: 137},
		{ID: 678, Name: "Ado", LGID: 137},

		// Ushongo (LGID: 138)
		{ID: 679, Name: "Ushongo", LGID: 138},
		{ID: 680, Name: "Naka", LGID: 138},
		{ID: 681, Name: "Ogbadibo", LGID: 138},
		{ID: 682, Name: "Mbakor", LGID: 138},
		{ID: 683, Name: "Gboko", LGID: 138},

		// Vandeikya (LGID: 139)
		{ID: 684, Name: "Vandeikya", LGID: 139},
		{ID: 685, Name: "Mbakor", LGID: 139},
		{ID: 686, Name: "Mbaakpu", LGID: 139},
		{ID: 687, Name: "Kwande", LGID: 139},
		{ID: 688, Name: "Gwer", LGID: 139},

				// Abadam (LGID: 140)
		{ID: 689, Name: "Abadam", LGID: 140},
		{ID: 690, Name: "Kari", LGID: 140},
		{ID: 691, Name: "Ngala", LGID: 140},
		{ID: 692, Name: "Dikwa", LGID: 140},
		{ID: 693, Name: "Gubio", LGID: 140},

		// Askira/Uba (LGID: 141)
		{ID: 694, Name: "Askira/Uba", LGID: 141},
		{ID: 695, Name: "Uba", LGID: 141},
		{ID: 696, Name: "Askira", LGID: 141},
		{ID: 697, Name: "Gulani", LGID: 141},
		{ID: 698, Name: "Maiduguri", LGID: 141},

		// Bama (LGID: 142)
		{ID: 699, Name: "Bama", LGID: 142},
		{ID: 700, Name: "Bama Town", LGID: 142},
		{ID: 701, Name: "Kwaya Kusar", LGID: 142},
		{ID: 702, Name: "Gwoza", LGID: 142},
		{ID: 703, Name: "Dikwa", LGID: 142},

		// Bayo (LGID: 143)
		{ID: 704, Name: "Bayo", LGID: 143},
		{ID: 705, Name: "Bayo Town", LGID: 143},
		{ID: 706, Name: "Wuyo", LGID: 143},
		{ID: 707, Name: "Kukawa", LGID: 143},
		{ID: 708, Name: "Mafa", LGID: 143},

		// Biu (LGID: 144)
		{ID: 709, Name: "Biu", LGID: 144},
		{ID: 710, Name: "Biu Town", LGID: 144},
		{ID: 711, Name: "Biu North", LGID: 144},
		{ID: 712, Name: "Biu South", LGID: 144},
		{ID: 713, Name: "Gubio", LGID: 144},

		// Chibok (LGID: 145)
		{ID: 714, Name: "Chibok", LGID: 145},
		{ID: 715, Name: "Chibok Town", LGID: 145},
		{ID: 716, Name: "Gwoza", LGID: 145},
		{ID: 717, Name: "Damboa", LGID: 145},
		{ID: 718, Name: "Maiduguri", LGID: 145},

		// Damboa (LGID: 146)
		{ID: 719, Name: "Damboa", LGID: 146},
		{ID: 720, Name: "Damboa Town", LGID: 146},
		{ID: 721, Name: "Gwoza", LGID: 146},
		{ID: 722, Name: "Maiduguri", LGID: 146},
		{ID: 723, Name: "Chibok", LGID: 146},

		// Dikwa (LGID: 147)
		{ID: 724, Name: "Dikwa", LGID: 147},
		{ID: 725, Name: "Dikwa Town", LGID: 147},
		{ID: 726, Name: "Gubio", LGID: 147},
		{ID: 727, Name: "Mafa", LGID: 147},
		{ID: 728, Name: "Bama", LGID: 147},

		// Gubio (LGID: 148)
		{ID: 729, Name: "Gubio", LGID: 148},
		{ID: 730, Name: "Gubio Town", LGID: 148},
		{ID: 731, Name: "Gamboru", LGID: 148},
		{ID: 732, Name: "Monguno", LGID: 148},
		{ID: 733, Name: "Bama", LGID: 148},

		// Guzamala (LGID: 149)
		{ID: 734, Name: "Guzamala", LGID: 149},
		{ID: 735, Name: "Guzamala Town", LGID: 149},
		{ID: 736, Name: "Gubio", LGID: 149},
		{ID: 737, Name: "Maiduguri", LGID: 149},
		{ID: 738, Name: "Kala/Balge", LGID: 149},

		// Gwoza (LGID: 150)
		{ID: 739, Name: "Gwoza", LGID: 150},
		{ID: 740, Name: "Gwoza Town", LGID: 150},
		{ID: 741, Name: "Maiduguri", LGID: 150},
		{ID: 742, Name: "Bama", LGID: 150},
		{ID: 743, Name: "Damboa", LGID: 150},

		// Hawul (LGID: 151)
		{ID: 744, Name: "Hawul", LGID: 151},
		{ID: 745, Name: "Hawul Town", LGID: 151},
		{ID: 746, Name: "Biu", LGID: 151},
		{ID: 747, Name: "Bama", LGID: 151},
		{ID: 748, Name: "Maiduguri", LGID: 151},

		// Jere (LGID: 152)
		{ID: 749, Name: "Jere", LGID: 152},
		{ID: 750, Name: "Jere Town", LGID: 152},
		{ID: 751, Name: "Maiduguri", LGID: 152},
		{ID: 752, Name: "Bama", LGID: 152},
		{ID: 753, Name: "Gubio", LGID: 152},

		// Kaga (LGID: 153)
		{ID: 754, Name: "Kaga", LGID: 153},
		{ID: 755, Name: "Kaga Town", LGID: 153},
		{ID: 756, Name: "Biu", LGID: 153},
		{ID: 757, Name: "Gubio", LGID: 153},
		{ID: 758, Name: "Maiduguri", LGID: 153},

		// Kala/Balge (LGID: 154)
		{ID: 759, Name: "Kala/Balge", LGID: 154},
		{ID: 760, Name: "Kala Town", LGID: 154},
		{ID: 761, Name: "Balge", LGID: 154},
		{ID: 762, Name: "Maiduguri", LGID: 154},
		{ID: 763, Name: "Damboa", LGID: 154},

		// Konduga (LGID: 155)
		{ID: 764, Name: "Konduga", LGID: 155},
		{ID: 765, Name: "Konduga Town", LGID: 155},
		{ID: 766, Name: "Gubio", LGID: 155},
		{ID: 767, Name: "Bama", LGID: 155},
		{ID: 768, Name: "Maiduguri", LGID: 155},

		// Kukawa (LGID: 156)
		{ID: 769, Name: "Kukawa", LGID: 156},
		{ID: 770, Name: "Kukawa Town", LGID: 156},
		{ID: 771, Name: "Gubio", LGID: 156},
		{ID: 772, Name: "Maiduguri", LGID: 156},
		{ID: 773, Name: "Dikwa", LGID: 156},

		// Kwaya Kusar (LGID: 157)
		{ID: 774, Name: "Kwaya Kusar", LGID: 157},
		{ID: 775, Name: "Kwaya Kusar Town", LGID: 157},
		{ID: 776, Name: "Biu", LGID: 157},
		{ID: 777, Name: "Maiduguri", LGID: 157},
		{ID: 778, Name: "Gwoza", LGID: 157},

		// Mafa (LGID: 158)
		{ID: 779, Name: "Mafa", LGID: 158},
		{ID: 780, Name: "Mafa Town", LGID: 158},
		{ID: 781, Name: "Dikwa", LGID: 158},
		{ID: 782, Name: "Gubio", LGID: 158},
		{ID: 783, Name: "Bama", LGID: 158},

		// Magumeri (LGID: 159)
		{ID: 784, Name: "Magumeri", LGID: 159},
		{ID: 785, Name: "Magumeri Town", LGID: 159},
		{ID: 786, Name: "Gubio", LGID: 159},
		{ID: 787, Name: "Dikwa", LGID: 159},
		{ID: 788, Name: "Maiduguri", LGID: 159},

		// Maiduguri (LGID: 160)
		{ID: 789, Name: "Maiduguri", LGID: 160},
		{ID: 790, Name: "Maiduguri Town", LGID: 160},
		{ID: 791, Name: "Jere", LGID: 160},
		{ID: 792, Name: "Biu", LGID: 160},
		{ID: 793, Name: "Gubio", LGID: 160},

		// Marte (LGID: 161)
		{ID: 794, Name: "Marte", LGID: 161},
		{ID: 795, Name: "Marte Town", LGID: 161},
		{ID: 796, Name: "Maiduguri", LGID: 161},
		{ID: 797, Name: "Gubio", LGID: 161},
		{ID: 798, Name: "Dikwa", LGID: 161},

		// Mobbar (LGID: 162)
		{ID: 799, Name: "Mobbar", LGID: 162},
		{ID: 800, Name: "Mobbar Town", LGID: 162},
		{ID: 801, Name: "Bama", LGID: 162},
		{ID: 802, Name: "Gubio", LGID: 162},
		{ID: 803, Name: "Dikwa", LGID: 162},

		// Monguno (LGID: 163)
		{ID: 804, Name: "Monguno", LGID: 163},
		{ID: 805, Name: "Monguno Town", LGID: 163},
		{ID: 806, Name: "Maiduguri", LGID: 163},
		{ID: 807, Name: "Gubio", LGID: 163},
		{ID: 808, Name: "Dikwa", LGID: 163},

		// Ngala (LGID: 164)
		{ID: 809, Name: "Ngala", LGID: 164},
		{ID: 810, Name: "Ngala Town", LGID: 164},
		{ID: 811, Name: "Maiduguri", LGID: 164},
		{ID: 812, Name: "Gubio", LGID: 164},
		{ID: 813, Name: "Bama", LGID: 164},

		// Nganzai (LGID: 165)
		{ID: 814, Name: "Nganzai", LGID: 165},
		{ID: 815, Name: "Nganzai Town", LGID: 165},
		{ID: 816, Name: "Maiduguri", LGID: 165},
		{ID: 817, Name: "Gubio", LGID: 165},
		{ID: 818, Name: "Bama", LGID: 165},

		// Shani (LGID: 166)
		{ID: 819, Name: "Shani", LGID: 166},
		{ID: 820, Name: "Shani Town", LGID: 166},
		{ID: 821, Name: "Maiduguri", LGID: 166},
		{ID: 822, Name: "Bama", LGID: 166},
		{ID: 823, Name: "Gubio", LGID: 166},

		// Abadam (LGID: 140)
		{ID: 689, Name: "Abadam", LGID: 140},
		{ID: 690, Name: "Kari", LGID: 140},
		{ID: 691, Name: "Ngala", LGID: 140},
		{ID: 692, Name: "Dikwa", LGID: 140},
		{ID: 693, Name: "Gubio", LGID: 140},

		// Askira/Uba (LGID: 141)
		{ID: 694, Name: "Askira/Uba", LGID: 141},
		{ID: 695, Name: "Uba", LGID: 141},
		{ID: 696, Name: "Askira", LGID: 141},
		{ID: 697, Name: "Gulani", LGID: 141},
		{ID: 698, Name: "Maiduguri", LGID: 141},

		// Bama (LGID: 142)
		{ID: 699, Name: "Bama", LGID: 142},
		{ID: 700, Name: "Bama Town", LGID: 142},
		{ID: 701, Name: "Kwaya Kusar", LGID: 142},
		{ID: 702, Name: "Gwoza", LGID: 142},
		{ID: 703, Name: "Dikwa", LGID: 142},

		// Bayo (LGID: 143)
		{ID: 704, Name: "Bayo", LGID: 143},
		{ID: 705, Name: "Bayo Town", LGID: 143},
		{ID: 706, Name: "Wuyo", LGID: 143},
		{ID: 707, Name: "Kukawa", LGID: 143},
		{ID: 708, Name: "Mafa", LGID: 143},

		// Biu (LGID: 144)
		{ID: 709, Name: "Biu", LGID: 144},
		{ID: 710, Name: "Biu Town", LGID: 144},
		{ID: 711, Name: "Biu North", LGID: 144},
		{ID: 712, Name: "Biu South", LGID: 144},
		{ID: 713, Name: "Gubio", LGID: 144},

		// Chibok (LGID: 145)
		{ID: 714, Name: "Chibok", LGID: 145},
		{ID: 715, Name: "Chibok Town", LGID: 145},
		{ID: 716, Name: "Gwoza", LGID: 145},
		{ID: 717, Name: "Damboa", LGID: 145},
		{ID: 718, Name: "Maiduguri", LGID: 145},

		// Damboa (LGID: 146)
		{ID: 719, Name: "Damboa", LGID: 146},
		{ID: 720, Name: "Damboa Town", LGID: 146},
		{ID: 721, Name: "Gwoza", LGID: 146},
		{ID: 722, Name: "Maiduguri", LGID: 146},
		{ID: 723, Name: "Chibok", LGID: 146},

		// Dikwa (LGID: 147)
		{ID: 724, Name: "Dikwa", LGID: 147},
		{ID: 725, Name: "Dikwa Town", LGID: 147},
		{ID: 726, Name: "Gubio", LGID: 147},
		{ID: 727, Name: "Mafa", LGID: 147},
		{ID: 728, Name: "Bama", LGID: 147},

		// Gubio (LGID: 148)
		{ID: 729, Name: "Gubio", LGID: 148},
		{ID: 730, Name: "Gubio Town", LGID: 148},
		{ID: 731, Name: "Gamboru", LGID: 148},
		{ID: 732, Name: "Monguno", LGID: 148},
		{ID: 733, Name: "Bama", LGID: 148},

		// Guzamala (LGID: 149)
		{ID: 734, Name: "Guzamala", LGID: 149},
		{ID: 735, Name: "Guzamala Town", LGID: 149},
		{ID: 736, Name: "Gubio", LGID: 149},
		{ID: 737, Name: "Maiduguri", LGID: 149},
		{ID: 738, Name: "Kala/Balge", LGID: 149},

		// Gwoza (LGID: 150)
		{ID: 739, Name: "Gwoza", LGID: 150},
		{ID: 740, Name: "Gwoza Town", LGID: 150},
		{ID: 741, Name: "Maiduguri", LGID: 150},
		{ID: 742, Name: "Bama", LGID: 150},
		{ID: 743, Name: "Damboa", LGID: 150},

		// Hawul (LGID: 151)
		{ID: 744, Name: "Hawul", LGID: 151},
		{ID: 745, Name: "Hawul Town", LGID: 151},
		{ID: 746, Name: "Biu", LGID: 151},
		{ID: 747, Name: "Bama", LGID: 151},
		{ID: 748, Name: "Maiduguri", LGID: 151},

		// Jere (LGID: 152)
		{ID: 749, Name: "Jere", LGID: 152},
		{ID: 750, Name: "Jere Town", LGID: 152},
		{ID: 751, Name: "Maiduguri", LGID: 152},
		{ID: 752, Name: "Bama", LGID: 152},
		{ID: 753, Name: "Gubio", LGID: 152},

		// Kaga (LGID: 153)
		{ID: 754, Name: "Kaga", LGID: 153},
		{ID: 755, Name: "Kaga Town", LGID: 153},
		{ID: 756, Name: "Biu", LGID: 153},
		{ID: 757, Name: "Gubio", LGID: 153},
		{ID: 758, Name: "Maiduguri", LGID: 153},

		// Kala/Balge (LGID: 154)
		{ID: 759, Name: "Kala/Balge", LGID: 154},
		{ID: 760, Name: "Kala Town", LGID: 154},
		{ID: 761, Name: "Balge", LGID: 154},
		{ID: 762, Name: "Maiduguri", LGID: 154},
		{ID: 763, Name: "Damboa", LGID: 154},

		// Konduga (LGID: 155)
		{ID: 764, Name: "Konduga", LGID: 155},
		{ID: 765, Name: "Konduga Town", LGID: 155},
		{ID: 766, Name: "Gubio", LGID: 155},
		{ID: 767, Name: "Bama", LGID: 155},
		{ID: 768, Name: "Maiduguri", LGID: 155},

		// Kukawa (LGID: 156)
		{ID: 769, Name: "Kukawa", LGID: 156},
		{ID: 770, Name: "Kukawa Town", LGID: 156},
		{ID: 771, Name: "Gubio", LGID: 156},
		{ID: 772, Name: "Maiduguri", LGID: 156},
		{ID: 773, Name: "Dikwa", LGID: 156},

		// Kwaya Kusar (LGID: 157)
		{ID: 774, Name: "Kwaya Kusar", LGID: 157},
		{ID: 775, Name: "Kwaya Kusar Town", LGID: 157},
		{ID: 776, Name: "Biu", LGID: 157},
		{ID: 777, Name: "Maiduguri", LGID: 157},
		{ID: 778, Name: "Gwoza", LGID: 157},

		// Mafa (LGID: 158)
		{ID: 779, Name: "Mafa", LGID: 158},
		{ID: 780, Name: "Mafa Town", LGID: 158},
		{ID: 781, Name: "Dikwa", LGID: 158},
		{ID: 782, Name: "Gubio", LGID: 158},
		{ID: 783, Name: "Bama", LGID: 158},

		// Magumeri (LGID: 159)
		{ID: 784, Name: "Magumeri", LGID: 159},
		{ID: 785, Name: "Magumeri Town", LGID: 159},
		{ID: 786, Name: "Gubio", LGID: 159},
		{ID: 787, Name: "Dikwa", LGID: 159},
		{ID: 788, Name: "Maiduguri", LGID: 159},

		// Maiduguri (LGID: 160)
		{ID: 789, Name: "Maiduguri", LGID: 160},
		{ID: 790, Name: "Maiduguri Town", LGID: 160},
		{ID: 791, Name: "Jere", LGID: 160},
		{ID: 792, Name: "Biu", LGID: 160},
		{ID: 793, Name: "Gubio", LGID: 160},

		// Marte (LGID: 161)
		{ID: 794, Name: "Marte", LGID: 161},
		{ID: 795, Name: "Marte Town", LGID: 161},
		{ID: 796, Name: "Maiduguri", LGID: 161},
		{ID: 797, Name: "Gubio", LGID: 161},
		{ID: 798, Name: "Dikwa", LGID: 161},

		// Mobbar (LGID: 162)
		{ID: 799, Name: "Mobbar", LGID: 162},
		{ID: 800, Name: "Mobbar Town", LGID: 162},
		{ID: 801, Name: "Bama", LGID: 162},
		{ID: 802, Name: "Gubio", LGID: 162},
		{ID: 803, Name: "Dikwa", LGID: 162},

		// Monguno (LGID: 163)
		{ID: 804, Name: "Monguno", LGID: 163},
		{ID: 805, Name: "Monguno Town", LGID: 163},
		{ID: 806, Name: "Maiduguri", LGID: 163},
		{ID: 807, Name: "Gubio", LGID: 163},
		{ID: 808, Name: "Dikwa", LGID: 163},

		// Ngala (LGID: 164)
		{ID: 809, Name: "Ngala", LGID: 164},
		{ID: 810, Name: "Ngala Town", LGID: 164},
		{ID: 811, Name: "Maiduguri", LGID: 164},
		{ID: 812, Name: "Gubio", LGID: 164},
		{ID: 813, Name: "Bama", LGID: 164},

		// Nganzai (LGID: 165)
		{ID: 814, Name: "Nganzai", LGID: 165},
		{ID: 815, Name: "Nganzai Town", LGID: 165},
		{ID: 816, Name: "Maiduguri", LGID: 165},
		{ID: 817, Name: "Gubio", LGID: 165},
		{ID: 818, Name: "Bama", LGID: 165},

		// Shani (LGID: 166)
		{ID: 819, Name: "Shani", LGID: 166},
		{ID: 820, Name: "Shani Town", LGID: 166},
		{ID: 821, Name: "Maiduguri", LGID: 166},
		{ID: 822, Name: "Bama", LGID: 166},
		{ID: 823, Name: "Gubio", LGID: 166},

		// Aniocha North (LGID: 185)
		{ID: 914, Name: "Aniocha North", LGID: 185},
		{ID: 915, Name: "Aniocha North Town", LGID: 185},
		{ID: 916, Name: "Aniocha South", LGID: 185},
		{ID: 917, Name: "Ethiope East", LGID: 185},
		{ID: 918, Name: "Ika North East", LGID: 185},

		// Aniocha South (LGID: 186)
		{ID: 919, Name: "Aniocha South", LGID: 186},
		{ID: 920, Name: "Aniocha South Town", LGID: 186},
		{ID: 921, Name: "Ethiope West", LGID: 186},
		{ID: 922, Name: "Ika South", LGID: 186},
		{ID: 923, Name: "Isoko North", LGID: 186},

		// Bomadi (LGID: 187)
		{ID: 924, Name: "Bomadi", LGID: 187},
		{ID: 925, Name: "Bomadi Town", LGID: 187},
		{ID: 926, Name: "Burutu", LGID: 187},
		{ID: 927, Name: "Okpe", LGID: 187},
		{ID: 928, Name: "Sapele", LGID: 187},

		// Burutu (LGID: 188)
		{ID: 929, Name: "Burutu", LGID: 188},
		{ID: 930, Name: "Burutu Town", LGID: 188},
		{ID: 931, Name: "Ethiope East", LGID: 188},
		{ID: 932, Name: "Isoko South", LGID: 188},
		{ID: 933, Name: "Udu", LGID: 188},

		// Ethiope East (LGID: 189)
		{ID: 934, Name: "Ethiope East", LGID: 189},
		{ID: 935, Name: "Ethiope East Town", LGID: 189},
		{ID: 936, Name: "Ika North East", LGID: 189},
		{ID: 937, Name: "Isoko North", LGID: 189},
		{ID: 938, Name: "Ukwuani", LGID: 189},

		// Ethiope West (LGID: 190)
		{ID: 939, Name: "Ethiope West", LGID: 190},
		{ID: 940, Name: "Ethiope West Town", LGID: 190},
		{ID: 941, Name: "Ika South", LGID: 190},
		{ID: 942, Name: "Isoko South", LGID: 190},
		{ID: 943, Name: "Uvwie", LGID: 190},

		// Ika North East (LGID: 191)
		{ID: 944, Name: "Ika North East", LGID: 191},
		{ID: 945, Name: "Ika North East Town", LGID: 191},
		{ID: 946, Name: "Ika South", LGID: 191},
		{ID: 947, Name: "Isoko North", LGID: 191},
		{ID: 948, Name: "Ukwuani", LGID: 191},

		// Ika South (LGID: 192)
		{ID: 949, Name: "Ika South", LGID: 192},
		{ID: 950, Name: "Ika South Town", LGID: 192},
		{ID: 951, Name: "Isoko North", LGID: 192},
		{ID: 952, Name: "Isoko South", LGID: 192},
		{ID: 953, Name: "Uvwie", LGID: 192},

		// Isoko North (LGID: 193)
		{ID: 954, Name: "Isoko North", LGID: 193},
		{ID: 955, Name: "Isoko North Town", LGID: 193},
		{ID: 956, Name: "Isoko South", LGID: 193},
		{ID: 957, Name: "Ukwuani", LGID: 193},
		{ID: 958, Name: "Uvwie", LGID: 193},

		// Isoko South (LGID: 194)
		{ID: 959, Name: "Isoko South", LGID: 194},
		{ID: 960, Name: "Isoko South Town", LGID: 194},
		{ID: 961, Name: "Udu", LGID: 194},
		{ID: 962, Name: "Ukwuani", LGID: 194},
		{ID: 963, Name: "Warri South", LGID: 194},

		// Ndokwa East (LGID: 195)
		{ID: 964, Name: "Ndokwa East", LGID: 195},
		{ID: 965, Name: "Ndokwa East Town", LGID: 195},
		{ID: 966, Name: "Ndokwa West", LGID: 195},
		{ID: 967, Name: "Okpe", LGID: 195},
		{ID: 968, Name: "Sapele", LGID: 195},

		// Ndokwa West (LGID: 196)
		{ID: 969, Name: "Ndokwa West", LGID: 196},
		{ID: 970, Name: "Ndokwa West Town", LGID: 196},
		{ID: 971, Name: "Okpe", LGID: 196},
		{ID: 972, Name: "Sapele", LGID: 196},
		{ID: 973, Name: "Warri South West", LGID: 196},

		// Okpe (LGID: 197)
		{ID: 974, Name: "Okpe", LGID: 197},
		{ID: 975, Name: "Okpe Town", LGID: 197},
		{ID: 976, Name: "Patani", LGID: 197},
		{ID: 977, Name: "Udu", LGID: 197},
		{ID: 978, Name: "Warri South", LGID: 197},

		// Oshimili North (LGID: 198)
		{ID: 979, Name: "Oshimili North", LGID: 198},
		{ID: 980, Name: "Oshimili North Town", LGID: 198},
		{ID: 981, Name: "Oshimili South", LGID: 198},
		{ID: 982, Name: "Sapele", LGID: 198},
		{ID: 983, Name: "Warri North", LGID: 198},

		// Oshimili South (LGID: 199)
		{ID: 984, Name: "Oshimili South", LGID: 199},
		{ID: 985, Name: "Oshimili South Town", LGID: 199},
		{ID: 986, Name: "Patani", LGID: 199},
		{ID: 987, Name: "Sapele", LGID: 199},
		{ID: 988, Name: "Warri South West", LGID: 199},

		// Patani (LGID: 200)
		{ID: 989, Name: "Patani", LGID: 200},
		{ID: 990, Name: "Patani Town", LGID: 200},
		{ID: 991, Name: "Sapele", LGID: 200},
		{ID: 992, Name: "Udu", LGID: 200},
		{ID: 993, Name: "Warri South West", LGID: 200},

		// Sapele (LGID: 201)
		{ID: 994, Name: "Sapele", LGID: 201},
		{ID: 995, Name: "Sapele Town", LGID: 201},
		{ID: 996, Name: "Udu", LGID: 201},
		{ID: 997, Name: "Ukwuani", LGID: 201},
		{ID: 998, Name: "Warri North", LGID: 201},

		// Udu (LGID: 202)
		{ID: 999, Name: "Udu", LGID: 202},
		{ID: 1000, Name: "Udu Town", LGID: 202},
		{ID: 1001, Name: "Ukwuani", LGID: 202},
		{ID: 1002, Name: "Uvwie", LGID: 202},
		{ID: 1003, Name: "Warri North", LGID: 202},

		// Ughelli North (LGID: 203)
		{ID: 1004, Name: "Ughelli North", LGID: 203},
		{ID: 1005, Name: "Ughelli North Town", LGID: 203},
		{ID: 1006, Name: "Ughelli South", LGID: 203},
		{ID: 1007, Name: "Udu", LGID: 203},
		{ID: 1008, Name: "Warri North", LGID: 203},

		// Ughelli South (LGID: 204)
		{ID: 1009, Name: "Ughelli South", LGID: 204},
		{ID: 1010, Name: "Ughelli South Town", LGID: 204},
		{ID: 1011, Name: "Udu", LGID: 204},
		{ID: 1012, Name: "Uvwie", LGID: 204},
		{ID: 1013, Name: "Warri South", LGID: 204},

		// Ukwuani (LGID: 205)
		{ID: 1014, Name: "Ukwuani", LGID: 205},
		{ID: 1015, Name: "Ukwuani Town", LGID: 205},
		{ID: 1016, Name: "Uvwie", LGID: 205},
		{ID: 1017, Name: "Warri North", LGID: 205},
		{ID: 1018, Name: "Warri South West", LGID: 205},

		// Uvwie (LGID: 206)
		{ID: 1019, Name: "Uvwie", LGID: 206},
		{ID: 1020, Name: "Uvwie Town", LGID: 206},
		{ID: 1021, Name: "Warri South", LGID: 206},
		{ID: 1022, Name: "Warri North", LGID: 206},
		{ID: 1023, Name: "Warri South West", LGID: 206},

		// Warri North (LGID: 207)
		{ID: 1024, Name: "Warri North", LGID: 207},
		{ID: 1025, Name: "Warri North Town", LGID: 207},
		{ID: 1026, Name: "Warri South", LGID: 207},
		{ID: 1027, Name: "Warri South West", LGID: 207},
		{ID: 1028, Name: "Sapele", LGID: 207},

		// Warri South (LGID: 208)
		{ID: 1029, Name: "Warri South", LGID: 208},
		{ID: 1030, Name: "Warri South Town", LGID: 208},
		{ID: 1031, Name: "Warri North", LGID: 208},
		{ID: 1032, Name: "Warri South West", LGID: 208},
		{ID: 1033, Name: "Sapele", LGID: 208},

		// Warri South West (LGID: 209)
		{ID: 1034, Name: "Warri South West", LGID: 209},
		{ID: 1035, Name: "Warri South West Town", LGID: 209},
		{ID: 1036, Name: "Warri North", LGID: 209},
		{ID: 1037, Name: "Warri South", LGID: 209},
		{ID: 1038, Name: "Sapele", LGID: 209},

		// Abakaliki (LGID: 210)
		{ID: 1039, Name: "Abakaliki", LGID: 210},
		{ID: 1040, Name: "Abakaliki Town", LGID: 210},
		{ID: 1041, Name: "Ebonyi", LGID: 210},
		{ID: 1042, Name: "Ishielu", LGID: 210},
		{ID: 1043, Name: "Onicha", LGID: 210},

		// Afikpo North (LGID: 211)
		{ID: 1044, Name: "Afikpo North", LGID: 211},
		{ID: 1045, Name: "Afikpo North Town", LGID: 211},
		{ID: 1046, Name: "Afikpo South", LGID: 211},
		{ID: 1047, Name: "Ivo", LGID: 211},
		{ID: 1048, Name: "Izzi", LGID: 211},

		// Afikpo South (LGID: 212)
		{ID: 1049, Name: "Afikpo South", LGID: 212},
		{ID: 1050, Name: "Afikpo South Town", LGID: 212},
		{ID: 1051, Name: "Ikwo", LGID: 212},
		{ID: 1052, Name: "Ivo", LGID: 212},
		{ID: 1053, Name: "Ohaozara", LGID: 212},

		// Ebonyi (LGID: 213)
		{ID: 1054, Name: "Ebonyi", LGID: 213},
		{ID: 1055, Name: "Ebonyi Town", LGID: 213},
		{ID: 1056, Name: "Ikwo", LGID: 213},
		{ID: 1057, Name: "Ohaozara", LGID: 213},
		{ID: 1058, Name: "Onicha", LGID: 213},

		// Ezza North (LGID: 214)
		{ID: 1059, Name: "Ezza North", LGID: 214},
		{ID: 1060, Name: "Ezza North Town", LGID: 214},
		{ID: 1061, Name: "Ezza South", LGID: 214},
		{ID: 1062, Name: "Ikwo", LGID: 214},
		{ID: 1063, Name: "Onicha", LGID: 214},

		// Ezza South (LGID: 215)
		{ID: 1064, Name: "Ezza South", LGID: 215},
		{ID: 1065, Name: "Ezza South Town", LGID: 215},
		{ID: 1066, Name: "Ikwo", LGID: 215},
		{ID: 1067, Name: "Ivo", LGID: 215},
		{ID: 1068, Name: "Onicha", LGID: 215},

		// Ikwo (LGID: 216)
		{ID: 1069, Name: "Ikwo", LGID: 216},
		{ID: 1070, Name: "Ikwo Town", LGID: 216},
		{ID: 1071, Name: "Ivo", LGID: 216},
		{ID: 1072, Name: "Izzi", LGID: 216},
		{ID: 1073, Name: "Ohaozara", LGID: 216},

		// Ishielu (LGID: 217)
		{ID: 1074, Name: "Ishielu", LGID: 217},
		{ID: 1075, Name: "Ishielu Town", LGID: 217},
		{ID: 1076, Name: "Ivo", LGID: 217},
		{ID: 1077, Name: "Izzi", LGID: 217},
		{ID: 1078, Name: "Onicha", LGID: 217},

		// Ivo (LGID: 218)
		{ID: 1079, Name: "Ivo", LGID: 218},
		{ID: 1080, Name: "Ivo Town", LGID: 218},
		{ID: 1081, Name: "Izzi", LGID: 218},
		{ID: 1082, Name: "Ohaozara", LGID: 218},
		{ID: 1083, Name: "Onicha", LGID: 218},

		// Izzi (LGID: 219)
		{ID: 1084, Name: "Izzi", LGID: 219},
		{ID: 1085, Name: "Izzi Town", LGID: 219},
		{ID: 1086, Name: "Ohaozara", LGID: 219},
		{ID: 1087, Name: "Onicha", LGID: 219},
		{ID: 1088, Name: "Ezza South", LGID: 219},

		// Ohaozara (LGID: 220)
		{ID: 1089, Name: "Ohaozara", LGID: 220},
		{ID: 1090, Name: "Ohaozara Town", LGID: 220},
		{ID: 1091, Name: "Onicha", LGID: 220},
		{ID: 1092, Name: "Ezza North", LGID: 220},
		{ID: 1093, Name: "Ivo", LGID: 220},

		// Ohaukwu (LGID: 221)
		{ID: 1094, Name: "Ohaukwu", LGID: 221},
		{ID: 1095, Name: "Ohaukwu Town", LGID: 221},
		{ID: 1096, Name: "Onicha", LGID: 221},
		{ID: 1097, Name: "Ezza South", LGID: 221},
		{ID: 1098, Name: "Ikwo", LGID: 221},

		// Onicha (LGID: 222)
		{ID: 1099, Name: "Onicha", LGID: 222},
		{ID: 1100, Name: "Onicha Town", LGID: 222},
		{ID: 1101, Name: "Izzi", LGID: 222},
		{ID: 1102, Name: "Ivo", LGID: 222},
		{ID: 1103, Name: "Ohaozara", LGID: 222},

		// Akoko Edo (LGID: 223)
		{ID: 1104, Name: "Akoko Edo", LGID: 223},
		{ID: 1105, Name: "Akoko Edo Town", LGID: 223},
		{ID: 1106, Name: "Boki", LGID: 223},
		{ID: 1107, Name: "Esan North-East", LGID: 223},
		{ID: 1108, Name: "Etsako West", LGID: 223},

		// Egor (LGID: 224)
		{ID: 1109, Name: "Egor", LGID: 224},
		{ID: 1110, Name: "Egor Town", LGID: 224},
		{ID: 1111, Name: "Ikpoba-Okha", LGID: 224},
		{ID: 1112, Name: "Oredo", LGID: 224},
		{ID: 1113, Name: "Ovia North-East", LGID: 224},

		// Esan Central (LGID: 225)
		{ID: 1114, Name: "Esan Central", LGID: 225},
		{ID: 1115, Name: "Esan Central Town", LGID: 225},
		{ID: 1116, Name: "Esan North-East", LGID: 225},
		{ID: 1117, Name: "Esan South-East", LGID: 225},
		{ID: 1118, Name: "Esan West", LGID: 225},

		// Esan North-East (LGID: 226)
		{ID: 1119, Name: "Esan North-East", LGID: 226},
		{ID: 1120, Name: "Esan North-East Town", LGID: 226},
		{ID: 1121, Name: "Esan Central", LGID: 226},
		{ID: 1122, Name: "Esan South-East", LGID: 226},
		{ID: 1123, Name: "Esan West", LGID: 226},

		// Esan South-East (LGID: 227)
		{ID: 1124, Name: "Esan South-East", LGID: 227},
		{ID: 1125, Name: "Esan South-East Town", LGID: 227},
		{ID: 1126, Name: "Esan Central", LGID: 227},
		{ID: 1127, Name: "Esan North-East", LGID: 227},
		{ID: 1128, Name: "Esan West", LGID: 227},

		// Esan West (LGID: 228)
		{ID: 1129, Name: "Esan West", LGID: 228},
		{ID: 1130, Name: "Esan West Town", LGID: 228},
		{ID: 1131, Name: "Esan Central", LGID: 228},
		{ID: 1132, Name: "Esan North-East", LGID: 228},
		{ID: 1133, Name: "Esan South-East", LGID: 228},

		// Etsako Central (LGID: 229)
		{ID: 1134, Name: "Etsako Central", LGID: 229},
		{ID: 1135, Name: "Etsako Central Town", LGID: 229},
		{ID: 1136, Name: "Etsako East", LGID: 229},
		{ID: 1137, Name: "Etsako West", LGID: 229},
		{ID: 1138, Name: "Owan East", LGID: 229},

		// Etsako East (LGID: 230)
		{ID: 1139, Name: "Etsako East", LGID: 230},
		{ID: 1140, Name: "Etsako East Town", LGID: 230},
		{ID: 1141, Name: "Etsako Central", LGID: 230},
		{ID: 1142, Name: "Etsako West", LGID: 230},
		{ID: 1143, Name: "Owan East", LGID: 230},

		// Etsako West (LGID: 231)
		{ID: 1144, Name: "Etsako West", LGID: 231},
		{ID: 1145, Name: "Etsako West Town", LGID: 231},
		{ID: 1146, Name: "Etsako Central", LGID: 231},
		{ID: 1147, Name: "Etsako East", LGID: 231},
		{ID: 1148, Name: "Owan East", LGID: 231},

		// Igueben (LGID: 232)
		{ID: 1149, Name: "Igueben", LGID: 232},
		{ID: 1150, Name: "Igueben Town", LGID: 232},
		{ID: 1151, Name: "Igueben East", LGID: 232},
		{ID: 1152, Name: "Owan West", LGID: 232},
		{ID: 1153, Name: "Uhunmwonde", LGID: 232},

		// Ikpoba-Okha (LGID: 233)
		{ID: 1154, Name: "Ikpoba-Okha", LGID: 233},
		{ID: 1155, Name: "Ikpoba-Okha Town", LGID: 233},
		{ID: 1156, Name: "Oredo", LGID: 233},
		{ID: 1157, Name: "Ovia North-East", LGID: 233},
		{ID: 1158, Name: "Uhunmwonde", LGID: 233},

		// Oredo (LGID: 234)
		{ID: 1159, Name: "Oredo", LGID: 234},
		{ID: 1160, Name: "Oredo Town", LGID: 234},
		{ID: 1161, Name: "Ikpoba-Okha", LGID: 234},
		{ID: 1162, Name: "Ovia North-East", LGID: 234},
		{ID: 1163, Name: "Uhunmwonde", LGID: 234},

		// Orhionmwon (LGID: 235)
		{ID: 1164, Name: "Orhionmwon", LGID: 235},
		{ID: 1165, Name: "Orhionmwon Town", LGID: 235},
		{ID: 1166, Name: "Etsako East", LGID: 235},
		{ID: 1167, Name: "Etsako West", LGID: 235},
		{ID: 1168, Name: "Owan West", LGID: 235},

		// Ovia North-East (LGID: 236)
		{ID: 1169, Name: "Ovia North-East", LGID: 236},
		{ID: 1170, Name: "Ovia North-East Town", LGID: 236},
		{ID: 1171, Name: "Ikpoba-Okha", LGID: 236},
		{ID: 1172, Name: "Oredo", LGID: 236},
		{ID: 1173, Name: "Uhunmwonde", LGID: 236},

		// Ovia South-West (LGID: 237)
		{ID: 1174, Name: "Ovia South-West", LGID: 237},
		{ID: 1175, Name: "Ovia South-West Town", LGID: 237},
		{ID: 1176, Name: "Etsako Central", LGID: 237},
		{ID: 1177, Name: "Oredo", LGID: 237},
		{ID: 1178, Name: "Uhunmwonde", LGID: 237},

		// Owan East (LGID: 238)
		{ID: 1179, Name: "Owan East", LGID: 238},
		{ID: 1180, Name: "Owan East Town", LGID: 238},
		{ID: 1181, Name: "Etsako Central", LGID: 238},
		{ID: 1182, Name: "Etsako West", LGID: 238},
		{ID: 1183, Name: "Uhunmwonde", LGID: 238},

		// Owan West (LGID: 239)
		{ID: 1184, Name: "Owan West", LGID: 239},
		{ID: 1185, Name: "Owan West Town", LGID: 239},
		{ID: 1186, Name: "Etsako Central", LGID: 239},
		{ID: 1187, Name: "Etsako East", LGID: 239},
		{ID: 1188, Name: "Uhunmwonde", LGID: 239},

		// Uhunmwonde (LGID: 240)
		{ID: 1189, Name: "Uhunmwonde", LGID: 240},
		{ID: 1190, Name: "Uhunmwonde Town", LGID: 240},
		{ID: 1191, Name: "Ovia South-West", LGID: 240},
		{ID: 1192, Name: "Owan East", LGID: 240},
		{ID: 1193, Name: "Owan West", LGID: 240},

		// Ado Ekiti (LGID: 241)
		{ID: 1194, Name: "Ado Ekiti", LGID: 241},
		{ID: 1195, Name: "Ado Town", LGID: 241},
		{ID: 1196, Name: "Ado Central", LGID: 241},
		{ID: 1197, Name: "Ado West", LGID: 241},
		{ID: 1198, Name: "Ado North", LGID: 241},

		// Efon (LGID: 242)
		{ID: 1199, Name: "Efon", LGID: 242},
		{ID: 1200, Name: "Efon Town", LGID: 242},
		{ID: 1201, Name: "Efon South", LGID: 242},
		{ID: 1202, Name: "Efon Central", LGID: 242},
		{ID: 1203, Name: "Efon West", LGID: 242},

		// Ekiti East (LGID: 243)
		{ID: 1204, Name: "Ekiti East", LGID: 243},
		{ID: 1205, Name: "Ekiti East Town", LGID: 243},
		{ID: 1206, Name: "Ekiti East North", LGID: 243},
		{ID: 1207, Name: "Ekiti East Central", LGID: 243},
		{ID: 1208, Name: "Ekiti East West", LGID: 243},

		// Ekiti South-West (LGID: 244)
		{ID: 1209, Name: "Ekiti South-West", LGID: 244},
		{ID: 1210, Name: "Ekiti South-West Town", LGID: 244},
		{ID: 1211, Name: "Ekiti South-West East", LGID: 244},
		{ID: 1212, Name: "Ekiti South-West Central", LGID: 244},
		{ID: 1213, Name: "Ekiti South-West West", LGID: 244},

		// Ekiti West (LGID: 245)
		{ID: 1214, Name: "Ekiti West", LGID: 245},
		{ID: 1215, Name: "Ekiti West Town", LGID: 245},
		{ID: 1216, Name: "Ekiti West North", LGID: 245},
		{ID: 1217, Name: "Ekiti West Central", LGID: 245},
		{ID: 1218, Name: "Ekiti West South", LGID: 245},

		// Emure (LGID: 246)
		{ID: 1219, Name: "Emure", LGID: 246},
		{ID: 1220, Name: "Emure Town", LGID: 246},
		{ID: 1221, Name: "Emure North", LGID: 246},
		{ID: 1222, Name: "Emure Central", LGID: 246},
		{ID: 1223, Name: "Emure South", LGID: 246},

		// Gbonyin (LGID: 247)
		{ID: 1224, Name: "Gbonyin", LGID: 247},
		{ID: 1225, Name: "Gbonyin Town", LGID: 247},
		{ID: 1226, Name: "Gbonyin North", LGID: 247},
		{ID: 1227, Name: "Gbonyin Central", LGID: 247},
		{ID: 1228, Name: "Gbonyin South", LGID: 247},

		// Ido Osi (LGID: 248)
		{ID: 1229, Name: "Ido Osi", LGID: 248},
		{ID: 1230, Name: "Ido Osi Town", LGID: 248},
		{ID: 1231, Name: "Ido Osi North", LGID: 248},
		{ID: 1232, Name: "Ido Osi Central", LGID: 248},
		{ID: 1233, Name: "Ido Osi South", LGID: 248},

		// Ijero (LGID: 249)
		{ID: 1234, Name: "Ijero", LGID: 249},
		{ID: 1235, Name: "Ijero Town", LGID: 249},
		{ID: 1236, Name: "Ijero North", LGID: 249},
		{ID: 1237, Name: "Ijero Central", LGID: 249},
		{ID: 1238, Name: "Ijero South", LGID: 249},

		// Ikere (LGID: 250)
		{ID: 1239, Name: "Ikere", LGID: 250},
		{ID: 1240, Name: "Ikere Town", LGID: 250},
		{ID: 1241, Name: "Ikere North", LGID: 250},
		{ID: 1242, Name: "Ikere Central", LGID: 250},
		{ID: 1243, Name: "Ikere South", LGID: 250},

		// Ikole (LGID: 251)
		{ID: 1244, Name: "Ikole", LGID: 251},
		{ID: 1245, Name: "Ikole Town", LGID: 251},
		{ID: 1246, Name: "Ikole North", LGID: 251},
		{ID: 1247, Name: "Ikole Central", LGID: 251},
		{ID: 1248, Name: "Ikole South", LGID: 251},

		// Ilejemeje (LGID: 252)
		{ID: 1249, Name: "Ilejemeje", LGID: 252},
		{ID: 1250, Name: "Ilejemeje Town", LGID: 252},
		{ID: 1251, Name: "Ilejemeje North", LGID: 252},
		{ID: 1252, Name: "Ilejemeje Central", LGID: 252},
		{ID: 1253, Name: "Ilejemeje South", LGID: 252},

		// Irepodun/Ifelodun (LGID: 253)
		{ID: 1254, Name: "Irepodun/Ifelodun", LGID: 253},
		{ID: 1255, Name: "Irepodun/Ifelodun Town", LGID: 253},
		{ID: 1256, Name: "Irepodun/Ifelodun North", LGID: 253},
		{ID: 1257, Name: "Irepodun/Ifelodun Central", LGID: 253},
		{ID: 1258, Name: "Irepodun/Ifelodun South", LGID: 253},

		// Ise/Orun (LGID: 254)
		{ID: 1259, Name: "Ise/Orun", LGID: 254},
		{ID: 1260, Name: "Ise/Orun Town", LGID: 254},
		{ID: 1261, Name: "Ise/Orun North", LGID: 254},
		{ID: 1262, Name: "Ise/Orun Central", LGID: 254},
		{ID: 1263, Name: "Ise/Orun South", LGID: 254},

		// Moba (LGID: 255)
		{ID: 1264, Name: "Moba", LGID: 255},
		{ID: 1265, Name: "Moba Town", LGID: 255},
		{ID: 1266, Name: "Moba North", LGID: 255},
		{ID: 1267, Name: "Moba Central", LGID: 255},
		{ID: 1268, Name: "Moba South", LGID: 255},

		// Oye (LGID: 256)
		{ID: 1269, Name: "Oye", LGID: 256},
		{ID: 1270, Name: "Oye Town", LGID: 256},
		{ID: 1271, Name: "Oye North", LGID: 256},
		{ID: 1272, Name: "Oye Central", LGID: 256},
		{ID: 1273, Name: "Oye South", LGID: 256},

		// Aninri
		{ID: 1274, Name: "Ndeaboh", LGID: 257},
		{ID: 1275, Name: "Oha Nkanu", LGID: 257},
		{ID: 1276, Name: "Mpu", LGID: 257},
		{ID: 1277, Name: "Oduma", LGID: 257},
		{ID: 1278, Name: "Nenwe", LGID: 257},

		// Awgu
		{ID: 1279, Name: "Agbogugu", LGID: 258},
		{ID: 1280, Name: "Agbudu", LGID: 258},
		{ID: 1281, Name: "Awgu", LGID: 258},
		{ID: 1282, Name: "Ihe", LGID: 258},
		{ID: 1283, Name: "Mgbidi", LGID: 258},
		{ID: 1284, Name: "Mgbowo", LGID: 258},
		{ID: 1285, Name: "Nkwe", LGID: 258},
		{ID: 1286, Name: "Owelli", LGID: 258},
		{ID: 1287, Name: "Oji", LGID: 258},
		{ID: 1288, Name: "Ugbo", LGID: 258},

		// Enugu East
		{ID: 1289, Name: "Nkwo Nike", LGID: 259},
		{ID: 1290, Name: "Abakpa Nike", LGID: 259},
		{ID: 1291, Name: "Emene", LGID: 259},
		{ID: 1292, Name: "Trans-Ekulu", LGID: 259},
		{ID: 1293, Name: "GRA", LGID: 259},

		// Enugu North
		{ID: 1294, Name: "Asata", LGID: 260},
		{ID: 1295, Name: "Ogui", LGID: 260},
		{ID: 1296, Name: "GRA", LGID: 260},
		{ID: 1297, Name: "New Haven", LGID: 260},
		{ID: 1298, Name: "Independence Layout", LGID: 260},

		// Enugu South
		{ID: 1299, Name: "Achara Layout", LGID: 261},
		{ID: 1300, Name: "Amechi", LGID: 261},
		{ID: 1301, Name: "Ugwuaji", LGID: 261},

		// Ezeagu
		{ID: 1302, Name: "Aguobu-Owa", LGID: 262},
		{ID: 1303, Name: "Ihuokpara", LGID: 262},
		{ID: 1304, Name: "Ihuoha", LGID: 262},
		{ID: 1305, Name: "Olo", LGID: 262},
		{ID: 1306, Name: "Obinofia", LGID: 262},
		{ID: 1307, Name: "Umana Ndiagu", LGID: 262},
		{ID: 1308, Name: "Udi", LGID: 262},

		// Igbo Etiti
		{ID: 1309, Name: "Ogbede", LGID: 263},
		{ID: 1310, Name: "Diogbe", LGID: 263},
		{ID: 1311, Name: "Ikolo", LGID: 263},
		{ID: 1312, Name: "Ukehe", LGID: 263},
		{ID: 1313, Name: "Ikem", LGID: 263},

		// Igbo Eze North
		{ID: 1314, Name: "Enugu-Ezike", LGID: 264},
		{ID: 1315, Name: "Ibagwa", LGID: 264},
		{ID: 1316, Name: "Umuitodo", LGID: 264},
		{ID: 1317, Name: "Ogurugu", LGID: 264},
		{ID: 1318, Name: "Aji", LGID: 264},

		// Igbo Eze South
		{ID: 1319, Name: "Ibagwa-Aka", LGID: 265},
		{ID: 1320, Name: "Ovoko", LGID: 265},
		{ID: 1321, Name: "Uhunowerre", LGID: 265},
		{ID: 1322, Name: "Iheaka", LGID: 265},
		{ID: 1323, Name: "Nkalagu-Obukpa", LGID: 265},

		// Isi Uzo
		{ID: 1324, Name: "Ikem", LGID: 266},
		{ID: 1325, Name: "Neke", LGID: 266},
		{ID: 1326, Name: "Umualor", LGID: 266},
		{ID: 1327, Name: "Eha-Amufu", LGID: 266},
		{ID: 1328, Name: "Mbu", LGID: 266},

		// Nkanu East
		{ID: 1329, Name: "Amankanu", LGID: 267},
		{ID: 1330, Name: "Akpugo", LGID: 267},
		{ID: 1331, Name: "Akpawfu", LGID: 267},
		{ID: 1332, Name: "Ugbawka", LGID: 267},
		{ID: 1333, Name: "Mburubu", LGID: 267},

		// Nkanu West
		{ID: 1334, Name: "Akegbe Ugwu", LGID: 268},
		{ID: 1335, Name: "Obe", LGID: 268},
		{ID: 1336, Name: "Ozalla", LGID: 268},
		{ID: 1337, Name: "Amurri", LGID: 268},
		{ID: 1338, Name: "Amaechi", LGID: 268},

		// Nsukka
		{ID: 1339, Name: "Nsukka", LGID: 269},
		{ID: 1340, Name: "Opi", LGID: 269},
		{ID: 1341, Name: "Ede-Oballa", LGID: 269},
		{ID: 1342, Name: "Ibagwa-Ani", LGID: 269},
		{ID: 1343, Name: "Nru", LGID: 269},

		// Oji River
		{ID: 1344, Name: "Achi", LGID: 270},
		{ID: 1345, Name: "Inyi", LGID: 270},
		{ID: 1346, Name: "Akpugoeze", LGID: 270},
		{ID: 1347, Name: "Oji River", LGID: 270},
		{ID: 1348, Name: "Umuneke", LGID: 270},

		// Udenu
		{ID: 1349, Name: "Obollo-Afor", LGID: 271},
		{ID: 1350, Name: "Orba", LGID: 271},
		{ID: 1351, Name: "Imilike", LGID: 271},
		{ID: 1352, Name: "Ezimo", LGID: 271},
		{ID: 1353, Name: "Ukehe", LGID: 271},

		// Udi
		{ID: 1354, Name: "Amokwe", LGID: 272},
		{ID: 1355, Name: "Udi", LGID: 272},
		{ID: 1356, Name: "Nachi", LGID: 272},
		{ID: 1357, Name: "Ebe", LGID: 272},
		{ID: 1358, Name: "Abia", LGID: 272},

		// Uzo Uwani
		{ID: 1359, Name: "Adani", LGID: 273},
		{ID: 1360, Name: "Umulokpa", LGID: 273},
		{ID: 1361, Name: "Ukpata", LGID: 273},
		{ID: 1362, Name: "Igga", LGID: 273},
		{ID: 1363, Name: "Nkpologwu", LGID: 273},
		{ID: 1364, Name: "Nimbo", LGID: 273},

		// Gombe
		{ID: 1365, Name: "Garko", LGID: 274},
		{ID: 1366, Name: "Kumo", LGID: 274},
		{ID: 1367, Name: "Pindiga", LGID: 274},
		{ID: 1368, Name: "Tumu", LGID: 274},

		{ID: 1369, Name: "Balanga", LGID: 275},
		{ID: 1370, Name: "Gelengu", LGID: 275},
		{ID: 1371, Name: "Bambam", LGID: 275},
		{ID: 1372, Name: "Tula", LGID: 275},

		{ID: 1373, Name: "Billiri", LGID: 276},
		{ID: 1374, Name: "Tal", LGID: 276},
		{ID: 1375, Name: "Kwalamba", LGID: 276},
		{ID: 1376, Name: "Todi", LGID: 276},

		{ID: 1377, Name: "Dukku", LGID: 277},
		{ID: 1378, Name: "Hashidu", LGID: 277},
		{ID: 1379, Name: "Jamal Gombe", LGID: 277},
		{ID: 1380, Name: "Wuro Tale", LGID: 277},

		{ID: 1381, Name: "Bajoga", LGID: 278},
		{ID: 1382, Name: "Ashaka", LGID: 278},
		{ID: 1383, Name: "Filiya", LGID: 278},
		{ID: 1384, Name: "Gundale", LGID: 278},

		{ID: 1385, Name: "Gombe", LGID: 279},
		{ID: 1386, Name: "Kwadon", LGID: 279},
		{ID: 1387, Name: "Tumfure", LGID: 279},
		{ID: 1388, Name: "Pantami", LGID: 279},

		{ID: 1389, Name: "Kaltungo", LGID: 280},
		{ID: 1390, Name: "Awak", LGID: 280},
		{ID: 1391, Name: "Ture", LGID: 280},
		{ID: 1392, Name: "Kamo", LGID: 280},

		{ID: 1393, Name: "Kwami", LGID: 281},
		{ID: 1394, Name: "Mallam Sidi", LGID: 281},
		{ID: 1395, Name: "Jauro Tukur", LGID: 281},
		{ID: 1396, Name: "Bojuwa", LGID: 281},

		{ID: 1397, Name: "Nafada", LGID: 282},
		{ID: 1398, Name: "Langa", LGID: 282},
		{ID: 1399, Name: "Zambuk", LGID: 282},
		{ID: 1400, Name: "Difa", LGID: 282},

		{ID: 1401, Name: "Shongom", LGID: 283},
		{ID: 1402, Name: "Bangoba", LGID: 283},
		{ID: 1403, Name: "Lapan", LGID: 283},
		{ID: 1404, Name: "Gungura", LGID: 283},

		{ID: 1405, Name: "Deba", LGID: 284},
		{ID: 1406, Name: "Yamaltu", LGID: 284},
		{ID: 1407, Name: "Hinna", LGID: 284},
		{ID: 1408, Name: "Kuri", LGID: 284},

		// Imo
		{ID: 1409, Name: "Ahiara", LGID: 285},
		{ID: 1410, Name: "Nguru", LGID: 285},
		{ID: 1411, Name: "Enyiogugu", LGID: 285},
		{ID: 1412, Name: "Ekwereazu", LGID: 285},

		{ID: 1413, Name: "Amuzi", LGID: 286},
		{ID: 1414, Name: "Umuhu", LGID: 286},
		{ID: 1415, Name: "Afor Oru", LGID: 286},
		{ID: 1416, Name: "Ibeku", LGID: 286},

		{ID: 1417, Name: "Umualumaku", LGID: 287},
		{ID: 1418, Name: "Umuezeala", LGID: 287},
		{ID: 1419, Name: "Agbaja", LGID: 287},
		{ID: 1420, Name: "Ihite", LGID: 287},

		{ID: 1421, Name: "Atta", LGID: 288},
		{ID: 1422, Name: "Itu", LGID: 288},
		{ID: 1423, Name: "Obizi", LGID: 288},
		{ID: 1424, Name: "Okpofe", LGID: 288},

		{ID: 1425, Name: "Arondizuogu", LGID: 289},
		{ID: 1426, Name: "Osina", LGID: 289},
		{ID: 1427, Name: "Akokwa", LGID: 289},
		{ID: 1428, Name: "Urualla", LGID: 289},

		{ID: 1429, Name: "Isiekenesi", LGID: 290},
		{ID: 1430, Name: "Ntueke", LGID: 290},
		{ID: 1431, Name: "Obodo Ukwu", LGID: 290},
		{ID: 1432, Name: "Umuma", LGID: 290},

		{ID: 1433, Name: "Amainyi", LGID: 291},
		{ID: 1434, Name: "Amaimo", LGID: 291},
		{ID: 1435, Name: "Umuezegwu", LGID: 291},
		{ID: 1436, Name: "Ihitte", LGID: 291},

		{ID: 1437, Name: "Amatta", LGID: 292},
		{ID: 1438, Name: "Iho", LGID: 292},
		{ID: 1439, Name: "Amaimo", LGID: 292},
		{ID: 1440, Name: "Avuvu", LGID: 292},

		{ID: 1441, Name: "Obolo", LGID: 293},
		{ID: 1442, Name: "Anara", LGID: 293},
		{ID: 1443, Name: "Umuduru", LGID: 293},
		{ID: 1444, Name: "Ibeme", LGID: 293},

		{ID: 1445, Name: "Umuoma", LGID: 294},
		{ID: 1446, Name: "Amaimo", LGID: 294},
		{ID: 1447, Name: "Umundugba", LGID: 294},
		{ID: 1448, Name: "Ekwe", LGID: 294},

		{ID: 1449, Name: "Ogwa", LGID: 295},
		{ID: 1450, Name: "Ubomiri", LGID: 295},
		{ID: 1451, Name: "Ifakala", LGID: 295},
		{ID: 1452, Name: "Orodo", LGID: 295},

		{ID: 1453, Name: "Ntu", LGID: 296},
		{ID: 1454, Name: "Umuneke", LGID: 296},
		{ID: 1455, Name: "Obiangwu", LGID: 296},
		{ID: 1456, Name: "Alulu", LGID: 296},

		{ID: 1457, Name: "Okwudor", LGID: 297},
		{ID: 1458, Name: "Nkume", LGID: 297},
		{ID: 1459, Name: "Atta", LGID: 297},
		{ID: 1460, Name: "Umuelemai", LGID: 297},

		{ID: 1461, Name: "Umudioka", LGID: 298},
		{ID: 1462, Name: "Eziama", LGID: 298},
		{ID: 1463, Name: "Umuezeala", LGID: 298},
		{ID: 1464, Name: "Umuozu", LGID: 298},

		{ID: 1465, Name: "Umueze", LGID: 299},
		{ID: 1466, Name: "Amaigbo", LGID: 299},
		{ID: 1467, Name: "Dimneze", LGID: 299},
		{ID: 1468, Name: "Eziama", LGID: 299},

		{ID: 1469, Name: "Obowo", LGID: 300},
		{ID: 1470, Name: "Umulogho", LGID: 300},
		{ID: 1471, Name: "Avutu", LGID: 300},
		{ID: 1472, Name: "Achara", LGID: 300},

		{ID: 1473, Name: "Oguta", LGID: 301},
		{ID: 1474, Name: "Awa", LGID: 301},
		{ID: 1475, Name: "Orsu Obodo", LGID: 301},
		{ID: 1476, Name: "Ejemekwuru", LGID: 301},

		{ID: 1477, Name: "Egbema", LGID: 302},
		{ID: 1478, Name: "Ohaji", LGID: 302},
		{ID: 1479, Name: "Assa", LGID: 302},
		{ID: 1480, Name: "Obitti", LGID: 302},

		{ID: 1481, Name: "Okigwe", LGID: 303},
		{ID: 1482, Name: "Amuro", LGID: 303},
		{ID: 1483, Name: "Ihube", LGID: 303},
		{ID: 1484, Name: "Ubahu", LGID: 303},

		{ID: 1485, Name: "Onuimo", LGID: 304},
		{ID: 1486, Name: "Okwe", LGID: 304},
		{ID: 1487, Name: "Owerre Nkworji", LGID: 304},
		{ID: 1488, Name: "Umunna", LGID: 304},

		{ID: 1489, Name: "Orlu", LGID: 305},
		{ID: 1490, Name: "Amakohia", LGID: 305},
		{ID: 1491, Name: "Eziachi", LGID: 305},
		{ID: 1492, Name: "Obibi", LGID: 305},

		{ID: 1493, Name: "Orsu", LGID: 306},
		{ID: 1494, Name: "Awo Idemili", LGID: 306},
		{ID: 1495, Name: "Ebenator", LGID: 306},
		{ID: 1496, Name: "Eziawa", LGID: 306},

		{ID: 1497, Name: "Omumma", LGID: 307},
		{ID: 1498, Name: "Amiri", LGID: 307},
		{ID: 1499, Name: "Ihite", LGID: 307},
		{ID: 1500, Name: "Iragbiji", LGID: 307},

		{ID: 1501, Name: "Mgbidi", LGID: 308},
		{ID: 1502, Name: "Ozara", LGID: 308},
		{ID: 1503, Name: "Amaokpara", LGID: 308},
		{ID: 1504, Name: "Nnempi", LGID: 308},

		{ID: 1505, Name: "Owerri", LGID: 309},
		{ID: 1506, Name: "New Owerri", LGID: 309},
		{ID: 1507, Name: "Owerri Municipal", LGID: 309},
		{ID: 1508, Name: "Egbu", LGID: 309},

		{ID: 1509, Name: "Naze", LGID: 310},
		{ID: 1510, Name: "Emekuku", LGID: 310},
		{ID: 1511, Name: "Emii", LGID: 310},
		{ID: 1512, Name: "Agbala", LGID: 310},

		{ID: 1513, Name: "Owerri West", LGID: 311},
		{ID: 1514, Name: "Irete", LGID: 311},
		{ID: 1515, Name: "Okuku", LGID: 311},
		{ID: 1516, Name: "Nekede", LGID: 311},

		// Jigawa
		// Auyo
		{ID: 1505, Name: "Auyo Town", LGID: 312},
		{ID: 1506, Name: "Auyo Village", LGID: 312},
		{ID: 1507, Name: "Auyo Suburb", LGID: 312},

		// Babura
		{ID: 1508, Name: "Babura Town", LGID: 313},
		{ID: 1509, Name: "Babura Village", LGID: 313},
		{ID: 1510, Name: "Babura Suburb", LGID: 313},

		// Biriniwa
		{ID: 1511, Name: "Biriniwa Town", LGID: 314},
		{ID: 1512, Name: "Biriniwa Village", LGID: 314},
		{ID: 1513, Name: "Biriniwa Suburb", LGID: 314},

		// Birnin Kudu
		{ID: 1514, Name: "Birnin Kudu Town", LGID: 315},
		{ID: 1515, Name: "Birnin Kudu Village", LGID: 315},
		{ID: 1516, Name: "Birnin Kudu Suburb", LGID: 315},

		// Buji
		{ID: 1517, Name: "Buji Town", LGID: 316},
		{ID: 1518, Name: "Buji Village", LGID: 316},
		{ID: 1519, Name: "Buji Suburb", LGID: 316},

		// Dutse
		{ID: 1520, Name: "Dutse Town", LGID: 317},
		{ID: 1521, Name: "Dutse Village", LGID: 317},
		{ID: 1522, Name: "Dutse Suburb", LGID: 317},

		// Gagarawa
		{ID: 1523, Name: "Gagarawa Town", LGID: 318},
		{ID: 1524, Name: "Gagarawa Village", LGID: 318},
		{ID: 1525, Name: "Gagarawa Suburb", LGID: 318},

		// Garki
		{ID: 1526, Name: "Garki Town", LGID: 319},
		{ID: 1527, Name: "Garki Village", LGID: 319},
		{ID: 1528, Name: "Garki Suburb", LGID: 319},

		// Gumel
		{ID: 1529, Name: "Gumel Town", LGID: 320},
		{ID: 1530, Name: "Gumel Village", LGID: 320},
		{ID: 1531, Name: "Gumel Suburb", LGID: 320},

		// Guri
		{ID: 1532, Name: "Guri Town", LGID: 321},
		{ID: 1533, Name: "Guri Village", LGID: 321},
		{ID: 1534, Name: "Guri Suburb", LGID: 321},

		// Gwaram
		{ID: 1535, Name: "Gwaram Town", LGID: 322},
		{ID: 1536, Name: "Gwaram Village", LGID: 322},
		{ID: 1537, Name: "Gwaram Suburb", LGID: 322},

		// Gwiwa
		{ID: 1538, Name: "Gwiwa Town", LGID: 323},
		{ID: 1539, Name: "Gwiwa Village", LGID: 323},
		{ID: 1540, Name: "Gwiwa Suburb", LGID: 323},

		// Hadejia
		{ID: 1541, Name: "Hadejia Town", LGID: 324},
		{ID: 1542, Name: "Hadejia Village", LGID: 324},
		{ID: 1543, Name: "Hadejia Suburb", LGID: 324},

		// Jahun
		{ID: 1544, Name: "Jahun Town", LGID: 325},
		{ID: 1545, Name: "Jahun Village", LGID: 325},
		{ID: 1546, Name: "Jahun Suburb", LGID: 325},

		// Kafin Hausa
		{ID: 1547, Name: "Kafin Hausa Town", LGID: 326},
		{ID: 1548, Name: "Kafin Hausa Village", LGID: 326},
		{ID: 1549, Name: "Kafin Hausa Suburb", LGID: 326},

		// Kaugama
		{ID: 1550, Name: "Kaugama Town", LGID: 327},
		{ID: 1551, Name: "Kaugama Village", LGID: 327},
		{ID: 1552, Name: "Kaugama Suburb", LGID: 327},

		// Kazaure
		{ID: 1553, Name: "Kazaure Town", LGID: 328},
		{ID: 1554, Name: "Kazaure Village", LGID: 328},
		{ID: 1555, Name: "Kazaure Suburb", LGID: 328},

		// Kiri Kasama
		{ID: 1556, Name: "Kiri Kasama Town", LGID: 329},
		{ID: 1557, Name: "Kiri Kasama Village", LGID: 329},
		{ID: 1558, Name: "Kiri Kasama Suburb", LGID: 329},

		// Kiyawa
		{ID: 1559, Name: "Kiyawa Town", LGID: 330},
		{ID: 1560, Name: "Kiyawa Village", LGID: 330},
		{ID: 1561, Name: "Kiyawa Suburb", LGID: 330},

		// Maigatari
		{ID: 1562, Name: "Maigatari Town", LGID: 331},
		{ID: 1563, Name: "Maigatari Village", LGID: 331},
		{ID: 1564, Name: "Maigatari Suburb", LGID: 331},

		// Malam Madori
		{ID: 1565, Name: "Malam Madori Town", LGID: 332},
		{ID: 1566, Name: "Malam Madori Village", LGID: 332},
		{ID: 1567, Name: "Malam Madori Suburb", LGID: 332},

		// Miga
		{ID: 1568, Name: "Miga Town", LGID: 333},
		{ID: 1569, Name: "Miga Village", LGID: 333},
		{ID: 1570, Name: "Miga Suburb", LGID: 333},

		// Ringim
		{ID: 1571, Name: "Ringim Town", LGID: 334},
		{ID: 1572, Name: "Ringim Village", LGID: 334},
		{ID: 1573, Name: "Ringim Suburb", LGID: 334},

		// Roni
		{ID: 1574, Name: "Roni Town", LGID: 335},
		{ID: 1575, Name: "Roni Village", LGID: 335},
		{ID: 1576, Name: "Roni Suburb", LGID: 335},

		// Sule Tankarkar
		{ID: 1577, Name: "Sule Tankarkar Town", LGID: 336},
		{ID: 1578, Name: "Sule Tankarkar Village", LGID: 336},
		{ID: 1579, Name: "Sule Tankarkar Suburb", LGID: 336},

		// Taura
		{ID: 1580, Name: "Taura Town", LGID: 337},
		{ID: 1581, Name: "Taura Village", LGID: 337},
		{ID: 1582, Name: "Taura Suburb", LGID: 337},

		// Yankwashi
		{ID: 1583, Name: "Yankwashi Town", LGID: 338},
		{ID: 1584, Name: "Yankwashi Village", LGID: 338},
		{ID: 1585, Name: "Yankwashi Suburb", LGID: 338},

	}
	

}
