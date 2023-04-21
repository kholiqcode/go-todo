package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/kholiqcode/go-todolist/utils"
	"golang.org/x/crypto/chacha20poly1305"
)

type PasetoMaker struct {
	symmetricKey paseto.V4SymmetricKey
}

func NewPasetoMaker(config *utils.BaseConfig) (Maker, error) {
	key := config.TokenSymmetricKey
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	symmetricKey, err := paseto.V4SymmetricKeyFromBytes([]byte(key))

	if err != nil {
		return nil, fmt.Errorf("invalid key: %w", err)
	}

	maker := &PasetoMaker{
		symmetricKey: symmetricKey,
	}

	return maker, nil
}

func (m *PasetoMaker) CreateToken(params PayloadParams, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(params, duration)

	mapPayload := map[string]interface{}{}

	err := utils.ConvertInterfaceE(payload, &mapPayload)

	if err != nil {
		return "", nil, fmt.Errorf("error convert payload to map: %w", err)
	}

	token, err := paseto.MakeToken(mapPayload, nil)
	if err != nil {
		return "", nil, fmt.Errorf("error make token: %w", err)
	}
	tokenEncrypt := token.V4Encrypt(m.symmetricKey, nil)

	return tokenEncrypt, payload, err

}

func (m *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	parser := paseto.NewParserWithoutExpiryCheck()

	verifiedToken, err := parser.ParseV4Local(m.symmetricKey, token, nil)

	if err != nil {
		return nil, ErrInvalidToken
	}

	err = utils.ConvertInterfaceE(verifiedToken.Claims(), payload)

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
