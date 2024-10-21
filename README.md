# Interbank-Go

A biblioteca `interbank-go` é uma implementação em Go para interagir com a API do Banco Inter. Ela fornece funcionalidades para autenticação, operações bancárias, cobranças e transações PIX.

## Como usar

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

## Estrutura do Projeto

- **auth**: Gerencia a autenticação e autorização.
- **backend**: Lida com as requisições HTTP e gerenciamento de tokens.
- **banking**: Fornece serviços relacionados a operações bancárias.
- **cobranca**: Gerencia a emissão e consulta de cobranças.
- **pix**: Implementa funcionalidades relacionadas ao sistema PIX.
- **erros**: Define estruturas para tratamento de erros.

## Autenticação

### auth/credentials.go

Gerencia as credenciais necessárias para autenticação na API do Banco Inter. As credenciais incluem `clientID`, `clientSecret`, `grantType` e `scopes`.

#### Variáveis de Ambiente

Para configurar as credenciais padrão, defina as seguintes variáveis de ambiente:

- `INTERBANK_CLIENT_ID`: O ID do cliente obtido no detalhe da tela de aplicações no IB.
- `INTERBANK_CLIENT_SECRET`: O segredo do cliente obtido no detalhe da tela de aplicações no IB.
- `INTERBANK_SCOPES`: Deve ser uma lista separada por vírgulas dos escopos desejados.
- `INTERBANK_TLS_PATH`: Caminho para um arquivo contendo o certificado e a chave (tls.key e tls.crt).

#### Funções Principais

- `NewCredentials`: Cria novas credenciais.
- `NewDefaultCredentials`: Carrega credenciais padrão a partir de variáveis de ambiente.
- `SetTLS`: Configura o certificado TLS.

### auth/token.go

Define a estrutura do token de acesso e métodos para verificar sua validade.

#### Funções Principais

- `Valid`: Verifica se o token ainda é válido.
- `GetAccessToken`: Retorna o token de acesso.

## Serviços Bancários

### banking/banking.go

Fornece métodos para consultar extratos e saldos.

#### Funções Principais

- `ExportarExtrato`: Exporta o extrato da conta.
- `ConsultarExtratoCompleto`: Consulta o extrato completo da conta.
- `ConsultarSaldo`: Consulta o saldo da conta.

## Cobrança

### cobranca/cobranca.go

Gerencia a emissão de cobranças.

#### Funções Principais

- `EmitirCobranca`: Emite uma nova cobrança.

## PIX

### pix/pix.go

Implementa funcionalidades para transações PIX, incluindo consultas e devoluções.

#### Funções Principais

- `ConsultarDevolucao`: Consulta a devolução de um PIX.
- `SolicitarDevolucao`: Solicita a devolução de um PIX.
- `ConsultarRecebidos`: Consulta PIX recebidos.

### pix/pix_cob_imediata.go

Gerencia cobranças imediatas via PIX.

#### Funções Principais

- `CriarCobrancaImediata`: Cria uma cobrança imediata.
- `ConsultarCobrancaImediata`: Consulta uma cobrança imediata.

### pix/pix_cobv_imediata.go

Gerencia cobranças com vencimento via PIX.

#### Funções Principais

- `CriarCobrancaImediataComVencimentoETxID`: Cria uma cobrança com vencimento e TxID.
- `ConsultarCobrancasImediatasComVencimento`: Consulta cobranças com vencimento.

## Requisitos

- Go 1.23
- Dependências externas listadas no `go.mod`.

## Conclusão

Esta biblioteca oferece uma interface robusta e fácil de usar para integrar aplicações Go com a API do Banco Inter, suportando operações bancárias, cobranças e transações PIX de forma segura e eficiente.
