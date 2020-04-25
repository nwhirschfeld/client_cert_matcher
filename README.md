# client_cert_matcher
caddy v2 http matcher, which checks CN of client certificate against list

## build caddy with the plugin
The best way to build caddy with the plugin is using [xcaddy](https://github.com/caddyserver/xcaddy) 
```
# xcaddy build --with github.com/nwhirschfeld/client_cert_matcher 
```

## usage
Currently Caddy v2 does not yet allow to require client certificates using the Caddyfile. As it is anyway nessecary to configure client certificates over the JSON interface, no Caddyfile parser is included in the plugin yet.

To use the plugin it is nessecary to enable client authentication. The following snippet shows a sample configuration:

```json
{
  "apps": {
    "http": {
      ...
      "servers": {
        "srv0": {
          ...
          "tls_connection_policies": [
            {
              ...
              "client_authentication": {
                "trusted_ca_certs": [
                  "Q0VSVElGSUNBVEU8Mwo="
                ],
                "mode":"require_and_verify"
              }
            }
          ]
        }
      }
    },
    ...
  },
  ...
}
```

now you can include CN names you want to allow in the match ruleset of your routes

```json
{
  "client_cert": [
    "alice",
    "bob",
    "eve"
  ],
  ...
}
```
