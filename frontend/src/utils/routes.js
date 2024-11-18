import Analytics from "../components/Analytics/Analytics"
import Auth from "../components/Auth/Auth"
import Cabinet from "../components/Cabinet/Cabinet"
import Clients from "../components/Clients/Clients"
import Orders from "../components/Orders/Orders"
import Products from "../components/Products/Products"

export const ROUTES = {
	HOME: '/',
	SIGNUP: '/sign-up',
	SIGNIN: '/sign-in',
	CABINET: '/cabinet',
	CABINETEDIT: '/cabinet-edit',
	ORDERS: '/orders',
	PRODUCTS: '/products',
	CLIENTS: '/clients',
	ANALYTICS: '/analytics',
}
export const authRoutes = [
	{
		path: ROUTES.CABINET,
		Component: Cabinet
	},
	
	{
		path: ROUTES.ORDERS,
		Component: Orders
	},
	{
		path: ROUTES.PRODUCTS,
		Component: Products
	},
	{
		path: ROUTES.CLIENTS,
		Component: Clients
	},
	{
		path: ROUTES.ANALYTICS,
		Component: Analytics
	}
]
export const publicRoutes = [
	{
		path: ROUTES.SIGNIN,
		Component: Auth,
	},
	{
		path: ROUTES.SIGNUP,
		Component: Auth,
	},
	{
		path: ROUTES.PRODUCTS + '/:id',
		
	},
	
]
