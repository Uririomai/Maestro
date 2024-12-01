import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { ROUTES } from '../../utils/routes'
import { Link, useLocation, useNavigate } from 'react-router-dom'

import { useAuth } from '../../hooks/useAuth'

import styles from './auth.module.scss'

import GOOGLE_IC from '../../assets/images/reg.sign-up/Google.svg'

import Button from '../button/button'
import { login, registration } from '../../redux/slices/userSlice'

/* import { registration, login } from '../../http/userAPI'
import { setUser } from '../../redux/slices/userSlice' */

const Auth = () => {
	const location = useLocation()
	const isLoginPage = location.pathname === ROUTES.SIGNIN

	const [email, setEmail] = useState('')
	const [password, setPassword] = useState('')
	const navigate = useNavigate()

	/* const { isAuth } = useAuth() */

	const { isAuth } = useSelector(state => state.user)
	const dispatch = useDispatch()

	const click = async () => {
		try {
			if (isLoginPage) {
				dispatch(login({ email, password }))
			} else {
				/* console.log(isAuth) */
				dispatch(registration({ email, password }))
				/* console.log(isAuth) */
			}
		} catch (e) {
			alert(e)
		}
	}

	

	useEffect(() => {
		console.log('pre redirect')
		if (isAuth) {
			console.log('redirect')
			navigate(ROUTES.CABINET)
		} else {
			navigate(ROUTES.SIGNIN)
		}
	}, [isAuth])

	return (
		<section className={styles.auth + ` container`}>
			<div className={styles.auth__wrapper}>
				<div>
					<h1 className={styles.auth__title}>
						{isLoginPage ? 'Добро пожаловать обратно' : 'Впервые здесь?'}
					</h1>
					<h2 className={styles.auth__subtitle}>
						{isLoginPage
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
					{!isLoginPage && (
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
						buttonText={isLoginPage ? 'Войти' : 'Зарегестрироваться'}
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
					{isLoginPage ? (
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
