package email

import (
	"errors"
	"net/smtp"
	"regexp"

	"github.com/jordan-wright/email"
)

// content 使用html格式
func SendMail(Toemail, title, conten_html string) error {
	smtpHost := "smtp.qq.com" // SMTP服务器地址
	smtpPort := "587"         // SMTP服务器端口
	toUserEmail := Toemail    // 接收者邮箱地址
	if !isEmailLegal(toUserEmail) {
		return errors.New("邮箱格式错误")
	}
	e := email.NewEmail()
	e.From = "shemao1114@qq.com"                                                                                // 发件人邮箱账号
	e.To = append(e.To, toUserEmail)                                                                            // 收件人邮箱地址                                                                               // 收件人邮箱地址
	e.Subject = title                                                                                           // 邮件主题
	e.Text = []byte("未处理的请求")                                                                                   // 邮件正文内容（纯文本）
	e.HTML = []byte(conten_html)                                                                                // 邮件正文内容（HTML格式）
	err := e.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", "shemao1114@qq.com", "rxduvmexcqsfbfdf", smtpHost)) // 发送邮件
	if err != nil {
		return err
	}
	return nil
}
func isEmailLegal(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	// 使用MatchString()函数来判断电子邮件地址是否匹配正则表达式
	return emailRegex.MatchString(email)
}
