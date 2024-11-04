import React from 'react'
import { NavLink } from 'react-router-dom'

import styles from './sidebar.module.scss'

const Sidebar = () => {
	return (
		<aside class={styles.sidebar}>
			<nav className={styles.nav}>
				<ul className={styles.nav__list}>
					<li className={styles.nav__item}>
						<NavLink to={`/cabinet`}>Кабинет</NavLink>
					</li>
					<li className={styles.nav__item}>
						<NavLink to={`/orders`}>Заказы</NavLink>
					</li>
					<li className={styles.nav__item}>
						<NavLink to={`/products`}>Товары</NavLink>
					</li>
					<li className={styles.nav__item}>
						<NavLink to={`/clients`}>Клиенты</NavLink>
					</li>
					<li className={styles.nav__item}>
						<NavLink to={`/analytics`}>Аналитика</NavLink>
					</li>
				</ul>
			</nav>
		</aside>
	)
}

export default Sidebar
