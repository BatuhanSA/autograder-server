package core

import (
	"github.com/edulinq/autograder/internal/db"
	"github.com/edulinq/autograder/internal/model"
)

// Return a user only in the case that the authentication is successful.
// If any error is retuturned, then the request should end and the response sent based on the error.
// This assumes basic validation has already been done on the request.
func (this *APIRequestUserContext) Auth() (*model.ServerUser, *APIError) {
	if this.UserEmail == model.RootUserEmail {
		return nil, NewAuthBadRequestError("-051", this, "Root is not allowed to authenticate.")
	}

	user, err := db.GetServerUser(this.UserEmail)
	if err != nil {
		return nil, NewAuthBadRequestError("-012", this, "Cannot Get User").Err(err)
	}

	if user == nil {
		return nil, NewAuthBadRequestError("-013", this, "Unknown User")
	}

	auth, err := user.Auth(this.UserPass)
	if err != nil {
		return nil, NewBareInternalError("-037", this.Endpoint, "User auth failed.").Err(err)
	}

	if !auth {
		return nil, NewAuthBadRequestError("-014", this, "Bad Password")
	}

	return user, nil
}
