import React, { useEffect, useState } from 'react'
import axios from 'axios'

import Table from '../Table/Table'

import stylesTable from '../Table/table.module.scss'

const Clients = () => {
	const [dataTable, setDataTable] = useState([])

	useEffect(() => {
		axios('https://jsonplaceholder.typicode.com/users')
			.then(response => {
				setDataTable(response.data)
			})
			.catch(error => console.log(`setDataTable`, error))
	}, [])

	const columns = [
		{ heading: 'id', value: 'id' },
		{ heading: 'name', value: 'name' },
		{ heading: 'email', value: 'email' },
		{ heading: 'phone', value: 'phone' },
		{ heading: 'website', value: 'website' },
	]

	/* const columns = [
		{ heading: 'Id', value: 'id' },
		{ heading: 'Сумма', value: 'summ' },
		{ heading: 'Дата', value: 'date' },
		{ heading: 'Клиент', value: 'client' },
		{ heading: 'Статус', value: 'status' },
	] */

	return <Table data={dataTable} column={columns} styles={stylesTable} />
}

export default Clients
