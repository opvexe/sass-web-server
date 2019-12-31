package tools

import (
	"bytes"
	"crypto/tls"
	"errors"
	"github.com/jordan-wright/email"
	"html/template"
	"net"
	"net/smtp"
	"pea-web/cmd"
)

var emailTemplate = `
<div style="background-color:white;border-top:2px solid #12ADDB;box-shadow:0 1px 3px #AAAAAA;line-height:180%;padding:0 15px 12px;width:500px;margin:50px auto;color:#555555;font-family:'Century Gothic','Trebuchet MS','Hiragino Sans GB',微软雅黑,'Microsoft Yahei',Tahoma,Helvetica,Arial,'SimSun',sans-serif;font-size:12px;">
    <h2 style="border-bottom:1px solid #DDD;font-size:14px;font-weight:normal;padding:13px 0 10px 8px;">
        <span style="color: #12ADDB;font-weight:bold;">
            {{.Title}}
        </span>
    </h2>
    <div style="padding:0 12px 0 12px; margin-top:18px;">
        {{if .Content}}
		<p>
            {{.Content}}
        </p>
		{{end}}
		{{if .QuoteContent}}
		<div style="background-color: #f5f5f5;padding: 10px 15px;margin:18px 0;word-wrap:break-word;">
            {{.QuoteContent}}
        </div>
		{{end}}
       
		{{if .Url}}
        <p>
            <a style="text-decoration:none; color:#12addb" href="{{.Url}}" target="_blank" rel="noopener">点击查看详情</a>
        </p>
		{{end}}
    </div>
</div>
`

// 发送邮箱
func SendTemplateEmail(to, subject, title, content, quoteContent, url string) error {
	tpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return err
	}
	var buff bytes.Buffer
	err = tpl.Execute(&buff, map[string]interface{}{
		"Title":        title,
		"Content":      content,
		"QuoteContent": quoteContent,
		"Url":          url,
	})
	if err != nil {
		return err
	}
	html := buff.String()
	if err := SendEmail(to, subject, html); err != nil {
		return err
	}
	return nil
}

//发送邮箱
func SendEmail(to, subject, html string) error {
	var (
		host      = cmd.Conf.SMTP.Host
		port      = cmd.Conf.SMTP.Port
		username  = cmd.Conf.SMTP.Username
		password  = cmd.Conf.SMTP.Password
		ssl       = true
		addr      = net.JoinHostPort(host, port)
		auth      = smtp.PlainAuth("", username, password, host)
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}
	)
	e := email.NewEmail()
	e.From = cmd.Conf.SMTP.Username
	e.To = []string{to}
	e.Subject = subject
	e.HTML = []byte(html)

	if ssl {
		if err := e.SendWithTLS(addr, auth, tlsConfig); err != nil {
			return errors.New("发送邮箱异常")
		}
	} else {
		if err := e.Send(addr, auth); err != nil {
			return errors.New("发送邮件异常")
		}
	}
	return nil
}
