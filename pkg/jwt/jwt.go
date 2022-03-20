package jwt

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"doctorx/pkg/db"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//VerifyToken
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the apikey from header
		token := c.Request.Header.Get("apikey")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":    "Access Denied - API Key is required!",
				"data":   nil,
			})
			c.Abort()
			return
		}

		if !IsValidToken(token) {
			c.JSON(http.StatusUnauthorized, gin.H {
				"status": 401,
				"msg": "Access Denied - the token is invalid!",
				"data": nil,
			})
			c.Abort()
			return
		}

		log.Print("get apikey: ", token)
		c.Next()

	}
}

func IsValidToken(token string) bool {
	// reserved super token
	if token == "ZG9jdG9yeC1zdXBlcnRva2VuCg==" {
		return true
	}

	// verify the token in db
	conn := db.GetSQLite3Connection()
	var count int
	statement := "select count(1) from token where token='" + token + "'"
	//log.Print("statement: ", statement)
	rs, err := db.RunQuery(conn, statement)
	CheckError(err)
	for rs.Next() {
		err := rs.Scan(&count)
		CheckError(err)
	}

	//close connection
	rs.Close()
	conn.Close()
	return count != 0
}


//CORS Management
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}