package auth

// func RegisterHandler(c *gin.Context) {
// var req RegisterRequest
// if err := c.ShouldBindJSON(&req); err != nil {
// 	common.ErrorValidation(c, err)
// 	return
// }

// // Validate unique email
// var count int64
// config.DB.Model(&user.User{}).Where("email = ?", req.Email).Count(&count)
// if count > 0 {
// 	common.ErrorResponse(c, []string{"Email already exists"}, 422)
// 	return
// }

// ip := c.GetHeader("X-Forwarded-For")
// newUser, errRegister := Register(req.Email, req.Password, req.Name, ip)
// if errRegister != nil {
// 	common.ErrorResponse(c, errRegister.Error(), 500)
// 	return
// }

// res := map[string]any{
// 	"id":    newUser.ID,
// 	"name":  newUser.Name,
// 	"email": newUser.Email,
// }
// common.SuccessResponse(c, res)
// }

// func LoginHandler(c *gin.Context) {
// 	var req LoginRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		common.ErrorValidation(c, err)
// 		return
// 	}

// 	token, errLogin := Login(req.Email, req.Password)
// 	if errLogin != nil {
// 		common.ErrorResponse(c, errLogin.Error(), 500)
// 		return
// 	}

// 	common.SuccessResponse(c, map[string]any{
// 		"token": token,
// 	})
// }

// func TestAja(c *gin.Context) {
// 	fmt.Println("TEST")
// 	c.JSON(200, gin.H{"test": "OK"})
// }

// func Hello(c *gin.Context) {
// 	c.JSON(200, gin.H{"message": "Hello, World!"})
// }
