import React from 'react'

import AppRoutes from '../Routes/routes'
import Header from '../Header/Header'
import RegWindow from '../reg.sign-up/regWindow'
import Sidebar from '../Sidebar/Sidebar'
import Orders from '../Orders/Orders'

import styles from './app.scss'

const App = () => {
	return (
		<div className={styles.app}>
			<div className={styles.w}>
				<Header />
				{/* <RegWindow /> */}

				<section className='personal-cabinet container'>
					<Sidebar />
					<Orders />
				</section>
				<AppRoutes />
			</div>
		</div>
	)
}

export default App
