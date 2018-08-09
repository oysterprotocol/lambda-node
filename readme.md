# Lambda Node

The brokernode leverages [AWS Lambda](https://aws.amazon.com/lambda/) to offload PoW.

# Setup

The lambda node uses the {serverless framework](https://serverless.com/framework/docs/providers/aws/guide/quick-start/).

To install the command line tools, run:

```bash
npm install -g serverless
```

You will need to setup AWS credentials in order to deploy or run functions:

_NOTE: Change the key and secret key to the values stored in the oyster 1 Password under `serverless-login`_

```
serverless config credentials --provider aws --key <AWS_KEY> --secret <AWS_SECRET_KEY>
```

# Deploy

```bash
make deploy
```
