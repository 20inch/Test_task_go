package main

import "fmt"
import "strconv"

// функция перевода римских чисел типа string в арабские типа int
func romanToInt(roman string) (int,bool) {
	
	var result int
    var typeMathRoman bool
		
		switch roman {
		case "I": result = 1
		case "II": result = 2
		case "III": result = 3
		case "IV": result = 4
		case "V": result = 5
		case "VI": result = 6
		case "VII": result = 7
		case "VIII": result = 8
		case "IX": result = 9
		case "X": result = 10
		default: panic("Неверный ввод данных: неверно ввели число") 
		}
		
		typeMathRoman = true
	    return result,typeMathRoman
	
}

// функция перевода переменной с арабскими числами типа string в тип int
func strToInt(input string) (int,bool) {
        var number int
        typeMathInt := false
        number, err := strconv.Atoi(input)
        
        if err != nil { 
            number,typeMathInt = romanToInt(input)
        }
        
        return number,typeMathInt
        
}

// функция конвертирования результата математической операции в римское представление
func arabToRoman(num int) string {
	var result string
	//число должно быть не меньше 1
	if num < 1 {
		fmt.Print("Ошибка вывода результата: в римской системе исчисления нет значений меньше 1")
	}
	// если число больше или равно 100
	for num >=100 {
	    result+="C"
	    num-=100
	}
	// если число больше или равно 90
	if num>=90{
	    result+="CX"
	    num-=90
	}
	// если число больше или равно 50
	for num >=50 {
    	result+="L"
    	num-=50
	}
    // если число больше или равно 40
	if num>=40{
	    result+="XL"
	    num-=40
	}
    // если число больше или равно 10
	for num >=10 {
    	result+="X"
    	num-=10
	}
    // если число больше или равно 9
	if num>=9{
	    result+="IX"
	    num-=9
	}
	// если число больше или равно 5
	for num >=5 {
    	result+="V"
    	num-=5
	}
	// если число больше или равно 4
	if num>=4{
	    result+="IV"
	    num-=4
	}
	// если число больше или равно 1
	for num >=1 {
    	result+="I"
    	num-=1
	}
	
	return result 
	
}

func main() {
	// объявляем переменные
	var inputUserNum1, inputUserMathOperation, inputUserNum2 string
	inputLeghtErr:=""
	var num1,num2,itog int
	var isRoman1,isRoman2 bool
	
	// считываем данные введенные пользователем в переменные
	fmt.Println("Введитете выражение в формате: число1 +-*/ число2")
	fmt.Println("Принимаются арабские числа от 1 до 10 и римские от I до X")
	fmt.Scanln(&inputUserNum1,&inputUserMathOperation,&inputUserNum2,&inputLeghtErr)
	
	// проверка полноты ввода данных
	if inputUserNum1=="" || inputUserNum2=="" || inputUserMathOperation==""{
	panic("Неверный ввод данных: строка не является математической операцией.")  
	}
	if inputLeghtErr!="" {
    panic("Неверный ввод данных: избыточный объём данных")
    }
	
	// конвертирование введенных пользователем значений в переменные типа int
	num1,isRoman1 = strToInt(inputUserNum1)
    num2,isRoman2 = strToInt(inputUserNum2)
    
    // проверка диапазона значений и системы исчесления
    if isRoman1!=isRoman2 {
    panic("Неверный ввод данных: используются одновременно разные системы исчисления")
    }
    
    if num1 < 1|| num1>10 || num2 < 1|| num2>10  {
    panic("Неверный ввод данных: неверно ввели число")
    }
    
    // выполняем математическую операцию
    switch inputUserMathOperation {
       case "+": itog = num1+num2
       case "-": itog = num1-num2
       case "*": itog = num1*num2
       case "/": itog = num1/num2
       default: panic("Неверный ввод данных: неверная математическая операция")
   }
   
   // проверка итогового результата для вывода в римской системе исчисления
   if isRoman1==true && isRoman2==true && itog<0{
       panic("Ошибка вывода результата: в римской системе исчисления нет отрицательных чисел")
   }
   
   // вывод результата операции в арабской или римской системе исчисления
   if isRoman1==false{
       fmt.Println(itog)
   } else {
       fmt.Println(arabToRoman(itog))
   }
   
}