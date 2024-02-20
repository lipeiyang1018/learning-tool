# learning-tool

> 程序用途和用法： 
1. 本程序是用golang 语言编写
2. 可以辅助学习英文单词和语文拼音 （可以延伸其他学科）
3. config.json 修改配置文件
4. words.json 添加学习的单词量
5. 有条件的同学们也可以用的程序基础上修改扩展
6. https://github.com/lipeiyang1018/learning-tool 本人代码开源仓库



> 配置文件 config.json

- "email": "your_email@example.com",  // 你的邮箱
- "password": "your_password",     // 你的邮箱密码
- "parent_email": "parent_email@example.com",  // 你的家长邮箱
- "teacher_email": "teacher_email@example.com", //  你老师的邮箱
- "smtp_server": "smtp.example.com", // 邮箱服务器地址 示例  "smtp-mail.outlook.com",
- "smtp_port": 587,   // 邮箱服务器端口   一般不用改
- "total_questions": 100,    // 每次做题的数量
- "passing_score": 95  // 分数线



程序打包
``` 
Mac 下编译 Linux 和 Windows 64位可执行程序
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build   
   CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

