package inter

/*
Set the following environment variables to configure the default credentials:

INTERBANK_CLIENT_ID
and
INTERBANK_CLIENT_SECRET

INTERBANK_SCOPES - must be a comma-separated list of scopes

INTERBANK_TLS_PATH - (required for mutual TLS)
must be a path to a file containing the certificate and key (tls.key and tls.crt)

```go
// Cria um cliente usando as variaveis de ambiente (INTERBANK_CLIENT_ID, INTERBANK_CLIENT_SECRET, INTERBANK_SCOPES e INTERBANK_TLS_PATH)
client, err := inter.NewClient()
// ...

// Use o ambiente sandbox (é o mesmo que client.SetURL("https://cdpj-sandbox.partners.uatinter.co"))
client.UseSandBox()

// Defina a conta corrente a ser usada (opcional)
client.SetAccountNumber("12345678")

// ConsultarSaldo
respSaldo, err := client.Banking.ConsultarSaldo(context.Background(), "2024-10-20")
// ...

log.Printf("Saldo: %+v", respSaldo.Disponivel)
```
*/
