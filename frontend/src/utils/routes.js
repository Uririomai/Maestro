import Analytics from "../components/Analytics/Analytics"
import Auth from "../components/Auth/Auth"
import Cabinet from "../components/Cabinet/Cabinet"
import Clients from "../components/Clients/Clients"
import MenuPage from "../components/MenuPage/MenuPage"
import Orders from "../components/Orders/Orders"
import PagesSite from "../components/PagesSite/PagesSite"
import Products from "../components/Products/Products"
import ThemePage from "../components/ThemePage/ThemePage"

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
	THEME: '/theme',
	PAGESSITE: '/pages',
	MENU: '/menu',
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
	},
	{
		path: ROUTES.THEME,
		Component: ThemePage
	},
	{
		path: ROUTES.PAGESSITE,
		Component: PagesSite
	},
	{
		path: ROUTES.MENU,
		Component: MenuPage
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
