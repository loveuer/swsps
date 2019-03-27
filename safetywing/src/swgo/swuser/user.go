package swuser

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

// Get user imformation by username
// return user.User and an error
func Get(uname string) (User, error) {
	sentence := fmt.Sprintf("select id, username, realname, profile_icon, last_login from sw_user where username = '%s'", uname)
	row := conn.QueryRow(sentence)
	var u User
	err = row.Scan(&u.ID, &u.UserName, &u.RealName, &u.ProfileIcon, &u.LastLogin)
	if err != nil {
		log.Println("  * get user by id scan error")
		return u, err
	}
	return u, nil
}

// New generate user
// return an error
func New(newu User, password string) error {
	encedPassword := genPassword(password)
	sentence := fmt.Sprintf("insert into sw_user(username,realname,profile_icon,last_login,password) values ('%s','%s','%s','%s','%s')", newu.UserName, newu.RealName, newu.ProfileIcon, newu.LastLogin.String(), encedPassword)
	_, err = conn.Exec(sentence)
	return err
}

// Check username and password while login
// return a bool
func Check(username, password string) bool {
	// 通过username拿到password
	dbPassword := getUserPasswordByUserName(username)
	result := comparePassword(password, dbPassword)
	if result {
		go updateUserLastLogin(username)
	}
	return result
}

func getUserPasswordByUserName(username string) string {
	sentence := fmt.Sprintf("select password from sw_user where username='%s'", username)
	row := conn.QueryRow(sentence)
	var dbPassword string
	err = row.Scan(&dbPassword)
	if err != nil {
		log.Println("  * scan db password error")
		return ""
	}
	return dbPassword
}

func comparePassword(inp, dbp string) bool {
	dbpSplit := strings.Split(dbp, "$")
	if len(dbpSplit) != 4 {
		log.Println("  * db password can't split to 4 parts")
		return false
	}

	iterations, err := strconv.Atoi(dbpSplit[1])
	if err != nil {
		log.Println("  * db password trans iteration to int error")
		return false
	}

	salt := dbpSplit[2]
	if encryptPassword(inp, salt, iterations) == dbp {
		return true
	}
	return false
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
func encryptPassword(pswd, salt string, iterations int) string {
	pwd := []byte(pswd)    // 用户设置的原始密码
	saltbs := []byte(salt) // 盐，是一个长12的随机字符串
	digest := sha256.New   // digest 算法，使用 sha256

	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key(pwd, saltbs, iterations, 32, digest)

	// 第二步：Base64 编码
	str := base64.StdEncoding.EncodeToString(dk)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	// 因为是从django迁移数据过来
	encPs := "pbkdf2_sha256" + "$" + strconv.FormatInt(int64(iterations), 10) + "$" + salt + "$" + str
	return encPs
}

func updateUserLastLogin(username string) {
	t := time.Now()
	timeStr := t.Format(baseTime)
	sentence := fmt.Sprintf("update sw_user set last_login = '%s' where username = '%s'", timeStr, username)
	_, err = conn.Exec(sentence)
	if err != nil {
		log.Println("  * update user last_login time execute error")
	}
}
