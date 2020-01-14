package main

import(
  ."fmt"
  "strings"
)

func Massimo(len1, len2 int)(int,int){
  if len1>len2{
    return len1,len2
  }else if len2>len1{
    return len2,len1
  }else{
    return len1,len2
  }
}

func main(){
  var s1,s2 string
  k1 := 0
  k2 := 0
  Print("inserire le parole: ")
  Scanln(&s1,&s2)

  str1 :=strings.Split(s1,"")
  str2 :=strings.Split(s2,"")
  Println(str1,str2)
  len1 := len(str1)
  len2 := len(str2)
  LenM,min := Massimo(len1,len2)
  somLen := len1+len2 + (LenM-min)
  str :=make([]string,somLen)
  for i:=1;i<somLen+1;i++{
    if i%2!=0{
      if k1<len1{
        str[i-1]=str1[k1]
        k1++
      }else{
        str[i-1]="-"
      }
    }else{
      if k2<len2{
          str[i-1]=str2[k2]
          k2++
      }else{
      str[i-1]="-"
      }
    }
  }
    Println(str)
  }
