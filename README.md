# Basic Go Bank API

A simple API reading rate from the EuroBank server.

## Running locally

```
git clone https://github.com/openwt/bankapi
PORT=6060 ./euroconv
```

Your app should now be running on [localhost:6060](http://localhost:6060/).

## Deploying to Heroku

```
heroku create
git push heroku master
heroku open
```

Alternatively, you can deploy your own copy of the app using the web-based flow:

[![Deploy to Heroku](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)
