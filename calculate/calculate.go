package calculate

import (
	"github.com/pkg/errors"
	"strconv"
)

func Result(s string) string{

	num, err := calculate(s)
	output:=""
	if err != nil {
		output=s + "=" + err.Error()
		return output
	}
	output=s + "=" + strconv.Itoa(num)
	return output

}

func calculate(s string) (int, error) {
	if len(s) ==0 {
		return 0, nil
	}
	if s[0]=='-' {
		return 0, errors.New("invalid")
	}
	substring := ""

	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '(':
			j := i + 1
			count := 1
			for count > 0 {
				if s[j] == '(' {
					count++
				} else if s[j] == ')' {
					count--
				}
				j++
			}
			num, err := calculate(s[i+1 : j-1])
			i = j
			if err != nil {
				return 0, err
			}
			substring += strconv.Itoa(num)
		default:
			for j:=0; j < len(s); j++ {
				if j>=1 && (s[j-1] >= '0' && s[j-1] <= '9') && (s[j] >= '0' && s[j] <= '9'){
					return 0, errors.New("invalid")
				}
				if j>=2 && s[j-2]=='-' && s[j-1]=='-' && s[j] >= '0' && s[j] <= '9' {
					return 0, errors.New("invalid")
				}
				if j>=2 && s[j-2]=='(' && s[j-1]=='-' && s[j] >= '0' && s[j] <= '9' {
					return 0, errors.New("invalid")
				}
			}
			substring += string(s[i])
			i++
		}
	}
	return calcWithoutParenthesis(substring),nil
}

func calcWithoutParenthesis(s string) int {
	stack := make([]int, 0)
	i := 0
	var num int
	for i<len(s) {
		switch s[i] {
		case '+':
			i++
		case '-':
			i++
			num, i = getNextNum(s, i)
			stack = append(stack, -num)
		case '*':
			i++
			num, i = getNextNum(s, i)
			stack[len(stack)-1] *= num
		case '/':
			i++
			num, i = getNextNum(s, i)
			stack[len(stack)-1] /= num
		default:
			num, i = getNextNum(s, i)
			stack = append(stack, num)
		}
	}
	result := 0
	for _, n := range stack {
		result += n
	}
	return result
}

func getNextNum(s string, i int) (int, int) {
	num := 0
	for ; i < len(s) && (s[i] >= '0' && s[i] <= '9'); i++ {
		num = num * 10 + int(s[i] - '0')
	}
	return num, i
}
