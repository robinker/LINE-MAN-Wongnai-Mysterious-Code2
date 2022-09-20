package main

import (
	"fmt"
	"github.com/robinker/LINE-MAN-Wongnai-Mysterious-Code2/util"
	"github.com/robinker/LINE-MAN-Wongnai-Mysterious-Code2/decoder"
)


func main() {
	var whatIsIt string;
	secret := "CYtZBsWZaZliYZocWLZlXuZZYWYeYXZsXeZXtXWpXeRYYYd!ZnYeWXoYXasnX,WXWrWPoAdWesnciGenWr";
	meaningLessChars := util.FindMeaningLessCharacter(secret);
	cipherText := util.NormalizeString(secret, meaningLessChars)
	fmt.Println("Normalized secret: ", cipherText)
	whatIsIt = decoder.DecodeZigZagPattern(cipherText, len(meaningLessChars))
	fmt.Println("whatIsIt: ", whatIsIt)
}

/*
	x-----x-----x-----x
	-x---x-x---x-x---x-
	--x-x---x-x---x-x--
	---x-----x-----x---
*/
