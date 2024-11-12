import React from 'react'

import AppRoutes from '../Routes/routes'
import Header from '../Header/Header'
import RegWindow from '../reg.sign-up/regWindow'
import Sidebar from '../Sidebar/Sidebar'
import Orders from '../Orders/Orders'

import styles from './app.scss'
import Clients from '../Clients/Clients'
import Products from '../Products/Products'

const App = () => {
	return (
		<div className={styles.app}>
			<div className={styles.w}>
				<Header />
				{/* <RegWindow /> */}

				<section className='personal-cabinet container'>
					<Sidebar />
					<AppRoutes />
				</section>
			</div>
		</div>
	)
}

export default App
