# Mailchimp to Google Form

When someone subscribes to a list in Mailchimp, post their email to a Google Form.

Tiny server which receives Mailchimp webhooks and posts to Google.

Configure by setting these environment variables:

```
export MAILCHIMP_SECRET="[set this to a long random string, e.g. a UUID]"
export GOOGLE_FORM_ACTION_URL="[form action= url]"
export GOOGLE_FORM_EMAIL_FIELD="[form email field name e.g. entry.123456]"
```

Then configure your Mailchimp webhook with the `MAILCHIMP_SECRET` you set above, for example:

```
https://your-new-heroku-app.herokuapp.com/the-long-secret-you-set
```

Run locally with `make run` then test with curl:

```
curl -v -X POST -d "data[email]=test@example.com" -d "type=subscribe" localhost:4747/fake_secret
```
