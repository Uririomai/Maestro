import React from 'react'

import styles from './button.module.scss'

const Button = props => {
	const colorElements = {
		background: props.colorBack,
		color: props.colorText,
	}

	if (props.imageSrc) {
		return (
			<>
				<button className={styles.button} style={colorElements}>
					<img
						className={styles.button__img}
						src={props.imageSrc}
						alt='img-button'
					/>

					<p className={styles.button__text}>{props.buttonText}</p>
				</button>
			</>
		)
	} else {
		return (
			<>
				<button className={styles.button} style={colorElements}>
					<p className={styles.button__text}>{props.buttonText}</p>
				</button>
			</>
		)
	}
}

export default Button
