import React, { useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import { useSelector } from 'react-redux'

import PENCIL from '../../assets/images/pencil.png'

import styles from './cabinet.module.scss'
import Sidebar from '../Sidebar/Sidebar'
import Button from '../button/button'
import Toggle from '../Toggle/Toggle'
import Select from '../Select/Select'
import { selectMessanger } from '../../utils/select.js'
import { useAuth } from '../../hooks/useAuth.js'
import Input from '../Input/Input.jsx'

const Cabinet = () => {
	const location = useLocation()
	const [isEditPage, setToggleEdit] = useState(false)

	const navigate = useNavigate()
	const email = useSelector(state => state.user.email)
	const phone = useSelector(state => state.user.phone)

	const toggleEdit = () => {
		isEditPage ? setToggleEdit(false) : setToggleEdit(true)
		console.log(isEditPage)
	}

	return (
		<>
			<section className={styles.personalCabinet + ` container`}>
				<Sidebar />
				<div className={styles.cabinet}>
					<div className={styles.cabinet__header}>
						<h2 className={styles.cabinet__title}>Контактная информация</h2>
						<Button
							className={styles.cabinet__buttonEdit}
							onClick={toggleEdit}
							colorBack={'var(--color-light)'}
							buttonText={'Редактировать'}
							width={'143px'}
							height={'36px'}
						></Button>
					</div>
					<form className={styles.cabinet__inputList}>
						<Input
							type={'email'}
							id={'email'}
							isDisable={!isEditPage}
							placeholder={email}
							label={'Электронная почта'}
						/>

						<Input
							type={'tel'}
							id={'phone'}
							isDisable={!isEditPage}
							placeholder={phone}
							label={'Телефон'}
						/>

						{/* {isEditPage && (
							<label className={styles.cabinet__label} htmlFor='password'>
								Пароль
								<input
									className={styles.cabinet__input}
									type='password'
									name='password'
									id='password'
									disable
									placeholder='************'
								/>{' '}
							
							</label>
						)} */}

						{/* 	{isEditPage && (
							<label className={styles.cabinet__label} htmlFor='newpassword'>
								Новый пароль
								<input
									className={styles.cabinet__input}
									type='password'
									name='newpassword'
									id='newpassword'
									disable
								/>{' '}
								
							</label>
						)} */}

						<label className={styles.cabinet__label} htmlFor='messanger'>
							Мессенджер
							<ul className={styles.cabinet__messangerList}>
								<li className={styles.cabinet__messangerItem}>
									<Select
										styles={styles}
										name={'messanger'}
										id={'messanger'}
										optionArray={selectMessanger}
									/>
								</li>
								<li className={styles.cabinet__messangerItem}>
									<input
										className={styles.cabinet__input}
										type='text'
										name='messanger'
										id='messanger'
										placeholder='@example0'
										disable={!isEditPage}
									/>{' '}
								</li>
							</ul>
						</label>
					</form>
					<div className={styles.cabinet__notification}>
						<h2 className={styles.cabinet__title}>Контактная информация</h2>
						<ul className={styles.cabinet__notificationList}>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}>
										Получать по электронной почте
									</h3>
									<p className={styles.cabinet__notificationDescription}>
										Получать уведомления по электронной почте
									</p>
								</div>
								<Toggle id={'not1'} />
							</li>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}>
										Получать через мессенджеры
									</h3>
									<p className={styles.cabinet__notificationDescription}>
										Получать уведомления через мессенджеры
									</p>
								</div>
								<Toggle id={'not2'} />
							</li>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}>
										Изменения статуса заказа
									</h3>
									<p className={styles.cabinet__notificationDescription}>
										Уведомить меня о изменении статуса заказа
									</p>
								</div>
								<Toggle id={'not3'} />
							</li>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}>
										Новый заказ
									</h3>
									<p className={styles.cabinet__notificationDescription}>
										Уведомить меня о новом размещенном заказе
									</p>
								</div>
								<Toggle id={'not4'} />
							</li>
							<li className={styles.cabinet__notificationItem}>
								<div className={styles.cabinet__notificationTextWrapper}>
									<h3 className={styles.cabinet__notificationTitle}>
										Новый возврат
									</h3>
									<p className={styles.cabinet__notificationDescription}>
										Уведомить меня о новом поступившем возврате
									</p>
								</div>
								<Toggle id={'not5'} />
							</li>
						</ul>
					</div>
				</div>
			</section>
		</>
	)
}

export default Cabinet
