package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"

	mail "gopkg.in/mail.v2"
)

func main() {
	err := readConfig("config.json")
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}
	learningTool()
}

//学习工具
func learningTool() {
	fmt.Println("欢迎使用单词学习工具！")

	// 读取配置文件中的单词列表
	words, err := readWordsFromFile("words.json")
	if err != nil {
		fmt.Println("从文件中读取单词时出错：", err)
		os.Exit(1)
	}

	fmt.Println("请选择模式:")
	fmt.Println("1. 简单模式（英文到中文）")
	fmt.Println("2. 困难模式（中文到英文）")
	var mode int
	fmt.Scanln(&mode)

	var correctAnswers int // 正确答案的数量
	var totalQuestions int // 总问题数

	for totalQuestions < config.TotalQuestions {
		var word Word

		// 根据模式选择单词
		if mode == 1 {
			word = randomWord(words)
			fmt.Printf("'%s' 的中文含义是什么？ ", word.Text)
		} else if mode == 2 {
			word = randomWord(words)
			fmt.Printf("'%s' 的英文单词是什么？ ", word.Meaning)
		} else {
			fmt.Println("无效的模式。请选择 1 或 2。")
			continue
		}

		var userAnswer string
		fmt.Scanln(&userAnswer)

		if mode == 1 {
			if userAnswer == word.Meaning {
				fmt.Println("回答正确！")
				correctAnswers++
			} else {
				fmt.Printf("回答错误。'%s' 的中文含义是 '%s'。\n", word.Text, word.Meaning)
			}
		} else if mode == 2 {
			if userAnswer == word.Text {
				fmt.Println("回答正确！")
				correctAnswers++
			} else {
				fmt.Printf("回答错误。'%s' 的英文单词是 '%s'。\n", word.Meaning, word.Text)
			}
		}

		totalQuestions++
	}

	// 计算正确率
	accuracy := float64(correctAnswers) / float64(totalQuestions) * 100
	fmt.Printf("你的正确率是 %.2f%%\n", accuracy)

	// 如果正确率达到95%，发送邮件给家长
	if accuracy >= float64(config.PassingScore) {
		sendEmail(accuracy)
	}
}

var config *Config

type Config struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	ParentEmail    string `json:"parent_email"`
	TeacherEmail   string `json:"teacher_email"`
	SMTPServer     string `json:"smtp_server"`
	SMTPPort       int    `json:"smtp_port"`
	TotalQuestions int    `json:"total_questions"`
	PassingScore   int    `json:"passing_score"`
}

// 单词结构体
type Word struct {
	Text    string `json:"text"`
	Meaning string `json:"meaning"`
}

// 从配置文件中读取单词列表
func readWordsFromFile(filename string) ([]Word, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var words []Word
	err = json.Unmarshal(file, &words)
	if err != nil {
		return nil, err
	}

	return words, nil
}

// 从单词列表中随机选择一个单词
func randomWord(words []Word) Word {
	randIndex := rand.Intn(len(words))
	return words[randIndex]
}
func readConfig(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return err
	}
	return nil
}

// 发送邮件
func sendEmail(accuracy float64) {
	// 邮件服务器配置
	m := mail.NewMessage()
	m.SetHeader("From", config.Email)     // 发送者邮箱
	m.SetHeader("To", config.ParentEmail) // 接受者邮箱
	m.SetHeader("Subject", "单词学习工具 - 成绩")
	m.SetBody("text/plain", fmt.Sprintf("你的孩子在单词学习工具中达到了 %.2f%% 的正确率。太棒了！", accuracy))

	d := mail.NewDialer(config.SMTPServer, config.SMTPPort, config.Email, config.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("发送邮件时出错：", err)
	}
}
