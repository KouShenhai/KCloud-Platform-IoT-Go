package core

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	// =======================基础操作=======================
	s := "hello"
	// 输出 => hello world
	fmt.Printf("%s %s", s, "world")
	fmt.Println()
	s += " world"
	// 输出 => hello world
	fmt.Println(s)
	// 输出 => 11
	fmt.Println(len(s))
	// 输出 => h
	fmt.Println(string(s[0]))
	// 输出 => ello world
	fmt.Println(s[1:])
	// =======================基础操作=======================

	// =======================字符串遍历=======================
	for index, v := range s {
		fmt.Printf("%d：%s", index, string(v))
		fmt.Println()
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d：%s", i, string(s[i]))
		fmt.Println()
	}
	// =======================字符串遍历=======================

	// =======================strings操作=======================
	// 输出 => true
	fmt.Println(strings.Contains("test", "es"))
	// 输出 => 2
	fmt.Println(strings.Count("test", "t"))
	// 输出 => 1
	fmt.Println(strings.Index("test", "e"))
	// 输出 => 1
	fmt.Println(strings.LastIndex("test", "e"))
	// 输出 => a-b
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// 输出 => tstt
	// -1表示替换所有
	// 1表示替换第一个
	// 0表示不替换
	fmt.Println(strings.Replace("test", "es", "st", -1))
	// 输出 => test
	fmt.Println(strings.Replace("test", "es", "st", 0))
	// 输出 => tstt
	fmt.Println(strings.Replace("test", "es", "st", 1))
	// 输出 => true
	fmt.Println(strings.HasSuffix("test", "st"))
	// 输出 => true
	fmt.Println(strings.HasPrefix("test", "te"))
	// 输出 => A B C
	fmt.Println(strings.Trim(" !!! A B C !!! ", " !!! "))
	// 输出 =>  !!! A B C
	fmt.Println(strings.TrimSuffix(" !!! A B C !!! ", " !!! "))
	// 输出 => A B C !!!
	fmt.Println(strings.TrimPrefix(" !!! A B C !!! ", " !!! "))
	// 输出 => !!! A B C !!!
	fmt.Println(strings.TrimSpace(" !!! A B C !!! "))
	// 输出 => A B C !!!
	fmt.Println(strings.TrimLeft(" !!! A B C !!! ", " !!! "))
	// 输出 =>  !!! A B C
	fmt.Println(strings.TrimRight(" !!! A B C !!! ", " !!! "))
	// 输出 => A B C !!!
	fmt.Println(strings.TrimLeftFunc(" !!! A B C !!! ", func(r rune) bool {
		return r == ' ' || r == '!'
	}))
	// 输出 =>  !!! A B C
	fmt.Println(strings.TrimRightFunc(" !!! A B C !!! ", func(r rune) bool {
		return r == ' ' || r == '!'
	}))
	// 输出 => [a b c]
	fmt.Println(strings.Split("a,b,c", ","))
	// 输出 => [a b,c]
	fmt.Println(strings.SplitN("a,b,c", ",", 2))
	// 输出 => [a, b, c]
	fmt.Println(strings.SplitAfter("a,b,c", ","))
	// 输出 => [a, b,c]
	fmt.Println(strings.SplitAfterN("a,b,c", ",", 2))
	// 输出 => [a b c d]
	fmt.Println(strings.Fields(" a b c\td"))
	// 输出 => [a b c d]
	fmt.Println(strings.FieldsFunc(" a b c\td", func(r rune) bool {
		return r == ' ' || r == '\t'
	}))
	// 输出 => AAb
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'A'
		}
		return r
	}, "aab"))
	// 输出 => AAb
	fmt.Println(strings.ToUpper("aab"))
	// 输出 => aab
	fmt.Println(strings.ToLower("AAB"))
	// 输出 => AAB
	fmt.Println(strings.ToTitle("AAB"))
	// 输出 => aab
	fmt.Println(strings.ToValidUTF8("aab", "b"))
	// =======================strings操作=======================

	// =======================字符串转换=======================
	// 输出 => abc
	fmt.Println(string([]byte{97, 98, 99}))
	// 输出 => abc
	fmt.Println(string([]rune{97, 98, 99}))
	// 输出 => abc
	fmt.Println(string([]rune{0x61, 0x62, 0x63}))
	// 输出 => 1
	numStr := "123"
	num, _ := fmt.Sscan(numStr, &numStr)
	fmt.Println(num)
	// 输出 => 123
	num, _ = strconv.Atoi(numStr)
	fmt.Println(num)
	// 输出 => 123
	numStr = strconv.Itoa(num)
	fmt.Println(numStr)
	// =======================字符串转换=======================

	// =======================字符串构建=======================
	// 输出 => hello world!
	builder := strings.Builder{}
	builder.WriteString("hello")
	builder.WriteRune(' ')
	builder.WriteString("world")
	builder.WriteString("!")
	fmt.Println(builder.String())
	// 输出 => hello
	builder.Reset()
	builder.WriteString("hello")
	fmt.Println(builder.String())
	// 输出 => hello world!
	var buf bytes.Buffer
	buf.WriteString("hello")
	buf.WriteRune(' ')
	buf.WriteString("world")
	buf.WriteString("!")
	fmt.Println(buf.String())
	// =======================字符串构建=======================
}
