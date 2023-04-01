# Workflow Agent

```markdown
name: OIDC Claims
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    permissions:
      id-token: write
    steps:
      - uses: actions/checkout@master
      - name: Self test
        id: selftest

        # Put your action repo here
        uses: lukehinds/workflow-agent@v0.0.8
```