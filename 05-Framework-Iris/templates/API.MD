# Sherpa API

The Sherpa API is for internal use at scalabilitysherpa.com. It allows polling of user
and lesson information in the context of the Scalability Sherpa course. This API is
the endpoint for all in app functionality within the course that interacts with the
Scalability Sherpa database.

Notes:
[x] get all users
- create a new user
- return user with list of lessons they have completed
- get all lessons
- get individual lesson with metadata and users who have completed
- mark lesson complete (post)
-
-

## Users Collection [/users]

### List All Users [GET]

Returns a list of all users.

+ Response 200 (application/json)

        {
          "users": [
            {
              "userid": 1,
              "first": "Quin",
              "last": "Kinser",
              "email": "quinkinser@gmail.com",
              "username": "quink",
              "lessonsCompleted": ["lid1", "lid2", "lid3"]

            },
            {
              "userid": 2,
              "first": "Wayne",
              "last": "Adams",
              "email": "quinkinser@gmail.com",
              "username": "wayney",
              "lessonsCompleted": ["lid1", "lid2"]
            }
             ,
            {
              "userid": 3,
              "first": "Jeremy",
              "last": "Toce",
              "email": "quinkinser@gmail.com",
              "username": "toasty",
              "lessonsCompleted": ["lid1"]
            }
          ]
        }



### Create OR Update a User [POST]

Create a new user using this action. This will happen at user sign up. It takes a
JSON object containing a first name, last name, email, username and password.

Send an array of users to be created.

+ Request (application/json)

        {
          "users": [
            {
              "userid": undefined,
              "first": "Quin",
              "last": "Kinser",
              "email": "quinkinser@gmail.com",
              "username": "quink",
              "lessonsCompleted": []

            },
            {
              "userid": 2,
              "first": "Wayne",
              "last": "Adams",
              "email": "quinkinser@gmail.com",
              "username": "wayney",
              "lessonsCompleted": []
            }
             ,
            {
              "userid": 3,
              "first": "Jeremy",
              "last": "Toce",
              "email": "quinkinser@gmail.com",
              "username": "toasty",
              "lessonsCompleted": []
            }
          ]
        }

+ Response 201 (application/json)

    + Headers



    + Body

            {
                "userid": 8,
                "url": "/users/8",
                "created_at": "2014-11-11T08:40:51.620Z"
            }

## User [/users/{userid}]

### Return Individual User [GET]

Returns a user with all information, including list of lessons they have completed. May need a session to complete
this request.

+ Request (application/json)

        {
            "userid": 1,
            "first": "Quin",
            "last": "Kinser",
            "email": "quinkinser@gmail.com",
            "username": "quink",
            "lessonsCompleted": ["lid1", "lid2", "lid3"]
        }

+ Response 201 (application/json)

    + Headers

            Location: /users/1

    + Body

            {
                "userid": 1,
                "first": "Quin",
                "last": "Kinser",
                "email": "quinkinser@gmail.com",
                "username": "quink",
                "lessonsCompleted": ["lid1", "lid2", "lid3"]
            }


### Mark a Lesson Completed [POST]

Adds a lesson ID to the user's completed lessons. In this example, lesson 8 is assed to the user's lessons

+ Request (application/json)

        {
            "userid": 1,
            "first": "Quin",
            "last": "Kinser",
            "email": "quinkinser@gmail.com",
            "username": "quink",
            "lessonsCompleted": ["lid1", "lid2", "lid3", "lid8"]
        }

+ Response 201 (application/json)

    + Headers

            Location: /users/1

    + Body

            {
                "userid": 1,
                "first": "Quin",
                "last": "Kinser",
                "email": "quinkinser@gmail.com",
                "username": "quink",
                "lessonsCompleted": ["lid1", "lid2", "lid3", "lid8"]
            }


