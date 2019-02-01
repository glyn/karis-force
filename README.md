# Karis Kids Saleforce Application

Under development - do not use.

## Authentication

The following environment variables need to be set:
```terminal
export KARIS_FORCE_CLIENT_ID="xxx"
export KARIS_FORCE_CLIENT_SECRET="xxx"
export KARIS_FORCE_USER_NAME="user.name@kariskids.org"
export KARIS_FORCE_PASSWORD="xxx"
export KARIS_FORCE_SECURITY_TOKEN="xxx"
```

Notes:
* The client id and client secret are the consumer key and consumer secret, respectively, for the connected app
"karisforce". See [this description](https://developer.salesforce.com/forums/?id=906F0000000AfcgIAC) of how to obtain those values.
* The user name and password are your usual login details for Karis Kids Saleforce.
You must have karisforce assigned as a connected app.
* The security token is emailed to you when you [reset your security token](https://help.salesforce.com/articleView?id=user_security_token.htm&type=5).