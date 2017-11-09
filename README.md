## Haiku Donation Tracker

This is a small project used to automate the tracking of donations to [Haiku](https://www.haiku-os.org).
Currently donations are primarily received from PayPal, with occasional checks or larger donations
made directly to the [Haiku, Inc.](http://www.haiku-inc.org) bank account.
A golang based donation analysis and reporting tool.

## Compiling

```shell
$ go build .
```

## config.json

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

## To Run

```
./donation_tracker
```

This requires PayPal API credentials in config.json as described above.

## Code Organization

* `config.go`: Contains code for loading a simple `config.json` file containing PayPal API credentials and other config information.
* `currency.go`: Contains the code for getting the EUR to USD conversion rate from "fixer.io".
* `main.go`: Contains the main function which orchestrates the overall process described above.
* `nvp.go`: Provides support for the unique "Name Value Pair" (NVP) format returned from PayPal API calls.
* `paypal.go`: Provides all the code for handling, sorting, filtering and summarizing the PayPal transactions.
* `paypal_api.go`: The code for calling the PayPal API which makes use of the NVP code.
* `summary.go`: All the code for overall summaries.
* `transactions.go`: Probably should be renamed. Loads the bank transactions from the CSV file described above.
