package handler

import (
	"fmt"
	"net/http"

	jsonwebtoken "github.com/yihune21/link-shortner/internal/jsonWebToken"
	"github.com/yihune21/link-shortner/internal/service"
)

type AuthHandler struct{
	as *service.AuthService
}

func NewAuthHandler(as *service.AuthService) *AuthHandler {
	return &AuthHandler{as:as}
}

func (a *AuthHandler)MiddlewareAuth(w http.ResponseWriter , r *http.Request) {
          access_token ,  err := a.as.GetToken(r.Header)
		  if err != nil{
			WriteError(w , 401 , fmt.Sprintf("Auth Error %s" , err))
			return
		  }
		  is_valid := jsonwebtoken.VerifyToken(access_token)
          if !is_valid{
            WriteError(w , 401 ,"ACCESS TOKEN EXPIRED!")
			return 
		  }
		  
		//   _ , err = apiConf.db.GetToken(r.Context(),access_token)

		//   if err == nil{
		// 	respondWithError(w , 400 , fmt.Sprintf("Token is blacklisted.%v",err))
		// 	return
		//   }

		//   user_id,err := jwtAuth.ExtractUserIDFromToken(access_token)
		//   if err != nil{
		// 	respondWithError(w , 400 , fmt.Sprintf("Error with extracting user id %v",err))
		// 	return
		//   }

		//   user , err := apiConf.db.GetUserById(r.Context() ,user_id)
        //   if err != nil{
		// 	 respondWithError(w , 404 , fmt.Sprintf("Couldn't found user %s" , err))
		// 	 return
		//   }
		//   handler(w, r , user)
	

}