# Feature Flags with Fiber

Este projeto demonstra como implementar feature flags dinâmicas em um aplicativo Go usando o framework Fiber.

## Estrutura do Projeto

- `main.go`: O ponto de entrada do aplicativo.
- `README.md`: Este arquivo que fornece informações sobre o projeto.

## Como Executar

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/feature-flags-fiber.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd feature-flags-fiber
   ```

3. Execute o aplicativo:

   ```bash
   go run main.go
   ```

   O aplicativo estará disponível em http://localhost:8080.

## Endpoints

- `/seu-endpoint-1`: Um endpoint protegido pela feature flag "exampleFlag1".
- `/seu-endpoint-2`: Um endpoint protegido pela feature flag "exampleFlag2".
- `/feature-status`: Retorna o estado atual de todas as feature flags.

## Atualizar uma Feature Flag

Você pode atualizar o estado de uma feature flag fazendo uma solicitação POST para `/update-feature/:flagName`. Por exemplo:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"enableFeature": true}' http://localhost:8080/update-feature/exampleFlag1
```

Substitua `exampleFlag1` pelo nome da feature flag que você deseja atualizar e ajuste `true` ou `false` conforme necessário.

## Contribuição

Sinta-se à vontade para contribuir para este projeto abrindo issues ou pull requests.

## Licença

Este projeto é licenciado sob a [MIT License](LICENSE).
