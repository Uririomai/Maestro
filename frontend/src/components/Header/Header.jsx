import React from 'react'
import { Link } from 'react-router-dom'

import { ROUTES } from '../../utils/routes'

import styles from './header.module.scss'

import LOGO from '../../assets/images/logo.png'

const Header = () => {
	return (
		<header className={styles.header + ` container`}>
			<div className={styles.header__wrapper}>
				<div className={styles.header__logo}>
					<Link to={ROUTES.HOME}>
						<img src={LOGO} alt='logo' />
					</Link>
				</div>
			</div>
		</header>
	)
}

export default Header
