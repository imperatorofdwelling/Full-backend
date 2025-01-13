package role

// Route Определяем структуру для маршрута и его методов
type Route struct {
	Path    string
	Methods []string
}

// Константы для ролей
const (
	TenantRole   = "tenant"
	LandlordRole = "landlord"
)

// TenantRoutes Константа для маршрутов роли Tenant
var TenantRoutes = []Route{
	{Path: "/advantages", Methods: []string{"GET", "POST"}},
	{Path: "/registration", Methods: []string{"POST"}},
	{Path: "/login", Methods: []string{"POST"}},
	{Path: "/contract", Methods: []string{"GET", "POST"}},
	{Path: "/favourites", Methods: []string{"GET", "POST", "DELETE"}},
	{Path: "/locations", Methods: []string{"GET"}},
	{Path: "/reservation", Methods: []string{"GET"}},
	{Path: "/history", Methods: []string{"GET", "POST"}},
	{Path: "/stays", Methods: []string{"GET"}},
	{Path: "/staysadvantage", Methods: []string{"GET"}},
	{Path: "/report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/staysreviews", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/user", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "/user/report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
}

// LandlordRoutes Константа для маршрутов роли Landlord
var LandlordRoutes = []Route{
	{Path: "/advantages", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/registration", Methods: []string{"POST"}},
	{Path: "/login", Methods: []string{"POST"}},
	{Path: "/contract", Methods: []string{"GET", "PUT", "POST"}},
	{Path: "/favourites", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/locations", Methods: []string{"GET", "DELETE"}},
	{Path: "/reservation", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/history", Methods: []string{"GET", "POST"}},
	{Path: "/stays", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/staysadvantage", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/staysreviews", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
	{Path: "/user", Methods: []string{"GET", "PUT", "DELETE"}},
	{Path: "/user/report", Methods: []string{"GET", "POST", "PUT", "DELETE"}},
}
