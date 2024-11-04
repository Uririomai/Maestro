import React from 'react'

import styles from './regWindow.module.scss'

import Button from '../button/button'

import GOOGLE_IC from '../../assets/images/reg.sign-up/Google.svg'

const regWindow = () => {
	return (
		<section className={styles.regWindow + ` container`}>
			<div className={styles.regWindow__wrapper}>
				<div>
					<h1 className={styles.regWindow__title}>Добро пожаловать обратно</h1>
					<h2 className={styles.regWindow__subtitle}>
						Пожалуйста, войдите в свою учетную запись
					</h2>
				</div>
				<form className={styles.regWindow__form}>
					<label htmlFor=''>
						<input
							className={styles.regWindow__field}
							type='email'
							name='email'
							placeholder='Электронная почта'
						/>
					</label>
					<label htmlFor=''>
						<input
							className={styles.regWindow__field}
							type='password'
							name='password'
							placeholder='Пароль'
						/>
					</label>
					<Button
						buttonText={'Войти'}
						colorBack={'var(--color-black)'}
						colorText={'var(--color-light)'}
					/>
				</form>
				<p className={styles.regWindow__continueWith}>или продолжить с</p>
				<Button imageSrc={GOOGLE_IC} buttonText={'Google'} />

				<label htmlFor='member-me' className={styles.regWindow__checkbox}>
					<input type='checkbox' id='member-me' />
					<p>Запомнить меня</p>
				</label>
			</div>
		</section>
	)
}

export default regWindow
