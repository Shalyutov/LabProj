package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
	orm "labproj/ydb"
	"net/http"
	"slices"
	"strings"
)

func Authorize(scopesRequired []string, db *orm.Orm) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("o9384u98vr8nfy93e8ur034u03h9458uy0469h56y0n9i6tpv394omd28d3y4rv9873b456b"), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}), jwt.WithIssuer("AuthLIS"),
			jwt.WithAudience("LIS"), jwt.WithIssuedAt(), jwt.WithExpirationRequired())

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userClaim, err := claims.GetSubject()
			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			q := `
			DECLARE $user AS Utf8;
			SELECT
				full_name, is_blocked
			FROM 
				users
			WHERE 
				username = $user;
			`
			params := query.WithParameters(
				ydb.ParamsBuilder().
					Param("$user").Text(userClaim).
					Build(),
			)
			users, err := orm.Query[orm.User](db, q, params)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if len(users) == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			if users[0].IsBlocked != nil && *users[0].IsBlocked == true {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			q = `
			DECLARE $user AS Utf8;
			SELECT
				scope
			FROM 
				user_scopes
			WHERE 
				username = $user;
			`
			params = query.WithParameters(
				ydb.ParamsBuilder().
					Param("$user").Text(userClaim).
					Build(),
			)
			scopes, err := orm.Query[orm.UserScope](db, q, params)
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			if len(scopes) == 0 {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			for _, scope := range scopesRequired {
				if !slices.Contains(scopes, orm.UserScope{Scope: scope}) {
					c.AbortWithStatus(http.StatusUnauthorized)
					return
				}
			}
			c.Set("username", userClaim)
			c.Set("FullName", users[0].FullName)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
