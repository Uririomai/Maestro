import React from 'react'

import AppRoutes from './Routes/routes'
import Header from './Header/Header'
import RegWindow from './reg.sign-up/regWindow'
import Sidebar from './Sidebar/Sidebar'

import styles from '../styles/scss/style.scss'

const App = () => {
	return (
		<div className={styles.app}>
			<div className={styles.w}>
				<Header />
				<RegWindow />
				<Sidebar />
				<AppRoutes />
			</div>
		</div>
	)
}

export default App
