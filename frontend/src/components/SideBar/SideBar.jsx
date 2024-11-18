import React from 'react'
import { NavLink } from 'react-router-dom'
import { ROUTES } from '../../utils/routes'

import styles from './sidebar.module.scss'
import stylesNav from './nav.module.scss'

const Sidebar = () => {
	return (
		<aside className={styles.sidebar}>
			<nav className={`${styles.sidebar__nav} ${stylesNav.nav}`}>
				<ul className={stylesNav.nav__list}>
					<li className={stylesNav.nav__item}>
						<NavLink active className={stylesNav.nav__link} to={ROUTES.CABINET}>
							Кабинет
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={ROUTES.ORDERS}>
							Заказы
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={ROUTES.PRODUCTS}>
							Товары
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={ROUTES.CLIENTS}>
							Клиенты
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={ROUTES.ANALYTICS}>
							Аналитика
						</NavLink>
					</li>
				</ul>
			</nav>
		</aside>
	)
}

export default Sidebar
