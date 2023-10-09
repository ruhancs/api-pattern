package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/ruhancs/api-pattern/internal/dto"
	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/ruhancs/api-pattern/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserRepository database.UserRepositoryInterface
	Jwt            *jwtauth.JWTAuth
	JwtExpiresIn   int
}

func NewUserHandler(repository database.UserRepositoryInterface, jwt *jwtauth.JWTAuth, jwtExpiration int) *UserHandler {
	return &UserHandler{
		UserRepository: repository,
		Jwt:            jwt,
		JwtExpiresIn:   jwtExpiration,
	}
}

// GetJwt godoc
// @Summary      Get a user jwt
// @Description  Get a user jwt to authenticate
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.GetJwtDto  true  "user request"
// @Success      200  {object}  dto.GetJwtOutputDto
// @Failure      400  {object}  Error
// @Failure      403  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users/get_token [post]
func (handler *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	var userPayload dto.GetJwtDto
	err := json.NewDecoder(r.Body).Decode(&userPayload)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOut := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errOut)
		return
	}
	user, err := handler.UserRepository.FindByEmail(userPayload.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		errOut := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errOut)
		return
	}
	if !user.ValidatePassword(userPayload.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenMap := map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(handler.JwtExpiresIn)).Unix(),
	}
	_, tokenString, _ := handler.Jwt.Encode(tokenMap)

	accessToken := dto.GetJwtOutputDto{AccessToken: tokenString}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser godoc
// @Summary      Create user
// @Description  create user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      dto.UserCreateDto  true  "user request"
// @Success      201
// @Failure      400  {object}  Error
// @Failure      500  {object}  Error
// @Router       /users [post]
func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDto dto.UserCreateDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errOut := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(errOut)
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserRepository.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
