package role

// Route Определяем структуру для маршрута и его методов
type Route struct {
	Path    string
	Methods []string
}

// Константы для ролей
const (
	TenantRole   = 0
	LandlordRole = 1
)

// TenantRoutes Константа для маршрутов роли Tenant
var TenantRoutes = []Route{
	{Path: "advantages", Methods: []string{"GET", "POST"}},
	{Path: "chat", Methods: []string{"GET", "POST"}},
	{Path: "registration", Methods: []string{"POST"}},
	{Path: "login", Methods: []string{"POST"}},
	{Path: "email", Methods: []string{"GET"}},
	{Path: "contract", Methods: []string{"GET", "POST"}},
	{Path: "favourites", Methods: []string{"GET", "POST", "DELETE"}},
	{Path: "history", Methods: []string{"GET", "POST"}},
	{Path: "locations", Methods: []string{"GET"}},
	{Path: "message", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "reservation", Methods: []string{"GET"}},
	{Path: "stays", Methods: []string{"GET"}},
	{Path: "staysadvantage", Methods: []string{"GET"}},
	{Path: "staysreviews", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "user", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "user/report", Methods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"}},
}

// LandlordRoutes Константа для маршрутов роли Landlord
var LandlordRoutes = []Route{
	{Path: "advantages", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "chat", Methods: []string{"GET", "POST"}},
	{Path: "registration", Methods: []string{"POST"}},
	{Path: "login", Methods: []string{"POST"}},
	{Path: "email", Methods: []string{"GET"}},
	{Path: "contract", Methods: []string{"GET", "PUT", "POST"}},
	{Path: "favourites", Methods: []string{"GET", "POST", "DELETE"}},
	{Path: "history", Methods: []string{"GET", "POST"}},
	{Path: "locations", Methods: []string{"GET", "DELETE"}},
	{Path: "message", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "reservation", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "stays", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "staysadvantage", Methods: []string{"GET", "DELETE"}},
	{Path: "staysreviews", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "user", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "user/report", Methods: []string{"GET", "POST", "PUT", "DELETE", "PATCH"}},
}
