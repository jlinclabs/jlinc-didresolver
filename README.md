# DID Resolver plugin

A DIF Universal Resolver plugin for the [JLINC DID spec](https://did-spec.jlinc.org/).
Consistent with the DIF Universal Resolver [API](https://github.com/decentralized-identity/universal-resolver/blob/master/swagger/api-driver.yml).

### Usage

```
cd dist/linux
./didresolver

```
Then GET http://localhost:8080/identifiers/{DID}
where {DID} is a valid JLINC DID.

For testing

```
http://localhost:8080/identifiers/did:jlinc:BkqgAVpbS_iJ6_tGXS8IrD0-yp5N24FZkq2ecR-vuDY

```
