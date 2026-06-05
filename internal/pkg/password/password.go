package password

import "golang.org/x/crypto/bcrypt"

func Hash(raw string) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(raw),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func Verify(
	hash string,
	raw string,
) bool {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(raw),
	)

	return err == nil
}
