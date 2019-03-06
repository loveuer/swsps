package user

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// Get user imformation by uid
func Get(uid string) (User, error) {
	sentence := fmt.Sprintf("select id, username, realname, profile_icon, last_login from swuser where id = %s", uid)
	row := conn.QueryRow(sentence)
	var u User
	err = row.Scan(&u.ID, &u.UserName, &u.RealName, &u.ProfileIcon, &u.LastLogin)
	if err != nil {
		return u, err
	}
	return u, nil
}

// New generate user
func New(newu User, password string) error {
	encedPassword := genPassword(password)
	sentence := fmt.Sprintf("insert into swuser(username,realname,profile_icon,last_login,password) values ('%s','%s','%s','%s','%s')", newu.UserName, newu.RealName, newu.ProfileIcon, newu.LastLogin.String(), encedPassword)
	_, err = conn.Exec(sentence)
	return err
}

func Check(username, password string) (User, error) {

}

// generate new password
func genPassword(pswd string) string {
	rand.Seed(time.Now().UnixNano())
	salt := func(n int) string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letterRunes[rand.Intn(len(letterRunes))]
		}
		return string(b)
	}(12)
	return encryptPassword(pswd, salt, 100000)
}

// iter 100000, len(salt)=12
func encryptPassword(pswd, salt string, iterations int64) string {
	pwd := []byte(pswd)    // 用户设置的原始密码
	saltbs := []byte(salt) // 盐，是一个长12的随机字符串
	digest := sha256.New   // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwd, saltbs, int(iterations), 32, digest)

	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	// 因为是从django迁移数据过来
	encPs := "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + string(salt) + "$" + str
	return encPs
}
