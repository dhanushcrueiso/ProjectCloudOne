# Run
1.Go to the root folder and => go run . to run the code

Testing through Postman :

Posting the url to get the shortned version of it:
1. http://localhost:3007/v1/shorten  (POST)
     Req body: {
              "url":"https://google.com"
                }

Posting the shortned version to get the original url:
2.http://localhost:3007/v1/getOriginal (POST)
      Req body:{
              "url": "http://shortn.ly/CXZHiCf"
              }


# Test 
1. To test the application with pre written unit test case , navigate to the test folder and
2. use : go test ; to run the test case automatically 
