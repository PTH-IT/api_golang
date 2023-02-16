package utils

import (
	"github.com/simia-tech/crypt"
)

func CryptPassword(password string) *string {
	settings := "$6$saltstring$svn8UoSVapNtMuq1ukKS4tPQd8iKwSMHWjl/O817G3uBnIFNjnQJuesI68u4OTLiBFdcbYEdFCoEOfaS35inz1" // salt = "saltsalt"

	encoded, err := crypt.Crypt(password, settings)
	if err != nil {
		return nil
	}
	return &encoded
}
