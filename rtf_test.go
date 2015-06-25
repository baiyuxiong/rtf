package rtf

import (
	"testing"
	"github.com/revel/revel"
	"time"
	"fmt"
	"strings"
)

func TestNumberFuncs(t *testing.T) {
	_,exists := revel.TemplateFuncs["nil"]
	if !exists{
		t.Error()
	}
	nilFunc := revel.TemplateFuncs["nil"].(func(interface{})bool)
	if !nilFunc(nil){
		t.Error()
	}
	if nilFunc("hello"){
		t.Error()
	}

	addFunc := revel.TemplateFuncs["add"].(func(int,int)int)
	if addFunc(1,2) != 3 {
		t.Error()
	}

	ueqFunc := revel.TemplateFuncs["ueq"].(func(interface{},interface{})bool)
	if (ueqFunc(addFunc(0,3),3)) {
		t.Error()
	}

	if (ueqFunc(nilFunc(nil),true)) {
		t.Error()
	}

	minusFunc := revel.TemplateFuncs["minus"].(func(int,int)int)
	if minusFunc(1,2) != -1 {
		t.Error()
	}


	multiplyFunc := revel.TemplateFuncs["multiply"].(func(...int)int)
	if multiplyFunc(1,2) != 2 {
		t.Error()
	}
	if multiplyFunc(0,1,2,3,4,5) != 0 {
		t.Error()
	}

	divideFunc := revel.TemplateFuncs["divide"].(func(...int)int)
	if divideFunc(6,2) != 3 {
		t.Error()
	}
	if divideFunc(32,2,2,2,2,2) != 1 {
		t.Error()
	}

	if divideFunc(3,2) != 1 {
		t.Error()
	}

	if divideFunc(1,2) != 0 {
		t.Error()
	}

	if divideFunc(1,0) != 0 {
		t.Error()
	}

	lessFunc := revel.TemplateFuncs["less"].(func(int,int)bool)
	if lessFunc(1,2) != true {
		t.Error()
	}
	if lessFunc(1,1) != false {
		t.Error()
	}
	if lessFunc(2,1) != false {
		t.Error()
	}

	lteFunc := revel.TemplateFuncs["lte"].(func(int,int)bool)
	if lteFunc(1,2) != true {
		t.Error()
	}
	if lteFunc(1,1) != true {
		t.Error()
	}
	if lteFunc(2,1) != false {
		t.Error()
	}

	gteFunc := revel.TemplateFuncs["gte"].(func(int,int)bool)
	if gteFunc(1,2) != false {
		t.Error()
	}
	if gteFunc(1,1) != true {
		t.Error()
	}
	if gteFunc(2,1) != true {
		t.Error()
	}

	percentageToIntFunc := revel.TemplateFuncs["percentageToInt"].(func(float32)int)
	if percentageToIntFunc(0.999) != 99 {
		t.Error()
	}

	percentFunc := revel.TemplateFuncs["percent"].(func(int64,int64)int64)
	if percentFunc(100,99) != 99 {
		t.Error()
	}
}


func TestDateFuncs(t *testing.T) {
	timestampFunc := revel.TemplateFuncs["timestamp"].(func()int64)
	now := time.Now().Unix();
	if now-1 > timestampFunc() || now+1 < timestampFunc(){
		t.Error()
	}

	yearFunc := revel.TemplateFuncs["year"].(func()int)
	if yearFunc()<2015 || yearFunc()>3000{
		t.Error()
	}

	monthFunc := revel.TemplateFuncs["month"].(func()string)
	if len(monthFunc()) <3 {
		t.Error()
	}

	weekFunc := revel.TemplateFuncs["week"].(func()string)

	if len(weekFunc())<0 || !strings.Contains(weekFunc(),"day"){
		t.Error()
	}

	dayFunc := revel.TemplateFuncs["day"].(func()int)
	if dayFunc()<0 || dayFunc()>31{
		t.Error()
	}

	todayFunc := revel.TemplateFuncs["today"].(func()string)
	if len(todayFunc()) != 10 || !strings.Contains(todayFunc(),fmt.Sprintf("%d",yearFunc())){
		t.Error()
	}

	nowFunc := revel.TemplateFuncs["now"].(func()string)
	if len(nowFunc()) != 16 || !strings.Contains(nowFunc(),fmt.Sprintf("%d",yearFunc())){
		t.Error()
	}

	formatTimeFunc := revel.TemplateFuncs["formatTime"].(func(time.Time,string)string)
	if len(formatTimeFunc(time.Now(),"2006-01-02 15:04:05")) != 19{
		t.Error()
	}

	dayLeftFunc := revel.TemplateFuncs["dayLeft"].(func(time.Time)int)
	if dayLeftFunc(time.Now().Add(time.Hour*25)) != 1{
		t.Error()
	}
	if dayLeftFunc(time.Now().Add(time.Hour*23)) != 0{
		t.Error()
	}
}


func TestStringFuncs(t *testing.T) {
	lowerFunc := revel.TemplateFuncs["lower"].(func(string)string)
	if lowerFunc("Hello") != "hello" {
		t.Error()
	}

	upperFunc := revel.TemplateFuncs["upper"].(func(string)string)
	if upperFunc("Hello") != "HELLO" {
		t.Error()
	}

	splitCommaFunc := revel.TemplateFuncs["splitComma"].(func(string)[]string)
	if len(splitCommaFunc("Hello")) != 1 {
		t.Error()
	}
	if len(splitCommaFunc("H,e,l,l,o")) != 5 {
		t.Error()
	}

	shortContentFunc := revel.TemplateFuncs["shortContent"].(func(string,int)string)
	if shortContentFunc("Hello",5) != "Hello"{
		t.Error()
	}
	if shortContentFunc("Hello",4) != "Hell..."{
		t.Error()
	}
	if shortContentFunc("中文english",3) != "中文e..."{
		t.Error()
	}

	joinFunc := revel.TemplateFuncs["join"].(func([]string,string)string)
	if joinFunc([]string{"a","b","c"},",") != "a,b,c" {
		t.Error()
	}

	containsFunc := revel.TemplateFuncs["contains"].(func(string,string)bool)
	if containsFunc("hello","h") != true {
		t.Error()
	}

	replaceFunc := revel.TemplateFuncs["replace"].(func(string,string,string)string)
	if replaceFunc("hello","h","H") != "Hello" {
		t.Error()
	}


	md5Func := revel.TemplateFuncs["md5"].(func(string)string)
	if md5Func("hello") != "5d41402abc4b2a76b9719d911017c592" {
		t.Error()
	}

	randomAlnumFunc := revel.TemplateFuncs["randomAlnum"].(func(int)string)
	tmp := randomAlnumFunc(5)
	for _,v := range tmp {
		if !((v >= 49 && v <= 57) || (v >= 65 && v <= 90) || (v >= 97 && v <= 122)){
			t.Error()
		}
	}
	if len(randomAlnumFunc(5)) != 5 {
		t.Error()
	}

	randomAlphasFunc := revel.TemplateFuncs["randomAlphas"].(func(int)string)
	tmp = randomAlphasFunc(5)
	for _,v := range tmp {
		if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122)){
			t.Error()
		}
	}
	if len(randomAlphasFunc(5)) != 5 {
		t.Error()
	}

	randomNumericFunc := revel.TemplateFuncs["randomNumeric"].(func(int)string)
	tmp = randomNumericFunc(5)
	for _,v := range tmp {
		if !(v >= 48 && v <= 57){
			t.Error()
		}
	}
	if len(randomNumericFunc(5)) != 5{
		t.Error()
	}
}