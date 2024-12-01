import React from 'react'

import styles from './input.module.scss'

const Input = (props) => {
  return (
    <>
        <label htmlFor=''>
						<input
							className={styles.auth__field}
							type='password'
							name='password'
							value={password}
							onChange={e => setPassword(e.target.value)}
							placeholder='Пароль'
						/>
					</label>
    </>
  )
}

export default Input