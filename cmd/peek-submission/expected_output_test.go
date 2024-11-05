package main

const SUBMISSION_1697406272 = `{
    "found-user": true,
    "found-submission": true,
    "submission-result": {
        "id": "course101::hw0::course-student@test.edulinq.org::1697406272",
        "short-id": "1697406272",
        "course-id": "course101",
        "assignment-id": "hw0",
        "user": "course-student@test.edulinq.org",
        "message": "",
        "max_points": 2,
        "score": 2,
        "name": "HW0",
        "questions": [
            {
                "name": "Q1",
                "max_points": 1,
                "score": 1,
                "message": "",
                "grading_start_time": 1697406273000,
                "grading_end_time": 1697406273000
            },
            {
                "name": "Q2",
                "max_points": 1,
                "score": 1,
                "message": "",
                "grading_start_time": 1697406273000,
                "grading_end_time": 1697406273000
            },
            {
                "name": "Style",
                "max_points": 0,
                "score": 0,
                "message": "Style is clean!",
                "grading_start_time": 1697406273000,
                "grading_end_time": 1697406273000
            }
        ],
        "grading_start_time": 1697406273000,
        "grading_end_time": 1697406273000,
        "additional-info": null
    }
}
`

const NO_SUBMISSION = `{
    "found-user": true,
    "found-submission": false,
    "submission-result": null
}
`

const INCORRECT_SUBMISSION = `{
    "found-user": true,
    "found-submission": false,
    "submission-result": null
}
`
