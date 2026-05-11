
package main 


func reverse(s string)string{

	str:=[]rune(s)
	i, j:=0, len(str)-1


	for i<j {
		str[i], str[j]=str[j], str[i]
		i++
		j--
	}

	return string(str)
}