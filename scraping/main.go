package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	uuid "github.com/satori/go.uuid"
	"github.com/sclevine/agouti"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"

	"./data"
)

func gormConnect() *gorm.DB {

	//TODO
	// 一旦DBをベタ書きしている
	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/treasure_app?parseTime=true&charset=utf8mb4&interpolateParams=true")

	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&data.Lecture{})
	return db
}

func dbInit() {
	db := gormConnect()
	defer db.Close()
}

func main() {

	// db初期化
	dbInit()

	// 番号を取得してoutput.txtに保存する。
	// 一度だけ行えばよく、output.txrがすでに存在する場合は、コメントアウトしてやればOK
	// ScrapingLectureNumber()

	// TODO
	/*
		**** DBの作成 ****
		1. outputの一行を読み込む（スライスに格納する）
		2. 読み取った一行と基本となるパスを組み合わせて、そのURLにアクセスする
		3. アクセスしたHTMLを取得する
		4. HTMLの中身を構造体に埋め込み、それをデータベースに格納する

		**** フロントにJSON形式でデータの受け渡し ****
		1. DBからデータをJSON形式でフロントに渡す

		**** 前項で単純なデータの受け渡しが可能なことが確認できたら、
		ルーティングも含めてページにアクセスされた時にどのようなデータをJSONで渡すかも含めて考える ****

	*/

	// linesにはoutput.txtの内容が一行ずつスライスとして格納されている
	lines := GetFileContents("output.txt")

	db, err := gorm.Open("mysql", "root:password@tcp(localhost:3306)/treasure_app?parseTime=true&charset=utf8mb4&interpolateParams=true")
	if err != nil {
		log.Printf("Failed to open DB: %v", err)
	}
	defer db.Close()

	// os.GETENV等で取得するように設定したい(暫定的に直接year変数に年代を代入するようにしている)
	year := "2019"

	baseURL := "https://syllabus.doshisha.ac.jp/html/" + year + "/"
	extension := ".html"

	// i を1にしておき、開発段階ではリクエストを1件しか送らずに、大学のサーバの負荷をなくす
	//NOTE
	for i := 0; i < 100; i++ {

		shortNumber := lines[i][1:5]

		rep := regexp.MustCompile(`-`)
		longNumber := rep.ReplaceAllString(lines[i], "")

		// output.txtの各行の末尾に空白が1文字分あるので、それを除くようにする
		longNumber = longNumber[0 : len(longNumber)-2]

		if len(longNumber) == 8 {
			longNumber = longNumber + "000"
		}

		// アクセスしたいURL
		fullURL := baseURL + shortNumber + "/" + longNumber + extension

		doc, _ := goquery.NewDocument(fullURL)

		year := "body > div > table:nth-child(2) > tbody > tr:nth-child(1) > td > p > b"
		id := "body > div > table:nth-child(2) > tbody > tr:nth-child(2) > td > table > tbody > tr > td:nth-child(1) > p > b"
		title := "body > div > table:nth-child(2) > tbody > tr:nth-child(2) > td > table > tbody > tr > td:nth-child(2) > p > font > b"
		titles := "body > div > table:nth-child(2) > tbody > tr:nth-child(2) > td > table > tbody > tr > td:nth-child(2) > p"
		teacher := "body > div > table:nth-child(2) > tbody > tr:nth-child(3) > td > table > tbody > tr > td:nth-child(2) > a"
		overview := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > p:nth-child(3)"
		goal := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > p:nth-child(5)"

		scehdule := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(7) > tbody"
		evaluationCriteria := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(10) > tbody"
		textbook := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(14) > tbody > tr:nth-child(1) > td > p"
		referenceURL := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(18) > tbody > tr:nth-child(1) > td > a"
		remarks := "body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > p:nth-child(20)"

		var lecture data.Lecture

		u2, err := uuid.NewV4()
		if err != nil {
			fmt.Printf("Something went wrong: %s", err)
			return
		}
		lecture.EvaluateID = u2.String()

		doc.Find(year).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			lecture.Year = body
		})
		doc.Find(id).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			// bodyの末に空白が1も自分入るのでそれを除去する
			body = body[0 : len(body)-2]
			lecture.LectureID = body
		})
		doc.Find(title).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			// ○ や △ という文字が先頭に含まれており、かな文字であるので3byteに値するので、先頭から3byte分取り除く
			body = body[3:]

			lecture.Title = body
		})
		doc.Find(titles).Each(func(i int, s *goquery.Selection) {

			// body には Title, SubTitle, EnglishTitle, Unit, Semester, Location, LectureStyleに値するものが入っている
			body := s.Text()

			// SubTitle に関する扱い
			if strings.Index(body, "(") != -1 {
				startSubTitle := strings.Index(body, "(")
				finnishSubTitle := strings.Index(body, ")")
				lecture.SubTitle = body[startSubTitle : finnishSubTitle+1]
			}

			// EnglishTitle に関する扱い
			alfabet := "abcdefghijklmnopqrstuvwxwzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			startEnglishTitle := strings.IndexAny(body, alfabet)
			finEnglishTitle := strings.Index(body[startEnglishTitle:], "\n") + startEnglishTitle
			lecture.EnglishTitle = body[startEnglishTitle:finEnglishTitle]

			// Unitに関する扱い
			unitIndex := strings.IndexAny(body, "単位/") - 1
			lecture.Unit, _ = strconv.Atoi(body[unitIndex : unitIndex+1])

			// Semesterに関する扱い
			// パターンとしては 春学期, 秋学期, 通年
			semesterIndex := strings.Index(body, "春学期/")
			if semesterIndex != -1 {
				lecture.Semester = body[semesterIndex : semesterIndex+9]
			}
			semesterIndex = strings.Index(body, "秋学期/")
			if semesterIndex != -1 {
				lecture.Semester = body[semesterIndex : semesterIndex+9]
			}
			semesterIndex = strings.Index(body, "春集中/")
			if semesterIndex != -1 {
				lecture.Semester = body[semesterIndex : semesterIndex+9]
			}
			semesterIndex = strings.Index(body, "秋集中/")
			if semesterIndex != -1 {
				lecture.Semester = body[semesterIndex : semesterIndex+9]
			}
			semesterIndex = strings.Index(body, "通年/")
			if semesterIndex != -1 {
				lecture.Semester = body[semesterIndex : semesterIndex+6]
			}

			// Locationに関する扱い
			// パターンとしては 京田辺, 今出川
			location := strings.Index(body, "京田辺/")
			if location != -1 {
				lecture.Location = body[location : location+9]
			}
			location = strings.Index(body, "今出川/")
			if location != -1 {
				lecture.Location = body[location : location+9]
			}

			// LectureStyleに関する扱い
			// パターンとしては 講義, 演習
			lectureStyle := strings.Index(body, "講義/")
			if lectureStyle != -1 {
				lecture.LectureStyle = body[lectureStyle : lectureStyle+6]
			}
			lectureStyle = strings.Index(body, "演習/")
			if lectureStyle != -1 {
				lecture.LectureStyle = body[lectureStyle : lectureStyle+6]
			}
		})
		doc.Find(teacher).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			// 改行を取り除く
			rep := regexp.MustCompile(`\n`)
			body = rep.ReplaceAllString(body, "")

			lecture.Teacher = body
		})
		doc.Find(overview).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			lecture.Overview = body
		})
		doc.Find(goal).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			lecture.Goal = body
		})

		// body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(7) > tbody > tr:nth-child(2)
		// body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(7) > tbody > tr:nth-child(2) > td:nth-child(2)
		// body > div > table:nth-child(2) > tbody > tr:nth-child(4) > td > table:nth-child(7) > tbody > tr:nth-child(3)
		var sce data.Scehdule
		s := doc.Find(scehdule)
		length := s.Find("tr ").Length()
		sce.LectureID = lecture.LectureID
		for i := 2; i <= length; i++ {
			st := s.Find("tr:nth-child(" + strconv.Itoa(i) + ")" + "> td:nth-child(2)")
			body := st.Text()

			// 改行を取り除く
			rep := regexp.MustCompile(`\n`)
			body = rep.ReplaceAllString(body, "")

			sce.Session = body
			db.Create(&sce)
		}

		var evaluate data.Evaluate
		s = doc.Find(evaluationCriteria)
		length = s.Find("tr > td:nth-child(1) ").Length()
		evaluate.ID = lecture.EvaluateID
		for i := 1; i <= length; i++ {
			st1 := s.Find("tr:nth-child(" + strconv.Itoa(i) + ")" + " > td:nth-child(1) ").First()
			body := st1.Text()

			// 改行を取り除く
			rep := regexp.MustCompile(`\n`)
			body = rep.ReplaceAllString(body, "")

			evaluate.Method = body

			st2 := s.Find("tr:nth-child(" + strconv.Itoa(i) + ")" + " > td:nth-child(2) ").First()
			body = st2.Text()

			// 改行を取り除く
			body = rep.ReplaceAllString(body, "")

			evaluate.Percentage = body

			st3 := s.Find("tr:nth-child(" + strconv.Itoa(i) + ")" + " > td:nth-child(3) ").First()
			body = st3.Text()

			// 改行を取り除く
			body = rep.ReplaceAllString(body, "")

			evaluate.Comment = body

			// ok := db.NewRecord(&evaluate)
			// if !ok {
			// 	log.Printf("Failed to create new data: %v", err)
			// 	panic("Failed to create new data")
			// }
			db.Create(&evaluate)
		}

		doc.Find(textbook).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			// 改行を取り除く
			rep := regexp.MustCompile(`\n`)
			body = rep.ReplaceAllString(body, "")

			// 両端に "  " が含まれているので、トリムする
			body = strings.Trim(body, "  ")

			lecture.Textbook = body
		})
		doc.Find(referenceURL).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			lecture.ReferenceURL = body
		})
		doc.Find(remarks).Each(func(i int, s *goquery.Selection) {

			body := s.Text()

			lecture.Remarks = body
		})

		ok := db.NewRecord(&lecture)
		if !ok {
			log.Printf("Failed to create new data: %v", err)
			panic("Failed to create new data")
		}

		db.Create(&lecture)
	}

}

// GetFileContents is
func GetFileContents(filePath string) []string {
	// ファイルを開く
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to file oepn: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	lines := make([]string, 0, 100000)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// linesスライスに一行ずつテキストを保存する
		lines = append(lines, scanner.Text())
	}
	if scanErr := scanner.Err(); scanErr != nil {
		log.Printf("Failed to file scan: %v", err)
	}
	return lines
}

// ScrapingLectureNumber 関数は講義番号を取得する関数である
func ScrapingLectureNumber() {
	// urlの指定を行う
	url := "https://syllabus.doshisha.ac.jp/"

	// chromeドライバーの起動
	driver := agouti.ChromeDriver()

	err := driver.Start()
	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	// 指定したURLにアクセスする
	err = page.Navigate(url)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	// getSearchAmountでデフォルトで20件となっているところを1000件にする
	getSearchAmount := "body > table:nth-child(4) > tbody > tr > td > form > table:nth-child(10) > tbody > tr > td:nth-child(2) > input[type=text]"
	page.Find(getSearchAmount).Fill("1000")

	// Buttonを調べて、そのボタンを押下する
	page.FindByButton("検索/Search").Click()

	//////////////	ページ遷移	/////////////////

	// contentには "検索/Search" を押下した先のページのHTMLの内容が入っている
	content, err := page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	reader := strings.NewReader(content)

	// 現状ブラウザで開いているページのDOMを取得する
	doc, _ := goquery.NewDocumentFromReader(reader)

	// ファイルを書き込み用にオープン (mode=0666)
	file, err := os.Create("./output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 1ページに1000件のデータがあるので、domNumberを1000とする
	for domNumber := 1; domNumber <= 1000; domNumber++ {
		// lectureNumber は講座番号のselectorである
		lectureNumber := "body > table:nth-child(4) > tbody > tr > td > table > tbody > tr:nth-child(" + strconv.Itoa(domNumber) + ") > td:nth-child(1)"
		doc.Find(lectureNumber).Each(func(i int, s *goquery.Selection) {

			// body にはlectureNumberにマッチした中にあるテキストが含まれている
			body := s.Text()

			// テキストを書き込む
			_, err = file.WriteString(body + "\n")
			if err != nil {
				panic(err)
			}
		})
	}

	// "次結果一覧/Next"Buttonがある場合
	for page.FindByButton("次結果一覧/Next") != nil {
		page.FindByButton("次結果一覧/Next").Click()

		//////////////	ページ遷移	/////////////////

		// contentには "検索/Search" を押下した先のページのHTMLの内容が入っている
		content, err := page.HTML()
		if err != nil {
			log.Printf("Failed to get html: %v", err)
		}

		reader := strings.NewReader(content)

		// 現状ブラウザで開いているページのDOMを取得する
		doc, _ := goquery.NewDocumentFromReader(reader)

		for domNumber := 1; domNumber <= 1000; domNumber++ {
			// lectureNumber は講座番号のselectorである
			lectureNumber := "body > table:nth-child(4) > tbody > tr > td > table > tbody > tr:nth-child(" + strconv.Itoa(domNumber) + ") > td:nth-child(1)"
			doc.Find(lectureNumber).Each(func(i int, s *goquery.Selection) {

				// body にはlectureNumberにマッチした中にあるテキストが含まれている
				body := s.Text()

				// テキストを書き込む
				_, err = file.WriteString(body + "\n")
				if err != nil {
					panic(err)
				}
			})
		}
	}
}
