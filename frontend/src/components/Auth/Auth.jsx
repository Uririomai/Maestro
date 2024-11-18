import React, { useState } from 'react'
import { ROUTES } from '../../utils/routes'
import { Link, useLocation, useNavigate } from 'react-router-dom'

import styles from './auth.module.scss'

import GOOGLE_IC from '../../assets/images/reg.sign-up/Google.svg'

import Button from '../button/button'

import { registration, login } from '../../http/userAPI'

const Auth = () => {
	const location = useLocation()
	const isLogin = location.pathname === ROUTES.SIGNIN

	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const navigate = useNavigate()

	const click = async () => {
		try {
			if (isLogin) {
				const response = await login(email, password)
				console.log(response)
			} else {
				const response = await registration(email, password)
				console.log(response)
			}
			navigate(ROUTES.CABINET)
		} catch (e) {}
	}

	return (
		<section className={styles.auth + ` container`}>
			<div className={styles.auth__wrapper}>
				<div>
					<h1 className={styles.auth__title}>
						{isLogin ? 'Добро пожаловать обратно' : 'Впервые здесь?'}
					</h1>
					<h2 className={styles.auth__subtitle}>
						{isLogin
							? 'Пожалуйста, войдите в свою учетную запись'
							: 'Создайте учетную запись, чтобы начать'}
					</h2>
				</div>
				<form className={styles.auth__form}>
					<label htmlFor=''>
						<input
							className={styles.auth__field}
							type='email'
							name='email'
							value={email}
							onChange={e => setEmail(e.target.value)}
							placeholder='Электронная почта'
						/>
					</label>
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
					{!isLogin && (
						<label htmlFor=''>
							<input
								className={styles.auth__field}
								type='password'
								name='password'
								value={password}
								onChange={e => setPassword(e.target.value)}
								placeholder='Повторите пароль'
							/>
						</label>
					)}
					<Button
						buttonText={isLogin ? 'Войти' : 'Зарегестрироваться'}
						colorBack={'var(--color-black)'}
						colorText={'var(--color-light)'}
						onClick={click}
					/>
				</form>
				<p className={styles.auth__continueWith}>или продолжить с</p>
				<Button imageSrc={GOOGLE_IC} buttonText={'Google'} />

				<label htmlFor='member-me' className={styles.auth__checkbox}>
					<input type='checkbox' id='member-me' />
					<p>Запомнить меня</p>
				</label>

				<div className={styles.auth__text}>
					{isLogin ? (
						<div>
							Нет аккаунта?{' '}
							<Link to={ROUTES.SIGNUP} className={styles.auth__link}>
								Зарегистрируйтесь!
							</Link>
						</div>
					) : (
						<div>
							Есть аккаунт?{' '}
							<Link to={ROUTES.SIGNIN} className={styles.auth__link}>
								Войдите!
							</Link>
						</div>
					)}
				</div>
			</div>
		</section>
	)
}

export default Auth
