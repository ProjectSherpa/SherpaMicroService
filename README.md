

#ScalibilitySherpa.com
##How to create a GO Microservice in 7 days 


![][mountainlogo.png]


##Problem
Scalabilitysherpa.com is expanding its footprint and they need a robust infrastructure. Currently the website is serving videos using a tightly coupled MERN stack. The company would like a more modern approach to persist data. 

![team][ScreenShot2016-08-15at2.05.14PM.png]

1.	Users- The website needs to track each individual users progress throughout the website, if they are paid or free content only. 

2.	Videos - Videos should contain a title, location of links , number of views.

3. Operational restrictions: The database needs to be secure, backups need to be provided as part of this service infrastructure change. 

4. Login should be handled by a third party oath 

![][ScreenShot2016-08-15at4.04.17PM.png]

##Current Infrastructure
Tightly coupled react framework with MYSQL backend. 

![][current.png]

##Proposed Solution
RESTfull Microservice

![future.png][future.png]

RESTful API ENDPOINTS
http://api.scalabilitysherpa.com/api/v1/

Resource , get, post, put, delete
/Users, returns list of all users, create new user, bulk update users, delete all
/Users/11/ , returns a user, 405 (not allowed), updates user, deletes user 
![][ScreenShot2016-08-15at5.25.35PM.png]
###Examples
####How to retrieve all users
1
2
3
45
CURL –X POST \
-H "Accept: application/json" \
-d '{"state":"running"}' \
https://api.scalabilitysherpa.com/api/v1/users



####How to retrieve a single user
1
2
3
45
CURL –X POST \
-H "Accept: application/json" \
-d '{"state":"running"}' \
https://api.scalabilitysherpa.com/api/v1/user/2863


####Partial Answers

Retrieve the first user with first name matching xyz 
1
2
3
4
5
6
7
8
GET /users?firstname=xyz
200 OK
{
"id":"007",
"Firstname”:”xyz”,
“LastName”:”Bond",
“AccountType”: “premium”, 
“Videos”:”[002,004,0783,0889]”
}

####Sorted Queries
Add sorted queries capability to the endpoint. 


####Searching
Add searching capability to the endpoint.



###Security 

	1.	 OAuth2 to secure your API endpoint. 



Notes
Maybe we can implement the following:
	•	API – https://api.{fakecompany}.com
	•	OAuth2 – https://oauth2.{fakecompany}.com
	•	Developer portal – https://developers.{fakecompany}.com

###RESTful API Resources
https://en.wikipedia.org/wiki/Representational_state_transfer
http://blog.mwaysolutions.com/2014/06/05/10-best-practices-for-better-restful-api/

http://blog.octo.com/en/design-a-rest-api/




##Schema Design


Users -
![][databaseplaceholder.png]
Videos -

###Database Backup 

-Automatic backups
-Send email to admin if there are database issues 

Scalability
	1.	Sharding
	2.	Globally accessible 
	3.	DNS
	4.	Monitoring

###Resources

Cloud Native Go: Building Web Applications and Microservices for the Cloud with Go and React

###Developer Portal
https://developers.scalabilitysherpa.com

Developer portal adds a location for developers share, update and document code, API end points , any idiomatic quirks etc.
![team][readmeimages/ScreenShot2016-08-15at4.01.56PM.png][mountainlogo.png]: readmeimages/mountainlogo.png
[team.png]: readmeimages/team.png
[current.png]: readmeimages/current.png
[future.png]: readmeimages/future.png 

[ScreenShot2016-08-15at2.05.14PM.png]: readmeimages/ScreenShot2016-08-15at2.05.14PM.png

[mountainlogo.png]: readmeimages/mountainlogo.png

[ScreenShot2016-08-15at4.04.17PM.png]: readmeimages/ScreenShot2016-08-15at4.04.17PM.png

[ScreenShot2016-08-15at5.23.12PM.png]: readmeimages/ScreenShot2016-08-15at5.23.12PM.png width=262px height=192px

[ScreenShot2016-08-15at5.28.40PM.png]: readmeimages/ScreenShot2016-08-15at5.28.40PM.png width=328px height=199px

[ScreenShot2016-08-15at5.25.35PM.png]: readmeimages/ScreenShot2016-08-15at5.25.35PM.png

[redislogo.png]: readmeimages/redislogo.png width=226px height=71px

[databaseplaceholder.png]: readmeimages/databaseplaceholder.png width=344px height=246px

[ScreenShot2016-08-15at4.01.56PM.png]: readmeimages/ScreenShot2016-08-15at4.01.56PM.png width=1494px height=993px