# go-flashtext
A basic implement for FlashText RESTful-API with Golang and Beego.

This service needs MySQL as its database, make sure your have set up that correctly and revise the ```main.go``` file's settings as well.

To launch this flash-text service, make sure you have already installed Beego(https://beego.me/) and run ```$ bee run``` in your terminal.


## Usage

```
# list all target words in database
http://localhost:8080/all

# add new target words into database
http://localhost:8080/addAll
{
  "word": ["This", "is", "a", "hello", "world", "message"]
}

# detect target words from an input list
http://localhost:8080/parse
{
  "word": ["Tony", "Stark", "says", "hello", "world"]
}
```

If you find this implement useful, please grant me a star and thank you in advance :)
