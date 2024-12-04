import React from 'react'

import styles from './toggle.module.scss'

const Toggle = (props) => {
	return (
		<>
			<label htmlFor={props.id} className={styles.toggle}>
				<input type='checkbox' id={props.id} className={styles.toggle__input} />
				<span className={styles.toggle__slider}></span>
			</label>
		</>
	)
}

export default Toggle
