import React from 'react'

const Table = ({ data, column, styles }) => {
	return (
		<section className={styles.table}>
			<div className={styles.table__filters}>
				<input placeholder='Поиск...' className={styles.table__search}></input>
				<div className={styles.table__sort}>
					Сортировать по <span className={styles.table__sortarrow}> </span>
				</div>
			</div>
			<table className={styles.table__body}>
				<thead>
					<tr>
						{column.map((item, index) => (
							<TableHeadItem item={item} styles={styles} />
						))}
					</tr>
				</thead>
				<tbody>
					{data.map((item, index) => (
						<TableRow item={item} column={column} styles={styles} />
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
