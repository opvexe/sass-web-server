package tools

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"pea-web/cmd"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

const (
	FMT_DATE_TIME    = "2006-01-02 15:04:05"
	FMT_DATE         = "2006-01-02"
	FMT_TIME         = "15:04:05"
	FMT_DATE_TIME_CN = "2006年01月02日 15时04分05秒"
	FMT_DATE_CN      = "2006年01月02日"
	FMT_TIME_CN      = "15时04分05秒"
)

// 是否生产环境
func IsProd() bool {
	return cmd.Conf.DevEnv == "prod"
}

//判断用户名的合法性
func IsValidateUsername(username string) error {
	if len(username) == 0 {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

// 验证是否是合法的邮箱
func IsValidateEmail(email string) (err error) {
	if len(email) == 0 {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return
}

// 是否是合法的密码
func IsValidatePassword(password, rePassword string) error {
	if len(password) == 0 {
		return errors.New("请输入密码")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}

	return nil
}

// MD5
func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// 判断是否为空
func IsEmpty(a interface{}) bool {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v.Interface() == reflect.Zero(v.Type())
}

/**
 * 将时间格式换成 xx秒前，xx分钟前...
 * 规则：
 * 59秒--->刚刚
 * 1-59分钟--->x分钟前（23分钟前）
 * 1-24小时--->x小时前（5小时前）
 * 昨天--->昨天 hh:mm（昨天 16:15）
 * 前天--->前天 hh:mm（前天 16:15）
 * 前天以后--->mm-dd（2月18日）
 */
func PrettyTime(timestamp int64) string {
	_time := TimeFromTimestamp(timestamp)
	_duration := (NowTimestamp() - timestamp) / 1000
	if _duration < 60 {
		return "刚刚"
	} else if _duration < 3600 {
		return strconv.FormatInt(_duration/60, 10) + "分钟前"
	} else if _duration < 86400 {
		return strconv.FormatInt(_duration/3600, 10) + "小时前"
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24))) <= timestamp {
		return "昨天 " + TimeFormat(_time, FMT_TIME)
	} else if Timestamp(WithTimeAsStartOfDay(time.Now().Add(-time.Hour*24*2))) <= timestamp {
		return "前天 " + TimeFormat(_time, FMT_TIME)
	} else {
		return TimeFormat(_time, FMT_DATE)
	}
}

func Timestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// 毫秒时间戳转时间
func TimeFromTimestamp(timestamp int64) time.Time {
	return time.Unix(0, timestamp*int64(time.Millisecond))
}

// 毫秒时间戳
func NowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

// 返回指定时间当天的开始时间
func WithTimeAsStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// 时间格式化
func TimeFormat(time time.Time, layout string) string {
	return time.Format(layout)
}

// 秒时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}
