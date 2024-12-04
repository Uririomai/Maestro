import React from 'react'

import styles from './select.module.scss'

const Select = ({ styles, name, id, optionArray = [] }) => {
	return (
		<select
			className={styles.select}
			type='select'
			name={name}
			id={id}
		>
			{optionArray.map(item => (
				<option className={styles.select__item} value={item}>
					{item}
				</option>
			))}
		</select>
	)
}

export default Select
