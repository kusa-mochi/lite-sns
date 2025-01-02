package api_server

import (
	"fmt"
	auth_utils "lite-sns/m/src/cmd/app_server/api_server_common/auth"
	"lite-sns/m/src/cmd/app_server/commands"
	"log"
	"net/http"
	"net/mail"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

func (s *ApiServer) validateEmailAddress(addr string) error {
	_, err := mail.ParseAddress(addr)
	if err != nil {
		log.Println("invalid email address")
	}

	// TODO: ブラックリストの適用

	return err
}

func (s *ApiServer) validateNickname(name string) error {
	if name == "" {
		log.Println("invalid nickname")
		return fmt.Errorf("")
	}
	if len(name) > 20 {
		log.Println("too long nickname")
		return fmt.Errorf("")
	}

	allWhiteSpace := true
	for i, r := range name {
		// 先頭が空白の場合、エラー
		if i == 0 {
			if unicode.IsSpace(r) {
				log.Println("nickname starts from a space")
				return fmt.Errorf("")
			}
		}

		// 末尾が空白の場合、エラー
		if i == len(name)-1 {
			if unicode.IsSpace(r) {
				log.Println("nickname ends with a space")
				return fmt.Errorf("")
			}
		}

		if !unicode.IsSpace(r) {
			allWhiteSpace = false
		}
	}
	// すべて空白の場合、エラー
	if allWhiteSpace {
		log.Println("nickname has only white space characters")
		return fmt.Errorf("")
	}

	return nil
}

func (s *ApiServer) validatePassword(password string) error {
	if password == "" {
		log.Println("invalid password")
		return fmt.Errorf("")
	}
	if len(password) > 128 {
		log.Println("too long password")
		return fmt.Errorf("")
	}
	if len(password) < 12 {
		log.Println("too short password")
		return fmt.Errorf("")
	}
	for i := 0; i < len(password); i++ {
		if password[i] > unicode.MaxASCII {
			log.Println("password has no-ASCII character")
			return fmt.Errorf("")
		}
	}
	if strings.Contains(password, " ") {
		log.Println("password contains whitespace")
		return fmt.Errorf("")
	}

	return nil
}

// ユーザーアカウント仮登録処理
func (s *ApiServer) Signup(c *gin.Context) {
	log.Println("server signup start")

	var (
		emailAddr string = c.PostForm("EmailAddr")
		nickname  string = c.PostForm("Nickname")
		password  string = c.PostForm("Password")
	)

	log.Println("email addr:", emailAddr)
	log.Println("nickname:", nickname)
	log.Println("password:", password)

	// 受信したデータのバリデーション
	// スループット確保のためバリデーションのみこのゴルーチンで処理する。

	// eメールアドレス のバリデーション
	err := s.validateEmailAddress(emailAddr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "invalid signup data",
		})
		return
	}

	// ニックネーム のバリデーション
	err = s.validateNickname(nickname)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "invalid signup data",
		})
		return
	}

	// パスワード のバリデーション
	// パスワードハッシュではなくパスワードをバリデーションする。
	err = s.validatePassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": "invalid signup data",
		})
		return
	}

	// パスワードハッシュの取得
	passwordHash := auth_utils.GetHashStringFrom(password)

	resCh := make(chan string)
	s.commandCh <- &commands.SignupCommand{
		EmailAddr:    emailAddr,
		Nickname:     nickname,
		PasswordHash: passwordHash,
		ResCh:        resCh,
	}
	result := <-resCh

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
