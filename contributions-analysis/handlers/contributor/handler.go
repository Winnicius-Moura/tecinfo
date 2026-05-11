package contributor

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/wnn-dev/contributions-analysis/email"
	"github.com/wnn-dev/contributions-analysis/objects"
	"github.com/wnn-dev/contributions-analysis/password"
)

type Handler struct {
	contributorService  objects.ContributorService
	contributionService objects.ContributionService
	emailService        email.Service
	jwtSecret           string
}

// NewHandler is the Contributor handler constructor
func NewHandler(
	contributorService objects.ContributorService,
	contributionService objects.ContributionService,
	emailService email.Service,
	jwtSecret string,
) *Handler {
	return &Handler{
		contributorService:  contributorService,
		contributionService: contributionService,
		emailService:        emailService,
		jwtSecret:           jwtSecret,
	}
}

func generateRandomPassword() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "123456"
	}
	return hex.EncodeToString(bytes)
}

func (h *Handler) SignUp(ctx *gin.Context, input *objects.ContributorRegistrationVM) (*objects.Contributor, error) {
	if len(input.Email) < 11 || input.Email[len(input.Email)-11:] != "@evl.com.br" {
		return nil, errors.New("apenas e-mails institucionais (@evl.com.br) são permitidos")
	}

	plainPassword := generateRandomPassword()
	pHashed, err := password.NewHashedPassword(plainPassword)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error hashing password"))
	}

	contributor := &objects.Contributor{
		FullName: input.FullName,
		Email:    input.Email,
		Password: pHashed,
	}

	err = h.contributorService.Create(ctx, contributor)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error creating contributor"))
	}

	corpo := fmt.Sprintf("Olá,\n\nSua conta foi criada. Sua senha de acesso é: %s\n\nPor favor, faça login e altere sua senha.", plainPassword)
	go h.emailService.SendEmail(contributor.Email, "Acesso ao sistema", corpo)

	return contributor, nil
}

func (h *Handler) Login(ctx *gin.Context, input *objects.ContributorLoginVM) (*objects.Contributor, error) {
	contributor, err := h.contributorService.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributor by email"))
	}

	if !contributor.Password.IsEqualTo(input.Password) {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": contributor.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		return nil, errors.Wrap(err, "error generating token")
	}

	contributor.Token = tokenString

	return contributor, nil
}

func (h *Handler) Contributor(ctx *gin.Context, id string) (*objects.Contributor, error) {
	contributor, err := h.contributorService.Contributor(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributor"))
	}
	return contributor, nil
}

func (h *Handler) Contributors(ctx *gin.Context) ([]*objects.Contributor, error) {
	contributors, err := h.contributorService.Contributors(ctx)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("error fetching contributors"))
	}
	return contributors, nil
}
