# Authenticator bundle

This bundle is used to authenticate users on your API.
It uses [Json Web Token](https://en.wikipedia.org/wiki/JSON_Web_Token).

## Setup

In order to work, the bundle needs the following variables to be set in your `.env` file :

| variable             | description                            |
| -------------------- | -------------------------------------- |
| TOKEN_VALID_DURATION | valid duration of the token in minutes |
| RSA_PUBLIC_PATH      | the path to the public.pem file        |
| RSA_PRIVATE_PATH     | the path to the private.pem file       |
| RSA_PASSWORD         | password of the encryption key         |
| DOMAIN               | the domain for the cookies             |

As the bundle uses RSA keys, you need to generate them :

```sh
# Generate private.pem
openssl genrsa -des3 -out private.pem 2048

# Generate public.pem
openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```

> 💡 You need to use the same password for both commands, this password is the one you set as **RSA_PASSWORD**

By default if you put your public.pem and private.pem in project root, then the path to them should be `../public.pem`.

## Available methods

The following methods are available :

| name          | description                                      |
| ------------- | ------------------------------------------------ |
| GenerateToken | generates a JWT from email                       |
| DecodeToken   | decode a jwt and returns the token and its claim |
| HashPassword  | hash password with MD5 method                    |

## Use

When installed, a middleware is automatically created for authentication.
You can use it inside the `api/routes/router.go` file :

```go
func Init(r *gin.Engine) {
  r.POST("/register", controllers.CreateCustomer)
  r.POST("/login", controllers.Login)

  api := r.Group("/api")
  api.Use(middlewares.CheckAuthorization)
  {
    api.GET("/", controllers.TestController)
  }
}
```

`login controller` is also created, feel free to modify it to your needs.

Your customer entity should have a password, you can use the authenticator.HashPassword to do so.
