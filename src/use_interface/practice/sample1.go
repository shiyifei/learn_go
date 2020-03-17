package practice

//接口定义
type VowersFinder interface {
	FindVowers() []rune   //接口中的方法声明 不用加func关键字
}

//由于string类型是系统的，这里自定义一种MyString类型，针对该类型创建一个FindVowers()方法
type MyString string

//针对MyString类型定义一个接口中定义的方法，那么该类型就实现了VowersFinder接口
func (ms MyString) FindVowers() []rune{
	var vowers []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowers = append(vowers, rune)
		}
	}
	return vowers
}





