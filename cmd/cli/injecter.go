package cli

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	ability_detail_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/ability_detail/service"
	armor_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/armor/service"
	cache_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/cache/service"
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/service"
	class_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/class/service"
	equipment_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/service"
	level_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/level/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	weapon_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/service"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph"
	"github.com/ballinwza/be-pradit-dnd-2025/middleware"
)

func InjecterAllService() *handler.Server {

	inMemoryCache := cache_service.NewInMemoryCache(15*time.Minute, 16*time.Minute)

	// Inject dependencies
	graphConfig := graph.Config{Resolvers: &graph.Resolver{
		UserService:          user_service.NewUserService(),
		CharacterService:     character_service.NewCharacterService(),
		AbilityDetailService: ability_detail_service.NewAbilityDetailService(),
		ArmorService:         armor_service.NewArmorService(inMemoryCache),
		WeaponService:        weapon_service.NewWeaponService(),
		EquipmentService:     equipment_service.NewEquipmentService(),
		ClassService:         class_service.NewClassService(),
		LevelService:         level_service.NewLevelService(),
		CacheService:         cache_service.NewCacheService(inMemoryCache),
	}}

	// Authentication
	graphConfig.Directives.Auth = graph.Auth

	srv := handler.New(graph.NewExecutableSchema(graphConfig))

	// Websocket
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// อนุญาตการเชื่อมต่อจากทุก Origin (เหมาะสำหรับ Development)
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			// initPayload คือ connectionParams ที่เราส่งมาจาก Client
			authHeader, ok := initPayload["Authorization"].(string)
			if !ok {
				// ไม่มี Authorization ส่งมา
				return ctx, nil, nil // ปล่อยผ่านไปก่อน ให้ resolver จัดการ
			}

			// ตรวจสอบ JWT (อาจจะดึง logic มาจาก http middleware)
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := middleware.VerifyToken(tokenString) // สร้างฟังก์ชัน verifyToken ขึ้นมาใช้ซ้ำ
			if err != nil {
				// Token ไม่ถูกต้อง, ไม่ต้องทำอะไร ปล่อย context ว่าง
				return ctx, nil, nil
			}

			// ถ้า Token ถูกต้อง, แนบ claims ไปกับ context ของ connection นี้
			return context.WithValue(ctx, middleware.UserCtxKey, claims), nil, nil
		},
	})

	// Http
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	// Execution handler
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		gqlErr := graphql.DefaultErrorPresenter(ctx, err)
		if gqlErr.Extensions == nil {
			gqlErr.Extensions = map[string]interface{}{}
		}

		gqlErr.Extensions["timestamp"] = time.Now().Format(time.RFC3339)

		if strings.Contains(gqlErr.Message, "does not exist in") {
			gqlErr.Extensions["friendly_message"] = "Error grqphql ใส่ field ไม่ถูกต้อง"
			gqlErr.Extensions["status_code"] = http.StatusBadRequest
		}

		return gqlErr
	})

	// Other
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
