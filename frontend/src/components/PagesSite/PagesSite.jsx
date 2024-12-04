import React, { useState, useEffect } from 'react'
import axios from 'axios'

import Table from '../Table/Table'
import Sidebar from '../Sidebar/Sidebar'

import stylesTable from '../Table/table.module.scss'
import cabinetStyles from '../Cabinet/cabinet.module.scss'
import styles from './pagesSite.module.scss'

const PagesSite = () => {
	const [dataTable, setDataTable] = useState([])

	useEffect(() => {
		axios('https://jsonplaceholder.typicode.com/users')
			.then(response => {
				setDataTable(response.data)
			})
			.catch(error => console.log(`setDataTable`, error))
	}, [])

	const columns = [
		{ heading: 'Cтраница', value: 'id' },
		{ heading: 'Отображение', value: 'name' },
		{ heading: 'Статус', value: 'username' },
		
	]

	return (
		<>
			<section className={cabinetStyles.personalCabinet + ` container`}>
				<Sidebar />
				<Table data={dataTable} column={columns} styles={stylesTable} />
			</section>
		</>
	)
}

export default PagesSite
