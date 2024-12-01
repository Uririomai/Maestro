import React, { useState } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'
import {useSelector} from 'react-redux'

import PENCIL from '../../assets/images/pencil.png'

import styles from './cabinet.module.scss'
import Sidebar from '../Sidebar/Sidebar'
import Button from '../button/button'
import Toggle from '../Toggle/Toggle'
import Select from '../Select/Select'
import { selectMessanger } from '../../utils/select.js'
import { useAuth } from '../../hooks/useAuth.js'

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
						<label className={styles.cabinet__label} htmlFor='email'>
							Электронная почта
							<input
								className={styles.cabinet__input}
								type='text'
								name='email'
								id='email'
								disable={!isEditPage}
								placeholder={useSelector(state => state.user.email)}
							/>{' '}
							{/* TODO email and phone from STATE */}
						</label>

						{isEditPage && (
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
								{/* TODO email and phone from STATE */}
							</label>
						)}

						<label className={styles.cabinet__label} htmlFor='phone'>
							Телефон
							<input
								className={styles.cabinet__input}
								type='tel'
								name='phone'
								id='phone'
								disable={!isEditPage}
								placeholder={useSelector(state => state.user.phone)}
							/>{' '}
							{/* TODO email and phone from STATE */}
						</label>

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

						<label className={styles.cabinet__label} htmlFor='messanger'>
							Мессенджер
							<ul className={styles.cabinet__messangerList}>
								<li className={styles.cabinet__messangerItem}>
									{/* <select
										className={styles.cabinet__input}
										type='select'
										name='messanger'
										id='messanger'
										
									>
										<option disable value='Добавить мессенджер' selected>
											Добавить мессенджер
										</option>
										<option value='Вконтакте'>Вконтакте</option>
										<option value='Телеграмм'>Телеграмм</option>
									</select> */}
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
							{/* TODO email and phone from STATE */}
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
