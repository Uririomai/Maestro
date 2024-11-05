import React from 'react'

import styles from './table.module.scss'

const Table = ({ data, column }) => {
	return (
		<section className={styles.table}>
			<div className={styles.table__filters}>
				<div className={styles.table__search}>Поиск заказов...</div>
				<div className={styles.table__sort}>
					Сортировать по <span className={styles.table__sortarrow}> </span>
				</div>
			</div>
			<table className={styles.table__body}>
				<tr>
					{column.map((item, index) => (
						<TableHeadItem item={item} />
					))}
				</tr>
				{data.map((item, index) => (
					<TableRow item={item} column={column} />
				))}
			</table>
		</section>
	)
}

const TableHeadItem = ({ item }) => (
	<th className={styles.table__header}>{item.heading}</th>
)
const TableRow = ({ item, column }) => (
	<tr className={styles.table__row}>
		{column.map((columnItem, index) => {
			return <td>{item[columnItem.value]}</td>
		})}
	</tr>
)

export default Table
