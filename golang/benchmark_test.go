package golang_test

import (
	"bytes"
	"math/rand"
	"strconv"
	"testing"
	"text/template"
	"time"
)

const (
	LOG_LEN       = 144
	WITH_DELAY    = true
	WITHOUT_DELAY = false
	ITEMS_COUNT   = 100
)

func Benchmark_WriteString(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.WriteString("123456789 ")
	}
}

func Benchmark_WriteStringAsBytes(b *testing.B) {
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.Write([]byte("123456789 "))
	}
}

func benchmark_buffer2buffer(b *testing.B, write func(source, target *bytes.Buffer)) {
	buf1 := new(bytes.Buffer)
	buf1.Write([]byte("123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 123456789 "))
	for i := 0; i < b.N; i++ {
		buf2 := new(bytes.Buffer)
		for i := 0; i < ITEMS_COUNT; i++ {
			write(buf2, buf1)
		}
		_ = buf2.String()
	}
}

func Benchmark_BufferString(b *testing.B) {
	// Writing as string is slower comparing to writing as bytes
	benchmark_buffer2buffer(b, func(source, target *bytes.Buffer) {
		target.WriteString(source.String())
	})
}

func Benchmark_BufferBytes(b *testing.B) {
	// Writing as bytes is faster comparing to writing as string
	benchmark_buffer2buffer(b, func(source, target *bytes.Buffer) {
		target.Write(source.Bytes())
	})
}

type Renderer interface {
	Render(writer *bytes.Buffer) error
}

type TemplateControl int

type UserCard struct {
	User UserData
}

func delay() {
	time.Sleep(time.Duration(10)*time.Millisecond + time.Duration(rand.Intn(10)))
}

func (t UserCard) Render(writer *bytes.Buffer) error {
	_, err := writer.Write([]byte("<div id=user"))
	_, err = writer.Write([]byte(strconv.Itoa(t.User.Id)))
	_, err = writer.Write([]byte(" class="))
	if t.User.Id%2 == 0 {
		_, err = writer.Write([]byte("even"))
	} else {
		_, err = writer.Write([]byte("odd "))
	}
	_, err = writer.Write([]byte(">"))
	_, err = writer.Write([]byte(t.User.Name()))
	_, err = writer.Write([]byte(", DOB: "))
	_, err = writer.Write([]byte(t.User.DateOfBirth.Format("02 January 2006")))
	_, err = writer.Write([]byte("</div>"))
	return err
}

func Test_renderCard(t *testing.T) {
	writer := new(bytes.Buffer)
	err := UserCard{User: UserData{}}.Render(writer)
	if err != nil {
		t.FailNow()
	}
	_ = writer.String()
}

func Test_SequentialVsGoRoutineVsStandard(t *testing.T) {
	users := getUsers(WITHOUT_DELAY)

	renderUsingGoRouting := func() string {
		buffer := new(bytes.Buffer)
		err := TemplateUserCardsPage{Users: users}.Render(buffer)
		if err != nil {
			t.FailNow()
		}
		return buffer.String()
	}

	renderUsingTemplateText := func() string {
		t1, err := template.New("simple").Funcs(funcMap()).Parse(STANDARD_GO_PAGE_TEMPLATE)
		if err != nil {
			t.FailNow()
		}
		buffer := new(bytes.Buffer)
		err = t1.ExecuteTemplate(buffer, "simple", users)
		if err != nil {
			t.FailNow()
		}
		return buffer.String()
	}

	sequentialResult, err := renderSequential(users)
	if err != nil {
		t.FailNow()
	}

	goRoutinResult := renderUsingGoRouting()

	if goRoutinResult != sequentialResult {
		t.Error("goRoutinResult != sequentialResult", len(goRoutinResult), len(sequentialResult))
		t.Log("GoRouting: " + goRoutinResult + "|")
		t.Log("Sequential: " + sequentialResult + "|")
	}

	textTemplateResult := renderUsingTemplateText()
	if textTemplateResult != sequentialResult {
		t.Error("textTemplateResult != sequentialResult", len(textTemplateResult), len(sequentialResult))
		t.Log("Template: " + textTemplateResult + "|")
		t.Log("Sequential: " + sequentialResult + "|")
	}
}

type UserData struct {
	Id          int
	name        string
	delay       bool
	DateOfBirth time.Time
}

func (u UserData) Name() string {
	if u.delay {
		delay()
	}
	return u.name
}

func getUsers(delay bool) []UserData {
	var items [ITEMS_COUNT]UserData

	for i := 1; i < ITEMS_COUNT+1; i++ {
		items[i-1] = UserData{
			Id:          i,
			delay:       delay,
			name:        "User #" + strconv.Itoa(i),
			DateOfBirth: time.Now().Add(-time.Duration(24*i) * time.Hour),
		}
	}
	return items[:]
}

func renderSequential(users []UserData) (string, error) {
	writer := new(bytes.Buffer)
	_, err := writer.Write([]byte("<ul>\n"))
	for _, user := range users {
		_, err = writer.Write([]byte("\t<li>"))
		if err != nil {
			return writer.String(), err
		}
		err = UserCard{User: user}.Render(writer)
		if err != nil {
			return writer.String(), err
		}
		_, err = writer.Write([]byte("</li>\n"))
		if err != nil {
			return writer.String(), err
		}
	}
	writer.Write([]byte("</ul>"))
	return writer.String(), err
}

func benchmark_Sequential(b *testing.B, delay bool) {
	b.StopTimer()
	users := getUsers(delay)
	b.StartTimer()
	var s string
	var err error
	for i := 0; i < b.N; i++ {
		s, err = renderSequential(users)
		if err != nil {
			b.FailNow()
		}
	}
	//time.Sleep(time.Second*1) // Why it affects number of operations?
	b.StopTimer()
	_ = s // b.Log(s[:LOG_LEN])
}

const STANDARD_GO_PAGE_TEMPLATE = "{{ define \"user_card\" }}<div id=user{{ .Id }} class={{mod2 .Id}}>{{ .Name }}, DOB: {{ .DateOfBirth.Format \"02 January 2006\" }}</div>{{ end }}<ul>\n{{range $user := .}}\t<li>{{ template \"user_card\" $user }}</li>\n{{end}}</ul>"

func funcMap() template.FuncMap {
	return template.FuncMap{
		"mod2": func(i int) string {
			if i%2 == 0 {
				return "even"
			} else {
				return "odd "
			}
		},
	}
}

/*
func Benchmark_StandardTemplateWithoutDelayColdStart(b *testing.B) {
	b.StopTimer()
	users := getUsers(WITHOUT_DELAY)
	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		t, _ := template.New("simple").Funcs(funcMap()).Parse(STANDARD_GO_PAGE_TEMPLATE)
		wr := new(bytes.Buffer)
		t.ExecuteTemplate(wr, "simple", users)
		s = wr.String()
	}
	b.StopTimer()
	_ = s // b.Log(s[:LOG_LEN])
}
*/

func Benchmark_StandardTemplateWithoutDelayWarmed(b *testing.B) {
	b.StopTimer()
	users := getUsers(WITHOUT_DELAY)
	t, _ := template.New("simple").Funcs(funcMap()).Parse(STANDARD_GO_PAGE_TEMPLATE)
	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		wr := new(bytes.Buffer)
		t.ExecuteTemplate(wr, "simple", users)
		s = wr.String()
	}
	b.StopTimer()
	_ = s // b.Log(s[:LOG_LEN])
}

func Benchmark_StandardTemplateWithtDelayWarmed(b *testing.B) {
	b.StopTimer()
	users := getUsers(WITH_DELAY)
	t, _ := template.New("simple").Funcs(funcMap()).Parse(STANDARD_GO_PAGE_TEMPLATE)
	b.StartTimer()
	var s string
	for i := 0; i < b.N; i++ {
		wr := new(bytes.Buffer)
		t.ExecuteTemplate(wr, "simple", users)
		s = wr.String()
	}
	b.StopTimer()
	_ = s // b.Log(s[:LOG_LEN])
}

func Benchmark_SequentialWithoutDelay(b *testing.B) {
	benchmark_Sequential(b, WITHOUT_DELAY)
}

func Benchmark_SequentialWithDelay(b *testing.B) {
	benchmark_Sequential(b, WITH_DELAY)
}

type RenderOutput struct {
	Buffer *bytes.Buffer
	Err    error
}

type RenderItem struct {
	Output chan RenderOutput
	Next   *RenderItem
}

func (item *RenderItem) ProcessChain(writer *bytes.Buffer) error {
	for {
		for output := range item.Output {
			if output.Err != nil {
				return output.Err
			}
			_, err := writer.Write(output.Buffer.Bytes())
			if err != nil {
				return err
			}
		}
		if item.Next == nil {
			return nil
		}
		item = item.Next
	}
}

type TemplateUserCardsPage struct {
	Users []UserData
}

func (t TemplateUserCardsPage) Render(writer *bytes.Buffer) error {
	writer.Write([]byte("<ul>\n"))
	var prevItem, firstItem *RenderItem

	_renderRangeItem := func(i int, renderItem RenderItem, userData UserData) {
		writer := new(bytes.Buffer)
		var err error
		defer func() {
			renderItem.Output <- RenderOutput{Buffer: writer, Err: err}
			close(renderItem.Output)
		}()
		_, err = writer.Write([]byte("\t<li>")) // From(line:10,pos:1)-To(line:10,pos:4)
		if err != nil {
			return
		}
		err = UserCard{User: userData}.Render(writer)
		if err != nil {
			return
		}
		_, err = writer.Write([]byte("</li>\n"))
	}

	for i, userData := range t.Users {
		renderItem := RenderItem{Output: make(chan RenderOutput, 1)}
		if i == 0 {
			firstItem = &renderItem
		} else {
			prevItem.Next = &renderItem
		}
		prevItem = &renderItem
		go _renderRangeItem(i, renderItem, userData)
	}
	if firstItem != nil {
		firstItem.ProcessChain(writer)
	}

	_, err := writer.Write([]byte("</ul>"))
	return err
}

func benchmark_GoRoutine(b *testing.B, delay bool) {
	var s string
	b.StopTimer()
	users := getUsers(delay)
	b.StartTimer()
	t := TemplateUserCardsPage{Users: users}
	for i := 0; i < b.N; i++ {
		writer := new(bytes.Buffer)
		_ = t.Render(writer)
		s = writer.String()
	}
	b.StopTimer()
	_ = s //b.Log(s[:LOG_LEN])
}
func Benchmark_GoRoutineWithoutDelay(b *testing.B) {
	benchmark_GoRoutine(b, WITHOUT_DELAY)
}
func Benchmark_GoRoutineWithDelay(b *testing.B) {
	benchmark_GoRoutine(b, WITH_DELAY)
}
