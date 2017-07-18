
## Haiku Donation Tracker

A golang based donation analysis and reporting tool.

### Compiling

```shell
$ go build .
```

### config.json

A configuration file needs to be placed in the working directory of the tool to obtain proper PayPal auth.

```json
{
    "ENDPOINT":"https://api.sandbox.paypal.com",
    "USER":"A_TEST_USER",
    "PWD":"A_TEST_CREDENTIAL",
    "SIGNATURE":"A_SIGNATURE_CREDENTIAL",
    "METHOD":"TransactionSearch"
}
```
