## RESTful API
#### Vote
POST `http(s)://<your likit ip>/api/v1/vote`

body

```
{
    "businessId": "businessId",
    "messageId": "messageId",
    "userId": "userId"
}
```

POST `http(s)://<your likit ip>/api/v1/unvote`

body

#### Unvote
```
{
    "businessId": "businessId",
    "messageId": "messageId",
    "userId": "userId"
}
```

#### Count the number of votes
GET `http(s)://<your likit ip>/api/v1/count/:businessId/:messageId`

#### List the users who voted
GET `http(s)://<your likit ip>/api/v1/list/:businessId/:messageId`

#### Check if the user has voted
GET `http(s)://<your likit ip>/api/v1/isVoted/:businessId/:messageId/:userId`

