import React, { useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'

import PENCIL from '../../assets/images/pencil.png'

import styles from './cabinet.module.scss'
import Sidebar from '../Sidebar/Sidebar'
import Button from '../button/button'
import Toggle from '../Toggle/Toggle'

const Cabinet = () => {
	const location = useLocation()
	const [isEditPage, setToggleEdit] = useState(true)

	const navigate = useNavigate()

	const toggleEdit = () => {
		isEditPage ? setToggleEdit(false) : setToggleEdit(true)
	}

	return (
		<>
			<section className={styles.personalCabinet + ` container`}>
				<Sidebar />
				<div className={styles.cabinet}>
					<div className={styles.cabinet__header}>
						<h2 className={styles.cabinet__title}>Контактная информация</h2>
						<button className={styles.cabinet__buttonEdit} onClick={toggleEdit}>
							<img
								className={styles.cabinet__imgButtonEdit}
								src={PENCIL}
								alt='edit profile'
							/>
						</button>
					</div>
					<form className={styles.cabinet__inputList}>
						{isEditPage && (
							<label className={styles.cabinet__label} htmlFor='password'>
								Пароль
								<input
									className={styles.cabinet__input}
									type='password'
									name='password'
									id='password'
									disable
								/>{' '}
								{/* TODO email and phone from STATE */}
							</label>
						)}
						{isEditPage && (
							<label className={styles.cabinet__label} htmlFor='newpassword'>
								Новый пароль
								<input
									className={styles.cabinet__input}
									type='password'
									name='newpassword'
									id='newpassword'
									disable
								/>{' '}
								{/* TODO email and phone from STATE */}
							</label>
						)}

						<label className={styles.cabinet__label} htmlFor='email'>
							Электронная почта
							<input
								className={styles.cabinet__input}
								type='text'
								name='email'
								id='email'
								disable={!isEditPage}
							/>{' '}
							{/* TODO email and phone from STATE */}
						</label>

						<label className={styles.cabinet__label} htmlFor='phone'>
							Телефон
							<input
								className={styles.cabinet__input}
								type='tel'
								name='phone'
								id='phone'
								disable={!isEditPage}
							/>{' '}
							{/* TODO email and phone from STATE */}
						</label>

						<label className={styles.cabinet__label} htmlFor='messager'>
							Мессенджер
							<input
								className={styles.cabinet__input}
								type='text'
								name='messager'
								id='messager'
								placeholder='Добавить мессенджер'
								disable={!isEditPage}
							/>{' '}
							<input
								className={styles.cabinet__input}
								type='text'
								name='messager'
								id='messager'
								placeholder='@example0'
								disable={!isEditPage}
							/>{' '}
							{/* TODO email and phone from STATE */}
						</label>
					</form>
					<div className={styles.cabinet__notification}>
						<h2 className={styles.cabinet__title}>Контактная информация</h2>
						<ul className={styles.cabinet_notificationList}>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}></h3>
									<p className={styles.cabinet__notificationDescript}></p>
								</div>
								<Toggle />
								
							</li>
						</ul>
					</div>
				</div>
			</section>
		</>
	)
}

export default Cabinet
