# SparkProctoringProto

Общие типы данных и API-контракты системы SparkProctoring.

Этот модуль импортируется **SparkProctoringAgent** (клиент), **SparkProctoringServer** (сервер) и другими компонентами для обеспечения типобезопасного обмена данными.

## Установка

```bash
go get github.com/SparkGuard/SparkProctoringProto
```

## Структура

```
SparkProctoringProto/
├── proto/                                    # Исходные .proto файлы
│   └── sparkproctoring/v1/
│       ├── common.proto                      # Общие типы: EventPayload, SessionInfo, ChunkMeta, ...
│       └── agent_service.proto               # gRPC-сервис AgentService (Agent → Server)
├── gen/                                      # Сгенерированный Go-код (НЕ РЕДАКТИРОВАТЬ ВРУЧНУЮ)
│   └── sparkproctoring/v1/
│       ├── common.pb.go
│       ├── agent_service.pb.go
│       └── agent_service_grpc.pb.go
├── pkg/proto/                                # Ручные Go-константы (для REST API и утилит)
│   ├── api.go                                # REST эндпоинты, request/response типы, коды ошибок
│   ├── jwt.go                                # JWT-константы и структура клеймов
│   └── types.go                              # Байтовые типы блоков, строковые статусы, роли
├── buf.yaml                                  # Конфигурация buf (линтинг, breaking changes)
├── buf.gen.yaml                              # Конфигурация кодогенерации (локальные плагины)
├── Makefile                                  # Команды: generate, deps, clean, lint
├── go.mod
└── go.sum
```

## Архитектура: два транспорта

Система использует два транспортных протокола, каждый со своим набором типов:

| Транспорт | Направление | Формат | Назначение |
|---|---|---|---|
| **gRPC** | Agent → Server | Protobuf | Телеметрия, видеочанки, heartbeat |
| **REST** | Web UI → Server | JSON | Авторизация преподавателей, просмотр сессий |

### Разделение `gen/` и `pkg/`

`.proto` файлы — **единственный источник правды** для типов данных:

| Нужно | Откуда импортировать | Пример |
|---|---|---|
| gRPC-клиент/сервер, protobuf-сообщения | `gen/sparkproctoring/v1` | `pb.AuthSessionRequest{}` |
| REST-эндпоинты, error-коды | `pkg/proto` | `proto.EndpointSessions` |
| Байтовые типы блоков Storage | `pkg/proto` | `proto.TypeVideo` |
| JWT-константы, TTL | `pkg/proto` | `proto.ClaimSessionID` |
| Строковые статусы, роли (REST JSON) | `pkg/proto` | `proto.StatusActive` |

## gRPC-сервис AgentService

Определён в `proto/sparkproctoring/v1/agent_service.proto`:

| Метод | Тип | Назначение |
|---|---|---|
| `AuthSession` | Unary | Аутентификация агента, получение JWT |
| `SendTelemetry` | Unary | Отправка батча событий |
| `UploadChunk` | Client stream | Загрузка видеочанка порциями |
| `StreamEvents` | Client stream | Live-стриминг событий |
| `Heartbeat` | Unary | Пульс агента |
| `EndSession` | Unary | Завершение сессии |

## Кодогенерация

Используется **buf** с локальными плагинами:

```bash
# Установить инструменты (один раз)
make deps

# Сгенерировать Go-код из .proto
make generate

# Линтинг .proto файлов
make lint

# Удалить сгенерированный код
make clean
```

Сгенерированные файлы коммитятся в репозиторий (папка `gen/`), чтобы другие модули могли импортировать без установки buf.

## Использование

### gRPC-клиент (Agent → Server)

```go
import (
    pb "github.com/SparkGuard/SparkProctoringProto/gen/sparkproctoring/v1"
    "google.golang.org/grpc"
)

conn, _ := grpc.Dial("server:50051", grpc.WithTransportCredentials(creds))
client := pb.NewAgentServiceClient(conn)

// Аутентификация
resp, _ := client.AuthSession(ctx, &pb.AuthSessionRequest{
    SessionKey:   "exam-token-123",
    AgentVersion: "1.0.0",
    Os:           "darwin",
})

// Отправка телеметрии
client.SendTelemetry(ctx, &pb.SendTelemetryRequest{
    SessionId: resp.SessionId,
    Events: []*pb.EventPayload{
        {
            EventType:   pb.EventType_EVENT_TYPE_KEYBOARD,
            TimestampMs: time.Now().UnixMilli(),
            PayloadJson: keyEventJSON,
        },
    },
})
```

### REST-константы (Web UI → Server)

```go
import "github.com/SparkGuard/SparkProctoringProto/pkg/proto"

// Маршрутизация
router.Post(proto.APIPrefix+proto.EndpointAuthLogin, handler.Login)
router.Get(proto.APIPrefix+proto.EndpointSessions, handler.ListSessions)

// Обработка логина
var req proto.LoginRequest
json.NewDecoder(r.Body).Decode(&req)

// Формирование ответа
resp := proto.LoginResponse{
    Token:     token,
    UserID:    user.ID,
    Role:      proto.RoleTeacher,
    ExpiresAt: expiry.Unix(),
}
```

### Константы для Storage

```go
import "github.com/SparkGuard/SparkProctoringProto/pkg/proto"

// Запись блока в локальное хранилище
storage.Write(proto.TypeVideo, chunkBytes)
storage.Write(proto.TypeEvent, eventJSON)

// JWT-клеймы
claims := proto.JWTClaims{
    SessionID: sessionID,
    UserID:    userID,
    Role:      proto.RoleStudent,
}
```

## Зависимости

- `google.golang.org/grpc` — gRPC фреймворк
- `google.golang.org/protobuf` — Protobuf runtime
- `buf.build/protocolbuffers/wellknowntypes` — стандартные protobuf-типы (google.protobuf.Timestamp и др.)
