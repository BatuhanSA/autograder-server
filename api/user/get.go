package user

import (
    "github.com/eriq-augustine/autograder/api/core"
)

type UserGetRequest struct {
    core.APIRequestCourseUserContext
    core.MinRoleGrader
    Users core.CourseUsers `json:"-"`

    TargetUser core.TargetUser `json:"target-email"`
}

type UserGetResponse struct {
    FoundUser bool `json:"found-user"`
    User *core.UserInfo `json:"user"`
}

func HandleUserGet(request *UserGetRequest) (*UserGetResponse, *core.APIError) {
    response := UserGetResponse{};

    if (!request.TargetUser.Found) {
        return &response, nil;
    }

    response.FoundUser = true;
    response.User = core.NewUserInfo(request.TargetUser.User);

    return &response, nil;
}
