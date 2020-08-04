# nest-secrets

Retrieve configuration from AWS Parameter Store by prefix. For example, if you
app is called 'pineapple' and you have parameters as so:

    /pineapple/prod/db.url: http://..
    /pineapple/prod/db.password: secret-password

Then run:

    $ nest-secrets --prefix /pineapple/prod

And, assuming you have AWS credentials, you'll get back:

    pineapple_prod_db_url=http://..
    pineapple_prod_db_password=secret-password

Note, '/' and '.' are converted to '\_' for the environment variable names, and
the leading '/' is removed.

If using with Docker, the recommendation is to use this with the
[env-file](https://docs.docker.com/compose/env-file/) option, something like:

    $ nest-secrets --prefix /pineapple/prod > .env
    $ docker run my-container --env-file .env
