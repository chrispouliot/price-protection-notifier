## Price Protection Notifier

Golang lambda function that runs daily, checks a list of product pages, and notifies you if the price you set it to watch for is reached.

Built with Golang, MongoDB, AWS Lambda

### Usage

Upload the `main.zip` to AWS Lambda and set it on a daily trigger (You can use Cloudwatch events for this)

I use it with the ESP Mailgun, but it should be usable with any email provider of your choice.

Necessary environment variables
```bash
MONGO_URL=mongodb://...
MONGO_DB=...
MONGO_COLLECTION=...
MAIL_EMAIL_ADDRESS=me@example.com
MAIL_API_URL=https://api.mailgun.net/v3/...
MAIL_API_KEY=...
MAIL_FROM_ADDRESS=...
```

### Interacting with the products database

The following commands let you interact with the database:
 - remove
 - insert
 - list
 
 They all require the MongoDB environment variables listed above

