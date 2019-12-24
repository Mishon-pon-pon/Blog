package main

import (
	"Blog/controllers"
	"Blog/libs/dataBase"
	"Blog/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DBExec(`Insert into Articles(Title, TextArticle) Values(
	// 	'New Post',
	// 	'Лорем ипсум долор сит амет, ест дицта албуциус ан, вел магна цоммуне рецусабо ет. Еум мутат солеат но, иудицабит перицулис тхеопхрастус еа нам. Вис аццусамус торяуатос еу, вис ат фацер тритани. Иудицо детрахит еурипидис вис ид, унум аудиам ест еи, ут вих велит волумус. Ут партем темпор лаореет вел, опортере суавитате дефинитионем еум еа. Еррор аеяуе лудус мел еу. Алияуид номинати еа сит. Алтера алтерум вис ет, реяуе цонвенире нам еу. Ерос персецути репримияуе вис ид, ет сале утрояуе нусяуам сит, сеа суммо маиорум сцаевола еа. Хис ид малис цаусае аццумсан, меи еу модо мазим. Еум цу дицат веритус инимицус, ин дицо яуидам мел, веро рецтеяуе суавитате еу хис. Еа яуи виде еиус, ат вел омниум дебитис волуптуа. Аццумсан елаборарет персеяуерис цум ет, дицтас нолуиссе продессет ут вим. Солута симилияуе ад меи, вих ан яуис сцрипсерит. Ад еум солум вениам лаореет, ех вис сале пондерум. Цу граеце алияуам адолесценс вис. Омниум фацилис антиопам иус ех, но цонсеяуат волуптатум сцрибентур про. Ид яуис волумус алиенум мел, еам порро сусципит перфецто еи, омнис лаборес интерессет про ет. Лабитур волумус ат хас, пер солеат ментитум еа. Про мовет денияуе те. Не иус алтера форенсибус интерпретарис, нам ут граеци перфецто, ест ан стет цонгуе адиписцинг. Сит лорем еффициенди ет, яуо ад сенсибус диссентиет пхилосопхиа. Цаусае фабеллас дефиниебас те хас, меи зрил оратио ин. Мелиоре вивендум цонсецтетуер не нец. Ет еум витае тантас темпорибус, пер легере ратионибус интеллегам но. Ат сеа одио зрил, неморе фацилиси инструцтиор ид яуо. Но вел садипсцинг персеяуерис. Цу евертитур инцидеринт сед, ет вих граеце интеллегам репудиандае.'
	// )`)
	oneMux := http.NewServeMux()

	oneMux.HandleFunc("/", indexHandler)
	oneMux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web"))))
	oneMux.HandleFunc("/test", controllers.PersonCreate)
	oneMux.HandleFunc("/admin", controllers.Admin)
	fmt.Println("server run on 3003")
	http.ListenAndServe(":3003", oneMux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	index := template.Must(template.ParseFiles(
		"web/index.html",
	))
	articles := []models.Article{}
	dataBase.DBQuery(`SELECT * FROM Articles`, func(result *sql.Rows) {
		a := models.Article{}
		for result.Next() {
			err := result.Scan(&a.ID, &a.Title, &a.TextArticle)
			if err != nil {
				panic(err)
			}
			articles = append(articles, a)
		}
	})
	index.Execute(w, articles)
}

func newPost(w http.ResponseWriter, r *http.Request) {
	var test struct {
		test string
	}
	err := json.NewDecoder(r.Body).Decode(&test)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(test)
}
