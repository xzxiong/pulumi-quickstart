# pulumi-quickstart


# testcase

## get output

- dev: run stackRefGetOutput, result: ok
- qa: run stackRefGetStringOutput: result: failed:
    - `error: an unhandled error occurred: program failed:`
    - `waiting for RPCs: stack reference output "not-exist-key" does not exist on stack "xzxiong/quickstart/dev"`
