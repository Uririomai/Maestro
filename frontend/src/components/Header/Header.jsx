import React, { useEffect, useState } from 'react'
import { Link, useLocation, useNavigate } from 'react-router-dom'

import { ROUTES } from '../../utils/routes'

import styles from './header.module.scss'

import LOGO from '../../assets/images/logo.png'
import Button from '../button/button'
import { useDispatch } from 'react-redux'
import { removeUser } from '../../redux/slices/userSlice'

const Header = () => {
	const location = useLocation()
	const [isSignPage, setClassHeader] = useState(false)

	const dispatch = useDispatch()
	const navigate = useNavigate()

	useEffect(() => {
		switch (location.pathname) {
			case '/sign-up':
				setClassHeader(true)
				break
			case '/sign-in':
				setClassHeader(true)
				break
			default:
				setClassHeader(false)
				break
		}
	}, [location.pathname])

	useEffect(() => {}, [isSignPage])

	const logout = () => {
		dispatch(removeUser())

		navigate(ROUTES.SIGNIN) // todo redirect all routes
	}

	return (
		<header className={styles.header + ` container`}>
			<div className={styles.header__wrapper}>
				<div
					className={
						isSignPage ? styles.header__logo : styles.header__logoPosition
					}
				>
					<Link to={ROUTES.SIGNUP}>
						<img src={LOGO} alt='logo' />
					</Link>
				</div>
				{!isSignPage && (
					<Button
						className={styles.header__logout}
						colorBack={'var(--color-black)'}
						width={24}
						height={24}
						onClick={logout}
					/>
				)}
			</div>
		</header>
	)
}

export default Header
