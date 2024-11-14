import React from 'react'
import { NavLink } from 'react-router-dom'

import styles from './sidebar.module.scss'
import stylesNav from './nav.module.scss'

const Sidebar = () => {
	return (
		<aside className={styles.sidebar}>
			<nav className={`${styles.sidebar__nav} ${stylesNav.nav}`}>
				<ul className={stylesNav.nav__list}>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={`/cabinet`}>
							Кабинет
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={`/orders`}>
							Заказы
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={`/products`}>
							Товары
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={`/clients`}>
							Клиенты
						</NavLink>
					</li>
					<li className={stylesNav.nav__item}>
						<NavLink className={stylesNav.nav__link} to={`/analytics`}>
							Аналитика
						</NavLink>
					</li>
				</ul>
			</nav>
		</aside>
	)
}

export default Sidebar
