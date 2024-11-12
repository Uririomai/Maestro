import React, { useEffect, useState } from 'react'

import axios from 'axios'

import Table from '../Table/Table'

const Products = () => {
	const [dataTable, setDataTable] = useState([])

	useEffect(() => {
		axios('https://jsonplaceholder.typicode.com/users')
			.then(response => {
				setDataTable(response.data)
			})
			.catch(error => console.log(`setDataTable`, error))
	}, [])

	const columns = [
		{ heading: 'Товар', value: 'name' },
		{ heading: 'id', value: 'id' },
		{ heading: 'Количество', value: 'username' },
		{ heading: 'Цена', value: 'phone' },
		{ heading: 'Статус', value: 'website' },
	]

	/* const columns = [
		{ heading: 'Id', value: 'id' },
		{ heading: 'Сумма', value: 'summ' },
		{ heading: 'Дата', value: 'date' },
		{ heading: 'Клиент', value: 'client' },
		{ heading: 'Статус', value: 'status' },
	] */

	return <Table data={dataTable} column={columns} />
}

export default Products
