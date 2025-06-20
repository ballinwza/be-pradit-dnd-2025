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
	character_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/character/service"
	equipment_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/equipment/service"
	user_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/user/service"
	weapon_service "github.com/ballinwza/be-pradit-dnd-2025/internal/features/weapon/service"
	"github.com/ballinwza/be-pradit-dnd-2025/internal/graph"
)

func InjecterAllService() *handler.Server {
	userService := user_service.NewUserService()
	characterService := character_service.NewCharacterService()
	abilityDetailService := ability_detail_service.NewAbilityDetailService()
	armorService := armor_service.NewArmorService()
	weaponService := weapon_service.NewWeaponService()
	equipmentService := equipment_service.NewEquipmentService()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserService:          userService,
		CharacterService:     characterService,
		AbilityDetailService: abilityDetailService,
		ArmorService:         armorService,
		WeaponService:        weaponService,
		EquipmentService:     equipmentService,
	}}))

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
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

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

	return srv
}
