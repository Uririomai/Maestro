import React from 'react'

import styles from './toggle.module.scss'

const Toggle = () => {
	return (
		<>
			<label htmlFor='qwe' className={styles.toggle}>
				<input type='checkbox' id='qwe' className={styles.toggle__input} />
				<span className={styles.toggle__slider}></span>
			</label>
		</>
	)
}

export default Toggle
