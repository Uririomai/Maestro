import React, { useEffect, useState } from 'react'
import axios from 'axios'

import Table from '../Table/Table'
import Graph from '../Graphic/Graphic'

import styles from './analytics.module.scss'
import stylesAnalytics from '../Table/tableAnalytics.module.scss'

const Analytics = () => {
	const [dataTable, setDataTable] = useState([])

	useEffect(() => {
		axios
			.get('https://jsonplaceholder.typicode.com/users')
			.then(response => {
				setDataTable(response.data)
			})
			.catch(error => console.log('setDataTable', error))
	}, [])

	const columns = [
		{ heading: 'Продукт', value: 'name' },
		{ heading: 'Продажи', value: 'id' },
		{ heading: 'Возвраты', value: 'username' },
		{ heading: 'Рейтинг', value: 'phone' },
	]

	return (
		<>
			<div className={styles.analytics}>
				<ul className={styles.analytics__graphsList}>
					<li className={styles.analytics__graphCard}>
						<h3 className={styles.analytics__titleCard}>Общая прибыль</h3>
						<p className={styles.analytics__number}>
							$120,000 <span className={styles.analytics__percent}>+5%</span>
						</p>
						<div className={styles.analytics__graphWrapper}>
							<Graph />
						</div>
					</li>
					<li className={styles.analytics__graphCard}>
						<h3 className={styles.analytics__titleCard}>Заказы</h3>
						<p className={styles.analytics__number}>
							1,500 <span className={styles.analytics__percent}>+3%</span>
						</p>
						<div className={styles.analytics__graphWrapper}>
							<Graph />
						</div>
					</li>
				</ul>
				<h2 className={styles.analytics__subtitle}>Ежедневная метрика</h2>
				<ul className={styles.analytics__dayCardList}>
					<li className={styles.analytics__dayCard}>
						<h3 className={styles.analytics__titleCard}>Посетители</h3>
						<p className={styles.analytics__number}>10,000</p>
						<p className={styles.analytics__percent}>-8%</p>
					</li>
					<li className={styles.analytics__dayCard}>
						<h3 className={styles.analytics__titleCard}>Возвраты</h3>
						<p className={styles.analytics__number}>200</p>
						<p className={styles.analytics__percent}>-7%</p>
					</li>
					<li className={styles.analytics__dayCard}>
						<h3 className={styles.analytics__titleCard}>Затраты</h3>
						<p className={styles.analytics__number}>$50,000</p>
						<p className={styles.analytics__percent}>+2%</p>
					</li>
					<li className={styles.analytics__dayCard}>
						<h3 className={styles.analytics__titleCard}>Средний чек</h3>
						<p className={styles.analytics__number}>$470</p>
						<p className={styles.analytics__percent}>+1%</p>
					</li>
				</ul>
				<div className={styles.analytics__table}>
					<Table data={dataTable} column={columns} styles={stylesAnalytics} />
				</div>
			</div>
		</>
	)
}

export default Analytics
