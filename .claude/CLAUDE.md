# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**CPay Payment Gateway Go SDK** is a Go library for integrating with CPay's payment gateway services, supporting both cryptocurrency (USDT, etc.) and credit card payments.

## Commands

### Build and Test
```bash
# Run all tests
go test ./...

# Run tests for a specific package
go test ./sdk
go test ./service/v1

# Run a specific test
go test -v ./sdk -run TestGenSign

# Run example tests (integration tests requiring credentials)
go test -v ./example
```

### Module Management
```bash
# Download dependencies
go mod download

# Tidy dependencies
go mod tidy

# Verify dependencies
go mod verify
```

## Architecture

### Layered Design

The SDK follows a clean layered architecture:

```
service/v1/          # API layer - Payment operations
    ↓
sdk/client.go        # HTTP client layer
    ↓
sdk/signer.go        # Signature generation
sdk/parameter_builder.go  # Request encoding (GET vs POST)
sdk/response/        # Response parsing
```

### Client Creation Flow

Two ways to create a client:

1. **With secret directly** (common):
   ```go
   client := service.NewClientWithSecret(merchantId, secretKey)
   client.WithConfig(sdk.NewDefaultSandBoxEnv())
   ```

2. **With explicit credential object**:
   ```go
   client := service.NewClient(config, credential)
   ```

### Request/Response Pattern

All API methods follow this pattern:

1. **Request struct** - Contains request parameters + `Sign` field (auto-filled)
2. **Response struct** - Contains response data (without wrapper)
3. **Method** - Wraps everything in `BaseRequest`/`BaseResponse`

```go
// Example from create_crypto_order.go
func (c *Client) CreateCryptoOrder(req *CreateCryptoOrderRequest) (*CreateCryptoOrderResponse, error) {
    baseReq := &request.BaseRequest{
        URL:     createOrderURI,      // e.g., "/createOrder"
        Method:  http.MethodPost,
        Version: sdk.VersionV1,       // "v1"
        Param:   req,
    }
    baseResp := response.NewBaseResponse(&CreateCryptoOrderResponse{})
    // c.Send() handles signing, HTTP call, parsing
    err := c.Send(baseReq, baseResp)
    return baseResp.Data.(*CreateCryptoOrderResponse), err
}
```

### Signature Generation (sdk/signer.go)

The signature is HMAC-SHA256 with these rules:

1. **Parameter processing**:
   - Remove `sign` parameter if present
   - Sort keys alphabetically
   - Filter out empty/whitespace-only values
   - Append `key={secretKey}` at the end

2. **String format**: `key1=value1&key2=value2&...&key={secretKey}`

3. **Security validations**:
   - Secret key must be ≥32 characters
   - Parameter keys validated (letters, digits, `_`, `-`, `.` only)
   - Max 1000 parameters (DoS protection)

4. **Type support**: string, int, int64, float32, float64, bool

### Parameter Builders (sdk/parameter_builder.go)

Two builders based on HTTP method:

- **WithoutBodyBuilder** (GET/DELETE/HEAD): Parameters in URL query string
- **WithURLBodyBuilder** (POST/PUT/etc): Parameters in form-encoded body

### Available APIs (service/v1/)

| Method | URI | HTTP Method | Description |
|--------|-----|-------------|-------------|
| `CreateCryptoOrder` | `/createOrder` | POST | Create crypto payment order |
| `CreateCreditCardOrder` | `/createOrderByCreditCard` | POST | Create credit card payment order |
| `GetPaymentOrder` | `/getOrderDetail` | GET | Query order status |
| `GetWalletAddress` | `/getWalletAddress` | GET | Get wallet addresses for deposits |

### Config Endpoints

Default endpoints (config.go):
- **Sandbox**: `test-api.cpay.ltd` (HTTP)
- **Production**: `api.cpay.ltd` (HTTPS)

Custom endpoint via `sdk.NewConfig().WithEndpoint(host:port)`

### Request Struct Conventions

- All request fields use `mapstructure` tag for struct-to-map conversion
- `Sign` field is auto-populated, never set manually
- Use `CreateTime int64` with `time.Now().UnixMilli()` for timestamps
- Amounts are `string` type (decimal precision)

### Response Parsing (sdk/response/)

All responses follow the CPay API wrapper format:
```json
{
  "code": 0,        // 0 = success, >0 = error
  "msg": "message",
  "data": {...},    // Unwrapped into response struct
  "traceid": "..."
}
```

Errors are returned as `*errors.CPaySDKError` with Code/Message/TraceId.

## Adding New APIs

To add a new payment API:

1. Create request/response structs in `service/v1/{api_name}.go`
2. Add `const {api}URI = "/path"`
3. Add method on `*Client`:
   ```go
   func (c *Client) NewAPI(req *NewAPIRequest) (*NewAPIResponse, error) {
       baseReq := &request.BaseRequest{
           URL:     newAPIURI,
           Method:  http.MethodPost,  // or GET
           Version: sdk.VersionV1,
           Param:   req,
       }
       baseResp := response.NewBaseResponse(&NewAPIResponse{})
       // Data extraction pattern
       resp := baseResp.Data.(*NewAPIResponse)
       err := c.Send(baseReq, baseResp)
       return resp, err
   }
   ```

## Git与版本控制 (Git & Version Control)

- **Commit Message规范**: **[严格遵循]** Conventional Commits 规范 (https://www.conventionalcommits.org/)。
  - 格式: `<type>(<scope>): <subject>`
  - 当被要求生成commit message时，必须遵循此格式。
