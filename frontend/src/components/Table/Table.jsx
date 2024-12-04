import React, { useEffect, useState } from 'react'
import Select from '../Select/Select'

import { productEdit, selectSort } from '../../utils/select'
import { useLocation } from 'react-router-dom'

/* const filterRows = (searchText, listData) => {
	if (!searchText) {
		return listData
	}
	console.log(listData.filter((item, index) => {
		
	}))

	return listData.filter(
		row =>
			row
				.keys()
				.map(item => item.toLowerCase().includes(searchText.toLowerCase()))
		row.map(item => item.toLowerCase().includes(searchText.toLowerCase()))
	)
} */

const Table = ({ data, column, styles }) => {
	/* const [listRow, setListRow] = useState(data)
	const [searchTerm, setSearchTerm] = useState('')

	useEffect(() => {
		const debounce = setTimeout(() => {
			const filteredRows = filterRows(searchTerm, data)
			setListRow(filteredRows)
		}, 300)
		return () => clearTimeout(debounce)
	}, [searchTerm]) */

	const location = useLocation()
	const [isProductsPage, setProductsPage] = useState(false)

	useEffect(() => {
		if (location.pathname === '/products') {
			setProductsPage(true)
		} else {
			setProductsPage(false)
		}
	}, [location])

	return (
		<section className={styles.table}>
			<div className={styles.table__head}>
				<div className={styles.table__headLeft}>
					<input
						placeholder='Поиск...'
						className={styles.table__search}
						/* value={searchTerm} */
						/* onChange={e => setSearchTerm(e.target.value)} */
					></input>
					<Select
						styles={styles}
						name={'sort'}
						id={'sort'}
						optionArray={column.map(column => column.heading)}
					/>
				</div>
				{isProductsPage && (
					<div className={styles.table__headRight}>
						<Select
							styles={styles}
							name={'productsSelect'}
							id={'productsSelect'}
							optionArray={productEdit}
						/>
					</div>
				)}
			</div>
			<table className={styles.table__body}>
				<thead>
					<tr>
						{column.map((item, index) => (
							<TableHeadItem key={index} item={item} styles={styles} />
						))}
					</tr>
				</thead>
				<tbody>
					{data.map((item, index) => (
						<TableRow key={index} item={item} column={column} styles={styles} />
					))}
				</tbody>
			</table>
		</section>
	)
}

const TableHeadItem = ({ item, styles }) => (
	<th className={styles.table__header} key={item.id}>
		{item.heading}
	</th>
)
const TableRow = ({ item, column, styles }) => (
	<tr className={styles.table__row}>
		{column.map((columnItem, index) => {
			return <td className={styles.table__cell}>{item[columnItem.value]}</td>
		})}
	</tr>
)

export default Table
