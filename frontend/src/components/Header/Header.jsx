import React, { useEffect, useState } from 'react'
import { Link, useLocation } from 'react-router-dom'

import { ROUTES } from '../../utils/routes'

import styles from './header.module.scss'

import LOGO from '../../assets/images/logo.png'

const Header = () => {
	const location = useLocation()
	const [logoPositionInHeader, setLogoPosition] = useState(styles.header)

	useEffect(() => {
		switch (location.pathname) {
			case '/sign-up':
				setLogoPosition(styles.headerLeft)
				break;
			default:
				setLogoPosition(styles.header)
				break
		}
	}, [location.pathname])

	return (
		<header className={styles.header + ` ` + logoPositionInHeader + ` container` }>
			<div className={styles.header__wrapper}>
				<div className={styles.header__logo}>
					<Link to={ROUTES.SIGNUP}>
						<img src={LOGO} alt='logo' />
					</Link>
				</div>
			</div>
		</header>
	)
}

export default Header
