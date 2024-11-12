import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Home from '../Home/Home'
import regWindow from '../reg.sign-up/regWindow'
import Orders from '../Orders/Orders'
import Products from '../Products/Products'
import Clients from '../Clients/Clients'
import Sidebar from '../Sidebar/Sidebar'

const AppRoutes = () => (
	<Routes>
		<Route index element={<Home />} />
		<Route path='/sign-up' element={<regWindow />} />
		{/* <Route path='/lk' element={<Sidebar />} /> */}
		<Route path='/orders' element={<Orders />} />
		<Route path='/clients' element={<Clients />} />
		<Route path='/products' element={<Products />} />
	</Routes>
)

export default AppRoutes
